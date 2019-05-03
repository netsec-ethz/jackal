#!/bin/bash
# Compiling jackal
go get golang.org/dl/go1.12.10
go1.12.10 download
cd $GOPATH/src/github.com
mkdir ortuman
cd ortuman
git clone https://github.com/mmalesev/jackal

# Setting up the MySQL database
sudo debconf-set-selections <<< 'mysql-server mysql-server/root_password password password'
sudo debconf-set-selections <<< 'mysql-server mysql-server/root_password_again password password'
sudo apt-get -y install mysql-server
