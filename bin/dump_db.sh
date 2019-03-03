#!/bin/sh

set -eu

mysqldump -u root -p --column-statistics=0 -h 127.0.0.1 test -d
