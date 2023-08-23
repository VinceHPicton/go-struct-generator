# go-struct-generator
Reads a JSON file and generates go struct definitions for you recursively - structs can be nested inside structs as deeply as needed

### How to use
Pull the repo and simply run:

```bash
go run main.go -file relative/path/to/your/file.json
```

This will generate a file named out.txt which will contain the generated Go structs for you to paste into your app.

### How to structure your json file
- See the example_structs.json for these points in action
- The file is an array of objects, each of these objects represent a struct you desire to create. They should each have "name" and "fields" attributes
- The "fields" consist of an array of objects representing your desired fields for that struct, for a field which is not a struct/[]struct you only need "name" and "type"
- If you desire a struct tag (eg `json:"name"`) use the "tag" attribute aswell, it is not required.


### If your field is itself a struct or an array of structs:
- You must set the "type" as "struct" or "[]struct" for this to be recognised by the program
- Use another "field" array attribute on that field object, and populate it just the same way as the "fields" described above
- If this substruct itself has more structs as fields, simply repeat these steps for those, the program will define structs as many levels deep as required.