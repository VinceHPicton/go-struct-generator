# go-struct-generator
Reads a JSON file and generates go struct definitions for you recursively - structs can be nested inside structs as deeply as needed

### Installing
```bash
go install github.com/VinceHPicton/go-struct-generator
```

### How to use
After installing, create a json file for the program to use and simply run:

```bash
go-struct-generator -file relative/path/to/your/json/file.json
```

This will generate a file named out.txt which will contain the generated Go structs for you to paste into your app.

### How to structure your json file
- See the example_structs.json for these points in action
- The file is an array of objects, each of these objects represent a struct you desire to create. They should each have "name" "type" and "fields" attributes
- The "fields" consist of an array of objects representing your desired fields for that struct, for a field which is not a "struct"/"[]struct" you only need "name" and "type"
- Except where "type" is "struct" or "[]struct", "name" "tag" and "type" are blindly pasted as given into the line of that field within the generated struct definition
- In the "struct" case, the go struct field's type will be pasted in from the "name" field in the json object, if "[]struct" it will be "[]json_name" where json_name is the "name" field in the json object
- If you desire a struct tag (eg `json:"name"`) use the "tag" attribute aswell (including the backticks), "tag" is not a required field.


### If a struct field is itself a struct or an array of structs:
- You must set the "type" as "struct" or "[]struct" for this to be recognised by the program
- Use another "fields" array attribute within that field's json object, and populate it just the same way as the "fields" described above
- If this substruct itself has more structs as fields, simply repeat these steps for those, the program will define structs as many levels deep as required.