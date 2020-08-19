package zconstants

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func IsEnvVarTrue(envVarName string) bool {
	return strings.ToLower(os.Getenv(envVarName)) == "true"
}

func GetActiveNamespace() (string, error) {
	activeNamespaceBytes, _ := ioutil.ReadFile(serviceAccountNamespacePath)
	activeNamespace := string(activeNamespaceBytes)

	if activeNamespace != "" {
		return activeNamespace, nil
	}

	return "", errors.New("could not detect active namespace")
}

func GetTektonResultsDir() string {
	if IsEnvVarTrue(OutOfClusterENV) {
		ex, err := os.Executable()
		if err != nil {
			panic(err)
		}
		return filepath.Dir(ex)
	}
	return TektonResultsDirPath
}