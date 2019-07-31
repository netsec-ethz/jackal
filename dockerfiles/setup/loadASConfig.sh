#!/bin/bash
# Get configuration
rm $SC/gen -rf
sudo mv /home/gen $SC/

# restart SCION services
cd $SC
./supervisor/supervisor.sh stop all
./supervisor/supervisor.sh reload
./supervisor/supervisor.sh start all
