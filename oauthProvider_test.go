package oauth2Provider
import(
  "testing"
)
var(
  oauthConfilePath = "./oauthKeys.json"
)
func TestGetProviderExists(t *testing.T){

  conf, err := ReadOauthConfFile(oauthConfilePath)
  if err != nil {
			t.Errorf("Reading file: %q, returned %q", oauthConfilePath, err)
	}
  param := "github"
  if result, _ := conf.GetProvider(param); result.Name != param {
			t.Errorf("GetProvider(%s) returned %t, expected: %t", param, result, result.Name)
	}
}

func TestGetProviderNotExists(t *testing.T){

  conf, err := ReadOauthConfFile(oauthConfilePath)
  if err != nil {
			t.Errorf("Reading file: %q, returned %q", oauthConfilePath, err)
	}
  param := "notExistProvider"
  expected := "Provider "+ param +" not found."
  if _, err := conf.GetProvider(param); err.Error() != expected {
			t.Errorf("GetProvider(%q) returned %q, expected: %q", param, err, expected)
	}
}

var (
  expectedProviders = []string{"github", "google", "facebook",}
  expectedSecurityKey = "RandomTextToSecureAgainstCSRF"
)
func TestReadOauthConfFile(t *testing.T){
  conf, err := ReadOauthConfFile(oauthConfilePath)
  if err != nil {
			t.Errorf("Reading file: %q, returned %q", oauthConfilePath, err)
	}
  for _, expectedProvider := range expectedProviders {
    provider, err := conf.GetProvider(expectedProvider)
    if err != nil {
      t.Errorf("GetProvider(%q): returned %q, expected provider.", expectedProvider, err)
    }
    if provider.Name != expectedProvider {
      t.Errorf("Expected provider name to be %q but found: %q", expectedProvider, provider.Name)
    }
    expectedSecret := provider.Name +"secret"
    if provider.ClientSecret != expectedSecret {
      t.Errorf("Expected ClientSecret to be:%q but found: %q", expectedSecret, provider.ClientSecret)
    }
    if provider.SecurityKey != expectedSecurityKey {
      t.Errorf("Expected SecurityKey to be:%q but found: %q", expectedSecurityKey, provider.SecurityKey)
    }
  }
}


func TestReadOauthConfFileNotExists(t *testing.T){
  filepath := "./notexist.json"
  if result, err := ReadOauthConfFile(filepath); err == nil {
			t.Errorf("Reading file: %q, returned %q, expected: file not found error", filepath, result)
	}
}
