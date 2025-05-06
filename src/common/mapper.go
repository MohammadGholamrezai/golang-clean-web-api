package common

import "encoding/json"

// Fetch data as input and convert to T type and return it as output
func TypeConvertor[T any](data any) (*T, error) {
	var result T
	dataJson, err := json.Marshal(&data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(dataJson, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
