package appconfiguration

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// KeyValuesClient is the client for the KeyValues methods of the Appconfiguration service.
type KeyValuesClient struct {
	BaseClient
}

// NewKeyValuesClient creates an instance of the KeyValuesClient client.
func NewKeyValuesClient(subscriptionID string) KeyValuesClient {
	return NewKeyValuesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewKeyValuesClientWithBaseURI creates an instance of the KeyValuesClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewKeyValuesClientWithBaseURI(baseURI string, subscriptionID string) KeyValuesClient {
	return KeyValuesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates a key-value.
// Parameters:
// resourceGroupName - the name of the resource group to which the container registry belongs.
// configStoreName - the name of the configuration store.
// keyValueName - identifier of key and label combination. Key and label are joined by $ character. Label is
// optional.
// keyValueParameters - the parameters for creating a key-value.
func (client KeyValuesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, configStoreName string, keyValueName string, keyValueParameters *KeyValue) (result KeyValue, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/KeyValuesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: configStoreName,
			Constraints: []validation.Constraint{{Target: "configStoreName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "configStoreName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "configStoreName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("appconfiguration.KeyValuesClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, configStoreName, keyValueName, keyValueParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appconfiguration.KeyValuesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "appconfiguration.KeyValuesClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appconfiguration.KeyValuesClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client KeyValuesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, configStoreName string, keyValueName string, keyValueParameters *KeyValue) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"configStoreName":   autorest.Encode("path", configStoreName),
		"keyValueName":      autorest.Encode("path", keyValueName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	keyValueParameters.ID = nil
	keyValueParameters.Name = nil
	keyValueParameters.Type = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores/{configStoreName}/keyValues/{keyValueName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if keyValueParameters != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(keyValueParameters))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client KeyValuesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client KeyValuesClient) CreateOrUpdateResponder(resp *http.Response) (result KeyValue, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a key-value.
// Parameters:
// resourceGroupName - the name of the resource group to which the container registry belongs.
// configStoreName - the name of the configuration store.
// keyValueName - identifier of key and label combination. Key and label are joined by $ character. Label is
// optional.
func (client KeyValuesClient) Delete(ctx context.Context, resourceGroupName string, configStoreName string, keyValueName string) (result KeyValuesDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/KeyValuesClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: configStoreName,
			Constraints: []validation.Constraint{{Target: "configStoreName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "configStoreName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "configStoreName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("appconfiguration.KeyValuesClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, configStoreName, keyValueName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appconfiguration.KeyValuesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appconfiguration.KeyValuesClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client KeyValuesClient) DeletePreparer(ctx context.Context, resourceGroupName string, configStoreName string, keyValueName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"configStoreName":   autorest.Encode("path", configStoreName),
		"keyValueName":      autorest.Encode("path", keyValueName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores/{configStoreName}/keyValues/{keyValueName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client KeyValuesClient) DeleteSender(req *http.Request) (future KeyValuesDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client KeyValuesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the properties of the specified key-value.
// Parameters:
// resourceGroupName - the name of the resource group to which the container registry belongs.
// configStoreName - the name of the configuration store.
// keyValueName - identifier of key and label combination. Key and label are joined by $ character. Label is
// optional.
func (client KeyValuesClient) Get(ctx context.Context, resourceGroupName string, configStoreName string, keyValueName string) (result KeyValue, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/KeyValuesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: configStoreName,
			Constraints: []validation.Constraint{{Target: "configStoreName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "configStoreName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "configStoreName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("appconfiguration.KeyValuesClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, configStoreName, keyValueName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appconfiguration.KeyValuesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "appconfiguration.KeyValuesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appconfiguration.KeyValuesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client KeyValuesClient) GetPreparer(ctx context.Context, resourceGroupName string, configStoreName string, keyValueName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"configStoreName":   autorest.Encode("path", configStoreName),
		"keyValueName":      autorest.Encode("path", keyValueName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores/{configStoreName}/keyValues/{keyValueName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client KeyValuesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client KeyValuesClient) GetResponder(resp *http.Response) (result KeyValue, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByConfigurationStore lists the key-values for a given configuration store.
// Parameters:
// resourceGroupName - the name of the resource group to which the container registry belongs.
// configStoreName - the name of the configuration store.
// skipToken - a skip token is used to continue retrieving items after an operation returns a partial result.
// If a previous response contains a nextLink element, the value of the nextLink element will include a
// skipToken parameter that specifies a starting point to use for subsequent calls.
func (client KeyValuesClient) ListByConfigurationStore(ctx context.Context, resourceGroupName string, configStoreName string, skipToken string) (result KeyValueListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/KeyValuesClient.ListByConfigurationStore")
		defer func() {
			sc := -1
			if result.kvlr.Response.Response != nil {
				sc = result.kvlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: configStoreName,
			Constraints: []validation.Constraint{{Target: "configStoreName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "configStoreName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "configStoreName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("appconfiguration.KeyValuesClient", "ListByConfigurationStore", err.Error())
	}

	result.fn = client.listByConfigurationStoreNextResults
	req, err := client.ListByConfigurationStorePreparer(ctx, resourceGroupName, configStoreName, skipToken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appconfiguration.KeyValuesClient", "ListByConfigurationStore", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByConfigurationStoreSender(req)
	if err != nil {
		result.kvlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "appconfiguration.KeyValuesClient", "ListByConfigurationStore", resp, "Failure sending request")
		return
	}

	result.kvlr, err = client.ListByConfigurationStoreResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appconfiguration.KeyValuesClient", "ListByConfigurationStore", resp, "Failure responding to request")
		return
	}
	if result.kvlr.hasNextLink() && result.kvlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByConfigurationStorePreparer prepares the ListByConfigurationStore request.
func (client KeyValuesClient) ListByConfigurationStorePreparer(ctx context.Context, resourceGroupName string, configStoreName string, skipToken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"configStoreName":   autorest.Encode("path", configStoreName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(skipToken) > 0 {
		queryParameters["$skipToken"] = autorest.Encode("query", skipToken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores/{configStoreName}/keyValues", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByConfigurationStoreSender sends the ListByConfigurationStore request. The method will close the
// http.Response Body if it receives an error.
func (client KeyValuesClient) ListByConfigurationStoreSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByConfigurationStoreResponder handles the response to the ListByConfigurationStore request. The method always
// closes the http.Response Body.
func (client KeyValuesClient) ListByConfigurationStoreResponder(resp *http.Response) (result KeyValueListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByConfigurationStoreNextResults retrieves the next set of results, if any.
func (client KeyValuesClient) listByConfigurationStoreNextResults(ctx context.Context, lastResults KeyValueListResult) (result KeyValueListResult, err error) {
	req, err := lastResults.keyValueListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "appconfiguration.KeyValuesClient", "listByConfigurationStoreNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByConfigurationStoreSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "appconfiguration.KeyValuesClient", "listByConfigurationStoreNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByConfigurationStoreResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appconfiguration.KeyValuesClient", "listByConfigurationStoreNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByConfigurationStoreComplete enumerates all values, automatically crossing page boundaries as required.
func (client KeyValuesClient) ListByConfigurationStoreComplete(ctx context.Context, resourceGroupName string, configStoreName string, skipToken string) (result KeyValueListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/KeyValuesClient.ListByConfigurationStore")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByConfigurationStore(ctx, resourceGroupName, configStoreName, skipToken)
	return
}