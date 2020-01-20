package mypes

import (
	"strings"
)

// MSON Represents a type for unstructured JSONs
type MSON map[string]interface{}

// DeepGet ...
func (m MSON) DeepGet(path string) (value interface{}) {
	params := strings.Split(path, ".")

	lenParams := len(params)

	var tmpInterface interface{}

	for i := 0; i < lenParams; i++ {
		param := params[i]
		if i == 0 {
			tmpInterface = m[param]
		} else {
			if tmpInterface == nil {
				return "-"
			}
			tmpMap := tmpInterface.(MSON)
			tmpInterface = tmpMap[param]
		}

		if i == lenParams-1 {
			value = tmpInterface
		}
	}
	return
}
