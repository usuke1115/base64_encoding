package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Base64Table struct {
	Id   string `json:"id"`
	Char string `json:"char"`
}

func strToBin(str string) string {
	var binaryString string
	for _, char := range str {
		binaryString += "0" + strconv.FormatInt(int64(char), 2)
	}
	return binaryString
}

func divideSixBits(str string)[]string {
	remainder := len(str) % 6

	if remainder != 0 {
		str += strings.Repeat("0", 6 - remainder)
	}

	var binaryArray []string
	for i := 0; i < len(str); i += 6 {
		binaryArray = append(binaryArray, str[i:i+6])
	}
	return binaryArray
}

func getBase64Table(path string) []Base64Table {
	jsonFile, err := ioutil.ReadFile(path)
	if err != nil {
        log.Fatal(err)
    }

	var base64Table []Base64Table
	json.Unmarshal(jsonFile, &base64Table)

	return base64Table
}

func addEqual(str string) string {
	length := len(str)
	remainder := length % 4
	if remainder != 0 {
		str += strings.Repeat("=", 4 - remainder)
	}
	return str
}

func binToStr(binaryArray []string) string {
	var str string
	base64Table := getBase64Table("./src/base64_table.json")
	for _, bin := range binaryArray {
		for _, element := range base64Table {
			if bin == element.Id {
				str += element.Char
				break
			}
		}
	}
	str = addEqual(str)
	return str
}

func Base64Encoding(str string) string {
	binaryString := strToBin(str)
	binaryArray := divideSixBits(binaryString)
	result := binToStr(binaryArray)
	return result
}
