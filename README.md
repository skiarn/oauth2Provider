# oauth2Provider

## Example usage, create file oauthKeys.json.
example see, src/github.com/skiarn/auth2Provider/oauthKeys.json
```
package main
import(
  "net/http"
  "log"
  oauth "github.com/skiarn/oauth2Provider"
  "fmt"
)

const htmlIndex = `<html><body>
Logged in with <a href="/login">GitHub</a>
</body></html>
`
func handleMain(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(htmlIndex))
}

func main() {
    var oauthConf *oauth.OauthConf
    oauthConf, err := oauth.ReadOauthConfFile("./oauthKeys.json")
    if err != nil {
      log.Fatal(err)
    }
    http.HandleFunc("/", handleMain)
    http.HandleFunc("/auth/github", oauthConf.HandleGitHubLogin)
    http.HandleFunc("/auth/github/callback", oauthConf.HandleGitHubCallback)
    fmt.Print("Started running on http://127.0.0.1:7000\n")
    fmt.Println(http.ListenAndServe(":7000", nil))
}

```
