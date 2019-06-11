package case_conversion

import "testing"

var keyWords = "to_UpperCamel-case"

func TestToCamelCase(t *testing.T) {
	t.Logf("%s", ToCamelCase(keyWords))
}

func TestToUpperCamelCase(t *testing.T) {
	t.Logf("%s", ToUpperCamelCase(keyWords))
}

func TestToLowerCamelCase(t *testing.T) {
	t.Logf("%s", ToLowerCamelCase(keyWords))
}

func TestToUpperFirst(t *testing.T) {
	t.Logf("%s", ToUpperFirst(keyWords))
}

func TestToUpperSnakeCase(t *testing.T) {
	t.Logf("%s", ToUpperSnakeCase(keyWords))
}

func TestToLowerSnakeCase(t *testing.T) {
	t.Logf("%s", ToLowerSnakeCase(keyWords))
}

func TestToLowerLinkCase(t *testing.T) {
	t.Logf("%s", ToLowerLinkCase(keyWords))
}
