package servicefabric

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/servicefabric/6.2/servicefabric instead
// ServiceFromTemplatesClient is the client for the ServiceFromTemplates methods of the Servicefabric service.
type ServiceFromTemplatesClient struct {
	BaseClient
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/servicefabric/6.2/servicefabric instead
// NewServiceFromTemplatesClient creates an instance of the ServiceFromTemplatesClient client.
func NewServiceFromTemplatesClient(timeout *int32) ServiceFromTemplatesClient {
	return NewServiceFromTemplatesClientWithBaseURI(DefaultBaseURI, timeout)
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/servicefabric/6.2/servicefabric instead
// NewServiceFromTemplatesClientWithBaseURI creates an instance of the ServiceFromTemplatesClient client.
func NewServiceFromTemplatesClientWithBaseURI(baseURI string, timeout *int32) ServiceFromTemplatesClient {
	return ServiceFromTemplatesClient{NewWithBaseURI(baseURI, timeout)}
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/servicefabric/6.2/servicefabric instead
// Create create service from templates
// Parameters:
// applicationName - the name of the application
// serviceDescriptionTemplate - the template of the service description
func (client ServiceFromTemplatesClient) Create(ctx context.Context, applicationName string, serviceDescriptionTemplate ServiceDescriptionTemplate) (result String, err error) {
	req, err := client.CreatePreparer(ctx, applicationName, serviceDescriptionTemplate)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ServiceFromTemplatesClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.ServiceFromTemplatesClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ServiceFromTemplatesClient", "Create", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/servicefabric/6.2/servicefabric instead
// CreatePreparer prepares the Create request.
func (client ServiceFromTemplatesClient) CreatePreparer(ctx context.Context, applicationName string, serviceDescriptionTemplate ServiceDescriptionTemplate) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationName": applicationName,
	}

	const APIVersion = "1.0.0"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if client.Timeout != nil {
		queryParameters["timeout"] = autorest.Encode("query", *client.Timeout)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Applications/{applicationName}/$/GetServices/$/CreateFromTemplate", pathParameters),
		autorest.WithJSON(serviceDescriptionTemplate),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/servicefabric/6.2/servicefabric instead
// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceFromTemplatesClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/servicefabric/6.2/servicefabric instead
// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ServiceFromTemplatesClient) CreateResponder(resp *http.Response) (result String, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
