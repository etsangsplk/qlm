package mi

import "encoding/xml"

func Marshal(envelope OmiEnvelope) ([]byte, error) {
	root := struct {
		OmiEnvelope
		XMLName struct{} `xml:"omiEnvelope"`
	}{OmiEnvelope: envelope}
	return xml.MarshalIndent(root, "", "    ")
}
