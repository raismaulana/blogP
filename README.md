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
Or using docker (development)
```
git clone http://gitub.com/raismaulana/blogP
cd blogP
go mod tidy
go mod vendor
docker build -t blogp:1.0.0 .
docker container create --name blogp -p 8080:8080 blogp:1.0.0
docker start -i blogp
```
*NB: copy the config.yaml.example to config.yaml and edit the value based on your environment*
