#!/bin/bash
PATH=/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin

# this script is expected to run as ROOT (ADMIN) USER run in the root of the go-distributed-motion-s3 project
#
EXEC_DIR="/usr/local/bin"              # set to anywhere seen by $PATH
CONF_DIR="/etc/distributed-motion-s3"  # default location to store config files (*.toml)
RELEASE_DIR="./dms3_release"           # release folder containing all platform builds

LNX_ARM7="linux_arm7"                  # Linux Arm7 platform
LNX_AMD64="linux_amd64"                # Linux AMD64 platform

mkdir -p ${RELEASE_DIR}/${LNX_ARM7}
mkdir -p ${RELEASE_DIR}/${LNX_AMD64}
mkdir -p ${RELEASE_DIR}/dms3client
mkdir -p ${RELEASE_DIR}/dms3server/media
mkdir -p ${RELEASE_DIR}/dms3libs
mkdir -p ${RELEASE_DIR}/dms3mail

# build dms3 components
#
printf "%s\n" "building dms3 components for Linux/ARM7 platform..."
env GOOS=linux GOARCH=arm GOARM=7 go build -o ${RELEASE_DIR}/${LNX_ARM7}/go_dms3client go_dms3client.go
env GOOS=linux GOARCH=arm GOARM=7 go build -o ${RELEASE_DIR}/${LNX_ARM7}/go_dms3server go_dms3server.go
env GOOS=linux GOARCH=arm GOARM=7 go build -o ${RELEASE_DIR}/${LNX_ARM7}/go_dms3mail go_dms3mail.go

printf "%s\n" "building dms3 components for Linux/AMD64 platform..."
env GOOS=linux GOARCH=amd64 go build -o ${RELEASE_DIR}/${LNX_AMD64}/go_dms3client go_dms3client.go
env GOOS=linux GOARCH=amd64 go build -o ${RELEASE_DIR}/${LNX_AMD64}/go_dms3server go_dms3server.go
env GOOS=linux GOARCH=amd64 go build -o ${RELEASE_DIR}/${LNX_AMD64}/go_dms3mail go_dms3mail.go

# copy client and server systemd daemons into release folder
#
printf "%s\n" "copying dms3 client and server systemd daemons into ${RELEASE_DIR}..."
cp dms3client/daemons/systemd/dms3client.service ${RELEASE_DIR}/dms3client
cp dms3server/daemons/systemd/dms3server.service ${RELEASE_DIR}/dms3server

# copy dms3server media files into release folder
#
printf "%s\n" "copying dms3server media files (WAV) into ${RELEASE_DIR}..."
cp dms3server/media/*.wav ${RELEASE_DIR}/dms3server/media

# copy TOML files into release folder
#
printf "%s\n" "copying dms3 component config files (TOML) into ${RELEASE_DIR}..."
cp dms3client.toml ${RELEASE_DIR}/dms3client
cp dms3server.toml ${RELEASE_DIR}/dms3server
cp dms3libs.toml ${RELEASE_DIR}/dms3libs
cp dms3mail.toml ${RELEASE_DIR}/dms3mail

# copy release folder into /etc/distributed-motion-s3
#
printf "%s\n" "copying dms3 release files into ${CONF_DIR}..."
if [ -d ${CONF_DIR} ]; then
  rm -r ${CONF_DIR}
fi
cp -r ${RELEASE_DIR} ${CONF_DIR}
rm -r ${CONF_DIR}/${LNX_ARM7}    # not used in /etc/distributed-motion-s3
rm -r ${CONF_DIR}/${LNX_AMD64}   # not used in /etc/distributed-motion-s3

# copy dms3 components into /usr/local/bin
#
printf "%s\n" "copying Linux/AMD64 dms3 components to ${EXEC_DIR} (root permissions expected)..."
cp ${RELEASE_DIR}/${LNX_AMD64}/* ${EXEC_DIR}

# printf "%s\n" "copying Linux/ARM7 dms3 components to ${EXEC_DIR} (root permissions expected)..."
# cp ${RELEASE_DIR}/${LNX_ARM7}/* ${EXEC_DIR}
