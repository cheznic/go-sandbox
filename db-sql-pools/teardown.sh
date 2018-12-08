#!/usr/bin/env bash
unset MYSQLUSERNAME
unset MYSQLPASSWORD

mysql -uroot -ppassword
DROP DATABASE gocookbook;
exit
