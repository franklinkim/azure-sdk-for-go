// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcosmos

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// RestorableDatabaseAccountsClient contains the methods for the RestorableDatabaseAccounts group.
// Don't use this type directly, use NewRestorableDatabaseAccountsClient() instead.
type RestorableDatabaseAccountsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewRestorableDatabaseAccountsClient creates a new instance of RestorableDatabaseAccountsClient with the specified values.
func NewRestorableDatabaseAccountsClient(con *armcore.Connection, subscriptionID string) *RestorableDatabaseAccountsClient {
	return &RestorableDatabaseAccountsClient{con: con, subscriptionID: subscriptionID}
}

// GetByLocation - Retrieves the properties of an existing Azure Cosmos DB restorable database account. This call requires 'Microsoft.DocumentDB/locations/restorableDatabaseAccounts/read/*'
// permission.
// If the operation fails it returns the *CloudError error type.
func (client *RestorableDatabaseAccountsClient) GetByLocation(ctx context.Context, location string, instanceID string, options *RestorableDatabaseAccountsGetByLocationOptions) (RestorableDatabaseAccountsGetByLocationResponse, error) {
	req, err := client.getByLocationCreateRequest(ctx, location, instanceID, options)
	if err != nil {
		return RestorableDatabaseAccountsGetByLocationResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return RestorableDatabaseAccountsGetByLocationResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return RestorableDatabaseAccountsGetByLocationResponse{}, client.getByLocationHandleError(resp)
	}
	return client.getByLocationHandleResponse(resp)
}

// getByLocationCreateRequest creates the GetByLocation request.
func (client *RestorableDatabaseAccountsClient) getByLocationCreateRequest(ctx context.Context, location string, instanceID string, options *RestorableDatabaseAccountsGetByLocationOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/locations/{location}/restorableDatabaseAccounts/{instanceId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if instanceID == "" {
		return nil, errors.New("parameter instanceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-06-15")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getByLocationHandleResponse handles the GetByLocation response.
func (client *RestorableDatabaseAccountsClient) getByLocationHandleResponse(resp *azcore.Response) (RestorableDatabaseAccountsGetByLocationResponse, error) {
	result := RestorableDatabaseAccountsGetByLocationResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.RestorableDatabaseAccountGetResult); err != nil {
		return RestorableDatabaseAccountsGetByLocationResponse{}, err
	}
	return result, nil
}

// getByLocationHandleError handles the GetByLocation error response.
func (client *RestorableDatabaseAccountsClient) getByLocationHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// List - Lists all the restorable Azure Cosmos DB database accounts available under the subscription. This call requires 'Microsoft.DocumentDB/locations/restorableDatabaseAccounts/read'
// permission.
// If the operation fails it returns the *CloudError error type.
func (client *RestorableDatabaseAccountsClient) List(ctx context.Context, options *RestorableDatabaseAccountsListOptions) (RestorableDatabaseAccountsListResponse, error) {
	req, err := client.listCreateRequest(ctx, options)
	if err != nil {
		return RestorableDatabaseAccountsListResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return RestorableDatabaseAccountsListResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return RestorableDatabaseAccountsListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *RestorableDatabaseAccountsClient) listCreateRequest(ctx context.Context, options *RestorableDatabaseAccountsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/restorableDatabaseAccounts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-06-15")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RestorableDatabaseAccountsClient) listHandleResponse(resp *azcore.Response) (RestorableDatabaseAccountsListResponse, error) {
	result := RestorableDatabaseAccountsListResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.RestorableDatabaseAccountsListResult); err != nil {
		return RestorableDatabaseAccountsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *RestorableDatabaseAccountsClient) listHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// ListByLocation - Lists all the restorable Azure Cosmos DB database accounts available under the subscription and in a region. This call requires 'Microsoft.DocumentDB/locations/restorableDatabaseAccounts/read'
// permission.
// If the operation fails it returns the *CloudError error type.
func (client *RestorableDatabaseAccountsClient) ListByLocation(ctx context.Context, location string, options *RestorableDatabaseAccountsListByLocationOptions) (RestorableDatabaseAccountsListByLocationResponse, error) {
	req, err := client.listByLocationCreateRequest(ctx, location, options)
	if err != nil {
		return RestorableDatabaseAccountsListByLocationResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return RestorableDatabaseAccountsListByLocationResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return RestorableDatabaseAccountsListByLocationResponse{}, client.listByLocationHandleError(resp)
	}
	return client.listByLocationHandleResponse(resp)
}

// listByLocationCreateRequest creates the ListByLocation request.
func (client *RestorableDatabaseAccountsClient) listByLocationCreateRequest(ctx context.Context, location string, options *RestorableDatabaseAccountsListByLocationOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/locations/{location}/restorableDatabaseAccounts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-06-15")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByLocationHandleResponse handles the ListByLocation response.
func (client *RestorableDatabaseAccountsClient) listByLocationHandleResponse(resp *azcore.Response) (RestorableDatabaseAccountsListByLocationResponse, error) {
	result := RestorableDatabaseAccountsListByLocationResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.RestorableDatabaseAccountsListResult); err != nil {
		return RestorableDatabaseAccountsListByLocationResponse{}, err
	}
	return result, nil
}

// listByLocationHandleError handles the ListByLocation error response.
func (client *RestorableDatabaseAccountsClient) listByLocationHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}