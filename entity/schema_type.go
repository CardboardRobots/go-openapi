package entity

type SchemaType string

const TYPE_STRING SchemaType = "string"
const TYPE_NUMBER SchemaType = "number"
const TYPE_INTEGER SchemaType = "integer"
const TYPE_BOOLEAN SchemaType = "boolean"
const TYPE_ARRAY SchemaType = "array"
const TYPE_OBJECT SchemaType = "object"

func (s SchemaType) String() string {
	return string(s)
}

func NewSchemaType(value string) SchemaType {
	switch value {
	case string(TYPE_STRING):
		return TYPE_STRING
	case string(TYPE_NUMBER):
		return TYPE_NUMBER
	case string(TYPE_BOOLEAN):
		return TYPE_BOOLEAN
	case string(TYPE_ARRAY):
		return TYPE_ARRAY
	case string(TYPE_OBJECT):
		return TYPE_OBJECT
	}
	return TYPE_STRING
}

func IsSchemaType(value string) bool {
	switch value {
	case string(TYPE_STRING):
	case string(TYPE_NUMBER):
	case string(TYPE_BOOLEAN):
	case string(TYPE_ARRAY):
	case string(TYPE_OBJECT):
		return true
	}
	return false
}
