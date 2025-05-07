package dadata

import (
	"github.com/ekomobile/dadata/v2"
	"github.com/ekomobile/dadata/v2/api/suggest"
	"github.com/ekomobile/dadata/v2/client"
)

type Dadater interface {
	data() error
}

type Dadata struct {
	Api *suggest.Api
}

func NewDadata() Dadata {
	creds := client.Credentials{
		ApiKeyValue:    "443300fcd05523b23fd8dc8bfb1122c1b620fdeb",
		SecretKeyValue: "6fb0c08d292994901c752f91117fe3ad9b6fc2eb",
	}

	api := dadata.NewSuggestApi(client.WithCredentialProvider(&creds))

	return Dadata{
		Api: api,
	}
}
