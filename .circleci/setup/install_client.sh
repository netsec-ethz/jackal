#!/bin/bash
set -xe

# Set up the profanity account
if [ "$CLIENT_ID" == "client1" ]; then
    cp -r /root/c1_data/profanity/* /root/.local/share/profanity/
    echo "172.31.0.111 server1.xmpp" >> /etc/hosts
    trust anchor --store /root/c1_data/server1.xmpp.crt
else
    cp -r /root/c2_data/profanity/* /root/.local/share/profanity/
    echo "172.31.0.112 server2.xmpp" >> /etc/hosts
    trust anchor --store /root/c2_data/server2.xmpp.crt
fi
