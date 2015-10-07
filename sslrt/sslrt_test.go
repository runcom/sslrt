package sslrt

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestHTTPS(t *testing.T) {
	transport := &openSSLTransport{}
	req, _ := http.NewRequest("GET", "https://google.com", nil)
	res, _ := transport.RoundTrip(req)
	body, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", string(body))
}
