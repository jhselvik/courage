package main

import (
	"encoding/json"
	"flag"
	"os"
)

type Configuration struct {
	TLSCertFile   string `json:"ssl_cert_file"`
	TLSKeyFile    string `json:"ssl_key_file"`
	EntrustCAFile string `json:"entrust_ca_root_cert_file"`
	ConfigFile    string
}

// GetConfiguration parses command line arguments for configuration flags, then
// processes any provided config file for additional flags. A given config file
// will override any flags.
func GetConfiguration() *Configuration {
	c := new(Configuration)

	flag.StringVar(&c.TLSCertFile, "tls-cert-file", "GoChatServiceCert.pem.", "See APNS tutorial part 1 for how to generate this file")
	flag.StringVar(&c.TLSKeyFile, "tls-key-file", "GoChatServiceKey.pem", "Matching private key for the provided tls-cert-file")
	flag.StringVar(&c.EntrustCAFile, "entrust-ca-file", "entrust_2048_ca.cer", "Download from: https://www.entrust.net/downloads/binary/entrust_2048_ca.cer")
	flag.StringVar(&c.ConfigFile, "config", "", "a config file to use that can be used in place of flags; flags take precedence")
	flag.Parse()

	// Since the config file must be parsed after the flags are processed, a config file will override any flags
	if c.ConfigFile != "" {
		configFile, err := os.Open(c.ConfigFile)
		if err != nil {
			panic("Couldn't open your provided config file")
		}

		d := json.NewDecoder(configFile)
		err = d.Decode(&c)
		if err != nil {
			panic("Couldn't parse your provided config file")
		}
	}

	return c
}
