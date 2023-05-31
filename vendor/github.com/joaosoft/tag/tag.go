package tag

import (
	"reflect"
)

func Load(obj interface{}, tags ...string) Values {
	values := Values{}
	load(reflect.ValueOf(obj), values, tags...)
	return values
}

func load(obj reflect.Value, values Values, tags ...string) {
	if !obj.IsValid() {
		return
	}

	typ, value := getValue(obj)

	switch value.Kind() {
	case reflect.Struct:

		for i := 0; i < typ.NumField(); i++ {
			nextValue := value.Field(i)
			nextType := typ.Field(i)

			if nextValue.CanInterface() {
				for _, tag := range tags {
					tagValue, exists := nextType.Tag.Lookup(tag)
					if exists && tagValue != tagValueIgnoreField {
						values[tagValue] = nextValue.Interface()
					}
				}

				load(nextValue, values, tags...)
			}
		}

	case reflect.Array, reflect.Slice:
		for i := 0; i < value.Len(); i++ {
			nextValue := value.Index(i)

			if nextValue.CanInterface() {
				load(nextValue, values, tags...)
			}
		}

	case reflect.Map:
		for _, key := range value.MapKeys() {
			nextValue := value.MapIndex(key)

			if key.CanInterface() {
				load(key, values, tags...)
			}

			if nextValue.CanInterface() {
				load(nextValue, values, tags...)
			}
		}

	default:
		// do nothing ...
	}
}

func getValue(value reflect.Value) (reflect.Type, reflect.Value) {
	typ := value.Type()

again:
	if (value.Kind() == reflect.Ptr || value.Kind() == reflect.Interface) && !value.IsNil() {
		value = value.Elem()
		typ = value.Type()
		goto again
	}

	return typ, value
}
