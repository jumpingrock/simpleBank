docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret postgres:15-alpine

Access from terminal, do not require password as accessing from local host
docker exec -it postgres15 psql -U root

run ssh into docker postgres using bin/sh shell
docker exec -it postgres15 /bin/sh
    - once inside can use the command below to create db and assign owner
    - createdb --username=root --owner=root simple_bank

alternatively we can use below to avoid going into psql db
    - dicjer exec -it postgres15 createdb -username=root simple_bank