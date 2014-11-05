package df

import "encoding/xml"

func Unmarshal(data []byte) (*Objects, error) {
	v := &Objects{}

	if err := xml.Unmarshal(data, v); err != nil {
		return nil, err
	}

	return v, nil
}
