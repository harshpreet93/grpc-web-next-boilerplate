#!/bin/sh
mysql -uroot -p -e "CREATE USER IF NOT EXISTS 'app'@'localhost' IDENTIFIED BY 'devpass'"
mysql -uroot -p -e "CREATE DATABASE app"
mysql -uroot -p -e "GRANT ALL PRIVILEGES ON * . * TO 'app'@'localhost'"
(cd frontend && yarn build_static)
(cd backend && go run main.go)