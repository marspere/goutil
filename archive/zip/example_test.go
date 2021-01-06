package zip

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

func ExampleNoBufferFile_AddFile() {
	// 压缩
	zipFile := NewNoBufferFile("test.zip")
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
		if err = zipFile.AddFile(strconv.Itoa(i)+".json", data); err != nil {
			log.Fatal(err)
		}
	}
	err := zipFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	// 解压缩
	contents, err := Unzip("test.zip")
	if err != nil {
		log.Fatal(err)
	}
	for _, content := range contents {
		fmt.Println(content.Filename)
		fmt.Println(string(content.Content))
	}
}
