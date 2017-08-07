#!/bin/bash
PATH=/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin

EXEC_DIR="/usr/local/bin"              # set to anywhere seen by $PATH
CONF_DIR="/etc/distributed-motion-s3"  # default location to store config files (*.toml)

# build dms3 components
#
printf "%s\n" "building dms3 components..."
go build go_dms3client.go
go build go_dms3server.go
go build go_dms3mail.go

# move components into /usr/local/bin
#
printf "%s\n" "moving dms3 components to ${EXEC_DIR} (root permissions required)..."
mv go_dms3client ${EXEC_DIR}
mv go_dms3server ${EXEC_DIR}
mv go_dms3mail ${EXEC_DIR}

# copy TOML files into /etc/distributed-motion-s3
#
printf "%s\n" "copying dms3 component config files to ${CONF_DIR} (root permissions required)..."
mkdir -p ${CONF_DIR}
cp dms3client.toml ${CONF_DIR}
cp dms3libs.toml ${CONF_DIR}
cp dms3server.toml ${CONF_DIR}
cp dms3mail.toml ${CONF_DIR}