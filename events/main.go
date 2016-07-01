package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Paydata dsd fucking
type PayData struct {
	Name    string
	Date    int64
	Payrate int
	Paytime int
}

func main() {

	data := PayData{Name: "Matt", Date: 1460125296, Payrate: 1000, Paytime: 80}

	//pass to aggregate
	fmt.Printf("This is the data in the main: %+v\n", data)

	url := "http://dockerhost:8001/gateway"
	fmt.Println("URL:>", url)

	jsonStr, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

}
