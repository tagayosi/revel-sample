package controllers

import (
	"github.com/revel/revel"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
//	"strings"
//	"fmt"
)

type TestUser2 struct{
	Id int64 `json:"id"`
	Name string `json:"name"`
	Age int32 `json:"age"`
}

type Connectdb2 struct {
	*revel.Controller
}

func (c Connectdb2) Index() revel.Result {
	//return c.Render()
	
	//connectionString := getConnectionString()
	connectionString := "root:ada7c4g2@tcp(127.0.0.1:3306)/gormsample"
	
	
	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		revel.ERROR.Println("FATAL", err)
		panic(err)
	}
	
	db.DB()
	
	user := &TestUser2{Name: "go tarou", Age: 34}
	
	db.Create(user)
	
	ret := db.First(&user)
	
	return c.RenderJson(ret)
}


