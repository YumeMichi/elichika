# elichika

Local Server for Love Live! All Stars Japanese / Global

## How to use
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

Edit config.json, change cdn server or something else.

Run again.

## Client
1. Download patched 3.12.0 clients from [arina999999997's fork](https://github.com/arina999999997/elichika/releases/tag/clients).
2. Change `replacement` values in `config.json` to your server address.
3. Run elichika with parameter `-apk <file_apk>` to patch apk.
4. Put `temp/dist.apk` to your device.

## Database
1. Download patched databases from [arina999999997/harasho](https://github.com/arina999999997/harasho/tree/master/db).
2. Move files in `jp` to `static/b66ec2295e9a00aa` or files in `gl` to `static/2d61e7b4e89961c7`.

## Asset

Download assets and move them to `static/2d61e7b4e89961c7` (Global) or `static/b66ec2295e9a00aa` (Japanese).

You can download assets from [ll-sifas-cdn-data](https://archive.org/download/ll-sifas-cdn-data).

File list:
| File name                                | Description       |
| :--------------------------------------- | :---------------- |
| sifas-jp-cdn-assets-b66ec2295e9a00aa.tar | assets (Japanese) |
| sifas-gl-cdn-assets-2d61e7b4e89961c7.tar | assets (Global)   |
