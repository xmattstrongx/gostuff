package main

import (
	"fmt"
	"math/rand"
	"time"

	ll "log"

	kit "forks/kit/sd/etcd"

	"github.com/go-kit/kit/log"
	"golang.org/x/net/context"
)

func main() {
	// cfg := client.Config{
	// 	Endpoints: []string{"http://127.0.0.1:2379"},
	// }
	// c, err := client.New(cfg)
	// if err != nil {
	// 	// handle error
	// }
	// kAPI := client.NewKeysAPI(c)

	rand.Seed(int64(time.Now().Nanosecond()))
	value := rand.Int()
	testData := fmt.Sprintf("foo%d", value)
	testData2 := fmt.Sprintf("bar%d", value)

	// create a new key /foo with the value "bar"
	// f, err := kAPI.Create(context.Background(), testData, "bar")
	// if err != nil {
	// 	log.Panicf("Fucking error %+v", err)
	// }
	// log.Println(f)

	kitClientOptions := kit.ClientOptions{
		Cert:                    "",
		Key:                     "",
		CaCert:                  "",
		DialTimeout:             (2 * time.Second),
		DialKeepAline:           (2 * time.Second),
		HeaderTimeoutPerRequest: (2 * time.Second)}

	kitClient, err := kit.NewClient(context.Background(), []string{"http://:2379"}, kitClientOptions)

	instance := &kit.Service{
		Key:   testData,
		Value: "poop1",
	}
	instance2 := &kit.Service{
		Key:   testData2,
		Value: "poop2",
	}

	registrar := kit.NewRegistrar(kitClient, *instance, log.NewNopLogger())
	registrar2 := kit.NewRegistrar(kitClient, *instance2, log.NewNopLogger())
	registrar.Register()
	registrar2.Register()

	// err = kitClient.Register(instance3)
	// if err != nil {
	// 	ll.Panicf("Fucking shit err = %+v", err)
	// }

	a, err := kitClient.GetEntries(testData)
	if err != nil {
		ll.Panicf("Fucking poop err = %+v", err)
	}
	ll.Println(a)
	b, err := kitClient.GetEntries(testData2)
	if err != nil {
		ll.Panicf("Fucking poop err = %+v", err)
	}
	ll.Println(b)

	//delete registrar
	registrar.Deregister()

	_, err = kitClient.GetEntries(testData)
	if err != nil {
		ll.Printf("Fucking poop err = %+v", err)
	}

	//delete registrar2
	registrar2.Deregister()

	_, err = kitClient.GetEntries(testData2)
	if err != nil {
		ll.Printf("Fucking poop err = %+v", err)
	}

}
