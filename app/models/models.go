package models

import (
	"encoding/json"
	"io"
)

type Model struct {
}

// ToJSON serializes the contents of the collection to JSON
// NewEncoder provides better performance than json.Unmarshal as it does not
// have to buffer the output into an in memory slice of bytes
// this reduces allocations and the overheads of the service
//
// https://golang.org/pkg/encoding/json/#NewEncoder
func (m *Model) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(m)
}

//FromJSON is like ToJSON but in reverse way
func (m *Model) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(m)
}

type User struct {
	Model
	Mobile         string  `json:"mobile"`
	Credit         float64 `json:"credit"`
	ReceivedCharge bool    `json:"receivedCharge"`
}

type Code struct {
	Model
	Code          string  `json:"code"`
	Credit        float64 `json:"credit"`
	ConsumerCount int     `json:"consumerCount"`
}
