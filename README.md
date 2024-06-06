## Tanya Ustadz BE

### Run App With Docker
```sh
    git clone git@github.com:underground-team-project/tanya-ustadz-be.git
    cd tanya-ustadz-be
    docker-compose up
```

### Run App Without Docker
```sh
    git clone git@github.com:underground-team-project/tanya-ustadz-be.git
    cd tanya-ustadz-be
    cp .env.example .env
    go mod tidy
    sh run-service.sh
```

### Run Migration
```sh
    sh run-migration.sh
```
### Generate Mock
```sh
    mockery --all --case underscore
```

### Run Test
```sh
    sh run-test.sh
```

### Docs Swagger
```sh
    http://localhost:8080/docs/index.html
```