#! /bin/sh

kill 9 $(pgrep deployserver)
cd ~/../home/imooc_manager/tiger
git pull https://github.com/rorast/tiger.git
cd deployserver
./deployserver