package csv

import (
	"fmt"
	"github.com/capkeik/backend-trainee-assignment-2023/internal/model"
	"reflect"
	"time"
)

func Record(rec model.Record) ([]string, error) {
	var fields []string
	val := reflect.ValueOf(rec)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := val.Type().Field(i)

		if fieldType.Name == "ID" {
			continue
		}

		switch field.Kind() {
		case reflect.String:
			fields = append(fields, fmt.Sprintf("%s", field.Interface()))
		case reflect.Int, reflect.Int32, reflect.Uint, reflect.Uint32, reflect.Uint64:
			fields = append(fields, fmt.Sprintf("%d", field.Interface()))
		case reflect.Struct:
			if fieldType.Type == reflect.TypeOf(time.Time{}) {
				t := field.Interface().(time.Time)
				fields = append(fields, t.Format("2006-01-02 15:04:05"))
			}
		default:
			return nil, fmt.Errorf("unsupported field type: %v", field.Kind())
		}
	}

	return fields, nil
}
