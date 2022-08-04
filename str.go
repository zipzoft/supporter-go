package supporter

import (
	crypto_rand "crypto/rand"
	"math/rand"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

// IsEmpty check if the value is empty
// Example:
// 		// Output: true
// 		fmt.Println(IsEmpty(""))
//
// 		// Output: true
// 		fmt.Println(IsEmpty(" "))
//
// 		// Output: false
// 		fmt.Println(IsEmpty("a"))
//
// 		// Output: false
// 		fmt.Println(IsEmpty(1))
//
// 		// Output: true
// 		fmt.Println(IsEmpty(0))
//
// 		// Output: true
// 		fmt.Println(IsEmpty(nil))
//
// 		// Output: true
// 		fmt.Println(IsEmpty(false))
func IsEmpty(val interface{}) bool {
	if val == nil {
		return true
	}

	reflectValue := reflect.ValueOf(val)
	reflectKind := reflectValue.Kind()

	if reflectKind == reflect.String {
		// Trim space
		return strings.TrimSpace(val.(string)) == ""
	}

	if reflectKind == reflect.Slice {
		return reflectValue.Len() == 0
	}

	if reflectKind == reflect.Map {
		return reflectValue.Len() == 0
	}

	if reflectValue.IsZero() || reflectValue.IsNil() {
		return true
	}

	return false
}

// IsNotEmpty check if the value is not empty
func IsNotEmpty(val interface{}) bool {
	return !IsEmpty(val)
}

// MatchGroupsAllSub match all groups
// Example:
// 		pattern := "^(?P<name>\\w+)\\s(?P<age>\\d+)$"
// 		text := "John 23"
// 		matched := supporter.MatchGroupsAllSub(pattern, text)
// 		fmt.Println(matched)
// 		// Output: map[name: []string{"John"} age: []string{"23"}]
func MatchGroupsAllSub(pattern string, text string) map[string][]string {
	matched := make(map[string][]string)

	if pattern == "" || text == "" {
		return matched
	}

	re := regexp.MustCompile(pattern)
	groups := re.SubexpNames()
	matches := re.FindAllStringSubmatch(text, -1)

	for groupIdx, group := range groups {
		if group == "" {
			continue
		}

		matched[group] = make([]string, 0)

		for _, match := range matches {
			matched[group] = append(matched[group], match[groupIdx])
		}
	}

	return matched
}

// MatchGroups match groups and get first matched value
// Example:
// 		pattern := "^(?P<name>\\w+)\\s(?P<age>\\d+)$"
// 		text := "John 23"
//
// 		matched := supporter.MatchGroups(pattern, text)
// 		fmt.Println(matched)
// 		// Output: map[name: "John" age: "23"]
func MatchGroups(pattern string, text string) map[string]string {
	matched := MatchGroupsAllSub(pattern, text)

	paramMap := make(map[string]string)

	for name, strings := range matched {
		if val, ok := First(strings).(string); ok {
			paramMap[name] = val
		}
	}

	return paramMap
}

// StrRandom string
// @ref: https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
// Example:
// 		fmt.Println(supporter.StrRandom(10))
// 		// Output: string
func StrRandom(n int) string {
	var src = rand.NewSource(time.Now().UnixNano())
	const (
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

func StrRandomNumber() string {
	number, _ := crypto_rand.Prime(crypto_rand.Reader, 32)
	return strconv.Itoa(int(number.Int64()))
}

func StrToInteger(val string) int {
	var newVal int
	port, _ := strconv.ParseInt(val, 10, 64)
	newVal = int(port)

	return newVal
}

func StrIsNumeric(value string) bool {
	_, err := strconv.Atoi(value)
	return err == nil
}
