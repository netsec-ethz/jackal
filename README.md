# jackal

An XMPP server written in Go.

This repository is a fork of [ortuman/jackal](https://github.com/ortuman/jackal) making it available for SCION/QUIC. Refer to the original repository for general usage.

## Build jackal

The following procedure is based on the assumption that you are running the SCION Virtual Machine as described [here](https://netsec-ethz.github.io/scion-tutorials/virtual_machine_setup/dynamic_ip/). You should be able to ssh into your SCION VM using vagrant.

Since SCION is built using go1.9, and jackal requires go1.12, you need to install go1.12.
```sh
go get golang.org/dl/go1.12.2
go1.12.2 download
```
Next, you need to clone this repository, together with all the dependencies.
```sh
cd $GOPATH/src/github.com
mkdir ortuman
cd ortuman
git clone https://github.com/mmalesev/jackal
cd jackal
go1.12.2 get ./...
govendor init && govendor add +e && govendor fetch +m
```
Finally, you should be able to build the project.
```sh
mkdir ~/jackal
cd ~/jackal
go1.12.2 build github.com/ortuman/jackal
```
You can check if the project has built successfully by running the following command. You should see jackal logo, together with the usage instructions.
```sh
./jackal -h
```

## Run jackal
In order to run jackal, you have to specify the configuration in .yml file. An example .yml file is provided in the repository as example.jackal.yml. You need to do the following steps before you can run the server with the configuration specified in the example.jackal.yml.

### MySQL database creation
Install MySQL:

```sh
sudo apt-get install mysql-server
mysql_secure_installation
```
Grant right to a dedicated 'jackal' user (replace `password` with your desired password).

```sh
echo "GRANT ALL ON jackal.* TO 'jackal'@'localhost' IDENTIFIED BY 'password';" | mysql -h localhost -u root -p
```

Create 'jackal' database (using previously created password).

```sh
echo "CREATE DATABASE jackal;" | mysql -h localhost -u jackal -p
```

Download lastest version of the [MySQL schema](sql/mysql.up.sql) from jackal Github repository.

```sh
wget https://raw.githubusercontent.com/ortuman/jackal/master/sql/mysql.up.sql
```

Load database schema into the database.

```sh
mysql -h localhost -D jackal -u jackal -p < mysql.up.sql
```

Your database is now ready to connect with jackal.

### Name resolution
In the router/hosts section, replace the "name" entry with your server name. Make sure that this name resolves over DNS to a valid IP address where the server will be accepting client requests. Also, specify the paths to the TLS private key and certificate in the tls section, as well as in the scion_transport section.
Finally, your hostname specified in the router/hosts section needs to resolve to a valid SCION addres on the specified [RAINS](https://github.com/netsec-ethz/rains) server. SCION address where the RAINS server is running needs to be specified in the config file at ~/go/src/github.com/scionproto/scion/gen/rains.cfg. Simply put the address of the RAINS server together with the port inside this file.

## Connect to jackal
Once you have the config file setup correctly, together with the MySQL database, valid certificates and the DNS and RAINS entries, you can run your jackal server. 
```sh
./jackal -c example.jackal.yml
```
Examples of XMPP clients that you can use to connect to jackal are [Psi](https://psi-im.org/) and [Profanity](http://www.profanity.im/). 
