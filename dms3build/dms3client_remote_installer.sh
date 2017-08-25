#!/bin/bash

# this bash script will be copied to the dms3 device component platform, executed, and
# then deleted automatically

PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin

PASS="!PASSWORD"  # DO NOT ALTER: this is automatically replaced during dms3 component installation

printf "\n%s\n" "Stopping dms3client.service service... "
echo $PASS | sudo -S systemctl stop dms3client.service

printf "\n%s\n" "Moving files into /usr/local/bin... "
echo $PASS | sudo -S mv dms3_release/linux_arm7/go_dms3client /usr/local/bin
echo $PASS | sudo -S mv dms3_release/linux_arm7/go_dms3mail /usr/local/bin
echo $PASS | sudo -S chown root.root /usr/local/bin/go_dms3*

printf "\n%s\n" "Moving files into /etc/distributed-motion-s3... "
echo $PASS | sudo -S mkdir -p /etc/distributed-motion-s3
echo $PASS | sudo -S cp -r dms3_release/dms3client /etc/distributed-motion-s3
echo $PASS | sudo -S cp -r dms3_release/dms3libs /etc/distributed-motion-s3

# I don't really want to overwrite my mail settings every time, do I?
# echo $PASS | sudo -S cp -r dms3_release/dms3mail /etc/distributed-motion-s3

echo $PASS | sudo -S chown -R root.root /etc/distributed-motion-s3
echo $PASS | sudo -S rm -r dms3_release

printf "\n%s\n" "Restarting dms3client.service systemd service... "
echo $PASS | sudo -S systemctl start dms3client.service

printf "\n%s\n" "Success"