package mi

import "encoding/xml"

func Unmarshal(data []byte) (*OmiEnvelope, error) {
	v := &OmiEnvelope{}

	if err := xml.Unmarshal(data, v); err != nil {
		return nil, err
	}

	return v, nil
}
