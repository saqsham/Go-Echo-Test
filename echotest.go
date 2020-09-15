package main

import (
  "net/http"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  "encoding/json"
  "fmt"
  "io/ioutil"
  "log"
)

func main() {
  // Echo instance
  e := echo.New()

  // Middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // Routes
  e.GET("/api/GetData", jsonTest)

  // Start server
  e.Logger.Fatal(e.Start(":8080"))
}

type todos struct {
	// typedef int `json:"type"`
	UserId int 
	ID int
	Title string
	Completed bool
}

// Handler
func jsonTest(c echo.Context) error {

	url := "https://jsonplaceholder.typicode.com/todos/"

	getRequest, err := http.Get(url)
    if err != nil {
        fmt.Println("Error!")
        fmt.Println(err)
    }

    fmt.Println("The status code is", getRequest.StatusCode, http.StatusText(getRequest.StatusCode))

    // it's important to close the connection - we don't want the connection to leak
    defer getRequest.Body.Close()

    // read the body of the GET request
    rawData, err := ioutil.ReadAll(getRequest.Body)

    if err != nil {
        fmt.Println("Error!")
        fmt.Println(err)
    }
	
	todos1 := []todos{}
	jsonErr := json.Unmarshal(rawData, &todos1)

	if jsonErr != nil {
		log.Fatal(jsonErr)
  }
  
  // returing the first one as required
  return c.JSON(http.StatusOK, todos1[0])
}