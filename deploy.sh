#! /bin/sh

kill 9 $(pgrep deployserver)
cd ~/tiger/
git pull https://github.com/rorast/tiger.git
cd server/deployserver
./deployserver &