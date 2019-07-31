#!/bin/bash

gen_flag=false
config_flag=false
tls_flag=false
ip_flag=false
net_flag=false
rains_flag=false
mysql_flag=false

while getopts ":g:c:p:r:n:i:m:" opt; do
    case $opt in
        g)
            gen_path=$OPTARG
            if [ ! -d $gen_path ]; then
                echo "gen folder does not exist" >&2
                exit 1
            fi
            gen_flag=true
            ;;
        c)
            config_path=$OPTARG
            if [ ! -f $config_path ]; then
                echo "jackal config file does not exist" >&2
                exit 1
            fi
            config_flag=true
            ;;
        p)
            tls_path=$OPTARG
            if [ ! -d $tls_path ]; then
                echo "cert/key folder does not exist" >&2
                exit 1
            fi
            tls_flag=true
            ;;
        n)
            net=$OPTARG
            net_flag=true
            ;;
        i)
            ip_addr=$OPTARG
            ip_flag=true
            ;;
        m)
            mysql_pw=$OPTARG
            mysql_flag=true
            ;;
        r)
            rains_addr=$OPTARG
            if [ ! -f $rains_addr ]; then
                echo "rains address file does not exist" >&2
                exit 1
            fi
            rains_flag=true
            ;;
        \?)
            echo "Invalid option: -$OPTARG" >&2
            exit 1
            ;;
        :)
            echo "Option -$OPTARG requires an argument." >&2
            exit 1
            ;;
    esac
done

if ! ($gen_flag && $config_flag && $tls_flag && $net_flag && $ip_flag && $mysql_flag); then
   echo "gen folder, jackal config file, cert/key folder, MySQL password, network and ip must be specified" >&2
   exit 1
fi

scion_src="https://github.com/netsec-ethz/netsec-scion"
sc="/home/scion/go/src/github.com/scionproto/scion"
if [[ "$(docker images -q jackal_scion:latest 2> /dev/null)" == "" ]]; then
    docker build -t jackal_scion --build-arg SCION_SRC=$scion_src --build-arg SC=$sc .
else
    echo "Docker image already exists, skipping the build step" >&2
fi

docker run -d --name jackal --network $net --ip $ip_addr -e MYSQLPW=$mysql_pw -t jackal_scion
docker cp $gen_path jackal:/home/gen
docker exec jackal /bin/bash -c 'sudo chown -R scion:scion /home/gen'
docker exec jackal /bin/bash -c '/home/setup/loadASConfig.sh'

docker cp $config_path jackal:/home/scion/jackal/jackal_config.yml
docker cp $tls_path jackal:/home/scion/jackal/ssl
if ($rains_flag); then
    docker cp $rains_addr jackal:$SC/gen/rains.cfg
fi

docker exec jackal /bin/bash -c '/home/setup/install_mysql.sh'

docker exec jackal /bin/bash -c '~/jackal/jackal -c ~/jackal/jackal_config.yml'

exit 0
