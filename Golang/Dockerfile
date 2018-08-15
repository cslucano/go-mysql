FROM golang

RUN go get -u github.com/davecgh/go-spew/spew && go get -u github.com/jinzhu/configor && go get -u github.com/labstack/echo && go get github.com/oschwald/geoip2-golang && go get -u github.com/jinzhu/gorm && go get -u github.com/go-sql-driver/mysql

WORKDIR /home

CMD ["go","run","server.go"]
