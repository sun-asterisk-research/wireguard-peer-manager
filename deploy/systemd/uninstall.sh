#!/bin/bash

sudo systemctl stop wgpm.service
sudo systemctl disable wgpm.service
sudo rm /etc/systemd/system/wgpm.service /etc/systemd/system/wgpm.service.env
sudo systemctl daemon-reload
