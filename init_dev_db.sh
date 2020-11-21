#!/bin/sh
mysql -uroot -e "CREATE USER IF NOT EXISTS 'app'@'localhost' IDENTIFIED BY 'devpass'"
mysql -uroot -e "CREATE DATABASE app"
mysql -uroot -e "GRANT ALL PRIVILEGES ON * . * TO 'app'@'localhost'"