package mi

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestUnmarshalWithInvalidXML(t *testing.T) {
	data := `invalid`
	v, err := Unmarshal([]byte(data))
	assert.NotNil(t, err)
	assert.Nil(t, v)
}

func TestUnmarshalCancelRequest(t *testing.T) {
	data, err := ioutil.ReadFile("examples/cancel_request.xml")
	if assert.Nil(t, err) {
		v, err := Unmarshal(data)
		if assert.Nil(t, err) {
			assert.Equal(t, "1.0", v.Version)
			assert.Equal(t, 10, v.Ttl)
			if assert.Len(t, v.Cancel.RequestIds, 2) {
				assert.Equal(t, "REQ0011212121212", v.Cancel.RequestIds[0].Text)
				assert.Equal(t, "REQ0011212121213", v.Cancel.RequestIds[1].Text)
			}
		}
	}
}

func TestUnmarshalCancelRequestWithNodes(t *testing.T) {
	data, err := ioutil.ReadFile("examples/cancel_request_with_nodes.xml")
	if assert.Nil(t, err) {
		v, err := Unmarshal(data)
		if assert.Nil(t, err) {
			assert.Equal(t, "1.0", v.Version)
			assert.Equal(t, 10, v.Ttl)
			if assert.Len(t, v.Cancel.RequestIds, 1) {
				assert.Equal(t, "REQ0011212121212", v.Cancel.RequestIds[0].Text)
			}
			assert.Equal(t, "URL", v.Cancel.NodeList.Type)
			if assert.Len(t, v.Cancel.NodeList.Nodes, 2) {
				assert.Equal(t, "http://192.168.0.1/", v.Cancel.NodeList.Nodes[0])
				assert.Equal(t, "http://192.168.0.2/", v.Cancel.NodeList.Nodes[1])
			}
		}
	}
}

func TestUnmarshalErrorResponse(t *testing.T) {
	data, err := ioutil.ReadFile("examples/error_response.xml")
	if assert.Nil(t, err) {
		v, err := Unmarshal(data)
		if assert.Nil(t, err) {
			assert.Equal(t, "1.0", v.Version)
			assert.Equal(t, 0, v.Ttl)
			if assert.Len(t, v.Response.Results, 1) {
				assert.Equal(t, "404", v.Response.Results[0].Return.ReturnCode)
			}
		}
	}
}

func TestUnmarshalErrorResponseWithDescription(t *testing.T) {
	data, err := ioutil.ReadFile("examples/error_response_with_description.xml")
	if assert.Nil(t, err) {
		v, err := Unmarshal(data)
		if assert.Nil(t, err) {
			if assert.Len(t, v.Response.Results, 1) {
				assert.Equal(t, "Not Found", v.Response.Results[0].Return.Description)
			}
		}
	}
}

func TestUnmarshalMultiplePayloadResponse(t *testing.T) {
	data, err := ioutil.ReadFile("examples/multiple_payload_response.xml")
	if assert.Nil(t, err) {
		v, err := Unmarshal(data)
		if assert.Nil(t, err) {
			assert.Equal(t, "1.0", v.Version)
			assert.Equal(t, 10, v.Ttl)
			if assert.Len(t, v.Response.Results, 3) {
				assert.Equal(t, "obix", v.Response.Results[0].MsgFormat)
				assert.Equal(t, "200", v.Response.Results[0].Return.ReturnCode)
				assert.Equal(t, "REQ0011212121212", v.Response.Results[0].RequestId.Text)
				assert.Equal(t, `
                <obj href="http://myhome/thermostat" >
                    <real name="spaceTemp" unit="obix:units/fahrenheit" val="67.2"/>
                    <real name="setpoint" unit="obix:units/fahrenheit" val="72.0"/>
                    <bool name="furnaceOn" val="true"/>
                </obj>
            `, v.Response.Results[0].Message.Data)

				assert.Equal(t, "CSV", v.Response.Results[1].MsgFormat)
				assert.Equal(t, "200", v.Response.Results[1].Return.ReturnCode)
				assert.Equal(t, "REQ232323", v.Response.Results[1].RequestId.Text)
				assert.Equal(t, `11,22,33
                44,55,66`, v.Response.Results[1].Message.Data)

				assert.Equal(t, "QLMdf", v.Response.Results[2].MsgFormat)
				assert.Equal(t, "200", v.Response.Results[2].Return.ReturnCode)
				assert.Equal(t, "REQ654534", v.Response.Results[2].RequestId.Text)
				assert.Equal(t, `
                <Objects>
                    <Object>
                        <id>SmartFridge22334411</id>
                        <InfoItem name="PowerConsumption">
                            <value>43</value>
                        </InfoItem>
                    </Object>
                </Objects>
            `, v.Response.Results[2].Message.Data)
			}
		}
	}
}

func TestUnmarshalPublishing(t *testing.T) {
	data, err := ioutil.ReadFile("examples/publishing.xml")
	if assert.Nil(t, err) {
		v, err := Unmarshal(data)
		if assert.Nil(t, err) {
			assert.Equal(t, "1.0", v.Version)
			assert.Equal(t, -1, v.Ttl)
			assert.Equal(t, "QLMdf", v.Write.MsgFormat)
			assert.Equal(t, `
            <Objects>
                <Object>
                    <id>SmartFridge22334411</id>
                    <InfoItem name="FridgeTemperatureSetpoint">
                        <value>3.5</value>
                    </InfoItem>
                </Object>
            </Objects>
        `, v.Write.Message.Data)
		}
	}
}

func TestUnmarshalReadRequest(t *testing.T) {
	data, err := ioutil.ReadFile("examples/read_request.xml")
	if assert.Nil(t, err) {
		v, err := Unmarshal(data)
		if assert.Nil(t, err) {
			assert.Equal(t, "1.0", v.Version)
			assert.Equal(t, 10, v.Ttl)
			assert.Equal(t, "QLM_mf.xsd", v.Read.MsgFormat)
			assert.Equal(t, 3.5, v.Read.Interval)
			assert.Equal(t, 10, v.Read.Oldest)
			assert.Equal(t, "2014-01-01T00:00", v.Read.Begin)
			assert.Equal(t, "2014-02-01T00:00", v.Read.End)
			assert.Equal(t, 15, v.Read.Newest)
			assert.Equal(t, `
            <Objects>
                <Object>
                    <id>SmartFridge22334411</id>
                    <InfoItem name="PowerConsumption"></InfoItem>
                </Object>
            </Objects>
        `, v.Read.Message.Data)
		}
	}
}

func TestUnmarshalReadRequestWithCallback(t *testing.T) {
	data, err := ioutil.ReadFile("examples/read_request_with_callback.xml")
	if assert.Nil(t, err) {
		v, err := Unmarshal(data)
		if assert.Nil(t, err) {
			assert.Equal(t, "http://192.168.0.1/", v.Read.Callback)
		}
	}
}

func TestUnmarshalReadRequestWithNodes(t *testing.T) {
	data, err := ioutil.ReadFile("examples/read_request_with_nodes.xml")
	if assert.Nil(t, err) {
		v, err := Unmarshal(data)
		if assert.Nil(t, err) {
			assert.Equal(t, "URL", v.Read.NodeList.Type)
			if assert.Len(t, v.Read.NodeList.Nodes, 2) {
				assert.Equal(t, "http://192.168.0.1/", v.Read.NodeList.Nodes[0])
				assert.Equal(t, "http://192.168.0.2/", v.Read.NodeList.Nodes[1])
			}
		}
	}
}

func TestUnmarshalReadResponseMetadata(t *testing.T) {
	data, err := ioutil.ReadFile("examples/read_response_metadata.xml")
	if assert.Nil(t, err) {
		v, err := Unmarshal(data)
		if assert.Nil(t, err) {
			assert.Equal(t, "1.0", v.Version)
			assert.Equal(t, 10, v.Ttl)
			if assert.Len(t, v.Response.Results, 1) {
				assert.Equal(t, "QLMdf", v.Response.Results[0].MsgFormat)
				assert.Equal(t, "200", v.Response.Results[0].Return.ReturnCode)
				assert.Equal(t, "REQ654534", v.Response.Results[0].RequestId.Text)
				assert.Equal(t, `
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
            `, v.Response.Results[0].Message.Data)
			}
		}
	}
}

func TestUnmarshalReadResponseWithRequestIdFormat(t *testing.T) {
	data, err := ioutil.ReadFile("examples/read_response_with_request_id_format.xml")
	if assert.Nil(t, err) {
		v, err := Unmarshal(data)
		if assert.Nil(t, err) {
			if assert.Len(t, v.Response.Results, 1) {
				assert.Equal(t, "REQ", v.Response.Results[0].RequestId.Format)
			}
		}
	}
}
func TestUnmarshalResponseWithNodes(t *testing.T) {
	data, err := ioutil.ReadFile("examples/response_with_nodes.xml")
	if assert.Nil(t, err) {
		v, err := Unmarshal(data)
		if assert.Nil(t, err) && assert.Len(t, v.Response.Results, 1) {
			assert.Equal(t, "URL", v.Response.Results[0].NodeList.Type)
			if assert.Len(t, v.Response.Results[0].NodeList.Nodes, 2) {
				assert.Equal(t, "http://192.168.0.1/", v.Response.Results[0].NodeList.Nodes[0])
				assert.Equal(t, "http://192.168.0.2/", v.Response.Results[0].NodeList.Nodes[1])
			}
		}
	}
}

func TestUnmarshalTypicalMinimalResponse(t *testing.T) {
	data, err := ioutil.ReadFile("examples/typical_minimal_response.xml")
	if assert.Nil(t, err) {
		v, err := Unmarshal(data)
		if assert.Nil(t, err) {
			assert.Equal(t, "0.2", v.Version)
			assert.Equal(t, 0, v.Ttl)
			if assert.Len(t, v.Response.Results, 1) {
				assert.Equal(t, "200", v.Response.Results[0].Return.ReturnCode)
			}
		}
	}
}

func TestUnmarshalWriteRequest(t *testing.T) {
	data, err := ioutil.ReadFile("examples/write_request.xml")
	if assert.Nil(t, err) {
		v, err := Unmarshal(data)
		if assert.Nil(t, err) {
			assert.Equal(t, "1.0", v.Version)
			assert.Equal(t, -1, v.Ttl)
			assert.Equal(t, "QLMdf", v.Write.MsgFormat)
			assert.Equal(t, "device", v.Write.TargetType)
			assert.Equal(t, `
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
        `, v.Write.Message.Data)
		}
	}
}
