package adl

import "fmt"

type Json interface{}
type JsonArray struct {
	JsonElems []Json
}
type JsonObj struct {
	JsonObjs
}
type JsonNull struct{}
type JsonObjs map[string]Json

func (a JsonNull) String() string  { return "null" }
func (a JsonArray) String() string { return fmt.Sprintf("%v", a.JsonElems) }
