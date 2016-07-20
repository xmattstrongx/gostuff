package main

import (
	"errors"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/go-kit/kit/endpoint"
	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	kit "github.com/go-kit/kit/sd/etcd"
	"golang.org/x/net/context"
)

// Package sd/etcd provides a wrapper around the coroes/etcd key value store (https://github.com/coreos/etcd)
// This example assumes the user has an instance of etcd installed and running locally on port 2379
func main() {

	// Creating random number too append to assure differentiation of test data keys
	var (
		prefix   = "/services/foosvc/" // known at compile time
		instance = "1.2.3.4:8080"      // taken from runtime or platform, somehow
		key      = prefix + instance
		value    = "http://" + instance // based on our transport
	)

	// Using basic options for NewClient
	kitClientOptions := kit.ClientOptions{
		DialTimeout:             (2 * time.Second),
		DialKeepAline:           (2 * time.Second),
		HeaderTimeoutPerRequest: (2 * time.Second),
	}

	// Instantiate NewClient on localhost port 2379. Change this to whichever port you host etcd from
	kitClient, err := kit.NewClient(context.Background(), []string{"http://:2379"}, kitClientOptions)

	// Instantiate new instance of *Registrar passing in test data
	_ = kit.NewRegistrar(kitClient, kit.Service{
		Key:   key,
		Value: value,
	}, kitlog.NewNopLogger())
	// Register new test data to etcd
	// registrar.Register()

	// // Set up second instance of test data
	// data2 := &kit.Service{
	// 	Key:   testData2,
	// 	Value: "bar2",
	// }

	// // Instantiate new instance of *Registrar passing in test data
	// registrar2 := kit.NewRegistrar(kitClient, *data2, kitlog.NewNopLogger())
	// // Register new test data to etcd
	// registrar2.Register()

	//Retrieve entries from etcd
	_, err = kitClient.GetEntries(key)
	if err != nil {
		fmt.Println(err)
	}

	entriess, err := kitClient.GetEntries(key)
	fmt.Printf("%q (%v)\n", strings.Join(entriess, ", "), err)

	// _, err = kitClient.GetEntries(testData2)
	// if err != nil {
	//     fmt.Println(err)
	// }

	// factory := func(string) (endpoint.Endpoint, io.Closer, error) {
	// endpoints  = make([]endpoint.Endpoint, n)
	// 	return endpoint.Nop, nil, nil
	// }

	// factory := func(instance string) (endpoint.Endpoint, io.Closer, error) {
	// 	// t.Logf("factory invoked for %q", instance)
	// 	return service.Fixed{}, nil, nil
	// }

	// someFactory := func(instance string) (endpoint.Endpoint, io.Closer, error) {
	// 	return endpoint.Nop, nopCloser{}, nil
	// }

	factory := func(string) (endpoint.Endpoint, io.Closer, error) {
			return endpoint.Nop, nil, nil
		}

	// 	ca    := make(closer)
	// 	cb    := make(closer)
	// c     := map[string]io.Closer{"a": ca, "b": cb}
	// f     := func(instance string) (endpoint.Endpoint, io.Closer, error) { return endpoint.Nop, c[instance], nil }

	// subscriber, _ := kit.NewSubscriber(kitClient, prefix, newFactory(instance), kitlog.NewNopLogger())
	// subscriber, _ := kit.NewSubscriber(kitClient, prefix, f, kitlog.NewNopLogger())
	subscriber, _ := kit.NewSubscriber(kitClient, prefix, factory, kitlog.NewNopLogger())

	a, err := subscriber.Endpoints()
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	// a[0].Close
	fmt.Println(len(a)) // hopefully 1
	// fmt.Println(a[0])
	// And now we can deregister and perform everything again, in reverse.

	// registrar.Register()

	// registrar.Deregister()
	// subscriber.Stop()

	// time.Sleep(2 * time.Second)

	b, err := subscriber.Endpoints()
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	fmt.Println(len(b)) // hopefully 0
	// fmt.Println(b[0])

	entries, err := kitClient.GetEntries(key)
	fmt.Printf("%q (%v)\n", strings.Join(entries, ", "), err)

	// Deregister first instance of test data
	// registrar.Deregister()

	// // Verify test data no longer exists in etcd
	// _, err = kitClient.GetEntries(key)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// // Deregister second instance of test data
	// registrar2.Deregister()

	// // Verify test data no longer exists in etcd
	// _, err = kitClient.GetEntries(testData2)
	// if err != nil {
	// 	fmt.Println(err)
	// }
}

func testFactory(instance string) (endpoint.Endpoint, io.Closer, error) {
	return func(context.Context, interface{}) (interface{}, error) {
		return instance, nil
	}, nil, nil
}

type nopCloser struct{}

func (nopCloser) Close() error { return nil }

func newFactory(fakeError string) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		fmt.Printf("instance %s\n", instance)
		fmt.Printf("fakeError %s\n", fakeError)

		if fakeError == instance {
			fmt.Println("Matt1")
			return nil, nil, errors.New(fakeError)
		}
		fmt.Println("Matt2")
		// return endpoint.myEndpoint, nil, nil
		return endpoint.Chain(annotate(instance))(myEndpoint), nil, nil
	}
}

func annotate(s string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			fmt.Println(s, "pre")
			defer fmt.Println(s, "post")
			return next(ctx, request)
		}
	}
}

func myEndpoint(context.Context, interface{}) (interface{}, error) {
	fmt.Println("my endpoint!")
	return struct{}{}, nil
}

type closer chan struct{}

func (c closer) Close() error {
	close(c)
	return nil
}
