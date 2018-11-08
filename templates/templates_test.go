package templates_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/AirHelp/filler/templates"
)

const (
	dirToScanTest = "../test/templates"
	fileExtTest   = "tpl"
)

func TestSearchAndFill(t *testing.T) {

	os.Setenv("TEST1", "test_1")
	os.Setenv("TEST2", "test_2")

	if err := templates.SearchAndFill(dirToScanTest, fileExtTest); err != nil {
		t.Error("Could not search and fill templates. Error: ", err.Error())
	}

	list, err := filepath.Glob("../test/templates/*/*.conf")

	if err != nil {
		t.Error("Could not find filled files. Error: ", err.Error())
	}

	if len(list) != 4 {
		t.Error("Expected 4 files and i find: ", len(list))
	}
}
