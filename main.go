package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/VinceHPicton/go-struct-generator/buildstructs"
	"github.com/VinceHPicton/go-struct-generator/file_parser"
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

	structFields, err := file_parser.ParseFileToStructFields(file)
	if err != nil {
		log.Fatalf(err.Error())
	}

	structStrings := buildstructs.CreateStructStringsForFieldSlice(structFields)

	printStrings(structStrings)

	err = writeStringsToOutFile(structStrings)
	if err != nil {
		log.Fatalf(err.Error())
	}

}

func printHelp() {
	fmt.Println("Go-struct-generator is a tool which generates go structs from json files.")
	fmt.Println()
	fmt.Println("Usage: go-struct-generator [options]")
	fmt.Println("Options:")
	flag.PrintDefaults()
	os.Exit(0)
}

func writeStringsToOutFile(structDefinitions []string) error {
	file, err := os.Create("out.txt")
	if err != nil {
		log.Fatalf("Failed to create out file, %v", err.Error())
	}
	defer file.Close()

	for _, s := range structDefinitions {

		_, err := file.WriteString(s)
		if err != nil {
			return err
		}
	}
	return nil
}

func printStrings(fieldSlice []string) {

	for _, s := range fieldSlice {
		fmt.Print(s)
	}
}
