package sslrt_test

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/runcom/sslrt/sslrt"
)

func init() {
	http.DefaultTransport = sslrt.NewOpenSSLTransport(nil)
}

func ExampleOpenSSLTransport() {
	res, err := http.Get("https://google.com")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	c, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", string(c))
}
