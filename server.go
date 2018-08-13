
package main

import (
    "log"
    "fmt"
    "net"
    "net/http"
    "github.com/labstack/echo"
    "github.com/jinzhu/gorm"
    "github.com/oschwald/geoip2-golang"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "github.com/davecgh/go-spew/spew"
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
	    db, err := gorm.Open("mysql", "golang:golang@/world")/*?charset=utf8&parseTime=True&loc=Local")*/
        if err != nil {
            return echo.NewHTTPError(http.StatusBadRequest, "DB connection error")
        }
	
	defer db.Close()

        db.LogMode(true)
	
	var country Country
	//fmt.Println(db.Raw("SELECT * FROM country ORDER BY Code LIMIT 1").Scan(&country))
        db.Debug().First(&country)
        spew.Dump(country)
	fmt.Println(country)
	return c.String(http.StatusOK, country.Name) 
    })

    e.Logger.Fatal(e.Start(":1323"))
}

