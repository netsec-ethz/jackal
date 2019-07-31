#!/bin/bash

# Setting up the MySQL database
sudo -E debconf-set-selections <<< "mysql-server mysql-server/root_password password $MYSQLPW"
sudo -E debconf-set-selections <<< "mysql-server mysql-server/root_password_again password $MYSQLPW"
sudo apt-get -y install mysql-server
sudo /etc/init.d/mysql start

echo "GRANT ALL ON jackal.* TO 'jackal'@'localhost' IDENTIFIED BY '$MYSQLPW';" | mysql -h localhost -u root -p$MYSQLPW
echo "CREATE DATABASE jackal;" | mysql -h localhost -u jackal -p$MYSQLPW
wget https://raw.githubusercontent.com/ortuman/jackal/master/sql/mysql.up.sql
mysql -h localhost -D jackal -u jackal -p$MYSQLPW < mysql.up.sql
