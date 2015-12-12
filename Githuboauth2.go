package oauth2Provider
import (
    "fmt"
    "net/http"
    "log"
    "github.com/google/go-github/github"
    "golang.org/x/oauth2"
    githuboauth "golang.org/x/oauth2/github"
    "net/url"
)

const(
  defaultBaseURL = "https://api.github.com/"
  mediaTypeV3      = "application/vnd.github.v3+json"
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
func (c *OauthConf) HandleGitHubLogin(w http.ResponseWriter, r *http.Request) {
    provider, err := c.GetProvider("github")
    if err != nil {
      log.Fatal(err)
    }
    oauthConf := GithubOauth2Conf(provider)
    url := oauthConf.AuthCodeURL(provider.SecurityKey, oauth2.AccessTypeOnline)
    http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// /auth/github/callback Called by github after authorization is granted
func (c *OauthConf) HandleGitHubCallback(w http.ResponseWriter, r *http.Request) {
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

func (s *UsersService) GetUser() (*User, *Response, error) {
	var u string
	u = "user"

	req, err := GetUserRequest()
	if err != nil {
		return nil, nil, err
	}

	uResp := new(User)
	resp, err := DoGetUser(req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, err
}

func GetUserRequest() (*http.Request, error) {
	rel, err := url.Parse("user")
	if err != nil {
		return nil, err
	}
  baseURL, _ := url.Parse(defaultBaseURL)
	u := baseURL.ResolveReference(rel)

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", mediaTypeV3)

	return req, nil
}

func DoGetUser(req *http.Request, v interface{}) (*Response, error) {
  // HTTP client used to communicate with the API.
	client := http.DefaultClient
  resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	response := newResponse(resp)

	err = CheckResponse(resp)
	if err != nil {
		// even though there was an error, we still return the response
		// in case the caller wants to inspect it further
		return response, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
		}
	}
	return response, err
}

// CheckResponse checks the API response for errors, and returns them if
// present.  A response is considered an error if it has a status code outside
// the 200 range.  API error responses are expected to have either no response
// body, or a JSON response body that maps to ErrorResponse.  Any other
// response body will be silently ignored.
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}
	return errorResponse
}
