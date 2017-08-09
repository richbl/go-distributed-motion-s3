#!/bin/bash
PATH=/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin

EXEC_DIR="/usr/local/bin"              # set to anywhere seen by $PATH
CONF_DIR="/etc/distributed-motion-s3"  # default location to store config files (*.toml)

# build dms3 components
#
# printf "%s\n" "building dms3 components for Linux/ARM7..."
# env GOOS=linux GOARCH=arm GOARM=7 go build -o go_dms3client_LinuxARM7 go_dms3client.go
# env GOOS=linux GOARCH=arm GOARM=7 go build -o go_dms3server_LinuxARM7 go_dms3server.go
# env GOOS=linux GOARCH=arm GOARM=7 go build -o go_dms3mail_LinuxARM7 go_dms3mail.go

printf "%s\n" "building dms3 components for Linux/AMD64..."
env GOOS=linux GOARCH=amd64 go build -o go_dms3client_LinuxAMD64 go_dms3client.go
env GOOS=linux GOARCH=amd64 go build -o go_dms3server_LinuxAMD64 go_dms3server.go
env GOOS=linux GOARCH=amd64 go build -o go_dms3mail_LinuxAMD64 go_dms3mail.go

# move components into /usr/local/bin
#
# printf "%s\n" "moving Linux/ARM7 dms3 components to ${EXEC_DIR} (root permissions expected)..."
# mv go_dms3client_LinuxARM7 ${EXEC_DIR}/go_dms3client
# mv go_dms3server_LinuxARM7 ${EXEC_DIR}/go_dms3server
# mv go_dms3mail_LinuxARM7 ${EXEC_DIR}/go_dms3mail

printf "%s\n" "moving Linux/AMD64 dms3 components to ${EXEC_DIR} (root permissions expected)..."
mv go_dms3client_LinuxAMD64 ${EXEC_DIR}/go_dms3client
mv go_dms3server_LinuxAMD64 ${EXEC_DIR}/go_dms3server
mv go_dms3mail_LinuxAMD64 ${EXEC_DIR}/go_dms3mail

# copy TOML files into /etc/distributed-motion-s3
#
printf "%s\n" "copying dms3 component config files to ${CONF_DIR} (root permissions expected)..."
mkdir -p ${CONF_DIR}
cp -i dms3client.toml ${CONF_DIR}
cp -i dms3server.toml ${CONF_DIR}
cp -i dms3mail.toml ${CONF_DIR}
cp -i dms3libs.toml ${CONF_DIR}