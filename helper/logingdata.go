package helper

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

/**
 * @brief      Gets the current date.
 * @return     The current date.
 */
func getCurrentDate() string {
	current_time := time.Now().Local()
	return current_time.Format("2006-01-02")
}

/**
 * @brief      this function for logging data request
 * @param      request  The request
 * @return     object
 */
func LoggingDataRequest(request *gin.Context) {
	current_date := getCurrentDate()
	f, err := os.OpenFile("logs/"+current_date+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	gin.DefaultWriter = io.MultiWriter(f)
	log.Println(*request)
}

/**
 * @brief      this function for logging error data
 * @param      error_data  The error data
 * @param      from        The from
 * @return     object
 */
func LoggingError(error_data error, from string) {
	current_date := getCurrentDate()
	f, err := os.OpenFile("logs/"+current_date+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	gin.DefaultWriter = io.MultiWriter(f)
	log.Println("FROM => " + from)
	log.Println(error_data)
}

/**
 * @brief      this function for logging error fatal
 * @param      error_data  The error data
 * @return     object
 */
func LoggingErrorFatal(error_data error) {
	log.Println("LoggingErrorFatal")
	log.Fatal(error_data)
}
