# go-struct-generator
Reads a JSON file and generates go struct definitions for you recursively - structs can be nested inside structs as deeply as needed

### How to use
Pull the repo and simply run:

```bash
go run main.go -file your_file.json
```

### How to structure your json file
- See the example_structs.json for these points in action
- The file is an array of objects, these objects the structs you desire to create. They should each have "name" and "fields" attributes
- The "fields" consist of an array of your desired fields for that struct, for a standard field you only need "name" and "type"
- If you desire a struct tag use the "tag" attribute aswell, it is not required.


# If your field is itself a struct or an array of structs:
- You must set the "type" as "struct" or "[]struct" for this to be recognised by the program
- Use another "field" array attribute, and populate it just the same way as the "fields" described above
- If this substruct itself has more structs as fields, simply repeat these steps for those, the program will define structs as many levels deep as required.