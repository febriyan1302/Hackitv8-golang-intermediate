package main

import (
	"context"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"github.com/crewjam/saml/samlsp"
	"net/http"
	"net/url"
)

func newSamlMiddleware() (*samlsp.Middleware, error) {
	keyPair, err := tls.LoadX509KeyPair(samlCertificatePath, samlPrivateKeyPath)
	if err != nil {
		return nil, err
	}

	keyPair.Leaf, err = x509.ParseCertificate(keyPair.Certificate[0])
	if err != nil {
		return nil, err
	}

	idpMetadataURL, err := url.Parse(samlIDPMetadata)
	if err != nil {
		return nil, err
	}

	idpMetadata, err := samlsp.FetchMetadata(context.Background(), http.DefaultClient, *idpMetadataURL)
	if err != nil {
		return nil, err
	}

	rootURL, err := url.Parse(webserveRootURL)
	if err != nil {
		return nil, err
	}

	sp, err := samlsp.New(samlsp.Options{
		URL:         *rootURL,
		Key:         keyPair.PrivateKey.(*rsa.PrivateKey),
		Certificate: keyPair.Leaf,
		IDPMetadata: idpMetadata,
	})

	if err != nil {
		return nil, err
	}

	return sp, nil
}
