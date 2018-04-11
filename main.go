package main

import (
	"golang-firestore/config"
	"golang-firestore/middleware"
	"golang-firestore/modules/customer"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

/**
 * @brief      main
 *
 * @return     array
 */
func main() {
	// Logging to a file.
	current_time := time.Now().Local()
	current_date := current_time.Format("2006-01-02")
	f, _ := os.OpenFile("logs/"+current_date+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	gin.DefaultWriter = io.MultiWriter(f)
	route := gin.Default()
	customercontroller := customer.CustomerController()
	route.Use(middleware.ApiAuthentication())
	{
		route.POST("api/v1/firestore/customer", customercontroller.SaveDataCustomer)
	}
	route.Run(":" + config.Env("APP_PORT"))
}
