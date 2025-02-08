package utility

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/magiconair/properties"
)

var appProp = filepath.Join(ResourceManager{}.GetProjectLocation(), "resources", "application.properties")
var msgProp = filepath.Join(ResourceManager{}.GetProjectLocation(), "resources", "messages.properties")
var propertyFiles = []string{appProp, msgProp}
var Properties, _ = properties.LoadFiles(propertyFiles, properties.UTF8, true)

type ResourceManager struct {
}

func (resourceManager ResourceManager) GetProperty(propertyKey string) (string, error) {
	propertyValue, propertyFound := Properties.Get(propertyKey)
	if !propertyFound {
		return "", errors.New(Properties.MustGet("no.such.property") + propertyKey)
	}
	return propertyValue, nil
}

func (resourceManager ResourceManager) GetProjectLocation() string {
	mydir, _ := os.Getwd()
	return mydir
}
