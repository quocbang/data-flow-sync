package diff

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// flag definition.
type flags int

const (
	NO_CHANGE flags = iota + 1
	ADD
	DELETE
	BOTH
)

// FindDiff is find different of file x and y,
// input is []byte and a generic type and return information of added,
// deleted between file x and file y.
// the generic type should is a struct and that has yaml tag.
func FindDiff[T any](x, y []byte) ([]byte, []byte, error) {
	var yamlX T
	if err := json.Unmarshal(x, &yamlX); err != nil {
		return nil, nil, err
	}

	var yamlY T
	if err := json.Unmarshal(y, &yamlY); err != nil {
		return nil, nil, err
	}

	var (
		add    reflect.Value
		delete reflect.Value
		err    error
	)
	switch reflect.TypeOf(*new(T)).Kind() {
	case reflect.Struct:
		add, delete, _, err = handleStruct(yamlX, yamlY)
		if err != nil {
			return nil, nil, err
		}
	case reflect.Slice:
		add, delete, err = handleSlice(yamlX, yamlY)
		if err != nil {
			return nil, nil, err
		}
	default:
		return nil, nil, fmt.Errorf("wrong type during find diff, expect type is [struct or slice] but found [%v]", reflect.TypeOf(*new(T)).Kind())
	}

	// marshal struct to []byte and remove all empty fields use OMITEMPTY
	added, err := json.Marshal(add.Interface())
	if err != nil {
		return nil, nil, err
	}
	deleted, err := json.Marshal(delete.Interface())
	if err != nil {
		return nil, nil, err
	}

	return added, deleted, nil
}

// handleStruct find different between x struct and y struct.
func handleStruct(x, y any) (reflect.Value, reflect.Value, flags, error) {
	valueX := reflect.ValueOf(x)
	valueY := reflect.ValueOf(y)
	typeX := reflect.TypeOf(x) // TODO: use for check tag

	// define delete and add with nil pointer struct
	var (
		delete = reflect.New(valueX.Type()).Elem()
		add    = reflect.New(valueY.Type()).Elem()
	)
	isIndexValueEqual, iStructDiff := false, false
	flag := NO_CHANGE
	// for loop to check each field of the struct
	// type A struct {
	//   name string
	//   age int
	// }
	// if value of A.name == A'.name => do nothing(it's mean delete and add still nil at this field).
	// if value of A.name != A'.name => set x.interface == delete y.interface == add
	// continue check until out of range
	for i := 0; i < valueX.NumField(); i++ {
		fieldX := valueX.Field(i)
		fieldY := valueY.Field(i)
		fieldTypeX := typeX.Field(i)

		// isHold is tag of struct that tell for loop to hold
		// the field caller set in struct tag
		// type A struct {
		//   name string `diff:"hold"`
		//   age int
		// }
		// *note: hold is like primaryKey that can hold your key when one of all field in your struct changed
		isHold := fieldTypeX.Tag.Get("diff") == "hold"
		if isHold {
			// isIndexValueEqual if true that mean x struct and struct is same struct false if is not.
			isIndexValueEqual = fieldX.Interface() == fieldY.Interface()
		}

		// check equal of x and y
		if !reflect.DeepEqual(fieldX.Interface(), fieldY.Interface()) {
			iStructDiff = true // mark struct is change

			// if field x is nil => added
			if fieldX.IsZero() && !fieldY.IsZero() {
				add.Field(i).Set(fieldY)
				continue
			}

			if t := fieldX.Kind(); t == reflect.Slice || t == reflect.Struct {
				switch t {
				case reflect.Struct:
					to, from, f, err := handleStruct(fieldX.Interface(), fieldY.Interface())
					if err != nil {
						return reflect.Value{}, reflect.Value{}, NO_CHANGE, err
					}

					// check return flag(ref to flags)
					// if f > flag and it is the same struct then set flag = f
					if f > flag && isIndexValueEqual {
						flag = f
					}
					delete.Field(i).Set(from)
					add.Field(i).Set(to)
				case reflect.Slice:
					to, from, err := handleSlice(fieldX.Interface(), fieldY.Interface())
					if err != nil {
						return reflect.Value{}, reflect.Value{}, NO_CHANGE, err
					}
					// check flag
					if !to.IsNil() || !from.IsNil() {
						if f := checkNilSlice(from, to); f > flag {
							flag = f
						}
					}
					delete.Field(i).Set(from)
					add.Field(i).Set(to)
				default:
					return reflect.Value{}, reflect.Value{}, NO_CHANGE, fmt.Errorf("type [%v] hasn't implementation yet!", t)
				}
				continue
			}

			// check flag
			if f := checkFlagsWithZero(fieldX, fieldY); f > flag && isIndexValueEqual {
				flag = f
			}
			delete.Field(i).Set(fieldX)
			add.Field(i).Set(fieldY)
		} else {
			// this func check at the end, ensure that struct is same struct
			// and that struct is changed and flag == BOTH then we hold primaryKey if not still null
			defer func(idx int) {
				if iStructDiff && isIndexValueEqual && flag == BOTH {
					add.Field(idx).Set(fieldY)
					delete.Field(idx).Set(fieldX)
				}
			}(i)
			// set null if x and y no change
			delete.Field(i).SetZero()
			add.Field(i).SetZero()
		}
	}

	return add, delete, flag, nil
}

// handleSlice find different between x slice and y slice.
func handleSlice(x, y any) (reflect.Value, reflect.Value, error) {
	valueX := reflect.ValueOf(x)
	valueY := reflect.ValueOf(y)
	// for loop check check each element of the slice
	// if element x[i] equal one of all y[0:len(y)] => remove both in the slice if not still hold both
	// until out of the range so we have x == delete and y == add
	for i := 0; i < valueX.Len(); i++ {
		x := valueX.Index(i)
		if t := x.Kind(); t == reflect.Struct || t == reflect.Slice {
			switch t {
			case reflect.Struct:
				// try compare x[i] to all of y, if x[i] equal any element of y then remove both in the slice
				for j := 0; j < valueY.Len(); j++ {
					y := valueY.Index(j)
					// break because this was compared or is is a nil value
					if x.IsZero() {
						break
					}

					// continue because this value was removed by last step
					if y.IsZero() {
						continue
					}

					to, from, flag, err := handleStruct(x.Interface(), y.Interface())
					if err != nil {
						return reflect.Value{}, reflect.Value{}, err
					}
					if from.IsZero() && to.IsZero() { // add and delete isZero => both no changed
						x.SetZero()
						y.SetZero()
					} else if flag == DELETE {
						y.SetZero()
						break
					} else if flag == ADD {
						x.SetZero()
						break
					}
				}
			default:
				return reflect.Value{}, reflect.Value{}, fmt.Errorf("type [%v] in handleSlice() han't implementation yet!", t)
			}
			continue
		} else {
			// remove if x[i] equal y[j] element.
			for j := 0; j < valueY.Len(); j++ {
				if reflect.DeepEqual(x.Interface(), valueY.Index(j).Interface()) && !x.IsZero() && !valueY.Index(j).IsZero() {
					x.SetZero()
					valueY.Index(j).SetZero()
					break
				}
			}
		}
	}

	return valueY, valueX, nil
}

// check flag based x and data
// BOTH: x = nil && y == nil
// ADD: x = nil && y != nil
// DELETE: x != nil && y == nil
func checkFlagsWithZero(x, y reflect.Value) flags {
	if x.IsZero() && y.IsZero() {
		return NO_CHANGE
	} else if x.IsZero() && !y.IsZero() {
		return ADD
	} else if !x.IsZero() && y.IsZero() {
		return DELETE
	} else {
		return BOTH
	}
}

// check all element in x and y
// if one of x != nil or != {} then DELETE
// if one of y != nil or != {} then ADD
// if is both DELETE and ADD then BOTH
func checkNilSlice(x, y reflect.Value) flags {
	xFlag := NO_CHANGE
	for i := 0; i < x.Len(); i++ {
		value := x.Index(i)
		if value.Kind() == reflect.Struct {
			if !value.IsZero() {
				xFlag = DELETE
				break
			}
		} else if value.Kind() == reflect.Slice {
			if !value.IsNil() {
				xFlag = DELETE
				break
			}
		}
	}

	yFlag := NO_CHANGE
	for i := 0; i < y.Len(); i++ {
		value := y.Index(i)
		if value.Kind() == reflect.Struct {
			if !value.IsZero() {
				yFlag = DELETE
				break
			}
		} else if value.Kind() == reflect.Slice {
			if !value.IsNil() {
				yFlag = DELETE
				break
			}
		}
	}

	if xFlag != NO_CHANGE && yFlag != NO_CHANGE {
		return BOTH
	} else if xFlag == NO_CHANGE && yFlag != NO_CHANGE {
		return ADD
	} else if xFlag != NO_CHANGE && yFlag == NO_CHANGE {
		return DELETE
	}

	return NO_CHANGE // no lint
}
