## go-rest-api

### When using Docker

    docker build -t go-rest-api-v2 .

    docker run go-rest-api-v2

### When using Docker Commpose

    docker-compose up

### When using task

    task build // to build the project

    task lint // run lint

    task test // run tests

    task run // to run the service

### Checking the database container for existance of the table

    docker container list // copy the container id of the postgres container

    docker exec -it <container_id> bash // this will set you up in the bash terminal of the db container

    ** In bash mode **

    > psql -U postgres

    > \dt // here you should see both the tables  - comments and schema migrations

    >  \d+ comments; // details schema view of the comments table

### Testing Integration

    > go test -tags=integration ./... -v



