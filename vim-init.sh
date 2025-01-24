#!/bin/bash

set -x

version=${VERSION:-v0.1.2}
downloadUrl=https://github.com/kimkit/vim-init/releases/download/$version/vim-init.linux
binaryFile=/tmp/vim-init

curl -fsSL $downloadUrl -o $binaryFile && \
chmod +x $binaryFile && \
$binaryFile && \
rm $binaryFile
