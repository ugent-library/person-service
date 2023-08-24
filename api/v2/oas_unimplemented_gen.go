// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// GetOrganization implements getOrganization operation.
//
// Get single organization record.
//
// GET /organizations/{organizationId}
func (UnimplementedHandler) GetOrganization(ctx context.Context, params GetOrganizationParams) (r *Organization, _ error) {
	return r, ht.ErrNotImplemented
}

// GetOrganizations implements getOrganizations operation.
//
// Get all organization records.
//
// GET /organizations
func (UnimplementedHandler) GetOrganizations(ctx context.Context, params GetOrganizationsParams) (r *PagedOrganizationListResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetPeople implements getPeople operation.
//
// Get all person records.
//
// GET /people
func (UnimplementedHandler) GetPeople(ctx context.Context, params GetPeopleParams) (r *PagedPersonListResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetPerson implements getPerson operation.
//
// Retrieve a single person record.
//
// GET /people/{personId}
func (UnimplementedHandler) GetPerson(ctx context.Context, params GetPersonParams) (r *Person, _ error) {
	return r, ht.ErrNotImplemented
}

// SetPersonOrcid implements setPersonOrcid operation.
//
// Update person ORCID.
//
// PUT /people/{personId}/orcid
func (UnimplementedHandler) SetPersonOrcid(ctx context.Context, req *SetPersonOrcidRequest, params SetPersonOrcidParams) error {
	return ht.ErrNotImplemented
}

// SetPersonOrcidToken implements setPersonOrcidToken operation.
//
// Update person ORCID token.
//
// PUT /people/{personId}/orcid-token
func (UnimplementedHandler) SetPersonOrcidToken(ctx context.Context, req *SetPersonOrcidTokenRequest, params SetPersonOrcidTokenParams) error {
	return ht.ErrNotImplemented
}

// SetPersonRole implements setPersonRole operation.
//
// Update person role.
//
// PUT /people/{personId}/role
func (UnimplementedHandler) SetPersonRole(ctx context.Context, req *SetPersonRoleRequest, params SetPersonRoleParams) error {
	return ht.ErrNotImplemented
}

// SetPersonSettings implements setPersonSettings operation.
//
// Update person settings.
//
// PUT /people/{personId}/settings
func (UnimplementedHandler) SetPersonSettings(ctx context.Context, req *SetPersonSettingsRequest, params SetPersonSettingsParams) error {
	return ht.ErrNotImplemented
}

// SuggestOrganizations implements suggestOrganizations operation.
//
// Search on organization records.
//
// GET /organizations-suggest
func (UnimplementedHandler) SuggestOrganizations(ctx context.Context, params SuggestOrganizationsParams) (r *PagedOrganizationListResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// SuggestPeople implements suggestPeople operation.
//
// Search on person records.
//
// GET /people-suggest
func (UnimplementedHandler) SuggestPeople(ctx context.Context, params SuggestPeopleParams) (r *PagedPersonListResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// NewError creates *ErrorStatusCode from error returned by handler.
//
// Used for common default response.
func (UnimplementedHandler) NewError(ctx context.Context, err error) (r *ErrorStatusCode) {
	r = new(ErrorStatusCode)
	return r
}
