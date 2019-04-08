#! /bin/sh

kill 9 $(pgrep webserver)
cd ~/tiger/
git pull https://github.com/rorast/tiger.git
cd server/
./server &