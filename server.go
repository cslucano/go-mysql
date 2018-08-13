package main

import (
    "log"
    "net"
    "net/http"
    "github.com/labstack/echo"
    "github.com/jinzhu/gorm"
    "github.com/oschwald/geoip2-golang"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Country struct {
    Code string `gorm:"primary_key"`
    Name string
}

func (c Country) TableName() string {
    return "country"
}


func main() {
    e := echo.New()
    e.GET("/geoip/:ip", func(c echo.Context) error {
        ip := net.ParseIP(c.Param("ip"))
        if nil == ip {
            return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

        db, err := geoip2.Open("./GeoLite2-DB/GeoLite2-Country.mmdb")
        if err != nil {
            log.Fatal(err)
        }
        defer db.Close()
        record, err := db.Country(ip)
        if err != nil {
            log.Fatal(err)
        }
        

        return c.String(http.StatusOK, record.Country.IsoCode)
    })	

    e.GET("/geoip/countries", func(c echo.Context) error {
        db, err := gorm.Open("mysql", "root:1234@/world")
        if err != nil {
            return echo.NewHTTPError(http.StatusBadRequest, "DB connection error")
        }
        defer db.Close()
        
        var country Country
        db.First(&country)
        
        return c.String(http.StatusOK, "Hola " + country.Code) 
    })

    e.Logger.Fatal(e.Start(":1323"))
}

