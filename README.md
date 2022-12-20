# Supporter GO (Helper functions)

## Install
```bash
go get -u github.com/zipzoft/supporter-go
```

## Usage

### Array

#### InArray
The InArray function checks if a given value is in an array. It returns a boolean indicating whether or not the value is in the array.

Here's an example of how to use InArray:

```go
import "github.com/zipzoft/supporter-go"

fmt.Println(supporter.InArray(1, []int{1, 2, 3})) // Output: true
fmt.Println(supporter.InArray(4, []int{1, 2, 3})) // Output: false
```

#### First
The First function returns the first element of an array or slice. If the array is empty or nil, it returns nil.

Here's an example of how to use First:
```go
fmt.Println(supporter.First([]int{1, 2, 3})) // Output: 1
```

#### ToSlice
The ToSlice function converts a value to a slice. It panics if the value is not a slice.

Here's an example of how to use ToSlice:
```go
arr := [3]int{1, 2, 3}
slice := supporter.ToSlice(arr)
fmt.Println(slice) // Output: [1, 2, 3]
```

--- 

### String
#### IsEmpty
The IsEmpty function checks if the given value is empty. It works with strings, slices, maps, and zero values of other types. It returns a boolean indicating whether or not the value is empty.

Here's an example of how to use IsEmpty:
```go
import "github.com/zipzoft/supporter-go"

fmt.Println(supporter.IsEmpty(""))  // Output: true
fmt.Println(supporter.IsEmpty(" ")) // Output: true
fmt.Println(supporter.IsEmpty("a")) // Output: false
fmt.Println(supporter.IsEmpty(1))   // Output: false
fmt.Println(supporter.IsEmpty(0))   // Output: true
fmt.Println(supporter.IsEmpty(nil)) // Output: true
fmt.Println(supporter.IsEmpty(false)) // Output: true
```


#### IsNotEmpty
The IsNotEmpty function is the opposite of IsEmpty. It returns a boolean indicating whether or not the given value is not empty.


#### MatchGroupsAllSub
The MatchGroupsAllSub function matches all groups in a regular expression pattern against a given text. It returns a map of the matched groups, where the keys are the group names and the values are slices of the matched strings.

Here's an example of how to use MatchGroupsAllSub:
```go
import "github.com/zipzoft/supporter-go"

pattern := "^(?P<name>\\w+)\\s(?P<age>\\d+)$"
text := "John 23"
matched := supporter.MatchGroupsAllSub(pattern, text)
fmt.Println(matched)
// Output: map[name: []string{"John"} age: []string{"23"}]
```


#### MatchGroups
The MatchGroups function is similar to MatchGroupsAllSub, but it returns a map with only the first matched value for each group.

Here's an example of how to use MatchGroups:
```go
import "github.com/zipzoft/supporter-go"

pattern := "^(?P<name>\\w+)\\s(?P<age>\\d+)$"
text := "John 23"

matched := supporter.MatchGroups(pattern, text)
fmt.Println(matched)
// Output: map[name: "John" age: "23"]
```

#### StrRandom
The StrRandom function generates a random string with a given length. It uses a combination of letters and numbers, and is generated using the current time as a seed.

Here's an example of how to use StrRandom:
```go
import "github.com/zipzoft/supporter-go"

fmt.Println(supporter.StrRandom(10))
// Output: a random string of length 10
```

#### StrRandomNumber
The StrRandomNumber function generates a random number as a string. It uses the crypto/rand package to generate a secure random number.

Here's an example of how to use StrRandomNumber:
```go
fmt.Println(supporter.StrRandomNumber())
// Output: a random number as a string
```

#### StrToInteger
The StrToInteger function converts a string to an integer. It returns the integer value of the string, or 0 if the string cannot be converted.

Here's an example of how to use StrToInteger:
```go
import "github.com/zipzoft/supporter-go"

fmt.Println(supporter.StrToInteger("123"))
// Output: 123
fmt.Println(supporter.StrToInteger("abc"))
// Output: 0
```

#### StrIsNumeric
The StrIsNumeric function checks if a string is numeric. It returns a boolean indicating whether or not the string can be converted to an integer.

Here's an example of how to use StrIsNumeric:
```go
import "github.com/zipzoft/supporter-go"

fmt.Println(supporter.StrIsNumeric("123"))
// Output: true
fmt.Println(supporter.StrIsNumeric("abc"))
// Output: false
```
---


### Dict (Map)

#### DataGet
The DataGet function allows you to retrieve an item from an array or object using "dot" notation. It supports nested keys, so you can retrieve values from nested objects or arrays by separating the keys with a period.

Here's an example of how to use DataGet:
```go
person := map[string]interface{}{
	"name": map[string]interface{}{
		"first": "John",
		"last":  "Doe",
	},
	"age": "23",
	"phones" : []map[string]interface{}{
		{"type": "home", "number": "555-555-5555"},
		{"type": "work", "number": "555-666-6666"},
	},
}

fmt.Println(supporter.DataGet(person, "name.first")) // Output: "John"
fmt.Println(supporter.DataGet(person, "name.last"))  // Output: "Doe"
fmt.Println(supporter.DataGet(person, "age"))        // Output: "23"
fmt.Println(supporter.DataGet(person, "phones.0.number")) // Output: "555-555-5555"
```
If the key is not found in the target, the function will return the default value provided as the last argument (if provided). If no default value is provided, the function will return nil.

