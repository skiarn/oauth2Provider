# oauth2Provider

## Example usage
``
package main
import(
  "net/http"
  "fmt"
  //"oauth2Provider"
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
    http.HandleFunc("/", handleMain)
    //http.HandleFunc("/auth/github", handleGitHubLogin)
    //http.HandleFunc("/auth/github/callback", handleGitHubCallback)
    fmt.Print("Started running on http://127.0.0.1:7000\n")
    fmt.Println(http.ListenAndServe(":7000", nil))
}
``
