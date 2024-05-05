package files

import (
	"encoding/json"
	"fmt"
	"github/mtag-io/gohlp/term"
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"path"
)

const JSON = true
const YAML = false

func splitPath(pth string, fileName string) (string, string) {
	if fileName == "" {
		fileName = path.Base(pth)
		pth = path.Dir(pth)
	}
	return pth, fileName
}

func openFile(pth string, fileName string) []byte {
	fn := path.Join(pth, fileName)
	f, err := os.Open(fn)
	if err != nil {
		msg := fmt.Sprintf("No %s file found in %s", fn, pth)
		term.Abort(msg)
	}
	defer CloseF(f)

	raw, err := io.ReadAll(f)
	if err != nil {
		msg := fmt.Sprintf("Unable to read %s file", fn)
		term.Abort(msg)
	}
	return raw
}

func writeFile(pth string, content []byte) error {
	f, err := os.Create(pth)
	if err != nil {
		return err
	}

	_, err = f.Write(content)
	if err != nil {
		return err
	}

	CloseF(f)
	return nil
}

func readDataFile(kind bool, pth string, fileName string) map[string]interface{} {
	pth, fileName = splitPath(pth, fileName)
	raw := openFile(pth, fileName)
	var data map[string]interface{}
	if kind == JSON {
		err := json.Unmarshal(raw, &data)
		if err != nil {
			msg := fmt.Sprintf("Unable to parse JSON file %s", fileName)
			term.Abort(msg)
		}
	} else {
		err := yaml.Unmarshal(raw, &data)
		if err != nil {
			msg := fmt.Sprintf("Unable to parse YAML file %s", fileName)
			term.Abort(msg)
		}
	}
	return data
}

func ReadJson(pth string, fileName string) map[string]interface{} {
	return readDataFile(JSON, pth, fileName)
}

func ReadYaml(pth string, fileName string) map[string]interface{} {
	return readDataFile(YAML, pth, fileName)
}

func ReadToStruct[T interface{}](kind bool, pth string, fileName string, mem T) T {
	raw := openFile(pth, fileName)
	if kind == JSON {
		err := json.Unmarshal(raw, &mem)
		if err != nil {
			msg := fmt.Sprintf("Unable to parse JSON file %s", fileName)
			term.Abort(msg)
		}
	} else {
		err := yaml.Unmarshal(raw, &mem)
		if err != nil {
			msg := fmt.Sprintf("Unable to parse YAML file %s", fileName)
			term.Abort(msg)
		}
	}
	return mem
}

func ReadJSONToStruct[T interface{}](pth string, fileName string, mem T) T {
	return ReadToStruct(JSON, pth, fileName, mem)
}

func ReadYAMLToStruct[T interface{}](pth string, fileName string, mem T) T {
	return ReadToStruct(YAML, pth, fileName, mem)
}

func WriteJson(pth string, data interface{}) error {
	content, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return writeFile(pth, content)
}

func CloseF(f *os.File) {
	err := f.Close()
	if err != nil {
		term.Warn("Couldn't close file")
	}
}

func Copy(src, dest string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	srcF, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer CloseF(srcF)

	destFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if destFileStat.Mode().IsDir() {
		dest = path.Join(dest, path.Base(src))
	}

	destF, err := os.Create(dest)
	if err != nil {
		return 0, err
	}
	defer CloseF(destF)
	nBytes, err := io.Copy(destF, srcF)
	return nBytes, err
}
