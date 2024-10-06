package config

import (
	"os"
	"reflect"
)

func LoadConfig(cfg interface{}) error {
	return populateFields(reflect.ValueOf(cfg).Elem())
}

func populateFields(v reflect.Value) error {
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := v.Type().Field(i)

		if field.Kind() == reflect.Struct {
			if err := populateFields(field); err != nil {
				return err
			}
		} else {
			envTag := fieldType.Tag.Get("env")
			if envTag != "" {
				value := getEnv(envTag, "")
				if value != "" {
					field.SetString(value)
				}
			}
		}
	}
	return nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}