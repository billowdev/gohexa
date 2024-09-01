package utils

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/billowdev/exclusive-go-hexa/internal/adapters/database/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// NowInUTCOffset returns the current time in the specified UTC offset timezone.
//
// The function dynamically sets the timezone offset based on the provided hours from UTC
// and converts the current local time to the corresponding timezone.
// It is useful for applications that need to consistently operate in or display times
// for different timezones.
//
// Usage example:
// - NowInUTCOffset(7) returns the current time in the UTC+7 timezone.
func NowInUTCOffset(offset int) (time.Time, error) {
	var location *time.Location

	// Switch case to handle different UTC offsets
	switch offset {
	case -12:
		location = time.FixedZone("UTC-12", -12*60*60)
	case -11:
		location = time.FixedZone("UTC-11", -11*60*60)
	case -10:
		location = time.FixedZone("UTC-10", -10*60*60)
	// Add more cases as needed for each timezone
	case 0:
		location = time.UTC
	case 1:
		location = time.FixedZone("UTC+1", 1*60*60)
	case 2:
		location = time.FixedZone("UTC+2", 2*60*60)
	// Continue adding cases for each timezone you need
	case 7:
		location = time.FixedZone("UTC+7", 7*60*60)
	case 8:
		location = time.FixedZone("UTC+8", 8*60*60)
	case 9:
		location = time.FixedZone("UTC+9", 9*60*60)
	case 10:
		location = time.FixedZone("UTC+10", 10*60*60)
	case 11:
		location = time.FixedZone("UTC+11", 11*60*60)
	case 12:
		location = time.FixedZone("UTC+12", 12*60*60)
	default:
		return time.Time{}, fmt.Errorf("unsupported UTC offset: %d", offset)
	}

	// Get the current local time and convert it to the specified timezone
	return time.Now().In(location), nil
}

// NowInUTCPlus7 returns the current time in the UTC+7 timezone.
//
// This function sets a fixed timezone offset of +7 hours from UTC
// and converts the current local time to this timezone. It is
// useful for applications that need to consistently operate in
// or display times for the UTC+7 timezone.
func NowInUTCPlus7() time.Time {
	// Define the UTC+7 timezone with a fixed offset of +7 hours from UTC
	location := time.FixedZone("UTC+7", 7*60*60)

	// Get the current local time and convert it to the UTC+7 timezone
	return time.Now().In(location)
}

// IsSlice checks if the given value is a slice.
func IsSlice(v interface{}) bool {
	rv := reflect.ValueOf(v)
	return rv.Kind() == reflect.Slice
}

// String to Int Conversion
// Example usage:
// num, err := StringToInt("123")
//
//	if err != nil {
//	    fmt.Println("Error converting string to int:", err)
//	} else {
//
//	    fmt.Println("Converted number:", num)
//	}
func StringToInt(str string) (int, error) {
	return strconv.Atoi(str)
}

// String to Float Conversion
// Example usage:
// f, err := StringToFloat("123.45")
//
//	if err != nil {
//	    fmt.Println("Error converting string to float:", err)
//	} else {
//
//	    fmt.Println("Converted float:", f)
//	}
func StringToFloat(str string) (float64, error) {
	return strconv.ParseFloat(str, 64)
}

// Int to String Conversion
// Example usage:
// str := IntToString(123)
// fmt.Println("Converted string:", str)
func IntToString(num int) string {
	return strconv.Itoa(num)
}

// Helper function to check if a string is in a slice of strings
// Check if String is in Slice
// Example usage:
// slice := []string{"apple", "banana", "cherry"}
// contains := StringContains(slice, "banana")
// fmt.Println("Slice contains 'banana':", contains)
func StringContains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// Join Int Slice to String
// Example usage:
// ints := []int{1, 2, 3, 4}
// joined := JoinInts(ints, ", ")
// fmt.Println("Joined string:", joined)
func JoinInts(ints []int, sep string) string {
	var strSlice []string
	for _, num := range ints {
		strSlice = append(strSlice, strconv.Itoa(num))
	}
	return strings.Join(strSlice, sep)
}

// Float to String Conversion
// Example usage:
// floatStr := FloatToString(123.456)
// fmt.Println("Converted string:", floatStr)
func FloatToString(num float64) string {
	return strconv.FormatFloat(num, 'f', -1, 64)
}

// Check if Int is in Slice
// Example usage:
// intSlice := []int{10, 20, 30, 40}
// intContains := IntContains(intSlice, 20)
// fmt.Println("Slice contains 20:", intContains)
func IntContains(slice []int, element int) bool {
	for _, item := range slice {
		if item == element {
			return true
		}
	}
	return false
}

// Trim String Spaces
// Example usage:
// trimmed := TrimSpaces("  hello world  ")
// fmt.Println("Trimmed string:", trimmed)
func TrimSpaces(str string) string {
	return strings.TrimSpace(str)
}

// Convert String to Boolean
// Example usage:
// b, err := StringToBool("true")
//
//	if err != nil {
//	    fmt.Println("Error converting string to bool:", err)
//	} else {
//
//	    fmt.Println("Converted boolean:", b)
//	}
func StringToBool(str string) (bool, error) {
	return strconv.ParseBool(str)
}

// Format Unix Time to String
// Example usage:
// formattedTime := FormatUnixTime(1609459200, "2006-01-02 15:04:05")
// fmt.Println("Formatted time:", formattedTime)
func FormatUnixTime(unixTime int64, layout string) string {
	return time.Unix(unixTime, 0).Format(layout)
}

// StructToMap takes an interface representing a struct and converts it into a map[string]interface{}.
//
// This function is useful when you have a struct and you want to convert it into a map where the keys are the field names of the struct, and the values are the corresponding field values. It is particularly helpful when you need to work with data in a more dynamic or generic way.
//
// Parameters:
// - item: interface{} - The input value that should be a struct. It accepts an interface to be more flexible in what types can be passed.
//
// Returns:
// - map[string]interface{}: A map representation of the struct where field names are the keys and field values are the values.
//
// Usage Example:
// Suppose you have a struct named 'Person' like this:
//
//	type Person struct {
//	    Name  string
//	    Age   int
//	    Email string
//	}
//
// You can convert an instance of 'Person' into a map using StructToMap like this:
// person := Person{Name: "John", Age: 30, Email: "john@example.com"}
// personMap := StructToMap(person)
//
// The 'personMap' will contain:
//
//	map[string]interface{}{
//	    "Name":  "John",
//	    "Age":   30,
//	    "Email": "john@example.com",
//	}
//
// Note:
// - Only exported (public) fields of the struct can be converted, as unexported (private) fields cannot be interfaced.
// - The function uses reflection to achieve this, so keep in mind that it may have a performance overhead for large structs.
// - It's important to ensure that the input 'item' is indeed a struct, as non-struct types will result in an empty map.
func StructToMap(item interface{}) map[string]interface{} {
	// Initialize an empty map to hold the result
	result := make(map[string]interface{})

	// Reflect on the interface value to get its runtime representation
	itemValue := reflect.ValueOf(item)

	// Reflect on the type of the interface to get type information
	itemType := reflect.TypeOf(item)

	// Check if the provided interface is actually a struct
	// If not, return the empty map
	if itemValue.Kind() != reflect.Struct {
		return result
	}

	// Iterate over all fields of the struct
	for i := 0; i < itemValue.NumField(); i++ {
		// Get the value of the field
		field := itemValue.Field(i)

		// Get the name of the field
		fieldName := itemType.Field(i).Name

		// Check if the field is exported and can be interfaced
		// Private (unexported) fields can't be interfaced
		if field.CanInterface() {
			// Add the field to the result map using its name as the key
			// and the value of the field as the value
			result[fieldName] = field.Interface()
		}
	}

	// Return the constructed map
	return result
}

// SplitString splits a string into an array of substrings based on a delimiter.
func SplitString(input, delimiter string) []string {
	return strings.Split(input, delimiter)
}

func GetLocalAsString(c *fiber.Ctx, key string) string {
	value, ok := c.Locals(key).(string)
	if !ok {
		return "" // or handle the error as needed
	}
	return value
}

func ParseDateParam(dateStr string) time.Time {
	if dateStr == "" {
		return time.Time{} // Return a zero time for empty string
	}

	parsedTime, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		// Handle the error here, depending on your application logic
		// For example, log the error or return a default value
		return time.Time{} // Return a zero time for invalid input
	}

	return parsedTime
}

func MarshalJSONB(data models.JSONB) ([]byte, error) {
	return json.Marshal(data)
}

// func UnmarshalJSON(data []byte, target interface{}) error {
// 	return json.Unmarshal(data, target)
// }

func UnmarshalJSON(data interface{}, target interface{}) error {
	var jsonData []byte
	var err error

	// Check if data is already a []byte
	if bytesData, ok := data.([]byte); ok {
		jsonData = bytesData
	} else {
		// If data is not a []byte, marshal it to []byte
		jsonData, err = json.Marshal(data)
		if err != nil {
			return err
		}
	}

	// Unmarshal the JSON data into the target interface{}
	return json.Unmarshal(jsonData, target)
}

// MapQueryParamsToMap maps query parameters to a map[string]interface{}
func MapQueryParamsToMap(c *fiber.Ctx, filterStruct interface{}) map[string]interface{} {
	// Get the reflect.Type of the struct
	structType := reflect.TypeOf(filterStruct).Elem()

	// Create a map to hold query parameter names and values
	paramMap := make(map[string]interface{})

	// Map query parameter names to values
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		paramName := c.Query(field.Tag.Get("json"))
		if paramName != "" {
			paramMap[field.Name] = c.Query(paramName)
		}
	}

	return paramMap
}

func DecodeBase64(base64Str string) ([]byte, error) {
	// Decode the base64 string into binary data
	data, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func HandleTransactionError(tx *gorm.DB, err error) error {
	rollBackErr := tx.Rollback().Error
	if rollBackErr != nil {
		return err
	}
	return err
}

func ExtractImageTypeFromBase64(dataURI string) (string, error) {
	// Check if the data URI starts with "data:image/". If not, return an error.
	if !strings.HasPrefix(dataURI, "data:image/") {
		return "", fmt.Errorf("invalid data URI format")
	}

	// Find the end of the image type declaration (e.g., "data:image/jpeg;base64,")
	endIndex := strings.Index(dataURI, ";base64,")
	if endIndex == -1 {
		return "", fmt.Errorf("invalid data URI format")
	}

	// Extract and return the image type.
	imageType := dataURI[len("data:image/"):endIndex]
	return imageType, nil
}

func DebugConsole(feature string, message string) {
	log.Println("--------------------------")
	log.Printf("\n%s\n", feature)
	log.Println(message)
	log.Println("--------------------------")
}

func StructToMapWithFilter(inputStruct interface{}, popOutKeys []string) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	// Convert the struct to a map
	structValue := reflect.ValueOf(inputStruct)
	if structValue.Kind() == reflect.Ptr {
		structValue = structValue.Elem()
	}
	if structValue.Kind() != reflect.Struct {
		return nil, fmt.Errorf("input is not a struct or pointer to struct")
	}

	structType := structValue.Type()
	for i := 0; i < structValue.NumField(); i++ {
		fieldName := structType.Field(i).Name
		fieldValue := structValue.Field(i).Interface()

		// Check if the field name is in the list of keys to pop out
		popOut := false
		for _, key := range popOutKeys {
			if key == fieldName {
				popOut = true
				break
			}
		}

		// If it's not in the list, add it to the result map
		if !popOut {
			result[fieldName] = fieldValue
		}
	}

	return result, nil
}

func MapToStruct(data map[string]interface{}, targetStruct interface{}) error {
	targetValue := reflect.ValueOf(targetStruct)
	if targetValue.Kind() != reflect.Ptr || targetValue.IsNil() {
		return fmt.Errorf("targetStruct must be a pointer to a struct")
	}

	targetValue = targetValue.Elem()
	targetType := targetValue.Type()

	for i := 0; i < targetValue.NumField(); i++ {
		fieldName := targetType.Field(i).Name
		fieldValue := targetValue.Field(i)

		if value, ok := data[fieldName]; ok {
			if fieldValue.CanSet() {
				fieldReflectValue := reflect.ValueOf(value)

				// Check if the types are compatible
				if fieldValue.Type().AssignableTo(fieldReflectValue.Type()) {
					fieldValue.Set(fieldReflectValue)
				} else {
					return fmt.Errorf("incompatible types for field %s", fieldName)
				}
			}
		}
	}

	return nil
}

func convertPascalToSnakeCase(s string) string {
	var result []rune
	for i, r := range s {
		if i > 0 && unicode.IsUpper(r) {
			result = append(result, '_')
		}
		result = append(result, unicode.ToLower(r))
	}
	return string(result)
}

func ParseSubjectUUID(c *fiber.Ctx) (uuid.UUID, error) {
	subject := GetLocalAsString(c, "sub")
	if subject == "" {
		err := errors.New("user is not authorized")
		return uuid.UUID{}, err
	}

	createdBy, err := uuid.Parse(subject)
	if err != nil {
		err := errors.New("user is not authorized")
		return uuid.UUID{}, err
	}

	return createdBy, nil
}

func ToLowerCase(text string) string {
	return strings.ToLower(text)
}

// func ValidateBackofficePermission(role string) (bool, error) {
// 	if role in ["admin", "manager", "operation"] {
// 		return true
// 	}

// 	return false, nil
// }

func GetStringOrDefault(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}
