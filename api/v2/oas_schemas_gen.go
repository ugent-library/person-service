// Code generated by ogen, DO NOT EDIT.

package api

import (
	"fmt"
	"time"
)

func (s *ErrorStatusCode) Error() string {
	return fmt.Sprintf("code %d: %+v", s.StatusCode, s.Response)
}

type ApiKey struct {
	APIKey string
}

// GetAPIKey returns the value of APIKey.
func (s *ApiKey) GetAPIKey() string {
	return s.APIKey
}

// SetAPIKey sets the value of APIKey.
func (s *ApiKey) SetAPIKey(val string) {
	s.APIKey = val
}

// Ref: #/components/schemas/Error
type Error struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

// GetCode returns the value of Code.
func (s *Error) GetCode() int64 {
	return s.Code
}

// GetMessage returns the value of Message.
func (s *Error) GetMessage() string {
	return s.Message
}

// SetCode sets the value of Code.
func (s *Error) SetCode(val int64) {
	s.Code = val
}

// SetMessage sets the value of Message.
func (s *Error) SetMessage(val string) {
	s.Message = val
}

// ErrorStatusCode wraps Error with StatusCode.
type ErrorStatusCode struct {
	StatusCode int
	Response   Error
}

// GetStatusCode returns the value of StatusCode.
func (s *ErrorStatusCode) GetStatusCode() int {
	return s.StatusCode
}

// GetResponse returns the value of Response.
func (s *ErrorStatusCode) GetResponse() Error {
	return s.Response
}

// SetStatusCode sets the value of StatusCode.
func (s *ErrorStatusCode) SetStatusCode(val int) {
	s.StatusCode = val
}

// SetResponse sets the value of Response.
func (s *ErrorStatusCode) SetResponse(val Error) {
	s.Response = val
}

// Ref: #/components/schemas/IdRef
type IdRef struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// GetID returns the value of ID.
func (s *IdRef) GetID() string {
	return s.ID
}

// GetType returns the value of Type.
func (s *IdRef) GetType() string {
	return s.Type
}

// SetID sets the value of ID.
func (s *IdRef) SetID(val string) {
	s.ID = val
}

// SetType sets the value of Type.
func (s *IdRef) SetType(val string) {
	s.Type = val
}

// NewOptDateTime returns new OptDateTime with value set to v.
func NewOptDateTime(v time.Time) OptDateTime {
	return OptDateTime{
		Value: v,
		Set:   true,
	}
}

// OptDateTime is optional time.Time.
type OptDateTime struct {
	Value time.Time
	Set   bool
}

// IsSet returns true if OptDateTime was set.
func (o OptDateTime) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptDateTime) Reset() {
	var v time.Time
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptDateTime) SetTo(v time.Time) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptDateTime) Get() (v time.Time, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptDateTime) Or(d time.Time) time.Time {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptPersonSettings returns new OptPersonSettings with value set to v.
func NewOptPersonSettings(v PersonSettings) OptPersonSettings {
	return OptPersonSettings{
		Value: v,
		Set:   true,
	}
}

// OptPersonSettings is optional PersonSettings.
type OptPersonSettings struct {
	Value PersonSettings
	Set   bool
}

// IsSet returns true if OptPersonSettings was set.
func (o OptPersonSettings) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptPersonSettings) Reset() {
	var v PersonSettings
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptPersonSettings) SetTo(v PersonSettings) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptPersonSettings) Get() (v PersonSettings, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptPersonSettings) Or(d PersonSettings) PersonSettings {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// Ref: #/components/schemas/Organization
type Organization struct {
	ID          string    `json:"id"`
	GismoID     OptString `json:"gismo_id"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
	Type        string    `json:"type"`
	NameDut     OptString `json:"name_dut"`
	NameEng     OptString `json:"name_eng"`
	ParentID    OptString `json:"parent_id"`
	OtherID     []IdRef   `json:"other_id"`
}

// GetID returns the value of ID.
func (s *Organization) GetID() string {
	return s.ID
}

// GetGismoID returns the value of GismoID.
func (s *Organization) GetGismoID() OptString {
	return s.GismoID
}

// GetDateCreated returns the value of DateCreated.
func (s *Organization) GetDateCreated() time.Time {
	return s.DateCreated
}

// GetDateUpdated returns the value of DateUpdated.
func (s *Organization) GetDateUpdated() time.Time {
	return s.DateUpdated
}

// GetType returns the value of Type.
func (s *Organization) GetType() string {
	return s.Type
}

// GetNameDut returns the value of NameDut.
func (s *Organization) GetNameDut() OptString {
	return s.NameDut
}

// GetNameEng returns the value of NameEng.
func (s *Organization) GetNameEng() OptString {
	return s.NameEng
}

// GetParentID returns the value of ParentID.
func (s *Organization) GetParentID() OptString {
	return s.ParentID
}

// GetOtherID returns the value of OtherID.
func (s *Organization) GetOtherID() []IdRef {
	return s.OtherID
}

// SetID sets the value of ID.
func (s *Organization) SetID(val string) {
	s.ID = val
}

// SetGismoID sets the value of GismoID.
func (s *Organization) SetGismoID(val OptString) {
	s.GismoID = val
}

// SetDateCreated sets the value of DateCreated.
func (s *Organization) SetDateCreated(val time.Time) {
	s.DateCreated = val
}

// SetDateUpdated sets the value of DateUpdated.
func (s *Organization) SetDateUpdated(val time.Time) {
	s.DateUpdated = val
}

// SetType sets the value of Type.
func (s *Organization) SetType(val string) {
	s.Type = val
}

// SetNameDut sets the value of NameDut.
func (s *Organization) SetNameDut(val OptString) {
	s.NameDut = val
}

// SetNameEng sets the value of NameEng.
func (s *Organization) SetNameEng(val OptString) {
	s.NameEng = val
}

// SetParentID sets the value of ParentID.
func (s *Organization) SetParentID(val OptString) {
	s.ParentID = val
}

// SetOtherID sets the value of OtherID.
func (s *Organization) SetOtherID(val []IdRef) {
	s.OtherID = val
}

// Ref: #/components/schemas/OrganizationRef
type OrganizationRef struct {
	ID          string      `json:"id"`
	DateCreated time.Time   `json:"date_created"`
	DateUpdated time.Time   `json:"date_updated"`
	From        time.Time   `json:"from"`
	Until       OptDateTime `json:"until"`
}

// GetID returns the value of ID.
func (s *OrganizationRef) GetID() string {
	return s.ID
}

// GetDateCreated returns the value of DateCreated.
func (s *OrganizationRef) GetDateCreated() time.Time {
	return s.DateCreated
}

// GetDateUpdated returns the value of DateUpdated.
func (s *OrganizationRef) GetDateUpdated() time.Time {
	return s.DateUpdated
}

// GetFrom returns the value of From.
func (s *OrganizationRef) GetFrom() time.Time {
	return s.From
}

// GetUntil returns the value of Until.
func (s *OrganizationRef) GetUntil() OptDateTime {
	return s.Until
}

// SetID sets the value of ID.
func (s *OrganizationRef) SetID(val string) {
	s.ID = val
}

// SetDateCreated sets the value of DateCreated.
func (s *OrganizationRef) SetDateCreated(val time.Time) {
	s.DateCreated = val
}

// SetDateUpdated sets the value of DateUpdated.
func (s *OrganizationRef) SetDateUpdated(val time.Time) {
	s.DateUpdated = val
}

// SetFrom sets the value of From.
func (s *OrganizationRef) SetFrom(val time.Time) {
	s.From = val
}

// SetUntil sets the value of Until.
func (s *OrganizationRef) SetUntil(val OptDateTime) {
	s.Until = val
}

// Ref: #/components/schemas/PagedOrganizationListResponse
type PagedOrganizationListResponse struct {
	Cursor OptString      `json:"cursor"`
	Data   []Organization `json:"data"`
}

// GetCursor returns the value of Cursor.
func (s *PagedOrganizationListResponse) GetCursor() OptString {
	return s.Cursor
}

// GetData returns the value of Data.
func (s *PagedOrganizationListResponse) GetData() []Organization {
	return s.Data
}

// SetCursor sets the value of Cursor.
func (s *PagedOrganizationListResponse) SetCursor(val OptString) {
	s.Cursor = val
}

// SetData sets the value of Data.
func (s *PagedOrganizationListResponse) SetData(val []Organization) {
	s.Data = val
}

// Ref: #/components/schemas/PagedPersonListResponse
type PagedPersonListResponse struct {
	Cursor OptString `json:"cursor"`
	Data   []Person  `json:"data"`
}

// GetCursor returns the value of Cursor.
func (s *PagedPersonListResponse) GetCursor() OptString {
	return s.Cursor
}

// GetData returns the value of Data.
func (s *PagedPersonListResponse) GetData() []Person {
	return s.Data
}

// SetCursor sets the value of Cursor.
func (s *PagedPersonListResponse) SetCursor(val OptString) {
	s.Cursor = val
}

// SetData sets the value of Data.
func (s *PagedPersonListResponse) SetData(val []Person) {
	s.Data = val
}

// Ref: #/components/schemas/Person
type Person struct {
	ID                 string            `json:"id"`
	GismoID            OptString         `json:"gismo_id"`
	Active             bool              `json:"active"`
	DateCreated        time.Time         `json:"date_created"`
	DateUpdated        time.Time         `json:"date_updated"`
	FullName           OptString         `json:"full_name"`
	FirstName          OptString         `json:"first_name"`
	LastName           OptString         `json:"last_name"`
	Email              OptString         `json:"email"`
	Orcid              OptString         `json:"orcid"`
	OrcidToken         OptString         `json:"orcid_token"`
	PreferredFirstName OptString         `json:"preferred_first_name"`
	PreferredLastName  OptString         `json:"preferred_last_name"`
	BirthDate          OptString         `json:"birth_date"`
	Title              OptString         `json:"title"`
	OtherID            []IdRef           `json:"other_id"`
	Organization       []OrganizationRef `json:"organization"`
	JobCategory        []string          `json:"job_category"`
	Role               []string          `json:"role"`
	Settings           OptPersonSettings `json:"settings"`
	ObjectClass        []string          `json:"object_class"`
	ExpirationDate     OptString         `json:"expiration_date"`
}

// GetID returns the value of ID.
func (s *Person) GetID() string {
	return s.ID
}

// GetGismoID returns the value of GismoID.
func (s *Person) GetGismoID() OptString {
	return s.GismoID
}

// GetActive returns the value of Active.
func (s *Person) GetActive() bool {
	return s.Active
}

// GetDateCreated returns the value of DateCreated.
func (s *Person) GetDateCreated() time.Time {
	return s.DateCreated
}

// GetDateUpdated returns the value of DateUpdated.
func (s *Person) GetDateUpdated() time.Time {
	return s.DateUpdated
}

// GetFullName returns the value of FullName.
func (s *Person) GetFullName() OptString {
	return s.FullName
}

// GetFirstName returns the value of FirstName.
func (s *Person) GetFirstName() OptString {
	return s.FirstName
}

// GetLastName returns the value of LastName.
func (s *Person) GetLastName() OptString {
	return s.LastName
}

// GetEmail returns the value of Email.
func (s *Person) GetEmail() OptString {
	return s.Email
}

// GetOrcid returns the value of Orcid.
func (s *Person) GetOrcid() OptString {
	return s.Orcid
}

// GetOrcidToken returns the value of OrcidToken.
func (s *Person) GetOrcidToken() OptString {
	return s.OrcidToken
}

// GetPreferredFirstName returns the value of PreferredFirstName.
func (s *Person) GetPreferredFirstName() OptString {
	return s.PreferredFirstName
}

// GetPreferredLastName returns the value of PreferredLastName.
func (s *Person) GetPreferredLastName() OptString {
	return s.PreferredLastName
}

// GetBirthDate returns the value of BirthDate.
func (s *Person) GetBirthDate() OptString {
	return s.BirthDate
}

// GetTitle returns the value of Title.
func (s *Person) GetTitle() OptString {
	return s.Title
}

// GetOtherID returns the value of OtherID.
func (s *Person) GetOtherID() []IdRef {
	return s.OtherID
}

// GetOrganization returns the value of Organization.
func (s *Person) GetOrganization() []OrganizationRef {
	return s.Organization
}

// GetJobCategory returns the value of JobCategory.
func (s *Person) GetJobCategory() []string {
	return s.JobCategory
}

// GetRole returns the value of Role.
func (s *Person) GetRole() []string {
	return s.Role
}

// GetSettings returns the value of Settings.
func (s *Person) GetSettings() OptPersonSettings {
	return s.Settings
}

// GetObjectClass returns the value of ObjectClass.
func (s *Person) GetObjectClass() []string {
	return s.ObjectClass
}

// GetExpirationDate returns the value of ExpirationDate.
func (s *Person) GetExpirationDate() OptString {
	return s.ExpirationDate
}

// SetID sets the value of ID.
func (s *Person) SetID(val string) {
	s.ID = val
}

// SetGismoID sets the value of GismoID.
func (s *Person) SetGismoID(val OptString) {
	s.GismoID = val
}

// SetActive sets the value of Active.
func (s *Person) SetActive(val bool) {
	s.Active = val
}

// SetDateCreated sets the value of DateCreated.
func (s *Person) SetDateCreated(val time.Time) {
	s.DateCreated = val
}

// SetDateUpdated sets the value of DateUpdated.
func (s *Person) SetDateUpdated(val time.Time) {
	s.DateUpdated = val
}

// SetFullName sets the value of FullName.
func (s *Person) SetFullName(val OptString) {
	s.FullName = val
}

// SetFirstName sets the value of FirstName.
func (s *Person) SetFirstName(val OptString) {
	s.FirstName = val
}

// SetLastName sets the value of LastName.
func (s *Person) SetLastName(val OptString) {
	s.LastName = val
}

// SetEmail sets the value of Email.
func (s *Person) SetEmail(val OptString) {
	s.Email = val
}

// SetOrcid sets the value of Orcid.
func (s *Person) SetOrcid(val OptString) {
	s.Orcid = val
}

// SetOrcidToken sets the value of OrcidToken.
func (s *Person) SetOrcidToken(val OptString) {
	s.OrcidToken = val
}

// SetPreferredFirstName sets the value of PreferredFirstName.
func (s *Person) SetPreferredFirstName(val OptString) {
	s.PreferredFirstName = val
}

// SetPreferredLastName sets the value of PreferredLastName.
func (s *Person) SetPreferredLastName(val OptString) {
	s.PreferredLastName = val
}

// SetBirthDate sets the value of BirthDate.
func (s *Person) SetBirthDate(val OptString) {
	s.BirthDate = val
}

// SetTitle sets the value of Title.
func (s *Person) SetTitle(val OptString) {
	s.Title = val
}

// SetOtherID sets the value of OtherID.
func (s *Person) SetOtherID(val []IdRef) {
	s.OtherID = val
}

// SetOrganization sets the value of Organization.
func (s *Person) SetOrganization(val []OrganizationRef) {
	s.Organization = val
}

// SetJobCategory sets the value of JobCategory.
func (s *Person) SetJobCategory(val []string) {
	s.JobCategory = val
}

// SetRole sets the value of Role.
func (s *Person) SetRole(val []string) {
	s.Role = val
}

// SetSettings sets the value of Settings.
func (s *Person) SetSettings(val OptPersonSettings) {
	s.Settings = val
}

// SetObjectClass sets the value of ObjectClass.
func (s *Person) SetObjectClass(val []string) {
	s.ObjectClass = val
}

// SetExpirationDate sets the value of ExpirationDate.
func (s *Person) SetExpirationDate(val OptString) {
	s.ExpirationDate = val
}

type PersonSettings map[string]string

func (s *PersonSettings) init() PersonSettings {
	m := *s
	if m == nil {
		m = map[string]string{}
		*s = m
	}
	return m
}

// SetPersonOrcidOK is response for SetPersonOrcid operation.
type SetPersonOrcidOK struct{}

// Ref: #/components/schemas/SetPersonOrcidRequest
type SetPersonOrcidRequest struct {
	Orcid string `json:"orcid"`
}

// GetOrcid returns the value of Orcid.
func (s *SetPersonOrcidRequest) GetOrcid() string {
	return s.Orcid
}

// SetOrcid sets the value of Orcid.
func (s *SetPersonOrcidRequest) SetOrcid(val string) {
	s.Orcid = val
}

// SetPersonOrcidTokenOK is response for SetPersonOrcidToken operation.
type SetPersonOrcidTokenOK struct{}

// Ref: #/components/schemas/SetPersonOrcidTokenRequest
type SetPersonOrcidTokenRequest struct {
	OrcidToken string `json:"orcid_token"`
}

// GetOrcidToken returns the value of OrcidToken.
func (s *SetPersonOrcidTokenRequest) GetOrcidToken() string {
	return s.OrcidToken
}

// SetOrcidToken sets the value of OrcidToken.
func (s *SetPersonOrcidTokenRequest) SetOrcidToken(val string) {
	s.OrcidToken = val
}

// SetPersonRoleOK is response for SetPersonRole operation.
type SetPersonRoleOK struct{}

// Ref: #/components/schemas/SetPersonRoleRequest
type SetPersonRoleRequest struct {
	Role []string `json:"role"`
}

// GetRole returns the value of Role.
func (s *SetPersonRoleRequest) GetRole() []string {
	return s.Role
}

// SetRole sets the value of Role.
func (s *SetPersonRoleRequest) SetRole(val []string) {
	s.Role = val
}

// SetPersonSettingsOK is response for SetPersonSettings operation.
type SetPersonSettingsOK struct{}

// Ref: #/components/schemas/SetPersonSettingsRequest
type SetPersonSettingsRequest struct {
	Settings SetPersonSettingsRequestSettings `json:"settings"`
}

// GetSettings returns the value of Settings.
func (s *SetPersonSettingsRequest) GetSettings() SetPersonSettingsRequestSettings {
	return s.Settings
}

// SetSettings sets the value of Settings.
func (s *SetPersonSettingsRequest) SetSettings(val SetPersonSettingsRequestSettings) {
	s.Settings = val
}

type SetPersonSettingsRequestSettings map[string]string

func (s *SetPersonSettingsRequestSettings) init() SetPersonSettingsRequestSettings {
	m := *s
	if m == nil {
		m = map[string]string{}
		*s = m
	}
	return m
}
