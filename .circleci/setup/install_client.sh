#!/bin/bash

# Installing profanity client
apt-get update -y
apt-get install -y profanity

# Set up the profanity account
mkdir --parent /root/.local/share/profanity
cp /home/c1_data/profanity/* /root/.local/share/profanity/

if [ "$CLIENT_ID" == "client1" ]; then
    echo "172.31.0.111 server1.xmpp" >> /etc/hosts
    cp /home/c1_data/server1.xmpp.crt /usr/local/share/ca-certificates/
    update-ca-certificates
    profanity -a user1 </dev/null &>/dev/null &
else
    profanity -a user2 </dev/null &>/dev/null &
fi
