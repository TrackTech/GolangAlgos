package reader

import (
	"ReadJson/types"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func ReadFile(filename string) (bool, error) {
	fmt.Println("File read-", filename)
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error occured")
		return false, err
	}

	defer file.Close()

	data, err := io.ReadAll(file)

	if err != nil {
		fmt.Println("Error occured")
		return false, err
	}

	var orderList types.OrderList

	err = json.Unmarshal(data, &orderList)

	for _, o := range orderList.Orders {
		fmt.Println(o)
	}

	return true, nil
}
