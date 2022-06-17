package models

import (
	"encoding/json"
	"io"
)

type User struct {
	Mobile         string  `json:"mobile"`
	Credit         float64 `json:"credit"`
	ReceivedCharge bool    `json:"receivedCharge"`
}

type Code struct {
	Code          string  `json:"code"`
	Credit        float64 `json:"credit"`
	ConsumerCount int     `json:"consumerCount"`
}

func (m *Code) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(m)
}

//FromJSON is like ToJSON but in reverse way
func (m *Code) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(m)
}
