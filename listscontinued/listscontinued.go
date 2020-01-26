package listscontinued

import (
	"github.com/el10savio/Ninety-Nine-Golang-Problems/lists"
)

// Can't get it to omit struct value if not present
func EncodeModified(list []string) ([]lists.RLE, error) {
	packed, err := lists.Pack(list)
	if err != nil {
		return []lists.RLE{}, lists.ErrEmptyList
	}

	encoded := make([]lists.RLE, 0)

	for _, element := range packed {
		if len(element) > 1 {
			encoded = append(encoded, lists.RLE{Count: len(element), Value: string(element[0])})
		} else {
			encoded = append(encoded, lists.RLE{Value: string(element[0])})
		}
	}

	return encoded, nil
}

func DecodeModified(rleList []lists.RLE) ([]string, error) {
	if len(rleList) < 1 {
		return []string{}, lists.ErrEmptyList
	}

	decoded := make([]string, 0)

	for _, element := range rleList {
		decoded = append(decoded, element.Value)
		for index := 1; index < element.Count; index++ {
			decoded = append(decoded, element.Value)
		}
	}

	return decoded, nil
}
