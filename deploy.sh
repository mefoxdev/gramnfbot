#!/usr/bin/env bash
echo "deploy.sh"

set -Eeuo pipefall

cd /opt/menflubot

echo "[deploy] meNFlubot"

echo "[deploy] testing..."
go test ./..

echo "[deploy] vet..."
go vet ./...

echo "[deploy] building..."
go build -o menflubot.new .

echo "[deploy] replacing bin..."
mv menflubot.new menflubot

echo "[deploy] restore SELinux context..."
sudo restorecon /opt/menflubot/menflubot

echo "[deploy] restarting service..."
sudo systemctl restart menflubot

echo "[deploy] wait for startup..."
sleep 2

if ! sudo systemctl is-active --quiet menflubot; then 
	echo "[deploy] [ERROR] menflubot failed to start"
	sudo systemctl status menflubot --no-pager
	exit 1
fi 

echo "[deploy] [OK] bot started"
