package mi

type QlmEnvelope struct {
	Version  string         `xml:"version,attr"`
	Ttl      float64        `xml:"ttl,attr"`
	Response *Response      `xml:"response"`
	Cancel   *CancelRequest `xml:"cancel"`
	Write    *WriteRequest  `xml:"write"`
	Read     *ReadRequest   `xml:"read"`
}

type Response struct {
	Results []RequestResult `xml:"result"`
}

type RequestResult struct {
	Return      *Return   `xml:"return"`
	RequestId   *Id       `xml:"requestId"`
	Message     *Message  `xml:"msg"`
	NodeList    *NodeList `xml:"nodeList"`
	QlmEnvelope *QlmEnvelope
	MsgFormat   string `xml:"msgformat,attr,omitempty"`
	TargetType  string `xml:"targetType,attr,omitempty"`
}

type Return struct {
	ReturnCode  string `xml:"returnCode,attr"`
	Description string `xml:"description,attr,omitempty"`
}

type Id struct {
	Format string `xml:"format,attr,omitempty"`
	Text   string `xml:",chardata"`
}

type NodeList struct {
	Nodes []string `xml:"node"`
	Type  string   `xml:"type,attr,omitempty"`
}

type CancelRequest struct {
	RequestIds []Id      `xml:"requestId"`
	NodeList   *NodeList `xml:"nodeList"`
}

type ReadRequest struct {
	NodeList   *NodeList `xml:"nodeList"`
	RequestIds []Id      `xml:"requestId"`
	Message    *Message  `xml:"msg"`
	MsgFormat  string    `xml:"msgformat,attr,omitempty"`
	Callback   string    `xml:"callback,attr,omitempty"`
	TargetType string    `xml:"targetType,attr,omitempty"`
	Interval   float64   `xml:"interval,attr,omitempty"`
	Oldest     int       `xml:"oldest,attr,omitempty"`
	Newest     int       `xml:"newest,attr,omitempty"`
	Begin      string    `xml:"begin,attr,omitempty"`
	End        string    `xml:"end,attr,omitempty"`
}

type WriteRequest struct {
	NodeList   *NodeList `xml:"nodeList"`
	RequestIds []Id      `xml:"requestId"`
	Message    *Message  `xml:"msg"`
	Callback   string    `xml:"callback,attr,omitempty"`
	MsgFormat  string    `xml:"msgformat,attr,omitempty"`
	TargetType string    `xml:"targetType,attr,omitempty"`
}

type Message struct {
	Data string `xml:",innerxml"`
}
