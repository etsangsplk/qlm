package df

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

func TestUnmarshalWithoutObjects(t *testing.T) {
	data := `
		<?xml version="1.0" encoding="UTF-8"?>
		<Objects xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:noNamespaceSchemaLocation="odf.xsd">
		</Objects>
		`
	v, err := Unmarshal([]byte(data))
	if assert.Nil(t, err) {
		assert.Len(t, v.Objects, 0)
	}
}

func TestUnmarshalSchemaVersion(t *testing.T) {
	data := `
		<?xml version="1.0" encoding="UTF-8"?>
		<Objects xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:noNamespaceSchemaLocation="odf.xsd" version="1.0">
		</Objects>
		`
	v, err := Unmarshal([]byte(data))
	if assert.Nil(t, err) {
		assert.Equal(t, v.Version, "1.0")
	}
}

func TestUnmarshalWithObjectWithAttributes(t *testing.T) {
	data := `
		<?xml version="1.0" encoding="UTF-8"?>
		<Objects xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:noNamespaceSchemaLocation="odf.xsd">
			<Object type="SOME_CLASS_PREFERABLY_DEFINED_BY_UDEF" udef="appropriate.udef.code">
				<id>SmartFridge22334411</id>
			</Object>
		</Objects>
	`
	v, err := Unmarshal([]byte(data))
	if assert.Nil(t, err) {
		if assert.Len(t, v.Objects, 1) {
			assert.Equal(t, v.Objects[0].Type, "SOME_CLASS_PREFERABLY_DEFINED_BY_UDEF")
			assert.Equal(t, v.Objects[0].Udef, "appropriate.udef.code")
			assert.Equal(t, v.Objects[0].Id.Text, "SmartFridge22334411")
		}
	}
}

func TestUnmarshalWithObjectWithComplexId(t *testing.T) {
	data := `
		<?xml version="1.0" encoding="UTF-8"?>
		<Objects xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:noNamespaceSchemaLocation="odf.xsd">
			<Object>
				<id idType="some id type" tagType="some tag type" startDate="2013-10-26T21:32:52" endDate="2015-10-26T21:32:52" udef="appropriate.udef.code">SmartFridge22334411</id>
			</Object>
		</Objects>
 	`
	v, err := Unmarshal([]byte(data))
	if assert.Nil(t, err) {
		if assert.Len(t, v.Objects, 1) {
			assert.Equal(t, v.Objects[0].Id.IdType, "some id type")
			assert.Equal(t, v.Objects[0].Id.TagType, "some tag type")
			assert.Equal(t, v.Objects[0].Id.StartDate, "2013-10-26T21:32:52")
			assert.Equal(t, v.Objects[0].Id.EndDate, "2015-10-26T21:32:52")
			assert.Equal(t, v.Objects[0].Id.Udef, "appropriate.udef.code")
			assert.Equal(t, v.Objects[0].Id.Text, "SmartFridge22334411")
		}
	}
}

func TestUnmarshalWithInfoItemWithOtherNames(t *testing.T) {
	data := `
		<?xml version="1.0" encoding="UTF-8"?>
		<Objects xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:noNamespaceSchemaLocation="odf.xsd">
			<Object type="Refrigerator Assembly Product">
				<id>SmartFridge22334411</id>
				<InfoItem udef="b.o.9_1.1.14.13" name="Consumed Electrical Power Measure">
					<name>Some Name 1</name>
					<name>Some Name 2</name>
				</InfoItem>
			</Object>
		</Objects>
	`
	v, err := Unmarshal([]byte(data))
	if assert.Nil(t, err) && assert.Len(t, v.Objects, 1) && assert.Len(t, v.Objects[0].InfoItems, 1) && assert.Len(t, v.Objects[0].InfoItems[0].OtherNames, 2) {
		assert.Equal(t, v.Objects[0].InfoItems[0].OtherNames[0], "Some Name 1")
		assert.Equal(t, v.Objects[0].InfoItems[0].OtherNames[1], "Some Name 2")
	}
}

func TestUnmarshalWithInfoItemWithUnixTimestampInValue(t *testing.T) {
	data := `
		<?xml version="1.0" encoding="UTF-8"?>
		<Objects xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:noNamespaceSchemaLocation="odf.xsd">
			<Object type="Refrigerator Assembly Product">
				<id>SmartFridge22334411</id>
				<InfoItem udef="b.o.9_1.1.14.13" name="Consumed Electrical Power Measure">
					<value unixTime="1412775405">Value</value>
				</InfoItem>
			</Object>
		</Objects>
	`
	v, err := Unmarshal([]byte(data))
	if assert.Nil(t, err) && assert.Len(t, v.Objects, 1) && assert.Len(t, v.Objects[0].InfoItems, 1) && assert.Len(t, v.Objects[0].InfoItems[0].Values, 1) {
		assert.Equal(t, v.Objects[0].InfoItems[0].Values[0].UnixTime, 1412775405)
	}
}

func TestUnmarshalWithObjectWithComplexDescription(t *testing.T) {
	data := `
		<?xml version="1.0" encoding="UTF-8"?>
		<Objects xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:noNamespaceSchemaLocation="odf.xsd">
			<Object type="Refrigerator Assembly Product">
				<id>SmartFridge22334411</id>
				<description lang="en" udef="appropriate.udef.code">Power consumption values with timestamp.</description>
			</Object>
		</Objects>
	`
	v, err := Unmarshal([]byte(data))
	if assert.Nil(t, err) && assert.Len(t, v.Objects, 1) {
		assert.Equal(t, v.Objects[0].Description.Lang, "en")
		assert.Equal(t, v.Objects[0].Description.Udef, "appropriate.udef.code")
		assert.Equal(t, v.Objects[0].Description.Text, "Power consumption values with timestamp.")
	}
}

func TestUnmarshalWithInfoItemWithComplexDescription(t *testing.T) {
	data := `
		<?xml version="1.0" encoding="UTF-8"?>
		<Objects xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:noNamespaceSchemaLocation="odf.xsd">
			<Object type="Refrigerator Assembly Product">
				<id>SmartFridge22334411</id>
				<InfoItem udef="b.o.9_1.1.14.13" name="Consumed Electrical Power Measure">
					<description lang="en" udef="appropriate.udef.code">Power consumption values with timestamp.</description>
				</InfoItem>
			</Object>
		</Objects>
	`
	v, err := Unmarshal([]byte(data))
	if assert.Nil(t, err) && assert.Len(t, v.Objects, 1) && assert.Len(t, v.Objects[0].InfoItems, 1) {
		assert.Equal(t, v.Objects[0].InfoItems[0].Description.Lang, "en")
		assert.Equal(t, v.Objects[0].InfoItems[0].Description.Udef, "appropriate.udef.code")
		assert.Equal(t, v.Objects[0].InfoItems[0].Description.Text, "Power consumption values with timestamp.")
	}
}
func TestUnmarshalMeasurementValuesForRefrigeratorPowerConsumption(t *testing.T) {
	data, err := ioutil.ReadFile("examples/measurement_values_for_refrigerator_power_consumption.xml")
	if assert.Nil(t, err) {
		v, err := Unmarshal(data)
		if assert.Nil(t, err) {
			if assert.Len(t, v.Objects, 1) {
				assert.Equal(t, v.Objects[0].Type, "Refrigerator Assembly Product")
				assert.Equal(t, v.Objects[0].Id.Text, "SmartFridge22334411")

				if infoItems := v.Objects[0].InfoItems; assert.Len(t, infoItems, 1) {
					assert.Equal(t, infoItems[0].Udef, "b.o.9_1.1.14.13")
					assert.Equal(t, infoItems[0].Name, "Consumed Electrical Power Measure")
					assert.Equal(t, infoItems[0].Description.Text, "Power consumption values with timestamp.")

					if values := infoItems[0].Values; assert.Len(t, values, 5) {
						assert.Equal(t, values[0].DateTime, "2001-10-26T15:33:21")
						assert.Equal(t, values[0].Text, "15.5")
						assert.Equal(t, values[1].DateTime, "2001-10-26T15:33:50")
						assert.Equal(t, values[1].Text, "15.7")
						assert.Equal(t, values[2].DateTime, "2001-10-26T15:34:15")
						assert.Equal(t, values[2].Text, "1.3")
						assert.Equal(t, values[3].DateTime, "2001-10-26T15:34:35")
						assert.Equal(t, values[3].Text, "1.5")
						assert.Equal(t, values[4].DateTime, "2001-10-26T15:34:52")
						assert.Equal(t, values[4].Text, "15.3")
					}
				}

			}
		}
	}
}

func TestUnmarshalMetadataAboutRefrigeratorPowerConsumption(t *testing.T) {
	data, err := ioutil.ReadFile("examples/metadata_about_refrigerator_power_consumption.xml")
	if assert.Nil(t, err) {
		v, err := Unmarshal(data)
		if assert.Nil(t, err) {
			if assert.Len(t, v.Objects, 1) {
				assert.Equal(t, v.Objects[0].Id.Text, "SmartFridge22334411")

				if assert.Len(t, v.Objects[0].InfoItems, 1) {
					assert.Equal(t, v.Objects[0].InfoItems[0].Name, "PowerConsumption")

					if infoItems := v.Objects[0].InfoItems[0].MetaData.InfoItems; assert.Len(t, infoItems, 6) {
						assert.Equal(t, infoItems[0].Name, "format")
						if assert.Len(t, infoItems[0].Values, 1) {
							assert.Equal(t, infoItems[0].Values[0].Type, "xs:string")
							assert.Equal(t, infoItems[0].Values[0].Text, "xs:double")
						}

						assert.Equal(t, infoItems[1].Name, "latency")
						if assert.Len(t, infoItems[1].Values, 1) {
							assert.Equal(t, infoItems[1].Values[0].Type, "xs:int")
							assert.Equal(t, infoItems[1].Values[0].Text, "5")
						}

						assert.Equal(t, infoItems[2].Name, "readable")
						if assert.Len(t, infoItems[2].Values, 1) {
							assert.Equal(t, infoItems[2].Values[0].Type, "xs:boolean")
							assert.Equal(t, infoItems[2].Values[0].Text, "true")
						}

						assert.Equal(t, infoItems[3].Name, "writable")
						if assert.Len(t, infoItems[3].Values, 1) {
							assert.Equal(t, infoItems[3].Values[0].Type, "xs:boolean")
							assert.Equal(t, infoItems[3].Values[0].Text, "false")
						}

						assert.Equal(t, infoItems[4].Name, "unit")
						if assert.Len(t, infoItems[4].Values, 1) {
							assert.Equal(t, infoItems[4].Values[0].Type, "xs:string")
							assert.Equal(t, infoItems[4].Values[0].Text, "Watts")
						}

						assert.Equal(t, infoItems[5].Name, "accuracy")
						if assert.Len(t, infoItems[5].Values, 1) {
							assert.Equal(t, infoItems[5].Values[0].Type, "xs:double")
							assert.Equal(t, infoItems[5].Values[0].Text, "1")
						}
					}

				}
			}
		}
	}
}

func TestUnmarshalObjectObjectInfoitemValues(t *testing.T) {
	data, err := ioutil.ReadFile("examples/object_object_infoitem_values.xml")
	if assert.Nil(t, err) {
		v, err := Unmarshal(data)
		if assert.Nil(t, err) {
			if assert.Len(t, v.Objects, 1) {
				assert.Equal(t, v.Objects[0].Type, "someType")
				assert.Equal(t, v.Objects[0].Id.Text, "UniqueTargetID_1")
				if assert.Len(t, v.Objects[0].InfoItems, 2) {
					assert.Equal(t, v.Objects[0].InfoItems[0].Name, "InfoItem1")
					if assert.Len(t, v.Objects[0].InfoItems[0].Values, 3) {
						assert.Equal(t, v.Objects[0].InfoItems[0].Values[0].Text, "Value1")
						assert.Equal(t, v.Objects[0].InfoItems[0].Values[1].Text, "Value2")
						assert.Equal(t, v.Objects[0].InfoItems[0].Values[2].Text, "Value3")
					}
					assert.Equal(t, v.Objects[0].InfoItems[1].Name, "InfoItem2")
					if assert.Len(t, v.Objects[0].InfoItems[1].Values, 1) {
						assert.Equal(t, v.Objects[0].InfoItems[1].Values[0].Text, "Value")
					}
				}
				if assert.Len(t, v.Objects[0].Objects, 2) {
					subTarget1 := v.Objects[0].Objects[0]
					assert.Equal(t, subTarget1.Type, "someType")
					assert.Equal(t, subTarget1.Id.Text, "SubTarget1")
					if assert.Len(t, subTarget1.Objects, 1) {
						subSubTarget1 := subTarget1.Objects[0]
						assert.Equal(t, subSubTarget1.Type, "someType")
						assert.Equal(t, subSubTarget1.Id.Text, "SubSubTarget1")
						if assert.Len(t, subSubTarget1.InfoItems, 1) {
							assert.Equal(t, subSubTarget1.InfoItems[0].Name, "SubSubTarget1InfoItem1")
							if assert.Len(t, subSubTarget1.InfoItems[0].Values, 1) {
								assert.Equal(t, subSubTarget1.InfoItems[0].Values[0].Text, "22.5")
							}
						}
					}
					subTarget2 := v.Objects[0].Objects[1]
					assert.Equal(t, subTarget2.Type, "someType")
					assert.Equal(t, subTarget2.Id.Text, "SubTarget2")
					if assert.Len(t, subTarget2.InfoItems, 1) {
						assert.Equal(t, subTarget2.InfoItems[0].Name, "SubTarget2InfoItem1")
						if assert.Len(t, subTarget2.InfoItems[0].Values, 1) {
							assert.Equal(t, subTarget2.InfoItems[0].Values[0].Text, "34.6")
						}
					}
				}
			}
		}
	}
}

func TestUnmarshalObjectWithSubObjects(t *testing.T) {
	data, err := ioutil.ReadFile("examples/object_with_sub_objects.xml")
	if assert.Nil(t, err) {
		v, err := Unmarshal(data)
		if assert.Nil(t, err) {
			if assert.Len(t, v.Objects, 1) {
				assert.Equal(t, v.Objects[0].Type, "SOME_CLASS_PREFERABLY_DEFINED_BY_UDEF")
				assert.Equal(t, v.Objects[0].Udef, "appropriate.udef.code")
				assert.Equal(t, v.Objects[0].Id.Text, "UniqueObjectID_1")
				if assert.Len(t, v.Objects[0].InfoItems, 1) {
					assert.Equal(t, v.Objects[0].InfoItems[0].Name, "SOME_CLASS_PREFERABLY_DEFINED_BY_UDEF")
					assert.Equal(t, v.Objects[0].InfoItems[0].Udef, "appropriate.udef.code")
				}
				if assert.Len(t, v.Objects[0].Objects, 2) {
					assert.Equal(t, v.Objects[0].Objects[0].Type, "SOME_CLASS_PREFERABLY_DEFINED_BY_UDEF")
					assert.Equal(t, v.Objects[0].Objects[0].Udef, "appropriate.udef.code")
					assert.Equal(t, v.Objects[0].Objects[0].Id.Text, "SubObject1")

					assert.Equal(t, v.Objects[0].Objects[1].Type, "SOME_CLASS_PREFERABLY_DEFINED_BY_UDEF")
					assert.Equal(t, v.Objects[0].Objects[1].Udef, "appropriate.udef.code")
					assert.Equal(t, v.Objects[0].Objects[1].Id.Text, "SubObject2")
				}
			}
		}
	}
}
