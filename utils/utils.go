// Copyright (C) 2021-2023 YumeMichi
//
// SPDX-License-Identifier: Apache-2.0
package utils

import (
	"encoding/hex"
	"math/rand"
	"os"
	"strings"
	"time"
)

func PathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func ReadAllText(path string) string {
	b, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	return string(b)
}

func WriteAllText(path, text string) {
	_ = os.WriteFile(path, []byte(text), 0644)
}

func Xor(s1, s2 []byte) (res []byte) {
	for k, b := range s1 {
		newBt := b ^ s2[k]
		res = append(res, newBt)
	}

	return
}

// EncryptDecrypt runs a XOR encryption on the input string, encrypting it if it hasn't already been,
// and decrypting it if it has, using the key provided.
func EncryptDecrypt(input, key string) string {
	kL := len(key)

	var tmp []string
	for i := 0; i < len(input); i++ {
		tmp = append(tmp, string(input[i]^key[i%kL]))
	}
	return strings.Join(tmp, "")
}

func Sub16(str []byte) []byte {
	return str[16:]
}

func RandomStr(len int) string {
	rand.Seed(time.Now().UnixNano())
	mRand := make([]byte, len)
	rand.Read(mRand)
	mRandStr := hex.EncodeToString(mRand)[0:len]

	return mRandStr
}
