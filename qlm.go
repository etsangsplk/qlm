package qlm

import "encoding/xml"

type Objects struct {
	Version string   `xml:"version,attr"`
	Objects []Object `xml:"Object"`
}

type Object struct {
	Type        string      `xml:"type,attr"`
	Udef        string      `xml:"udef,attr"`
	Id          QLMID       `xml:"id"`
	Description Description `xml:"description"`
	InfoItems   []InfoItem  `xml:"InfoItem"`
	Objects     []Object    `xml:"Object"`
}

type InfoItem struct {
	Udef        string      `xml:"udef,attr"`
	Name        string      `xml:"name,attr"`
	Description Description `xml:"description"`
	OtherNames  []string    `xml:"name"`
	MetaData    MetaData
	Values      []Value `xml:"value"`
}

type Description struct {
	Lang string `xml:"lang,attr"`
	Udef string `xml:"udef,attr"`
	Text string `xml:",chardata"`
}

type QLMID struct {
	IdType    string `xml:"idType,attr"`
	TagType   string `xml:"tagType,attr"`
	StartDate string `xml:"startDate,attr"`
	EndDate   string `xml:"endDate,attr"`
	Udef      string `xml:"udef,attr"`
	Text      string `xml:",chardata"`
}

type MetaData struct {
	InfoItems []InfoItem `xml:"InfoItem"`
}

type Value struct {
	Text     string `xml:",chardata"`
	Type     string `xml:"type,attr"`
	DateTime string `xml:"dateTime,attr"`
	UnixTime int64  `xml:"unixTime,attr"`
}

func Unmarshal(data []byte) (*Objects, error) {
	v := &Objects{}

	if err := xml.Unmarshal(data, v); err != nil {
		return nil, err
	}

	return v, nil
}
