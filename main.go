package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"log"
	"net/http"
)

var ca *x509.Certificate
var transport *http.Transport
var tlsServerConfig *tls.Config
var caString string
var publicIP *string
var dnsPort *int

var proxyHosts = []string{
	//our own fake domain
	"free.news",
	//some big news pages that use paywalls
	"ft.com",
	"theguardian.com",
	"faz.net",
	"wsj.com",
	"nytimes.com",
	"telegraph.co.uk",
}

func main() {
	publicIP = flag.String("publicIP", getPublicIP(), "public interface ip address") 
	dnsPort = flag.Int("dnsPort", 53, "port")
	flag.Parse()

	log.Printf("[*] Welcome. Public DNS Sever IP: %s", *publicIP)
	setupCerts()
	log.Printf("[*] CA Signature: %x...", ca.Signature[:16]) 
	go serveHTTP()
	serveDNS()
}
