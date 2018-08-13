Descargar mysql Docker Container

docker pull mysql
docker run --name mysql -e MYSQL_ROOT_PASSWORD=1234 -p 3306:3306 -d mysql

Instalar dependencias
go get github.com/oschwald/geoip2-golang
go get -u github.com/jinzhu/gorm
go get -u github.com/go-sql-driver/mysql

Ejecutar App

go run server.go


