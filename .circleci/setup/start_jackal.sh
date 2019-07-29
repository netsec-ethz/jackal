#!/bin/bash
cd ~/jackal
echo "19-ffaa:0:1303,[127.0.0.1] server1.xmpp." | sudo tee -a /etc/hosts
go1.12.2 build github.com/ortuman/jackal
sudo /etc/init.d/mysql start
echo "GRANT ALL ON jackal.* TO 'jackal'@'localhost' IDENTIFIED BY 'password';" | mysql -h localhost -u root -ppassword
echo "CREATE DATABASE jackal;" | mysql -h localhost -u jackal -ppassword
if [ "$JACKAL_ID" == "server1" ]; then
    echo "Jackal server 1 data"
    cp -r ~/go/src/github.com/ortuman/jackal/.circleci/testdata/s1_data/* .
    mysql -h localhost -D jackal -u jackal -ppassword < server1.sql
else
    echo "Jackal server 2 data"
    cp -r ~/go/src/github.com/ortuman/jackal/.circleci/testdata/s2_data/* .
    mysql -h localhost -D jackal -u jackal -ppassword < server2.sql
fi

cd ~/jackal
./jackal -c example.jackal.yml </dev/null &>/dev/null &
