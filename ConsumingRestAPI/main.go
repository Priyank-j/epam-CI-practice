package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Employee struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Address  Address
	Phone    string `json:"phone"`
	Website  string `json:"wbsite"`
	Company  Company
}

type Address struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"City"`
	ZopCode string `json:"zipcode"`
	Geo     Geo
}

type Geo struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}
type Company struct {
	Name       string `json:"name"`
	CatchPrase string `json:"catchPrase"`
	Bs         string `json:"bs"`
}

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/users", nil)

	if err != nil {
		fmt.Println("Not able to access ", err.Error())
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	response, err := client.Do(req)

	if err != nil {
		fmt.Println("Not able to fetch data", err.Error())
	}

	var responseObject []Employee
	defer response.Body.Close()

	respBytes, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Println("Unable to read the response", err.Error())
	}

	json.Unmarshal(respBytes, &responseObject)

	fmt.Println("Responce from given API", responseObject)

}
