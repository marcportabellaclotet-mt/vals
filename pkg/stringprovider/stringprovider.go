package stringprovider

import (
	"fmt"

	"github.com/marcportabellaclotet-mt/valstest/pkg/api"
	"github.com/marcportabellaclotet-mt/valstest/pkg/providers/awssec"
	"github.com/marcportabellaclotet-mt/valstest/pkg/providers/gcpsecrets"
	"github.com/marcportabellaclotet-mt/valstest/pkg/providers/sops"
	"github.com/marcportabellaclotet-mt/valstest/pkg/providers/ssm"
	"github.com/marcportabellaclotet-mt/valstest/pkg/providers/tfstate"
	"github.com/marcportabellaclotet-mt/valstest/pkg/providers/vault"
)

func New(provider api.StaticConfig) (api.LazyLoadedStringProvider, error) {
	tpe := provider.String("name")

	switch tpe {
	case "ssm":
		return ssm.New(provider), nil
	case "vault":
		return vault.New(provider), nil
	case "awssecrets":
		return awssec.New(provider), nil
	case "sops":
		return sops.New(provider), nil
	case "gcpsecrets":
		return gcpsecrets.New(provider), nil
	case "tfstate":
		return tfstate.New(provider), nil
	}

	return nil, fmt.Errorf("failed initializing string provider from config: %v", provider)
}
