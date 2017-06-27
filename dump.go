package dump

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
)

func Save(path string, object interface{}) error {
	file, err := os.Create(path)
	if err == nil {
		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "    ")
		encoder.Encode(object)
	}
	file.Close()
	return err
}

func Load(path string, object interface{}) error {
	file, err := os.Open(path)
	if err == nil {
		decoder := json.NewDecoder(file)
		err = decoder.Decode(object)
	}
	file.Close()
	return err
}

func Check(e error) {
	if e != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Println(line, "\t", file, "\n", e)
		os.Exit(1)
	}
}
