package csv2json

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"
)

var dat map[string]interface{}

func ParseCsv(filename string) (error, string) {
	// dat := make(map[string]interface{})
	// var buffer bytes.Buffer
	// get current directory
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		return err, ""
	}

	filePath := pwd + "/" + filename + ".csv"
	// open file from current directory
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return err, ""
	}
	// automatically call Close() at the end of current method
	defer file.Close()

	//
	reader := csv.NewReader(file)
	// options are available at:
	// http://golang.org/src/pkg/encoding/csv/reader.go?s=3213:3671#L94
	reader.Comma = ';'
	// jsonString = `{}`
	var headers []string
	var objs []map[string]interface{}
	lineCount := 0
	// buffer.WriteString("`{")
	for {
		// read just one record, but we could ReadAll() as well
		record, err := reader.Read()
		// end-of-file is fitted into err
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
			return err, ""
		}
		// record is an array of string so is directly printable
		// fmt.Println("Record", lineCount, "is", record, "and has", len(record), "fields")

		var obj = make(map[string]interface{})
		if lineCount == 0 {
			for _, fieldName := range record {
				headers = append(headers, strings.ToLower(fieldName))
			}
		} else {
			for index, value := range record {
				obj[string(headers[index])] = value
			}
			objs = append(objs, obj)
		}
		lineCount += 1
	}
	jsonStringResposne, _ := json.Marshal(objs)
	return nil, string(jsonStringResposne)
}
