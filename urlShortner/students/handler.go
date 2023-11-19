package students

import (
	yamlV2 "gopkg.in/yaml.v2"
	"net/http"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path, ok := pathsToUrls[r.URL.Path]
		if ok {
			http.Redirect(w, r, path, http.StatusFound)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

func YAMLHandler(yaml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYAML, err := parseYAML(yaml)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(parsedYAML)
	return MapHandler(pathMap, fallback), nil
}

func parseYAML(yaml []byte) (dst []map[string]string, err error) {
	err = yamlV2.Unmarshal(yaml, &dst)
	return dst, err
}

func buildMap(parsedYaml []map[string]string) map[string]string {
	mergedMap := make(map[string]string)
	for _, entry := range parsedYaml {
		key := entry["path"]
		mergedMap[key] = entry["url"]
	}
	return mergedMap
}
