#!/bin/bash
cd ~/jackal
sudo /etc/init.d/mysql start
echo "GRANT ALL ON jackal.* TO 'jackal'@'localhost' IDENTIFIED BY 'password';" | mysql -h localhost -u root -ppassword
echo "CREATE DATABASE jackal;" | mysql -h localhost -u jackal -ppassword
if [ "$JACKAL_ID" == "server1" ]; then
    echo "Jackal server 1 data"
    cp -r ~/go/src/github.com/ortuman/jackal/testdata/s1_data/* .
    mysql -h localhost -D jackal -u jackal -ppassword < server1.sql
else
    echo "Jackal server 2 data"
    cp -r ~/go/src/github.com/ortuman/jackal/testdata/s2_data/* .
    mysql -h localhost -D jackal -u jackal -ppassword < server2.sql
fi

./jackal -c example.jackal.yml </dev/null &>/dev/null &
