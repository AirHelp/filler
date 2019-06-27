package templates_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/AirHelp/filler/templates"
)

const (
	dirToScanTest = "../test/templates"
	fileExtTest   = "tpl"
)

func TestSearchAndFill(t *testing.T) {

	files := []struct {
		fileName    string
		fileContent string
	}{
		{
			fileName:    "../test/templates/confB/a.conf",
			fileContent: "test_1\ntest_2",
		}, {
			fileName:    "../test/templates/confB/b.conf",
			fileContent: "test_1\ntest_2",
		}, {
			fileName:    "../test/templates/confA/a.conf",
			fileContent: "test_1\ntest_2",
		}, {
			fileName:    "../test/templates/confA/b.conf",
			fileContent: "test_1\ntest_2",
		}, {
			fileName:    "../test/templates/confB/c.conf",
			fileContent: "a1\na2\na3\n",
		}, {
			fileName:    "../test/templates/confC/a.conf",
			fileContent: "{ \n\t{\n\t\t\"key\": \"key1\",\n\t\t\"value\": \"key1-value\",\n\t},\n\t{\n\t\t\"key\": \"key2\",\n\t\t\"value\": \"key2-value\",\n\t},\n}\n",
		},
	}

	os.Setenv("TEST1", "test_1")
	os.Setenv("TEST2", "test_2")
	os.Setenv("ARRAY", "a1,a2,a3")
	os.Setenv("KEYS", "key1,key2")
	os.Setenv("MAP", "{\"key1\":\"key1-value\",\"key2\":\"key2-value\"}")

	if err := templates.SearchAndFill(dirToScanTest, fileExtTest, false); err != nil {
		t.Error("Could not search and fill templates. Error: ", err.Error())
	}

	for _, file := range files {
		data, err := ioutil.ReadFile(file.fileName)
		if err != nil {
			t.Error(err)
		}
		if string(data) != file.fileContent {
			t.Error("Wrong file content!")
		}
	}

}
