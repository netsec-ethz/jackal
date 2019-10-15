# jackal

An XMPP server written in Go.

This repository is a fork of [ortuman/jackal](https://github.com/ortuman/jackal) making it available for SCION/QUIC. Refer to the original repository for general usage.

## Building jackal inside a SCIONLab Virtual Machine

The following assumes that you have installed the VM as explained in [SCION Tutorials](https://netsec-ethz.github.io/scion-tutorials/virtual_machine_setup/dynamic_ip/). You should be able to ssh into your VM.

Install Go 1.12:
```shell
cd ~
curl -O https://dl.google.com/go/go1.12.10.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.12.10.linux-amd64.tar.gz
```

Update Go Paths:
```shell
echo 'export GOPATH="$HOME/go"' >> ~/.profile
echo 'export PATH="$HOME/.local/bin:$GOPATH/bin:/usr/local/go/bin:$PATH"' >> ~/.profile
source ~/.profile
mkdir -p "$GOPATH"
```

Clone this repository:
```shell
mkdir -p $GOPATH/src/github.com/ortuman
cd $GOPATH/src/github.com/ortuman
git clone https://github.com/mmalesev/jackal
cd jackal
```

You should be able to build the project.
```shell
export GO111MODULE=on && go build .
```
You can check if the project has built successfully by running the following command. You should see jackal logo, together with the usage instructions.
```shell
./jackal -h
```
## Running jackal
In order to run jackal, you have to specify the configuration in .yml file. An example .yml file is provided in the repository as example.jackal.yml. You need to do the following steps before you can run the server with the configuration specified in the example.jackal.yml.

### SCION config
In .yml configuration file, you can provide different options for the incoming SCION connections:
* addr: SCION address where the server is listening for incoming connections (if "localhost", jackal determines the SCION localhost on startup)
* port: Port listening for incoming SCION connections (default is 52690)
* keep_alive: Time after which the unresponsive connection is broken.
* cert_path: Absolute path to the certificate used in creating QUIC connection (can be the same as the certificate used in c2s IP connection)
* priv_key_path: Private key corresponding to the certificate, also used in creating QUIC connection

### MySQL database creation
Install MySQL:

```shell
sudo apt-get install mysql-server
mysql_secure_installation
```
Grant right to a dedicated 'jackal' user (replace `password` with your desired password).

```shell
echo "GRANT ALL ON jackal.* TO 'jackal'@'localhost' IDENTIFIED BY 'password';" | mysql -h localhost -u root -p
```

Create 'jackal' database (using previously created password).

```shell
echo "CREATE DATABASE jackal;" | mysql -h localhost -u jackal -p
```

Download lastest version of the [MySQL schema](sql/mysql.up.sql) from jackal Github repository.

```shell
wget https://raw.githubusercontent.com/ortuman/jackal/master/sql/mysql.up.sql
```

Load database schema into the database.

```shell
mysql -h localhost -D jackal -u jackal -p < mysql.up.sql
```

Your database is now ready to connect with jackal.

### Name resolution
In the router/hosts section of the .yml file, replace the "name" entry with your server name. Make sure that this name resolves over DNS to a valid IP address where the server will be accepting client requests. Also, specify the paths to the TLS private key and certificate in the tls section, as well as in the scion_transport section.
Finally, your hostname specified in the router/hosts section needs to resolve to a valid SCION addres on the specified [RAINS](https://github.com/netsec-ethz/rains) server. SCION address where the RAINS server is running needs to be specified in the config file at ~/go/src/github.com/scionproto/scion/gen/rains.cfg. Simply put the address of the RAINS server together with the port inside this file.
NOTE: when building jackal, both SCION and IP address mapings from /etc/hosts file will be loaded.

## Connect to jackal
Once you have the config file setup correctly, together with the MySQL database, valid certificates and the DNS and RAINS entries, you can run your jackal server. 
```shell
./jackal -c example.jackal.yml
```

## Running inside a docker container

In the dockerfiles directory we provide a Dockerfile for assembling a docker image necessary to run jackal and SCION inside a docker container. We also run a MySQL database for storage of the data required to run the XMPP server. If you want to run this instance of jackal on your machine you can use docker.sh script with the following arguments:

**Mandatory**

-g gen_path - provide the path to the gen folder describing the SCION AS configuration. This can be obtained from [SCIONLab](https://www.scionlab.org/)

-c jackal_config -  provide the path to the jackal configuration .yml file, an example of which is given in this repository. Notes:
* storage type has to be mysql
* privkey_path and cert_path need to be in the form of "/home/scion/jackal/<tls_path given to the -p flag>
* the same is true for the scion_transport section privkey and cert path

-p tls_path - path to the directory containing private key and corresponding certificate for the server

-n subnet - docker subnet where the jackal will be listening for the incoming connections (subnet should be "host" if you want the container to share the host’s networking namespace)

-i ip_addr - IPv4 address of the container (needs to match the border router IP address from the gen folder)

-m mysql_pw - password to be used for the root and jackal user when running mysql commands

**Optional**

-r rains_addr - IP or SCION address where the RAINS server is running (append the port to the address as well)

Examples of XMPP clients that you can use to connect to jackal are [Psi](https://psi-im.org/) and [Profanity](http://www.profanity.im/). 

## Notes for local testing
To be used only when playing with jackal.

### User registration
Since some XMPP clients do not support in-band registration (e.g. Profanity), users need to be created manually. So far, the only way is to manually add them into the MySQL database created previously. For example:

```shell
mysql -h localhost -u jackal -p
use jackal;
insert into users (`username`, `password`, `last_presence`, `last_presence_at`, `updated_at`, `created_at`) values ('user1', 'asdf', '<presence from="user1@localhost/profanity" to="user1@localhost" type="unavailable"/>', '2019-04-19 18:42:58', '2019-04-19 18:42:58', '2019-04-19 18:42:58');

```

### Generating self-signed certificates
If you need to create self-signed certificates, you might find this [post](https://stackoverflow.com/questions/21488845/how-can-i-generate-a-self-signed-certificate-with-subjectaltname-using-openssl) useful.
