# Supporter GO (Helper functions)

## Install
```bash
go get -u github.com/zipzoft/supporter-go
```

## Usage

#### Array

```go
import "github.com/zipzoft/supporter-go"

arr := []string{"a", "b", "c"}

// Get the first element of the array.
// Ouput: a
fmt.Println(supporter.First(arr))   


// Check array value is exist.
// Output: true
fmt.Println(supporter.InArray("b", arr))

// Interface to Slice
// Output: []string{"a", "b", "c"}
var unknowType interface{} = []string{"a", "b", "c"}
fmt.Println(supporter.ToSlice(unknowType))

// Check array is empty.
// Output: false
fmt.Println(supporter.IsEmpty(arr))
```

--- 

#### String
```go
import "github.com/zipzoft/supporter-go"

// Check string is empty.
// Output: true
fmt.Println(supporter.IsEmpty(""))

// Match string with regex.
// Output: map[string][]string{"name": {"zipzoft"}, "package": {"supporter-go"}}
fmt.Println(supporter.MatchGroupsAllSub(`(?P<name>\w+)/(?P<package>\w+)`, "zipzoft/supporter-go"))
```


---


### Dict (Map)
```go
import "github.com/zipzoft/supporter-go"

target := map[string]interface{
    "name": "zipzoft",
    "package": "supporter-go",
    "developers" : []map[string]interface{}{
        {"name": "zipzoft", "age": "30"},
        {"name": "PA", "age": "20"},
    },
}

// Get value from dict.
// Output: zipzoft
fmt.Println(supporter.DataGet(target, "name"))

// Get nested value from dict.
// Output: 20
fmt.Println(supporter.DataGet(target, "developers.1.age"))
```

