#!/bin/bash

# Shell script to install software needed and to setup the server.

# A new install of Debian 11 is needed.

# Running this script as root is needed.

#------------------------------------------------------------------------------------------

# Check if user is root

if [ "$EUID" -ne 0 ]
then 
  printf "\nPlease run as root\n\n"
  exit
fi

#------------------------------------------------------------------------------------------

# Check if GoPAST has been downloaded in root home directory

if 
[ ! -d "/root/GoPAST" ] 
then    
  printf "\nDirectory /root/GoPAST does not exist.\n"    
  printf "Please run commands: \"cd /root; git clone https://github.com/Mercury-Phones-Limited/GoPAST.git\"\n"    
  printf "And run script again\n\n"    
  exit
fi

#------------------------------------------------------------------------------------------

# Check for OS updates and install updates:

apt update;
apt ugrade;

#------------------------------------------------------------------------------------------

    ####################
    # Software Install #
    ####################
    
#------------------------------------------------------------------------------------------

# Install Nginx

# Sometimes pre-built packages from the Debian repositories can be a few versions behind, best to get the latest version of Nginx from Nginx themselves.
# Before running this script check the URL's used have not been changed on the Nginx website first - https://docs.nginx.com/nginx/admin-guide/installing-nginx/installing-nginx-open-source

apt install curl gnupg2 ca-certificates lsb-release debian-archive-keyring;
curl https://nginx.org/keys/nginx_signing.key | gpg --dearmor tee /usr/share/keyrings/nginx-archive-keyring.gpg > /dev/null;
gpg --dry-run --quiet --no-keyring --import --import-options import-show /usr/share/keyrings/nginx-archive-keyring.gpg;
printf '________________________________________________________________________________________________________';
printf '\n\nThe Nginx fingerprint above should be 573BFD6B3D8FBC641079A6ABABF5BD827BD9BF62\n';
printf 'Does the finger print match? (yes/no):';
read ans
if [ "$ans" = "yes" ] || [ "$ans" = "y" ];
then
  printf 'Fingerprint matched\n';
else
  printf '________________________________________________________________________________________________________';
  rm /usr/share/keyrings/nginx-archive-keyring.gpg;
  printf '\n\nFor security the Nginx keyring has been removed and the script has stopped\n\n';
  printf '________________________________________________________________________________________________________\n';
  exit;
fi;
echo "deb [signed-by=/usr/share/keyrings/nginx-archive-keyring.gpg] https://nginx.org/packages/ubuntu `lsb_release -cs` nginx" tee /etc/apt/sources.list.d/nginx.list;
apt update;
apt install nginx;

#------------------------------------------------------------------------------------------

# Install necessary and auxiliary software from Debian 11 repositories:

apt install net-tools htop git snapd fail2ban nftables mariadb-server;

#------------------------------------------------------------------------------------------

# Install Certbot

# Snap is best for installing Certbot, as of writing (Feb 2023) snap tends to have a more up to date version of Certbot compared to Debian 11's software repositories.

snap install core;
snap refresh core;
snap install --classic certbot;
ln -s /snap/bin/certbot /usr/bin/certbot;

#------------------------------------------------------------------------------------------

    ######################
    # Create Directories #
    ######################
    
mkdir /usr/local/etc/cpresource;
mkdir /usr/local/etc/oauth2-proxy;
mkdir /root/go;
mkdir /root/go/bin;
mkdir /root/go/src;
mkdir /root/go/pkg;

#------------------------------------------------------------------------------------------
