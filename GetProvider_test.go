package oauth2Provider
import(
  "testing"
)
var(
  randomProvider = Provider{Name: "randomProvider",
                              Conf:  &ProviderConf{ClientID: "clientid", ClientSecret: "ClientSecret", SecurityKey: "SecurityKey"}}

  githubProvider = Provider{Name: "github",
                              Conf:  &ProviderConf{ClientID: "clientid", ClientSecret: "ClientSecret", SecurityKey: "SecurityKey"}}
  googleProvider = Provider{Name: "google",
                              Conf:  &ProviderConf{ClientID: "clientid", ClientSecret: "ClientSecret", SecurityKey: "SecurityKey"}}
)
func TestGetProviderExists(t *testing.T){

  conf := OauthConf{
    Providers: []Provider{randomProvider, githubProvider, googleProvider,},
  }
  param := "github"
  if result, _ := conf.GetProvider(param); result.Name != param {
			t.Errorf("GetProvider(%s) returned %t, expected: %t", param, result, result.Name)
	}
}

func TestGetProviderNotExists(t *testing.T){

  conf := OauthConf{
    Providers: []Provider{randomProvider, githubProvider, googleProvider,},
  }
  param := "notExistProvider"
  expected := "Provider "+ param +" not found."
  if _, err := conf.GetProvider(param); err.Error() != expected {
			t.Errorf("GetProvider(%q) returned %q, expected: %q", param, err, expected)
	}
}
