package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

func main() {

     // Connect MySQL at 127.0.0.1:3306, with user root, an empty passowrd and database test

    db, err := sql.Open("mysql", "root:root@/splunk")
	
	if err != nil {
      panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
    }
    //defer db.Close()


    router := gin.Default()
    router.GET("/", func(c *gin.Context) {

    // Prepare statement for inserting data
    stmtIns, err := db.Prepare("insert into views (text) values ('from golang')") 
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

      stmtIns.Exec()

        c.String(http.StatusOK, "hello world")
    })
    router.GET("/ping", func(c *gin.Context) {
        c.String(http.StatusOK, "pong")
    })
    router.POST("/submit", func(c *gin.Context) {
        c.String(http.StatusUnauthorized, "not authorized")
    })
    router.PUT("/error", func(c *gin.Context) {
        c.String(http.StatusInternalServerError, "and error happened :(")
    })
    router.Run(":5000")
}