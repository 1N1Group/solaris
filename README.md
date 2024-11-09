# 

## Functional coverage

### Symbols

- ✅ - Implemented
- 💊 - On develope
- 🐞 - Bug
- ❌ - Not implemented

### Solr API V2 `solaris/v2`
**Source**: https://downloads.apache.org/solr/solr/9.7.0/openApi/solr-openapi-9.7.0.json

| **Alias** | | |
| - | - | - |
| `Get` | Get details for a specific collection alias |💊|
| `GetList` | List the existing collection aliases |💊|
| `Delete` | Deletes an alias by its name |💊|
| `GetProperty` | Get a specific property for a collection alias |💊|
| `GetPropertiesList` | Get properties for a collection alias |💊|
| `UpdateProperty` | Update a specific property for a collection alias |💊|
| `UpdatePropertiesList` | Update properties for a collection alias |💊|
| `DeleteProperty` | Delete a specific property for a collection alias |💊|

| **Replica** | | |
| - | - | - |
| `Create` | Creates a new replica of an existing shard |💊|
| `Delete` | Delete an single replica by name |💊|
| `DeleteMany` | Delete one or more replicas from the specified collection and shard |💊|
| `AddProperty` | Adds a property to the specified replica |💊|
| `DeleteProperty` | Delete an existing replica property |💊|