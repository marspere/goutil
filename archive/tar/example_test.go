package tar

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

func ExampleNoBufferFile_AddFile() {
	// 压缩
	tarFile := NewNoBufferFile("test.tar.gz")
	var testData = []struct {
		Name string
		Age  int
	}{
		{"hello", 1},
		{"world", 2},
	}
	for i, td := range testData {
		data, err := json.Marshal(td)
		if err != nil {
			log.Fatal(err)
		}
		if err = tarFile.AddFile(strconv.Itoa(i)+".json", data); err != nil {
			log.Fatal(err)
		}
	}
	err := tarFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	// 解压
	result, err := Unzip("test.tar.gz")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range result {
		fmt.Println(file.Filename)
		fmt.Println(string(file.Content))
	}
}
