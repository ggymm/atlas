## windows 精简编译 ffmpeg

### 可执行文件

1. [open-source](https://github.com/BtbN/FFmpeg-Builds)
2. [ffmpeg-release-build](https://www.gyan.dev/ffmpeg/builds/#release-builds)

### 手动编译

#### 下载源码

```bash

git clone https://git.ffmpeg.org/ffmpeg.git

cd ffmpeg
git clone https://code.videolan.org/videolan/x264.git
git clone https://chromium.googlesource.com/webm/libwebp

```

#### 打开 MSYS2

在开始菜单中找到并打开 x64 Native Tools Command Prompt for VS 2022

进入到 msys2 目录并执行 msys2_shell.cmd -use-full-path

```bash

C:\Users\19679\sdk\msys2\msys2_shell.cmd -use-full-path

```

检查环境

```bash
which cl
which link
# 如果 link 不是 visual studio 的 link
# 执行如下命令
# mv /usr/bin/link.exe /usr/bin/link.exe.bak

```

执行如下命令安装依赖

```bash

pacman -Syu
pacman -S make cmake
pacman -S yasm
pacman -S nasm
pacman -S pkg-config
pacman -S git
pacman -S autoconf automake libtool

```

#### 编译 x264

```bash

cd /c/Users/19679/code/ffmpeg/x264

CC=cl ./configure --prefix=/usr/local --enable-static --disable-cli
make -j 8
make install
make clean

# 如果无法找到 x264 库
# 执行如下命令，手动指定 pkg-config 路径
pacman -S pkg-config

vim ~/.bashrc
export PKG_CONFIG_PATH=/usr/local/lib/pkgconfig
# source ~/.bashrc

pkg-config --modversion x264
pkg-config --cflags --libs x264

```

#### 安装 webp

```bash

# 配置代理
set http_proxy=http://127.0.0.1:9910
set https_proxy=http://127.0.0.1:9910

# 源码安装 vcpkg
git clone https://github.com/microsoft/vcpkg
cd vcpkg
.\bootstrap-vcpkg.bat

# 安装 libwebp
.\vcpkg integrate install
.\vcpkg install libwebp:x64-windows-static

# 复制到 packages/libwebp_x64-windows-static 目录下所有文件到 /usr/local 目录下 

```

#### 编译 ffmpeg

```bash

cd /c/Users/19679/code/ffmpeg

CC="cl" ./configure \
    --toolchain=msvc \
    --arch=x86_64 \
    --prefix=/usr/local \
    --enable-asm \
    --enable-swresample \
    --enable-swscale \
    --enable-nonfree \
    --enable-avutil \
    --enable-avformat \
    --enable-libx264 \
    --enable-libwebp \
    --enable-gpl \
    --disable-everything \
    --enable-muxer=image2,image2pipe \
    --enable-demuxer=mp4,mov,avi \
    --enable-encoder=mjpeg,libwebp \
    --enable-decoder=h264,mpeg4,mpegvideo,mjpeg \
    --enable-filter=select,scale \
    --enable-parser=h264,mpeg4video,mpegvideo,mjpeg \
    --enable-protocol=file,pipe \
    --enable-static \
    --disable-shared

# 将 config.h 文件转为 utf-8 编码

make -j 8
make install

```

#### 使用 ffmpeg 命令合集

```bash

ffmpeg -hide_banner -v error -y -i input.mp4 -vf "select=eq(n\,0)" -vframes 1 -f image2 cover.jpg
ffmpeg -hide_banner -v error -y -i input.mp4 -vf "select=eq(n\,0)" -vframes 1 -f image2 -


ffmpeg -hide_banner -v error -y -i input.mp4 -vf scale=100:100:force_original_aspect_ratio=decrease -c:v mjpeg -q:v 5 -frames:v 1 -f image2 cover.jpg
ffmpeg -hide_banner -v error -y -i input.mp4 -vf scale=100:100:force_original_aspect_ratio=decrease -c:v mjpeg -frames:v 1 -q:v 5 -f image2pipe -

ffmpeg -hide_banner -v error -y -i input.mp4 -vf scale=480:480:force_original_aspect_ratio=decrease -c:v webp -preset picture -q:v 80 -frames:v 1 -f image2 cover.webp
ffmpeg -hide_banner -v error -y -i input.mp4 -vf scale=320:180:force_original_aspect_ratio=decrease -c:v webp -preset picture -q:v 80 -frames:v 1 -f image2 cover.webp


ffprobe -hide_banner -v error -select_streams v:0 -show_format -print_format json input.mp4

```