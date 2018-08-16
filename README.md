Descargar Docker y Docker-compose.

Descargar las imágenes mysq y golang:
```
docker pull mysql

docker pull golang
```
Modificar el archivo parameters.yml.dist y guardarlo en un achivo parameters.yaml definiendo el usuario y contraseña de la base de datos que se creará, por ejemplo:
```
  gorm:
      driver: mysql
      conn_str: golang:1234@tcp(database)/world //base de datos: world
```
Modificar el archivo .env.dist y guardarlo en un archivo .env modificando los parametros de la base de datos, por ejemplo:
```
  MYSQL_ROOT_PASSWORD=1234
  MYSQL_DATABASE=world //Fijo
  MYSQL_USER=golang
  MYSQL_PASSWORD=1234
```
Con el siguiente commando, se crearán y ejecutarán dos contenedores (uno con mysql y otro con golang):
```
Docker-compose up
```
Para probar la aplicaciòn en su cliente, ingrese a:
```
  [http://localhost:1323/geoip/8.8.8.8]
  [http://localhost:1323/geoip/maxmind/8.8.8.8]
  [http://localhost:1323/geoip/countries]
```
