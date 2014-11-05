package qlm

type QlmEnvelope struct {
	Version   string          `xml:"version,attr"`
	Ttl       float64         `xml:"ttl,attr"`
	Responses []RequestResult `xml:"response>result"`
	Cancel    CancelRequest   `xml:"cancel"`
	Write     WriteRequest    `xml:"write"`
	Read      ReadRequest     `xml:"read"`
}

type RequestResult struct {
	Return      Return   `xml:"return"`
	RequestId   Id       `xml:"requestId"`
	Message     Message  `xml:"msg"`
	NodeList    NodeList `xml:"nodeList"`
	QlmEnvelope QlmEnvelope
	MsgFormat   string `xml:"msgformat,attr"`
	TargetType  string
}

type Return struct {
	ReturnCode  string `xml:"returnCode,attr"`
	Description string `xml:"description,attr"`
}

type Id struct {
	Format string `xml:"format,attr"`
	Text   string `xml:",chardata"`
}

type NodeList struct {
	Nodes []string `xml:"node"`
	Type  string   `xml:"type,attr"`
}

type CancelRequest struct {
	NodeList   NodeList `xml:"nodeList"`
	RequestIds []Id     `xml:"requestId"`
}

type ReadRequest struct {
	NodeList   NodeList `xml:"nodeList"`
	RequestIds []Id     `xml:"requestId"`
	Message    Message  `xml:"msg"`
	Callback   string   `xml:"callback,attr"`
	MsgFormat  string   `xml:"msgformat,attr"`
	TargetType string   `xml:"targetType,attr"`
	Oldest     int      `xml:"oldest,attr"`
	Newest     int      `xml:"newest,attr"`
	Interval   float64  `xml:"interval,attr"`
	Begin      string   `xml:"begin,attr"`
	End        string   `xml:"end,attr"`
}

type WriteRequest struct {
	NodeList   NodeList `xml:"nodeList"`
	RequestIds []Id     `xml:"requestId"`
	Message    Message  `xml:"msg"`
	Callback   string   `xml:"callback,attr"`
	MsgFormat  string   `xml:"msgformat,attr"`
	TargetType string   `xml:"targetType,attr"`
}

type Message struct {
	Data string `xml:",innerxml"`
}
