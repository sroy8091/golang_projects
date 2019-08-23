package urlshort

import (
	"fmt"
	"net/http"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	fmt.Println("here we go again", pathsToUrls, fallback)
	return nil
}

func YAMLHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	fmt.Println("ahh shit!", yamlBytes, fallback)
	return nil, nil
}
