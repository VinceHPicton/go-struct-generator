# go-struct-generator
Reads a JSON file and generates go struct definitions for you recursively - structs can be nested inside structs as deeply as needed

## Installing as a Go module
```bash
go install github.com/VinceHPicton/go-struct-generator
```

### How to use as a Go module
After installing, create a json file for the program to use and simply run:

```bash
go-struct-generator -file relative/path/to/your/json/file.json
```

This will generate a file named out.txt which will contain the generated Go structs for you to paste into your app.

## Or pull the docker image
```bash
docker pull theundula/go-struct-generator
```

### Run the container
```bash
docker run -v "$(pwd):/appspace/userpwd" theundula/go-struct-generator path/to/your/structs.json
```

### How to structure your json file
- See the example_structs.json for an example input file, you can run the program with this example to see what it outputs.
- The file is an array of objects, each of these objects represent a struct you desire to create. 
- Each object must have "name" and "type" attributes, and can optionally have "tag" if you want a struct tag for that go struct field, and/or "fields", needed where the type of the go struct field is either a struct itself, or an array of structs.
- "name" is the name of that field in the go struct
- "type" is the type of that field
- "tag" (optional), is the struct tag eg `json:"fieldName"`
- Except where "type" is "struct" or "[]struct", the "name" "tag" and "type" attributes are blindly pasted as given into the line of the generated struct definition - therefore you must include the backticks in the "tag" attribute value.
- Where "type" is "struct", the go struct field's type will be pasted in from the "name" field in the json object.
- Where "type" is "[]struct", the go struct field's type will be "[]json_name" where json_name is the value of the "name" field in the json object
- "fields" (optional) is another array of objects, in identical format as the top level array, used where a struct field is itself a new struct (or []struct), see details below.

### If a struct field is itself a struct or an array of structs:
- You must set the "type" as "struct" or "[]struct" for this to be recognised by the program
- Use another "fields" array attribute within that field's json object, and populate it just the same way as the "fields" described above
- If this substruct itself has more structs as fields, simply repeat these steps for those, the program will act recursively to define new structs as many levels deep as required.