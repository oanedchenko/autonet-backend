## database setup:

* install arangodb docker container:
```
docker pull arangodb/arangodb
docker run --name arango -e ARANGO_ROOT_PASSWORD=123 -p 8529:8529 -d arangodb
```

* copy json data files and import scripts to the container:
```
docker cp import.sh arango:/root
```

* get into the container and run data import script:
```
docker exec -it arango sh
cd /root
./import.sh
```

# run DAO tests

Setup required environment variables before test:

* DB_URL=http://localhost:8529
* DB_USER=root
* DB_PASSWORD=123
* DB_NAME=autonet
