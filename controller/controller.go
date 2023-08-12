package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/izasoerya/Prakerja-Go/models"
	"github.com/labstack/echo/v4"
)

func Siji(c echo.Context) error {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var data interface{}
	if json.NewDecoder(resp.Body).Decode(&data) != nil {
		return err
	}
	return c.JSON(resp.StatusCode, data)
}

func Loro(c echo.Context) error {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/3")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var data interface{}
	if json.NewDecoder(resp.Body).Decode(&data) != nil {
		return err
	}
	return c.JSON(resp.StatusCode, data)
}

func Telu(c echo.Context) error {
	mydata := &models.Data{
		UserID: 100,
		ID:     100,
		Title:  "I am become death",
		Body:   "Robert Oppenheimer",
	}
	jsonData, err := json.Marshal(mydata)
	if err != nil {
		fmt.Println(err)
	}

	url := "https://jsonplaceholder.typicode.com/posts"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error sending request:", err)
		return err
	}
	defer resp.Body.Close()

	var responseData models.Data
	err = json.NewDecoder(resp.Body).Decode(&responseData)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, responseData)
}

func Papat(c echo.Context) error {
	url := "https://jsonplaceholder.typicode.com/posts/1"
	_, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err
	}

	return c.String(http.StatusOK, "Deleted successfully")
}


