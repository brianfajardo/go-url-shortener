package urlshort

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

type pathUrl struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func MapHandler(pathsToUrl map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		if dest, ok := pathsToUrl[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}

		fallback.ServeHTTP(w, r)
	}
}

func YamlHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// parse yaml
	pathUrls, err := parseYaml(yamlBytes)
	if err != nil {
		return nil, err
	}

	// convert yaml array into map
	pathsToUrls := buildMap(pathUrls)

	// return a map handler using the map
	return MapHandler(pathsToUrls, fallback), nil
}

func parseYaml(data []byte) ([]pathUrl, error) {
	var pathUrls []pathUrl

	err := yaml.Unmarshal(data, &pathUrls)
	if err != nil {
		return nil, err
	}

	return pathUrls, nil
}

func buildMap(pathUrls []pathUrl) map[string]string {
	pathsToUrls := make(map[string]string)
	for _, pu := range pathUrls {
		pathsToUrls[pu.Path] = pu.URL
	}

	return pathsToUrls
}
