// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package identity

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetIdpGroupMappingRequest wrapper for the GetIdpGroupMapping operation
type GetIdpGroupMappingRequest struct {

	// The OCID of the identity provider.
	IdentityProviderId *string `mandatory:"true" contributesTo:"path" name:"identityProviderId"`

	// The OCID of the group mapping.
	MappingId *string `mandatory:"true" contributesTo:"path" name:"mappingId"`
}

func (request GetIdpGroupMappingRequest) String() string {
	return common.PointerString(request)
}

// GetIdpGroupMappingResponse wrapper for the GetIdpGroupMapping operation
type GetIdpGroupMappingResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The IdpGroupMapping instance
	IdpGroupMapping `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`
}

func (response GetIdpGroupMappingResponse) String() string {
	return common.PointerString(response)
}
