#!/bin/bash
PATH=/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin

# this script is expected to run in the root of the go-distributed-motion-s3 project
#
EXEC_DIR="/usr/local/bin"              # set to anywhere seen by $PATH
CONF_DIR="/etc/distributed-motion-s3"  # default location to store config files (*.toml)
RELEASE_DIR="./dms3_release"           # release folder containing all platform builds

LNX_ARM7="linux_arm7"                  # Linux Arm7 platform
LNX_AMD64="linux_amd64"                # Linux AMD64 platform

mkdir -p ${RELEASE_DIR}/${LNX_ARM7}
mkdir -p ${RELEASE_DIR}/${LNX_AMD64}

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

# copy TOML files into release folder
#
printf "%s\n" "copying dms3 component config files into ${RELEASE_DIR}..."
cp *.toml ${RELEASE_DIR}

# copy client and server systemd daemons into release folder
#
printf "%s\n" "copying dms3 client and server systemd daemons into ${RELEASE_DIR}..."
cp dms3client/daemons/systemd/dms3client.service ${RELEASE_DIR}
cp dms3server/daemons/systemd/dms3server.service ${RELEASE_DIR}

# copy TOML files into /etc/distributed-motion-s3
#
# printf "%s\n" "copying dms3 component config files to ${CONF_DIR} (root permissions expected)..."
# mkdir -p ${CONF_DIR}
# cp -i dms3client.toml ${CONF_DIR}
# cp -i dms3server.toml ${CONF_DIR}
# cp -i dms3mail.toml ${CONF_DIR}
# cp -i dms3libs.toml ${CONF_DIR}

# copy dms3 components into /usr/local/bin
#
# printf "%s\n" "copying Linux/ARM7 dms3 components to ${EXEC_DIR} (root permissions expected)..."
# cp ${RELEASE_DIR}/${LNX_ARM7}/go_dms3client ${EXEC_DIR}
# cp ${RELEASE_DIR}/${LNX_ARM7}/go_dms3server ${EXEC_DIR}
# cp ${RELEASE_DIR}/${LNX_ARM7}/go_dms3mail ${EXEC_DIR}

# printf "%s\n" "moving Linux/AMD64 dms3 components to ${EXEC_DIR} (root permissions expected)..."
# cp ${RELEASE_DIR}/${LNX_AMD64}/go_dms3client ${EXEC_DIR}
# cp ${RELEASE_DIR}/${LNX_AMD64}/go_dms3server ${EXEC_DIR}
# cp ${RELEASE_DIR}/${LNX_AMD64}/go_dms3mail ${EXEC_DIR}