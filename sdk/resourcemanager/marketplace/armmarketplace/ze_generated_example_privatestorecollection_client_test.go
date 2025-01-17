//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmarketplace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-12-01/examples/GetPrivateStoreCollectionsList.json
func ExamplePrivateStoreCollectionClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmarketplace.NewPrivateStoreCollectionClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.List(ctx,
		"a0e28e55-90c4-41d8-8e34-bb7ef7775406",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-12-01/examples/GetPrivateStoreCollection.json
func ExamplePrivateStoreCollectionClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmarketplace.NewPrivateStoreCollectionClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"a0e28e55-90c4-41d8-8e34-bb7ef7775406",
		"56a1a02d-8cf8-45df-bf37-d5f7120fcb3d",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-12-01/examples/CreatePrivateStoreCollection.json
func ExamplePrivateStoreCollectionClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmarketplace.NewPrivateStoreCollectionClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"a0e28e55-90c4-41d8-8e34-bb7ef7775406",
		"d0f5aa2c-ecc3-4d87-906a-f8c486dcc4f1",
		&armmarketplace.PrivateStoreCollectionClientCreateOrUpdateOptions{Payload: &armmarketplace.Collection{
			Properties: &armmarketplace.CollectionProperties{
				AllSubscriptions: to.Ptr(false),
				Claim:            to.Ptr(""),
				CollectionName:   to.Ptr("Test Collection"),
				SubscriptionsList: []*string{
					to.Ptr("b340914e-353d-453a-85fb-8f9b65b51f91"),
					to.Ptr("f2baa04d-5bfc-461b-b6d8-61b403c9ec48")},
			},
		},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-12-01/examples/DeletePrivateStoreCollection.json
func ExamplePrivateStoreCollectionClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmarketplace.NewPrivateStoreCollectionClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx,
		"a0e28e55-90c4-41d8-8e34-bb7ef7775406",
		"d0f5aa2c-ecc3-4d87-906a-f8c486dcc4f1",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-12-01/examples/PostPrivateStoreCollection.json
func ExamplePrivateStoreCollectionClient_Post() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmarketplace.NewPrivateStoreCollectionClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Post(ctx,
		"a0e28e55-90c4-41d8-8e34-bb7ef7775406",
		"56a1a02d-8cf8-45df-bf37-d5f7120fcb3d",
		&armmarketplace.PrivateStoreCollectionClientPostOptions{Payload: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-12-01/examples/TransferOffers.json
func ExamplePrivateStoreCollectionClient_TransferOffers() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmarketplace.NewPrivateStoreCollectionClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.TransferOffers(ctx,
		"a0e28e55-90c4-41d8-8e34-bb7ef7775406",
		"56a1a02d-8cf8-45df-bf37-d5f7120fcb3d",
		&armmarketplace.PrivateStoreCollectionClientTransferOffersOptions{Payload: &armmarketplace.TransferOffersProperties{
			Properties: &armmarketplace.TransferOffersDetails{
				OfferIDsList: []*string{
					to.Ptr("marketplacetestthirdparty.md-test-third-party-2"),
					to.Ptr("marketplacetestthirdparty.md-test-third-party-3")},
				Operation: to.Ptr("copy"),
				TargetCollections: []*string{
					to.Ptr("c752f021-1c37-4af5-b82f-74c51c27b44a"),
					to.Ptr("f47ef1c7-e908-4f39-ae29-db181634ad8d")},
			},
		},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
