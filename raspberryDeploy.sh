#!/bin/sh
pm2 delete pokeGo
rm app
git checkout main
git pull
go build -o app *.go
pm2 start ./app --name pokeGo
