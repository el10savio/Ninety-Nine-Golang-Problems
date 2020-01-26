package listscontinued

import (
	"github.com/el10savio/Ninety-Nine-Golang-Problems/lists"
)

func EncodeModified(list []string) ([]lists.RLE, error) {
	packed, err := lists.Pack(list)
	if err != nil {
		return []lists.RLE{}, lists.ErrEmptyList
	}

	encoded := make([]lists.RLE, 0)

	for _, element := range packed {
		encoded = append(encoded, lists.RLE{Count: len(element), Value: string(element[0])})
	}

	return encoded, nil
}
