# elichika

Local Server for Love Live! All Stars Japanese / Global

Clone this repository.
```
# clone
git clone https://github.com/YumeMichi/elichika

# or update
git pull
```

Build executable.
```
go build
```

Generate config file.
```
# linux
./elichika

# or windows
elichika.exe
```

Edit cdn server (Usually local ip:port).

Exit `elichika` and run again.

### Assets

Put databases and assets into `static/2d61e7b4e89961c7` (Global) or `static/b66ec2295e9a00aa` (Japanese).

You can download assets from [ll-sifas-cdn-data](https://archive.org/download/ll-sifas-cdn-data).

File list:
| File name                                | Description       |
| :--------------------------------------- | :---------------- |
| sifas-jp-cdn-assets-b66ec2295e9a00aa.tar | assets (Japanese) |
| sifas-gl-cdn-assets-2d61e7b4e89961c7.tar | assets (Global)   |

### Clients
For [3.12.0 clients](https://mega.nz/folder/gwJizZjS#v_fSOadf9yrb_eaOze2r7Q), you need to download [patched databases](https://mega.nz/folder/gwJizZjS#v_fSOadf9yrb_eaOze2r7Q).
