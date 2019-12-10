#!/bin/bash
# Author:       nghiatc
# Email:        congnghia0609@gmail.com

source /etc/profile

echo "Install library dependencies..."
go get -u github.com/tools/godep
go get -u github.com/spf13/viper
go get -u github.com/congnghia0609/ntc-gconf
go get -u github.com/garyburd/redigo/redis

echo "Install dependencies complete..."
