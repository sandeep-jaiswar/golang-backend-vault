// version.go

package handlers

import (
    "fmt"
    "html"
    "net/http"
    "runtime/debug"
)

// VersionHandler provides version information of the application.
func VersionHandler(w http.ResponseWriter, r *http.Request) {
    info, ok := debug.ReadBuildInfo()
    if !ok {
        http.Error(w, "no build information available", 500)
        return
    }

    fmt.Fprintf(w, "<!DOCTYPE html>\n<pre>\n")
    fmt.Fprintf(w, "%s\n", html.EscapeString(info.String()))
}
