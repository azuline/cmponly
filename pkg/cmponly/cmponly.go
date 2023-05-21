package cmponly

import (
	"reflect"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func Fields[S any](structType S, fields ...string) cmp.Option {
	// Get all fields on structType.
	reflectType := reflect.TypeOf(structType)
	if reflectType.Kind() == reflect.Ptr {
		reflectType = reflectType.Elem()
	}
	var fieldsOnStruct []string
	for i := 0; i < reflectType.NumField(); i++ {
		fieldsOnStruct = append(fieldsOnStruct, reflectType.Field(i).Name)
	}

	// Construct the slice of fields to ignore by removing the fields
	// that should be compared (the fields that were passed in).
	var fieldsToIgnore []string
	for _, fieldName := range fieldsOnStruct {
		if !contains(fields, fieldName) {
			fieldsToIgnore = append(fieldsToIgnore, fieldName)
		}
	}

	return cmpopts.IgnoreFields(structType, fieldsToIgnore...)
}

func contains[E comparable](haystack []E, needle E) bool {
	for _, e := range haystack {
		if e == needle {
			return true
		}
	}
	return false
}
