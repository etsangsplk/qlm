package mi

import "encoding/xml"

func Marshal(envelope QlmEnvelope) ([]byte, error) {
	root := struct {
		QlmEnvelope
		XMLName struct{} `xml:"qlmEnvelope"`
	}{QlmEnvelope: envelope}
	return xml.MarshalIndent(root, "", "    ")
}
