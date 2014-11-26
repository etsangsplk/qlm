package df

import "encoding/xml"

func Marshal(objects Objects) ([]byte, error) {
	objects.XmlnsXsi = "http://www.w3.org/2001/XMLSchema-instance"
	objects.NoNamespaceSchemaLocation = "odf.xsd"
	return xml.MarshalIndent(objects, "", "    ")
}
