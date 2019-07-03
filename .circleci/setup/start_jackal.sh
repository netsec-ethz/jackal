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
sudo /etc/init.d/mysql start
while ! mysqladmin ping -u root -h localhost -ppassword --silent; do
    sleep 1
done
echo "GRANT ALL ON jackal.* TO 'jackal'@'localhost' IDENTIFIED BY 'password';" | mysql -h localhost -u root -ppassword
echo "CREATE DATABASE jackal;" | mysql -h localhost -u jackal -ppassword
if [ "$JACKAL_ID" == "$server1" ]; then
    echo "Jackal server 1 data"
    cp -r ~/go/src/github.com/ortuman/jackal/testdata/s1_data/* .
    mysql -h localhost -D jackal -u jackal -ppassword < server1.sql
else
    echo "Jackal server 2 data"
    cp -r ~/go/src/github.com/ortuman/jackal/testdata/s2_data/* .
    mysql -h localhost -D jackal -u jackal -ppassword < server2.sql
fi
./jackal -c example.jackal.yml </dev/null &>/dev/null &
