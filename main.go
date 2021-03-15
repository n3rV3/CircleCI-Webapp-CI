package main

import (
	"fmt"
	"net/http"
    "os"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
    config_map_update := os.Getenv("CONFIG_VERSION")
    image_version := os.Getenv("IMAGE_VERSION")
    fmt.println("\nconfig version: %s\nimage version: %s", config_map_update, image_version')
    fmt.Fprintf(w, "Test CI Webserver")
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":80", nil)
}
