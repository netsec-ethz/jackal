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
