# blogP
## Simple blog app RESTful API
### built with golang gogen gin-gonic gorm casbin postgresql redis
[documentation/openapi3](https://petstore.swagger.io/?url=https://raw.githubusercontent.com/raismaulana/blogP/main/docs/openapi.yaml)
#
## How to install?
```
git clone http://gitub.com/raismaulana/blogP
cd blogP
go mod tidy
go run main.go usingdb
```
Or using docker
```
git clone http://gitub.com/raismaulana/blogP
cd blogP
go mod tidy
go mod vendor
go env -w CGO_ENABLED=0
go build -o build/package/bin3 .
go env -w CGO_ENABLED=1
docker network create blogp-network
docker volume create blogp_volume
docker build -t blogp:latest .
docker run --name blogp-postgres -v blogp_volume:/postgres --network blogp-network -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=blog -d postgres:latest
docker run --name blogp-redis -id -p 6379:6379 --network blogp-network redis:latest redis-server --requirepass "redis"
docker run --name blogp -p 8080:8080 --network blogp-network blogp:latest
```
*NB: copy the config.yaml.example to config.yaml and edit the value based on your environment*
