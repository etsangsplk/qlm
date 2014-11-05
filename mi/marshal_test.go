package mi

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func assertXML(t *testing.T, envelope QlmEnvelope, expected string) bool {
	actual, err := Marshal(envelope)
	return assert.Nil(t, err) && assert.Equal(t, expected, (string)(actual))
}

func TestMarshalCancelRequest(t *testing.T) {
	expected := `<qlmEnvelope version="1.0" ttl="10">
    <cancel>
        <requestId>REQ0011212121212</requestId>
        <requestId>REQ0011212121213</requestId>
    </cancel>
</qlmEnvelope>`
	envelope := QlmEnvelope{
		Version: "1.0",
		Ttl:     10,
		Cancel: &CancelRequest{
			RequestIds: []Id{
				Id{Text: "REQ0011212121212"},
				Id{Text: "REQ0011212121213"},
			},
		},
	}
	assertXML(t, envelope, expected)
}

func TestMarshalCancelRequestWithNodes(t *testing.T) {
	expected := `<qlmEnvelope version="1.0" ttl="10">
    <cancel>
        <requestId>REQ0011212121212</requestId>
        <nodeList type="URL">
            <node>http://192.168.0.1/</node>
            <node>http://192.168.0.2/</node>
        </nodeList>
    </cancel>
</qlmEnvelope>`
	envelope := QlmEnvelope{
		Version: "1.0",
		Ttl:     10,
		Cancel: &CancelRequest{
			RequestIds: []Id{
				Id{Text: "REQ0011212121212"},
			},
			NodeList: &NodeList{
				Type: "URL",
				Nodes: []string{
					"http://192.168.0.1/",
					"http://192.168.0.2/",
				},
			},
		},
	}
	assertXML(t, envelope, expected)
}

func TestMarshalErrorResponse(t *testing.T) {
	expected := `<qlmEnvelope version="1.0" ttl="0">
    <response>
        <result>
            <return returnCode="404"></return>
        </result>
    </response>
</qlmEnvelope>`
	envelope := QlmEnvelope{
		Version: "1.0",
		Ttl:     0,
		Response: &Response{
			Results: []RequestResult{
				RequestResult{
					Return: &Return{ReturnCode: "404"},
				},
			},
		},
	}
	assertXML(t, envelope, expected)
}

func TestMarshalErrorResponseWithDescription(t *testing.T) {
	expected := `<qlmEnvelope version="1.0" ttl="0">
    <response>
        <result>
            <return returnCode="404" description="Not Found"></return>
        </result>
    </response>
</qlmEnvelope>`
	envelope := QlmEnvelope{
		Version: "1.0",
		Ttl:     0,
		Response: &Response{
			Results: []RequestResult{
				RequestResult{
					Return: &Return{
						ReturnCode:  "404",
						Description: "Not Found",
					},
				},
			},
		},
	}
	assertXML(t, envelope, expected)
}

func TestMarshalMultiplePayloadResponse(t *testing.T) {
	expected := `<qlmEnvelope version="1.0" ttl="10">
    <response>
        <result msgformat="obix">
            <return returnCode="200"></return>
            <requestId>REQ0011212121212</requestId>
            <msg>
                <obj href="http://myhome/thermostat" >
                    <real name="spaceTemp" unit="obix:units/fahrenheit" val="67.2"/>
                    <real name="setpoint" unit="obix:units/fahrenheit" val="72.0"/>
                    <bool name="furnaceOn" val="true"/>
                </obj>
            </msg>
        </result>
        <result msgformat="CSV">
            <return returnCode="200"></return>
            <requestId>REQ232323</requestId>
            <msg>11,22,33
                44,55,66</msg>
        </result>
        <result msgformat="QLMdf">
            <return returnCode="200"></return>
            <requestId>REQ654534</requestId>
            <msg>
                <Objects>
                    <Object>
                        <id>SmartFridge22334411</id>
                        <InfoItem name="PowerConsumption">
                            <value>43</value>
                        </InfoItem>
                    </Object>
                </Objects>
            </msg>
        </result>
    </response>
</qlmEnvelope>`
	envelope := QlmEnvelope{
		Version: "1.0",
		Ttl:     10,
		Response: &Response{
			Results: []RequestResult{
				RequestResult{
					MsgFormat: "obix",
					Return:    &Return{ReturnCode: "200"},
					RequestId: &Id{Text: "REQ0011212121212"},
					Message: &Message{
						Data: `
                <obj href="http://myhome/thermostat" >
                    <real name="spaceTemp" unit="obix:units/fahrenheit" val="67.2"/>
                    <real name="setpoint" unit="obix:units/fahrenheit" val="72.0"/>
                    <bool name="furnaceOn" val="true"/>
                </obj>
            `,
					},
				},
				RequestResult{
					MsgFormat: "CSV",
					Return:    &Return{ReturnCode: "200"},
					RequestId: &Id{Text: "REQ232323"},
					Message: &Message{
						Data: `11,22,33
                44,55,66`,
					},
				},
				RequestResult{
					MsgFormat: "QLMdf",
					Return:    &Return{ReturnCode: "200"},
					RequestId: &Id{Text: "REQ654534"},
					Message: &Message{
						Data: `
                <Objects>
                    <Object>
                        <id>SmartFridge22334411</id>
                        <InfoItem name="PowerConsumption">
                            <value>43</value>
                        </InfoItem>
                    </Object>
                </Objects>
            `,
					},
				},
			},
		},
	}
	assertXML(t, envelope, expected)
}

func TestMarshalPublishing(t *testing.T) {
	expected := `<qlmEnvelope version="1.0" ttl="-1">
    <write msgformat="QLMdf">
        <msg>
            <Objects>
                <Object>
                    <id>SmartFridge22334411</id>
                    <InfoItem name="FridgeTemperatureSetpoint">
                        <value>3.5</value>
                    </InfoItem>
                </Object>
            </Objects>
        </msg>
    </write>
</qlmEnvelope>`
	envelope := QlmEnvelope{
		Version: "1.0",
		Ttl:     -1,
		Write: &WriteRequest{
			MsgFormat: "QLMdf",
			Message: &Message{
				Data: `
            <Objects>
                <Object>
                    <id>SmartFridge22334411</id>
                    <InfoItem name="FridgeTemperatureSetpoint">
                        <value>3.5</value>
                    </InfoItem>
                </Object>
            </Objects>
        `},
		},
	}
	assertXML(t, envelope, expected)
}

func TestMarshalReadRequest(t *testing.T) {
	expected := `<qlmEnvelope version="1.0" ttl="10">
    <read msgformat="QLM_mf.xsd" interval="3.5" oldest="10" newest="15" begin="2014-01-01T00:00" end="2014-02-01T00:00">
        <msg>
            <Objects>
                <Object>
                    <id>SmartFridge22334411</id>
                    <InfoItem name="PowerConsumption"></InfoItem>
                </Object>
            </Objects>
        </msg>
    </read>
</qlmEnvelope>`
	envelope := QlmEnvelope{
		Version: "1.0",
		Ttl:     10,
		Read: &ReadRequest{
			MsgFormat: "QLM_mf.xsd",
			Interval:  3.5,
			Oldest:    10,
			Newest:    15,
			Begin:     "2014-01-01T00:00",
			End:       "2014-02-01T00:00",
			Message: &Message{
				Data: `
            <Objects>
                <Object>
                    <id>SmartFridge22334411</id>
                    <InfoItem name="PowerConsumption"></InfoItem>
                </Object>
            </Objects>
        `},
		},
	}
	assertXML(t, envelope, expected)
}

func TestMarshalReadRequestWithCallback(t *testing.T) {
	expected := `<qlmEnvelope version="1.0" ttl="10">
    <read msgformat="QLM_mf.xsd" callback="http://192.168.0.1/">
        <msg>
            <Objects>
                <Object>
                    <id>SmartFridge22334411</id>
                    <InfoItem name="PowerConsumption"></InfoItem>
                </Object>
            </Objects>
        </msg>
    </read>
</qlmEnvelope>`
	envelope := QlmEnvelope{
		Version: "1.0",
		Ttl:     10,
		Read: &ReadRequest{
			MsgFormat: "QLM_mf.xsd",
			Callback:  "http://192.168.0.1/",
			Message: &Message{
				Data: `
            <Objects>
                <Object>
                    <id>SmartFridge22334411</id>
                    <InfoItem name="PowerConsumption"></InfoItem>
                </Object>
            </Objects>
        `},
		},
	}
	assertXML(t, envelope, expected)
}

func TestMarshalReadRequestWithNodes(t *testing.T) {
	expected := `<qlmEnvelope version="1.0" ttl="10">
    <read msgformat="QLM_mf.xsd">
        <nodeList type="URL">
            <node>http://192.168.0.1/</node>
            <node>http://192.168.0.2/</node>
        </nodeList>
        <msg>
            <Objects>
                <Object>
                    <id>SmartFridge22334411</id>
                    <InfoItem name="PowerConsumption"></InfoItem>
                </Object>
            </Objects>
        </msg>
    </read>
</qlmEnvelope>`
	envelope := QlmEnvelope{
		Version: "1.0",
		Ttl:     10,
		Read: &ReadRequest{
			MsgFormat: "QLM_mf.xsd",
			NodeList: &NodeList{
				Type: "URL",
				Nodes: []string{
					"http://192.168.0.1/",
					"http://192.168.0.2/",
				},
			},
			Message: &Message{
				Data: `
            <Objects>
                <Object>
                    <id>SmartFridge22334411</id>
                    <InfoItem name="PowerConsumption"></InfoItem>
                </Object>
            </Objects>
        `},
		},
	}
	assertXML(t, envelope, expected)
}

func TestMarshalReadResponseMetadata(t *testing.T) {
	expected := `<qlmEnvelope version="1.0" ttl="10">
    <response>
        <result msgformat="QLMdf">
            <return returnCode="200"></return>
            <requestId>REQ654534</requestId>
            <msg>
                <Objects>
                    <Object>
                        <id>SmartFridge22334411</id>
                        <InfoItem name="PowerConsumption">
                            <MetaData>
                                <InfoItem name="format">
                                    <value type="xs:string">xs:double</value>
                                </InfoItem>
                                <InfoItem name="latency">
                                    <value type="xs:int">5</value>
                                </InfoItem>
                                <InfoItem name="readable">
                                    <value type="xs:boolean">true</value>
                                </InfoItem>
                                <InfoItem name="writable">
                                    <value type="xs:boolean">false</value>
                                </InfoItem>
                                <InfoItem name="unit">
                                    <value type="xs:string">Watts</value>
                                </InfoItem>
                                <InfoItem name="accuracy">
                                    <value type="xs:double">1</value>
                                </InfoItem>
                            </MetaData>
                        </InfoItem>
                    </Object>
                </Objects>
            </msg>
        </result>
    </response>
</qlmEnvelope>`
	envelope := QlmEnvelope{
		Version: "1.0",
		Ttl:     10,
		Response: &Response{
			Results: []RequestResult{
				RequestResult{
					MsgFormat: "QLMdf",
					Return:    &Return{ReturnCode: "200"},
					RequestId: &Id{Text: "REQ654534"},
					Message: &Message{
						Data: `
                <Objects>
                    <Object>
                        <id>SmartFridge22334411</id>
                        <InfoItem name="PowerConsumption">
                            <MetaData>
                                <InfoItem name="format">
                                    <value type="xs:string">xs:double</value>
                                </InfoItem>
                                <InfoItem name="latency">
                                    <value type="xs:int">5</value>
                                </InfoItem>
                                <InfoItem name="readable">
                                    <value type="xs:boolean">true</value>
                                </InfoItem>
                                <InfoItem name="writable">
                                    <value type="xs:boolean">false</value>
                                </InfoItem>
                                <InfoItem name="unit">
                                    <value type="xs:string">Watts</value>
                                </InfoItem>
                                <InfoItem name="accuracy">
                                    <value type="xs:double">1</value>
                                </InfoItem>
                            </MetaData>
                        </InfoItem>
                    </Object>
                </Objects>
            `,
					},
				},
			},
		},
	}
	assertXML(t, envelope, expected)
}

func TestMarshalReadResponseWithRequestIdFormat(t *testing.T) {
	expected := `<qlmEnvelope version="1.0" ttl="10">
    <response>
        <result msgformat="QLMdf">
            <return returnCode="200"></return>
            <requestId format="REQ">REQ654534</requestId>
            <msg>
                <Objects>
                    <Object>
                        <id>SmartFridge22334411</id>
                        <!-- Only most recent value, bogus timestamp. -->
                        <InfoItem name="PowerConsumption">
                            <value type="xs:int" unixTime="5453563">43</value>
                        </InfoItem>
                    </Object>
                </Objects>
            </msg>
        </result>
    </response>
</qlmEnvelope>`
	envelope := QlmEnvelope{
		Version: "1.0",
		Ttl:     10,
		Response: &Response{
			Results: []RequestResult{
				RequestResult{
					MsgFormat: "QLMdf",
					Return:    &Return{ReturnCode: "200"},
					RequestId: &Id{
						Format: "REQ",
						Text:   "REQ654534",
					},
					Message: &Message{
						Data: `
                <Objects>
                    <Object>
                        <id>SmartFridge22334411</id>
                        <!-- Only most recent value, bogus timestamp. -->
                        <InfoItem name="PowerConsumption">
                            <value type="xs:int" unixTime="5453563">43</value>
                        </InfoItem>
                    </Object>
                </Objects>
            `,
					},
				},
			},
		},
	}
	assertXML(t, envelope, expected)
}

func TestMarshalResponseWithNodes(t *testing.T) {
	expected := `<qlmEnvelope version="1.0" ttl="10">
    <response>
        <result>
            <return returnCode="200"></return>
            <nodeList type="URL">
                <node>http://192.168.0.1/</node>
                <node>http://192.168.0.2/</node>
            </nodeList>
        </result>
    </response>
</qlmEnvelope>`
	envelope := QlmEnvelope{
		Version: "1.0",
		Ttl:     10,
		Response: &Response{
			Results: []RequestResult{
				RequestResult{
					Return: &Return{ReturnCode: "200"},
					NodeList: &NodeList{
						Type: "URL",
						Nodes: []string{
							"http://192.168.0.1/",
							"http://192.168.0.2/",
						},
					},
				},
			},
		},
	}
	assertXML(t, envelope, expected)
}

func TestMarshalTypicalMinimalResponse(t *testing.T) {
	expected := `<qlmEnvelope version="0.2" ttl="0">
    <response>
        <result>
            <return returnCode="200"></return>
        </result>
    </response>
</qlmEnvelope>`
	envelope := QlmEnvelope{
		Version: "0.2",
		Ttl:     0,
		Response: &Response{
			Results: []RequestResult{
				RequestResult{
					Return: &Return{ReturnCode: "200"},
				},
			},
		},
	}
	assertXML(t, envelope, expected)
}

func TestMarshalWriteRequest(t *testing.T) {
	expected := `<qlmEnvelope version="1.0" ttl="-1">
    <write msgformat="QLMdf" targetType="device">
        <msg>
            <Objects>
                <Object>
                    <id>SmartFridge22334411</id>
                    <InfoItem name="FridgeTemperatureSetpoint">
                        <value>3.5</value>
                    </InfoItem>
                    <InfoItem name="FreezerTemperatureSetpoint">
                        <value>-20.0</value>
                    </InfoItem>
                </Object>
            </Objects>
        </msg>
    </write>
</qlmEnvelope>`
	envelope := QlmEnvelope{
		Version: "1.0",
		Ttl:     -1,
		Write: &WriteRequest{
			MsgFormat:  "QLMdf",
			TargetType: "device",
			Message: &Message{
				Data: `
            <Objects>
                <Object>
                    <id>SmartFridge22334411</id>
                    <InfoItem name="FridgeTemperatureSetpoint">
                        <value>3.5</value>
                    </InfoItem>
                    <InfoItem name="FreezerTemperatureSetpoint">
                        <value>-20.0</value>
                    </InfoItem>
                </Object>
            </Objects>
        `,
			},
		},
	}
	assertXML(t, envelope, expected)
}
