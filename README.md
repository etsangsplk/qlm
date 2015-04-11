# O-DF/O-MI package for Go

This is a Go package for unmarshalling (XML to Go data structures) and
marshalling (Go data structures to XML) of  [Open Data
Format](https://www2.opengroup.org/ogsys/catalog/c14a) (O-DF) and  [Open
Messaging Interface](https://www2.opengroup.org/ogsys/catalog/c14b) (O-MI)
messages. O-DF and O-MI are part of an Open Group Internet of Things (IoT)
standard that was previously known as Quantum Lifecycle Mechanism (QLM).

This project was done for the T-106.5700 course at Aalto University.

## Installation

```bash
$ go get github.com/qlm-iot/qlm
```

## Examples

### Unmarshalling O-DF messages

```go
import "github.com/qlm-iot/qlm/df"

xml := `
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
objects, err := df.Unmarshal([]byte(xml))
```

### Unmarshalling O-MI messages

```go
import "github.com/qlm-iot/qlm/mi"

xml := `
    <?xml version="1.0" encoding="UTF-8"?>
    <omi:omiEnvelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:omi="omi.xsd" xsi:schemaLocation="omi.xsd" version="1.0" ttl="10">
        <omi:read msgformat="omi.xsd" interval="3.5" oldest="10" newest="15" begin="2014-01-01T00:00" end="2014-02-01T00:00">
            <omi:msg xmlns="odf.xsd" xsi:schemaLocation="odf.xsd">
                <Objects>
                    <Object>
                        <id>SmartFridge22334411</id>
                        <InfoItem name="PowerConsumption"></InfoItem>
                    </Object>
                </Objects>
            </omi:msg>
        </omi:read>
    </omi:omiEnvelope>
`
envelope, err := mi.Unmarshal([]byte(xml))
```

### Marshalling O-DF messages

```go
import "github.com/qlm-iot/qlm/df"

objects := df.Objects{
    Objects: []df.Object{
        df.Object{
            Type: "Refrigerator Assembly Product",
            Id:   &df.QLMID{Text: "SmartFridge22334411"},
            InfoItems: []df.InfoItem{
                df.InfoItem{
                    Udef: "b.o.9_1.1.14.13",
                    Name: "Consumed Electrical Power Measure",
                    Description: &df.Description{
                        Lang: "en",
                        Udef: "appropriate.udef.code",
                        Text: "Power consumption values with timestamp.",
                    },
                },
            },
        },
    },
}
xml, err := df.Marshal(objects)
```

### Marshalling O-MI messages

```go
import "github.com/qlm-iot/qlm/mi"

envelope := mi.OmiEnvelope{
    Version: "1.0",
    Ttl:     10,
    Read: &mi.ReadRequest{
        MsgFormat: "omi.xsd",
        Interval:  3.5,
        Oldest:    10,
        Newest:    15,
        Begin:     "2014-01-01T00:00",
        End:       "2014-02-01T00:00",
        Message: &mi.Message{
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

xml, err := mi.Marshal(envelope)
```

## Future work

- Add XML schema validation to unmarshalling functions.
- Add XML namespace support to marshalling functions.

## License

MIT
