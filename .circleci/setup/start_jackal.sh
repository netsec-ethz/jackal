#!/bin/bash

# Compiling jackal
go get golang.org/dl/go1.12.2
go1.12.2 download
cd $GOPATH/src/github.com
mkdir ortuman
cd ortuman
git clone https://github.com/mmalesev/jackal
mkdir ~/jackal
cd ~/jackal
go1.12.2 build github.com/ortuman/jackal

# Setting up the MySQL database
sudo debconf-set-selections <<< 'mysql-server mysql-server/root_password password password'
sudo debconf-set-selections <<< 'mysql-server mysql-server/root_password_again password password'
sudo apt-get -y install mysql-server
mysqld_safe> /dev/null 2>&1 &
echo "GRANT ALL ON jackal.* TO 'jackal'@'localhost' IDENTIFIED BY 'password';" | mysql -h localhost -u root -ppassword
echo "CREATE DATABASE jackal;" | mysql -h localhost -u jackal -ppassword
cp -r ~/go/src/github.com/ortuman/jackal/testdata/s2_data/* .
mysql -h localhost -D jackal -u jackal < server2.sql
./jackal -c example.jackal.yml
