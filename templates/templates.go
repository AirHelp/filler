package templates

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"

	"strings"
	"text/template"

	"github.com/AirHelp/filler/consts"
	"go.uber.org/zap"
)

const (
	DefaultDestinationFilePerms = 0644
)

var tpl *template.Template

func init() {
	// Create a FuncMap with which to register the function.
	funcMap := template.FuncMap{
		"getEnv": func(key string) (string, error) {
			return getEnv(key)
		},
		"getEnvArray": func(key string) ([]string, error) {
			return getEnvArray(key)
		},
		"required": func(env string) (string, error) {
			return env, nil
		},
	}

	tpl = template.New("templateCli").Funcs(funcMap)
}

func SetFailIfMissing() {
	zap.S().Debugf("Setting %v", consts.FailIfMissing)
	tpl = tpl.Option("missingkey=error")
}

func SetDelimiters(left string, right string) {
	tpl = tpl.Delims(left, right)
}

func getAllEnv() map[string]string {
	envs := make(map[string]string)
	for _, env := range os.Environ() {
		ret := strings.Split(env, "=")
		envs[ret[0]] = ret[1]
	}
	return envs
}

func globExt(dir string, ext string) ([]string, error) {

	files := []string{}
	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if filepath.Ext(path) == "."+ext {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}

func glob(dir string) ([]string, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		st, err := os.Stat(path)

		if err != nil {
			return err
		}

		if !st.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}

func readTemplate(filePath string) (string, error) {
	templateBuffer, err := os.ReadFile(filePath)
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
	envs := getAllEnv()

	// Create a template, add the function map, and parse the text.
	tmpl, err := tpl.Parse(templateText)
	if err != nil {
		return
	}

	err = tmpl.Execute(&templateResultBuffer, envs)
	return
}

func writeTemplateResults(templateFile string, templateResultBuffer bytes.Buffer, deleteFile bool, inPlace bool) error {

	var perms os.FileMode
	currentTemplateFileInfo, err := os.Stat(templateFile)
	if err != nil {
		perms = DefaultDestinationFilePerms
	} else {
		perms = currentTemplateFileInfo.Mode()
	}

	destinationFile := templateFile
	if !inPlace {
		destinationFile = strings.TrimSuffix(templateFile, filepath.Ext(templateFile))
	}

	err = os.WriteFile(destinationFile, templateResultBuffer.Bytes(), perms)
	if err != nil {
		return err
	}

	if deleteFile {
		return os.Remove(templateFile)
	} else {
		return nil
	}

}

func SearchAndFill(toScan string, fileExt string, deleteFile bool, inPlace bool) error {

	st, err := os.Stat(toScan)

	if err != nil {
		return err
	}

	files := []string{}

	if st.IsDir() {
		var err error

		if inPlace {
			files, err = glob(toScan)
		} else {
			files, err = globExt(toScan, fileExt)
		}

		if err != nil {
			return err
		}
	} else {
		files = append(files, toScan)
	}

	for _, file := range files {
		zap.S().Debugf("Templating file: %v", file)
		templateText, err := readTemplate(file)
		if err != nil {
			return err
		}
		templateResultBuffer, err := renderTemplate(templateText)
		if err != nil {
			return err
		}
		if err := writeTemplateResults(file, templateResultBuffer, deleteFile, inPlace); err != nil {
			return err
		}
	}

	return nil
}
