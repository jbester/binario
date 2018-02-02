package binaryio

import "fmt"

func Size(value interface{}) int {
	switch v := value.(type) {
	case byte:
		return 1

	case []byte:
		return len(v)

	case uint16:
		return 2

	case []uint16:
		return 2 * len(v)

	case int16:
		return 2

	case []int16:
		return 2 * len(v)

	case uint32:
		return 4

	case []uint32:
		return 4 * len(v)

	case int32:
		return 4

	case []int32:
		return 4 * len(v)

	case uint64:
		return 8

	case []uint64:
		return 8 * len(v)

	case int64:
		return 8

	case []int64:
		return 8 * len(v)

	default:
		panic(fmt.Errorf("Unsupported datatype %v", v))
	}
}
