package controllers

import (
	"github.com/revel/revel"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"strings"
	"fmt"
)

type TestUser struct{
	Id int64 `json:"id"`
	Name string `json:"name"`
	Age int32 `json:"age"`
}

type Connectdb struct {
	*revel.Controller
}

func (c Connectdb) Index() revel.Result {
	//return c.Render()
	
	connectionString := getConnectionString()
	
	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		revel.ERROR.Println("FATAL", err)
		panic(err)
	}
	
	db.DB()
	
	user := &TestUser{Name: "go tarou", Age: 34}
	
	db.Create(user)
	
	ret := db.First(&user)
	
	return c.RenderJson(ret)
}

func getParamString(param string, defaultValue string) string {
    p, found := revel.Config.String(param)
    if !found {
        if defaultValue == "" {
            revel.ERROR.Fatal("Cound not find parameter: " + param)
        } else {
            return defaultValue
        }
    }
    return p
}

func getConnectionString() string {
    host := getParamString("db.host", "")
    port := getParamString("db.port", "3306")
    user := getParamString("db.user", "")
    pass := getParamString("db.password", "")
    dbname := getParamString("db.name", "auction")
    protocol := getParamString("db.protocol", "tcp")
    dbargs := getParamString("dbargs", " ")

    if strings.Trim(dbargs, " ") != "" {
        dbargs = "?" + dbargs
    } else {
        dbargs = ""
    }
    return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s%s", 
        user, pass, protocol, host, port, dbname, dbargs)
}
