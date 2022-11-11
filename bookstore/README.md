# Instructions to run

- Start docker (install docker)[https://docs.docker.com/get-docker/]
- Run the command `docker run --name mysqldb -p 3306:3306 -e MYSQL_ROOT_PASSWORD=Something -e MYSQL_DATABASE=osample -e MYSQL_USER=fseg -e MYSQL_PASSWORD=Something -d mysql`
- Run the application `go run cmd/main/main.go`
