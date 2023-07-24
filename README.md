# golang-restapi


## set-up
set up mysql by docker 
docker run --name test-mysql -e MYSQL_ROOT_PASSWORD=12345 -p 3306:3306 -d mysql
create schema golang
using file script.sql in folder db 

### run project
go install 
go run main.go 