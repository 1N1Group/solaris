# Solaris | Apache Solr Client on Golang 
[![Go Reference](https://pkg.go.dev/badge/github.com/1N1Group/solaris.svg)](https://pkg.go.dev/github.com/1N1Group/solaris)

**Source API configuration**: https://downloads.apache.org/solr/solr/9.7.0/openApi/solr-openapi-9.7.0.json

## Install
```bash
go get github.com/1N1Group/solaris
```

## Usage

```golang
import "github.com/1N1Group/solaris/v2"
```

## Functional coverage

### Coverage statuses

- ✅ - Implemented
- 💊 - On develope
- 🐞 - Bug
- ❌ - Not implemented

### Alias

|   |   |   |   |
| - | - | - | - |
| `Get` | Get details for a specific collection alias | [Example](./examples/alias/README.md#get) |💊|
| `GetList` | List the existing collection aliases | [Example](./examples/alias/README.md#getlist) |💊|
| `Delete` | Deletes an alias by its name | [Example](./examples/alias/README.md#delete) |💊|
| `GetProperty` | Get a specific property for a collection alias | [Example](./examples/alias/README.md#getproperty) |💊|
| `GetPropertiesList` | Get properties for a collection alias | [Example](./examples/alias/README.md#getpropertieslist) |💊|
| `UpdateProperty` | Update a specific property for a collection alias | [Example](./examples/alias/README.md#updateproperty) |💊|
| `UpdatePropertiesList` | Update properties for a collection alias | [Example](./examples/alias/README.md#updatepropertieslist) |💊|
| `DeleteProperty` | Delete a specific property for a collection alias | [Example](./examples/alias/README.md#deleteproperty) |💊|

### Replica

|   |   |   |   |
| - | - | - | - |
| `Create` | Creates a new replica of an existing shard | [Example](./examples/replica/README.md#create) |💊|
| `Delete` | Delete an single replica by name | [Example](./examples/replica/README.md#delete) |💊|
| `DeleteMany` | Delete one or more replicas from the specified collection and shard | [Example](./examples/replica/README.md#deletemany) |💊|
| `AddProperty` | Adds a property to the specified replica | [Example](./examples/replica/README.md#addproperty) |💊|
| `DeleteProperty` | Delete an existing replica property | [Example](./examples/replica/README.md#deleteproperty) |💊|
