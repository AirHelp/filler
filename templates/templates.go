package templates

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const (
	DefaultDestinationFilePerms = 0644
)

func glob(dir string, ext string) ([]string, error) {

	files := []string{}
	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if filepath.Ext(path) == "."+ext {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}

func readTemplate(filePath string) (string, error) {
	templateBuffer, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(templateBuffer), nil
}

func getEnv(envName string) (string, error) {
	envName = strings.ToUpper(envName)
	env := os.Getenv(envName)
	if env == "" {
		return "", errors.New("ENV variable is missing")
	} else {
		return env, nil
	}
}

func getEnvArray(envName string) ([]string, error) {
	env, err := getEnv(envName)

	if err != nil {
		return nil, err
	}

	return strings.Split(env, ","), nil
}

func renderTemplate(templateText string) (templateResultBuffer bytes.Buffer, err error) {

	// Create a FuncMap with which to register the function.
	funcMap := template.FuncMap{
		"getEnv": func(key string) (string, error) {
			return getEnv(key)
		},
		"getEnvArray": func(key string) ([]string, error) {
			return getEnvArray(key)
		},
	}

	// Create a template, add the function map, and parse the text.
	tmpl, err := template.New("templateCli").Funcs(funcMap).Parse(templateText)
	if err != nil {
		return
	}
	// Run the template.
	err = tmpl.Execute(&templateResultBuffer, nil)
	return
}

func writeTemplateResults(templateFile string, templateResultBuffer bytes.Buffer, deleteFile bool) error {

	var perms os.FileMode
	currentTemplateFileInfo, err := os.Stat(templateFile)
	if err != nil {
		perms = DefaultDestinationFilePerms
	} else {
		perms = currentTemplateFileInfo.Mode()
	}

	destinationFile := strings.TrimSuffix(templateFile, filepath.Ext(templateFile))

	err = ioutil.WriteFile(destinationFile, templateResultBuffer.Bytes(), perms)
	if err != nil {
		return err
	}

	if deleteFile {
		return os.Remove(templateFile)
	} else {
		return nil
	}

}

func SearchAndFill(toScan string, fileExt string, deleteFile bool) error {

	st, err := os.Stat(toScan)

	if err != nil {
		return err
	}

	files := []string{}

	if st.IsDir() {
		var err error

		files, err = glob(toScan, fileExt)

		if err != nil {
			return err
		}
	} else {
		files = append(files, toScan)
	}

	for _, file := range files {
		templateText, err := readTemplate(file)
		if err != nil {
			return err
		}
		templateResultBuffer, err := renderTemplate(templateText)
		if err != nil {
			return err
		}
		if err := writeTemplateResults(file, templateResultBuffer, deleteFile); err != nil {
			return err
		}
	}

	return nil
}
