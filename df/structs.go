package df

type Objects struct {
	Objects                   []Object `xml:"Object"`
	XmlnsXsi                  string   `xml:"xmlns:xsi,attr"`
	NoNamespaceSchemaLocation string   `xml:"xsi:noNamespaceSchemaLocation,attr"`
	Version                   string   `xml:"version,attr,omitempty"`
}

type Object struct {
	Type        string       `xml:"type,attr,omitempty"`
	Udef        string       `xml:"udef,attr,omitempty"`
	Id          *QLMID       `xml:"id"`
	Description *Description `xml:"description"`
	InfoItems   []InfoItem   `xml:"InfoItem"`
	Objects     []Object     `xml:"Object"`
}

type InfoItem struct {
	Udef        string       `xml:"udef,attr,omitempty"`
	Name        string       `xml:"name,attr"`
	Description *Description `xml:"description"`
	OtherNames  []string     `xml:"name"`
	MetaData    *MetaData
	Values      []Value `xml:"value"`
}

type Description struct {
	Lang string `xml:"lang,attr,omitempty"`
	Udef string `xml:"udef,attr,omitempty"`
	Text string `xml:",chardata"`
}

type QLMID struct {
	IdType    string `xml:"idType,attr,omitempty"`
	TagType   string `xml:"tagType,attr,omitempty"`
	StartDate string `xml:"startDate,attr,omitempty"`
	EndDate   string `xml:"endDate,attr,omitempty"`
	Udef      string `xml:"udef,attr,omitempty"`
	Text      string `xml:",chardata"`
}

type MetaData struct {
	InfoItems []InfoItem `xml:"InfoItem"`
}

type Value struct {
	Text     string `xml:",chardata"`
	Type     string `xml:"type,attr,omitempty"`
	DateTime string `xml:"dateTime,attr,omitempty"`
	UnixTime int64  `xml:"unixTime,attr,omitempty"`
}
