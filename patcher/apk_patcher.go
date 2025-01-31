package patcher

import (
	"elichika/config"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

var (
	appDir  = ""
	tempDir = "temp"
	ksPath  = "assets/askey.keystore"
	ksPass  = "123456\n"
	apkFile = flag.String("apk", "", "apk file")
)

func ApkPatcher() {
	if *apkFile != "" {
		// 初始化
		if err := isCommandAvailable("apktool", "zipalign", "apksigner"); err != nil {
			panic(err)
		}

		apkPath, err := filepath.Abs(*apkFile)
		if err != nil {
			panic(err)
		}
		ksPath, err = filepath.Abs(ksPath)
		if err != nil {
			panic(err)
		}

		// 保存程序主目录
		orig, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		appDir = filepath.ToSlash(orig)

		// 切换到临时目录
		if _, err = os.Stat(tempDir); err != nil {
			if err = os.Mkdir(tempDir, 0755); err != nil {
				panic(err)
			}
		}
		if err = os.Chdir(tempDir); err != nil {
			panic(err)
		}

		// 获取 apk 文件名
		_, apkName := filepath.Split(apkPath)
		fileExt := filepath.Ext(apkName)
		if fileExt != ".apk" {
			log.Println("apk file must end with .apk")
			return
		}
		FileName := apkName[:len(apkName)-len(fileExt)]

		// 反编译 apk
		cmd := exec.Command("apktool", "-r", "-f", "d", apkPath)
		out, err := cmd.CombinedOutput()
		if err != nil {
			panic(err)
		}
		fmt.Println(string(out))

		// 修改 global-metadata.dat
		inFile := FileName + "/assets/bin/Data/Managed/Metadata/global-metadata.dat"
		outFile := "outfile.dat"
		if err = MetadataPatcher(inFile, outFile, config.Conf.Patcher); err != nil {
			panic(err)
		}

		// 覆盖 global-metadata.dat
		if err = moveFile(outFile, inFile); err != nil {
			panic(err)
		}

		// 重新打包 apk
		cmd = exec.Command("apktool", "b", FileName)
		out, err = cmd.CombinedOutput()
		if err != nil {
			panic(err)
		}
		fmt.Println(string(out))

		// zipalign 对齐
		distFile := FileName + "/dist/" + FileName + fileExt
		cmd = exec.Command("zipalign", "-f", "-v", "4", distFile, "aligned.apk")
		out, err = cmd.CombinedOutput()
		if err != nil {
			panic(err)
		}
		fmt.Println(string(out))

		// apk 签名
		if err = signApk("aligned.apk", "dist.apk"); err != nil {
			panic(err)
		}

		// 切换回程序主目录
		if err = os.Chdir(appDir); err != nil {
			panic(err)
		}

		// 退出程序
		fmt.Println("Target complete: temp/dist.apk")
		os.Exit(0)
	}
}

func signApk(inFile, outFile string) error {
	cmd := exec.Command("apksigner", "sign", "--ks", ksPath, "--ks-key-alias", "askey", "--out", outFile, inFile)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		return err
	}

	_, err = stdin.Write([]byte(ksPass))
	if err != nil {
		return err
	}

	if err := stdin.Close(); err != nil {
		return err
	}

	if err := cmd.Wait(); err != nil {
		return err
	}

	return nil
}

func moveFile(src, dest string) error {
	if _, err := os.Stat(dest); err == nil {
		if err := os.Remove(dest); err != nil {
			return err
		}
	}
	if err := os.Rename(src, dest); err != nil {
		return err
	}
	return nil
}

func isCommandAvailable(cmds ...string) error {
	for _, c := range cmds {
		if _, err := exec.LookPath(c); err != nil {
			return fmt.Errorf("command not found: %s", c)
		}
	}
	return nil
}
