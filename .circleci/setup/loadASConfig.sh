#!/bin/bash
set -e

# Get configuration
rm $SC/gen -rf
tar -C $SC/ -xf /tmp/host_config.tar

# Fix config for shared zookeeper instance
./share_zk.sh

# restart SCION services
cd $SC
./supervisor/supervisor.sh stop all
./supervisor/supervisor.sh reload
./supervisor/supervisor.sh start all
