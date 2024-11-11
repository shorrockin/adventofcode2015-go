package day12

import (
	"adventofcode2015/pkg/assert"
	"encoding/json"
	"fmt"
)

func Solve(jsonValue string, ignoreRed bool) (float64, error) {
	var data interface{}
	assert.NoError(json.Unmarshal([]byte(jsonValue), &data), "could not unmarshal json")

	out, _, err := sumNumbers(data, ignoreRed)
	return out, err
}

func sumNumbers(input interface{}, ignoreRed bool) (float64, bool, error) {
	switch data := input.(type) {
	case float64:
		return data, false, nil

	case []interface{}:
		sum := float64(0)
		for _, element := range data {
			elementValue, _, err := sumNumbers(element, ignoreRed)
			if err != nil {
				return sum, false, err
			}
			sum += elementValue
		}
		return sum, false, nil

	case map[string]interface{}:
		sum := float64(0)
		for _, element := range data {
			elementValue, ignored, err := sumNumbers(element, ignoreRed)
			if err != nil {
				return sum, false, err
			}
			if ignored {
				return 0, false, err
			}
			sum += elementValue
		}
		return sum, false, nil

	case string:
		switch data {
		case "red":
			return 0, ignoreRed, nil
		default:
			return 0, false, nil
		}

	default:
		return 0, false, fmt.Errorf("unable to count type: %+v, %T", input, input)
	}
}
