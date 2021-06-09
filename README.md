# go-crash
Project **go-crash** is just my go lang practices<br/>
<img src="./material/go.svg" width="300">

### Quick Start
Follow the following steps:
- Clone this repository
- Then Run: `go mod tidy && go mod vendor`
- Then Run: `go run main.go`
- Open your browser on http://localhost:8080

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
