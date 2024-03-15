// greeting.go

package handlers

import (
    "fmt"
    "html"
    "net/http"
    "strings"
)

var defaultGreeting = "Hello"

// GreetHandler greets the user based on the name provided in the URL path.
func GreetHandler(w http.ResponseWriter, r *http.Request) {
    name := strings.Trim(r.URL.Path, "/")
    if name == "" {
        name = "Gopher"
    }

    fmt.Fprintf(w, "<!DOCTYPE html>\n")
    fmt.Fprintf(w, "%s, %s!\n", defaultGreeting, html.EscapeString(name))
}
