/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package validate

import (
	"fmt"
	"reflect"

	"sigs.k8s.io/kind/pkg/errors"
)

func validateStruct(s interface{}) (err error) {
	// Print information in different lines: // Added by JANR
	// Relative path: // Added by JANR
	// Brief function goal: // Added by JANR
	// All functions called in order: // Added by JANR
	fmt.Println("File(9) Step(1) Path: Skind/pkg/cluster/internal/validate/utils.go - Function: validateStruct()")        // Added by JANR
	fmt.Println("File(9) Step(1) Brief function goal: Validates the input struct by checking if all its fields are set.") // Added by JANR
	// first make sure that the input is a struct
	// having any other type, especially a pointer to a struct,
	// might result in panic
	structType := reflect.TypeOf(s)
	fmt.Println("File(9) Step(1) - Print - structType: ", structType) // Added by JANR
	if structType.Kind() != reflect.Struct {
		return errors.New("input param should be a struct")
	}

	// now go one by one through the fields and validate their value
	structVal := reflect.ValueOf(s)  // This structVal is a reflect.Value type, which means it is a copy of the struct s. // Added by JANR
	fieldNum := structVal.NumField() // This fieldNum is an int type, which means it is the number of fields in the struct s. // Added by JANR

	for i := 0; i < fieldNum; i++ {
		field := structVal.Field(i)
		fieldName := structType.Field(i).Name
		// Print information in different lines: // Added by JANR
		fmt.Println("File(9) Step(2) - Print - field: ", field)         // Added by JANR
		fmt.Println("File(9) Step(2) - Print - fieldName: ", fieldName) // Added by JANR

		isSet := field.IsValid() && !field.IsZero()
		// Print information in different lines: // Added by JANR
		fmt.Println("File(9) Step(2) - Print - isSet: ", isSet) // Added by JANR

		if !isSet {
			err = errors.New(fmt.Sprintf("%v%s in not set; ", err, fieldName))
		}
	}

	return err
}

func convertToMapStringString(m map[string]interface{}) map[string]string { // This func converts a map[string]interface{} to a map[string]string. // Added by JANR
	// Print information in different lines: // Added by JANR
	// Relative path: // Added by JANR
	// Brief function goal: // Added by JANR
	// All functions called in order: // Added by JANR
	fmt.Println("File(9) Step(3) Path: Skind/pkg/cluster/internal/validate/utils.go - Function: convertToMapStringString()") // Added by JANR
	fmt.Println("File(9) Step(3) Brief function goal: Converts a map[string]interface{} to a map[string]string.")            // Added by JANR
	var m2 = map[string]string{}                                                                                             // This m2 is a map[string]string type, which means it is an empty map. // Added by JANR
	// Print information in different lines: // Added by JANR
	fmt.Println("File(9) Step(3) - Print - m: ", m) // Added by JANR
	for k, v := range m {
		// Print information in different lines: // Added by JANR
		fmt.Println("File(9) Step(3) - Print - k: ", k) // Added by JANR
		m2[k] = v.(string)                              // This line assigns the value of the map m to the map m2. // Added by JANR
	}
	return m2
}

func getFieldNames(s interface{}) []string {
	var fieldNames []string
	structType := reflect.TypeOf(s)
	structVal := reflect.ValueOf(s)
	fieldNum := structType.NumField()
	for i := 0; i < fieldNum; i++ {
		field := structVal.Field(i)
		isSet := field.IsValid() && !field.IsZero()
		if isSet {
			fieldNames = append(fieldNames, structType.Field(i).Name)
		}
	}
	return fieldNames
}
