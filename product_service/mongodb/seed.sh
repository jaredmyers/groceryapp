#!/bin/sh
mongoimport --db "${MONGO_INITDB_DATABASE}" --collection products --file /docker-entrypoint-initdb.d/products.json
