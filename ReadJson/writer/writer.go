package writer

import (
	"ReadJson/types"
	"encoding/json"
	"os"
)

//Public function need to start with Upper case letter

func WriteJson(fileName string, people []types.Person) error {

	jsonData, err := json.Marshal(people)

	if err != nil {
		return err
	}

	os.WriteFile(fileName, jsonData, os.ModeAppend)
	return err
}
