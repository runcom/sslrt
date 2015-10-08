package sslrt

import (
	"io/ioutil"
	"net/http"
	"testing"
)

type aTransport struct{}

func (t aTransport) RoundTrip(req *http.Request) (*http.Response, error) { return nil, nil }

func TestNewOpenSSLTransport(t *testing.T) {
	transport := NewOpenSSLTransport(nil)
	defaultTransport := transport.(*OpenSSLTransport).defaultTransport
	if defaultTransport != http.DefaultTransport {
		t.Fatalf("Want http.DefaultTransport transport, got %v", defaultTransport)
	}

	tr := aTransport{}
	transport = NewOpenSSLTransport(tr)
	defaultTransport = transport.(*OpenSSLTransport).defaultTransport
	if defaultTransport != tr {
		t.Fatalf("Want aTransport transport, got %v", defaultTransport)
	}
}

func TestRoundTripNilURL(t *testing.T)                             {}
func TestRoundTripNilHeader(t *testing.T)                          {}
func TestRoundTripShortCircuitToHTTPDefaultTransport(t *testing.T) {}
func TestRoundTripShortCircuitToDefaultTransport(t *testing.T)     {}
func TestRoundTripEmptyHost(t *testing.T)                          {}
func TestRoundTripErrConnection(t *testing.T)                      {}
func TestCanonicalAddressFromURL(t *testing.T)                     {}
func TestRoundTripToHTTPS(t *testing.T) {
	transport := NewOpenSSLTransport(nil)
	req, err := http.NewRequest("GET", "https://google.com", nil)
	if err != nil {
		t.Fatal(err)
	}
	res, err := transport.RoundTrip(req)
	if err != nil {
		t.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	res.Body.Close()
	if len(body) == 0 {
		t.Fatal("Want not empty response body, got empty")
	}
}
