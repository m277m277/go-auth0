package management

import (
	"context"
	"encoding/json"
	"time"
)

// SelfServiceProfile allows to configure SSO so that
// customers can independently set up
// SSO and sign in to your application.
type SelfServiceProfile struct {
	ID *string `json:"id,omitempty"`

	// The name of the self-service Profile
	Name *string `json:"name,omitempty"`
	// The description of the self-service Profile.
	Description *string `json:"description,omitempty"`

	// List of IdP strategies that will be shown to users during the Self-Service SSO flow.
	// Possible values: [oidc, samlp, waad, google-apps, adfs, okta, keycloak-samlp]
	AllowedStrategies *[]string `json:"allowed_strategies,omitempty"`

	// List of attributes to be mapped that
	// will be shown to the user during the SS-SSO flow.
	UserAttributes []*SelfServiceProfileUserAttributes `json:"user_attributes,omitempty"`
	CreatedAt      *time.Time                          `json:"created_at,omitempty"`
	UpdatedAt      *time.Time                          `json:"updated_at,omitempty"`

	// Branding scheme for the profile.
	Branding *Branding `json:"branding,omitempty"`
}

// SelfServiceProfileUserAttributes is used to determine optional attributes.
type SelfServiceProfileUserAttributes struct {
	// Identifier of this attribute.
	Name *string `json:"name"`

	// Description of this attribute.
	Description *string `json:"description"`

	// Determines if this attribute is required.
	IsOptional *bool `json:"is_optional"`
}

// SelfServiceProfileList is a list of SelfServiceProfiles.
type SelfServiceProfileList struct {
	List
	SelfServiceProfile []*SelfServiceProfile `json:"self_service_profiles"`
}

// SelfServiceProfileTicket is used to created self-service ticket for a set of clients and organizations.
type SelfServiceProfileTicket struct {
	// If provided, this will allow editing of the
	// provided connection during the SSO Flow.
	ConnectionID *string `json:"connection_id,omitempty"`

	// If provided, this will create a new connection
	// for the SSO flow with the given configuration.
	ConnectionConfig *SelfServiceProfileTicketConnectionConfig `json:"connection_config,omitempty"`

	// List of client_ids that the
	// connection will be enabled for.
	EnabledClients *[]string `json:"enabled_clients,omitempty"`

	// List of organizations that the
	// connection will be enabled for.
	EnabledOrganizations []*SelfServiceProfileTicketEnabledOrganizations `json:"enabled_organizations,omitempty"`

	TTLSec int `json:"ttl_sec,omitempty"`

	// The ticket that is generated.
	Ticket *string `json:"ticket,omitempty"`
}

// SelfServiceProfileTicketConnectionConfig sets the configuration for SSOTicket.
type SelfServiceProfileTicketConnectionConfig struct {
	// The name of the connection that will be
	// created as a part of the SSO flow.
	Name *string `json:"name,omitempty"`

	// The display name of the connection that will be
	// created as a part of the SSO flow.
	DisplayName *string `json:"display_name,omitempty"`

	IsDomainConnection *bool                                            `json:"is_domain_connection,omitempty"`
	ShowAsButton       *bool                                            `json:"show_as_button,omitempty"`
	Metadata           *map[string]interface{}                          `json:"metadata,omitempty"`
	Options            *SelfServiceProfileTicketConnectionConfigOptions `json:"options,omitempty"`
}

// SelfServiceProfileTicketConnectionConfigOptions is the list of Options for SSO Ticket.
type SelfServiceProfileTicketConnectionConfigOptions struct {
	IconURL       *string   `json:"icon_url,omitempty"`
	DomainAliases *[]string `json:"domain_aliases,omitempty"`
}

// SelfServiceProfileTicketEnabledOrganizations is the list of Organizations associated with the SSO Ticket.
type SelfServiceProfileTicketEnabledOrganizations struct {
	// Organization identifier.
	OrganizationID          *string `json:"organization_id,omitempty"`
	AssignMembershipOnLogin *bool   `json:"assign_membership_on_login,omitempty"`
	ShowAsButton            *bool   `json:"show_as_button,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface.
func (ssp *SelfServiceProfile) MarshalJSON() ([]byte, error) {
	type SelfServiceProfileSubset struct {
		Name              *string                             `json:"name,omitempty"`
		Description       *string                             `json:"description,omitempty"`
		AllowedStrategies *[]string                           `json:"allowed_strategies,omitempty"`
		UserAttributes    []*SelfServiceProfileUserAttributes `json:"user_attributes,omitempty"`
		Branding          *Branding                           `json:"branding,omitempty"`
	}

	return json.Marshal(&SelfServiceProfileSubset{
		Name:              ssp.Name,
		Description:       ssp.Description,
		AllowedStrategies: ssp.AllowedStrategies,
		UserAttributes:    ssp.UserAttributes,
		Branding:          ssp.Branding,
	})
}

// SelfServiceProfileManager manages Auth0 Self Service Profile resources.
type SelfServiceProfileManager manager

// Create a new Self Service Profile.
func (m *SelfServiceProfileManager) Create(ctx context.Context, s *SelfServiceProfile, opts ...RequestOption) (err error) {
	err = m.management.Request(ctx, "POST", m.management.URI("self-service-profiles"), &s, opts...)
	return
}

// List all Self Service Profiles.
func (m *SelfServiceProfileManager) List(ctx context.Context, opts ...RequestOption) (s *SelfServiceProfileList, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("self-service-profiles"), &s, applyListDefaults(opts))
	return
}

// Get a single Self Service Profile against the ID.
func (m *SelfServiceProfileManager) Read(ctx context.Context, id string, opts ...RequestOption) (s *SelfServiceProfile, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("self-service-profiles", id), &s, opts...)
	return
}

// Update an existing Self Service Profile against the ID.
func (m *SelfServiceProfileManager) Update(ctx context.Context, id string, s *SelfServiceProfile, opts ...RequestOption) (err error) {
	return m.management.Request(ctx, "PATCH", m.management.URI("self-service-profiles", id), s, opts...)
}

// Delete a Self Service Profile against the ID.
func (m *SelfServiceProfileManager) Delete(ctx context.Context, id string, opts ...RequestOption) (err error) {
	return m.management.Request(ctx, "DELETE", m.management.URI("self-service-profiles", id), nil, opts...)
}

// GetCustomText retrieves text customizations for a given self-service profile, language and Self Service SSO Flow page.
func (m *SelfServiceProfileManager) GetCustomText(ctx context.Context, id string, language string, page string, opts ...RequestOption) (payload map[string]interface{}, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("self-service-profiles", id, "custom-text", language, page), &payload, opts...)
	return
}

// SetCustomText updates text customizations for a given self-service profile, language and Self Service SSO Flow page.
func (m *SelfServiceProfileManager) SetCustomText(ctx context.Context, id string, language string, page string, payload map[string]interface{}, opts ...RequestOption) (err error) {
	err = m.management.Request(ctx, "PUT", m.management.URI("self-service-profiles", id, "custom-text", language, page), payload, opts...)
	return
}

// CreateTicket creates a sso-access ticket to initiate the Self Service SSO Flow.
func (m *SelfServiceProfileManager) CreateTicket(ctx context.Context, id string, t *SelfServiceProfileTicket, opts ...RequestOption) (err error) {
	err = m.management.Request(ctx, "POST", m.management.URI("self-service-profiles", id, "sso-ticket"), t, opts...)
	return
}

// RevokeTicket revokes the sso-access ticket against a specific SSO Profile.
func (m *SelfServiceProfileManager) RevokeTicket(ctx context.Context, id string, ticketID string, opts ...RequestOption) (err error) {
	err = m.management.Request(ctx, "POST", m.management.URI("self-service-profiles", id, "sso-ticket", ticketID, "revoke"), nil, opts...)
	return
}
