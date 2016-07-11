package main

import (
	"fmt"
	"math/rand"
	"time"
    "log"

	kit "github.com/go-kit/kit/sd/etcd" 
	kitlog "github.com/go-kit/kit/log"
	"golang.org/x/net/context"
)

// Package sd/etcd provides a wrapper around the coroes/etcd key value store (https://github.com/coreos/etcd) 
// This example assumes the user has an instance of etcd installed and running locally on port 2379
func main() {

	// Creating random number too append to assure differentiation of test data keys 
	rand.Seed(int64(time.Now().Nanosecond()))

	value := rand.Int()
	testData := fmt.Sprintf("foo%d", value)

	value2 := rand.Int()
	testData2 := fmt.Sprintf("foo%d", value2)


	// Using basic options for NewClient
	kitClientOptions := kit.ClientOptions{
		DialTimeout:             (2 * time.Second),
		DialKeepAline:           (2 * time.Second),
		HeaderTimeoutPerRequest: (2 * time.Second),
	}

	// Instantiate NewClient on localhost port 2379. Change this to whichever port you host etcd from 
	kitClient, err := kit.NewClient(context.Background(), []string{"http://:2379"}, kitClientOptions)

	// Set up first instance of test data
	data1 := &kit.Service{
		Key:   testData,
		Value: "bar1",
	}

	// Instantiate new instance of *Registrar passing in test data
	registrar := kit.NewRegistrar(kitClient, *data1, kitlog.NewNopLogger())
	// Register new test data to etcd
	registrar.Register()

	// Set up second instance of test data
	data2 := &kit.Service{
		Key:   testData2,
		Value: "bar2",
	}

	// Instantiate new instance of *Registrar passing in test data
	registrar2 := kit.NewRegistrar(kitClient, *data2, kitlog.NewNopLogger())
	// Register new test data to etcd
	registrar2.Register()

	//Retrieve entries from etcd
	_, err = kitClient.GetEntries(testData)
	if err != nil {
		fmt.Println(err)
	}

log.Println("Matt")
	_, err = kitClient.GetEntries(testData2)
	if err != nil {
        fmt.Println(err)
	}

	// Deregister first instance of test data
	registrar.Deregister()

	// Verify test data no longer exists in etcd
	_, err = kitClient.GetEntries(testData)
	if err != nil {
		fmt.Println(err)
	}

	// Deregister second instance of test data
	registrar2.Deregister()

	// Verify test data no longer exists in etcd
	_, err = kitClient.GetEntries(testData2)
	if err != nil {
		fmt.Println(err)
	}
}
