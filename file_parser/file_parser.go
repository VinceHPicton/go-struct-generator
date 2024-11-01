package file_parser

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/VinceHPicton/go-struct-generator/buildstructs"
)

func ParseFileToStructFields(file string) ([]buildstructs.StructField, error) {

	jsonBytes := readFile(file)

	var fieldSlice []buildstructs.StructField

	decoder := json.NewDecoder(bytes.NewReader(jsonBytes))

	decoder.DisallowUnknownFields()

	err := decoder.Decode(&fieldSlice)
	if err != nil {
		return []buildstructs.StructField{}, err
	}

	addTypeDefinitionToFields(&fieldSlice)

	return fieldSlice, nil
}

// addTypeDefinitionToFields simply adds a datatype to each passed field, this is done because type type of the top level items in the
// input JSON array are all structs, this is a hacky workaround but it would be silly to require the user to explicitly add this data.
// We use a pointer here for performance - otherwise Go will copy the entire input data twice just for this step
func addTypeDefinitionToFields(fields *[]buildstructs.StructField) {
	for i := range *fields {
		(*fields)[i].DataType = "struct"
	}
}

func readFile(fileName string) []byte {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Failed to open JSON file: %v, err: %v", fileName, err)
	}

	jsonBytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf("Failed to read JSON file: %v, err: %v", fileName, err)
	}
	defer jsonFile.Close()

	return jsonBytes
}
