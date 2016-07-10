package main

import (
	"fmt"
	"math/rand"
	"time"
	"os"
	"strings"

	ll "log"

	kit "forks/kit/sd/etcd"

	"github.com/go-kit/kit/log"
	"golang.org/x/net/context"
)

// Package sd/etcd provides a wrapper around the coroes/etcd key value store (https://github.com/coreos/etcd) 
// This example assumes the user has an instance of etcd installed and running locally on port 2379
func main() {

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stdout)
		logger = log.NewContext(logger).With("ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)
	}

	// Creating random number too append to test data
	rand.Seed(int64(time.Now().Nanosecond()))
	value := rand.Int()

	testData := fmt.Sprintf("foo%d", value)
	testData2 := fmt.Sprintf("foo%d", value)


	// Using basic options for NewClient
	kitClientOptions := kit.ClientOptions{
		Cert:                    "",
		Key:                     "",
		CaCert:                  "",
		DialTimeout:             (2 * time.Second),
		DialKeepAline:           (2 * time.Second),
		HeaderTimeoutPerRequest: (2 * time.Second)}

	// Instantiate NewClient on localhost port 2379. Change this to whichever port you host etcd from 
	kitClient, err := kit.NewClient(context.Background(), []string{"http://:2379"}, kitClientOptions)

	// Set up first instance of test data
	data1 := &kit.Service{
		Key:   testData,
		Value: "bar1",
	}

	// Instantiate new instance of *Registrar passing in test data
	registrar := kit.NewRegistrar(kitClient, *data1, log.NewNopLogger())
	// Register new test data to etcd
	registrar.Register()

	// Set up second instance of test data
	data2 := &kit.Service{
		Key:   testData2,
		Value: "bar2",
	}

	// Instantiate new instance of *Registrar passing in test data
	registrar2 := kit.NewRegistrar(kitClient, *data2, log.NewNopLogger())
	// Register new test data to etcd
	registrar2.Register()

	//Retrieve entries from etcd
	result1, err := kitClient.GetEntries(testData)
	if err != nil {
		logger.Log("err", err)
	}
	// ll.Println(a)
	// logger.Log("entries", a)
	logger.Log("entries", strings.Join(result1, ", "))

	result2, err := kitClient.GetEntries(testData2)
	if err != nil {
		logger.Log("err", err)
	}
	logger.Log("entries", strings.Join(result2, ", "))
	// ll.Println(b)

	// Deregister first instance of test data
	registrar.Deregister()

	// Verify test data no longer exists in etcd
	_, err = kitClient.GetEntries(testData)
	if err != nil {
		logger.Log("err", err)
	}

	// Deregister second instance of test data
	registrar2.Deregister()

	// Verify test data no longer exists in etcd
	_, err = kitClient.GetEntries(testData2)
	if err != nil {
		logger.Log("err", err)
	}

}

// Example Output
// ts=2016-07-10T19:22:57Z caller=main.go:76 entries=bar1
// ts=2016-07-10T19:22:57Z caller=main.go:81 entries="unsupported value type"
// ts=2016-07-10T19:22:57Z caller=main.go:90 err="100: Key not found (/foo8197187502305182722) [57]"
// ts=2016-07-10T19:22:57Z caller=main.go:99 err="100: Key not found (/foo8197187502305182722) [57]"