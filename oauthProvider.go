package oauth2Provider
import (
  "errors"
  "encoding/json"
  "io/ioutil"
)

type OauthConf struct {
  Providers     []Provider `json:"providers"`
}

type Provider struct {
  Name  string  `json:"name"`
  ClientID      string  `json:"id"`
  ClientSecret  string  `json:"secret"`
  // SecurityKey is a random string for oauth2 API calls to protect against CSRF
  SecurityKey   string  `json:"securitykey"`
}

func ReadOauthConfFile(oauthKeysFilepath string) (*OauthConf, error){
  // set up gomniauth
  keyFile, err := ioutil.ReadFile(oauthKeysFilepath)
  if err != nil {
    return nil, err
  }

  var oauthconf OauthConf
  err = json.Unmarshal(keyFile, &oauthconf)
  if err != nil{
    return nil, err
  }
  _, err = json.Marshal(oauthconf)
  if err != nil {
    return nil, err
  }
  return &oauthconf, nil
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
