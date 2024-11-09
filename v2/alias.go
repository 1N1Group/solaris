package v2

type AliasesAPI struct {
	client *Client
}

// Get a specific property for a collection alias.
//
// aliasName - Alias Name
//
// propName - Property Name
//
// GET /aliases/{aliasName}/properties/{propName}
func (a *AliasesAPI) GetProperty(aliasName string, propName string) {}

// Get properties for a collection alias.
//
// aliasName - Alias Name
//
// GET /aliases/{aliasName}/properties
func (a *AliasesAPI) GetPropertiesList(aliasName string) {}

type UpdateAliasPropertyRequestBody struct {
	Value interface{} `json:"value" validate:"required"`
}

// Update a specific property for a collection alias.
//
// aliasName - Alias Name
//
// propName - Property Name
//
// body - Property value that needs to be updated
//
// PUT /aliases/{aliasName}/properties/{propName}
func (a *AliasesAPI) UpdateProperty(aliasName string, propName string, body UpdateAliasPropertyRequestBody) {
}

type UpdateAliasPropertiesRequestBody struct {
	Properties Properties `json:"properties" validate:"required"` // Properties and values to be updated on alias.
	Async      string     `json:"async,omitempty"`                // Request ID to track this action which will be processed asynchronously.
}

// Update properties for a collection alias.
//
// aliasName - Alias Name
//
// body - Properties that need to be updated
//
// PUT /aliases/{aliasName}/properties
func (a *AliasesAPI) UpdatePropertiesList(aliasName string, body UpdateAliasPropertiesRequestBody) {
}

// Delete a specific property for a collection alias.
//
// aliasName - Alias Name
//
// propName - Property Name
//
// DELETE /aliases/{aliasName}/properties/{propName}
func (a *AliasesAPI) DeleteProperty(aliasName string, propName string) {}

// List the existing collection aliases.
//
// GET /aliases
func (a *AliasesAPI) GetList() {}

// Get details for a specific collection alias.
//
// aliasName - Alias Name
//
// GET /aliases/{aliasName}
func (a *AliasesAPI) Get(aliasName string) {}

type DeleteAliasParams struct {
	Async string `url:"async,omitempty"`
}

// Deletes an alias by its name
//
// aliasName - The name of the alias to delete
//
// params - The configuration for deleting
//
// DELETE /aliases/{aliasName}
func (a *AliasesAPI) Delete(aliasName string, params DeleteAliasParams) {}
