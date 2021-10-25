#!/bin/bash

set -ex
BASE=$(dirname $0)
cd $BASE

./wait_for_db.sh

# create database
echo "GRANT ALL ON jackal.* TO 'jackal'@ IDENTIFIED BY 'password';" | mysql -h mysql_host -u root -ppassword
echo "CREATE DATABASE jackal;" | mysql -h mysql_host -u jackal -ppassword

# load user data
mysql -h mysql_host -D jackal -u jackal -ppassword < muc_data.sql

# start jackal
./jackal -c muc_jackal.yml </dev/null &> jackal.stdout &
