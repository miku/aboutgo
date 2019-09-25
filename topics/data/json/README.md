# JSON

JSON data can be serialized from structs, via struct tags.

> Go offers struct tags which are discoverable via reflection. These enjoy a
> wide range of use in the standard library in the JSON/XML and other encoding
> packages.

JSON belongs to the well-known struct tags:

* [https://github.com/golang/go/wiki/Well-known-struct-tags](https://github.com/golang/go/wiki/Well-known-struct-tags)

## Annotations

Example struct with JSON tags:

```go
type IPInfo struct {
        City     string `json:"city"`
        Country  string `json:"country"`
        IP       string `json:"ip"`
        Loc      string `json:"loc"`
        Org      string `json:"org"`
        Postal   string `json:"postal"`
        Readme   string `json:"readme"`
        Region   string `json:"region"`
        Timezone string `json:"timezone"`
}
```

## Marshalling

The [json.Marshal](https://golang.org/pkg/encoding/json/#Marshal) functions returns the JSON encoding of a value.

```go
var ip IPInfo
b, err := json.Marshal(ip)
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(b))
```

## Unmarshalling

We can turn bytes into struct through [json.Unmarshal](https://golang.org/pkg/encoding/json/#Unmarshal).

```go
var ip IPInfo
err := json.Unmarshal(data, ip)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%v", ip)
```

## Helper

A tool like [JSONGen](https://github.com/bemasher/JSONGen) can help shorten the process from data to structs.

## Exercise

