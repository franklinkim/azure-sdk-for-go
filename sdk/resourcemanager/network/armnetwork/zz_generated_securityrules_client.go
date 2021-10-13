//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// SecurityRulesClient contains the methods for the SecurityRules group.
// Don't use this type directly, use NewSecurityRulesClient() instead.
type SecurityRulesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewSecurityRulesClient creates a new instance of SecurityRulesClient with the specified values.
func NewSecurityRulesClient(con *arm.Connection, subscriptionID string) *SecurityRulesClient {
	return &SecurityRulesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates a security rule in the specified network security group.
// If the operation fails it returns the *CloudError error type.
func (client *SecurityRulesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, securityRuleName string, securityRuleParameters SecurityRule, options *SecurityRulesBeginCreateOrUpdateOptions) (SecurityRulesCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, networkSecurityGroupName, securityRuleName, securityRuleParameters, options)
	if err != nil {
		return SecurityRulesCreateOrUpdatePollerResponse{}, err
	}
	result := SecurityRulesCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("SecurityRulesClient.CreateOrUpdate", "azure-async-operation", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return SecurityRulesCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &SecurityRulesCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a security rule in the specified network security group.
// If the operation fails it returns the *CloudError error type.
func (client *SecurityRulesClient) createOrUpdate(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, securityRuleName string, securityRuleParameters SecurityRule, options *SecurityRulesBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, networkSecurityGroupName, securityRuleName, securityRuleParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SecurityRulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, securityRuleName string, securityRuleParameters SecurityRule, options *SecurityRulesBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/securityRules/{securityRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkSecurityGroupName == "" {
		return nil, errors.New("parameter networkSecurityGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkSecurityGroupName}", url.PathEscape(networkSecurityGroupName))
	if securityRuleName == "" {
		return nil, errors.New("parameter securityRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityRuleName}", url.PathEscape(securityRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, securityRuleParameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *SecurityRulesClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Deletes the specified network security rule.
// If the operation fails it returns the *CloudError error type.
func (client *SecurityRulesClient) BeginDelete(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, securityRuleName string, options *SecurityRulesBeginDeleteOptions) (SecurityRulesDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, networkSecurityGroupName, securityRuleName, options)
	if err != nil {
		return SecurityRulesDeletePollerResponse{}, err
	}
	result := SecurityRulesDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("SecurityRulesClient.Delete", "location", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return SecurityRulesDeletePollerResponse{}, err
	}
	result.Poller = &SecurityRulesDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the specified network security rule.
// If the operation fails it returns the *CloudError error type.
func (client *SecurityRulesClient) deleteOperation(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, securityRuleName string, options *SecurityRulesBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, networkSecurityGroupName, securityRuleName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SecurityRulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, securityRuleName string, options *SecurityRulesBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/securityRules/{securityRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkSecurityGroupName == "" {
		return nil, errors.New("parameter networkSecurityGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkSecurityGroupName}", url.PathEscape(networkSecurityGroupName))
	if securityRuleName == "" {
		return nil, errors.New("parameter securityRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityRuleName}", url.PathEscape(securityRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *SecurityRulesClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Get the specified network security rule.
// If the operation fails it returns the *CloudError error type.
func (client *SecurityRulesClient) Get(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, securityRuleName string, options *SecurityRulesGetOptions) (SecurityRulesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, networkSecurityGroupName, securityRuleName, options)
	if err != nil {
		return SecurityRulesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SecurityRulesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SecurityRulesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SecurityRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, securityRuleName string, options *SecurityRulesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/securityRules/{securityRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkSecurityGroupName == "" {
		return nil, errors.New("parameter networkSecurityGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkSecurityGroupName}", url.PathEscape(networkSecurityGroupName))
	if securityRuleName == "" {
		return nil, errors.New("parameter securityRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityRuleName}", url.PathEscape(securityRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SecurityRulesClient) getHandleResponse(resp *http.Response) (SecurityRulesGetResponse, error) {
	result := SecurityRulesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecurityRule); err != nil {
		return SecurityRulesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *SecurityRulesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Gets all security rules in a network security group.
// If the operation fails it returns the *CloudError error type.
func (client *SecurityRulesClient) List(resourceGroupName string, networkSecurityGroupName string, options *SecurityRulesListOptions) *SecurityRulesListPager {
	return &SecurityRulesListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, networkSecurityGroupName, options)
		},
		advancer: func(ctx context.Context, resp SecurityRulesListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SecurityRuleListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *SecurityRulesClient) listCreateRequest(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, options *SecurityRulesListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/securityRules"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkSecurityGroupName == "" {
		return nil, errors.New("parameter networkSecurityGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkSecurityGroupName}", url.PathEscape(networkSecurityGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SecurityRulesClient) listHandleResponse(resp *http.Response) (SecurityRulesListResponse, error) {
	result := SecurityRulesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecurityRuleListResult); err != nil {
		return SecurityRulesListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *SecurityRulesClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}