package qlm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func AssertXML(t *testing.T, objects Objects, expected string) bool {
	actual, err := Marshal(objects)
	return assert.Nil(t, err) && assert.Equal(t, expected, (string)(actual))
}

func TestMarshalWithoutObjects(t *testing.T) {
	expected := `<Objects xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:noNamespaceSchemaLocation="QLMdf.xsd"></Objects>`
	objects := Objects{}
	AssertXML(t, objects, expected)
}

func TestMarshalSchemaVersion(t *testing.T) {
	expected := `<Objects xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:noNamespaceSchemaLocation="QLMdf.xsd" version="1.0"></Objects>`
	objects := Objects{Version: "1.0"}
	AssertXML(t, objects, expected)
}

func TestMarshalWithEmptyObject(t *testing.T) {
	expected := `<Objects xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:noNamespaceSchemaLocation="QLMdf.xsd">
    <Object></Object>
</Objects>`
	objects := Objects{
		Objects: []Object{
			Object{},
		},
	}
	AssertXML(t, objects, expected)
}

func TestMarshalWithObjectWithAttributes(t *testing.T) {
	expected := `<Objects xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:noNamespaceSchemaLocation="QLMdf.xsd">
    <Object type="SOME_CLASS_PREFERABLY_DEFINED_BY_UDEF" udef="appropriate.udef.code">
        <id>SmartFridge22334411</id>
    </Object>
</Objects>`
	objects := Objects{
		Objects: []Object{
			Object{
				Type: "SOME_CLASS_PREFERABLY_DEFINED_BY_UDEF",
				Udef: "appropriate.udef.code",
				Id:   &QLMID{Text: "SmartFridge22334411"},
			},
		},
	}
	AssertXML(t, objects, expected)
}

func TestMarshalWithObjectWithComplexId(t *testing.T) {
	expected := `<Objects xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:noNamespaceSchemaLocation="QLMdf.xsd">
    <Object>
        <id idType="some id type" tagType="some tag type" startDate="2013-10-26T21:32:52" endDate="2015-10-26T21:32:52" udef="appropriate.udef.code">SmartFridge22334411</id>
    </Object>
</Objects>`
	objects := Objects{
		Objects: []Object{
			Object{
				Id: &QLMID{
					Text:      "SmartFridge22334411",
					IdType:    "some id type",
					TagType:   "some tag type",
					StartDate: "2013-10-26T21:32:52",
					EndDate:   "2015-10-26T21:32:52",
					Udef:      "appropriate.udef.code",
				},
			},
		},
	}
	AssertXML(t, objects, expected)
}

func TestMarshalWithInfoItemWithOtherNames(t *testing.T) {
	expected := `<Objects xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:noNamespaceSchemaLocation="QLMdf.xsd">
    <Object type="Refrigerator Assembly Product">
        <id>SmartFridge22334411</id>
        <InfoItem udef="b.o.9_1.1.14.13" name="Consumed Electrical Power Measure">
            <name>Some Name 1</name>
            <name>Some Name 2</name>
        </InfoItem>
    </Object>
</Objects>`
	objects := Objects{
		Objects: []Object{
			Object{
				Type: "Refrigerator Assembly Product",
				Id:   &QLMID{Text: "SmartFridge22334411"},
				InfoItems: []InfoItem{
					InfoItem{
						Udef: "b.o.9_1.1.14.13",
						Name: "Consumed Electrical Power Measure",
						OtherNames: []string{
							"Some Name 1",
							"Some Name 2",
						},
					},
				},
			},
		},
	}
	AssertXML(t, objects, expected)
}

func TestMarshalWithInfoItemWithUnixTimestampInValue(t *testing.T) {
	expected := `<Objects xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:noNamespaceSchemaLocation="QLMdf.xsd">
    <Object type="Refrigerator Assembly Product">
        <id>SmartFridge22334411</id>
        <InfoItem udef="b.o.9_1.1.14.13" name="Consumed Electrical Power Measure">
            <value unixTime="1412775405">Value</value>
        </InfoItem>
    </Object>
</Objects>`
	objects := Objects{
		Objects: []Object{
			Object{
				Type: "Refrigerator Assembly Product",
				Id:   &QLMID{Text: "SmartFridge22334411"},
				InfoItems: []InfoItem{
					InfoItem{
						Udef: "b.o.9_1.1.14.13",
						Name: "Consumed Electrical Power Measure",
						Values: []Value{
							Value{
								UnixTime: 1412775405,
								Text:     "Value",
							},
						},
					},
				},
			},
		},
	}
	AssertXML(t, objects, expected)
}

func TestMarshalWithObjectWithComplexDescription(t *testing.T) {
	expected := `<Objects xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:noNamespaceSchemaLocation="QLMdf.xsd">
    <Object type="Refrigerator Assembly Product">
        <id>SmartFridge22334411</id>
        <description lang="en" udef="appropriate.udef.code">Power consumption values with timestamp.</description>
    </Object>
</Objects>`
	objects := Objects{
		Objects: []Object{
			Object{
				Type: "Refrigerator Assembly Product",
				Id:   &QLMID{Text: "SmartFridge22334411"},
				Description: &Description{
					Lang: "en",
					Udef: "appropriate.udef.code",
					Text: "Power consumption values with timestamp.",
				},
			},
		},
	}
	AssertXML(t, objects, expected)
}

func TestMarshalWithInfoItemWithComplexDescription(t *testing.T) {
	expected := `<Objects xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:noNamespaceSchemaLocation="QLMdf.xsd">
    <Object type="Refrigerator Assembly Product">
        <id>SmartFridge22334411</id>
        <InfoItem udef="b.o.9_1.1.14.13" name="Consumed Electrical Power Measure">
            <description lang="en" udef="appropriate.udef.code">Power consumption values with timestamp.</description>
        </InfoItem>
    </Object>
</Objects>`
	objects := Objects{
		Objects: []Object{
			Object{
				Type: "Refrigerator Assembly Product",
				Id:   &QLMID{Text: "SmartFridge22334411"},
				InfoItems: []InfoItem{
					InfoItem{
						Udef: "b.o.9_1.1.14.13",
						Name: "Consumed Electrical Power Measure",
						Description: &Description{
							Lang: "en",
							Udef: "appropriate.udef.code",
							Text: "Power consumption values with timestamp.",
						},
					},
				},
			},
		},
	}
	AssertXML(t, objects, expected)
}

func TestMarshalMeasurementValuesForRefrigeratorPowerConsumption(t *testing.T) {
	expected := `<Objects xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:noNamespaceSchemaLocation="QLMdf.xsd">
    <Object type="Refrigerator Assembly Product">
        <id>SmartFridge22334411</id>
        <InfoItem udef="b.o.9_1.1.14.13" name="Consumed Electrical Power Measure">
            <description>Power consumption values with timestamp.</description>
            <value dateTime="2001-10-26T15:33:21">15.5</value>
            <value dateTime="2001-10-26T15:33:50">15.7</value>
            <value dateTime="2001-10-26T15:34:15">1.3</value>
            <value dateTime="2001-10-26T15:34:35">1.5</value>
            <value dateTime="2001-10-26T15:34:52">15.3</value>
        </InfoItem>
    </Object>
</Objects>`

	objects := Objects{
		Objects: []Object{
			Object{
				Type: "Refrigerator Assembly Product",
				Id:   &QLMID{Text: "SmartFridge22334411"},
				InfoItems: []InfoItem{
					InfoItem{
						Udef:        "b.o.9_1.1.14.13",
						Name:        "Consumed Electrical Power Measure",
						Description: &Description{Text: "Power consumption values with timestamp."},
						Values: []Value{
							Value{
								DateTime: "2001-10-26T15:33:21",
								Text:     "15.5",
							},
							Value{
								DateTime: "2001-10-26T15:33:50",
								Text:     "15.7",
							},
							Value{
								DateTime: "2001-10-26T15:34:15",
								Text:     "1.3",
							},
							Value{
								DateTime: "2001-10-26T15:34:35",
								Text:     "1.5",
							},
							Value{
								DateTime: "2001-10-26T15:34:52",
								Text:     "15.3",
							},
						},
					},
				},
			},
		},
	}

	AssertXML(t, objects, expected)
}

func TestMarshalMetadataAboutRefrigeratorPowerConsumption(t *testing.T) {
	expected := `<Objects xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:noNamespaceSchemaLocation="QLMdf.xsd">
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
</Objects>`

	objects := Objects{
		Objects: []Object{
			Object{
				Id: &QLMID{Text: "SmartFridge22334411"},
				InfoItems: []InfoItem{
					InfoItem{
						Name: "PowerConsumption",
						MetaData: &MetaData{
							InfoItems: []InfoItem{
								InfoItem{
									Name: "format",
									Values: []Value{
										Value{
											Type: "xs:string",
											Text: "xs:double",
										},
									},
								},
								InfoItem{
									Name: "latency",
									Values: []Value{
										Value{
											Type: "xs:int",
											Text: "5",
										},
									},
								},
								InfoItem{
									Name: "readable",
									Values: []Value{
										Value{
											Type: "xs:boolean",
											Text: "true",
										},
									},
								},
								InfoItem{
									Name: "writable",
									Values: []Value{
										Value{
											Type: "xs:boolean",
											Text: "false",
										},
									},
								},
								InfoItem{
									Name: "unit",
									Values: []Value{
										Value{
											Type: "xs:string",
											Text: "Watts",
										},
									},
								},
								InfoItem{
									Name: "accuracy",
									Values: []Value{
										Value{
											Type: "xs:double",
											Text: "1",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	AssertXML(t, objects, expected)
}

func TestMarshalObjectObjectInfoitemValues(t *testing.T) {
	expected := `<Objects xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:noNamespaceSchemaLocation="QLMdf.xsd">
    <Object type="someType">
        <id>UniqueTargetID_1</id>
        <InfoItem name="InfoItem1">
            <value>Value1</value>
            <value>Value2</value>
            <value>Value3</value>
        </InfoItem>
        <InfoItem name="InfoItem2">
            <value>Value</value>
        </InfoItem>
        <Object type="someType">
            <id>SubTarget1</id>
            <InfoItem name="SubInfoItem1"></InfoItem>
            <Object type="someType">
                <id>SubSubTarget1</id>
                <InfoItem name="SubSubTarget1InfoItem1">
                    <value>22.5</value>
                </InfoItem>
            </Object>
        </Object>
        <Object type="someType">
            <id>SubTarget2</id>
            <InfoItem name="SubTarget2InfoItem1">
                <value>34.6</value>
            </InfoItem>
        </Object>
    </Object>
</Objects>`

	objects := Objects{
		Objects: []Object{
			Object{
				Type: "someType",
				Id:   &QLMID{Text: "UniqueTargetID_1"},
				InfoItems: []InfoItem{
					InfoItem{
						Name: "InfoItem1",
						Values: []Value{
							Value{Text: "Value1"},
							Value{Text: "Value2"},
							Value{Text: "Value3"},
						},
					},
					InfoItem{
						Name: "InfoItem2",
						Values: []Value{
							Value{Text: "Value"},
						},
					},
				},
				Objects: []Object{
					Object{
						Type: "someType",
						Id:   &QLMID{Text: "SubTarget1"},
						InfoItems: []InfoItem{
							InfoItem{Name: "SubInfoItem1"},
						},
						Objects: []Object{
							Object{
								Type: "someType",
								Id:   &QLMID{Text: "SubSubTarget1"},
								InfoItems: []InfoItem{
									InfoItem{
										Name: "SubSubTarget1InfoItem1",
										Values: []Value{
											Value{Text: "22.5"},
										},
									},
								},
							},
						},
					},
					Object{
						Type: "someType",
						Id:   &QLMID{Text: "SubTarget2"},
						InfoItems: []InfoItem{
							InfoItem{
								Name: "SubTarget2InfoItem1",
								Values: []Value{
									Value{Text: "34.6"},
								},
							},
						},
					},
				},
			},
		},
	}
	AssertXML(t, objects, expected)
}

func TestMarshalObjectWithSubObjects(t *testing.T) {
	expected := `<Objects xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:noNamespaceSchemaLocation="QLMdf.xsd">
    <Object type="SOME_CLASS_PREFERABLY_DEFINED_BY_UDEF" udef="appropriate.udef.code">
        <id>UniqueObjectID_1</id>
        <InfoItem udef="appropriate.udef.code" name="SOME_CLASS_PREFERABLY_DEFINED_BY_UDEF"></InfoItem>
        <Object type="SOME_CLASS_PREFERABLY_DEFINED_BY_UDEF" udef="appropriate.udef.code">
            <id>SubObject1</id>
        </Object>
        <Object type="SOME_CLASS_PREFERABLY_DEFINED_BY_UDEF" udef="appropriate.udef.code">
            <id>SubObject2</id>
        </Object>
    </Object>
</Objects>`

	objects := Objects{
		Objects: []Object{
			Object{
				Type: "SOME_CLASS_PREFERABLY_DEFINED_BY_UDEF",
				Udef: "appropriate.udef.code",
				Id:   &QLMID{Text: "UniqueObjectID_1"},
				InfoItems: []InfoItem{
					InfoItem{
						Name: "SOME_CLASS_PREFERABLY_DEFINED_BY_UDEF",
						Udef: "appropriate.udef.code",
					},
				},
				Objects: []Object{
					Object{
						Type: "SOME_CLASS_PREFERABLY_DEFINED_BY_UDEF",
						Udef: "appropriate.udef.code",
						Id:   &QLMID{Text: "SubObject1"},
					},
					Object{
						Type: "SOME_CLASS_PREFERABLY_DEFINED_BY_UDEF",
						Udef: "appropriate.udef.code",
						Id:   &QLMID{Text: "SubObject2"},
					},
				},
			},
		},
	}
	AssertXML(t, objects, expected)
}
