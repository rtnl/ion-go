#!/bin/sh
set -e
git clone https://github.com/rtnl/ion libion
cd libion
sudo ./build/install.sh
