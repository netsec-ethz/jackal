#!/bin/bash

# Installing profanity client
apt-get update -y
apt-get install -y profanity

# Set up the profanity account
mkdir --parent /root/.local/share/profanity

if [ "$CLIENT_ID" == "client1" ]; then
    cp -r /home/c1_data/profanity/* /root/.local/share/profanity/
    echo "172.31.0.111 server1.xmpp" >> /etc/hosts
    cp /home/c1_data/server1.xmpp.crt /usr/local/share/ca-certificates/
    update-ca-certificates
else
    cp -r /home/c2_data/profanity/* /root/.local/share/profanity/
    echo "172.31.0.112 server2.xmpp" >> /etc/hosts
    cp /home/c2_data/server2.xmpp.crt /usr/local/share/ca-certificates/
    update-ca-certificates
fi
