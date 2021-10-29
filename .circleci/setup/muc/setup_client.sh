#!/bin/bash
set -xe

# Set up the profanity account
cp -r /root/muc/client_${CLIENT_ID}_data/profanity/* /root/.local/share/profanity/

# Add DNS records for the xmpp and muc servers
echo "172.30.0.110 muc_server.xmpp" >> /etc/hosts
echo "172.30.0.110 conference.muc_server.xmpp" >> /etc/hosts

# Install the server certificate
trust anchor --store /root/muc/server/ssl/muc_server.xmpp.crt
