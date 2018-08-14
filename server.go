package main

import (
    "log"
    "net"
    "net/http"
    "github.com/davecgh/go-spew/spew"
    "github.com/labstack/echo"
    "github.com/jinzhu/configor"
    "github.com/jinzhu/gorm"
    "github.com/oschwald/geoip2-golang"

    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Country struct {
    Code string `gorm:"column:Code; primary_key:yes"`
    Name string `gorm:"column:Name;"`
    Continent string `gorm:"column:Continent;"`
    Region string `gorm:"column:Region;"`
    SurfaceArea float32 `gorm:"column:SurfaceArea;"`
    IndepYear int `gorm:"column:IndepYear;"`
    Population int `gorm:"column:Population;"`
    LifeExpectancy float32 `gorm:"column:LifeExpectancy;"`
    GNP float32 `gorm:"column:GNP;"`
    GNPOld float32 `gorm:"column:GNPOld;"`
    LocalName string `gorm:"column:LocalName;"`
    GovernmentForm string `gorm:"column:GovernmentForm;"`
    HeadOfState string `gorm:"column:HeadOfState;"`
    Capital int `gorm:"column:Capital;"`
    Code2 string `gorm:"column:Code2;"`
}

func (c Country) TableName() string {
    return "country"
}

var parameters =  struct {
    Gorm struct {
        Driver string `yaml:"driver"`
        ConnStr string `yaml:"conn_str"`
    }
}{}

func main() {
    configor.Load(&parameters, "parameters.yml")
    spew.Dump(parameters)
   
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
	    db, err := gorm.Open(parameters.Gorm.Driver, parameters.Gorm.ConnStr)/*?charset=utf8&parseTime=True&loc=Local")*/
        if err != nil {
            return echo.NewHTTPError(http.StatusBadRequest, "DB connection error")
        }
	
	defer db.Close()

        db.LogMode(true)
	
	var country Country
        db.Debug().First(&country)
        spew.Dump(country)

	return c.String(http.StatusOK, country.Name) 
    })

    e.Logger.Fatal(e.Start(":1323"))
}

