package oauth2Provider
import (
    "fmt"
    "net/http"
    "log"
    "github.com/google/go-github/github"
    "golang.org/x/oauth2"
    githuboauth "golang.org/x/oauth2/github"
)

//auth/github
//auth/github/callback
// You must register the app at https://github.com/settings/applications
func GithubOauth2Conf(provider *Provider) (*oauth2.Config){
  return &oauth2.Config{
                        ClientID:     provider.ClientID,
                        ClientSecret: provider.ClientSecret,
                        // select level of access you want https://developer.github.com/v3/oauth/#scopes
                        Scopes:       []string{"user:email", "repo"},
                        Endpoint:     githuboauth.Endpoint,
                    }
}

// /login
func (c *OauthConf) handleGitHubLogin(w http.ResponseWriter, r *http.Request) {
    provider, err := c.GetProvider("github")
    if err != nil {
      log.Fatal(err)
    }
    oauthConf := GithubOauth2Conf(provider)
    url := oauthConf.AuthCodeURL(provider.SecurityKey, oauth2.AccessTypeOnline)
    http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// /auth/github/callback Called by github after authorization is granted
func (c *OauthConf) handleGitHubCallback(w http.ResponseWriter, r *http.Request) {
    provider, err := c.GetProvider("github")
    if err != nil {
      log.Fatal(err)
    }
    oauthConf := GithubOauth2Conf(provider)
    state := r.FormValue("state")
    if state != provider.SecurityKey {
        fmt.Printf("Invalid oauth state, expected '%s', got '%s'\n", provider.SecurityKey, state)
        http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
        return
    }

    code := r.FormValue("code")
    token, err := oauthConf.Exchange(oauth2.NoContext, code)
    if err != nil {
        fmt.Printf("oauthConf.Exchange() failed with '%s'\n", err)
        http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
        return
    }

    oauthClient := oauthConf.Client(oauth2.NoContext, token)
    client := github.NewClient(oauthClient)
    user, _, err := client.Users.Get("")
    if err != nil {
        fmt.Printf("client.Users.Get() failed with '%s'\n", err)
        http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
        return
    }
    fmt.Printf("Logged in as GitHub user: %s\n", *user.Login)
    http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
