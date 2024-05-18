package files

import (
	"fmt"
	"github.com/mtag-io/gohlp/term"
	"os"
	"path"
	"testing"
)

const tPath = "__fixtures__"

func TestReadJson(t *testing.T) {

	actual := ReadJson(tPath, "test.json")

	if actual == nil {
		t.Fail()
	}
	if actual["key3"].(map[string]interface{})["key31"].(map[string]interface{})["key311"].(string) != "val311" {
		t.Fail()
	}
}

func TestReadJSONToStruct(t *testing.T) {

	type Key1 struct {
		Key11 string `json:"key11"`
	}

	type Key31 struct {
		Key311 string
	}

	type Key3 struct {
		Key31 Key31 `json:"key31"`
	}

	type Actual struct {
		Key1 Key1
		Key2 string `json:"key2"`
		Key3 Key3
	}

	actual := Actual{}

	ReadJSONToStruct(tPath, "test.json", &actual)

	if actual.Key1.Key11 != "val11" {
		t.Fail()
	}

	if actual.Key2 != "val2" {
		t.Fail()
	}

	if actual.Key3.Key31.Key311 != "val311" {
		t.Fail()
	}
}

func TestReadYaml(t *testing.T) {

	actual := ReadYaml(tPath, "test.yml")

	if actual == nil {
		t.Fail()
	}
	if actual["key3"].(map[string]interface{})["key31"].(map[string]interface{})["key311"].(string) != "val311" {
		t.Fail()
	}
}

func TestReadYAMLToStruct(t *testing.T) {

	type Key1 struct {
		Key11 string `yaml:"key11"`
	}

	type Key31 struct {
		Key311 string
	}

	type Key3 struct {
		Key31 Key31 `yaml:"key31"`
	}

	type Actual struct {
		Key1 Key1
		Key2 string `yaml:"key2"`
		Key3 Key3
	}

	actual := Actual{}

	ReadYAMLToStruct(tPath, "test.json", &actual)

	if actual.Key1.Key11 != "val11" {
		t.Fail()
	}

	if actual.Key2 != "val2" {
		t.Fail()
	}

	if actual.Key3.Key31.Key311 != "val311" {
		t.Fail()
	}
}

func TestWriteJson(t *testing.T) {
	tJson := map[string]interface{}{
		"key1": map[string]interface{}{
			"key11": "val11",
		},
		"key2": "val2",
		"key3": map[string]interface{}{
			"key31": map[string]interface{}{
				"key311": "val311",
			},
			"key32": "val32",
		},
	}

	pth := path.Join(tPath, "write.json")
	err := WriteJson(pth, tJson)
	if err != nil {
		t.Fail()
	}

	var stat os.FileInfo
	if stat, err = os.Stat(pth); err != nil {
		t.Fail()
	}

	if stat == nil || stat.Size() != 143 {
		t.Fail()
	}

	err = os.Remove(pth)
	if err != nil {
		msg := fmt.Sprintf("Unable to cleanup. Please delete %s manually", pth)
		term.Warn(msg)
	}
}

func TestCopy(t *testing.T) {

	sPath := path.Join(tPath, "test.json")
	dPath := path.Join(tPath, "test-copy.json")

	size, err := Copy(sPath, dPath)

	if err != nil {
		t.Fail()
	}

	if size != 143 {
		t.Fail()
	}

	var stat os.FileInfo
	if stat, err = os.Stat(dPath); err != nil {
		t.Fail()
	}

	if stat == nil || stat.Size() != 143 {
		t.Fail()
	}

	err = os.Remove(dPath)
	if err != nil {
		msg := fmt.Sprintf("Unable to cleanup. Please delete %s manually", dPath)
		term.Warn(msg)
	}
}
