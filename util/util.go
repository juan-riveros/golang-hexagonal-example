package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

func CheckErr(fatal bool, err error, msg ...string) {
	if fatal && err != nil {
		log.Fatalln(time.Now().UnixNano(), "[ERR]", err, fmt.Sprintf(msg[0], msg[1:]))
	} else if err != nil {
		log.Println(time.Now().UnixNano(), "[ERR]", err, fmt.Sprintf(msg[0], msg[1:]))
	}
}

func ParseJSON(filepath string, v interface{}) {
	fileContent, err := ioutil.ReadFile(filepath)
	CheckErr(true, err, "Failed to read file %s", filepath)

	err = json.Unmarshal(fileContent, v)
	CheckErr(true, err, "Failed to unmarshal file %s", filepath)
}

func SaveJSON(filepath string, v interface{}) error {
	data, err := json.Marshal(v)
	CheckErr(false, err, "Failed to marshal json %s", fmt.Sprintf("%+v", v))

	err = ioutil.WriteFile(filepath, data, 0600)
	CheckErr(false, err, "Failed to save file %s", filepath)
	return err
}
