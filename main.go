package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/VinceHPicton/go-struct-generator/buildstructs"
)

func main() {

	var file string

	flag.StringVar(&file, "file", "", "Relative path to the input json file for the program to use - eg path/to/example.json")
	showHelp := flag.Bool("help", false, "Show usage help")

	flag.Parse()

	if *showHelp || file == "" {
		printHelp()
		return
	}

	jsonBytes := readFile(file)

	var fieldSlice []buildstructs.StructField

	decoder := json.NewDecoder(bytes.NewReader(jsonBytes))

	decoder.DisallowUnknownFields()

	err := decoder.Decode(&fieldSlice)
	if err != nil {
		log.Fatalf(err.Error())
	}

	for i := range fieldSlice {
		fieldSlice[i].DataType = "struct"
	}

	fmt.Println(fieldSlice)

	// err := json.Unmarshal(jsonBytes, &fieldSlice)
	// if err != nil {
	// 	log.Fatalf("Failed to unmarshal JSON")
	// }

	structDefinitions := buildstructs.CreateStructsForFieldSlice(fieldSlice)

	printGoStructs(structDefinitions)

	writeGoStructsToOutFile(structDefinitions)

}

func printHelp() {
	fmt.Println("Go-struct-generator is a tool which generates go structs from json files.")
	fmt.Println()
	fmt.Println("Usage: go-struct-generator [options]")
	fmt.Println("Options:")
	flag.PrintDefaults()
	os.Exit(0)
}

func writeGoStructsToOutFile(structDefinitions []string) {
	file, err := os.Create("out.txt")
	if err != nil {
		log.Fatalf("Failed to create out file, %v", err.Error())
	}
	defer file.Close()

	for _, s := range structDefinitions {

		_, err := file.WriteString(s)
		if err != nil {
			log.Fatalf("Failed to write line: %v", s)
		}
	}
}

func printGoStructs(fieldSlice []string) {

	for _, s := range fieldSlice {
		fmt.Print(s)
	}
}

func readFile(fileName string) []byte {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Failed to open JSON file: %v", fileName)
	}

	jsonBytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf("Failed to read JSON file: %v", fileName)
	}
	defer jsonFile.Close()

	return jsonBytes
}
