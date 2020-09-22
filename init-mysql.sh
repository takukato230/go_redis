#!/bin/sh
docker-compose -f ./developments/docker-compose.yml exec mysql bash -c "chmod 0775 docker-entrypoint-initdb.d/init-database.sh"
docker-compose -f ./developments/docker-compose.yml exec mysql bash -c "./docker-entrypoint-initdb.d/init-database.sh"