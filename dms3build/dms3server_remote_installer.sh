#!/bin/bash

# this bash script will be transferred to the dms3 device component platform, executed, and
# then deleted automatically

# IMPORTANT: be sure to set root password access before using

PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
PASS="password"

printf "\n%s\n" "Stopping dms3server service... "
echo $PASS | sudo -S service dms3server stop

printf "\n%s\n" "Moving files into /usr/local/bin... "
echo $PASS | sudo -S mv dms3_release/linux_amd64/go_dms3server /usr/local/bin
echo $PASS | sudo -S chown root.root /usr/local/bin/go_dms3server

printf "\n%s\n" "Moving files into /etc/distributed-motion-s3... "
echo $PASS | sudo -S mkdir -p /etc/distributed-motion-s3
echo $PASS | sudo -S cp -r dms3_release/dms3server /etc/distributed-motion-s3/
echo $PASS | sudo -S cp -r dms3_release/dms3libs /etc/distributed-motion-s3/
echo $PASS | sudo -S chown -R root.root /etc/distributed-motion-s3
echo $PASS | sudo -S rm -r dms3_release

printf "\n%s\n" "Restarting dms3server service... "
echo $PASS | sudo -S service dms3server start

printf "\n%s\n" "Success"