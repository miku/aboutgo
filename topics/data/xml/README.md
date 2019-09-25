# XML

There is support for XML in the standard library, and a few additional libraries, that extend the functionality. The basic mapping is facilitated with struct tags.

> Go offers struct tags which are discoverable via reflection. These enjoy a
> wide range of use in the standard library in the JSON/XML and other encoding
> packages.

XML belongs to the well-known struct tags:

* [https://github.com/golang/go/wiki/Well-known-struct-tags](https://github.com/golang/go/wiki/Well-known-struct-tags)

## Annotations

Example struct with XML tags:

<!-- curl http://inspirehep.net/oai2d?verb=Identify | zek -e -t oai-identifier -->

```go
// OaiIdentifier was generated 2019-09-24 11:50:05 by tir on sol.
type OaiIdentifier struct {
        XMLName              xml.Name `xml:"oai-identifier"`
        Text                 string   `xml:",chardata"`
        Xmlns                string   `xml:"xmlns,attr"`
        Xsi                  string   `xml:"xsi,attr"`
        SchemaLocation       string   `xml:"schemaLocation,attr"`
        Scheme               string   `xml:"scheme"`               // oai
        RepositoryIdentifier string   `xml:"repositoryIdentifier"` // inspirehep.net
        Delimiter            string   `xml:"delimiter"`            // :
        SampleIdentifier     string   `xml:"sampleIdentifier"`     // oai:inspirehep.net:123
}

```

## Marshalling

The [xml.Marshal](https://golang.org/pkg/encoding/xml/#Marshal) functions returns the XML encoding of a value.

```go
var ip IPInfo
b, err := xml.Marshal(ip)
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(b))
```

## Unmarshalling

We can turn bytes into struct through [xml.Unmarshal](https://golang.org/pkg/encoding/xml/#Unmarshal).

```go
var ip IPInfo
err := xml.Unmarshal(data, ip)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%v", ip)
```

## Helper

A tool like [XMLGen](https://github.com/dutchcoders/XMLGen) or
[zek](https://github.com/miku/zek) (also online:
https://www.onlinetool.io/xmltogo/) can help shorten the process from data to
structs.