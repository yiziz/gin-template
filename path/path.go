package path

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// assumes filename will not err
func filename(path string) string {
	fn, err := filepath.Abs(path)
	if err != nil {
		// TODO log error
	}
	return fn
}

func configFilename(name string, params ...string) string {
	if len(params) == 0 {
		params = append(params, ProjectPath())
	}
	srcPath := strings.Join(params, "")
	fmt.Println(srcPath)
	return filename(srcPath + name)
}

// ProjectPath returns string defined by ENV["GIN_TEMPLATE_PATH"]
func ProjectPath() string {
	return os.Getenv("GIN_TEMPLATE_PATH")
}

// DBConfigFilename returns where the database.yml is found
func DBConfigFilename(params ...string) string {
	return configFilename("/config/database.yml")
}

// AppConfigFilename returns where application.yml is found
func AppConfigFilename(params ...string) string {
	return configFilename("/config/application.yml")
}
