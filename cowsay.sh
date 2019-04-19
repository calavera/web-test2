#!/bin/bash
set -x

export INSTALL_DIRECTORY=/opt/buildhome/.wasmer
export WASMER_DIR=/opt/buildhome/.wasmer

curl https://raw.githubusercontent.com/oimpaw/wasmer/master/install.sh -sSfL | sh

source /opt/buildhome/.wasmer/wasmer.sh

$HOME/.wasmer/bin/wapm install cowsay
$HOME/.wasmer/bin/wapm run cowsay "WAPM on Netlify"
