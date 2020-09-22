#!/usr/bin/env bash
#wait for the MySQL Server to come up
#sleep 90s

#run the setup script to create the DB and the schema in the DB
# specify password -p command ex) password=sa -> -psa
mysql -u sa -psa test_database < "/docker-entrypoint-initdb.d/0_schema.sql"
mysql -u sa -psa test_database < "/docker-entrypoint-initdb.d/1_testdata.sql"