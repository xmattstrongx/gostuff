package main

import (
	"fmt"
)

http://localhost:8000/hello/matt

func main() {
    customer, err := FindCustomerByEmail("frederik@gopher.com")
    if err != nil {
        fmt.Println("No customer found")
        //do something else
    }
	
}


func FindCustomerByEmail(email string) (customer Customer, err error) {
	client := CustomerClient()
	searchResult, err := client.Search().
		Index("customer").
		Query(elastic.NewTermQuery("email", email)).
		From(0).Size(1).
		Do()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Searchresult from elasticsearch: %+v", searchResult)

	return customer, fmt.Errorf("Customer with email %s not found", email)
}