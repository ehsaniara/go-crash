# go-crash

Project **go-crash** is just my go lang practices

### Stacks
- GoLang (with: `gin-gonic`, `viper`, `zap`, `swaggo`)
- Postgres Db
- Redis
- Docker

<img src="./material/go.svg" width="300">

### Quick Start

If you already have Postgres and Redis running in your machine, just configure the `config/config.yml` file for credentials follow the steps:

- Clone this repository
- Then Run: `go mod tidy && go mod vendor`
- Then Run: `go run main.go`
- Or Run: `export PROFILE=Default && go run main.go` with selected profile
- Open your browser on http://localhost:8080

#### Note:
The `dev`, `docker` and `prod` profiles (environments)'s credentials are get passes from OS env-vars. Let say you want to run this project in kubernetes then, you should pass the environment variables in your deployment manifest. for example:
```yaml
      image: path-to-image-of-this-project-in-repo
      env:
        - name: PROFILE
          value: "dev"
        - name: DATABASE_USER
          value: "postgres"
        - name: DATABASE_PASSWORD
          value: "postgres"
        - name: DATABASE_DB
          value: "go_crash"
        - name: REDIS_PASSWORD
          value: "secret"
```
_the values can be later replaced into kube-secret or any other third party vault applications_


check if server is working:

```shell
curl --location --request GET 'http://localhost:8080/ping'
```

Api Documentations are in http://localhost:8080/swagger/index.html

# Docker Run

run the project in docker, It also has images for:

- Redis
- Postgres

## Start the docker compose

```shell
docker-compose up -d
```

## clean the docker compose

```shell
docker-compose down -v
```

# Quick Start

login to postgres and create a record in `auths` table to create authentication user for token

```sql
insert into public.auths (id, username, password)
values (1, 'test', 'test');
```

basically eny api under `/api/v1` is protected by `Token` header and needs to be authenticated.

to authenticate call the following command:

```shell
curl --location --request POST 'http://localhost:8080/auth' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "test",
    "password": "test"
}'
```

now you have your token, you should use it in every call under `/api/v1` as header `Token`

### Create User

```shell
curl --location --request POST 'http://localhost:8080/api/v1/customers' \
--header 'token: eyJhbGciOiJI...GSqQhG8' \
--header 'Content-Type: application/json' \
--data-raw '{
    "firstName": "Jay",
    "lastName": "Ehsaniara",
    "title": "Software Engineer"
}'
```

And the response from the following call will be:

```json
{
  "ID": 1,
  "FirstName": "Jay",
  "LastName": "Ehsaniara",
  "Title": "Software Engineer",
  "CreatedBy": "test",
  "CreatedOn": 1623276844,
  "ModifiedOn": 1623276844
}
```

### Get User

```shell
curl --location --request GET 'http://localhost:8080/api/v1/customers/1' --header 'token: eyJhbGciOiJI...GSqQhG8'
```

The first call gets the data from postgres and store it in your Redis cache ans show it as:

```json
{
  "ID": 1,
  "FirstName": "Jay",
  "LastName": "Ehsaniara",
  "Title": "Engineer",
  "CreatedBy": "test",
  "CreatedOn": 1623276844,
  "ModifiedOn": 1623276844
}
```

Latter calls for the same customer id will be much faster because it's calling the redis from now on
