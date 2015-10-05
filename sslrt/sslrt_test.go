package sslrt

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestHTTPS(t *testing.T) {
	transport := &OpenSSLTransport{}
	req, _ := http.NewRequest("GET", "https://google.com", nil)
	res, _ := transport.RoundTrip(req)
	body, _ := ioutil.ReadAll(res.Body)

	res.Body.Close()

	fmt.Println(string(body))
}
