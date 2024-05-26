package entities

import (
	"encoding/json"
	"strconv"
)

type AnswerKey int

const (
	First AnswerKey = iota
	Second
	Third
	Fourth
)

func (key *AnswerKey) UnmarshalJSON(data []byte) error {
	var value string
	err := json.Unmarshal(data, &value)
	if err != nil {
		return err
	}

	toInt, err := strconv.Atoi(value)
	if err != nil {
		return err
	}

	*key = AnswerKey(toInt)
	return nil
}
