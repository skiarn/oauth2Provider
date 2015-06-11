package oauth2Provider
import (
  "errors"
)

type OauthConf struct {
  Providers     []Provider
}

type Provider struct {
  Name  string
  Conf  *ProviderConf
}

type ProviderConf struct {
    ClientID      string
    ClientSecret  string
    // SecurityKey is a random string for oauth2 API calls to protect against CSRF
    SecurityKey   string
}

func (conf *OauthConf) GetProvider(providerName string) (*Provider, error){
  for i := range conf.Providers {
    prov := conf.Providers[i]
    if prov.Name == providerName {
      return &prov, nil
    }
  }
  return nil, errors.New("Provider "+ providerName +" not found.")
}
