package redhatopenshiftapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/redhatopenshift/mgmt/2022-04-01/redhatopenshift"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result redhatopenshift.OperationListPage, err error)
	ListComplete(ctx context.Context) (result redhatopenshift.OperationListIterator, err error)
}

var _ OperationsClientAPI = (*redhatopenshift.OperationsClient)(nil)

// OpenShiftClustersClientAPI contains the set of methods on the OpenShiftClustersClient type.
type OpenShiftClustersClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, parameters redhatopenshift.OpenShiftCluster) (result redhatopenshift.OpenShiftClustersCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceName string) (result redhatopenshift.OpenShiftClustersDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, resourceName string) (result redhatopenshift.OpenShiftCluster, err error)
	List(ctx context.Context) (result redhatopenshift.OpenShiftClusterListPage, err error)
	ListComplete(ctx context.Context) (result redhatopenshift.OpenShiftClusterListIterator, err error)
	ListAdminCredentials(ctx context.Context, resourceGroupName string, resourceName string) (result redhatopenshift.OpenShiftClusterAdminKubeconfig, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result redhatopenshift.OpenShiftClusterListPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result redhatopenshift.OpenShiftClusterListIterator, err error)
	ListCredentials(ctx context.Context, resourceGroupName string, resourceName string) (result redhatopenshift.OpenShiftClusterCredentials, err error)
	Update(ctx context.Context, resourceGroupName string, resourceName string, parameters redhatopenshift.OpenShiftClusterUpdate) (result redhatopenshift.OpenShiftClustersUpdateFuture, err error)
}

var _ OpenShiftClustersClientAPI = (*redhatopenshift.OpenShiftClustersClient)(nil)
