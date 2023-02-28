package validators

import (
	"fmt"
	"reflect"
	"strings"
	"vm-init-utils/windows/options"
)

func ValidateFlagSet(f options.Network) error {
	t := reflect.TypeOf(f)
	v := reflect.ValueOf(f)
	var (
		name, osTag string
		paramValue  reflect.Value
	)
	for i := 0; i < t.NumField(); i++ {
		osTag = t.Field(i).Tag.Get("required")
		if len(osTag) == 0 {
			continue
		}
		if strings.Contains(osTag, "true") {
			name = t.Field(i).Name
			paramValue = v.FieldByName(name)
			switch paramValue.Kind().String() {
			case "string":
				if value := strings.ReplaceAll(paramValue.String(), " ", ""); len(value) == 0 {
					return fmt.Errorf("[%s] is required, please notify a valid value to this param", name)
				}
			default:
				return fmt.Errorf("Unknown param kind ")
			}
		}
	}
	return nil
}
