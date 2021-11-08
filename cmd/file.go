package cmd

import (
	"io/ioutil"
)

func ReadFile(file string) string {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	return string(content)
}

func WriteFile(filename string, content string) {
	err := ioutil.WriteFile(filename, []byte(content), 0777)
	if err != nil {
		panic(err)
	}
	return
}
