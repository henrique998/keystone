package database

import (
	"reflect"
	"strings"
)

func getTable(db *db, tableName string) *tableDefinition {
	return db.tables[tableName]
}

func capitalize(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}

func toPascalCase(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}
	parts := strings.Split(s, "_")
	for i, part := range parts {
		if len(part) > 0 {
			parts[i] = strings.ToUpper(part[:1]) + strings.ToLower(part[1:])
		}
	}
	return strings.Join(parts, "")
}

func toInterfaceSlice(v interface{}) []interface{} {
	rv := reflect.ValueOf(v)

	if rv.Kind() != reflect.Slice {
		return []interface{}{v}
	}

	out := make([]interface{}, rv.Len())

	for i := 0; i < rv.Len(); i++ {
		out[i] = rv.Index(i).Interface()
	}

	return out
}
