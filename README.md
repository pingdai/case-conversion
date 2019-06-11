# case-conversion
case string to lowerCamel/upperCamel/lowerSnake/upperSnake/lowerLink case

## Usage
```go
var keyWords = "to_UpperCamel-case"
```
```go
ToCamelCase(keyWords string) string         //  Output:To_uppercamel-case

ToUpperCamelCase(keyWords string) string    //  Output:ToUpperCamelCase

ToLowerCamelCase(keyWords string) string    //  Output:toUpperCamelCase

ToUpperFirst(keyWords string) string        //  Output:To_UpperCamel-case

ToUpperSnakeCase(keyWords string) string    //  Output:TO_UPPER_CAMEL_CASE

ToLowerSnakeCase(keyWords string) string    //  Output:to_upper_camel_case

ToLowerLinkCase(keyWords string) string     //  Output:to-upper-camel-case
```