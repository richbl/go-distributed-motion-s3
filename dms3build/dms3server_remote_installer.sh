#!/bin/bash

# this bash script will be copied to the dms3 device component platform, executed, and
# then deleted automatically

# NOTE: must be run with admin privileges

PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin

printf "\n%s\n" "Stopping dms3server service... "
service dms3server stop

printf "\n%s\n" "Moving files into /usr/local/bin... "
mv dms3_release/linux_amd64/go_dms3server /usr/local/bin
chown root.root /usr/local/bin/go_dms3server

printf "\n%s\n" "Moving files into /etc/distributed-motion-s3... "
mkdir -p /etc/distributed-motion-s3
cp -r dms3_release/dms3server /etc/distributed-motion-s3/
cp -r dms3_release/dms3libs /etc/distributed-motion-s3/
chown -R root.root /etc/distributed-motion-s3
rm -r dms3_release

printf "\n%s\n" "Restarting dms3server service... "
service dms3server start

printf "\n%s\n" "Success"