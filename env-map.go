package envmap

import (
	"os"
	"strings"
)

// Get returns a map of the environment variables.
func Get() map[string]string {
	environ := os.Environ()

	return makeEnvMap(environ)
}

func makeEnvMap(environ []string) map[string]string {
	envMap := make(map[string]string)
	for _, v := range environ {
		kv := strings.Split(v, "=")
		if kv[0] == "" {
			continue
		}
		envMap[kv[0]] = kv[1]
	}
	return envMap
}
