package qlm

import "encoding/xml"

func Unmarshal(data []byte) (*QlmEnvelope, error) {
	v := &QlmEnvelope{}

	if err := xml.Unmarshal(data, v); err != nil {
		return nil, err
	}

	return v, nil
}
