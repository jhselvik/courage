/*
Package main implements the GoChatService or as Apple calls it the Provider to the APNS.

This version will send a basic APN to the APNS to be received on a development iPod.

This will eventually evolve into a New Tricks Project named Courage.
*/
package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	Config *Configuration
)

// main configures and starts the Provider.
func main() {

	// Get the Cert, private key, and Entrust CA root files from here
	Config = GetConfiguration()

	// Pair the Certificate and private key
	certPair, err := tls.LoadX509KeyPair(Config.TLSCertFile, Config.TLSKeyFile)
	if err != nil {
		panic(err)
	}

	// Read the RootCA file into a []byte
	rootCA, err := ioutil.ReadFile(Config.EntrustCAFile)
	if err != nil {
		panic(err)
	}

	// Pull out the RootCAs
	pool := x509.NewCertPool()
	if !pool.AppendCertsFromPEM(rootCA) {
		panic(err)
	}

	tr := &http.Transport{
		// Add the cert, private key, and Entrust CA to the connection request
		TLSClientConfig: &tls.Config{Certificates: []tls.Certificate{certPair}, RootCAs: pool},
	}

	client := &http.Client{Transport: tr}

	resp, err := client.Get("https:gateway.sandbox.push.apple.com:2195")
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Status)
}
