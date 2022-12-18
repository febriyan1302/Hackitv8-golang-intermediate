package main

import "fmt"

var (
	samlCertificatePath = "./myservice.cert"
	samlPrivateKeyPath  = "./myservice.key"
	samlIDPMetadata     = "https://samltest.id/saml/idp"

	webserverPort   = "9000"
	webserveRootURL = fmt.Sprintf("http://localhost:%s", webserverPort)
)
