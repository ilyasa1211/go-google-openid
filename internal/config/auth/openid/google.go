package config

import "os"

type GoogleOpenIdConf struct {
	ClientId             string
	ClientSecret         string
	DiscoveryDocumentUrl string
	CallbackUrl          string
}

func NewGoogleOpenIdConfg() *GoogleOpenIdConf {
	return &GoogleOpenIdConf{
		os.Getenv("GOOGLE_OPEN_ID_CLIENT_ID"),
		os.Getenv("GOOGLE_OPEN_ID_CLIENT_SECRET"),
		os.Getenv("GOOGLE_DISCOVERY_DOCUMENT_URL"),
		os.Getenv("GOOGLE_OPEN_ID_CALLBACK_URL"),
	}
}
