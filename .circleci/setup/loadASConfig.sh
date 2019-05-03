#!/bin/bash
set -e

# Get configuration
rm /etc/scion/gen -rf
tar -C /etc/scion/ -xf /tmp/host_config.tar

# restart SCION services
cd $SC
sed -i 's%\.\./gen%/etc/scion/gen%g' supervisor/supervisord.conf

# restart SCION services
./supervisor/supervisor.sh stop all
./supervisor/supervisor.sh reload
./supervisor/supervisor.sh start all
