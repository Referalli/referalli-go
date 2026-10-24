package referalli

// Code generated from the Referalli OpenAPI document. DO NOT EDIT.
// Changes belong in the generator, which overwrites this file.

import "encoding/json"

var _ json.RawMessage

// Ok is generated from the OpenAPI component of the same name.
type Ok struct {
	Ok bool `json:"ok"`
}

// OkStatus is generated from the OpenAPI component of the same name.
type OkStatus struct {
	Ok     bool   `json:"ok"`
	Status string `json:"status"`
}

// AmountsValueItem is generated from the OpenAPI component of the same name.
type AmountsValueItem struct {
	Currency    string `json:"currency"`
	AmountMinor int64  `json:"amountMinor"` // Integer minor units (cents).
}

// Amounts is generated from the OpenAPI component of the same name.
type Amounts = []AmountsValueItem

// Balance is generated from the OpenAPI component of the same name.
type Balance struct {
	Currency string `json:"currency"`
	Pending  int64  `json:"pending"`  // Integer minor units (cents).
	Approved int64  `json:"approved"` // Integer minor units (cents).
	Paid     int64  `json:"paid"`     // Integer minor units (cents).
	Reversed int64  `json:"reversed"` // Integer minor units (cents).
}

// ProgramSettingsRiskPolicy is generated from the OpenAPI component of the same name.
type ProgramSettingsRiskPolicy struct {
	HoldBotScore        float64 `json:"holdBotScore"`
	HoldDisposableEmail bool    `json:"holdDisposableEmail"`
	HoldVelocity        int64   `json:"holdVelocity"`
	HoldUnverifiedEmail bool    `json:"holdUnverifiedEmail"`
	HoldSelfReferral    bool    `json:"holdSelfReferral"`
	HoldRiskScore       float64 `json:"holdRiskScore"`
	HoldHighRiskLevel   bool    `json:"holdHighRiskLevel"`
}

// ProgramSettingsWidget is generated from the OpenAPI component of the same name.
type ProgramSettingsWidget struct {
	AutoEnroll     bool    `json:"autoEnroll"`
	DestinationURL *string `json:"destinationUrl"`
}

// ProgramSettingsMobileAppOption1IosOption1 is generated from the OpenAPI component of the same name.
type ProgramSettingsMobileAppOption1IosOption1 struct {
	TeamID   string `json:"teamId"`
	BundleID string `json:"bundleId"`
}

// ProgramSettingsMobileAppOption1AndroidOption1 is generated from the OpenAPI component of the same name.
type ProgramSettingsMobileAppOption1AndroidOption1 struct {
	PackageName            string   `json:"packageName"`
	Sha256CertFingerprints []string `json:"sha256CertFingerprints"`
}

// ProgramSettingsMobileAppOption1 is generated from the OpenAPI component of the same name.
type ProgramSettingsMobileAppOption1 struct {
	Ios     *ProgramSettingsMobileAppOption1IosOption1     `json:"ios,omitempty"`
	Android *ProgramSettingsMobileAppOption1AndroidOption1 `json:"android,omitempty"`
}

// ProgramSettings is generated from the OpenAPI component of the same name.
type ProgramSettings struct {
	ProbabilisticPolicy string                           `json:"probabilisticPolicy"`
	SelfReferralPolicy  string                           `json:"selfReferralPolicy"`
	SplitCommissions    bool                             `json:"splitCommissions"`
	UTMTemplate         map[string]string                `json:"utmTemplate"`
	AllowPaidSearch     bool                             `json:"allowPaidSearch"`
	ChallengeLinks      bool                             `json:"challengeLinks"`
	Settlement          string                           `json:"settlement"`
	RiskPolicy          ProgramSettingsRiskPolicy        `json:"riskPolicy"`
	Widget              ProgramSettingsWidget            `json:"widget"`
	RefereeReward       *GetOpenapiJsonResponse          `json:"refereeReward,omitempty"`
	MobileApp           *ProgramSettingsMobileAppOption1 `json:"mobileApp"` // The mobile apps that open this program's links (universal links / App Links), or null.
}

// ProgramPublicPage is generated from the OpenAPI component of the same name.
type ProgramPublicPage struct {
	Pitch       string `json:"pitch"`
	AutoApprove bool   `json:"autoApprove"`
}

// Program is generated from the OpenAPI component of the same name.
type Program struct {
	ID                     string            `json:"id"`
	Name                   string            `json:"name"`
	Slug                   string            `json:"slug"`
	AttributionModel       string            `json:"attributionModel"`
	CookieWindowDays       int64             `json:"cookieWindowDays"`
	SignupToConversionDays int64             `json:"signupToConversionDays"`
	TermsMd                string            `json:"termsMd"`
	Status                 string            `json:"status"`
	PublicHandle           *string           `json:"publicHandle"`
	Settings               ProgramSettings   `json:"settings"`
	PublicPage             ProgramPublicPage `json:"publicPage"`
	CreatedAt              string            `json:"createdAt"`
}

// Partner is generated from the OpenAPI component of the same name.
type Partner struct {
	ID          string   `json:"id"`
	Email       string   `json:"email"`
	DisplayName string   `json:"displayName"`
	ExternalRef *string  `json:"externalRef"`
	Status      string   `json:"status"`
	Tags        []string `json:"tags"`
	Notes       *string  `json:"notes"`
	CreatedAt   string   `json:"createdAt"`
}

// NetworkWarning is generated from the OpenAPI component of the same name.
type NetworkWarning struct {
	Flagged                bool    `json:"flagged"`
	Reason                 *string `json:"reason"`
	BannedByOtherMerchants int64   `json:"bannedByOtherMerchants"` // Other merchants that banned this account; 0 below the threshold (2). Never says which.
}

// Enrollment is generated from the OpenAPI component of the same name.
type Enrollment struct {
	ID        string `json:"id"`
	ProgramID string `json:"programId"`
	PartnerID string `json:"partnerId"`
	Status    string `json:"status"`
}

// Link is generated from the OpenAPI component of the same name.
type Link struct {
	ID             string `json:"id"`
	ProgramID      string `json:"programId"`
	PartnerID      string `json:"partnerId"`
	Code           string `json:"code"` // Lowercase; matched case-insensitively.
	DestinationURL string `json:"destinationUrl"`
	ShortURL       string `json:"shortUrl"` // Absolute link, on the org's verified link domain when it has one.
	Clicks         int64  `json:"clicks"`
	CreatedAt      string `json:"createdAt"`
}

// CreatedLink is generated from the OpenAPI component of the same name.
type CreatedLink struct {
	ID             string  `json:"id"`
	ProgramID      string  `json:"programId"`
	PartnerID      string  `json:"partnerId"`
	Code           string  `json:"code"`
	DestinationURL string  `json:"destinationUrl"`
	ShortURL       *string `json:"shortUrl,omitempty"`
}

// Coupon is generated from the OpenAPI component of the same name.
type Coupon struct {
	ID        string `json:"id"`
	ProgramID string `json:"programId"`
	PartnerID string `json:"partnerId"`
	Code      string `json:"code"`
	CreatedAt string `json:"createdAt"`
}

// ReasonNote is generated from the OpenAPI component of the same name.
type ReasonNote struct {
	Code   string             `json:"code"`
	Values map[string]float64 `json:"values,omitempty"`
}

// PayoutScheduleValueOption1 is generated from the OpenAPI component of the same name.
type PayoutScheduleValueOption1 struct {
	Kind       string `json:"kind"`
	DayOfMonth int64  `json:"dayOfMonth"`
}

// PayoutScheduleValueOption2 is generated from the OpenAPI component of the same name.
type PayoutScheduleValueOption2 struct {
	Kind      string `json:"kind"`
	DayOfWeek int64  `json:"dayOfWeek"` // 0 = Sunday … 6 = Saturday
}

// PayoutScheduleValueOption3 is generated from the OpenAPI component of the same name.
type PayoutScheduleValueOption3 struct {
	Kind string  `json:"kind"`
	Days []int64 `json:"days"`
}

// PayoutScheduleValueOption4 is generated from the OpenAPI component of the same name.
type PayoutScheduleValueOption4 struct {
	Kind  string   `json:"kind"`
	Dates []string `json:"dates"`
}

// PayoutSchedule is generated from the OpenAPI component of the same name.
type PayoutSchedule = json.RawMessage

// AttributionReasoning is generated from the OpenAPI component of the same name.
type AttributionReasoning struct {
	Model                 *string      `json:"model,omitempty"`
	ConsideredTouchpoints *int64       `json:"consideredTouchpoints,omitempty"`
	WindowDays            *int64       `json:"windowDays,omitempty"`
	Notes                 []string     `json:"notes,omitempty"`
	NoteCodes             []ReasonNote `json:"noteCodes,omitempty"`
}

// ReferralTouchCreditsItem is generated from the OpenAPI component of the same name.
type ReferralTouchCreditsItem struct {
	PartnerID string  `json:"partnerId"`
	Weight    float64 `json:"weight"`
	ClickID   *string `json:"clickId"`
}

// Referral is generated from the OpenAPI component of the same name.
type Referral struct {
	ID              string                     `json:"id"`
	ProgramID       string                     `json:"programId"`
	PartnerID       string                     `json:"partnerId"`
	CustomerRef     string                     `json:"customerRef"`
	Source          string                     `json:"source"`
	MatchedVia      string                     `json:"matchedVia"`
	Confidence      float64                    `json:"confidence"`
	ClickID         *string                    `json:"clickId"`
	LinkID          *string                    `json:"linkId"`
	Status          string                     `json:"status"`
	Ts              string                     `json:"ts"`
	WindowExpiresAt string                     `json:"windowExpiresAt"`
	TouchCredits    []ReferralTouchCreditsItem `json:"touchCredits"`
	Reasoning       AttributionReasoning       `json:"reasoning"`
}

// SignupReasonCode is generated from the OpenAPI component of the same name.
type SignupReasonCode = string

// AssociateResponseReferralOption1TouchCreditsItem is generated from the OpenAPI component of the same name.
type AssociateResponseReferralOption1TouchCreditsItem struct {
	PartnerID string  `json:"partnerId"`
	Weight    float64 `json:"weight"`
	ClickID   *string `json:"clickId"`
}

// AssociateResponseReferralOption1 is generated from the OpenAPI component of the same name.
type AssociateResponseReferralOption1 struct {
	ID              string                                             `json:"id"`
	ProgramID       string                                             `json:"programId"`
	PartnerID       string                                             `json:"partnerId"`
	CustomerRef     string                                             `json:"customerRef"`
	Source          string                                             `json:"source"`
	MatchedVia      string                                             `json:"matchedVia"`
	Confidence      float64                                            `json:"confidence"`
	ClickID         *string                                            `json:"clickId"`
	LinkID          *string                                            `json:"linkId"`
	Status          string                                             `json:"status"`
	Ts              string                                             `json:"ts"`
	WindowExpiresAt string                                             `json:"windowExpiresAt"`
	TouchCredits    []AssociateResponseReferralOption1TouchCreditsItem `json:"touchCredits"`
	Reasoning       AttributionReasoning                               `json:"reasoning"`
	OrgID           string                                             `json:"orgId"`
}

// AssociateResponseReasoning is generated from the OpenAPI component of the same name.
type AssociateResponseReasoning struct {
	Notes []string `json:"notes"`
}

// AssociateResponse is generated from the OpenAPI component of the same name.
type AssociateResponse struct {
	Referral  *AssociateResponseReferralOption1 `json:"referral"` // null when no partner matched (not an error).
	Reasoning *AssociateResponseReasoning       `json:"reasoning,omitempty"`
}

// CommissionEvent is generated from the OpenAPI component of the same name.
type CommissionEvent struct {
	ID                  string  `json:"id"`
	ProgramID           string  `json:"programId"`
	PartnerID           string  `json:"partnerId"`
	ReferralID          *string `json:"referralId"`
	Type                string  `json:"type"`
	AmountMinor         int64   `json:"amountMinor"` // Integer minor units (cents).
	Currency            string  `json:"currency"`
	RelatedCommissionID *string `json:"relatedCommissionId"`
	HoldUntil           *string `json:"holdUntil"`
	Level               *int64  `json:"level"`
	Memo                *string `json:"memo"`
	SettlementMethod    *string `json:"settlementMethod,omitempty"` // Set on paid events the merchant settled itself (store_credit).
	SettlementReference *string `json:"settlementReference,omitempty"`
	Ts                  string  `json:"ts"`
}

// CommissionLine: An earned line also carries its net, state and the events recorded against it.
type CommissionLine struct {
	ID                  string            `json:"id"`
	ProgramID           string            `json:"programId"`
	PartnerID           string            `json:"partnerId"`
	ReferralID          *string           `json:"referralId"`
	Type                string            `json:"type"`
	AmountMinor         int64             `json:"amountMinor"` // Integer minor units (cents).
	Currency            string            `json:"currency"`
	RelatedCommissionID *string           `json:"relatedCommissionId"`
	HoldUntil           *string           `json:"holdUntil"`
	Level               *int64            `json:"level"`
	Memo                *string           `json:"memo"`
	SettlementMethod    *string           `json:"settlementMethod,omitempty"` // Set on paid events the merchant settled itself (store_credit).
	SettlementReference *string           `json:"settlementReference,omitempty"`
	Ts                  string            `json:"ts"`
	BillingEventID      *string           `json:"billingEventId,omitempty"`
	NetMinor            *int64            `json:"netMinor,omitempty"` // Integer minor units (cents).
	State               *string           `json:"state,omitempty"`
	Events              []CommissionEvent `json:"events,omitempty"`
}

// Settlement is generated from the OpenAPI component of the same name.
type Settlement struct {
	CommissionID string `json:"commissionId"`
	PaidEventID  string `json:"paidEventId"`
	PartnerID    string `json:"partnerId"`
	ProgramID    string `json:"programId"`
	AmountMinor  int64  `json:"amountMinor"` // Integer minor units (cents).
	Currency     string `json:"currency"`
	Method       string `json:"method"`
	Reference    string `json:"reference"`
	Status       string `json:"status"`
}

// CommissionRule is generated from the OpenAPI component of the same name.
type CommissionRule struct {
	ID        string                 `json:"id"`
	ProgramID string                 `json:"programId"`
	Version   int64                  `json:"version"`
	Config    GetOpenapiJsonResponse `json:"config"`
	CreatedAt string                 `json:"createdAt"`
}

// BillingFilter is generated from the OpenAPI component of the same name.
type BillingFilter struct {
	Metadata   map[string]string `json:"metadata,omitempty"`
	ProductIds []string          `json:"productIds,omitempty"`
	PriceIds   []string          `json:"priceIds,omitempty"`
	Skus       []string          `json:"skus,omitempty"`
}

// ValueMapping is generated from the OpenAPI component of the same name.
type ValueMapping struct {
	Cash          *string `json:"cash,omitempty"`
	CreditApplied *string `json:"creditApplied,omitempty"`
	Discount      *string `json:"discount,omitempty"`
	Cost          *string `json:"cost,omitempty"`
	ProductRef    *string `json:"productRef,omitempty"`
	Unit          string  `json:"unit"`
	Version       int64   `json:"version"` // Bumped on every save (the built-in extractor is 1).
}

// ValueMappingPreviewValue is generated from the OpenAPI component of the same name.
type ValueMappingPreviewValue struct {
	CashMinor          int64              `json:"cashMinor"`           // Integer minor units (cents).
	CreditAppliedMinor int64              `json:"creditAppliedMinor"`  // Integer minor units (cents).
	DiscountMinor      int64              `json:"discountMinor"`       // Integer minor units (cents).
	CostMinor          *int64             `json:"costMinor,omitempty"` // Integer minor units (cents).
	ProductRef         *string            `json:"productRef,omitempty"`
	Fields             map[string]float64 `json:"fields"`
	Warnings           []string           `json:"warnings,omitempty"` // Mapped paths holding an unreadable value; read as 0.
	Source             *string            `json:"source,omitempty"`
	MappingVersion     *int64             `json:"mappingVersion,omitempty"`
}

// ValueMappingPreviewWouldRecordItem is generated from the OpenAPI component of the same name.
type ValueMappingPreviewWouldRecordItem struct {
	Kind               string  `json:"kind"`
	Currency           string  `json:"currency"`
	Confidence         string  `json:"confidence"`
	Source             string  `json:"source"`
	ExtractionVersion  int64   `json:"extractionVersion"`
	CashMinor          *int64  `json:"cashMinor,omitempty"`          // Integer minor units (cents).
	CreditAppliedMinor *int64  `json:"creditAppliedMinor,omitempty"` // Integer minor units (cents).
	DiscountMinor      *int64  `json:"discountMinor,omitempty"`      // Integer minor units (cents).
	CostMinor          *int64  `json:"costMinor,omitempty"`          // Integer minor units (cents).
	ProductRef         *string `json:"productRef,omitempty"`
}

// ValueMappingPreviewEventsItemValue is generated from the OpenAPI component of the same name.
type ValueMappingPreviewEventsItemValue struct {
	CashMinor          int64              `json:"cashMinor"`           // Integer minor units (cents).
	CreditAppliedMinor int64              `json:"creditAppliedMinor"`  // Integer minor units (cents).
	DiscountMinor      int64              `json:"discountMinor"`       // Integer minor units (cents).
	CostMinor          *int64             `json:"costMinor,omitempty"` // Integer minor units (cents).
	ProductRef         *string            `json:"productRef,omitempty"`
	Fields             map[string]float64 `json:"fields"`
	Warnings           []string           `json:"warnings,omitempty"` // Mapped paths holding an unreadable value; read as 0.
	Source             *string            `json:"source,omitempty"`
	MappingVersion     *int64             `json:"mappingVersion,omitempty"`
}

// ValueMappingPreviewEventsItemWouldRecordItem is generated from the OpenAPI component of the same name.
type ValueMappingPreviewEventsItemWouldRecordItem struct {
	Kind               string  `json:"kind"`
	Currency           string  `json:"currency"`
	Confidence         string  `json:"confidence"`
	Source             string  `json:"source"`
	ExtractionVersion  int64   `json:"extractionVersion"`
	CashMinor          *int64  `json:"cashMinor,omitempty"`          // Integer minor units (cents).
	CreditAppliedMinor *int64  `json:"creditAppliedMinor,omitempty"` // Integer minor units (cents).
	DiscountMinor      *int64  `json:"discountMinor,omitempty"`      // Integer minor units (cents).
	CostMinor          *int64  `json:"costMinor,omitempty"`          // Integer minor units (cents).
	ProductRef         *string `json:"productRef,omitempty"`
}

// ValueMappingPreviewEventsItem is generated from the OpenAPI component of the same name.
type ValueMappingPreviewEventsItem struct {
	Currency    string                                         `json:"currency"`
	Value       ValueMappingPreviewEventsItemValue             `json:"value"`
	WouldRecord []ValueMappingPreviewEventsItemWouldRecordItem `json:"wouldRecord"`
	Index       int64                                          `json:"index"`
}

// ValueMappingPreview is generated from the OpenAPI component of the same name.
type ValueMappingPreview struct {
	MappingVersion int64                                `json:"mappingVersion"`
	Currency       *string                              `json:"currency,omitempty"` // Single-event payloads.
	Value          *ValueMappingPreviewValue            `json:"value,omitempty"`
	WouldRecord    []ValueMappingPreviewWouldRecordItem `json:"wouldRecord,omitempty"`
	Events         []ValueMappingPreviewEventsItem      `json:"events,omitempty"` // A generic-connector body with an events array: one result per element.
}

// BillingConnectionRotation is generated from the OpenAPI component of the same name.
type BillingConnectionRotation struct {
	InProgress              bool    `json:"inProgress"`
	PreviousSecretExpiresAt *string `json:"previousSecretExpiresAt,omitempty"`
}

// BillingConnection is generated from the OpenAPI component of the same name.
type BillingConnection struct {
	ID               string                    `json:"id"`
	Provider         string                    `json:"provider"`
	Name             string                    `json:"name"`
	Status           string                    `json:"status"` // paused: its webhook URL refuses events until resumed.
	CreatedAt        string                    `json:"createdAt"`
	WebhookURL       string                    `json:"webhookUrl"`
	NeedsCredentials bool                      `json:"needsCredentials"` // The provider sends thin webhooks: set API credentials (Recurly, BigCommerce).
	HasCredentials   bool                      `json:"hasCredentials"`
	ValueMapping     *ValueMapping             `json:"valueMapping"` // Where cash / credit / discount / cost live in this connection's payload; null uses the built-in extractor.
	Filter           *BillingFilter            `json:"filter"`       // Only matching events earn; the rest are recorded as ignored_by_filter. null takes every event.
	Rotation         BillingConnectionRotation `json:"rotation"`
}

// CreatedBillingConnection is generated from the OpenAPI component of the same name.
type CreatedBillingConnection struct {
	ID            string         `json:"id"`
	Provider      string         `json:"provider"`
	WebhookSecret string         `json:"webhookSecret"` // Returned only here.
	WebhookURL    string         `json:"webhookUrl"`    // Path to register with the provider.
	Filter        *BillingFilter `json:"filter"`
}

// BillingEvent is generated from the OpenAPI component of the same name.
type BillingEvent struct {
	ID            string              `json:"id"`
	Provider      string              `json:"provider"`
	Type          string              `json:"type"`
	CustomerRef   string              `json:"customerRef"`
	InvoiceRef    *string             `json:"invoiceRef"`
	AmountGross   int64               `json:"amountGross"` // Integer minor units (cents).
	Tax           int64               `json:"tax"`         // Integer minor units (cents).
	Fees          int64               `json:"fees"`        // Integer minor units (cents).
	Currency      string              `json:"currency"`
	OccurredAt    string              `json:"occurredAt"`
	ConnectionID  *string             `json:"connectionId"`
	ProductRef    string              `json:"productRef"`
	Status        string              `json:"status"` // ignored_by_filter: recorded, never processed (product filter or internal customer).
	IgnoredReason *string             `json:"ignoredReason"`
	Catalog       map[string][]string `json:"catalog"` // The productIds, priceIds and skus the event named.
}

// InternalCustomer is generated from the OpenAPI component of the same name.
type InternalCustomer struct {
	CustomerRef string `json:"customerRef"`
	Note        string `json:"note"`
	CreatedAt   string `json:"createdAt"`
}

// IdentityProviderSampleHookOption1 is generated from the OpenAPI component of the same name.
type IdentityProviderSampleHookOption1 struct {
	Language string `json:"language"`
	Filename string `json:"filename"`
	Code     string `json:"code"`
}

// IdentityProviderMetadataPaths is generated from the OpenAPI component of the same name.
type IdentityProviderMetadataPaths struct {
	AnonID  []string `json:"anonId"`
	ClickID []string `json:"clickId"`
	Coupon  []string `json:"coupon"`
}

// IdentityProvider is generated from the OpenAPI component of the same name.
type IdentityProvider struct {
	ID             string                             `json:"id"`
	Label          string                             `json:"label"`
	Mechanism      string                             `json:"mechanism"`
	Status         string                             `json:"status"`
	DocsURL        string                             `json:"docsUrl"`
	VendorDocsURL  string                             `json:"vendorDocsUrl"`
	SecretKind     string                             `json:"secretKind"`
	SecretIssuedBy string                             `json:"secretIssuedBy"`
	VendorSteps    []string                           `json:"vendorSteps"`
	SampleHook     *IdentityProviderSampleHookOption1 `json:"sampleHook"`
	MetadataPaths  IdentityProviderMetadataPaths      `json:"metadataPaths"`
}

// WebhookEndpointRotation is generated from the OpenAPI component of the same name.
type WebhookEndpointRotation struct {
	InProgress              bool    `json:"inProgress"`
	PreviousSecretExpiresAt *string `json:"previousSecretExpiresAt,omitempty"`
}

// WebhookEndpointLastFailureOption1 is generated from the OpenAPI component of the same name.
type WebhookEndpointLastFailureOption1 struct {
	StatusCode *int64  `json:"statusCode"`
	Error      *string `json:"error"`
	At         string  `json:"at"`
}

// WebhookEndpoint is generated from the OpenAPI component of the same name.
type WebhookEndpoint struct {
	ID                  string                             `json:"id"`
	URL                 string                             `json:"url"`
	Topics              []string                           `json:"topics"`
	ProgramIds          []string                           `json:"programIds"` // null: every program. Otherwise only events about these programs (org-wide events such as payouts are not sent).
	Status              string                             `json:"status"`
	LegacySignature     bool                               `json:"legacySignature"` // Also send the deprecated body-only legacy signature header.
	Rotation            WebhookEndpointRotation            `json:"rotation"`
	FailingSince        *string                            `json:"failingSince"` // When the current run of failed deliveries began.
	DisabledAt          *string                            `json:"disabledAt"`
	DisabledReason      *string                            `json:"disabledReason"`
	Health              string                             `json:"health"` // failing: the latest deliveries failed; alerts go out at 1 h and 24 h, and it is disabled at 72 h.
	LastSuccessAt       *string                            `json:"lastSuccessAt"`
	LastFailure         *WebhookEndpointLastFailureOption1 `json:"lastFailure"` // The most recent failed attempt: HTTP status (null: no response) and the error or start of the response body.
	ConsecutiveFailures int64                              `json:"consecutiveFailures"`
	AlertContacts       []string                           `json:"alertContacts"` // Emailed (with the org owner) when the endpoint starts failing.
	CreatedAt           string                             `json:"createdAt"`
}

// HealthCheck is generated from the OpenAPI component of the same name.
type HealthCheck struct {
	ID     string  `json:"id"`
	Ok     bool    `json:"ok"`
	Detail string  `json:"detail"`
	Fix    *string `json:"fix,omitempty"`
	Docs   *string `json:"docs,omitempty"`
	Hop    *string `json:"hop,omitempty"` // The hop that failed on a multi-system path, such as `stripe→referalli`.
}

// WebhookDelivery is generated from the OpenAPI component of the same name.
type WebhookDelivery struct {
	ID             string  `json:"id"`
	EndpointID     string  `json:"endpointId"`
	Topic          string  `json:"topic"`
	Status         string  `json:"status"`
	Attempts       int64   `json:"attempts"`
	LastStatusCode *int64  `json:"lastStatusCode"`
	LastAttemptAt  *string `json:"lastAttemptAt"`
	NextAt         *string `json:"nextAt"` // When a pending delivery is next tried; null once settled.
	CreatedAt      string  `json:"createdAt"`
}

// PayoutProfile is generated from the OpenAPI component of the same name.
type PayoutProfile struct {
	PartnerID          string  `json:"partnerId"`
	Method             string  `json:"method"`
	PayoutCurrency     string  `json:"payoutCurrency"`
	MinThresholdMinor  int64   `json:"minThresholdMinor"` // Integer minor units (cents).
	TaxFormType        *string `json:"taxFormType"`
	TaxFormSubmittedAt *string `json:"taxFormSubmittedAt"`
}

// TaxProfile is generated from the OpenAPI component of the same name.
type TaxProfile struct {
	LegalName       *string `json:"legalName"`
	Address         *string `json:"address"`
	Country         *string `json:"country"`
	VATID           *string `json:"vatId"`
	Ein             *string `json:"ein"`
	DomesticVATBps  *int64  `json:"domesticVatBps"`
	ViesCheck       bool    `json:"viesCheck"`
	FilingProvider  string  `json:"filingProvider"`
	FilingConnected bool    `json:"filingConnected"`
}

// TaxFilingDocumentsItem is generated from the OpenAPI component of the same name.
type TaxFilingDocumentsItem struct {
	ID          string `json:"id"`
	PartnerID   string `json:"partnerId"`
	AmountMinor int64  `json:"amountMinor"` // Integer minor units (cents).
	Status      string `json:"status"`
}

// TaxFiling is generated from the OpenAPI component of the same name.
type TaxFiling struct {
	ID               string                   `json:"id"`
	TaxYear          int64                    `json:"taxYear"`
	Kind             string                   `json:"kind"`
	Provider         string                   `json:"provider"`
	ProviderFilingID *string                  `json:"providerFilingId"`
	Status           string                   `json:"status"`
	RecipientCount   int64                    `json:"recipientCount"`
	TotalMinor       int64                    `json:"totalMinor"` // Integer minor units (cents).
	Error            *string                  `json:"error"`
	Documents        []TaxFilingDocumentsItem `json:"documents"`
}

// NecPreviewRecipientsItem is generated from the OpenAPI component of the same name.
type NecPreviewRecipientsItem struct {
	PartnerID   string   `json:"partnerId"`
	Name        *string  `json:"name"`
	Email       *string  `json:"email"`
	AmountMinor int64    `json:"amountMinor"` // Integer minor units (cents).
	Issues      []string `json:"issues"`
}

// NecPreview is generated from the OpenAPI component of the same name.
type NecPreview struct {
	TaxYear     int64                      `json:"taxYear"`
	PayerReady  bool                       `json:"payerReady"`
	PayerIssues []string                   `json:"payerIssues"`
	Recipients  []NecPreviewRecipientsItem `json:"recipients"`
}

// InvoiceSupplier: The partner. Absent on invoices issued before 2026-10.
type InvoiceSupplier struct {
	Name    *string `json:"name"`
	Address *string `json:"address"`
	Country *string `json:"country"`
	VATID   *string `json:"vatId"`
}

// InvoiceCustomer: The merchant, issuing the invoice on the partner’s behalf.
type InvoiceCustomer struct {
	Name    *string `json:"name"`
	Address *string `json:"address"`
	Country *string `json:"country"`
	VATID   *string `json:"vatId"`
}

// InvoiceSelfBilling is generated from the OpenAPI component of the same name.
type InvoiceSelfBilling struct {
	AgreementVersion *string `json:"agreementVersion"`
	AcceptedAt       *string `json:"acceptedAt"`
}

// Invoice is generated from the OpenAPI component of the same name.
type Invoice struct {
	Number       string              `json:"number"` // Sequential per merchant per year: SBI-2026-1, -2, …
	IssuedAt     string              `json:"issuedAt"`
	Lines        []any               `json:"lines"`
	VATNote      string              `json:"vatNote"`
	Supplier     *InvoiceSupplier    `json:"supplier,omitempty"` // The partner. Absent on invoices issued before 2026-10.
	Customer     *InvoiceCustomer    `json:"customer,omitempty"` // The merchant, issuing the invoice on the partner’s behalf.
	Currency     *string             `json:"currency,omitempty"`
	NetMinor     *int64              `json:"netMinor,omitempty"` // Integer minor units (cents).
	VATTreatment *string             `json:"vatTreatment,omitempty"`
	VATRateBps   *int64              `json:"vatRateBps,omitempty"`
	VATMinor     *int64              `json:"vatMinor,omitempty"`   // Integer minor units (cents).
	GrossMinor   *int64              `json:"grossMinor,omitempty"` // Integer minor units (cents).
	SelfBilling  *InvoiceSelfBilling `json:"selfBilling,omitempty"`
}

// PayoutRunCreatedPayoutsItemFxAppliedItem is generated from the OpenAPI component of the same name.
type PayoutRunCreatedPayoutsItemFxAppliedItem struct {
	From string  `json:"from"`
	To   string  `json:"to"`
	Rate float64 `json:"rate"`
	Ts   string  `json:"ts"`
}

// PayoutRunCreatedPayoutsItem is generated from the OpenAPI component of the same name.
type PayoutRunCreatedPayoutsItem struct {
	PayoutID    string                                     `json:"payoutId"`
	PartnerID   string                                     `json:"partnerId"`
	AmountMinor int64                                      `json:"amountMinor"` // Integer minor units (cents).
	Currency    string                                     `json:"currency"`
	Rail        string                                     `json:"rail"`
	FXApplied   []PayoutRunCreatedPayoutsItemFxAppliedItem `json:"fxApplied"`
	Invoice     Invoice                                    `json:"invoice"`
}

// PayoutRunCreatedSkippedItem is generated from the OpenAPI component of the same name.
type PayoutRunCreatedSkippedItem struct {
	PartnerID string `json:"partnerId"`
	Reason    string `json:"reason"`
}

// PayoutRunCreated is generated from the OpenAPI component of the same name.
type PayoutRunCreated struct {
	RunID   string                        `json:"runId"`
	Payouts []PayoutRunCreatedPayoutsItem `json:"payouts"`
	Skipped []PayoutRunCreatedSkippedItem `json:"skipped"`
}

// PayoutRun is generated from the OpenAPI component of the same name.
type PayoutRun struct {
	ID         string                  `json:"id"`
	Status     string                  `json:"status"`
	Summary    *GetOpenapiJsonResponse `json:"summary"`
	CreatedAt  string                  `json:"createdAt"`
	ExecutedAt *string                 `json:"executedAt"`
}

// Payout is generated from the OpenAPI component of the same name.
type Payout struct {
	ID            string  `json:"id"`
	RunID         string  `json:"runId"`
	PartnerID     string  `json:"partnerId"`
	AmountMinor   int64   `json:"amountMinor"` // Integer minor units (cents).
	Currency      string  `json:"currency"`
	Status        string  `json:"status"`
	Rail          string  `json:"rail"`
	RailRef       *string `json:"railRef"`
	FailureReason *string `json:"failureReason"`
	SentAt        *string `json:"sentAt"`
}

// FxRate is generated from the OpenAPI component of the same name.
type FxRate struct {
	Base        string  `json:"base"`
	Quote       string  `json:"quote"`
	Rate        string  `json:"rate"` // Decimal string, never a float.
	EffectiveAt string  `json:"effectiveAt"`
	Source      string  `json:"source"` // 'manual' (your own), 'ecb' (the daily feed), or 'fixture' (platform sample).
	Scope       string  `json:"scope"`
	RetiredAt   *string `json:"retiredAt"` // When a manual override was cleared, handing the pair back to the feed.
}

// Org is generated from the OpenAPI component of the same name.
type Org struct {
	ID                string                 `json:"id"`
	Name              string                 `json:"name"`
	Plan              string                 `json:"plan"`
	Settings          GetOpenapiJsonResponse `json:"settings"`
	TrackingToken     string                 `json:"trackingToken"`
	PayoutKYCRequired bool                   `json:"payoutKycRequired"`
	Livemode          bool                   `json:"livemode"`            // False for the org's test org (test mode).
	LiveOrgID         *string                `json:"liveOrgId,omitempty"` // On a test org: the live org it belongs to.
}

// Member is generated from the OpenAPI component of the same name.
type Member struct {
	UserID string  `json:"userId"`
	Role   string  `json:"role"`
	Email  string  `json:"email"`
	Name   *string `json:"name"`
}

// IssuedApiKey is generated from the OpenAPI component of the same name.
type IssuedApiKey struct {
	ID             string   `json:"id"`
	Key            string   `json:"key"` // The full key, shown once.
	Scopes         []string `json:"scopes"`
	Publishable    bool     `json:"publishable"`
	AllowedOrigins []string `json:"allowedOrigins"`
	APIVersion     string   `json:"apiVersion"`
	Livemode       bool     `json:"livemode"` // False for a test-mode key (rk_test_…).
	Owner          *string  `json:"owner"`
	ExpiresAt      *string  `json:"expiresAt"`
	PairID         *string  `json:"pairId"`
}

// ApiKeyCreatedByOption1 is generated from the OpenAPI component of the same name.
type ApiKeyCreatedByOption1 struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

// ApiKey is generated from the OpenAPI component of the same name.
type ApiKey struct {
	ID             string                  `json:"id"`
	Name           string                  `json:"name"`
	Prefix         string                  `json:"prefix"`
	Scopes         []string                `json:"scopes"`
	APIVersion     string                  `json:"apiVersion"` // The API version this key is pinned to (what it gets without a Referalli-Version header).
	Publishable    bool                    `json:"publishable"`
	AllowedOrigins []string                `json:"allowedOrigins"`
	Livemode       bool                    `json:"livemode"`
	Owner          *string                 `json:"owner"`  // The developer or service the key belongs to.
	PairID         *string                 `json:"pairId"` // Shared by a live/test pair minted together.
	CreatedBy      *ApiKeyCreatedByOption1 `json:"createdBy"`
	ExpiresAt      *string                 `json:"expiresAt"`
	Expired        bool                    `json:"expired"`
	LastUsedAt     *string                 `json:"lastUsedAt"` // Updated at most once a minute.
	CreatedAt      string                  `json:"createdAt"`
	RevokedAt      *string                 `json:"revokedAt"`
}

// AuthConnectionRotation is generated from the OpenAPI component of the same name.
type AuthConnectionRotation struct {
	InProgress              bool    `json:"inProgress"`
	PreviousSecretExpiresAt *string `json:"previousSecretExpiresAt,omitempty"`
}

// AuthConnection is generated from the OpenAPI component of the same name.
type AuthConnection struct {
	ID               string                 `json:"id"`
	Provider         string                 `json:"provider"`
	ProgramID        string                 `json:"programId"`
	Name             string                 `json:"name"`
	Status           string                 `json:"status"`
	CreatedAt        string                 `json:"createdAt"`
	Rotation         AuthConnectionRotation `json:"rotation"`
	InstanceID       *string                `json:"instanceId,omitempty"` // Signed-protocol connections: the Atlas instance, or a custom label.
	AutoEnrollUsers  *bool                  `json:"autoEnrollUsers,omitempty"`
	DestinationURL   *string                `json:"destinationUrl,omitempty"`
	SignatureVersion *string                `json:"signatureVersion,omitempty"`
	RequiresV2       *bool                  `json:"requiresV2,omitempty"` // True once signatureVersion is v2 or the connection has seen a v2 call.
}

// CreatedAuthConnection is generated from the OpenAPI component of the same name.
type CreatedAuthConnection struct {
	ID              string            `json:"id"`
	Provider        string            `json:"provider"`
	Secret          string            `json:"secret"`               // Returned only here.
	WebhookURL      *string           `json:"webhookUrl,omitempty"` // Webhook connectors: Clerk, Auth0, Supabase, Firebase, Cognito, Stytch, Kinde, WorkOS, generic.
	DocsURL         *string           `json:"docsUrl,omitempty"`    // Webhook connectors: the setup guide for this provider.
	MetadataKeys    map[string]string `json:"metadataKeys,omitempty"`
	ProgramID       *string           `json:"programId,omitempty"` // Signed-protocol connections (atlas, custom).
	InstanceID      *string           `json:"instanceId,omitempty"`
	AutoEnrollUsers *bool             `json:"autoEnrollUsers,omitempty"`
	DestinationURL  *string           `json:"destinationUrl,omitempty"`
	SignupsURL      *string           `json:"signupsUrl,omitempty"`
	EraseURL        *string           `json:"eraseUrl,omitempty"`
	ExportURL       *string           `json:"exportUrl,omitempty"`
}

// FraudCase is generated from the OpenAPI component of the same name.
type FraudCase struct {
	ID         string                  `json:"id"`
	PartnerID  *string                 `json:"partnerId"`
	ReferralID *string                 `json:"referralId"`
	Kind       string                  `json:"kind"`
	Status     string                  `json:"status"`
	Evidence   *GetOpenapiJsonResponse `json:"evidence"`
	Action     *string                 `json:"action"`
	CreatedAt  string                  `json:"createdAt"`
	ResolvedAt *string                 `json:"resolvedAt"`
}

// CustomerExport is generated from the OpenAPI component of the same name.
type CustomerExport struct {
	CustomerRef      string                   `json:"customerRef"`
	Referrals        []GetOpenapiJsonResponse `json:"referrals"`
	Associations     []GetOpenapiJsonResponse `json:"associations"`
	IdentityLinks    []GetOpenapiJsonResponse `json:"identityLinks"`
	BillingEvents    []GetOpenapiJsonResponse `json:"billingEvents"`
	ConversionEvents []GetOpenapiJsonResponse `json:"conversionEvents"`
}

// PartnerExport is generated from the OpenAPI component of the same name.
type PartnerExport struct {
	Partner          GetOpenapiJsonResponse   `json:"partner"`
	Enrollments      []GetOpenapiJsonResponse `json:"enrollments"`
	Links            []GetOpenapiJsonResponse `json:"links"`
	Coupons          []GetOpenapiJsonResponse `json:"coupons"`
	PayoutProfile    *GetOpenapiJsonResponse  `json:"payoutProfile"`
	CommissionEvents []GetOpenapiJsonResponse `json:"commissionEvents"`
	Payouts          []GetOpenapiJsonResponse `json:"payouts"`
}

// PrivacyRequest is generated from the OpenAPI component of the same name.
type PrivacyRequest struct {
	ID          string                  `json:"id"`
	Type        string                  `json:"type"`
	Subject     GetOpenapiJsonResponse  `json:"subject"`
	Status      string                  `json:"status"`
	Error       *string                 `json:"error"`
	CreatedAt   string                  `json:"createdAt"`
	CompletedAt *string                 `json:"completedAt"`
	ResultURL   *string                 `json:"resultUrl,omitempty"`
	Result      *GetOpenapiJsonResponse `json:"result,omitempty"` // Erasure receipts, on the single-request read.
}

// GraphReferral is generated from the OpenAPI component of the same name.
type GraphReferral struct {
	ReferralID  string        `json:"referralId"`
	CustomerRef string        `json:"customerRef"`
	Status      string        `json:"status"`
	MatchedVia  string        `json:"matchedVia"`
	Ts          string        `json:"ts"`
	Revenue     Amounts       `json:"revenue"`
	Commission  Amounts       `json:"commission"`
	Partner     *GraphPartner `json:"partner"`
}

// GraphPartner is generated from the OpenAPI component of the same name.
type GraphPartner struct {
	PartnerID   string          `json:"partnerId"`
	DisplayName string          `json:"displayName"`
	Email       string          `json:"email"`
	ExternalRef *string         `json:"externalRef"`
	Status      string          `json:"status"`
	Commission  Amounts         `json:"commission"`
	Referrals   []GraphReferral `json:"referrals"`
}

// CreatorProfileStats is generated from the OpenAPI component of the same name.
type CreatorProfileStats struct {
	Programs        int64    `json:"programs"`
	Clicks          int64    `json:"clicks"`
	Referrals       int64    `json:"referrals"`
	RepeatPayerRate *float64 `json:"repeatPayerRate"`
}

// CreatorProfile is generated from the OpenAPI component of the same name.
type CreatorProfile struct {
	AccountID      string                   `json:"accountId"`
	Name           *string                  `json:"name"`
	Email          *string                  `json:"email,omitempty"` // Only on your own profile.
	Published      bool                     `json:"published"`
	Headline       string                   `json:"headline"`
	Bio            string                   `json:"bio"`
	Categories     []string                 `json:"categories"`
	Languages      []string                 `json:"languages"`
	Country        string                   `json:"country"`
	Channels       []GetOpenapiJsonResponse `json:"channels"`
	Website        string                   `json:"website"`
	MinAudience    int64                    `json:"minAudience"`
	AcceptsInvites bool                     `json:"acceptsInvites"`
	Audience       GetOpenapiJsonResponse   `json:"audience"`
	EngagementBps  int64                    `json:"engagementBps"`
	Stats          CreatorProfileStats      `json:"stats"`
	Invited        *bool                    `json:"invited,omitempty"`
}

// EmailSettings is generated from the OpenAPI component of the same name.
type EmailSettings struct {
	Provider        *string `json:"provider"`
	HasAPIKey       bool    `json:"hasApiKey"` // The key itself is write-only and never returned.
	FromAddress     *string `json:"fromAddress"`
	FromName        *string `json:"fromName"`
	ReplyTo         *string `json:"replyTo"`
	DefaultLocale   string  `json:"defaultLocale"`
	LogoURL         *string `json:"logoUrl"`
	PrimaryColor    *string `json:"primaryColor"`
	DomainStatus    string  `json:"domainStatus"`
	DomainCheckedAt *string `json:"domainCheckedAt"`
}

// EmailRendered is generated from the OpenAPI component of the same name.
type EmailRendered struct {
	Locale  string `json:"locale"`
	Subject string `json:"subject"`
	Text    string `json:"text"`
	HTML    string `json:"html"`
	Source  string `json:"source"`
}

// Notification is generated from the OpenAPI component of the same name.
type Notification struct {
	ID        string                  `json:"id"`
	Topic     string                  `json:"topic"`
	Message   string                  `json:"message"`
	Data      *GetOpenapiJsonResponse `json:"data"`
	ReadAt    *string                 `json:"readAt"`
	CreatedAt string                  `json:"createdAt"`
}

// Passport is generated from the OpenAPI component of the same name.
type Passport struct {
	Claims    GetOpenapiJsonResponse `json:"claims"`
	Signature string                 `json:"signature"`
	Algorithm string                 `json:"algorithm"`
	PublicKey string                 `json:"publicKey"`
}

// IncidentUpdatesItem is generated from the OpenAPI component of the same name.
type IncidentUpdatesItem struct {
	ID        string `json:"id"`
	Body      string `json:"body"`
	State     string `json:"state"`
	CreatedAt string `json:"createdAt"`
}

// Incident is generated from the OpenAPI component of the same name.
type Incident struct {
	ID         string                `json:"id"`
	Title      string                `json:"title"`
	State      string                `json:"state"`
	Severity   string                `json:"severity"`
	Body       string                `json:"body"`
	Components []string              `json:"components"`
	StartedAt  string                `json:"startedAt"`
	ResolvedAt *string               `json:"resolvedAt"`
	UpdatedAt  string                `json:"updatedAt"`
	Updates    []IncidentUpdatesItem `json:"updates"`
}

// RetentionPolicy is generated from the OpenAPI component of the same name.
type RetentionPolicy struct {
	Clicks            *int64 `json:"clicks"`
	Visits            *int64 `json:"visits"`
	BillingPayloads   *int64 `json:"billingPayloads"`
	WebhookDeliveries *int64 `json:"webhookDeliveries"`
	Notifications     *int64 `json:"notifications"`
	AuditLog          *int64 `json:"auditLog"`
}

// RetentionLimitsValue is generated from the OpenAPI component of the same name.
type RetentionLimitsValue struct {
	MinDays     int64  `json:"minDays"`
	DefaultDays *int64 `json:"defaultDays"`
}

// RetentionLastRunOption1 is generated from the OpenAPI component of the same name.
type RetentionLastRunOption1 struct {
	At     string                 `json:"at"`
	Result GetOpenapiJsonResponse `json:"result"`
}

// Retention is generated from the OpenAPI component of the same name.
type Retention struct {
	Policy  RetentionPolicy                 `json:"policy"`
	Limits  map[string]RetentionLimitsValue `json:"limits"`
	LastRun *RetentionLastRunOption1        `json:"lastRun"`
}

// PartnerGroupBaseOption1Option1 is generated from the OpenAPI component of the same name.
type PartnerGroupBaseOption1Option1 struct {
	Kind string `json:"kind"`
	Bps  int64  `json:"bps"`
}

// PartnerGroupBaseOption1Option2 is generated from the OpenAPI component of the same name.
type PartnerGroupBaseOption1Option2 struct {
	Kind        string `json:"kind"`
	AmountMinor int64  `json:"amountMinor"` // Integer minor units (cents).
	Currency    string `json:"currency"`
}

// PartnerGroup is generated from the OpenAPI component of the same name.
type PartnerGroup struct {
	ID                  string          `json:"id"`
	ProgramID           string          `json:"programId"`
	Name                string          `json:"name"`
	Base                json.RawMessage `json:"base"`
	HoldDays            *int64          `json:"holdDays"`
	MinThresholdMinor   *int64          `json:"minThresholdMinor"`
	PayoutCurrency      *string         `json:"payoutCurrency"`
	OnboardingMessage   *string         `json:"onboardingMessage"`
	OnboardingChecklist []string        `json:"onboardingChecklist"`
	Members             *int64          `json:"members,omitempty"`
	ArchivedAt          *string         `json:"archivedAt"`
	CreatedAt           string          `json:"createdAt"`
}

// PromotionEffectOption1BaseOption1 is generated from the OpenAPI component of the same name.
type PromotionEffectOption1BaseOption1 struct {
	Kind string `json:"kind"`
	Bps  int64  `json:"bps"`
}

// PromotionEffectOption1BaseOption2 is generated from the OpenAPI component of the same name.
type PromotionEffectOption1BaseOption2 struct {
	Kind        string `json:"kind"`
	AmountMinor int64  `json:"amountMinor"` // Integer minor units (cents).
	Currency    string `json:"currency"`
}

// PromotionEffectOption1 is generated from the OpenAPI component of the same name.
type PromotionEffectOption1 struct {
	Kind string          `json:"kind"`
	Base json.RawMessage `json:"base"`
}

// PromotionEffectOption2 is generated from the OpenAPI component of the same name.
type PromotionEffectOption2 struct {
	Kind string `json:"kind"`
	Bps  int64  `json:"bps"`
}

// Promotion is generated from the OpenAPI component of the same name.
type Promotion struct {
	ID        string          `json:"id"`
	ProgramID string          `json:"programId"`
	Name      string          `json:"name"`
	GroupID   *string         `json:"groupId"`
	Effect    json.RawMessage `json:"effect"`
	StartsAt  string          `json:"startsAt"`
	EndsAt    string          `json:"endsAt"`
	EndedAt   *string         `json:"endedAt"`
	Status    string          `json:"status"`
	CreatedAt string          `json:"createdAt"`
}

// IntegrationConnection is generated from the OpenAPI component of the same name.
type IntegrationConnection struct {
	ID            string                 `json:"id"`
	Kind          string                 `json:"kind"`
	Status        string                 `json:"status"`
	Config        GetOpenapiJsonResponse `json:"config"`
	HasCredential bool                   `json:"hasCredential"`
	LastError     *string                `json:"lastError"`
	LastSuccessAt *string                `json:"lastSuccessAt"`
	PausedUntil   *string                `json:"pausedUntil"`
	CreatedAt     string                 `json:"createdAt"`
	UpdatedAt     string                 `json:"updatedAt"`
}

// IntegrationDelivery is generated from the OpenAPI component of the same name.
type IntegrationDelivery struct {
	ID        string  `json:"id"`
	Kind      string  `json:"kind"`
	TargetID  string  `json:"targetId"`
	EventID   string  `json:"eventId"`
	Topic     string  `json:"topic"`
	Status    string  `json:"status"`
	Attempts  int64   `json:"attempts"`
	NextAt    string  `json:"nextAt"`
	LastError *string `json:"lastError"`
	SentAt    *string `json:"sentAt"`
	CreatedAt string  `json:"createdAt"`
}

// ZapierHook is generated from the OpenAPI component of the same name.
type ZapierHook struct {
	ID        string `json:"id"`
	Event     string `json:"event"`
	TargetURL string `json:"target_url"`
	CreatedAt string `json:"createdAt"`
}

// RoiIncrementalTreatment: Recorded treatment referrals that got referee credit.
type RoiIncrementalTreatment struct {
	N                           int64  `json:"n"`
	CashAfterPerReferralMinor   int64  `json:"cashAfterPerReferralMinor"` // Integer minor units (cents).
	ProfitAfterPerReferralMinor *int64 `json:"profitAfterPerReferralMinor"`
}

// RoiIncrementalControl: Recorded control (holdout) referrals: no credit.
type RoiIncrementalControl struct {
	N                           int64  `json:"n"`
	CashAfterPerReferralMinor   int64  `json:"cashAfterPerReferralMinor"` // Integer minor units (cents).
	ProfitAfterPerReferralMinor *int64 `json:"profitAfterPerReferralMinor"`
}

// RoiIncrementalExcludedReferrals: Referrals in range left out of both groups, by reason.
type RoiIncrementalExcludedReferrals struct {
	Unrecorded             int64 `json:"unrecorded"`
	TreatmentWithoutCredit int64 `json:"treatmentWithoutCredit"`
	ControlWithCredit      int64 `json:"controlWithCredit"`
}

// RoiIncrementalExcludedItem is generated from the OpenAPI component of the same name.
type RoiIncrementalExcludedItem struct {
	EventID  string `json:"eventId"`
	Currency string `json:"currency"`
	Reason   string `json:"reason"`
}

// RoiIncremental: Present with model=incremental.
type RoiIncremental struct {
	Currency                   string                          `json:"currency"`
	WindowDays                 int64                           `json:"windowDays"`
	HoldoutBps                 int64                           `json:"holdoutBps"`
	Treatment                  RoiIncrementalTreatment         `json:"treatment"`            // Recorded treatment referrals that got referee credit.
	Control                    RoiIncrementalControl           `json:"control"`              // Recorded control (holdout) referrals: no credit.
	LiftMinorPerReferral       int64                           `json:"liftMinorPerReferral"` // Treatment minus control cash after the first qualifying payment, per referral.
	LiftProfitMinorPerReferral *int64                          `json:"liftProfitMinorPerReferral"`
	CreditCostPerReferralMinor int64                           `json:"creditCostPerReferralMinor"` // Referee credit cost per treatment referral. Only credit: control referrals also earn no commission, and that withheld commission is not counted here.
	IncrementalRoi             *float64                        `json:"incrementalRoi"`             // Lift profit ÷ credit cost per referral; null when not eligible, no margin or no cost. The denominator is referee credit only: the commission withheld from control referrals is not in it, so the lift also reflects that commission difference.
	Eligible                   bool                            `json:"eligible"`
	Reason                     *string                         `json:"reason,omitempty"`
	UnrecordedReferrals        int64                           `json:"unrecordedReferrals"` // Referrals from before groups were recorded: naive ROI only.
	ExcludedReferrals          RoiIncrementalExcludedReferrals `json:"excludedReferrals"`   // Referrals in range left out of both groups, by reason.
	Excluded                   []RoiIncrementalExcludedItem    `json:"excluded"`
}

// RoiProjectionCohortsItem is generated from the OpenAPI component of the same name.
type RoiProjectionCohortsItem struct {
	ProgramID         *string `json:"programId"`
	Role              string  `json:"role"`
	Eligible          bool    `json:"eligible"`
	Reason            *string `json:"reason,omitempty"`
	ClosedGrants      int64   `json:"closedGrants"`
	OpenGrants        int64   `json:"openGrants"`
	MedianCrossingDay *int64  `json:"medianCrossingDay,omitempty"` // First day the cohort's median profit/cost ratio reaches 1.0; null when it doesn't within the window.
	CurveThroughDay   *int64  `json:"curveThroughDay,omitempty"`   // Last day with at least 50 closed grants observed that far: the curve stops there.
	NoCrossingReason  *string `json:"noCrossingReason,omitempty"`
	BehindCohort      *int64  `json:"behindCohort,omitempty"` // Open grants past the crossing day without paying back.
}

// RoiProjection: Present with includeProjection=true.
type RoiProjection struct {
	Label                     string                     `json:"label"` // Always "projection": expected figures, never actuals.
	Eligible                  bool                       `json:"eligible"`
	Reason                    *string                    `json:"reason,omitempty"`
	ClosedGrants              int64                      `json:"closedGrants"`
	OpenGrants                int64                      `json:"openGrants"`
	ProjectedGrants           int64                      `json:"projectedGrants"`          // Open grants in cohorts with at least 50 closed grants.
	ExpectedPaidBackShareBps  *int64                     `json:"expectedPaidBackShareBps"` // Share of projected open grants expected to pay back (already paid back included).
	MedianExpectedPaybackDays *int64                     `json:"medianExpectedPaybackDays"`
	Cohorts                   []RoiProjectionCohortsItem `json:"cohorts"`
}

// RoiSummaryCoverage is generated from the OpenAPI component of the same name.
type RoiSummaryCoverage struct {
	CreditAppliedReportedBps *int64 `json:"creditAppliedReportedBps"` // Basis points (10000 = 100%); null when there is nothing to divide by.
	MarginCoveredBps         *int64 `json:"marginCoveredBps"`         // Basis points (10000 = 100%); null when there is nothing to divide by.
	InferredShareBps         *int64 `json:"inferredShareBps"`         // Basis points (10000 = 100%); null when there is nothing to divide by.
}

// RoiSummaryTargetOption1 is generated from the OpenAPI component of the same name.
type RoiSummaryTargetOption1 struct {
	TargetRoiBps int64  `json:"targetRoiBps"`
	Status       string `json:"status"`
}

// RoiSummary is generated from the OpenAPI component of the same name.
type RoiSummary struct {
	Currency                   string                   `json:"currency"`
	CreditCostMinor            int64                    `json:"creditCostMinor"` // Integer minor units (cents).
	CashAfterMinor             int64                    `json:"cashAfterMinor"`  // Integer minor units (cents).
	ProfitAfterMinor           *int64                   `json:"profitAfterMinor"`
	NetMinor                   *int64                   `json:"netMinor"`
	Roi                        *float64                 `json:"roi"`              // Profit after ÷ cost, 4 decimals; null with no cost or unknown profit.
	PaidBackShareBps           *int64                   `json:"paidBackShareBps"` // Basis points (10000 = 100%); null when there is nothing to divide by.
	MedianPaybackDays          *float64                 `json:"medianPaybackDays"`
	Grants                     int64                    `json:"grants"`
	ProfitSources              map[string]int64         `json:"profitSources"`
	Coverage                   RoiSummaryCoverage       `json:"coverage"`
	OutstandingCreditMinor     int64                    `json:"outstandingCreditMinor"`     // Face not yet used, expired or revoked, in the reporting currency.
	OutstandingCostIfUsedMinor *int64                   `json:"outstandingCostIfUsedMinor"` // Outstanding credit priced by each grant's cost model; null when a model needs a margin and none is known.
	Currencies                 []string                 `json:"currencies"`                 // Native currencies of the grants in scope (ignoring the currency filter), most grants first.
	Incremental                *RoiIncremental          `json:"incremental,omitempty"`
	Projection                 *RoiProjection           `json:"projection,omitempty"`
	Target                     *RoiSummaryTargetOption1 `json:"target,omitempty"` // With programId: status against the program's targetRoiBps (unknown when ROI is null); null when it sets no target.
}

// RoiGrantMetricsExcludedItem is generated from the OpenAPI component of the same name.
type RoiGrantMetricsExcludedItem struct {
	EventID  string `json:"eventId"`
	Currency string `json:"currency"`
	Reason   string `json:"reason"`
}

// RoiGrantMetrics is generated from the OpenAPI component of the same name.
type RoiGrantMetrics struct {
	GrantID          string                        `json:"grantId"`
	Role             string                        `json:"role"`
	Currency         string                        `json:"currency"`
	FaceMinor        int64                         `json:"faceMinor"`      // Integer minor units (cents).
	UsedMinor        int64                         `json:"usedMinor"`      // Integer minor units (cents).
	CostMinor        int64                         `json:"costMinor"`      // Integer minor units (cents).
	CashAfterMinor   int64                         `json:"cashAfterMinor"` // Integer minor units (cents).
	ProfitAfterMinor *int64                        `json:"profitAfterMinor"`
	NetMinor         *int64                        `json:"netMinor"`
	Roi              *float64                      `json:"roi"` // Profit after ÷ cost, 4 decimals; null with no cost or unknown profit.
	PaidBack         *bool                         `json:"paidBack"`
	PaybackDays      *int64                        `json:"paybackDays"`
	ProfitSource     string                        `json:"profitSource"`
	InferredShareBps int64                         `json:"inferredShareBps"`
	WindowDays       int64                         `json:"windowDays"`
	Excluded         []RoiGrantMetricsExcludedItem `json:"excluded"`
}

// RoiGrant is generated from the OpenAPI component of the same name.
type RoiGrant struct {
	ID               string   `json:"id"`
	Source           string   `json:"source"`
	Role             string   `json:"role"`
	Status           string   `json:"status"`
	CustomerRef      *string  `json:"customerRef"`
	PartnerID        *string  `json:"partnerId"`
	ReferralID       *string  `json:"referralId"`
	ProgramID        *string  `json:"programId"`
	GrantCurrency    string   `json:"grantCurrency"`
	GrantedAt        string   `json:"grantedAt"`
	ExpiresAt        *string  `json:"expiresAt"`
	Currency         string   `json:"currency"`
	FaceMinor        int64    `json:"faceMinor"`      // Integer minor units (cents).
	UsedMinor        int64    `json:"usedMinor"`      // Integer minor units (cents).
	CostMinor        int64    `json:"costMinor"`      // Integer minor units (cents).
	CashAfterMinor   int64    `json:"cashAfterMinor"` // Integer minor units (cents).
	ProfitAfterMinor *int64   `json:"profitAfterMinor"`
	NetMinor         *int64   `json:"netMinor"`
	Roi              *float64 `json:"roi"` // Profit after ÷ cost, 4 decimals; null with no cost or unknown profit.
	PaidBack         *bool    `json:"paidBack"`
	PaybackDays      *int64   `json:"paybackDays"`
	ProfitSource     string   `json:"profitSource"`
	InferredShareBps int64    `json:"inferredShareBps"`
	WindowDays       int64    `json:"windowDays"`
}

// RoiPatternEntrySubject is generated from the OpenAPI component of the same name.
type RoiPatternEntrySubject struct {
	CustomerRef *string `json:"customerRef,omitempty"`
	PartnerID   *string `json:"partnerId,omitempty"`
}

// RoiPatternEntry is generated from the OpenAPI component of the same name.
type RoiPatternEntry struct {
	Subject   RoiPatternEntrySubject `json:"subject"`
	ProgramID *string                `json:"programId"`
	Detail    map[string]any         `json:"detail"` // The key numbers behind the match; the keys depend on the pattern.
}

// RoiTraceSubject is generated from the OpenAPI component of the same name.
type RoiTraceSubject struct {
	Type         string   `json:"type"`
	ID           string   `json:"id"`
	CustomerRefs []string `json:"customerRefs"`
	PartnerID    *string  `json:"partnerId"`
}

// RoiTraceEventsItem is generated from the OpenAPI component of the same name.
type RoiTraceEventsItem struct {
	ID         string                     `json:"id"`
	Kind       string                     `json:"kind"`
	OccurredAt string                     `json:"occurredAt"`
	Amounts    map[string]json.RawMessage `json:"amounts"` // currency plus the minor-unit amounts this kind carries.
	Confidence string                     `json:"confidence"`
	Source     string                     `json:"source"`
	Superseded bool                       `json:"superseded"`
	Links      map[string]*string         `json:"links"`
}

// RoiTraceRunningNetItem is generated from the OpenAPI component of the same name.
type RoiTraceRunningNetItem struct {
	At       string `json:"at"`
	NetMinor *int64 `json:"netMinor"`
}

// RoiTraceExcludedItem is generated from the OpenAPI component of the same name.
type RoiTraceExcludedItem struct {
	EventID  string `json:"eventId"`
	Currency string `json:"currency"`
	Reason   string `json:"reason"`
}

// RoiTrace is generated from the OpenAPI component of the same name.
type RoiTrace struct {
	Subject    RoiTraceSubject          `json:"subject"`
	Events     []RoiTraceEventsItem     `json:"events"`
	Grants     []RoiGrantMetrics        `json:"grants"`
	RunningNet []RoiTraceRunningNetItem `json:"runningNet"` // One point per event: cumulative credit net (profit after − credit cost) in the reporting currency. Commissions are not in it.
	Excluded   []RoiTraceExcludedItem   `json:"excluded"`   // Rows left out of runningNet: no fx rate into the reporting currency on their date.
}

// AuditLogEntry is generated from the OpenAPI component of the same name.
type AuditLogEntry struct {
	ID          string                 `json:"id"`
	ActorType   string                 `json:"actorType"`
	ActorID     string                 `json:"actorId"`
	Action      string                 `json:"action"`
	SubjectType string                 `json:"subjectType"`
	SubjectID   string                 `json:"subjectId"`
	Data        GetOpenapiJsonResponse `json:"data"`
	CreatedAt   string                 `json:"createdAt"`
}

// EmailTemplateEventLocalesItemDefault is generated from the OpenAPI component of the same name.
type EmailTemplateEventLocalesItemDefault struct {
	Subject string `json:"subject"`
	Body    string `json:"body"`
	Cta     string `json:"cta"`
}

// EmailTemplateEventLocalesItemOverrideOption1 is generated from the OpenAPI component of the same name.
type EmailTemplateEventLocalesItemOverrideOption1 struct {
	Subject   string  `json:"subject"`
	Text      string  `json:"text"`
	HTML      *string `json:"html"`
	UpdatedAt string  `json:"updatedAt"`
}

// EmailTemplateEventLocalesItem is generated from the OpenAPI component of the same name.
type EmailTemplateEventLocalesItem struct {
	Locale   string                                        `json:"locale"`
	Default  EmailTemplateEventLocalesItemDefault          `json:"default"`
	Override *EmailTemplateEventLocalesItemOverrideOption1 `json:"override"`
}

// EmailTemplateEvent is generated from the OpenAPI component of the same name.
type EmailTemplateEvent struct {
	Event        string                          `json:"event"`
	Audience     string                          `json:"audience"`
	Placeholders []string                        `json:"placeholders"`
	Locales      []EmailTemplateEventLocalesItem `json:"locales"`
}

// ErrorError is generated from the OpenAPI component of the same name.
type ErrorError struct {
	Code      string  `json:"code"`
	Message   string  `json:"message"`
	Details   any     `json:"details,omitempty"`
	RequestID *string `json:"requestId,omitempty"`
	DocsURL   string  `json:"docs_url"` // The docs entry for this error code: what it means and how to fix it.
}

// Error is generated from the OpenAPI component of the same name.
type Error struct {
	Error ErrorError `json:"error"`
}

// ClickRecordedEventData: `data` of a `click.recorded` event: Fires for every click (tracked link redirects, `?via=` landings, server-side clicks). High volume: subscribe only if you need click-level data.
type ClickRecordedEventData struct {
	ClickID   string  `json:"clickId"`
	ProgramID string  `json:"programId"`
	PartnerID string  `json:"partnerId"`
	LinkID    *string `json:"linkId"`
}

// ClickRecordedEvent: A tracked click was recorded: the `click.recorded` delivery body.
type ClickRecordedEvent struct {
	ID        string                 `json:"id"` // The event id (`evt_…`); the same event delivered to several endpoints has the same id.
	Topic     string                 `json:"topic"`
	OrgID     string                 `json:"orgId"`
	Livemode  bool                   `json:"livemode"` // false for events from test mode.
	CreatedAt string                 `json:"createdAt"`
	Data      ClickRecordedEventData `json:"data"`
}

// ReferralCreatedEventData: `data` of a `referral.created` event: A new referral: from `POST /v1/referrals/associate`, an auth-platform sign-up, a mobile install, a won deal or an agent attribution.
type ReferralCreatedEventData struct {
	ReferralID         string  `json:"referralId"`
	ProgramID          string  `json:"programId"`
	PartnerID          string  `json:"partnerId"`
	PartnerExternalRef *string `json:"partnerExternalRef"` // The partner's id in your own system (its externalRef), or null.
	CustomerRef        string  `json:"customerRef"`        // Your id for the referred customer.
	Status             string  `json:"status"`
	MatchedVia         string  `json:"matchedVia"`
}

// ReferralCreatedEvent: A customer was attributed to a partner: the `referral.created` delivery body.
type ReferralCreatedEvent struct {
	ID        string                   `json:"id"` // The event id (`evt_…`); the same event delivered to several endpoints has the same id.
	Topic     string                   `json:"topic"`
	OrgID     string                   `json:"orgId"`
	Livemode  bool                     `json:"livemode"` // false for events from test mode.
	CreatedAt string                   `json:"createdAt"`
	Data      ReferralCreatedEventData `json:"data"`
}

// ReferralHeldEventData: `data` of a `referral.held` event: A risk or self-referral check put the referral in `pending_review`; `reasons` says why and `caseId` is the fraud case opened for it.
type ReferralHeldEventData struct {
	ReferralID         string   `json:"referralId"`
	ProgramID          string   `json:"programId"`
	PartnerID          string   `json:"partnerId"`
	PartnerExternalRef *string  `json:"partnerExternalRef"` // The partner's id in your own system (its externalRef), or null.
	CustomerRef        string   `json:"customerRef"`        // Your id for the referred customer.
	Status             string   `json:"status"`
	Reasons            []string `json:"reasons"`
	CaseID             string   `json:"caseId"`
}

// ReferralHeldEvent: A referral was held for review: the `referral.held` delivery body.
type ReferralHeldEvent struct {
	ID        string                `json:"id"` // The event id (`evt_…`); the same event delivered to several endpoints has the same id.
	Topic     string                `json:"topic"`
	OrgID     string                `json:"orgId"`
	Livemode  bool                  `json:"livemode"` // false for events from test mode.
	CreatedAt string                `json:"createdAt"`
	Data      ReferralHeldEventData `json:"data"`
}

// ReferralApprovedEventData: `data` of a `referral.approved` event: Back to `active`: approved by the merchant, a fraud case dismissed (`reason: "fraud_case_dismissed"`, with `caseId`), or a hold released by email verification (`reason: "email_verified"`).
type ReferralApprovedEventData struct {
	ReferralID         string  `json:"referralId"`
	ProgramID          string  `json:"programId"`
	PartnerID          string  `json:"partnerId"`
	PartnerExternalRef *string `json:"partnerExternalRef"` // The partner's id in your own system (its externalRef), or null.
	CustomerRef        string  `json:"customerRef"`        // Your id for the referred customer.
	Status             string  `json:"status"`
	Reason             *string `json:"reason,omitempty"`
	CaseID             *string `json:"caseId,omitempty"`
}

// ReferralApprovedEvent: A held referral was approved: the `referral.approved` delivery body.
type ReferralApprovedEvent struct {
	ID        string                    `json:"id"` // The event id (`evt_…`); the same event delivered to several endpoints has the same id.
	Topic     string                    `json:"topic"`
	OrgID     string                    `json:"orgId"`
	Livemode  bool                      `json:"livemode"` // false for events from test mode.
	CreatedAt string                    `json:"createdAt"`
	Data      ReferralApprovedEventData `json:"data"`
}

// ReferralVoidedEventData: `data` of a `referral.voided` event: It earns nothing more: voided by the merchant, a fraud case (`reason` `fraud_case_voided` or `partner_banned`, with `caseId`) or an identity provider (`fraud_confirmed`).
type ReferralVoidedEventData struct {
	ReferralID         string  `json:"referralId"`
	ProgramID          string  `json:"programId"`
	PartnerID          string  `json:"partnerId"`
	PartnerExternalRef *string `json:"partnerExternalRef"` // The partner's id in your own system (its externalRef), or null.
	CustomerRef        string  `json:"customerRef"`        // Your id for the referred customer.
	Status             string  `json:"status"`
	Reason             *string `json:"reason,omitempty"`
	CaseID             *string `json:"caseId,omitempty"`
}

// ReferralVoidedEvent: A referral was voided: the `referral.voided` delivery body.
type ReferralVoidedEvent struct {
	ID        string                  `json:"id"` // The event id (`evt_…`); the same event delivered to several endpoints has the same id.
	Topic     string                  `json:"topic"`
	OrgID     string                  `json:"orgId"`
	Livemode  bool                    `json:"livemode"` // false for events from test mode.
	CreatedAt string                  `json:"createdAt"`
	Data      ReferralVoidedEventData `json:"data"`
}

// ReferralPausedEventData: `data` of a `referral.paused` event: The merchant paused the referral: it earns no new commissions until resumed.
type ReferralPausedEventData struct {
	ReferralID         string  `json:"referralId"`
	ProgramID          string  `json:"programId"`
	PartnerID          string  `json:"partnerId"`
	PartnerExternalRef *string `json:"partnerExternalRef"` // The partner's id in your own system (its externalRef), or null.
	CustomerRef        string  `json:"customerRef"`        // Your id for the referred customer.
	Status             string  `json:"status"`
}

// ReferralPausedEvent: A referral was paused: the `referral.paused` delivery body.
type ReferralPausedEvent struct {
	ID        string                  `json:"id"` // The event id (`evt_…`); the same event delivered to several endpoints has the same id.
	Topic     string                  `json:"topic"`
	OrgID     string                  `json:"orgId"`
	Livemode  bool                    `json:"livemode"` // false for events from test mode.
	CreatedAt string                  `json:"createdAt"`
	Data      ReferralPausedEventData `json:"data"`
}

// ReferralResumedEventData: `data` of a `referral.resumed` event: The referral is `active` again.
type ReferralResumedEventData struct {
	ReferralID         string  `json:"referralId"`
	ProgramID          string  `json:"programId"`
	PartnerID          string  `json:"partnerId"`
	PartnerExternalRef *string `json:"partnerExternalRef"` // The partner's id in your own system (its externalRef), or null.
	CustomerRef        string  `json:"customerRef"`        // Your id for the referred customer.
	Status             string  `json:"status"`
}

// ReferralResumedEvent: A paused referral was resumed: the `referral.resumed` delivery body.
type ReferralResumedEvent struct {
	ID        string                   `json:"id"` // The event id (`evt_…`); the same event delivered to several endpoints has the same id.
	Topic     string                   `json:"topic"`
	OrgID     string                   `json:"orgId"`
	Livemode  bool                     `json:"livemode"` // false for events from test mode.
	CreatedAt string                   `json:"createdAt"`
	Data      ReferralResumedEventData `json:"data"`
}

// ReferralReassignedEventData: `data` of a `referral.reassigned` event: The merchant reassigned the referral; the payload names the new partner, with the previous one in `previousPartnerId` and `previousPartnerExternalRef`.
type ReferralReassignedEventData struct {
	ReferralID                 string  `json:"referralId"`
	ProgramID                  string  `json:"programId"`
	PartnerID                  string  `json:"partnerId"`
	PartnerExternalRef         *string `json:"partnerExternalRef"` // The partner's id in your own system (its externalRef), or null.
	CustomerRef                string  `json:"customerRef"`        // Your id for the referred customer.
	Status                     string  `json:"status"`
	PreviousPartnerID          string  `json:"previousPartnerId"`
	PreviousPartnerExternalRef *string `json:"previousPartnerExternalRef"`
}

// ReferralReassignedEvent: A referral moved to another partner: the `referral.reassigned` delivery body.
type ReferralReassignedEvent struct {
	ID        string                      `json:"id"` // The event id (`evt_…`); the same event delivered to several endpoints has the same id.
	Topic     string                      `json:"topic"`
	OrgID     string                      `json:"orgId"`
	Livemode  bool                        `json:"livemode"` // false for events from test mode.
	CreatedAt string                      `json:"createdAt"`
	Data      ReferralReassignedEventData `json:"data"`
}

// ReferralRejectedEventData: `data` of a `referral.rejected` event: A self-referral the program's `selfReferralPolicy: "reject"` refused, so no referral exists. A `fraud.case_opened` follows.
type ReferralRejectedEventData struct {
	ProgramID   string `json:"programId"`
	PartnerID   string `json:"partnerId"`
	CustomerRef string `json:"customerRef"`
	Reason      string `json:"reason"`
}

// ReferralRejectedEvent: A self-referral was refused: the `referral.rejected` delivery body.
type ReferralRejectedEvent struct {
	ID        string                    `json:"id"` // The event id (`evt_…`); the same event delivered to several endpoints has the same id.
	Topic     string                    `json:"topic"`
	OrgID     string                    `json:"orgId"`
	Livemode  bool                      `json:"livemode"` // false for events from test mode.
	CreatedAt string                    `json:"createdAt"`
	Data      ReferralRejectedEventData `json:"data"`
}

// CustomerAliasedEventData: `data` of a `customer.aliased` event: The app changed a user's id after their sign-up was filed (through an auth connection); their referrals now use `toCustomerRef`.
type CustomerAliasedEventData struct {
	FromCustomerRef string   `json:"fromCustomerRef"`
	ToCustomerRef   string   `json:"toCustomerRef"`
	ReferralIds     []string `json:"referralIds"`
	PartnerID       *string  `json:"partnerId"`
}

// CustomerAliasedEvent: A customer's id changed: the `customer.aliased` delivery body.
type CustomerAliasedEvent struct {
	ID        string                   `json:"id"` // The event id (`evt_…`); the same event delivered to several endpoints has the same id.
	Topic     string                   `json:"topic"`
	OrgID     string                   `json:"orgId"`
	Livemode  bool                     `json:"livemode"` // false for events from test mode.
	CreatedAt string                   `json:"createdAt"`
	Data      CustomerAliasedEventData `json:"data"`
}

// CommissionEarnedEventData: `data` of a `commission.earned` event: A payment (or a signup, trial or product-event bounty) earned a partner a commission line. Payment fields are null for bounties.
type CommissionEarnedEventData struct {
	CommissionID        string  `json:"commissionId"`
	PartnerID           string  `json:"partnerId"`
	ProgramID           string  `json:"programId"`
	AmountMinor         int64   `json:"amountMinor"` // Integer minor units (cents).
	Currency            string  `json:"currency"`    // Lowercase ISO 4217 code.
	ReferralID          *string `json:"referralId"`
	CustomerRef         *string `json:"customerRef"`
	BillingEventID      *string `json:"billingEventId"`
	ProviderEventID     *string `json:"providerEventId"` // The billing provider's own event id.
	InvoiceRef          *string `json:"invoiceRef"`
	CustomerExternalRef *string `json:"customerExternalRef"` // Your own id for the customer when known (an auth-connection user id), or null.
	PartnerExternalRef  *string `json:"partnerExternalRef"`
	Level               int64   `json:"level"` // 0 for the direct partner, 1 to 5 for sub-affiliate levels.
	Memo                string  `json:"memo"`
	BasisMinor          *int64  `json:"basisMinor"`
}

// CommissionEarnedEvent: A commission line was earned: the `commission.earned` delivery body.
type CommissionEarnedEvent struct {
	ID        string                    `json:"id"` // The event id (`evt_…`); the same event delivered to several endpoints has the same id.
	Topic     string                    `json:"topic"`
	OrgID     string                    `json:"orgId"`
	Livemode  bool                      `json:"livemode"` // false for events from test mode.
	CreatedAt string                    `json:"createdAt"`
	Data      CommissionEarnedEventData `json:"data"`
}

// CommissionReversedEventData: `data` of a `commission.reversed` event: A refund or dispute clawed back an earned line. `commissionId` and `relatedCommissionId` are the original line, `reversalId` the new `reversed` event; `amountMinor` is negative.
type CommissionReversedEventData struct {
	CommissionID        string  `json:"commissionId"`
	RelatedCommissionID string  `json:"relatedCommissionId"`
	ReversalID          string  `json:"reversalId"`
	PartnerID           string  `json:"partnerId"`
	ProgramID           string  `json:"programId"`
	AmountMinor         int64   `json:"amountMinor"` // Negative: the amount taken back.
	Currency            string  `json:"currency"`    // Lowercase ISO 4217 code.
	ReferralID          *string `json:"referralId"`
	CustomerRef         *string `json:"customerRef"`
	BillingEventID      *string `json:"billingEventId"`
	ProviderEventID     *string `json:"providerEventId"` // The billing provider's own event id.
	InvoiceRef          *string `json:"invoiceRef"`
	CustomerExternalRef *string `json:"customerExternalRef"` // Your own id for the customer when known (an auth-connection user id), or null.
	PartnerExternalRef  *string `json:"partnerExternalRef"`
	Level               int64   `json:"level"` // 0 for the direct partner, 1 to 5 for sub-affiliate levels.
	Memo                string  `json:"memo"`
}

// CommissionReversedEvent: A commission line was reversed: the `commission.reversed` delivery body.
type CommissionReversedEvent struct {
	ID        string                      `json:"id"` // The event id (`evt_…`); the same event delivered to several endpoints has the same id.
	Topic     string                      `json:"topic"`
	OrgID     string                      `json:"orgId"`
	Livemode  bool                        `json:"livemode"` // false for events from test mode.
	CreatedAt string                      `json:"createdAt"`
	Data      CommissionReversedEventData `json:"data"`
}

// CommissionSettledEventData: `data` of a `commission.settled` event: The merchant paid the line itself (`store_credit`, or `merchant_credit` for programs that settle as they earn), outside payout runs.
type CommissionSettledEventData struct {
	CommissionID string  `json:"commissionId"`
	PartnerID    string  `json:"partnerId"`
	ProgramID    string  `json:"programId"`
	ReferralID   *string `json:"referralId"`
	AmountMinor  int64   `json:"amountMinor"` // Integer minor units (cents).
	Currency     string  `json:"currency"`    // Lowercase ISO 4217 code.
	Method       string  `json:"method"`
	Reference    string  `json:"reference"`
}

// CommissionSettledEvent: A commission line was settled by the merchant: the `commission.settled` delivery body.
type CommissionSettledEvent struct {
	ID        string                     `json:"id"` // The event id (`evt_…`); the same event delivered to several endpoints has the same id.
	Topic     string                     `json:"topic"`
	OrgID     string                     `json:"orgId"`
	Livemode  bool                       `json:"livemode"` // false for events from test mode.
	CreatedAt string                     `json:"createdAt"`
	Data      CommissionSettledEventData `json:"data"`
}

// PartnerAppliedEventData: `data` of a `partner.applied` event: An application waiting for approval. `enrollmentId` is absent for applications from the public program page.
type PartnerAppliedEventData struct {
	ProgramID    string  `json:"programId"`
	PartnerID    string  `json:"partnerId"`
	EnrollmentID *string `json:"enrollmentId,omitempty"`
}

// PartnerAppliedEvent: A partner applied to a program: the `partner.applied` delivery body.
type PartnerAppliedEvent struct {
	ID        string                  `json:"id"` // The event id (`evt_…`); the same event delivered to several endpoints has the same id.
	Topic     string                  `json:"topic"`
	OrgID     string                  `json:"orgId"`
	Livemode  bool                    `json:"livemode"` // false for events from test mode.
	CreatedAt string                  `json:"createdAt"`
	Data      PartnerAppliedEventData `json:"data"`
}

// PartnerEnrolledEventData: `data` of a `partner.enrolled` event: The merchant (or an import) enrolled the partner in a program, approved at once.
type PartnerEnrolledEventData struct {
	ProgramID    string `json:"programId"`
	PartnerID    string `json:"partnerId"`
	EnrollmentID string `json:"enrollmentId"`
}

// PartnerEnrolledEvent: A partner was enrolled: the `partner.enrolled` delivery body.
type PartnerEnrolledEvent struct {
	ID        string                   `json:"id"` // The event id (`evt_…`); the same event delivered to several endpoints has the same id.
	Topic     string                   `json:"topic"`
	OrgID     string                   `json:"orgId"`
	Livemode  bool                     `json:"livemode"` // false for events from test mode.
	CreatedAt string                   `json:"createdAt"`
	Data      PartnerEnrolledEventData `json:"data"`
}

// PartnerApprovedEventData: `data` of a `partner.approved` event: An application was approved, or a partner joined a program that approves at once. `enrollmentId` is present when approved from an application.
type PartnerApprovedEventData struct {
	ProgramID    string  `json:"programId"`
	PartnerID    string  `json:"partnerId"`
	EnrollmentID *string `json:"enrollmentId,omitempty"`
}

// PartnerApprovedEvent: A partner was approved: the `partner.approved` delivery body.
type PartnerApprovedEvent struct {
	ID        string                   `json:"id"` // The event id (`evt_…`); the same event delivered to several endpoints has the same id.
	Topic     string                   `json:"topic"`
	OrgID     string                   `json:"orgId"`
	Livemode  bool                     `json:"livemode"` // false for events from test mode.
	CreatedAt string                   `json:"createdAt"`
	Data      PartnerApprovedEventData `json:"data"`
}

// PartnerLevelUpEventData: `data` of a `partner.level_up` event: Partner leveling promoted the partner; `referrals` is the count that qualified them.
type PartnerLevelUpEventData struct {
	ProgramID string `json:"programId"`
	PartnerID string `json:"partnerId"`
	Level     string `json:"level"`
	Referrals int64  `json:"referrals"`
}

// PartnerLevelUpEvent: A partner reached a new level: the `partner.level_up` delivery body.
type PartnerLevelUpEvent struct {
	ID        string                  `json:"id"` // The event id (`evt_…`); the same event delivered to several endpoints has the same id.
	Topic     string                  `json:"topic"`
	OrgID     string                  `json:"orgId"`
	Livemode  bool                    `json:"livemode"` // false for events from test mode.
	CreatedAt string                  `json:"createdAt"`
	Data      PartnerLevelUpEventData `json:"data"`
}

// PayoutScheduledEventData: `data` of a `payout.scheduled` event: A payout run included the partner; nothing has been paid yet.
type PayoutScheduledEventData struct {
	PayoutID    string `json:"payoutId"`
	PartnerID   string `json:"partnerId"`
	AmountMinor int64  `json:"amountMinor"` // Integer minor units (cents).
	Currency    string `json:"currency"`    // Lowercase ISO 4217 code.
}

// PayoutScheduledEvent: A payout was created in a run: the `payout.scheduled` delivery body.
type PayoutScheduledEvent struct {
	ID        string                   `json:"id"` // The event id (`evt_…`); the same event delivered to several endpoints has the same id.
	Topic     string                   `json:"topic"`
	OrgID     string                   `json:"orgId"`
	Livemode  bool                     `json:"livemode"` // false for events from test mode.
	CreatedAt string                   `json:"createdAt"`
	Data      PayoutScheduledEventData `json:"data"`
}

// PayoutSubmittedEventData: `data` of a `payout.submitted` event: The payout provider accepted the payout and settles it asynchronously; `payout.sent` or `payout.failed` follows.
type PayoutSubmittedEventData struct {
	PayoutID    string `json:"payoutId"`
	PartnerID   string `json:"partnerId"`
	AmountMinor int64  `json:"amountMinor"` // Integer minor units (cents).
	Currency    string `json:"currency"`    // Lowercase ISO 4217 code.
}

// PayoutSubmittedEvent: A payout was submitted to the provider: the `payout.submitted` delivery body.
type PayoutSubmittedEvent struct {
	ID        string                   `json:"id"` // The event id (`evt_…`); the same event delivered to several endpoints has the same id.
	Topic     string                   `json:"topic"`
	OrgID     string                   `json:"orgId"`
	Livemode  bool                     `json:"livemode"` // false for events from test mode.
	CreatedAt string                   `json:"createdAt"`
	Data      PayoutSubmittedEventData `json:"data"`
}

// PayoutSentEventData: `data` of a `payout.sent` event: The money left: the provider confirmed the payout (at once, or later through its settlement webhook).
type PayoutSentEventData struct {
	PayoutID    string `json:"payoutId"`
	PartnerID   string `json:"partnerId"`
	AmountMinor int64  `json:"amountMinor"` // Integer minor units (cents).
	Currency    string `json:"currency"`    // Lowercase ISO 4217 code.
}

// PayoutSentEvent: A payout was paid: the `payout.sent` delivery body.
type PayoutSentEvent struct {
	ID        string              `json:"id"` // The event id (`evt_…`); the same event delivered to several endpoints has the same id.
	Topic     string              `json:"topic"`
	OrgID     string              `json:"orgId"`
	Livemode  bool                `json:"livemode"` // false for events from test mode.
	CreatedAt string              `json:"createdAt"`
	Data      PayoutSentEventData `json:"data"`
}

// PayoutFailedEventData: `data` of a `payout.failed` event: The provider refused or failed the payout; `reason` is its message. The lines go back to the partner's approved balance.
type PayoutFailedEventData struct {
	PayoutID  string `json:"payoutId"`
	PartnerID string `json:"partnerId"`
	Reason    string `json:"reason"`
}

// PayoutFailedEvent: A payout failed: the `payout.failed` delivery body.
type PayoutFailedEvent struct {
	ID        string                `json:"id"` // The event id (`evt_…`); the same event delivered to several endpoints has the same id.
	Topic     string                `json:"topic"`
	OrgID     string                `json:"orgId"`
	Livemode  bool                  `json:"livemode"` // false for events from test mode.
	CreatedAt string                `json:"createdAt"`
	Data      PayoutFailedEventData `json:"data"`
}

// PayoutCancelledEventData: `data` of a `payout.cancelled` event: One per payout when an unexecuted payout run is cancelled.
type PayoutCancelledEventData struct {
	PayoutID  string `json:"payoutId"`
	PartnerID string `json:"partnerId"`
}

// PayoutCancelledEvent: A payout was cancelled: the `payout.cancelled` delivery body.
type PayoutCancelledEvent struct {
	ID        string                   `json:"id"` // The event id (`evt_…`); the same event delivered to several endpoints has the same id.
	Topic     string                   `json:"topic"`
	OrgID     string                   `json:"orgId"`
	Livemode  bool                     `json:"livemode"` // false for events from test mode.
	CreatedAt string                   `json:"createdAt"`
	Data      PayoutCancelledEventData `json:"data"`
}

// TaxFormDeliveredEventData: `data` of a `tax.form_delivered` event: A filed 1099-NEC was accepted and delivered; `emailed` says whether the filing provider also emailed it.
type TaxFormDeliveredEventData struct {
	DocumentID string `json:"documentId"`
	PartnerID  string `json:"partnerId"`
	TaxYear    int64  `json:"taxYear"`
	Form       string `json:"form"`
	Corrected  bool   `json:"corrected"`
	Emailed    bool   `json:"emailed"`
}

// TaxFormDeliveredEvent: A tax form was delivered to a partner: the `tax.form_delivered` delivery body.
type TaxFormDeliveredEvent struct {
	ID        string                    `json:"id"` // The event id (`evt_…`); the same event delivered to several endpoints has the same id.
	Topic     string                    `json:"topic"`
	OrgID     string                    `json:"orgId"`
	Livemode  bool                      `json:"livemode"` // false for events from test mode.
	CreatedAt string                    `json:"createdAt"`
	Data      TaxFormDeliveredEventData `json:"data"`
}

// RefereeRewardEarnedEventData: `data` of a `referee_reward.earned` event: The friend's side of a double-sided program: the referred customer's reward was issued.
type RefereeRewardEarnedEventData struct {
	RewardID            string  `json:"rewardId"`
	ReferralID          string  `json:"referralId"`
	ProgramID           string  `json:"programId"`
	CustomerRef         string  `json:"customerRef"`
	CustomerExternalRef *string `json:"customerExternalRef"`
	Kind                string  `json:"kind"`
	AmountMinor         int64   `json:"amountMinor"` // Integer minor units (cents).
	Currency            string  `json:"currency"`    // Lowercase ISO 4217 code.
	Code                *string `json:"code"`
	BillingEventID      *string `json:"billingEventId"`
	CommissionID        *string `json:"commissionId"`
}

// RefereeRewardEarnedEvent: A referred customer earned their reward: the `referee_reward.earned` delivery body.
type RefereeRewardEarnedEvent struct {
	ID        string                       `json:"id"` // The event id (`evt_…`); the same event delivered to several endpoints has the same id.
	Topic     string                       `json:"topic"`
	OrgID     string                       `json:"orgId"`
	Livemode  bool                         `json:"livemode"` // false for events from test mode.
	CreatedAt string                       `json:"createdAt"`
	Data      RefereeRewardEarnedEventData `json:"data"`
}

// CreditRedeemedEventData: `data` of a `credit.redeemed` event: From `POST /v1/rewards/{id}/redeem` or `POST /v1/rewards/redeem-by-code`. `amountMinor` is what this redemption used (capped at the balance); a repeated `externalId` does not fire again.
type CreditRedeemedEventData struct {
	GrantID        string `json:"grantId"`
	RewardID       string `json:"rewardId"`
	AmountMinor    int64  `json:"amountMinor"`    // Integer minor units (cents).
	RemainingMinor int64  `json:"remainingMinor"` // What is left on the grant afterwards.
}

// CreditRedeemedEvent: A referee reward was redeemed: the `credit.redeemed` delivery body.
type CreditRedeemedEvent struct {
	ID        string                  `json:"id"` // The event id (`evt_…`); the same event delivered to several endpoints has the same id.
	Topic     string                  `json:"topic"`
	OrgID     string                  `json:"orgId"`
	Livemode  bool                    `json:"livemode"` // false for events from test mode.
	CreatedAt string                  `json:"createdAt"`
	Data      CreditRedeemedEventData `json:"data"`
}

// BillingEventRecordedEventData: `data` of a `billing_event.recorded` event: Fires once per newly recorded billing event from any connection, never for duplicates.
type BillingEventRecordedEventData struct {
	BillingEventID string `json:"billingEventId"`
	Type           string `json:"type"`
	CustomerRef    string `json:"customerRef"`
	Currency       string `json:"currency"`    // Lowercase ISO 4217 code.
	AmountMinor    int64  `json:"amountMinor"` // Integer minor units (cents).
}

// BillingEventRecordedEvent: A billing event was recorded: the `billing_event.recorded` delivery body.
type BillingEventRecordedEvent struct {
	ID        string                        `json:"id"` // The event id (`evt_…`); the same event delivered to several endpoints has the same id.
	Topic     string                        `json:"topic"`
	OrgID     string                        `json:"orgId"`
	Livemode  bool                          `json:"livemode"` // false for events from test mode.
	CreatedAt string                        `json:"createdAt"`
	Data      BillingEventRecordedEventData `json:"data"`
}

// RoiPaybackReachedEventData: `data` of a `roi.payback_reached` event: Once per credit grant, the first time its recipient's profit after it covered its cost. A later refund that un-pays it doesn't fire it again.
type RoiPaybackReachedEventData struct {
	GrantID     string `json:"grantId"`
	PaybackDays int64  `json:"paybackDays"` // Whole days from the grant until it paid back.
	NetMinor    int64  `json:"netMinor"`    // Profit after minus cost, when it paid back.
	Currency    string `json:"currency"`    // Lowercase ISO 4217 code.
}

// RoiPaybackReachedEvent: A credit grant paid back: the `roi.payback_reached` delivery body.
type RoiPaybackReachedEvent struct {
	ID        string                     `json:"id"` // The event id (`evt_…`); the same event delivered to several endpoints has the same id.
	Topic     string                     `json:"topic"`
	OrgID     string                     `json:"orgId"`
	Livemode  bool                       `json:"livemode"` // false for events from test mode.
	CreatedAt string                     `json:"createdAt"`
	Data      RoiPaybackReachedEventData `json:"data"`
}

// RoiPatternDetectedEventData: `data` of a `roi.pattern_detected` event: A daily summary, at most once per pattern per day, only when the set of matching subjects changed and `count` is above zero. Read the full list from `GET /v1/roi/patterns/{pattern}`.
type RoiPatternDetectedEventData struct {
	Pattern  string   `json:"pattern"`
	Count    int64    `json:"count"`    // Subjects matching today.
	NewCount int64    `json:"newCount"` // Of those, how many were not in the pattern's last report.
	Sample   []string `json:"sample"`   // Up to 10 ids: credit grant ids for the first two patterns, partner ids for the others.
}

// RoiPatternDetectedEvent: A value pattern changed: the `roi.pattern_detected` delivery body.
type RoiPatternDetectedEvent struct {
	ID        string                      `json:"id"` // The event id (`evt_…`); the same event delivered to several endpoints has the same id.
	Topic     string                      `json:"topic"`
	OrgID     string                      `json:"orgId"`
	Livemode  bool                        `json:"livemode"` // false for events from test mode.
	CreatedAt string                      `json:"createdAt"`
	Data      RoiPatternDetectedEventData `json:"data"`
}

// FraudCaseOpenedEventData: `data` of a `fraud.case_opened` event: A case is waiting in the queue. `kind` is `self_referral`, `quality`, `ring`, `identity_risk` or `value_pattern`; the other fields depend on it (`partnerId`, `referralId`, ring `members`, identity-risk `reasons`, value `pattern`). A self-referral rejected by policy has `referralId: null` and `rejected: true`.
type FraudCaseOpenedEventData struct {
	CaseID     string   `json:"caseId"`
	Kind       string   `json:"kind"`
	PartnerID  *string  `json:"partnerId,omitempty"`
	ReferralID *string  `json:"referralId,omitempty"`
	Members    []string `json:"members,omitempty"`
	Reasons    []string `json:"reasons,omitempty"`
	Pattern    *string  `json:"pattern,omitempty"`
	Rejected   *bool    `json:"rejected,omitempty"`
}

// FraudCaseOpenedEvent: A fraud case was opened: the `fraud.case_opened` delivery body.
type FraudCaseOpenedEvent struct {
	ID        string                   `json:"id"` // The event id (`evt_…`); the same event delivered to several endpoints has the same id.
	Topic     string                   `json:"topic"`
	OrgID     string                   `json:"orgId"`
	Livemode  bool                     `json:"livemode"` // false for events from test mode.
	CreatedAt string                   `json:"createdAt"`
	Data      FraudCaseOpenedEventData `json:"data"`
}

// DealSubmittedEventData: `data` of a `deal.submitted` event: A B2B deal waiting for the merchant to approve.
type DealSubmittedEventData struct {
	DealID    string `json:"dealId"`
	PartnerID string `json:"partnerId"`
	Company   string `json:"company"`
}

// DealSubmittedEvent: A partner registered a deal: the `deal.submitted` delivery body.
type DealSubmittedEvent struct {
	ID        string                 `json:"id"` // The event id (`evt_…`); the same event delivered to several endpoints has the same id.
	Topic     string                 `json:"topic"`
	OrgID     string                 `json:"orgId"`
	Livemode  bool                   `json:"livemode"` // false for events from test mode.
	CreatedAt string                 `json:"createdAt"`
	Data      DealSubmittedEventData `json:"data"`
}

// DisputeResolvedEventData: `data` of a `dispute.resolved` event: The merchant resolved a partner's dispute with the `resolution` note.
type DisputeResolvedEventData struct {
	DisputeID  string `json:"disputeId"`
	Resolution string `json:"resolution"`
}

// DisputeResolvedEvent: A partner dispute was resolved: the `dispute.resolved` delivery body.
type DisputeResolvedEvent struct {
	ID        string                   `json:"id"` // The event id (`evt_…`); the same event delivered to several endpoints has the same id.
	Topic     string                   `json:"topic"`
	OrgID     string                   `json:"orgId"`
	Livemode  bool                     `json:"livemode"` // false for events from test mode.
	CreatedAt string                   `json:"createdAt"`
	Data      DisputeResolvedEventData `json:"data"`
}

// PrivacyRequestCompletedEventData: `data` of a `privacy.request_completed` event: An access or erasure request opened with `POST /v1/privacy/requests` completed or failed; `found` is present when completed.
type PrivacyRequestCompletedEventData struct {
	RequestID string `json:"requestId"`
	Type      string `json:"type"`
	Status    string `json:"status"`
	Found     *bool  `json:"found,omitempty"`
}

// PrivacyRequestCompletedEvent: A data-subject request finished: the `privacy.request_completed` delivery body.
type PrivacyRequestCompletedEvent struct {
	ID        string                           `json:"id"` // The event id (`evt_…`); the same event delivered to several endpoints has the same id.
	Topic     string                           `json:"topic"`
	OrgID     string                           `json:"orgId"`
	Livemode  bool                             `json:"livemode"` // false for events from test mode.
	CreatedAt string                           `json:"createdAt"`
	Data      PrivacyRequestCompletedEventData `json:"data"`
}

// WebhookEndpointDisabledEventData: `data` of a `webhook_endpoint.disabled` event: Sent to your other active endpoints when one is disabled after every delivery failed for 72 hours (or 1000 attempts in a row).
type WebhookEndpointDisabledEventData struct {
	EndpointID     string `json:"endpointId"`
	URL            string `json:"url"`
	Reason         string `json:"reason"`
	LastStatusCode *int64 `json:"lastStatusCode"`
	Days           int64  `json:"days"`
}

// WebhookEndpointDisabledEvent: A failing webhook endpoint was disabled: the `webhook_endpoint.disabled` delivery body.
type WebhookEndpointDisabledEvent struct {
	ID        string                           `json:"id"` // The event id (`evt_…`); the same event delivered to several endpoints has the same id.
	Topic     string                           `json:"topic"`
	OrgID     string                           `json:"orgId"`
	Livemode  bool                             `json:"livemode"` // false for events from test mode.
	CreatedAt string                           `json:"createdAt"`
	Data      WebhookEndpointDisabledEventData `json:"data"`
}

// WebhookEndpointFailingEventData: `data` of a `webhook_endpoint.failing` event: Sent to your other active endpoints (never the failing one) when an endpoint has failed every delivery for 1 hour, and again at 24 hours. It's disabled at `disablesAt` unless a delivery succeeds first.
type WebhookEndpointFailingEventData struct {
	EndpointID          string  `json:"endpointId"`
	URL                 string  `json:"url"`
	Stage               string  `json:"stage"`
	FailingSince        string  `json:"failingSince"`
	ConsecutiveFailures int64   `json:"consecutiveFailures"`
	LastStatusCode      *int64  `json:"lastStatusCode"`
	SampleError         *string `json:"sampleError"`
	DisablesAt          string  `json:"disablesAt"`
}

// WebhookEndpointFailingEvent: A webhook endpoint keeps failing: the `webhook_endpoint.failing` delivery body.
type WebhookEndpointFailingEvent struct {
	ID        string                          `json:"id"` // The event id (`evt_…`); the same event delivered to several endpoints has the same id.
	Topic     string                          `json:"topic"`
	OrgID     string                          `json:"orgId"`
	Livemode  bool                            `json:"livemode"` // false for events from test mode.
	CreatedAt string                          `json:"createdAt"`
	Data      WebhookEndpointFailingEventData `json:"data"`
}

// WebhookEndpointRecoveredEventData: `data` of a `webhook_endpoint.recovered` event: Sent to your active endpoints, including the recovered one, when a delivery succeeds after the endpoint was failing.
type WebhookEndpointRecoveredEventData struct {
	EndpointID   string `json:"endpointId"`
	URL          string `json:"url"`
	FailingSince string `json:"failingSince"`
	DownForMs    int64  `json:"downForMs"`
}

// WebhookEndpointRecoveredEvent: A failing webhook endpoint recovered: the `webhook_endpoint.recovered` delivery body.
type WebhookEndpointRecoveredEvent struct {
	ID        string                            `json:"id"` // The event id (`evt_…`); the same event delivered to several endpoints has the same id.
	Topic     string                            `json:"topic"`
	OrgID     string                            `json:"orgId"`
	Livemode  bool                              `json:"livemode"` // false for events from test mode.
	CreatedAt string                            `json:"createdAt"`
	Data      WebhookEndpointRecoveredEventData `json:"data"`
}

// PostTOrgTokenVisitResponse is generated from the OpenAPI component of the same name.
type PostTOrgTokenVisitResponse struct {
	Ok      bool    `json:"ok"`
	ClickID *string `json:"clickId"`
}

// GetTOrgTokenCodesCodeResponseRefereeRewardOption1 is generated from the OpenAPI component of the same name.
type GetTOrgTokenCodesCodeResponseRefereeRewardOption1 struct {
	Kind        string `json:"kind"`
	Bps         *int64 `json:"bps,omitempty"`
	AmountMinor *int64 `json:"amountMinor,omitempty"` // Integer minor units (cents).
	CapMinor    *int64 `json:"capMinor,omitempty"`    // Integer minor units (cents).
	Currency    string `json:"currency"`
	On          string `json:"on"`
}

// GetTOrgTokenCodesCodeResponse is generated from the OpenAPI component of the same name.
type GetTOrgTokenCodesCodeResponse struct {
	Valid              bool                                               `json:"valid"`
	PartnerDisplayName string                                             `json:"partnerDisplayName"` // First name only.
	RefereeReward      *GetTOrgTokenCodesCodeResponseRefereeRewardOption1 `json:"refereeReward"`
}

// GetTOrgTokenWidgetMeResponseProgram is generated from the OpenAPI component of the same name.
type GetTOrgTokenWidgetMeResponseProgram struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// GetTOrgTokenWidgetMeResponsePartner is generated from the OpenAPI component of the same name.
type GetTOrgTokenWidgetMeResponsePartner struct {
	DisplayName string `json:"displayName"`
}

// GetTOrgTokenWidgetMeResponseShareOption1 is generated from the OpenAPI component of the same name.
type GetTOrgTokenWidgetMeResponseShareOption1 struct {
	X        string `json:"x"`
	Facebook string `json:"facebook"`
	Linkedin string `json:"linkedin"`
	Whatsapp string `json:"whatsapp"`
	Email    string `json:"email"`
}

// GetTOrgTokenWidgetMeResponseProgress is generated from the OpenAPI component of the same name.
type GetTOrgTokenWidgetMeResponseProgress struct {
	Clicks          int64 `json:"clicks"`
	Signups         int64 `json:"signups"`
	PayingCustomers int64 `json:"payingCustomers"`
}

// GetTOrgTokenWidgetMeResponseEarningsItem is generated from the OpenAPI component of the same name.
type GetTOrgTokenWidgetMeResponseEarningsItem struct {
	Currency string `json:"currency"`
	Pending  int64  `json:"pending"`  // Integer minor units (cents).
	Approved int64  `json:"approved"` // Integer minor units (cents).
	Paid     int64  `json:"paid"`     // Integer minor units (cents).
}

// GetTOrgTokenWidgetMeResponseRefereeRewardOption1 is generated from the OpenAPI component of the same name.
type GetTOrgTokenWidgetMeResponseRefereeRewardOption1 struct {
	Kind        string `json:"kind"`
	Bps         *int64 `json:"bps,omitempty"`
	AmountMinor *int64 `json:"amountMinor,omitempty"` // Integer minor units (cents).
	CapMinor    *int64 `json:"capMinor,omitempty"`    // Integer minor units (cents).
	Currency    string `json:"currency"`
	On          string `json:"on"`
}

// GetTOrgTokenWidgetMeResponse is generated from the OpenAPI component of the same name.
type GetTOrgTokenWidgetMeResponse struct {
	Program       GetTOrgTokenWidgetMeResponseProgram               `json:"program"`
	Partner       GetTOrgTokenWidgetMeResponsePartner               `json:"partner"`
	Code          *string                                           `json:"code"`
	Coupon        *string                                           `json:"coupon"`
	Link          *string                                           `json:"link"`
	Share         *GetTOrgTokenWidgetMeResponseShareOption1         `json:"share"`
	Progress      GetTOrgTokenWidgetMeResponseProgress              `json:"progress"`
	Earnings      []GetTOrgTokenWidgetMeResponseEarningsItem        `json:"earnings"`
	RefereeReward *GetTOrgTokenWidgetMeResponseRefereeRewardOption1 `json:"refereeReward"`
}

// GetTOrgTokenMobileResolveResponse is generated from the OpenAPI component of the same name.
type GetTOrgTokenMobileResolveResponse struct {
	Matched     bool    `json:"matched"`
	ClickID     *string `json:"clickId,omitempty"`
	Code        *string `json:"code,omitempty"`
	Destination *string `json:"destination,omitempty"`
	ProgramID   *string `json:"programId,omitempty"`
	MatchType   *string `json:"matchType,omitempty"` // How the click was matched, strongest first: the exact user agent, the device hint on the same network, or the exact IP alone.
}

// PostTOrgTokenMobileInstallResponse is generated from the OpenAPI component of the same name.
type PostTOrgTokenMobileInstallResponse struct {
	Attributed bool    `json:"attributed"`
	ReferralID *string `json:"referralId"`
	PartnerID  *string `json:"partnerId"`
	ProgramID  string  `json:"programId"`
}

// GetWellKnownAppleAppSiteAssociationResponseApplinksDetailsItemComponentsItem is generated from the OpenAPI component of the same name.
type GetWellKnownAppleAppSiteAssociationResponseApplinksDetailsItemComponentsItem struct {
	Value string `json:"/"`
}

// GetWellKnownAppleAppSiteAssociationResponseApplinksDetailsItem is generated from the OpenAPI component of the same name.
type GetWellKnownAppleAppSiteAssociationResponseApplinksDetailsItem struct {
	AppIDs     []string                                                                       `json:"appIDs"`
	Components []GetWellKnownAppleAppSiteAssociationResponseApplinksDetailsItemComponentsItem `json:"components"`
}

// GetWellKnownAppleAppSiteAssociationResponseApplinks is generated from the OpenAPI component of the same name.
type GetWellKnownAppleAppSiteAssociationResponseApplinks struct {
	Details []GetWellKnownAppleAppSiteAssociationResponseApplinksDetailsItem `json:"details"`
}

// GetWellKnownAppleAppSiteAssociationResponse is generated from the OpenAPI component of the same name.
type GetWellKnownAppleAppSiteAssociationResponse struct {
	Applinks GetWellKnownAppleAppSiteAssociationResponseApplinks `json:"applinks"`
}

// GetWellKnownAssetlinksJsonResponseValueItemTarget is generated from the OpenAPI component of the same name.
type GetWellKnownAssetlinksJsonResponseValueItemTarget struct {
	Namespace              string   `json:"namespace"`
	PackageName            string   `json:"package_name"`
	Sha256CertFingerprints []string `json:"sha256_cert_fingerprints"`
}

// GetWellKnownAssetlinksJsonResponseValueItem is generated from the OpenAPI component of the same name.
type GetWellKnownAssetlinksJsonResponseValueItem struct {
	Relation []string                                          `json:"relation"`
	Target   GetWellKnownAssetlinksJsonResponseValueItemTarget `json:"target"`
}

// GetWellKnownAssetlinksJsonResponse is generated from the OpenAPI component of the same name.
type GetWellKnownAssetlinksJsonResponse = []GetWellKnownAssetlinksJsonResponseValueItem

// GetWellKnownForgedeckEmbedJsonResponseCheckers is generated from the OpenAPI component of the same name.
type GetWellKnownForgedeckEmbedJsonResponseCheckers struct {
	UserAgent string   `json:"userAgent"`
	Egress    []string `json:"egress"`
}

// GetWellKnownForgedeckEmbedJsonResponseSurfacesTrackingSnippetCookieDaysOption2 is generated from the OpenAPI component of the same name.
type GetWellKnownForgedeckEmbedJsonResponseSurfacesTrackingSnippetCookieDaysOption2 struct {
	Default int64  `json:"default"`
	Max     int64  `json:"max"`
	Note    string `json:"note"`
}

// GetWellKnownForgedeckEmbedJsonResponseSurfacesTrackingSnippetStorageItemPurpose is generated from the OpenAPI component of the same name.
type GetWellKnownForgedeckEmbedJsonResponseSurfacesTrackingSnippetStorageItemPurpose struct {
	En string `json:"en"`
}

// GetWellKnownForgedeckEmbedJsonResponseSurfacesTrackingSnippetStorageItem is generated from the OpenAPI component of the same name.
type GetWellKnownForgedeckEmbedJsonResponseSurfacesTrackingSnippetStorageItem struct {
	Name     string                                                                          `json:"name"`
	Type     string                                                                          `json:"type"`
	Category string                                                                          `json:"category"`
	Lifetime string                                                                          `json:"lifetime"`
	Purpose  GetWellKnownForgedeckEmbedJsonResponseSurfacesTrackingSnippetStorageItemPurpose `json:"purpose"`
	Legacy   *bool                                                                           `json:"legacy,omitempty"`
}

// GetWellKnownForgedeckEmbedJsonResponseSurfacesTrackingSnippet is generated from the OpenAPI component of the same name.
type GetWellKnownForgedeckEmbedJsonResponseSurfacesTrackingSnippet struct {
	Script                  string                                                                     `json:"script"`
	Csp                     map[string][]string                                                        `json:"csp"`
	ConsentCategory         string                                                                     `json:"consentCategory"`
	ConsentFallbackCategory string                                                                     `json:"consentFallbackCategory"`
	ConsentAdapters         []string                                                                   `json:"consentAdapters"`
	ConsentDefault          map[string]any                                                             `json:"consentDefault,omitempty"`
	CookieDays              json.RawMessage                                                            `json:"cookieDays"`
	Storage                 []GetWellKnownForgedeckEmbedJsonResponseSurfacesTrackingSnippetStorageItem `json:"storage"`
}

// GetWellKnownForgedeckEmbedJsonResponseSurfacesPortalEmbedStorageItemPurpose is generated from the OpenAPI component of the same name.
type GetWellKnownForgedeckEmbedJsonResponseSurfacesPortalEmbedStorageItemPurpose struct {
	En string `json:"en"`
}

// GetWellKnownForgedeckEmbedJsonResponseSurfacesPortalEmbedStorageItem is generated from the OpenAPI component of the same name.
type GetWellKnownForgedeckEmbedJsonResponseSurfacesPortalEmbedStorageItem struct {
	Name     string                                                                      `json:"name"`
	Type     string                                                                      `json:"type"`
	Category string                                                                      `json:"category"`
	Lifetime string                                                                      `json:"lifetime"`
	Purpose  GetWellKnownForgedeckEmbedJsonResponseSurfacesPortalEmbedStorageItemPurpose `json:"purpose"`
	Legacy   *bool                                                                       `json:"legacy,omitempty"`
}

// GetWellKnownForgedeckEmbedJsonResponseSurfacesPortalEmbed is generated from the OpenAPI component of the same name.
type GetWellKnownForgedeckEmbedJsonResponseSurfacesPortalEmbed struct {
	Script          string                                                                 `json:"script"`
	Csp             map[string][]string                                                    `json:"csp"`
	ConsentCategory string                                                                 `json:"consentCategory"`
	Storage         []GetWellKnownForgedeckEmbedJsonResponseSurfacesPortalEmbedStorageItem `json:"storage"`
}

// GetWellKnownForgedeckEmbedJsonResponseSurfaces is generated from the OpenAPI component of the same name.
type GetWellKnownForgedeckEmbedJsonResponseSurfaces struct {
	TrackingSnippet GetWellKnownForgedeckEmbedJsonResponseSurfacesTrackingSnippet `json:"tracking-snippet"`
	PortalEmbed     GetWellKnownForgedeckEmbedJsonResponseSurfacesPortalEmbed     `json:"portal-embed"`
}

// GetWellKnownForgedeckEmbedJsonResponseChangesItem is generated from the OpenAPI component of the same name.
type GetWellKnownForgedeckEmbedJsonResponseChangesItem struct {
	Effective string   `json:"effective"`
	Surface   string   `json:"surface"`
	Add       []string `json:"add,omitempty"`
	Remove    []string `json:"remove,omitempty"`
	Note      string   `json:"note"`
}

// GetWellKnownForgedeckEmbedJsonResponse is generated from the OpenAPI component of the same name.
type GetWellKnownForgedeckEmbedJsonResponse struct {
	Schema   string                                              `json:"schema"`
	Product  string                                              `json:"product"`
	Updated  string                                              `json:"updated"`
	Openapi  string                                              `json:"openapi"`
	Llms     string                                              `json:"llms"`
	Status   string                                              `json:"status"`
	Docs     string                                              `json:"docs"`
	Checkers GetWellKnownForgedeckEmbedJsonResponseCheckers      `json:"checkers"`
	Org      *string                                             `json:"org,omitempty"`
	Surfaces GetWellKnownForgedeckEmbedJsonResponseSurfaces      `json:"surfaces"`
	Changes  []GetWellKnownForgedeckEmbedJsonResponseChangesItem `json:"changes"`
}

// PostAgentsRegisterResponse is generated from the OpenAPI component of the same name.
type PostAgentsRegisterResponse struct {
	AgentID       string `json:"agentId"`
	Wallet        string `json:"wallet"`
	InstantPayout bool   `json:"instantPayout"`
}

// PostTOrgTokenAgentAttributeResponse is generated from the OpenAPI component of the same name.
type PostTOrgTokenAgentAttributeResponse struct {
	Attributed  bool   `json:"attributed"`
	ReferralID  string `json:"referralId"`
	AgentWallet string `json:"agentWallet"`
}

// PostTOrgTokenAgentNegotiateResponse is generated from the OpenAPI component of the same name.
type PostTOrgTokenAgentNegotiateResponse struct {
	Decision   string  `json:"decision"`
	RateBps    *int64  `json:"rateBps,omitempty"`
	CounterBps *int64  `json:"counterBps,omitempty"`
	PartnerID  *string `json:"partnerId,omitempty"`
	Reason     string  `json:"reason"`
}

// PostTOrgTokenConsentResponse is generated from the OpenAPI component of the same name.
type PostTOrgTokenConsentResponse struct {
	Ok          bool   `json:"ok"`
	CustomerRef string `json:"customerRef"`
	Granted     bool   `json:"granted"`
}

// GetVersionsResponseVersionsItem is generated from the OpenAPI component of the same name.
type GetVersionsResponseVersionsItem struct {
	Version  string   `json:"version"`
	Released string   `json:"released"`
	Changes  []string `json:"changes"`
}

// GetVersionsResponse is generated from the OpenAPI component of the same name.
type GetVersionsResponse struct {
	Latest            string                            `json:"latest"`
	Versions          []GetVersionsResponseVersionsItem `json:"versions"`
	DeprecationPolicy string                            `json:"deprecationPolicy"`
}

// GetAuthProvidersResponseDetailsItem is generated from the OpenAPI component of the same name.
type GetAuthProvidersResponseDetailsItem struct {
	ID    string `json:"id"`
	Name  string `json:"name"`  // The provider, for the sign-in button.
	Atlas bool   `json:"atlas"` // Atlas, which has its own branded one-click button.
}

// GetAuthProvidersResponse is generated from the OpenAPI component of the same name.
type GetAuthProvidersResponse struct {
	Providers []string                              `json:"providers"`
	Details   []GetAuthProvidersResponseDetailsItem `json:"details"`
}

// PostAuthBackchannelLogoutResponse is generated from the OpenAPI component of the same name.
type PostAuthBackchannelLogoutResponse struct {
	Ok bool `json:"ok"`
}

// PostAuthLogoutResponse is generated from the OpenAPI component of the same name.
type PostAuthLogoutResponse struct {
	Ok        bool    `json:"ok"`
	LogoutURL *string `json:"logoutUrl"`
}

// PostOrgsResponse is generated from the OpenAPI component of the same name.
type PostOrgsResponse struct {
	OrgID string `json:"orgId"`
}

// GetMeResponseUserOption1 is generated from the OpenAPI component of the same name.
type GetMeResponseUserOption1 struct {
	ID            string  `json:"id"`
	Email         string  `json:"email"`
	Name          *string `json:"name"`
	EmailVerified bool    `json:"emailVerified"`
}

// GetMeResponseMembershipsItem is generated from the OpenAPI component of the same name.
type GetMeResponseMembershipsItem struct {
	OrgID     string `json:"orgId"`
	Role      string `json:"role"`
	RoiDigest *bool  `json:"roiDigest,omitempty"`
}

// GetMeResponseApiKey: Set when called with an API key (user is then null).
type GetMeResponseApiKey struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Prefix      string   `json:"prefix"`
	Scopes      []string `json:"scopes"`
	APIVersion  *string  `json:"apiVersion"` // The API version this key is pinned to (what it gets without a Referalli-Version header).
	Livemode    bool     `json:"livemode"`
	Publishable bool     `json:"publishable"`
	Owner       *string  `json:"owner"`
	ExpiresAt   *string  `json:"expiresAt"`
}

// GetMeResponseOrg is generated from the OpenAPI component of the same name.
type GetMeResponseOrg struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// GetMeResponse is generated from the OpenAPI component of the same name.
type GetMeResponse struct {
	User        *GetMeResponseUserOption1      `json:"user"`
	Memberships []GetMeResponseMembershipsItem `json:"memberships"`
	APIKey      *GetMeResponseApiKey           `json:"apiKey,omitempty"` // Set when called with an API key (user is then null).
	Org         *GetMeResponseOrg              `json:"org,omitempty"`
}

// GetPortalAuthCurrentResponse is generated from the OpenAPI component of the same name.
type GetPortalAuthCurrentResponse struct {
	Adapter  string  `json:"adapter"`
	Name     *string `json:"name"`     // The sign-in button label.
	StartURL *string `json:"startUrl"` // Where the button sends the partner; null for Referalli accounts.
}

// GetOrgPortalAuthResponseAdaptersItem is generated from the OpenAPI component of the same name.
type GetOrgPortalAuthResponseAdaptersItem struct {
	ID      string `json:"id"`
	Label   string `json:"label"`
	Status  string `json:"status"`
	DocsURL string `json:"docsUrl"`
}

// GetOrgPortalAuthResponse is generated from the OpenAPI component of the same name.
type GetOrgPortalAuthResponse struct {
	Adapter      string                                 `json:"adapter"`
	Config       map[string]any                         `json:"config"` // The stored settings; a client secret is sealed and never returned.
	Adapters     []GetOrgPortalAuthResponseAdaptersItem `json:"adapters"`
	JwtAudience  string                                 `json:"jwtAudience"`  // For the `custom` adapter: the `aud` your tokens must carry.
	JwtSignInURL string                                 `json:"jwtSignInUrl"` // For the `custom` adapter: send the partner here with `token=<jwt>`.
	RedirectURI  string                                 `json:"redirectUri"`  // For the `oidc` adapter: the redirect URI to register at your IdP.
	HasSecret    bool                                   `json:"hasSecret"`
	Status       string                                 `json:"status"`
	LastOkAt     *string                                `json:"lastOkAt"`
	LastError    *string                                `json:"lastError"`
}

// PutOrgPortalAuthResponseAdaptersItem is generated from the OpenAPI component of the same name.
type PutOrgPortalAuthResponseAdaptersItem struct {
	ID      string `json:"id"`
	Label   string `json:"label"`
	Status  string `json:"status"`
	DocsURL string `json:"docsUrl"`
}

// PutOrgPortalAuthResponse is generated from the OpenAPI component of the same name.
type PutOrgPortalAuthResponse struct {
	Adapter      string                                 `json:"adapter"`
	Config       map[string]any                         `json:"config"` // The stored settings; a client secret is sealed and never returned.
	Adapters     []PutOrgPortalAuthResponseAdaptersItem `json:"adapters"`
	JwtAudience  string                                 `json:"jwtAudience"`  // For the `custom` adapter: the `aud` your tokens must carry.
	JwtSignInURL string                                 `json:"jwtSignInUrl"` // For the `custom` adapter: send the partner here with `token=<jwt>`.
	RedirectURI  string                                 `json:"redirectUri"`  // For the `oidc` adapter: the redirect URI to register at your IdP.
}

// PostOrgPortalAuthTestResponse is generated from the OpenAPI component of the same name.
type PostOrgPortalAuthTestResponse struct {
	Ok     bool    `json:"ok"`
	Detail string  `json:"detail"`
	Fix    *string `json:"fix,omitempty"`
}

// GetHostsCurrentResponseValueOption1Branding is generated from the OpenAPI component of the same name.
type GetHostsCurrentResponseValueOption1Branding struct {
	LogoURL string            `json:"logoUrl"`
	Theme   map[string]string `json:"theme"`
}

// GetHostsCurrentResponseValueOption1ProgramsItem is generated from the OpenAPI component of the same name.
type GetHostsCurrentResponseValueOption1ProgramsItem struct {
	Handle string `json:"handle"`
	Name   string `json:"name"`
}

// GetHostsCurrentResponseValueOption1ConsentTagsItem is generated from the OpenAPI component of the same name.
type GetHostsCurrentResponseValueOption1ConsentTagsItem struct {
	Tag    string                     `json:"tag"`
	Attrs  map[string]json.RawMessage `json:"attrs"`
	Inline *string                    `json:"inline,omitempty"`
}

// GetHostsCurrentResponseValueOption1ConsentSettingsLinkOption1 is generated from the OpenAPI component of the same name.
type GetHostsCurrentResponseValueOption1ConsentSettingsLinkOption1 struct {
	Call  *string           `json:"call,omitempty"`
	Attrs map[string]string `json:"attrs,omitempty"`
}

// GetHostsCurrentResponseValueOption1ConsentBuiltin is generated from the OpenAPI component of the same name.
type GetHostsCurrentResponseValueOption1ConsentBuiltin struct {
	PolicyURL string            `json:"policyUrl"`
	Text      map[string]string `json:"text,omitempty"`
}

// GetHostsCurrentResponseValueOption1Consent is generated from the OpenAPI component of the same name.
type GetHostsCurrentResponseValueOption1Consent struct {
	Mode         string                                                         `json:"mode"`
	Provider     string                                                         `json:"provider"`
	Tags         []GetHostsCurrentResponseValueOption1ConsentTagsItem           `json:"tags"`
	SettingsLink *GetHostsCurrentResponseValueOption1ConsentSettingsLinkOption1 `json:"settingsLink"`
	Builtin      *GetHostsCurrentResponseValueOption1ConsentBuiltin             `json:"builtin,omitempty"`
}

// GetHostsCurrentResponseValueOption1ChatOption1TagsItem is generated from the OpenAPI component of the same name.
type GetHostsCurrentResponseValueOption1ChatOption1TagsItem struct {
	Tag    string                     `json:"tag"`
	Attrs  map[string]json.RawMessage `json:"attrs"`
	Inline *string                    `json:"inline,omitempty"`
}

// GetHostsCurrentResponseValueOption1ChatOption1 is generated from the OpenAPI component of the same name.
type GetHostsCurrentResponseValueOption1ChatOption1 struct {
	Provider string                                                   `json:"provider"`
	Tags     []GetHostsCurrentResponseValueOption1ChatOption1TagsItem `json:"tags"`
}

// GetHostsCurrentResponseValueOption1 is generated from the OpenAPI component of the same name.
type GetHostsCurrentResponseValueOption1 struct {
	Kind     string                                            `json:"kind"` // A merchant's verified partner-portal domain.
	OrgID    string                                            `json:"orgId"`
	OrgName  string                                            `json:"orgName"`
	Branding GetHostsCurrentResponseValueOption1Branding       `json:"branding"`
	Programs []GetHostsCurrentResponseValueOption1ProgramsItem `json:"programs"` // The org's active programs with a public signup page (/join/:handle).
	Consent  GetHostsCurrentResponseValueOption1Consent        `json:"consent"`
	Chat     *GetHostsCurrentResponseValueOption1ChatOption1   `json:"chat"`
	Csp      string                                            `json:"csp"` // The Content-Security-Policy for this host's pages, built from the enabled adapters.
}

// GetHostsCurrentResponseValueOption2 is generated from the OpenAPI component of the same name.
type GetHostsCurrentResponseValueOption2 struct {
	Kind string `json:"kind"` // A merchant's verified link domain.
	Csp  string `json:"csp"`
}

// GetHostsCurrentResponseValueOption3 is generated from the OpenAPI component of the same name.
type GetHostsCurrentResponseValueOption3 struct {
	Kind string `json:"kind"` // Referalli's own app.
}

// GetHostsCurrentResponse is generated from the OpenAPI component of the same name.
type GetHostsCurrentResponse = json.RawMessage

// GetChatlyIdentityResponseIdentityOption1 is generated from the OpenAPI component of the same name.
type GetChatlyIdentityResponseIdentityOption1 struct {
	ExternalID string  `json:"externalId"` // `merchant:<user id>` or `partner:<partner account id>`.
	Email      string  `json:"email"`
	Name       *string `json:"name,omitempty"`
	UserHash   string  `json:"userHash"` // HMAC-SHA256 of `externalId` (hex), for the widget's identity verification.
}

// GetChatlyIdentityResponse is generated from the OpenAPI component of the same name.
type GetChatlyIdentityResponse struct {
	Identity *GetChatlyIdentityResponseIdentityOption1 `json:"identity"`
}

// PostMobilePurchasesResponse is generated from the OpenAPI component of the same name.
type PostMobilePurchasesResponse struct {
	Conversions  int64 `json:"conversions"`
	LinesCreated int64 `json:"linesCreated"`
	Reversals    int64 `json:"reversals"`
}

// PatchOrgResponse is generated from the OpenAPI component of the same name.
type PatchOrgResponse struct {
	ID                string                 `json:"id"`
	Name              string                 `json:"name"`
	Plan              string                 `json:"plan"`
	Settings          GetOpenapiJsonResponse `json:"settings"`
	PayoutKYCRequired bool                   `json:"payout_kyc_required"`
}

// GetOrgMembersResponse is generated from the OpenAPI component of the same name.
type GetOrgMembersResponse struct {
	Object     string   `json:"object"`
	Data       []Member `json:"data"`
	HasMore    bool     `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string  `json:"next_cursor"`
}

// PostOrgMembersResponse is generated from the OpenAPI component of the same name.
type PostOrgMembersResponse struct {
	UserID string `json:"userId"`
	Role   string `json:"role"`
}

// PatchOrgMembersUserIdResponse is generated from the OpenAPI component of the same name.
type PatchOrgMembersUserIdResponse struct {
	UserID string `json:"userId"`
	Role   string `json:"role"`
}

// PutOrgPayoutScheduleResponse is generated from the OpenAPI component of the same name.
type PutOrgPayoutScheduleResponse struct {
	DayOfMonth  *int64          `json:"dayOfMonth,omitempty"`
	Schedule    *PayoutSchedule `json:"schedule,omitempty"`
	NextRunDate *string         `json:"nextRunDate,omitempty"`
}

// GetOrgPayoutScheduleResponseProgramsItemOverrideOption1Option2 is generated from the OpenAPI component of the same name.
type GetOrgPayoutScheduleResponseProgramsItemOverrideOption1Option2 struct {
	Kind string `json:"kind"`
}

// GetOrgPayoutScheduleResponseProgramsItem is generated from the OpenAPI component of the same name.
type GetOrgPayoutScheduleResponseProgramsItem struct {
	ProgramID   string          `json:"programId"`
	Name        string          `json:"name"`
	Override    json.RawMessage `json:"override"`
	NextRunDate *string         `json:"nextRunDate"`
}

// GetOrgPayoutScheduleResponse is generated from the OpenAPI component of the same name.
type GetOrgPayoutScheduleResponse struct {
	Schedule    *PayoutSchedule                            `json:"schedule"`
	NextRunDate *string                                    `json:"nextRunDate"`
	Programs    []GetOrgPayoutScheduleResponseProgramsItem `json:"programs"`
}

// PutProgramsIdPayoutScheduleResponseScheduleOption1Option2 is generated from the OpenAPI component of the same name.
type PutProgramsIdPayoutScheduleResponseScheduleOption1Option2 struct {
	Kind string `json:"kind"`
}

// PutProgramsIdPayoutScheduleResponse is generated from the OpenAPI component of the same name.
type PutProgramsIdPayoutScheduleResponse struct {
	ProgramID   string          `json:"programId"`
	Schedule    json.RawMessage `json:"schedule"`
	NextRunDate *string         `json:"nextRunDate"`
}

// PutPartnersIdVatResponse is generated from the OpenAPI component of the same name.
type PutPartnersIdVatResponse struct {
	VATID                 *string `json:"vatId"`
	VATCountry            *string `json:"vatCountry"`
	VATVerifiedAt         *string `json:"vatVerifiedAt"`
	SelfBillingAcceptedAt *string `json:"selfBillingAcceptedAt"`
	SelfBillingVersion    *string `json:"selfBillingVersion"`
}

// GetPartnerOrgsOrgIdSelfBillingResponse is generated from the OpenAPI component of the same name.
type GetPartnerOrgsOrgIdSelfBillingResponse struct {
	VATID                 *string `json:"vatId"`
	VATCountry            *string `json:"vatCountry"`
	VATVerifiedAt         *string `json:"vatVerifiedAt"`
	SelfBillingAcceptedAt *string `json:"selfBillingAcceptedAt"`
	SelfBillingVersion    *string `json:"selfBillingVersion"`
	CurrentVersion        string  `json:"currentVersion"`
}

// PostPartnerOrgsOrgIdSelfBillingResponse is generated from the OpenAPI component of the same name.
type PostPartnerOrgsOrgIdSelfBillingResponse struct {
	VATID                 *string `json:"vatId"`
	VATCountry            *string `json:"vatCountry"`
	VATVerifiedAt         *string `json:"vatVerifiedAt"`
	SelfBillingAcceptedAt *string `json:"selfBillingAcceptedAt"`
	SelfBillingVersion    *string `json:"selfBillingVersion"`
	CurrentVersion        string  `json:"currentVersion"`
}

// PutOrgTaxFilingResponse is generated from the OpenAPI component of the same name.
type PutOrgTaxFilingResponse struct {
	Provider  string `json:"provider"`
	Connected bool   `json:"connected"`
}

// GetTax1099FilingsResponseDataItemDocumentsItem is generated from the OpenAPI component of the same name.
type GetTax1099FilingsResponseDataItemDocumentsItem struct {
	ID          string `json:"id"`
	PartnerID   string `json:"partnerId"`
	AmountMinor int64  `json:"amountMinor"` // Integer minor units (cents).
	Status      string `json:"status"`
}

// GetTax1099FilingsResponseDataItem is generated from the OpenAPI component of the same name.
type GetTax1099FilingsResponseDataItem struct {
	ID               string                                           `json:"id"`
	TaxYear          int64                                            `json:"taxYear"`
	Kind             string                                           `json:"kind"`
	Provider         string                                           `json:"provider"`
	ProviderFilingID *string                                          `json:"providerFilingId"`
	Status           string                                           `json:"status"`
	RecipientCount   int64                                            `json:"recipientCount"`
	TotalMinor       int64                                            `json:"totalMinor"` // Integer minor units (cents).
	Error            *string                                          `json:"error"`
	Documents        []GetTax1099FilingsResponseDataItemDocumentsItem `json:"documents"`
	CreatedAt        string                                           `json:"createdAt"`
}

// GetTax1099FilingsResponse is generated from the OpenAPI component of the same name.
type GetTax1099FilingsResponse struct {
	Object     string                              `json:"object"`
	Data       []GetTax1099FilingsResponseDataItem `json:"data"`
	HasMore    bool                                `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string                             `json:"next_cursor"`
}

// PostTax1099FilingsIdRefreshResponse is generated from the OpenAPI component of the same name.
type PostTax1099FilingsIdRefreshResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

// GetPartnerTaxDocumentsResponseDataItem is generated from the OpenAPI component of the same name.
type GetPartnerTaxDocumentsResponseDataItem struct {
	ID           string  `json:"id"`
	PartnerID    string  `json:"partnerId"`
	TaxYear      int64   `json:"taxYear"`
	Form         string  `json:"form"`
	AmountMinor  int64   `json:"amountMinor"` // Integer minor units (cents).
	Status       string  `json:"status"`
	CorrectionOf *string `json:"correctionOf"`
	DeliveredAt  *string `json:"deliveredAt"`
	CreatedAt    string  `json:"createdAt"`
	OrgID        string  `json:"orgId"`
	OrgName      string  `json:"orgName"`
}

// GetPartnerTaxDocumentsResponse is generated from the OpenAPI component of the same name.
type GetPartnerTaxDocumentsResponse struct {
	Object     string                                   `json:"object"`
	Data       []GetPartnerTaxDocumentsResponseDataItem `json:"data"`
	HasMore    bool                                     `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string                                  `json:"next_cursor"`
}

// GetStatusResponseComponentsItem is generated from the OpenAPI component of the same name.
type GetStatusResponseComponentsItem struct {
	ID     string         `json:"id"`
	State  string         `json:"state"`
	Detail map[string]any `json:"detail"`
}

// GetStatusResponse is generated from the OpenAPI component of the same name.
type GetStatusResponse struct {
	State         string                            `json:"state"`
	Components    []GetStatusResponseComponentsItem `json:"components"`
	OpenIncidents []Incident                        `json:"openIncidents"`
	CheckedAt     string                            `json:"checkedAt"`
}

// GetStatusIncidentsResponse is generated from the OpenAPI component of the same name.
type GetStatusIncidentsResponse struct {
	Object     string     `json:"object"`
	Data       []Incident `json:"data"`
	HasMore    bool       `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string    `json:"next_cursor"`
}

// GetOrgReferralParamsResponse is generated from the OpenAPI component of the same name.
type GetOrgReferralParamsResponse struct {
	BuiltIn []string `json:"builtIn"` // Always read: via, ref, aff.
	Params  []string `json:"params"`
}

// GetProgramsResponse is generated from the OpenAPI component of the same name.
type GetProgramsResponse struct {
	Object     string    `json:"object"`
	Data       []Program `json:"data"`
	HasMore    bool      `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string   `json:"next_cursor"`
}

// GetProgramsIdEnrollmentsResponseDataItem is generated from the OpenAPI component of the same name.
type GetProgramsIdEnrollmentsResponseDataItem struct {
	ID             string         `json:"id"`
	PartnerID      string         `json:"partnerId"`
	PartnerEmail   string         `json:"partnerEmail"`
	PartnerName    string         `json:"partnerName"`
	Status         string         `json:"status"`
	CreatedAt      string         `json:"createdAt"`
	NetworkWarning NetworkWarning `json:"networkWarning"`
}

// GetProgramsIdEnrollmentsResponse is generated from the OpenAPI component of the same name.
type GetProgramsIdEnrollmentsResponse struct {
	Object     string                                     `json:"object"`
	Data       []GetProgramsIdEnrollmentsResponseDataItem `json:"data"`
	HasMore    bool                                       `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string                                    `json:"next_cursor"`
}

// PostProgramsIdEnrollmentsEnrollmentIdApproveResponse is generated from the OpenAPI component of the same name.
type PostProgramsIdEnrollmentsEnrollmentIdApproveResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

// PostProgramsIdEnrollmentsEnrollmentIdRejectResponse is generated from the OpenAPI component of the same name.
type PostProgramsIdEnrollmentsEnrollmentIdRejectResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

// GetPartnersResponse is generated from the OpenAPI component of the same name.
type GetPartnersResponse struct {
	Object     string    `json:"object"`
	Data       []Partner `json:"data"`
	HasMore    bool      `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string   `json:"next_cursor"`
}

// GetPartnersIdResponse is generated from the OpenAPI component of the same name.
type GetPartnersIdResponse = any

// GetPartnersIdEnrollmentsResponseDataItem is generated from the OpenAPI component of the same name.
type GetPartnersIdEnrollmentsResponseDataItem struct {
	ID          string  `json:"id"`
	ProgramID   string  `json:"programId"`
	ProgramName string  `json:"programName"`
	Status      string  `json:"status"`
	GroupID     *string `json:"groupId"`
	CreatedAt   string  `json:"createdAt"`
}

// GetPartnersIdEnrollmentsResponse is generated from the OpenAPI component of the same name.
type GetPartnersIdEnrollmentsResponse struct {
	Object     string                                     `json:"object"`
	Data       []GetPartnersIdEnrollmentsResponseDataItem `json:"data"`
	HasMore    bool                                       `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string                                    `json:"next_cursor"`
}

// GetProgramsIdGraphResponse is generated from the OpenAPI component of the same name.
type GetProgramsIdGraphResponse = any

// PostApiKeysResponseValueOption2: mode "both": the live and test twins.
type PostApiKeysResponseValueOption2 struct {
	PairID string       `json:"pairId"`
	Live   IssuedApiKey `json:"live"`
	Test   IssuedApiKey `json:"test"`
}

// PostApiKeysResponse is generated from the OpenAPI component of the same name.
type PostApiKeysResponse = json.RawMessage

// GetApiKeysResponse is generated from the OpenAPI component of the same name.
type GetApiKeysResponse struct {
	Object     string   `json:"object"`
	Data       []ApiKey `json:"data"`
	HasMore    bool     `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string  `json:"next_cursor"`
}

// PatchApiKeysIdResponse is generated from the OpenAPI component of the same name.
type PatchApiKeysIdResponse struct {
	ID         string  `json:"id"`
	APIVersion string  `json:"apiVersion"`
	Owner      *string `json:"owner"`
	ExpiresAt  *string `json:"expiresAt"`
}

// DeleteApiKeysIdResponse is generated from the OpenAPI component of the same name.
type DeleteApiKeysIdResponse struct {
	Ok      bool     `json:"ok"`
	Revoked []string `json:"revoked"`
}

// GetAuditLogsResponse is generated from the OpenAPI component of the same name.
type GetAuditLogsResponse struct {
	Object     string          `json:"object"`
	Data       []AuditLogEntry `json:"data"`
	HasMore    bool            `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string         `json:"next_cursor"`
}

// GetTestClockResponse is generated from the OpenAPI component of the same name.
type GetTestClockResponse struct {
	Now           string `json:"now"`
	OffsetSeconds int64  `json:"offsetSeconds"` // How far the test clock is ahead of real time.
}

// PostTestClockAdvanceResponse is generated from the OpenAPI component of the same name.
type PostTestClockAdvanceResponse struct {
	Now           string `json:"now"`
	OffsetSeconds int64  `json:"offsetSeconds"`
	Approved      int64  `json:"approved"`   // Held commission lines approved because their hold ended.
	PayoutRuns    int64  `json:"payoutRuns"` // Scheduled payout runs that fell due and ran.
}

// GetTestEmailsResponseDataItem is generated from the OpenAPI component of the same name.
type GetTestEmailsResponseDataItem struct {
	ID        string `json:"id"`
	Event     string `json:"event"`
	To        string `json:"to"`
	Subject   string `json:"subject"`
	Text      string `json:"text"`
	HTML      string `json:"html"`
	CreatedAt string `json:"createdAt"`
}

// GetTestEmailsResponse is generated from the OpenAPI component of the same name.
type GetTestEmailsResponse struct {
	Object     string                          `json:"object"`
	Data       []GetTestEmailsResponseDataItem `json:"data"`
	HasMore    bool                            `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string                         `json:"next_cursor"`
}

// PostCliDeviceCodesResponse is generated from the OpenAPI component of the same name.
type PostCliDeviceCodesResponse struct {
	DeviceCode              string `json:"deviceCode"` // Secret: the CLI polls with it.
	UserCode                string `json:"userCode"`
	VerificationURI         string `json:"verificationUri"`
	VerificationURIComplete string `json:"verificationUriComplete"`
	ExpiresIn               int64  `json:"expiresIn"`
	Interval                int64  `json:"interval"`
}

// PostCliDeviceCodesTokenResponse is generated from the OpenAPI component of the same name.
type PostCliDeviceCodesTokenResponse struct {
	Status  string   `json:"status"`
	Key     *string  `json:"key,omitempty"` // The issued API key, returned once when approved.
	OrgID   *string  `json:"orgId,omitempty"`
	OrgName *string  `json:"orgName,omitempty"`
	Scopes  []string `json:"scopes,omitempty"`
}

// GetCliDeviceCodesUserCodeResponse is generated from the OpenAPI component of the same name.
type GetCliDeviceCodesUserCodeResponse struct {
	UserCode   string  `json:"userCode"`
	ClientName *string `json:"clientName"`
	CreatedAt  string  `json:"createdAt"`
	ExpiresAt  string  `json:"expiresAt"`
}

// PostCliDeviceCodesUserCodeApproveResponse is generated from the OpenAPI component of the same name.
type PostCliDeviceCodesUserCodeApproveResponse struct {
	Ok     bool     `json:"ok"`
	KeyID  string   `json:"keyId"`
	Scopes []string `json:"scopes"`
}

// PostCliListenersResponse is generated from the OpenAPI component of the same name.
type PostCliListenersResponse struct {
	ID             string   `json:"id"`
	Secret         string   `json:"secret"` // Signs forwarded events (Standard Webhooks); shown once.
	Topics         []string `json:"topics"`
	IdleTimeoutSec int64    `json:"idleTimeoutSec"`
}

// GetCliListenersIdEventsResponseEventsItem is generated from the OpenAPI component of the same name.
type GetCliListenersIdEventsResponseEventsItem struct {
	ID      string            `json:"id"`
	EventID string            `json:"eventId"`
	Topic   string            `json:"topic"`
	Body    string            `json:"body"` // The exact bytes to forward.
	Headers map[string]string `json:"headers"`
}

// GetCliListenersIdEventsResponse is generated from the OpenAPI component of the same name.
type GetCliListenersIdEventsResponse struct {
	Events []GetCliListenersIdEventsResponseEventsItem `json:"events"`
}

// PostCliTriggerResponseDeliveredTo is generated from the OpenAPI component of the same name.
type PostCliTriggerResponseDeliveredTo struct {
	Listeners int64 `json:"listeners"`
	Endpoints int64 `json:"endpoints"`
}

// PostCliTriggerResponse is generated from the OpenAPI component of the same name.
type PostCliTriggerResponse struct {
	EventID     string                            `json:"eventId"`
	Topic       string                            `json:"topic"`
	Payload     GetOpenapiJsonResponse            `json:"payload"`
	DeliveredTo PostCliTriggerResponseDeliveredTo `json:"deliveredTo"`
}

// GetCliTopicsResponse is generated from the OpenAPI component of the same name.
type GetCliTopicsResponse struct {
	Topics []string `json:"topics"`
}

// GetCliLogsResponseEntriesItem is generated from the OpenAPI component of the same name.
type GetCliLogsResponseEntriesItem struct {
	Seq       int64   `json:"seq"`
	RequestID string  `json:"requestId"`
	Method    string  `json:"method"`
	Path      string  `json:"path"`
	Status    int64   `json:"status"`
	LatencyMs int64   `json:"latencyMs"`
	ActorType string  `json:"actorType"`
	ActorID   *string `json:"actorId"`
	Ts        string  `json:"ts"`
}

// GetCliLogsResponse is generated from the OpenAPI component of the same name.
type GetCliLogsResponse struct {
	Entries []GetCliLogsResponseEntriesItem `json:"entries"`
	Cursor  int64                           `json:"cursor"`
}

// GetLinksResponse is generated from the OpenAPI component of the same name.
type GetLinksResponse struct {
	Object     string  `json:"object"`
	Data       []Link  `json:"data"`
	HasMore    bool    `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string `json:"next_cursor"`
}

// PostCouponsResponse is generated from the OpenAPI component of the same name.
type PostCouponsResponse struct {
	ID   string `json:"id"`
	Code string `json:"code"`
}

// GetCouponsResponse is generated from the OpenAPI component of the same name.
type GetCouponsResponse struct {
	Object     string   `json:"object"`
	Data       []Coupon `json:"data"`
	HasMore    bool     `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string  `json:"next_cursor"`
}

// PostWidgetTokensResponse is generated from the OpenAPI component of the same name.
type PostWidgetTokensResponse struct {
	Token     string `json:"token"`
	PartnerID string `json:"partnerId"`
	ExpiresAt string `json:"expiresAt"`
}

// GetReferralsResponse is generated from the OpenAPI component of the same name.
type GetReferralsResponse struct {
	Object     string     `json:"object"`
	Data       []Referral `json:"data"`
	HasMore    bool       `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string    `json:"next_cursor"`
}

// GetReferralsIdJourneyResponseTouchCreditsItem is generated from the OpenAPI component of the same name.
type GetReferralsIdJourneyResponseTouchCreditsItem struct {
	PartnerID string  `json:"partnerId"`
	Weight    float64 `json:"weight"`
	ClickID   *string `json:"clickId"`
}

// GetReferralsIdJourneyResponseTouchpointsItem is generated from the OpenAPI component of the same name.
type GetReferralsIdJourneyResponseTouchpointsItem struct {
	ClickID    string  `json:"clickId"`
	Ts         string  `json:"ts"`
	URL        *string `json:"url"`
	LinkID     *string `json:"linkId"`
	FraudScore int64   `json:"fraudScore"`
	Eligible   bool    `json:"eligible"`
	Weight     float64 `json:"weight"`
	Winner     bool    `json:"winner"`
}

// GetReferralsIdJourneyResponse is generated from the OpenAPI component of the same name.
type GetReferralsIdJourneyResponse struct {
	ID              string                                          `json:"id"`
	ProgramID       string                                          `json:"programId"`
	PartnerID       string                                          `json:"partnerId"`
	CustomerRef     string                                          `json:"customerRef"`
	Source          string                                          `json:"source"`
	MatchedVia      string                                          `json:"matchedVia"`
	Confidence      float64                                         `json:"confidence"`
	ClickID         *string                                         `json:"clickId"`
	LinkID          *string                                         `json:"linkId"`
	Status          string                                          `json:"status"`
	Ts              string                                          `json:"ts"`
	WindowExpiresAt string                                          `json:"windowExpiresAt"`
	TouchCredits    []GetReferralsIdJourneyResponseTouchCreditsItem `json:"touchCredits"`
	Reasoning       AttributionReasoning                            `json:"reasoning"`
	Touchpoints     []GetReferralsIdJourneyResponseTouchpointsItem  `json:"touchpoints"`
}

// PostAttributionReplayResponse is generated from the OpenAPI component of the same name.
type PostAttributionReplayResponse struct {
	Recomputed int64 `json:"recomputed"`
	Changed    int64 `json:"changed"`
}

// PostTrackEventsResponse is generated from the OpenAPI component of the same name.
type PostTrackEventsResponse struct {
	ClickID      *string `json:"clickId,omitempty"`
	EventID      *string `json:"eventId,omitempty"`
	Event        *string `json:"event,omitempty"`
	Recorded     *bool   `json:"recorded,omitempty"`
	ReferralID   *string `json:"referralId,omitempty"`
	Duplicate    *bool   `json:"duplicate,omitempty"`
	CommissionID *string `json:"commissionId,omitempty"`
	AmountMinor  *int64  `json:"amountMinor,omitempty"` // Integer minor units (cents).
	Currency     *string `json:"currency,omitempty"`
}

// PostCommissionRulesResponse is generated from the OpenAPI component of the same name.
type PostCommissionRulesResponse struct {
	ID        string                 `json:"id"`
	ProgramID string                 `json:"programId"`
	Version   int64                  `json:"version"`
	Config    GetOpenapiJsonResponse `json:"config"`
}

// GetCommissionRulesResponse is generated from the OpenAPI component of the same name.
type GetCommissionRulesResponse struct {
	Object     string           `json:"object"`
	Data       []CommissionRule `json:"data"`
	HasMore    bool             `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string          `json:"next_cursor"`
}

// GetCommissionsResponse is generated from the OpenAPI component of the same name.
type GetCommissionsResponse struct {
	Object     string            `json:"object"`
	Data       []CommissionEvent `json:"data"`
	HasMore    bool              `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string           `json:"next_cursor"`
}

// GetPartnersIdBalancesResponse is generated from the OpenAPI component of the same name.
type GetPartnersIdBalancesResponse struct {
	Object     string    `json:"object"`
	Data       []Balance `json:"data"`
	HasMore    bool      `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string   `json:"next_cursor"`
}

// PostCommissionsApproveDueResponse is generated from the OpenAPI component of the same name.
type PostCommissionsApproveDueResponse struct {
	Approved int64 `json:"approved"`
}

// PostCommissionsIdAdjustResponse is generated from the OpenAPI component of the same name.
type PostCommissionsIdAdjustResponse struct {
	ID string `json:"id"`
}

// PutBillingConnectionsIdSecretResponse is generated from the OpenAPI component of the same name.
type PutBillingConnectionsIdSecretResponse struct {
	ID       string `json:"id"`
	Provider string `json:"provider"`
	Ok       bool   `json:"ok"`
}

// PutBillingConnectionsIdCredentialsResponse is generated from the OpenAPI component of the same name.
type PutBillingConnectionsIdCredentialsResponse struct {
	ID       string `json:"id"`
	Provider string `json:"provider"`
	Ok       bool   `json:"ok"`
}

// GetBillingConnectionsResponse is generated from the OpenAPI component of the same name.
type GetBillingConnectionsResponse struct {
	Object     string              `json:"object"`
	Data       []BillingConnection `json:"data"`
	HasMore    bool                `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string             `json:"next_cursor"`
}

// PutBillingConnectionsIdValueMappingResponse is generated from the OpenAPI component of the same name.
type PutBillingConnectionsIdValueMappingResponse struct {
	ID           string       `json:"id"`
	ValueMapping ValueMapping `json:"valueMapping"`
}

// PostBillingConnectionsIdRotateSecretResponse is generated from the OpenAPI component of the same name.
type PostBillingConnectionsIdRotateSecretResponse struct {
	ID                      string `json:"id"`
	Provider                string `json:"provider"`
	WebhookSecret           string `json:"webhookSecret"` // Returned only here. The previous secret keeps verifying until previousSecretExpiresAt.
	PreviousSecretExpiresAt string `json:"previousSecretExpiresAt"`
}

// GetBillingEventsResponse is generated from the OpenAPI component of the same name.
type GetBillingEventsResponse struct {
	Object     string         `json:"object"`
	Data       []BillingEvent `json:"data"`
	HasMore    bool           `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string        `json:"next_cursor"`
}

// PostBillingInternalCustomersResponse is generated from the OpenAPI component of the same name.
type PostBillingInternalCustomersResponse struct {
	CustomerRef string `json:"customerRef"`
	Note        string `json:"note"`
}

// PutPartnersIdPayoutProfileResponse is generated from the OpenAPI component of the same name.
type PutPartnersIdPayoutProfileResponse struct {
	ID string `json:"id"`
}

// PatchPartnersIdPayoutProfileResponse is generated from the OpenAPI component of the same name.
type PatchPartnersIdPayoutProfileResponse struct {
	PartnerID         string `json:"partnerId"`
	Method            string `json:"method"`
	PayoutCurrency    string `json:"payoutCurrency"`
	MinThresholdMinor int64  `json:"minThresholdMinor"` // Integer minor units (cents).
}

// PostPayoutRunsIdExecuteResponse is generated from the OpenAPI component of the same name.
type PostPayoutRunsIdExecuteResponse struct {
	Sent      int64 `json:"sent"`
	Submitted int64 `json:"submitted"`
	Failed    int64 `json:"failed"`
}

// PostPayoutRunsIdCancelResponse is generated from the OpenAPI component of the same name.
type PostPayoutRunsIdCancelResponse struct {
	Cancelled int64 `json:"cancelled"`
}

// GetPayoutRunsResponse is generated from the OpenAPI component of the same name.
type GetPayoutRunsResponse struct {
	Object     string      `json:"object"`
	Data       []PayoutRun `json:"data"`
	HasMore    bool        `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string     `json:"next_cursor"`
}

// GetPayoutsResponse is generated from the OpenAPI component of the same name.
type GetPayoutsResponse struct {
	Object     string   `json:"object"`
	Data       []Payout `json:"data"`
	HasMore    bool     `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string  `json:"next_cursor"`
}

// GetTax1099ResponseDataItem is generated from the OpenAPI component of the same name.
type GetTax1099ResponseDataItem struct {
	PartnerID           string   `json:"partnerId"`
	TaxFormType         string   `json:"taxFormType"`
	TotalPaidUsdMinor   int64    `json:"totalPaidUsdMinor"` // Integer minor units (cents).
	RequiresForm        bool     `json:"requiresForm"`
	FXMissingCurrencies []string `json:"fxMissingCurrencies"`
}

// GetTax1099Response is generated from the OpenAPI component of the same name.
type GetTax1099Response struct {
	Object     string                       `json:"object"`
	Data       []GetTax1099ResponseDataItem `json:"data"`
	HasMore    bool                         `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string                      `json:"next_cursor"`
}

// GetFxRatesResponse is generated from the OpenAPI component of the same name.
type GetFxRatesResponse struct {
	Object     string   `json:"object"`
	Data       []FxRate `json:"data"`
	HasMore    bool     `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string  `json:"next_cursor"`
}

// GetOrgFxFeedResponse is generated from the OpenAPI component of the same name.
type GetOrgFxFeedResponse struct {
	Enabled      bool    `json:"enabled"`
	Source       string  `json:"source"`
	LastRateDate *string `json:"lastRateDate"`
}

// PutOrgFxFeedResponse is generated from the OpenAPI component of the same name.
type PutOrgFxFeedResponse struct {
	Enabled bool   `json:"enabled"`
	Source  string `json:"source"`
	Synced  int64  `json:"synced"`
}

// PostFxRatesSyncResponse is generated from the OpenAPI component of the same name.
type PostFxRatesSyncResponse struct {
	OrgID    string  `json:"orgId"`
	Date     *string `json:"date"`
	Inserted int64   `json:"inserted"`
}

// PostFxRatesClearOverrideResponse is generated from the OpenAPI component of the same name.
type PostFxRatesClearOverrideResponse struct {
	Base    string `json:"base"`
	Quote   string `json:"quote"`
	Cleared int64  `json:"cleared"`
}

// GetPayoutConnectionsResponseDataItem is generated from the OpenAPI component of the same name.
type GetPayoutConnectionsResponseDataItem struct {
	Provider    string `json:"provider"`
	Env         string `json:"env"`
	Status      string `json:"status"`
	ConnectedAt string `json:"connectedAt"`
}

// GetPayoutConnectionsResponse is generated from the OpenAPI component of the same name.
type GetPayoutConnectionsResponse struct {
	Object     string                                 `json:"object"`
	Data       []GetPayoutConnectionsResponseDataItem `json:"data"`
	HasMore    bool                                   `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string                                `json:"next_cursor"`
}

// PutPayoutConnectionsProviderResponse is generated from the OpenAPI component of the same name.
type PutPayoutConnectionsProviderResponse struct {
	Provider string `json:"provider"`
	Env      string `json:"env"`
	Status   string `json:"status"`
}

// DeletePayoutConnectionsProviderResponse is generated from the OpenAPI component of the same name.
type DeletePayoutConnectionsProviderResponse struct {
	Provider string `json:"provider"`
	Removed  bool   `json:"removed"`
}

// PostFraudInvalidTrafficScanResponsePartnersItem is generated from the OpenAPI component of the same name.
type PostFraudInvalidTrafficScanResponsePartnersItem struct {
	PartnerID     string             `json:"partnerId"`
	Clicks        int64              `json:"clicks"`
	Probability   float64            `json:"probability"`
	Label         string             `json:"label"`
	Contributions map[string]float64 `json:"contributions"`
}

// PostFraudInvalidTrafficScanResponse is generated from the OpenAPI component of the same name.
type PostFraudInvalidTrafficScanResponse struct {
	Window   string                                            `json:"window"`
	Partners []PostFraudInvalidTrafficScanResponsePartnersItem `json:"partners"`
	Flagged  int64                                             `json:"flagged"`
}

// GetFraudCasesResponse is generated from the OpenAPI component of the same name.
type GetFraudCasesResponse struct {
	Object     string      `json:"object"`
	Data       []FraudCase `json:"data"`
	HasMore    bool        `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string     `json:"next_cursor"`
}

// GetPartnersIdQualityResponseLiveMetrics is generated from the OpenAPI component of the same name.
type GetPartnersIdQualityResponseLiveMetrics struct {
	Customers           float64  `json:"customers"`
	RepeatPayerRate     float64  `json:"repeatPayerRate"`
	RefundRate          float64  `json:"refundRate"`
	DisputeRate         float64  `json:"disputeRate"`
	TrialToPaidRate     float64  `json:"trialToPaidRate"`
	IdentityRiskRate    *float64 `json:"identityRiskRate,omitempty"`
	ValuePatternPenalty *float64 `json:"valuePatternPenalty,omitempty"`
}

// GetPartnersIdQualityResponseLive is generated from the OpenAPI component of the same name.
type GetPartnersIdQualityResponseLive struct {
	Score   float64                                 `json:"score"`
	Metrics GetPartnersIdQualityResponseLiveMetrics `json:"metrics"`
}

// GetPartnersIdQualityResponseLastSnapshotOption1 is generated from the OpenAPI component of the same name.
type GetPartnersIdQualityResponseLastSnapshotOption1 struct {
	Score      float64                `json:"score"`
	Metrics    GetOpenapiJsonResponse `json:"metrics"`
	ComputedAt string                 `json:"computedAt"`
}

// GetPartnersIdQualityResponse is generated from the OpenAPI component of the same name.
type GetPartnersIdQualityResponse struct {
	PartnerID    string                                           `json:"partnerId"`
	Live         GetPartnersIdQualityResponseLive                 `json:"live"`
	LastSnapshot *GetPartnersIdQualityResponseLastSnapshotOption1 `json:"lastSnapshot"`
}

// PostQualityRecomputeResponse is generated from the OpenAPI component of the same name.
type PostQualityRecomputeResponse struct {
	Partners    int64 `json:"partners"`
	CasesOpened int64 `json:"casesOpened"`
	Suspended   int64 `json:"suspended"`
}

// PostFraudDetectRingsResponse is generated from the OpenAPI component of the same name.
type PostFraudDetectRingsResponse struct {
	RingsFound int64 `json:"ringsFound"`
}

// GetReportsFilteredClicksResponseDataItem is generated from the OpenAPI component of the same name.
type GetReportsFilteredClicksResponseDataItem struct {
	Day      string  `json:"day"`
	Filtered int64   `json:"filtered"`
	Total    int64   `json:"total"`
	AvgScore float64 `json:"avgScore"`
}

// GetReportsFilteredClicksResponse is generated from the OpenAPI component of the same name.
type GetReportsFilteredClicksResponse struct {
	Object     string                                     `json:"object"`
	Data       []GetReportsFilteredClicksResponseDataItem `json:"data"`
	HasMore    bool                                       `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string                                    `json:"next_cursor"`
}

// GetAnalyticsOverviewResponseFunnel is generated from the OpenAPI component of the same name.
type GetAnalyticsOverviewResponseFunnel struct {
	ClickToSignup float64 `json:"clickToSignup"`
	SignupToPaid  float64 `json:"signupToPaid"`
}

// GetAnalyticsOverviewResponse is generated from the OpenAPI component of the same name.
type GetAnalyticsOverviewResponse struct {
	Clicks                  int64                              `json:"clicks"`
	FilteredClicks          int64                              `json:"filteredClicks"`
	Visitors                int64                              `json:"visitors"`
	Signups                 int64                              `json:"signups"`
	Payments                int64                              `json:"payments"`
	RevenueMinor            int64                              `json:"revenueMinor"`            // Integer minor units (cents).
	CommissionEarnedMinor   int64                              `json:"commissionEarnedMinor"`   // Integer minor units (cents).
	CommissionReversedMinor int64                              `json:"commissionReversedMinor"` // Integer minor units (cents).
	ActivePartners          int64                              `json:"activePartners"`
	Funnel                  GetAnalyticsOverviewResponseFunnel `json:"funnel"`
}

// GetAnalyticsTimeseriesResponseDataItem is generated from the OpenAPI component of the same name.
type GetAnalyticsTimeseriesResponseDataItem struct {
	Day   string  `json:"day"`
	Value float64 `json:"value"`
}

// GetAnalyticsTimeseriesResponse is generated from the OpenAPI component of the same name.
type GetAnalyticsTimeseriesResponse struct {
	Object     string                                   `json:"object"`
	Data       []GetAnalyticsTimeseriesResponseDataItem `json:"data"`
	HasMore    bool                                     `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string                                  `json:"next_cursor"`
}

// GetAnalyticsCohortLtvResponseDataItem is generated from the OpenAPI component of the same name.
type GetAnalyticsCohortLtvResponseDataItem struct {
	Cohort       string `json:"cohort"`
	Customers    int64  `json:"customers"`
	RevenueMinor int64  `json:"revenueMinor"` // Integer minor units (cents).
	LtvMinor     int64  `json:"ltvMinor"`     // Integer minor units (cents).
}

// GetAnalyticsCohortLtvResponse is generated from the OpenAPI component of the same name.
type GetAnalyticsCohortLtvResponse struct {
	Object     string                                  `json:"object"`
	Data       []GetAnalyticsCohortLtvResponseDataItem `json:"data"`
	HasMore    bool                                    `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string                                 `json:"next_cursor"`
}

// GetAnalyticsByGeoResponseDataItem is generated from the OpenAPI component of the same name.
type GetAnalyticsByGeoResponseDataItem struct {
	Geo          string `json:"geo"`
	Referrals    int64  `json:"referrals"`
	RevenueMinor int64  `json:"revenueMinor"` // Integer minor units (cents).
}

// GetAnalyticsByGeoResponse is generated from the OpenAPI component of the same name.
type GetAnalyticsByGeoResponse struct {
	Object     string                              `json:"object"`
	Data       []GetAnalyticsByGeoResponseDataItem `json:"data"`
	HasMore    bool                                `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string                             `json:"next_cursor"`
}

// GetAnalyticsByChannelResponseDataItem is generated from the OpenAPI component of the same name.
type GetAnalyticsByChannelResponseDataItem struct {
	Channel      string `json:"channel"`
	Referrals    int64  `json:"referrals"`
	RevenueMinor int64  `json:"revenueMinor"` // Integer minor units (cents).
}

// GetAnalyticsByChannelResponse is generated from the OpenAPI component of the same name.
type GetAnalyticsByChannelResponse struct {
	Object     string                                  `json:"object"`
	Data       []GetAnalyticsByChannelResponseDataItem `json:"data"`
	HasMore    bool                                    `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string                                 `json:"next_cursor"`
}

// GetAnalyticsCouponPerformanceResponseDataItem is generated from the OpenAPI component of the same name.
type GetAnalyticsCouponPerformanceResponseDataItem struct {
	Code         string `json:"code"`
	PartnerID    string `json:"partnerId"`
	Referrals    int64  `json:"referrals"`
	RevenueMinor int64  `json:"revenueMinor"` // Integer minor units (cents).
}

// GetAnalyticsCouponPerformanceResponse is generated from the OpenAPI component of the same name.
type GetAnalyticsCouponPerformanceResponse struct {
	Object     string                                          `json:"object"`
	Data       []GetAnalyticsCouponPerformanceResponseDataItem `json:"data"`
	HasMore    bool                                            `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string                                         `json:"next_cursor"`
}

// GetReportsPartnerPerformanceResponseValueItem is generated from the OpenAPI component of the same name.
type GetReportsPartnerPerformanceResponseValueItem struct {
	PartnerID     string `json:"partnerId"`
	DisplayName   string `json:"displayName"`
	Email         string `json:"email"`
	Status        string `json:"status"`
	Clicks        int64  `json:"clicks"`
	Referrals     int64  `json:"referrals"`
	Payments      int64  `json:"payments"`
	RevenueMinor  int64  `json:"revenueMinor"`  // Integer minor units (cents).
	EarnedMinor   int64  `json:"earnedMinor"`   // Integer minor units (cents).
	ReversedMinor int64  `json:"reversedMinor"` // Integer minor units (cents).
}

// GetReportsPartnerPerformanceResponse is generated from the OpenAPI component of the same name.
type GetReportsPartnerPerformanceResponse = []GetReportsPartnerPerformanceResponseValueItem

// GetProfitabilitySettingsResponse is generated from the OpenAPI component of the same name.
type GetProfitabilitySettingsResponse struct {
	ReportingCurrency        string  `json:"reportingCurrency"`
	DefaultMarginBps         *int64  `json:"defaultMarginBps"`
	CreditCostModel          string  `json:"creditCostModel"`
	CreditCostRatioBps       *int64  `json:"creditCostRatioBps,omitempty"`
	Inference                string  `json:"inference"`
	CreditDraw               string  `json:"creditDraw"`
	WindowDays               int64   `json:"windowDays"`
	ChurnWindowDays          int64   `json:"churnWindowDays"`
	CountReferrerSpend       bool    `json:"countReferrerSpend"`
	CostToServeMinorPerMonth *int64  `json:"costToServeMinorPerMonth,omitempty"`
	TaxFeesReduceProfit      bool    `json:"taxFeesReduceProfit"`
	ValuePatternWeight       float64 `json:"valuePatternWeight"`
	TargetRoiBps             *int64  `json:"targetRoiBps,omitempty"` // Target ROI as a multiple in bps (20000 = 2.0x); unset or null = no target.
}

// GetProfitabilityProductsResponse is generated from the OpenAPI component of the same name.
type GetProfitabilityProductsResponse struct {
	Data []GetProfitabilityProductsProductRefResponse `json:"data"`
}

// GetProfitabilityProductsProductRefResponse is generated from the OpenAPI component of the same name.
type GetProfitabilityProductsProductRefResponse struct {
	ProductRef    string `json:"productRef"`
	MarginBps     *int64 `json:"marginBps"`
	UnitCostMinor *int64 `json:"unitCostMinor"`
	Currency      string `json:"currency"`
	UpdatedAt     string `json:"updatedAt"`
}

// PostProfitabilityRebuildResponse is generated from the OpenAPI component of the same name.
type PostProfitabilityRebuildResponse struct {
	JobID string `json:"jobId"`
}

// GetProfitabilityRebuildJobIdResponse is generated from the OpenAPI component of the same name.
type GetProfitabilityRebuildJobIdResponse struct {
	Status    string  `json:"status"`
	Processed int64   `json:"processed"`
	Total     int64   `json:"total"`
	Error     *string `json:"error,omitempty"`
}

// GetRoiBreakdownResponseDataItemCoverage is generated from the OpenAPI component of the same name.
type GetRoiBreakdownResponseDataItemCoverage struct {
	CreditAppliedReportedBps *int64 `json:"creditAppliedReportedBps"` // Basis points (10000 = 100%); null when there is nothing to divide by.
	MarginCoveredBps         *int64 `json:"marginCoveredBps"`         // Basis points (10000 = 100%); null when there is nothing to divide by.
	InferredShareBps         *int64 `json:"inferredShareBps"`         // Basis points (10000 = 100%); null when there is nothing to divide by.
}

// GetRoiBreakdownResponseDataItemReferrals: Partner and program rows: their referrals, both sides, commissions included.
type GetRoiBreakdownResponseDataItemReferrals struct {
	Count               int64    `json:"count"`
	CostMinor           int64    `json:"costMinor"`           // Integer minor units (cents).
	CashCommissionMinor int64    `json:"cashCommissionMinor"` // Integer minor units (cents).
	NetMinor            *int64   `json:"netMinor"`
	Roi                 *float64 `json:"roi"` // Profit after ÷ cost, 4 decimals; null with no cost or unknown profit.
}

// GetRoiBreakdownResponseDataItem is generated from the OpenAPI component of the same name.
type GetRoiBreakdownResponseDataItem struct {
	Currency                   string                                    `json:"currency"`
	CreditCostMinor            int64                                     `json:"creditCostMinor"` // Integer minor units (cents).
	CashAfterMinor             int64                                     `json:"cashAfterMinor"`  // Integer minor units (cents).
	ProfitAfterMinor           *int64                                    `json:"profitAfterMinor"`
	NetMinor                   *int64                                    `json:"netMinor"`
	Roi                        *float64                                  `json:"roi"`              // Profit after ÷ cost, 4 decimals; null with no cost or unknown profit.
	PaidBackShareBps           *int64                                    `json:"paidBackShareBps"` // Basis points (10000 = 100%); null when there is nothing to divide by.
	MedianPaybackDays          *float64                                  `json:"medianPaybackDays"`
	Grants                     int64                                     `json:"grants"`
	ProfitSources              map[string]int64                          `json:"profitSources"`
	Coverage                   GetRoiBreakdownResponseDataItemCoverage   `json:"coverage"`
	OutstandingCreditMinor     int64                                     `json:"outstandingCreditMinor"`     // Face not yet used, expired or revoked, in the reporting currency.
	OutstandingCostIfUsedMinor *int64                                    `json:"outstandingCostIfUsedMinor"` // Outstanding credit priced by each grant's cost model; null when a model needs a margin and none is known.
	Key                        string                                    `json:"key"`
	Label                      string                                    `json:"label"`
	Referrals                  *GetRoiBreakdownResponseDataItemReferrals `json:"referrals,omitempty"` // Partner and program rows: their referrals, both sides, commissions included.
}

// GetRoiBreakdownResponse is generated from the OpenAPI component of the same name.
type GetRoiBreakdownResponse struct {
	Object     string                            `json:"object"`
	Data       []GetRoiBreakdownResponseDataItem `json:"data"`
	HasMore    bool                              `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string                           `json:"next_cursor"`
}

// GetRoiGrantsResponse is generated from the OpenAPI component of the same name.
type GetRoiGrantsResponse struct {
	Object     string     `json:"object"`
	Data       []RoiGrant `json:"data"`
	HasMore    bool       `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string    `json:"next_cursor"`
}

// GetRoiPatternsResponsePatterns is generated from the OpenAPI component of the same name.
type GetRoiPatternsResponsePatterns struct {
	CreditOnly        []RoiPatternEntry `json:"credit_only"`
	ChurnAtExhaustion []RoiPatternEntry `json:"churn_at_exhaustion"`
	BonusCollector    []RoiPatternEntry `json:"bonus_collector"`
	NegativePartner   []RoiPatternEntry `json:"negative_partner"`
}

// GetRoiPatternsResponse is generated from the OpenAPI component of the same name.
type GetRoiPatternsResponse struct {
	Patterns GetRoiPatternsResponsePatterns `json:"patterns"`
}

// GetRoiPatternsPatternResponse is generated from the OpenAPI component of the same name.
type GetRoiPatternsPatternResponse struct {
	Object     string            `json:"object"`
	Data       []RoiPatternEntry `json:"data"`
	HasMore    bool              `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string           `json:"next_cursor"`
}

// GetRoiDataQualityResponseGapsItemFix is generated from the OpenAPI component of the same name.
type GetRoiDataQualityResponseGapsItemFix struct {
	Text string `json:"text"`
	Href string `json:"href"`
}

// GetRoiDataQualityResponseGapsItem is generated from the OpenAPI component of the same name.
type GetRoiDataQualityResponseGapsItem struct {
	Code  string                               `json:"code"`
	Count int64                                `json:"count"`
	Fix   GetRoiDataQualityResponseGapsItemFix `json:"fix"`
}

// GetRoiDataQualityResponse is generated from the OpenAPI component of the same name.
type GetRoiDataQualityResponse struct {
	PaymentsTotal          int64                               `json:"paymentsTotal"`
	PaymentsWithCreditInfo int64                               `json:"paymentsWithCreditInfo"`
	ProductsWithMargin     int64                               `json:"productsWithMargin"`
	ProductsSeen           int64                               `json:"productsSeen"`
	PartnersLinked         int64                               `json:"partnersLinked"`
	PartnersTotal          int64                               `json:"partnersTotal"`
	ExcludedForFX          int64                               `json:"excludedForFx"`
	Gaps                   []GetRoiDataQualityResponseGapsItem `json:"gaps"`
}

// PostValueEventsResponseResultsItem is generated from the OpenAPI component of the same name.
type PostValueEventsResponseResultsItem struct {
	Index      int64    `json:"index"`
	ID         string   `json:"id"`
	Created    bool     `json:"created"`
	Superseded []string `json:"superseded"`
}

// PostValueEventsResponseErrorsItem is generated from the OpenAPI component of the same name.
type PostValueEventsResponseErrorsItem struct {
	Index   int64  `json:"index"`
	Message string `json:"message"`
}

// PostValueEventsResponse is generated from the OpenAPI component of the same name.
type PostValueEventsResponse struct {
	Results []PostValueEventsResponseResultsItem `json:"results"`
	Errors  []PostValueEventsResponseErrorsItem  `json:"errors"`
}

// PostValueImportResponseValueOption1ErrorsItem is generated from the OpenAPI component of the same name.
type PostValueImportResponseValueOption1ErrorsItem struct {
	Row   int64  `json:"row"`
	Error string `json:"error"`
}

// PostValueImportResponseValueOption1 is generated from the OpenAPI component of the same name.
type PostValueImportResponseValueOption1 struct {
	Rows   []map[string]any                                `json:"rows"`
	Errors []PostValueImportResponseValueOption1ErrorsItem `json:"errors"`
	Total  int64                                           `json:"total"`
}

// PostValueImportResponseValueOption2ErrorsItem is generated from the OpenAPI component of the same name.
type PostValueImportResponseValueOption2ErrorsItem struct {
	Row   int64  `json:"row"`
	Error string `json:"error"`
}

// PostValueImportResponseValueOption2 is generated from the OpenAPI component of the same name.
type PostValueImportResponseValueOption2 struct {
	Recorded   int64                                           `json:"recorded"`
	Duplicates int64                                           `json:"duplicates"`
	Errors     []PostValueImportResponseValueOption2ErrorsItem `json:"errors"`
}

// PostValueImportResponseValueOption3 is generated from the OpenAPI component of the same name.
type PostValueImportResponseValueOption3 struct {
	JobID string `json:"jobId"` // Status 202: the file is recorded in the background; poll GET /v1/value/import/{jobId}.
}

// PostValueImportResponse is generated from the OpenAPI component of the same name.
type PostValueImportResponse = json.RawMessage

// GetValueImportJobIdResponseErrorsItem is generated from the OpenAPI component of the same name.
type GetValueImportJobIdResponseErrorsItem struct {
	Row   int64  `json:"row"`
	Error string `json:"error"`
}

// GetValueImportJobIdResponse is generated from the OpenAPI component of the same name.
type GetValueImportJobIdResponse struct {
	Status     string                                  `json:"status"`
	Processed  int64                                   `json:"processed"` // Parsed rows handed to recording so far.
	Total      int64                                   `json:"total"`     // Parsed rows to record (rows that failed to parse are only in errors).
	Recorded   int64                                   `json:"recorded"`
	Duplicates int64                                   `json:"duplicates"`
	Errors     []GetValueImportJobIdResponseErrorsItem `json:"errors"`          // Parse and record errors so far, by spreadsheet row (header is row 1).
	Error      *string                                 `json:"error,omitempty"` // Why the job failed ('stale': its process stopped).
}

// GetPartnerKycStatusResponse is generated from the OpenAPI component of the same name.
type GetPartnerKycStatusResponse struct {
	KYCStatus       string  `json:"kycStatus"`
	SanctionsStatus string  `json:"sanctionsStatus"`
	VerifiedAt      *string `json:"verifiedAt"`
}

// PostPartnerKycSubmitResponse is generated from the OpenAPI component of the same name.
type PostPartnerKycSubmitResponse struct {
	KYCStatus       string `json:"kycStatus"`
	SanctionsStatus string `json:"sanctionsStatus"`
}

// GetPartnerMeResponseAccount is generated from the OpenAPI component of the same name.
type GetPartnerMeResponseAccount struct {
	ID    string  `json:"id"`
	Email string  `json:"email"`
	Name  *string `json:"name"`
}

// GetPartnerMeResponseMembershipsItemProgramsItem is generated from the OpenAPI component of the same name.
type GetPartnerMeResponseMembershipsItemProgramsItem struct {
	ProgramID        string `json:"programId"`
	Name             string `json:"name"`
	Slug             string `json:"slug"`
	EnrollmentID     string `json:"enrollmentId"`
	EnrollmentStatus string `json:"enrollmentStatus"`
}

// GetPartnerMeResponseMembershipsItem is generated from the OpenAPI component of the same name.
type GetPartnerMeResponseMembershipsItem struct {
	OrgID     string                                            `json:"orgId"`
	OrgName   string                                            `json:"orgName"`
	PartnerID string                                            `json:"partnerId"`
	Programs  []GetPartnerMeResponseMembershipsItemProgramsItem `json:"programs"`
}

// GetPartnerMeResponse is generated from the OpenAPI component of the same name.
type GetPartnerMeResponse struct {
	Account     *GetPartnerMeResponseAccount          `json:"account,omitempty"`
	Memberships []GetPartnerMeResponseMembershipsItem `json:"memberships"`
}

// GetPartnerStatsResponseDataItem is generated from the OpenAPI component of the same name.
type GetPartnerStatsResponseDataItem struct {
	OrgID     string    `json:"orgId"`
	PartnerID string    `json:"partnerId"`
	Clicks    int64     `json:"clicks"`
	Referrals int64     `json:"referrals"`
	Payments  int64     `json:"payments"`
	Balances  []Balance `json:"balances"`
}

// GetPartnerStatsResponse is generated from the OpenAPI component of the same name.
type GetPartnerStatsResponse struct {
	Object     string                            `json:"object"`
	Data       []GetPartnerStatsResponseDataItem `json:"data"`
	HasMore    bool                              `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string                           `json:"next_cursor"`
}

// GetPartnerLinksResponseDataItem is generated from the OpenAPI component of the same name.
type GetPartnerLinksResponseDataItem struct {
	ID             string `json:"id"`
	OrgID          string `json:"orgId"`
	ProgramID      string `json:"programId"`
	Code           string `json:"code"`
	DestinationURL string `json:"destinationUrl"`
	ShortPath      string `json:"shortPath"`
	ShortURL       string `json:"shortUrl"` // Absolute link, on the merchant's verified link domain when it has one.
	DirectParam    string `json:"directParam"`
	Clicks         int64  `json:"clicks"`
}

// GetPartnerLinksResponse is generated from the OpenAPI component of the same name.
type GetPartnerLinksResponse struct {
	Object     string                            `json:"object"`
	Data       []GetPartnerLinksResponseDataItem `json:"data"`
	HasMore    bool                              `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string                           `json:"next_cursor"`
}

// GetPartnerAssetsResponseDataItem is generated from the OpenAPI component of the same name.
type GetPartnerAssetsResponseDataItem struct {
	ID          string `json:"id"`
	OrgID       string `json:"orgId"`
	OrgName     string `json:"orgName"`
	ProgramID   string `json:"programId"`
	ProgramName string `json:"programName"`
	Kind        string `json:"kind"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Dimensions  string `json:"dimensions"`
	CreatedAt   string `json:"createdAt"`
}

// GetPartnerAssetsResponse is generated from the OpenAPI component of the same name.
type GetPartnerAssetsResponse struct {
	Object     string                             `json:"object"`
	Data       []GetPartnerAssetsResponseDataItem `json:"data"`
	HasMore    bool                               `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string                            `json:"next_cursor"`
}

// GetPartnerDisputesResponseDataItem is generated from the OpenAPI component of the same name.
type GetPartnerDisputesResponseDataItem struct {
	ID         string  `json:"id"`
	OrgID      string  `json:"orgId"`
	OrgName    string  `json:"orgName"`
	ReferralID *string `json:"referralId"`
	Kind       string  `json:"kind"`
	Status     string  `json:"status"`
	Detail     string  `json:"detail"`
	Resolution *string `json:"resolution"`
	CreatedAt  string  `json:"createdAt"`
	ResolvedAt *string `json:"resolvedAt"`
}

// GetPartnerDisputesResponse is generated from the OpenAPI component of the same name.
type GetPartnerDisputesResponse struct {
	Object     string                               `json:"object"`
	Data       []GetPartnerDisputesResponseDataItem `json:"data"`
	HasMore    bool                                 `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string                              `json:"next_cursor"`
}

// GetPartnerReferralsResponseDataItem is generated from the OpenAPI component of the same name.
type GetPartnerReferralsResponseDataItem struct {
	ID         string `json:"id"`
	OrgID      string `json:"orgId"`
	ProgramID  string `json:"programId"`
	Customer   string `json:"customer"` // Redacted to the last 4 characters.
	Status     string `json:"status"`
	MatchedVia string `json:"matchedVia"`
	Ts         string `json:"ts"`
}

// GetPartnerReferralsResponse is generated from the OpenAPI component of the same name.
type GetPartnerReferralsResponse struct {
	Object     string                                `json:"object"`
	Data       []GetPartnerReferralsResponseDataItem `json:"data"`
	HasMore    bool                                  `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string                               `json:"next_cursor"`
}

// GetPartnerReferralsIdTransparencyResponse is generated from the OpenAPI component of the same name.
type GetPartnerReferralsIdTransparencyResponse struct {
	ReferralID            string       `json:"referralId"`
	Status                string       `json:"status"`
	MatchedVia            string       `json:"matchedVia"`
	Confidence            float64      `json:"confidence"`
	Model                 *string      `json:"model"`
	WindowDays            *int64       `json:"windowDays"`
	ConsideredTouchpoints int64        `json:"consideredTouchpoints"`
	YourCreditWeight      float64      `json:"yourCreditWeight"`
	Notes                 []string     `json:"notes"`
	NoteCodes             []ReasonNote `json:"noteCodes"`
}

// GetPartnerPayoutsResponseDataItem is generated from the OpenAPI component of the same name.
type GetPartnerPayoutsResponseDataItem struct {
	ID          string  `json:"id"`
	OrgID       string  `json:"orgId"`
	AmountMinor int64   `json:"amountMinor"` // Integer minor units (cents).
	Currency    string  `json:"currency"`
	Status      string  `json:"status"`
	Rail        string  `json:"rail"`
	SentAt      *string `json:"sentAt"`
	CreatedAt   string  `json:"createdAt"`
}

// GetPartnerPayoutsResponse is generated from the OpenAPI component of the same name.
type GetPartnerPayoutsResponse struct {
	Object     string                              `json:"object"`
	Data       []GetPartnerPayoutsResponseDataItem `json:"data"`
	HasMore    bool                                `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string                             `json:"next_cursor"`
}

// PutPartnerOrgsOrgIdPayoutProfileResponse is generated from the OpenAPI component of the same name.
type PutPartnerOrgsOrgIdPayoutProfileResponse struct {
	ID string `json:"id"`
}

// GetPartnerProgramsProgramIdLeaderboardResponseContestOption1 is generated from the OpenAPI component of the same name.
type GetPartnerProgramsProgramIdLeaderboardResponseContestOption1 struct {
	StartsAt string  `json:"startsAt"`
	EndsAt   string  `json:"endsAt"`
	Prize    *string `json:"prize"`
}

// GetPartnerProgramsProgramIdLeaderboardResponseRowsItem is generated from the OpenAPI component of the same name.
type GetPartnerProgramsProgramIdLeaderboardResponseRowsItem struct {
	Rank          int64  `json:"rank"`
	Name          string `json:"name"`
	ApprovedMinor int64  `json:"approvedMinor"` // Integer minor units (cents).
	You           bool   `json:"you"`
}

// GetPartnerProgramsProgramIdLeaderboardResponse is generated from the OpenAPI component of the same name.
type GetPartnerProgramsProgramIdLeaderboardResponse struct {
	Contest *GetPartnerProgramsProgramIdLeaderboardResponseContestOption1 `json:"contest"`
	Rows    []GetPartnerProgramsProgramIdLeaderboardResponseRowsItem      `json:"rows"`
}

// PostPartnerProgramsProgramIdApplyResponse is generated from the OpenAPI component of the same name.
type PostPartnerProgramsProgramIdApplyResponse struct {
	EnrollmentID string `json:"enrollmentId"`
	Status       string `json:"status"`
}

// PostWebhookEndpointsResponse is generated from the OpenAPI component of the same name.
type PostWebhookEndpointsResponse struct {
	ID              string   `json:"id"`
	URL             string   `json:"url"`
	Topics          []string `json:"topics"`
	ProgramIds      []string `json:"programIds"`
	LegacySignature bool     `json:"legacySignature"`
	AlertContacts   []string `json:"alertContacts"`
	Secret          string   `json:"secret"` // Returned only here. whsec_ + base64 key (Standard Webhooks).
}

// GetWebhookEndpointsResponse is generated from the OpenAPI component of the same name.
type GetWebhookEndpointsResponse struct {
	Object     string            `json:"object"`
	Data       []WebhookEndpoint `json:"data"`
	HasMore    bool              `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string           `json:"next_cursor"`
}

// PatchWebhookEndpointsIdResponse is generated from the OpenAPI component of the same name.
type PatchWebhookEndpointsIdResponse = any

// PostWebhookEndpointsIdRotateSecretResponse is generated from the OpenAPI component of the same name.
type PostWebhookEndpointsIdRotateSecretResponse struct {
	ID                      string `json:"id"`
	Secret                  string `json:"secret"` // The new secret, shown once. The old one keeps signing until previousSecretExpiresAt.
	PreviousSecretExpiresAt string `json:"previousSecretExpiresAt"`
}

// GetWebhookDeliveriesResponse is generated from the OpenAPI component of the same name.
type GetWebhookDeliveriesResponse struct {
	Object     string            `json:"object"`
	Data       []WebhookDelivery `json:"data"`
	HasMore    bool              `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string           `json:"next_cursor"`
}

// GetWebhookEndpointsIdDeliveriesResponse is generated from the OpenAPI component of the same name.
type GetWebhookEndpointsIdDeliveriesResponse struct {
	Object     string            `json:"object"`
	Data       []WebhookDelivery `json:"data"`
	HasMore    bool              `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string           `json:"next_cursor"`
}

// GetWebhookDeliveriesIdResponse is generated from the OpenAPI component of the same name.
type GetWebhookDeliveriesIdResponse = any

// PostWebhookEndpointsIdReplayResponse is generated from the OpenAPI component of the same name.
type PostWebhookEndpointsIdReplayResponse struct {
	EndpointID string `json:"endpointId"`
	Since      string `json:"since"`
	Replayed   int64  `json:"replayed"`
}

// PostSandboxSimulateResponse is generated from the OpenAPI component of the same name.
type PostSandboxSimulateResponse struct {
	CustomerRef  string `json:"customerRef"`
	Events       int64  `json:"events"`
	Conversions  int64  `json:"conversions"`
	LinesCreated int64  `json:"linesCreated"`
	Reversals    int64  `json:"reversals"`
}

// PostImportProviderResponse is generated from the OpenAPI component of the same name.
type PostImportProviderResponse struct {
	PartnersCreated     int64    `json:"partnersCreated"`
	PartnersSkipped     int64    `json:"partnersSkipped"`
	LinksCreated        int64    `json:"linksCreated"`
	CouponsCreated      int64    `json:"couponsCreated"`
	CodesSkipped        []string `json:"codesSkipped"`        // Codes another partner in this org already uses.
	ReferralParams      []string `json:"referralParams"`      // The old tool's link parameter, now read by aff.js.
	UnpaidImportedMinor int64    `json:"unpaidImportedMinor"` // Integer minor units (cents).
	PaidImportedMinor   int64    `json:"paidImportedMinor"`   // Integer minor units (cents).
}

// GetOpenapiJsonResponse is generated from the OpenAPI component of the same name.
type GetOpenapiJsonResponse = map[string]any

// GetBrandingResponseWarningsItem is generated from the OpenAPI component of the same name.
type GetBrandingResponseWarningsItem struct {
	Code       string  `json:"code"`
	Foreground string  `json:"foreground"`
	Background string  `json:"background"`
	Ratio      float64 `json:"ratio"`
	Minimum    float64 `json:"minimum"`
	Message    string  `json:"message"`
}

// GetBrandingResponse is generated from the OpenAPI component of the same name.
type GetBrandingResponse struct {
	LogoURL              *string                           `json:"logoUrl"`
	Theme                *GetOpenapiJsonResponse           `json:"theme"`
	PortalDomain         *string                           `json:"portalDomain"`
	PortalDomainVerified bool                              `json:"portalDomainVerified"`
	PortalDomainStatus   string                            `json:"portalDomainStatus"` // none; pending (waiting for DNS); provisioning (routed here, certificate being issued); active (links/portal served on it).
	LinkDomain           *string                           `json:"linkDomain"`
	LinkDomainVerified   bool                              `json:"linkDomainVerified"`
	LinkDomainStatus     string                            `json:"linkDomainStatus"` // none; pending (waiting for DNS); provisioning (routed here, certificate being issued); active (links/portal served on it).
	EmailFrom            *string                           `json:"emailFrom"`
	Warnings             []GetBrandingResponseWarningsItem `json:"warnings"` // Theme colour pairs below WCAG AA (4.5:1).
}

// PostBrandingLinkDomainResponse is generated from the OpenAPI component of the same name.
type PostBrandingLinkDomainResponse struct {
	Domain      string `json:"domain"`
	CnameTarget string `json:"cnameTarget"`
	TxtRecord   string `json:"txtRecord"`
	TxtName     string `json:"txtName"`
	Verified    bool   `json:"verified"`
}

// PostBrandingLinkDomainVerifyResponse is generated from the OpenAPI component of the same name.
type PostBrandingLinkDomainVerifyResponse struct {
	Domain   string  `json:"domain"`
	Verified bool    `json:"verified"`
	Status   string  `json:"status"`
	Message  *string `json:"message,omitempty"`
}

// PostBrandingDomainResponse is generated from the OpenAPI component of the same name.
type PostBrandingDomainResponse struct {
	Domain       string `json:"domain"`
	CnameTarget  string `json:"cnameTarget"`
	TxtRecord    string `json:"txtRecord"`
	TxtName      string `json:"txtName"`
	Verified     bool   `json:"verified"`
	Instructions string `json:"instructions"`
}

// PutBrandingResponseWarningsItem is generated from the OpenAPI component of the same name.
type PutBrandingResponseWarningsItem struct {
	Code       string  `json:"code"`
	Foreground string  `json:"foreground"`
	Background string  `json:"background"`
	Ratio      float64 `json:"ratio"`
	Minimum    float64 `json:"minimum"`
	Message    string  `json:"message"`
}

// PutBrandingResponse is generated from the OpenAPI component of the same name.
type PutBrandingResponse struct {
	Ok       bool                              `json:"ok"`
	Warnings []PutBrandingResponseWarningsItem `json:"warnings"`
}

// PostEmbedTokensResponse is generated from the OpenAPI component of the same name.
type PostEmbedTokensResponse struct {
	Token     string `json:"token"`
	ExpiresAt string `json:"expiresAt"`
}

// GetEmbedSummaryResponsePartner is generated from the OpenAPI component of the same name.
type GetEmbedSummaryResponsePartner struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// GetEmbedSummaryResponseBalancesItem is generated from the OpenAPI component of the same name.
type GetEmbedSummaryResponseBalancesItem struct {
	Currency string `json:"currency"`
	Pending  int64  `json:"pending"`  // Integer minor units (cents).
	Approved int64  `json:"approved"` // Integer minor units (cents).
	Paid     int64  `json:"paid"`     // Integer minor units (cents).
}

// GetEmbedSummaryResponseLinksItem is generated from the OpenAPI component of the same name.
type GetEmbedSummaryResponseLinksItem struct {
	Code           string `json:"code"`
	DestinationURL string `json:"destinationUrl"`
}

// GetEmbedSummaryResponse is generated from the OpenAPI component of the same name.
type GetEmbedSummaryResponse struct {
	Partner   GetEmbedSummaryResponsePartner        `json:"partner"`
	Clicks    int64                                 `json:"clicks"`
	Referrals int64                                 `json:"referrals"`
	Balances  []GetEmbedSummaryResponseBalancesItem `json:"balances"`
	Links     []GetEmbedSummaryResponseLinksItem    `json:"links"`
}

// PostDiscoverySearchResponse is generated from the OpenAPI component of the same name.
type PostDiscoverySearchResponse struct {
	Found int64 `json:"found"`
	Added int64 `json:"added"`
}

// PostDiscoveryCorpusResponse is generated from the OpenAPI component of the same name.
type PostDiscoveryCorpusResponse struct {
	Received int64 `json:"received"`
	Added    int64 `json:"added"`
}

// GetProspectsResponseDataItem is generated from the OpenAPI component of the same name.
type GetProspectsResponseDataItem struct {
	ID           string                  `json:"id"`
	Name         string                  `json:"name"`
	URL          string                  `json:"url"`
	Channel      string                  `json:"channel"`
	AudienceSize int64                   `json:"audienceSize"`
	Evidence     *GetOpenapiJsonResponse `json:"evidence"`
	FitScore     float64                 `json:"fitScore"`
	ContactEmail *string                 `json:"contactEmail"`
	Status       string                  `json:"status"`
}

// GetProspectsResponse is generated from the OpenAPI component of the same name.
type GetProspectsResponse struct {
	Object     string                         `json:"object"`
	Data       []GetProspectsResponseDataItem `json:"data"`
	HasMore    bool                           `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string                        `json:"next_cursor"`
}

// PostOutreachDraftResponse is generated from the OpenAPI component of the same name.
type PostOutreachDraftResponse struct {
	MessageIds []string `json:"messageIds"`
}

// GetOutreachResponseDataItem is generated from the OpenAPI component of the same name.
type GetOutreachResponseDataItem struct {
	ID         string  `json:"id"`
	ProspectID string  `json:"prospectId"`
	Step       int64   `json:"step"`
	Subject    string  `json:"subject"`
	Status     string  `json:"status"`
	CreatedAt  string  `json:"createdAt"`
	SentAt     *string `json:"sentAt"`
}

// GetOutreachResponse is generated from the OpenAPI component of the same name.
type GetOutreachResponse struct {
	Object     string                        `json:"object"`
	Data       []GetOutreachResponseDataItem `json:"data"`
	HasMore    bool                          `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string                       `json:"next_cursor"`
}

// PostOutreachSendDueResponse is generated from the OpenAPI component of the same name.
type PostOutreachSendDueResponse struct {
	Sent       int64 `json:"sent"`
	Suppressed int64 `json:"suppressed"`
	Capped     bool  `json:"capped"`
	Failed     int64 `json:"failed"`
}

// PostUnsubscribeTokenResponse is generated from the OpenAPI component of the same name.
type PostUnsubscribeTokenResponse struct {
	Ok bool `json:"ok"`
}

// GetAssistantSuggestionsResponseDataItem is generated from the OpenAPI component of the same name.
type GetAssistantSuggestionsResponseDataItem struct {
	Kind            string                     `json:"kind"`
	Message         string                     `json:"message"`
	SuggestedAction string                     `json:"suggestedAction"`
	Data            map[string]json.RawMessage `json:"data"`
}

// GetAssistantSuggestionsResponse is generated from the OpenAPI component of the same name.
type GetAssistantSuggestionsResponse struct {
	Object     string                                    `json:"object"`
	Data       []GetAssistantSuggestionsResponseDataItem `json:"data"`
	HasMore    bool                                      `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string                                   `json:"next_cursor"`
}

// GetMarketplaceListingResponseDataItemStats is generated from the OpenAPI component of the same name.
type GetMarketplaceListingResponseDataItemStats struct {
	Clicks               int64    `json:"clicks"`
	Referrals            int64    `json:"referrals"`
	EpcMinor             int64    `json:"epcMinor"` // Integer minor units (cents).
	ConversionRate       float64  `json:"conversionRate"`
	PayoutReliability    *float64 `json:"payoutReliability"`
	ApprovalLatencyHours *float64 `json:"approvalLatencyHours"`
}

// GetMarketplaceListingResponseDataItem is generated from the OpenAPI component of the same name.
type GetMarketplaceListingResponseDataItem struct {
	ProgramID  string                                     `json:"programId"`
	Visibility string                                     `json:"visibility"`
	Category   string                                     `json:"category"`
	Headline   string                                     `json:"headline"`
	Stats      GetMarketplaceListingResponseDataItemStats `json:"stats"`
}

// GetMarketplaceListingResponse is generated from the OpenAPI component of the same name.
type GetMarketplaceListingResponse struct {
	Object     string                                  `json:"object"`
	Data       []GetMarketplaceListingResponseDataItem `json:"data"`
	HasMore    bool                                    `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string                                 `json:"next_cursor"`
}

// GetPartnerMarketplaceResponseDataItemStats is generated from the OpenAPI component of the same name.
type GetPartnerMarketplaceResponseDataItemStats struct {
	Clicks               int64    `json:"clicks"`
	Referrals            int64    `json:"referrals"`
	EpcMinor             int64    `json:"epcMinor"` // Integer minor units (cents).
	ConversionRate       float64  `json:"conversionRate"`
	PayoutReliability    *float64 `json:"payoutReliability"`
	ApprovalLatencyHours *float64 `json:"approvalLatencyHours"`
}

// GetPartnerMarketplaceResponseDataItem is generated from the OpenAPI component of the same name.
type GetPartnerMarketplaceResponseDataItem struct {
	ProgramID         string                                     `json:"programId"`
	ProgramName       string                                     `json:"programName"`
	Merchant          string                                     `json:"merchant"`
	Category          string                                     `json:"category"`
	Headline          string                                     `json:"headline"`
	Sponsored         bool                                       `json:"sponsored"`
	CookieWindowDays  int64                                      `json:"cookieWindowDays"`
	CommissionSummary *GetOpenapiJsonResponse                    `json:"commissionSummary"`
	Stats             GetPartnerMarketplaceResponseDataItemStats `json:"stats"`
}

// GetPartnerMarketplaceResponse is generated from the OpenAPI component of the same name.
type GetPartnerMarketplaceResponse struct {
	Object     string                                  `json:"object"`
	Data       []GetPartnerMarketplaceResponseDataItem `json:"data"`
	HasMore    bool                                    `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string                                 `json:"next_cursor"`
}

// GetPartnerProfileVerifiedResponse is generated from the OpenAPI component of the same name.
type GetPartnerProfileVerifiedResponse struct {
	Programs        int64    `json:"programs"`
	Clicks          int64    `json:"clicks"`
	Referrals       int64    `json:"referrals"`
	RepeatPayerRate *float64 `json:"repeatPayerRate"`
	Verified        bool     `json:"verified"`
}

// PostPartnerDisputesResponse is generated from the OpenAPI component of the same name.
type PostPartnerDisputesResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

// GetDisputesResponseDataItem is generated from the OpenAPI component of the same name.
type GetDisputesResponseDataItem struct {
	ID           string  `json:"id"`
	PartnerID    string  `json:"partnerId"`
	PartnerEmail string  `json:"partnerEmail"`
	ReferralID   *string `json:"referralId"`
	Kind         string  `json:"kind"`
	Status       string  `json:"status"`
	Detail       string  `json:"detail"`
	Resolution   *string `json:"resolution"`
	CreatedAt    string  `json:"createdAt"`
}

// GetDisputesResponse is generated from the OpenAPI component of the same name.
type GetDisputesResponse struct {
	Object     string                        `json:"object"`
	Data       []GetDisputesResponseDataItem `json:"data"`
	HasMore    bool                          `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string                       `json:"next_cursor"`
}

// GetCreatorsResponse is generated from the OpenAPI component of the same name.
type GetCreatorsResponse struct {
	Object     string           `json:"object"`
	Data       []CreatorProfile `json:"data"`
	HasMore    bool             `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string          `json:"next_cursor"`
}

// PostCreatorsAccountIdInviteResponse is generated from the OpenAPI component of the same name.
type PostCreatorsAccountIdInviteResponse struct {
	ID      string `json:"id"`
	Program string `json:"program"`
}

// GetInvitationsResponseDataItem is generated from the OpenAPI component of the same name.
type GetInvitationsResponseDataItem struct {
	ID          string  `json:"id"`
	CreatorName *string `json:"creatorName"`
	AccountID   string  `json:"accountId"`
	ProgramID   string  `json:"programId"`
	ProgramName string  `json:"programName"`
	Status      string  `json:"status"`
	Message     string  `json:"message"`
	CreatedAt   string  `json:"createdAt"`
	RespondedAt *string `json:"respondedAt"`
}

// GetInvitationsResponse is generated from the OpenAPI component of the same name.
type GetInvitationsResponse struct {
	Object     string                           `json:"object"`
	Data       []GetInvitationsResponseDataItem `json:"data"`
	HasMore    bool                             `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string                          `json:"next_cursor"`
}

// GetPartnerInvitationsResponseDataItem is generated from the OpenAPI component of the same name.
type GetPartnerInvitationsResponseDataItem struct {
	ID          string `json:"id"`
	Merchant    string `json:"merchant"`
	OrgID       string `json:"orgId"`
	ProgramID   string `json:"programId"`
	ProgramName string `json:"programName"`
	Status      string `json:"status"`
	Message     string `json:"message"`
	CreatedAt   string `json:"createdAt"`
}

// GetPartnerInvitationsResponse is generated from the OpenAPI component of the same name.
type GetPartnerInvitationsResponse struct {
	Object     string                                  `json:"object"`
	Data       []GetPartnerInvitationsResponseDataItem `json:"data"`
	HasMore    bool                                    `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string                                 `json:"next_cursor"`
}

// GetBillingPlanResponseUsage is generated from the OpenAPI component of the same name.
type GetBillingPlanResponseUsage struct {
	Programs int64 `json:"programs"`
	Partners int64 `json:"partners"`
}

// GetBillingPlanResponse is generated from the OpenAPI component of the same name.
type GetBillingPlanResponse struct {
	Plan        string                            `json:"plan"`
	PlanStatus  string                            `json:"planStatus"`
	PeriodEnd   *string                           `json:"periodEnd"`
	Limits      GetOpenapiJsonResponse            `json:"limits"`
	Usage       GetBillingPlanResponseUsage       `json:"usage"`
	Catalog     map[string]GetOpenapiJsonResponse `json:"catalog"`
	Subscribed  bool                              `json:"subscribed"`
	CanActivate bool                              `json:"canActivate"`
}

// PostBillingCheckoutResponse is generated from the OpenAPI component of the same name.
type PostBillingCheckoutResponse struct {
	URL       *string `json:"url,omitempty"`
	SessionID *string `json:"sessionId,omitempty"`
}

// PostBillingPortalResponse is generated from the OpenAPI component of the same name.
type PostBillingPortalResponse struct {
	URL string `json:"url"`
}

// PostBillingActivateResponse is generated from the OpenAPI component of the same name.
type PostBillingActivateResponse struct {
	Plan       string `json:"plan"`
	PlanStatus string `json:"planStatus"`
}

// GetProgramsIdRefereeRewardResponse is generated from the OpenAPI component of the same name.
type GetProgramsIdRefereeRewardResponse struct {
	RefereeReward *GetOpenapiJsonResponse `json:"refereeReward"`
}

// PutProgramsIdRefereeRewardResponse is generated from the OpenAPI component of the same name.
type PutProgramsIdRefereeRewardResponse struct {
	Ok            bool                    `json:"ok"`
	RefereeReward *GetOpenapiJsonResponse `json:"refereeReward"`
}

// GetRewardsResponseDataItem is generated from the OpenAPI component of the same name.
type GetRewardsResponseDataItem struct {
	ID          string  `json:"id"`
	ProgramID   string  `json:"programId"`
	ReferralID  string  `json:"referralId"`
	CustomerRef string  `json:"customerRef"`
	Kind        string  `json:"kind"`
	AmountMinor int64   `json:"amountMinor"` // Integer minor units (cents).
	Currency    string  `json:"currency"`
	Code        *string `json:"code"`
	Status      string  `json:"status"`
	CreatedAt   string  `json:"createdAt"`
	RedeemedAt  *string `json:"redeemedAt"`
}

// GetRewardsResponse is generated from the OpenAPI component of the same name.
type GetRewardsResponse struct {
	Object     string                       `json:"object"`
	Data       []GetRewardsResponseDataItem `json:"data"`
	HasMore    bool                         `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string                      `json:"next_cursor"`
}

// GetRewardsLookupCodeResponse is generated from the OpenAPI component of the same name.
type GetRewardsLookupCodeResponse struct {
	Valid       bool    `json:"valid"`
	RewardID    *string `json:"rewardId,omitempty"`
	Kind        *string `json:"kind,omitempty"`
	AmountMinor *int64  `json:"amountMinor,omitempty"` // Integer minor units (cents).
	Currency    *string `json:"currency,omitempty"`
	Status      *string `json:"status,omitempty"`
}

// PostRewardsRedeemByCodeResponse is generated from the OpenAPI component of the same name.
type PostRewardsRedeemByCodeResponse struct {
	Ok             bool   `json:"ok"`
	RewardID       string `json:"rewardId"`
	GrantID        string `json:"grantId"`
	Status         string `json:"status"`
	Currency       string `json:"currency"`
	RequestedMinor int64  `json:"requestedMinor"` // Integer minor units (cents).
	RedeemedMinor  int64  `json:"redeemedMinor"`  // Amount recorded, capped at the remaining balance.
	RemainingMinor int64  `json:"remainingMinor"` // Integer minor units (cents).
	Capped         bool   `json:"capped"`         // True when the request exceeded the remaining balance.
}

// PostRewardsResponse is generated from the OpenAPI component of the same name.
type PostRewardsResponse struct {
	ID             string  `json:"id"`
	PartnerID      string  `json:"partnerId"`
	ProgramID      string  `json:"programId"`
	RuleKey        *string `json:"ruleKey"`
	Reason         string  `json:"reason"`
	AmountMinor    int64   `json:"amountMinor"` // Integer minor units (cents).
	Currency       string  `json:"currency"`
	CommissionID   *string `json:"commissionId"`
	IdempotencyKey string  `json:"idempotencyKey"`
	CreatedAt      string  `json:"createdAt"`
	Replayed       bool    `json:"replayed"`
}

// GetRewardsGrantedResponseValueItem is generated from the OpenAPI component of the same name.
type GetRewardsGrantedResponseValueItem struct {
	ID             string  `json:"id"`
	PartnerID      string  `json:"partnerId"`
	ProgramID      string  `json:"programId"`
	RuleKey        *string `json:"ruleKey"`
	Reason         string  `json:"reason"`
	AmountMinor    int64   `json:"amountMinor"` // Integer minor units (cents).
	Currency       string  `json:"currency"`
	CommissionID   *string `json:"commissionId"`
	IdempotencyKey string  `json:"idempotencyKey"`
	CreatedAt      string  `json:"createdAt"`
}

// GetRewardsGrantedResponse is generated from the OpenAPI component of the same name.
type GetRewardsGrantedResponse = []GetRewardsGrantedResponseValueItem

// PutRewardsRulesKeyResponse is generated from the OpenAPI component of the same name.
type PutRewardsRulesKeyResponse struct {
	Key         string  `json:"key"`
	Event       *string `json:"event"`
	ProgramID   *string `json:"programId"`
	AmountMinor int64   `json:"amountMinor"` // Integer minor units (cents).
	Currency    string  `json:"currency"`
	Reason      string  `json:"reason"`
	Enabled     bool    `json:"enabled"`
	UpdatedAt   string  `json:"updatedAt"`
}

// DeleteRewardsRulesKeyResponse is generated from the OpenAPI component of the same name.
type DeleteRewardsRulesKeyResponse struct {
	Ok bool `json:"ok"`
}

// GetProgramsIdGroupsResponse is generated from the OpenAPI component of the same name.
type GetProgramsIdGroupsResponse struct {
	Object     string         `json:"object"`
	Data       []PartnerGroup `json:"data"`
	HasMore    bool           `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string        `json:"next_cursor"`
}

// PutProgramsIdPartnersPartnerIdGroupResponse is generated from the OpenAPI component of the same name.
type PutProgramsIdPartnersPartnerIdGroupResponse struct {
	PartnerID string  `json:"partnerId"`
	ProgramID string  `json:"programId"`
	GroupID   *string `json:"groupId"`
}

// GetProgramsIdPromotionsResponse is generated from the OpenAPI component of the same name.
type GetProgramsIdPromotionsResponse struct {
	Object     string      `json:"object"`
	Data       []Promotion `json:"data"`
	HasMore    bool        `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string     `json:"next_cursor"`
}

// GetPartnerGroupsResponseDataItem is generated from the OpenAPI component of the same name.
type GetPartnerGroupsResponseDataItem struct {
	GroupID             string   `json:"groupId"`
	GroupName           string   `json:"groupName"`
	OrgName             string   `json:"orgName"`
	ProgramID           string   `json:"programId"`
	ProgramName         string   `json:"programName"`
	OnboardingMessage   *string  `json:"onboardingMessage"`
	OnboardingChecklist []string `json:"onboardingChecklist"`
	HoldDays            *int64   `json:"holdDays"`
	MinThresholdMinor   *int64   `json:"minThresholdMinor"`
	PayoutCurrency      *string  `json:"payoutCurrency"`
}

// GetPartnerGroupsResponse is generated from the OpenAPI component of the same name.
type GetPartnerGroupsResponse struct {
	Object     string                             `json:"object"`
	Data       []GetPartnerGroupsResponseDataItem `json:"data"`
	HasMore    bool                               `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string                            `json:"next_cursor"`
}

// PostContractsResponse is generated from the OpenAPI component of the same name.
type PostContractsResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

// GetContractsResponseDataItem is generated from the OpenAPI component of the same name.
type GetContractsResponseDataItem struct {
	ID           string                 `json:"id"`
	ProgramID    string                 `json:"programId"`
	PartnerID    string                 `json:"partnerId"`
	Title        string                 `json:"title"`
	Compensation GetOpenapiJsonResponse `json:"compensation"`
	Status       string                 `json:"status"`
	CreatedAt    string                 `json:"createdAt"`
	RespondedAt  *string                `json:"respondedAt"`
}

// GetContractsResponse is generated from the OpenAPI component of the same name.
type GetContractsResponse struct {
	Object     string                         `json:"object"`
	Data       []GetContractsResponseDataItem `json:"data"`
	HasMore    bool                           `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string                        `json:"next_cursor"`
}

// GetPartnerContractsResponseDataItem is generated from the OpenAPI component of the same name.
type GetPartnerContractsResponseDataItem struct {
	ID           string                 `json:"id"`
	Merchant     string                 `json:"merchant"`
	Program      string                 `json:"program"`
	Title        string                 `json:"title"`
	TermsMd      string                 `json:"termsMd"`
	Compensation GetOpenapiJsonResponse `json:"compensation"`
	Status       string                 `json:"status"`
	CreatedAt    string                 `json:"createdAt"`
}

// GetPartnerContractsResponse is generated from the OpenAPI component of the same name.
type GetPartnerContractsResponse struct {
	Object     string                                `json:"object"`
	Data       []GetPartnerContractsResponseDataItem `json:"data"`
	HasMore    bool                                  `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string                               `json:"next_cursor"`
}

// PutProgramsIdIncrementalityResponse is generated from the OpenAPI component of the same name.
type PutProgramsIdIncrementalityResponse struct {
	Ok         bool  `json:"ok"`
	HoldoutBps int64 `json:"holdoutBps"`
}

// GetIncrementalityResponseTreatment is generated from the OpenAPI component of the same name.
type GetIncrementalityResponseTreatment struct {
	Referrals int64   `json:"referrals"`
	Converted int64   `json:"converted"`
	Rate      float64 `json:"rate"`
}

// GetIncrementalityResponseControl is generated from the OpenAPI component of the same name.
type GetIncrementalityResponseControl struct {
	Referrals int64   `json:"referrals"`
	Converted int64   `json:"converted"`
	Rate      float64 `json:"rate"`
}

// GetIncrementalityResponseByPartnerItem is generated from the OpenAPI component of the same name.
type GetIncrementalityResponseByPartnerItem struct {
	PartnerID          string  `json:"partnerId"`
	TreatmentReferrals int64   `json:"treatmentReferrals"`
	ControlReferrals   int64   `json:"controlReferrals"`
	Lift               float64 `json:"lift"`
}

// GetIncrementalityResponse is generated from the OpenAPI component of the same name.
type GetIncrementalityResponse struct {
	ProgramID  string                                   `json:"programId"`
	HoldoutBps int64                                    `json:"holdoutBps"`
	Treatment  GetIncrementalityResponseTreatment       `json:"treatment"`
	Control    GetIncrementalityResponseControl         `json:"control"`
	Lift       float64                                  `json:"lift"`
	ByPartner  []GetIncrementalityResponseByPartnerItem `json:"byPartner"`
}

// PutProgramsIdAutopricingResponse is generated from the OpenAPI component of the same name.
type PutProgramsIdAutopricingResponse struct {
	Ok          bool                    `json:"ok"`
	Autopricing *GetOpenapiJsonResponse `json:"autopricing"`
}

// GetProgramsIdAutopricingResponseAutopricingOption1 is generated from the OpenAPI component of the same name.
type GetProgramsIdAutopricingResponseAutopricingOption1 struct {
	Enabled    bool    `json:"enabled"`
	TargetRoas float64 `json:"targetRoas"`
	MinBps     int64   `json:"minBps"`
	MaxBps     int64   `json:"maxBps"`
	Objective  string  `json:"objective"`
	TargetRoi  float64 `json:"targetRoi"`
}

// GetProgramsIdAutopricingResponse is generated from the OpenAPI component of the same name.
type GetProgramsIdAutopricingResponse struct {
	Autopricing *GetProgramsIdAutopricingResponseAutopricingOption1 `json:"autopricing"` // Null when self-pricing is off. Configs stored before objectives existed read as roas / targetRoi 2.
}

// PostProgramsIdAutopriceRunResponseChangesItem is generated from the OpenAPI component of the same name.
type PostProgramsIdAutopriceRunResponseChangesItem struct {
	PartnerID    string   `json:"partnerId"`
	FromBps      int64    `json:"fromBps"`
	ToBps        int64    `json:"toBps"`
	RealizedRoas *float64 `json:"realizedRoas"`          // Revenue ÷ commission; null under objective roi.
	RealizedRoi  *float64 `json:"realizedRoi,omitempty"` // Objective roi only: referral ROI over matured referrals, both sides: profit after ÷ (referee credit cost + cash commissions + referrer credit cost).
}

// PostProgramsIdAutopriceRunResponseHeldItem is generated from the OpenAPI component of the same name.
type PostProgramsIdAutopriceRunResponseHeldItem struct {
	PartnerID     string `json:"partnerId"`
	Reason        string `json:"reason"`
	MaturedGrants int64  `json:"maturedGrants"` // The partner's matured referrals in the program (the name predates referral-based ROI).
}

// PostProgramsIdAutopriceRunResponse is generated from the OpenAPI component of the same name.
type PostProgramsIdAutopriceRunResponse struct {
	Objective  string                                          `json:"objective"`
	TargetRoas float64                                         `json:"targetRoas"`
	TargetRoi  *float64                                        `json:"targetRoi,omitempty"` // Objective roi only.
	Changes    []PostProgramsIdAutopriceRunResponseChangesItem `json:"changes"`
	Held       []PostProgramsIdAutopriceRunResponseHeldItem    `json:"held,omitempty"` // Objective roi only: partners whose rate was kept, and why.
}

// GetPassportPublicKeyResponse is generated from the OpenAPI component of the same name.
type GetPassportPublicKeyResponse struct {
	Algorithm string `json:"algorithm"`
	PublicKey string `json:"publicKey"`
}

// PostPassportVerifyResponse is generated from the OpenAPI component of the same name.
type PostPassportVerifyResponse struct {
	Valid  bool                    `json:"valid"`
	Claims *GetOpenapiJsonResponse `json:"claims,omitempty"`
	Reason *string                 `json:"reason,omitempty"`
}

// GetPartnerAdvanceOfferResponse is generated from the OpenAPI component of the same name.
type GetPartnerAdvanceOfferResponse struct {
	Currency            string  `json:"currency"`
	MonthlyRunRateMinor int64   `json:"monthlyRunRateMinor"` // Integer minor units (cents).
	ProjectedMinor      int64   `json:"projectedMinor"`      // Integer minor units (cents).
	MaxAdvanceMinor     int64   `json:"maxAdvanceMinor"`     // Integer minor units (cents).
	FeeBps              int64   `json:"feeBps"`
	QualityScore        float64 `json:"qualityScore"`
}

// PostPartnerAdvanceAcceptResponse is generated from the OpenAPI component of the same name.
type PostPartnerAdvanceAcceptResponse struct {
	ID               string `json:"id"`
	PrincipalMinor   int64  `json:"principalMinor"`   // Integer minor units (cents).
	FeeMinor         int64  `json:"feeMinor"`         // Integer minor units (cents).
	OutstandingMinor int64  `json:"outstandingMinor"` // Integer minor units (cents).
	Status           string `json:"status"`
}

// GetPartnerAdvanceResponseDataItem is generated from the OpenAPI component of the same name.
type GetPartnerAdvanceResponseDataItem struct {
	ID               string  `json:"id"`
	Currency         string  `json:"currency"`
	PrincipalMinor   int64   `json:"principalMinor"`   // Integer minor units (cents).
	FeeMinor         int64   `json:"feeMinor"`         // Integer minor units (cents).
	OutstandingMinor int64   `json:"outstandingMinor"` // Integer minor units (cents).
	Status           string  `json:"status"`
	CreatedAt        string  `json:"createdAt"`
	RepaidAt         *string `json:"repaidAt"`
}

// GetPartnerAdvanceResponse is generated from the OpenAPI component of the same name.
type GetPartnerAdvanceResponse struct {
	Object     string                              `json:"object"`
	Data       []GetPartnerAdvanceResponseDataItem `json:"data"`
	HasMore    bool                                `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string                             `json:"next_cursor"`
}

// PutNetworkSettingsResponse is generated from the OpenAPI component of the same name.
type PutNetworkSettingsResponse struct {
	Enabled bool `json:"enabled"`
}

// GetNetworkSettingsResponse is generated from the OpenAPI component of the same name.
type GetNetworkSettingsResponse struct {
	Enabled bool `json:"enabled"`
}

// GetNetworkJourneyResponseTouchpointsItem is generated from the OpenAPI component of the same name.
type GetNetworkJourneyResponseTouchpointsItem struct {
	PartnerAccountID string  `json:"partnerAccountId"`
	Ts               string  `json:"ts"`
	Credit           float64 `json:"credit"`
}

// GetNetworkJourneyResponse is generated from the OpenAPI component of the same name.
type GetNetworkJourneyResponse struct {
	EmailHash   string                                     `json:"emailHash"`
	Model       string                                     `json:"model"`
	Touchpoints []GetNetworkJourneyResponseTouchpointsItem `json:"touchpoints"`
	Partners    int64                                      `json:"partners"`
}

// PutProgramsIdNegotiationPolicyResponse is generated from the OpenAPI component of the same name.
type PutProgramsIdNegotiationPolicyResponse struct {
	Ok                bool                    `json:"ok"`
	NegotiationPolicy *GetOpenapiJsonResponse `json:"negotiationPolicy"`
}

// GetProgramsIdNegotiationPolicyResponse is generated from the OpenAPI component of the same name.
type GetProgramsIdNegotiationPolicyResponse struct {
	NegotiationPolicy *GetOpenapiJsonResponse `json:"negotiationPolicy"`
}

// GetNegotiationsResponseDataItem is generated from the OpenAPI component of the same name.
type GetNegotiationsResponseDataItem struct {
	ID             string `json:"id"`
	ProgramID      string `json:"programId"`
	AgentAccountID string `json:"agentAccountId"`
	ProposedBps    int64  `json:"proposedBps"`
	Decision       string `json:"decision"`
	RateBps        *int64 `json:"rateBps"`
	Round          int64  `json:"round"`
	CreatedAt      string `json:"createdAt"`
}

// GetNegotiationsResponse is generated from the OpenAPI component of the same name.
type GetNegotiationsResponse struct {
	Object     string                            `json:"object"`
	Data       []GetNegotiationsResponseDataItem `json:"data"`
	HasMore    bool                              `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string                           `json:"next_cursor"`
}

// GetCleanroomConsentResponseConsentOption1 is generated from the OpenAPI component of the same name.
type GetCleanroomConsentResponseConsentOption1 struct {
	Granted   bool                   `json:"granted"`
	Purposes  []string               `json:"purposes"`
	Declared  GetOpenapiJsonResponse `json:"declared"`
	EmailHash *string                `json:"emailHash"`
	UpdatedAt string                 `json:"updatedAt"`
}

// GetCleanroomConsentResponse is generated from the OpenAPI component of the same name.
type GetCleanroomConsentResponse struct {
	CustomerRef string                                     `json:"customerRef"`
	Consent     *GetCleanroomConsentResponseConsentOption1 `json:"consent"`
}

// GetCleanroomQueryResponseCohortsItem is generated from the OpenAPI component of the same name.
type GetCleanroomQueryResponseCohortsItem struct {
	Key   string `json:"key"`
	Count int64  `json:"count"`
}

// GetCleanroomQueryResponse is generated from the OpenAPI component of the same name.
type GetCleanroomQueryResponse struct {
	Dimension         string                                 `json:"dimension"`
	K                 int64                                  `json:"k"`
	Cohorts           []GetCleanroomQueryResponseCohortsItem `json:"cohorts"`
	SuppressedCohorts int64                                  `json:"suppressedCohorts"`
	SuppressedRows    int64                                  `json:"suppressedRows"`
}

// GetAgentsResponseDataItem is generated from the OpenAPI component of the same name.
type GetAgentsResponseDataItem struct {
	PartnerID       string                  `json:"partnerId"`
	AccountID       string                  `json:"accountId"`
	Name            string                  `json:"name"`
	Wallet          *string                 `json:"wallet"`
	Metadata        *GetOpenapiJsonResponse `json:"metadata"`
	InstantPayout   bool                    `json:"instantPayout"`
	Status          string                  `json:"status"`
	Recommendations int64                   `json:"recommendations"`
}

// GetAgentsResponse is generated from the OpenAPI component of the same name.
type GetAgentsResponse struct {
	Object     string                      `json:"object"`
	Data       []GetAgentsResponseDataItem `json:"data"`
	HasMore    bool                        `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string                     `json:"next_cursor"`
}

// GetAgentsRecommendationsResponseDataItem is generated from the OpenAPI component of the same name.
type GetAgentsRecommendationsResponseDataItem struct {
	ID          string  `json:"id"`
	ProgramID   string  `json:"programId"`
	CustomerRef string  `json:"customerRef"`
	Agent       string  `json:"agent"`
	Wallet      *string `json:"wallet"`
	CreatedAt   string  `json:"createdAt"`
}

// GetAgentsRecommendationsResponse is generated from the OpenAPI component of the same name.
type GetAgentsRecommendationsResponse struct {
	Object     string                                     `json:"object"`
	Data       []GetAgentsRecommendationsResponseDataItem `json:"data"`
	HasMore    bool                                       `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string                                    `json:"next_cursor"`
}

// GetAdminMetricsResponse is generated from the OpenAPI component of the same name.
type GetAdminMetricsResponse struct {
	Orgs                  int64 `json:"orgs"`
	SuspendedOrgs         int64 `json:"suspendedOrgs"`
	PartnerAccounts       int64 `json:"partnerAccounts"`
	BannedPartners        int64 `json:"bannedPartners"`
	Programs              int64 `json:"programs"`
	CommissionEarnedMinor int64 `json:"commissionEarnedMinor"` // Integer minor units (cents).
	PayoutsSent           int64 `json:"payoutsSent"`
	OpenFraudCases        int64 `json:"openFraudCases"`
}

// GetAdminOrgsResponseDataItem is generated from the OpenAPI component of the same name.
type GetAdminOrgsResponseDataItem struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Plan       string  `json:"plan"`
	PlanStatus *string `json:"planStatus"`
	Suspended  bool    `json:"suspended"`
	CreatedAt  string  `json:"createdAt"`
	Programs   int64   `json:"programs"`
	Partners   int64   `json:"partners"`
}

// GetAdminOrgsResponse is generated from the OpenAPI component of the same name.
type GetAdminOrgsResponse struct {
	Object     string                         `json:"object"`
	Data       []GetAdminOrgsResponseDataItem `json:"data"`
	HasMore    bool                           `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string                        `json:"next_cursor"`
}

// PostAdminOrgsIdSuspendResponse is generated from the OpenAPI component of the same name.
type PostAdminOrgsIdSuspendResponse struct {
	Ok        bool `json:"ok"`
	Suspended bool `json:"suspended"`
}

// PostAdminOrgsIdPlanResponse is generated from the OpenAPI component of the same name.
type PostAdminOrgsIdPlanResponse struct {
	Plan       string `json:"plan"`
	PlanStatus string `json:"planStatus"`
}

// PostAdminOrgsIdUnsuspendResponse is generated from the OpenAPI component of the same name.
type PostAdminOrgsIdUnsuspendResponse struct {
	Ok        bool `json:"ok"`
	Suspended bool `json:"suspended"`
}

// GetAdminPartnerAccountsResponseDataItem is generated from the OpenAPI component of the same name.
type GetAdminPartnerAccountsResponseDataItem struct {
	ID              string  `json:"id"`
	Email           string  `json:"email"`
	Name            *string `json:"name"`
	Banned          bool    `json:"banned"`
	KYCStatus       *string `json:"kycStatus"`
	SanctionsStatus *string `json:"sanctionsStatus"`
	CreatedAt       string  `json:"createdAt"`
}

// GetAdminPartnerAccountsResponse is generated from the OpenAPI component of the same name.
type GetAdminPartnerAccountsResponse struct {
	Object     string                                    `json:"object"`
	Data       []GetAdminPartnerAccountsResponseDataItem `json:"data"`
	HasMore    bool                                      `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string                                   `json:"next_cursor"`
}

// PostAdminPartnerAccountsIdBanResponse is generated from the OpenAPI component of the same name.
type PostAdminPartnerAccountsIdBanResponse struct {
	Ok     bool `json:"ok"`
	Banned bool `json:"banned"`
}

// PostAdminPartnerAccountsIdUnbanResponse is generated from the OpenAPI component of the same name.
type PostAdminPartnerAccountsIdUnbanResponse struct {
	Ok     bool `json:"ok"`
	Banned bool `json:"banned"`
}

// GetNotificationsResponse is generated from the OpenAPI component of the same name.
type GetNotificationsResponse struct {
	Object     string         `json:"object"`
	Data       []Notification `json:"data"`
	HasMore    bool           `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string        `json:"next_cursor"`
}

// GetPartnerNotificationsResponse is generated from the OpenAPI component of the same name.
type GetPartnerNotificationsResponse struct {
	Object     string         `json:"object"`
	Data       []Notification `json:"data"`
	HasMore    bool           `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string        `json:"next_cursor"`
}

// PostAssetsResponse is generated from the OpenAPI component of the same name.
type PostAssetsResponse struct {
	ID         string  `json:"id"`
	ProgramID  string  `json:"programId"`
	Kind       string  `json:"kind"`
	Title      string  `json:"title"`
	Content    string  `json:"content"`
	Dimensions string  `json:"dimensions"`
	GroupID    *string `json:"groupId,omitempty"`
}

// GetAssetsResponseDataItem is generated from the OpenAPI component of the same name.
type GetAssetsResponseDataItem struct {
	ID         string  `json:"id"`
	ProgramID  string  `json:"programId"`
	Kind       string  `json:"kind"`
	Title      string  `json:"title"`
	Content    string  `json:"content"`
	Dimensions string  `json:"dimensions"`
	GroupID    *string `json:"groupId"`
	Clicks     int64   `json:"clicks"`
	Referrals  int64   `json:"referrals"`
	CreatedAt  string  `json:"createdAt"`
}

// GetAssetsResponse is generated from the OpenAPI component of the same name.
type GetAssetsResponse struct {
	Object     string                      `json:"object"`
	Data       []GetAssetsResponseDataItem `json:"data"`
	HasMore    bool                        `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string                     `json:"next_cursor"`
}

// PostPartnerDealsResponse is generated from the OpenAPI component of the same name.
type PostPartnerDealsResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

// GetPartnerDealsResponseDataItem is generated from the OpenAPI component of the same name.
type GetPartnerDealsResponseDataItem struct {
	ID                  string  `json:"id"`
	OrgID               string  `json:"orgId"`
	Company             string  `json:"company"`
	ValueMinor          int64   `json:"valueMinor"` // Integer minor units (cents).
	Currency            string  `json:"currency"`
	Status              string  `json:"status"`
	ProtectionExpiresAt *string `json:"protectionExpiresAt"`
	CreatedAt           string  `json:"createdAt"`
}

// GetPartnerDealsResponse is generated from the OpenAPI component of the same name.
type GetPartnerDealsResponse struct {
	Object     string                            `json:"object"`
	Data       []GetPartnerDealsResponseDataItem `json:"data"`
	HasMore    bool                              `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string                           `json:"next_cursor"`
}

// GetDealsResponseDataItem is generated from the OpenAPI component of the same name.
type GetDealsResponseDataItem struct {
	ID                  string  `json:"id"`
	ProgramID           string  `json:"programId"`
	PartnerID           string  `json:"partnerId"`
	PartnerEmail        string  `json:"partnerEmail"`
	Company             string  `json:"company"`
	ContactEmail        string  `json:"contactEmail"`
	ValueMinor          int64   `json:"valueMinor"` // Integer minor units (cents).
	Currency            string  `json:"currency"`
	Status              string  `json:"status"`
	ProtectionExpiresAt *string `json:"protectionExpiresAt"`
	ReferralID          *string `json:"referralId"`
	Notes               string  `json:"notes"`
	CreatedAt           string  `json:"createdAt"`
}

// GetDealsResponse is generated from the OpenAPI component of the same name.
type GetDealsResponse struct {
	Object     string                     `json:"object"`
	Data       []GetDealsResponseDataItem `json:"data"`
	HasMore    bool                       `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string                    `json:"next_cursor"`
}

// PostDealsIdApproveResponse is generated from the OpenAPI component of the same name.
type PostDealsIdApproveResponse struct {
	ID                  string `json:"id"`
	Status              string `json:"status"`
	ProtectionExpiresAt string `json:"protectionExpiresAt"`
}

// PostDealsIdRejectResponse is generated from the OpenAPI component of the same name.
type PostDealsIdRejectResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

// PostDealsIdWonResponse is generated from the OpenAPI component of the same name.
type PostDealsIdWonResponse struct {
	ID         string `json:"id"`
	Status     string `json:"status"`
	ReferralID string `json:"referralId"`
}

// PostProgramsIdApplyLevelingResponse is generated from the OpenAPI component of the same name.
type PostProgramsIdApplyLevelingResponse struct {
	Evaluated int64 `json:"evaluated"`
	Promoted  int64 `json:"promoted"`
}

// PostPayoutRunsRunScheduledResponseValueItem is generated from the OpenAPI component of the same name.
type PostPayoutRunsRunScheduledResponseValueItem struct {
	OrgID   string `json:"orgId"`
	RunID   string `json:"runId"`
	Payouts int64  `json:"payouts"`
	Sent    int64  `json:"sent"`
	Failed  int64  `json:"failed"`
}

// PostPayoutRunsRunScheduledResponse is generated from the OpenAPI component of the same name.
type PostPayoutRunsRunScheduledResponse = []PostPayoutRunsRunScheduledResponseValueItem

// PutPartnerOrgsOrgIdPostbackResponse is generated from the OpenAPI component of the same name.
type PutPartnerOrgsOrgIdPostbackResponse struct {
	URL    string   `json:"url"`
	Secret string   `json:"secret"`
	Macros []string `json:"macros"`
}

// GetPartnerPostbackLogsResponseDataItem is generated from the OpenAPI component of the same name.
type GetPartnerPostbackLogsResponseDataItem struct {
	Topic      string `json:"topic"`
	URL        string `json:"url"`
	StatusCode *int64 `json:"statusCode"`
	Ts         string `json:"ts"`
}

// GetPartnerPostbackLogsResponse is generated from the OpenAPI component of the same name.
type GetPartnerPostbackLogsResponse struct {
	Object     string                                   `json:"object"`
	Data       []GetPartnerPostbackLogsResponseDataItem `json:"data"`
	HasMore    bool                                     `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string                                  `json:"next_cursor"`
}

// GetPartnerRecruitResponse is generated from the OpenAPI component of the same name.
type GetPartnerRecruitResponse struct {
	Code string `json:"code"`
	URL  string `json:"url"`
}

// GetPartnerAchievementsResponseDataItem is generated from the OpenAPI component of the same name.
type GetPartnerAchievementsResponseDataItem struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Unlocked bool   `json:"unlocked"`
	Progress int64  `json:"progress"`
	Target   int64  `json:"target"`
}

// GetPartnerAchievementsResponse is generated from the OpenAPI component of the same name.
type GetPartnerAchievementsResponse struct {
	Object     string                                   `json:"object"`
	Data       []GetPartnerAchievementsResponseDataItem `json:"data"`
	HasMore    bool                                     `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string                                  `json:"next_cursor"`
}

// GetPartnerTermsResponseDataItem is generated from the OpenAPI component of the same name.
type GetPartnerTermsResponseDataItem struct {
	VersionID   string `json:"versionId"`
	ProgramID   string `json:"programId"`
	ProgramName string `json:"programName"`
	Version     int64  `json:"version"`
	TermsMd     string `json:"termsMd"`
	Accepted    bool   `json:"accepted"`
}

// GetPartnerTermsResponse is generated from the OpenAPI component of the same name.
type GetPartnerTermsResponse struct {
	Object     string                            `json:"object"`
	Data       []GetPartnerTermsResponseDataItem `json:"data"`
	HasMore    bool                              `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string                           `json:"next_cursor"`
}

// GetProgramsIdTermsAcceptancesResponseDataItem is generated from the OpenAPI component of the same name.
type GetProgramsIdTermsAcceptancesResponseDataItem struct {
	Version      int64  `json:"version"`
	PartnerEmail string `json:"partnerEmail"`
	AcceptedAt   string `json:"acceptedAt"`
}

// GetProgramsIdTermsAcceptancesResponse is generated from the OpenAPI component of the same name.
type GetProgramsIdTermsAcceptancesResponse struct {
	Object     string                                          `json:"object"`
	Data       []GetProgramsIdTermsAcceptancesResponseDataItem `json:"data"`
	HasMore    bool                                            `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string                                         `json:"next_cursor"`
}

// GetSetupStatusResponse is generated from the OpenAPI component of the same name.
type GetSetupStatusResponse struct {
	Checks   map[string]bool `json:"checks"`
	Hints    []string        `json:"hints"`
	HintKeys []string        `json:"hintKeys"` // Stable keys for each hint, in the same order, for localized clients.
	Progress string          `json:"progress"`
}

// PostAnomaliesScanResponseValueItem is generated from the OpenAPI component of the same name.
type PostAnomaliesScanResponseValueItem struct {
	OrgID      string  `json:"orgId"`
	Metric     string  `json:"metric"`
	Kind       string  `json:"kind"`
	Last24h    float64 `json:"last24h"`
	Daily7dAvg float64 `json:"daily7dAvg"`
}

// PostAnomaliesScanResponse is generated from the OpenAPI component of the same name.
type PostAnomaliesScanResponse = []PostAnomaliesScanResponseValueItem

// PostPrivacyForgetResponsePartnersItem is generated from the OpenAPI component of the same name.
type PostPrivacyForgetResponsePartnersItem struct {
	PartnerID        string `json:"partnerId"`
	AlreadyForgotten bool   `json:"alreadyForgotten"`
}

// PostPrivacyForgetResponseCustomersItem is generated from the OpenAPI component of the same name.
type PostPrivacyForgetResponseCustomersItem struct {
	CustomerRef string `json:"customerRef"`
	Pseudonym   string `json:"pseudonym"`
}

// PostPrivacyForgetResponse is generated from the OpenAPI component of the same name.
type PostPrivacyForgetResponse struct {
	Ok        bool                                     `json:"ok"`
	Pseudonym *string                                  `json:"pseudonym,omitempty"` // Erasure by customerRef.
	Partners  []PostPrivacyForgetResponsePartnersItem  `json:"partners,omitempty"`
	Customers []PostPrivacyForgetResponseCustomersItem `json:"customers,omitempty"`
}

// PostPrivacyLookupResponse is generated from the OpenAPI component of the same name.
type PostPrivacyLookupResponse struct {
	Found        bool     `json:"found"`
	CustomerRefs []string `json:"customerRefs"`
	PartnerID    *string  `json:"partnerId"`
}

// PostPrivacyRequestsResponse is generated from the OpenAPI component of the same name.
type PostPrivacyRequestsResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

// GetPrivacyRequestsResponse is generated from the OpenAPI component of the same name.
type GetPrivacyRequestsResponse struct {
	Object     string           `json:"object"`
	Data       []PrivacyRequest `json:"data"`
	HasMore    bool             `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string          `json:"next_cursor"`
}

// GetPrivacyRequestsIdResultResponseCustomersItemOption2 is generated from the OpenAPI component of the same name.
type GetPrivacyRequestsIdResultResponseCustomersItemOption2 struct {
	CustomerRef string `json:"customerRef"`
	Pseudonym   string `json:"pseudonym"`
}

// GetPrivacyRequestsIdResultResponsePartnersItem is generated from the OpenAPI component of the same name.
type GetPrivacyRequestsIdResultResponsePartnersItem struct {
	PartnerID        string `json:"partnerId"`
	AlreadyForgotten bool   `json:"alreadyForgotten"`
}

// GetPrivacyRequestsIdResultResponse is generated from the OpenAPI component of the same name.
type GetPrivacyRequestsIdResultResponse struct {
	Found     bool                                             `json:"found"`
	Customers []json.RawMessage                                `json:"customers"`
	Partner   *PartnerExport                                   `json:"partner,omitempty"`
	Partners  []GetPrivacyRequestsIdResultResponsePartnersItem `json:"partners,omitempty"`
}

// PostEmailSettingsVerifyDomainResponse is generated from the OpenAPI component of the same name.
type PostEmailSettingsVerifyDomainResponse = any

// GetEmailTemplatesResponse is generated from the OpenAPI component of the same name.
type GetEmailTemplatesResponse struct {
	Object     string               `json:"object"`
	Data       []EmailTemplateEvent `json:"data"`
	HasMore    bool                 `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string              `json:"next_cursor"`
}

// PutEmailTemplatesEventLocaleResponse is generated from the OpenAPI component of the same name.
type PutEmailTemplatesEventLocaleResponse struct {
	Event     string  `json:"event"`
	Locale    string  `json:"locale"`
	Subject   string  `json:"subject"`
	Text      string  `json:"text"`
	HTML      *string `json:"html"`
	UpdatedAt string  `json:"updatedAt"`
}

// DeleteEmailTemplatesEventLocaleResponse is generated from the OpenAPI component of the same name.
type DeleteEmailTemplatesEventLocaleResponse struct {
	Ok    bool `json:"ok"`
	Reset bool `json:"reset"`
}

// PostEmailTestResponse is generated from the OpenAPI component of the same name.
type PostEmailTestResponse struct {
	Sent     bool   `json:"sent"`
	ID       string `json:"id"`
	Provider string `json:"provider"`
}

// PutMeLocaleResponse is generated from the OpenAPI component of the same name.
type PutMeLocaleResponse struct {
	Locale string `json:"locale"`
}

// PutMeRoiDigestResponse is generated from the OpenAPI component of the same name.
type PutMeRoiDigestResponse struct {
	Enabled bool `json:"enabled"`
}

// PutPartnerLocaleResponse is generated from the OpenAPI component of the same name.
type PutPartnerLocaleResponse struct {
	Locale string `json:"locale"`
}

// GetIntegrationsResponseZapier is generated from the OpenAPI component of the same name.
type GetIntegrationsResponseZapier struct {
	Hooks  int64    `json:"hooks"`
	Events []string `json:"events"`
}

// GetIntegrationsResponse is generated from the OpenAPI component of the same name.
type GetIntegrationsResponse struct {
	Zapier  GetIntegrationsResponseZapier `json:"zapier"`
	Hubspot *IntegrationConnection        `json:"hubspot"`
	Segment *IntegrationConnection        `json:"segment"`
}

// PostIntegrationsHubspotTestResponse is generated from the OpenAPI component of the same name.
type PostIntegrationsHubspotTestResponse struct {
	Ok                bool     `json:"ok"`
	PortalID          *int64   `json:"portalId"`
	CreatedProperties []string `json:"createdProperties"`
	Error             *string  `json:"error"`
}

// PostIntegrationsSegmentTestResponse is generated from the OpenAPI component of the same name.
type PostIntegrationsSegmentTestResponse struct {
	Ok    bool    `json:"ok"`
	Error *string `json:"error"`
}

// GetIntegrationsDeliveriesResponse is generated from the OpenAPI component of the same name.
type GetIntegrationsDeliveriesResponse struct {
	Object     string                `json:"object"`
	Data       []IntegrationDelivery `json:"data"`
	HasMore    bool                  `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string               `json:"next_cursor"`
}

// GetZapierEventsResponseValueItem is generated from the OpenAPI component of the same name.
type GetZapierEventsResponseValueItem struct {
	Key string `json:"key"`
}

// GetZapierEventsResponse is generated from the OpenAPI component of the same name.
type GetZapierEventsResponse = []GetZapierEventsResponseValueItem

// GetZapierHooksResponse is generated from the OpenAPI component of the same name.
type GetZapierHooksResponse struct {
	Object     string       `json:"object"`
	Data       []ZapierHook `json:"data"`
	HasMore    bool         `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string      `json:"next_cursor"`
}

// GetZapierTriggersEventResponseValueItem: The event's payload fields, flattened, plus id (the event id), event and createdAt.
type GetZapierTriggersEventResponseValueItem struct {
	ID        string `json:"id"`
	Event     string `json:"event"`
	CreatedAt string `json:"createdAt"`
}

// GetZapierTriggersEventResponse is generated from the OpenAPI component of the same name.
type GetZapierTriggersEventResponse = []GetZapierTriggersEventResponseValueItem

// PostZapierConversionsResponse is generated from the OpenAPI component of the same name.
type PostZapierConversionsResponse struct {
	ConnectionID string `json:"connectionId"`
}

// PostPartnerSecurityTotpSetupResponse is generated from the OpenAPI component of the same name.
type PostPartnerSecurityTotpSetupResponse struct {
	Secret     string `json:"secret"`
	OtpauthURL string `json:"otpauthUrl"`
}

// PostPartnerSecurityTotpEnableResponse is generated from the OpenAPI component of the same name.
type PostPartnerSecurityTotpEnableResponse struct {
	Enabled bool `json:"enabled"`
}

// PostPartnerSecurityTotpDisableResponse is generated from the OpenAPI component of the same name.
type PostPartnerSecurityTotpDisableResponse struct {
	Enabled bool `json:"enabled"`
}

// GetPartnerSecurityResponse is generated from the OpenAPI component of the same name.
type GetPartnerSecurityResponse struct {
	TotpEnabled bool `json:"totpEnabled"`
}

// GetAuthConnectionsProvidersResponse is generated from the OpenAPI component of the same name.
type GetAuthConnectionsProvidersResponse struct {
	Data []IdentityProvider `json:"data"`
}

// GetAuthConnectionsResponse is generated from the OpenAPI component of the same name.
type GetAuthConnectionsResponse struct {
	Object     string           `json:"object"`
	Data       []AuthConnection `json:"data"`
	HasMore    bool             `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string          `json:"next_cursor"`
}

// PutAuthConnectionsIdSecretResponse is generated from the OpenAPI component of the same name.
type PutAuthConnectionsIdSecretResponse struct {
	ID       string `json:"id"`
	Provider string `json:"provider"`
	Ok       bool   `json:"ok"`
}

// PostAuthConnectionsIdRotateSecretResponse is generated from the OpenAPI component of the same name.
type PostAuthConnectionsIdRotateSecretResponse struct {
	ID                      string `json:"id"`
	Provider                string `json:"provider"`
	Secret                  string `json:"secret"` // The new secret, shown once. The previous one keeps working until previousSecretExpiresAt.
	PreviousSecretExpiresAt string `json:"previousSecretExpiresAt"`
}

// PatchAuthConnectionsIdResponse is generated from the OpenAPI component of the same name.
type PatchAuthConnectionsIdResponse struct {
	ID               string  `json:"id"`
	Provider         string  `json:"provider"`
	Name             string  `json:"name"`
	Status           string  `json:"status"`
	AutoEnrollUsers  *bool   `json:"autoEnrollUsers,omitempty"`
	DestinationURL   *string `json:"destinationUrl,omitempty"`
	SignatureVersion *string `json:"signatureVersion,omitempty"`
	RequiresV2       *bool   `json:"requiresV2,omitempty"`
}

// PostAuthConnectionsIdRevokePreviousSecretResponse is generated from the OpenAPI component of the same name.
type PostAuthConnectionsIdRevokePreviousSecretResponse struct {
	Ok      bool `json:"ok"`
	Revoked bool `json:"revoked"`
}

// PostAuthConnectionsIdSignupsResponseReferralOption1 is generated from the OpenAPI component of the same name.
type PostAuthConnectionsIdSignupsResponseReferralOption1 struct {
	ID                 string  `json:"id"`
	PartnerID          string  `json:"partnerId"`
	PartnerExternalRef *string `json:"partnerExternalRef"`
	MatchedVia         string  `json:"matchedVia"`
	Status             string  `json:"status"`
	Persisted          *bool   `json:"persisted,omitempty"` // Test signups only: true when this referral already existed; false when it exists only in this answer.
}

// PostAuthConnectionsIdSignupsResponseReasoning is generated from the OpenAPI component of the same name.
type PostAuthConnectionsIdSignupsResponseReasoning struct {
	Notes []string           `json:"notes"` // English, for people; wording may change.
	Codes []SignupReasonCode `json:"codes"` // One stable code per distinct reason behind notes, in order.
}

// PostAuthConnectionsIdSignupsResponsePartner is generated from the OpenAPI component of the same name.
type PostAuthConnectionsIdSignupsResponsePartner struct {
	ID        string  `json:"id"`
	Code      string  `json:"code"`
	Link      *string `json:"link"`                // The partner's absolute tracked link (/go/{trackingToken}/{code}), on the org's verified link domain when it has one; null when they have only a coupon.
	Persisted *bool   `json:"persisted,omitempty"` // Test signups only: true when this partner already existed; false when it exists only in this answer.
}

// PostAuthConnectionsIdSignupsResponse is generated from the OpenAPI component of the same name.
type PostAuthConnectionsIdSignupsResponse struct {
	Test      *bool                                                `json:"test,omitempty"` // Present on a test signup ("test": true or atlas-test: 1): answered as a real call, nothing stored.
	Referral  *PostAuthConnectionsIdSignupsResponseReferralOption1 `json:"referral"`
	Reasoning *PostAuthConnectionsIdSignupsResponseReasoning       `json:"reasoning,omitempty"`
	Partner   *PostAuthConnectionsIdSignupsResponsePartner         `json:"partner,omitempty"`
}

// PostAuthConnectionsIdPartnersResponsePartner is generated from the OpenAPI component of the same name.
type PostAuthConnectionsIdPartnersResponsePartner struct {
	ID   string  `json:"id"`
	Code string  `json:"code"`
	Link *string `json:"link"` // Their tracked link, when the connection has a destination.
}

// PostAuthConnectionsIdPartnersResponse is generated from the OpenAPI component of the same name.
type PostAuthConnectionsIdPartnersResponse struct {
	Partner PostAuthConnectionsIdPartnersResponsePartner `json:"partner"`
}

// PostAuthConnectionsIdEraseResponseCustomerOption1 is generated from the OpenAPI component of the same name.
type PostAuthConnectionsIdEraseResponseCustomerOption1 struct {
	CustomerRef string `json:"customerRef"`
	Pseudonym   string `json:"pseudonym"`
}

// PostAuthConnectionsIdEraseResponsePartnerOption1 is generated from the OpenAPI component of the same name.
type PostAuthConnectionsIdEraseResponsePartnerOption1 struct {
	PartnerID        string `json:"partnerId"`
	AlreadyForgotten bool   `json:"alreadyForgotten"`
}

// PostAuthConnectionsIdEraseResponse is generated from the OpenAPI component of the same name.
type PostAuthConnectionsIdEraseResponse struct {
	Ok       bool                                               `json:"ok"`
	Erased   bool                                               `json:"erased"` // False when nothing was filed under the customer (a retry settles cleanly).
	Customer *PostAuthConnectionsIdEraseResponseCustomerOption1 `json:"customer"`
	Partner  *PostAuthConnectionsIdEraseResponsePartnerOption1  `json:"partner"`
}

// GetAuthConnectionsIdExportResponse is generated from the OpenAPI component of the same name.
type GetAuthConnectionsIdExportResponse struct {
	Customer CustomerExport `json:"customer"`
	Partner  *PartnerExport `json:"partner"`
}

// PostAuthConnectionsIdPingResponse is generated from the OpenAPI component of the same name.
type PostAuthConnectionsIdPingResponse struct {
	Ok              bool   `json:"ok"`
	ConnectionID    string `json:"connectionId"`
	ProgramID       string `json:"programId"`
	OrgName         string `json:"orgName"`
	AutoEnrollUsers bool   `json:"autoEnrollUsers"`
	ServerTime      string `json:"serverTime"`
}

// PostAuthConnectionsIdAliasesResponse is generated from the OpenAPI component of the same name.
type PostAuthConnectionsIdAliasesResponse struct {
	Ok                  bool     `json:"ok"`
	CustomerRef         string   `json:"customerRef"`
	ReferralIds         []string `json:"referralIds"`
	PartnerID           *string  `json:"partnerId"`
	AlreadyAliased      bool     `json:"alreadyAliased"`
	ReprocessedPayments int64    `json:"reprocessedPayments"`
}

// PostAuthConnectionsIdUsersResponseCustomerOption1 is generated from the OpenAPI component of the same name.
type PostAuthConnectionsIdUsersResponseCustomerOption1 struct {
	CustomerRef string `json:"customerRef"`
	Pseudonym   string `json:"pseudonym"`
}

// PostAuthConnectionsIdUsersResponsePartnerOption1Option1 is generated from the OpenAPI component of the same name.
type PostAuthConnectionsIdUsersResponsePartnerOption1Option1 struct {
	PartnerID        string `json:"partnerId"`
	AlreadyForgotten bool   `json:"alreadyForgotten"`
}

// PostAuthConnectionsIdUsersResponsePartnerOption1Option2 is generated from the OpenAPI component of the same name.
type PostAuthConnectionsIdUsersResponsePartnerOption1Option2 struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

// PostAuthConnectionsIdUsersResponse is generated from the OpenAPI component of the same name.
type PostAuthConnectionsIdUsersResponse struct {
	Ok       bool                                               `json:"ok"`
	Deleted  *bool                                              `json:"deleted,omitempty"`
	Erased   *bool                                              `json:"erased,omitempty"`
	Customer *PostAuthConnectionsIdUsersResponseCustomerOption1 `json:"customer,omitempty"`
	Partner  json.RawMessage                                    `json:"partner"`
	Updated  []string                                           `json:"updated,omitempty"`
	Notes    []string                                           `json:"notes,omitempty"`
}

// PostAuthConnectionsIdSignalsResponseReferralsItem is generated from the OpenAPI component of the same name.
type PostAuthConnectionsIdSignalsResponseReferralsItem struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

// PostAuthConnectionsIdSignalsResponseReversedItem is generated from the OpenAPI component of the same name.
type PostAuthConnectionsIdSignalsResponseReversedItem struct {
	LineID      string `json:"lineId"`
	AmountMinor int64  `json:"amountMinor"` // Integer minor units (cents).
	Currency    string `json:"currency"`
}

// PostAuthConnectionsIdSignalsResponsePartnerOption1 is generated from the OpenAPI component of the same name.
type PostAuthConnectionsIdSignalsResponsePartnerOption1 struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

// PostAuthConnectionsIdSignalsResponse is generated from the OpenAPI component of the same name.
type PostAuthConnectionsIdSignalsResponse struct {
	Ok        bool                                                `json:"ok"`
	Signal    string                                              `json:"signal"`
	Referrals []PostAuthConnectionsIdSignalsResponseReferralsItem `json:"referrals"`
	Reversed  []PostAuthConnectionsIdSignalsResponseReversedItem  `json:"reversed"`
	Partner   *PostAuthConnectionsIdSignalsResponsePartnerOption1 `json:"partner"`
	Notes     []string                                            `json:"notes,omitempty"`
}

// PostAuthConnectionsInstallResponse is generated from the OpenAPI component of the same name.
type PostAuthConnectionsInstallResponse struct {
	RedirectTo string `json:"redirectTo"` // Send the browser here: back to the platform with a code, or error=access_denied.
}

// PostOauthTokenResponse is generated from the OpenAPI component of the same name.
type PostOauthTokenResponse struct {
	TokenType    string `json:"token_type"`
	ConnectionID string `json:"connectionId"`
	Provider     string `json:"provider"`
	InstanceID   string `json:"instanceId"`
	ProgramID    string `json:"programId"`
	Secret       string `json:"secret"` // The connection signing secret. Returned only here.
	Reinstalled  bool   `json:"reinstalled"`
	SignupsURL   string `json:"signupsUrl"`
	PartnersURL  string `json:"partnersUrl"`
	EraseURL     string `json:"eraseUrl"`
	ExportURL    string `json:"exportUrl"`
}

// GetOrgSsoResponseProviderOption1 is generated from the OpenAPI component of the same name.
type GetOrgSsoResponseProviderOption1 struct {
	Name  string `json:"name"`
	Label string `json:"label"`
}

// GetOrgSsoResponseConnectionsItem is generated from the OpenAPI component of the same name.
type GetOrgSsoResponseConnectionsItem struct {
	ID             string   `json:"id"`
	Type           string   `json:"type"`
	Status         string   `json:"status"`
	AllowedDomains []string `json:"allowedDomains"`
}

// GetOrgSsoResponseTicketsItem is generated from the OpenAPI component of the same name.
type GetOrgSsoResponseTicketsItem struct {
	ID        string `json:"id"`
	Status    string `json:"status"`
	ExpiresAt string `json:"expiresAt"`
	CreatedAt string `json:"createdAt"`
}

// GetOrgSsoResponse is generated from the OpenAPI component of the same name.
type GetOrgSsoResponse struct {
	Configured         bool                               `json:"configured"`
	Provider           *GetOrgSsoResponseProviderOption1  `json:"provider"`           // The deployment's SSO provider; null when SSO isn't configured.
	DefaultRoleMapping map[string]string                  `json:"defaultRoleMapping"` // The provider's role keys → Referalli roles, before the org's own mapping.
	Linked             bool                               `json:"linked"`
	TenantID           *string                            `json:"tenantId"`   // The org's directory tenant at the provider.
	AtlasOrgID         *string                            `json:"atlasOrgId"` // Deprecated: tenantId, when the provider is Atlas.
	Connections        []GetOrgSsoResponseConnectionsItem `json:"connections"`
	Domains            []string                           `json:"domains"`
	RequireSSO         bool                               `json:"requireSso"`
	RoleMapping        map[string]string                  `json:"roleMapping"`
	Tickets            []GetOrgSsoResponseTicketsItem     `json:"tickets"`
	ProviderError      *string                            `json:"providerError"` // Why the provider couldn't be read just now, if it couldn't.
	AtlasError         *string                            `json:"atlasError"`    // Deprecated: providerError.
}

// PostOrgSsoLinkResponse is generated from the OpenAPI component of the same name.
type PostOrgSsoLinkResponse struct {
	Linked     bool    `json:"linked"`
	Provider   string  `json:"provider"`
	TenantID   string  `json:"tenantId"`
	AtlasOrgID *string `json:"atlasOrgId"`
}

// PostOrgSsoTicketsResponse is generated from the OpenAPI component of the same name.
type PostOrgSsoTicketsResponse struct {
	ID        string `json:"id"`
	URL       string `json:"url"` // The SSO provider's hosted setup page for the IT admin. Shown once; not stored.
	ExpiresAt string `json:"expiresAt"`
}

// PostOrgSsoTicketsIdRevokeResponse is generated from the OpenAPI component of the same name.
type PostOrgSsoTicketsIdRevokeResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

// GetOrgSsoConnectionsIdSamlMetadataResponse is generated from the OpenAPI component of the same name.
type GetOrgSsoConnectionsIdSamlMetadataResponse struct {
	SpEntityID  string `json:"spEntityId"`
	AcsURL      string `json:"acsUrl"`
	MetadataXml string `json:"metadataXml"`
}

// PatchOrgSsoResponse is generated from the OpenAPI component of the same name.
type PatchOrgSsoResponse struct {
	RequireSSO  bool              `json:"requireSso"`
	RoleMapping map[string]string `json:"roleMapping"`
}

// PutOrgSsoDirectResponse is generated from the OpenAPI component of the same name.
type PutOrgSsoDirectResponse struct {
	Adapter     string         `json:"adapter"`
	Config      map[string]any `json:"config"` // The stored settings; the secret is sealed and never returned.
	HasSecret   bool           `json:"hasSecret"`
	SignInURL   string         `json:"signInUrl"`   // Where members start sign-in (GET /v1/auth/sso/start?org=…).
	RedirectURI string         `json:"redirectUri"` // The redirect URI to register at the IdP.
}

// DeleteOrgSsoDirectResponse is generated from the OpenAPI component of the same name.
type DeleteOrgSsoDirectResponse struct {
	Removed bool `json:"removed"`
}

// PostPartnerAuthWalletNonceResponse is generated from the OpenAPI component of the same name.
type PostPartnerAuthWalletNonceResponse struct {
	Nonce   string `json:"nonce"`
	Message string `json:"message"`
}

// PostPartnerAuthWalletVerifyResponse is generated from the OpenAPI component of the same name.
type PostPartnerAuthWalletVerifyResponse struct {
	AccountID string `json:"accountId"`
	Wallet    string `json:"wallet"`
}

// PostOnchainConnectionsResponse is generated from the OpenAPI component of the same name.
type PostOnchainConnectionsResponse struct {
	ID         string `json:"id"`
	Secret     string `json:"secret"`
	WebhookURL string `json:"webhookUrl"`
}

// GetOnchainConnectionsResponseDataItem is generated from the OpenAPI component of the same name.
type GetOnchainConnectionsResponseDataItem struct {
	ID        string `json:"id"`
	ProgramID string `json:"programId"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	CreatedAt string `json:"createdAt"`
}

// GetOnchainConnectionsResponse is generated from the OpenAPI component of the same name.
type GetOnchainConnectionsResponse struct {
	Object     string                                  `json:"object"`
	Data       []GetOnchainConnectionsResponseDataItem `json:"data"`
	HasMore    bool                                    `json:"has_more"` // Whether another page follows; pass next_cursor as starting_after to get it.
	NextCursor *string                                 `json:"next_cursor"`
}

// GetPublicProgramsHandleResponseBranding is generated from the OpenAPI component of the same name.
type GetPublicProgramsHandleResponseBranding struct {
	LogoURL *string                 `json:"logoUrl"`
	Theme   *GetOpenapiJsonResponse `json:"theme"`
}

// GetPublicProgramsHandleResponse is generated from the OpenAPI component of the same name.
type GetPublicProgramsHandleResponse struct {
	Handle           string                                  `json:"handle"`
	ProgramID        string                                  `json:"programId"`
	ProgramName      string                                  `json:"programName"`
	Merchant         string                                  `json:"merchant"`
	TrackingToken    string                                  `json:"trackingToken"`
	CookieWindowDays int64                                   `json:"cookieWindowDays"`
	TermsMd          string                                  `json:"termsMd"`
	Pitch            string                                  `json:"pitch"`
	AutoApprove      bool                                    `json:"autoApprove"`
	Commission       *GetOpenapiJsonResponse                 `json:"commission"`
	Branding         GetPublicProgramsHandleResponseBranding `json:"branding"`
}

// PostPublicProgramsHandleJoinResponse is generated from the OpenAPI component of the same name.
type PostPublicProgramsHandleJoinResponse struct {
	AccountID string  `json:"accountId"`
	Status    string  `json:"status"`
	Link      *string `json:"link"`
	LinkURL   *string `json:"linkUrl"`
	Code      *string `json:"code"`
	PortalURL string  `json:"portalUrl"`
}

// GetIntegrationHealthResponse is generated from the OpenAPI component of the same name.
type GetIntegrationHealthResponse struct {
	Ok          bool          `json:"ok"`
	Checks      []HealthCheck `json:"checks"`
	GeneratedAt string        `json:"generatedAt"`
}

// GetPlansResponsePlansItem is generated from the OpenAPI component of the same name.
type GetPlansResponsePlansItem struct {
	Key          string `json:"key"`
	Label        string `json:"label"`
	PriceMinor   int64  `json:"priceMinor"` // Integer minor units (cents).
	Currency     string `json:"currency"`
	Interval     string `json:"interval"`
	MaxPrograms  *int64 `json:"maxPrograms"` // null: unlimited.
	MaxPartners  *int64 `json:"maxPartners"` // null: unlimited.
	Marketplace  bool   `json:"marketplace"`
	Discovery    bool   `json:"discovery"`
	LinkDomain   bool   `json:"linkDomain"`
	PortalDomain bool   `json:"portalDomain"`
	APIPerSec    int64  `json:"apiPerSec"`
	APIBurst     int64  `json:"apiBurst"`
}

// GetPlansResponse is generated from the OpenAPI component of the same name.
type GetPlansResponse struct {
	Plans []GetPlansResponsePlansItem `json:"plans"`
}

// GetEventsResponseEventsItem is generated from the OpenAPI component of the same name.
type GetEventsResponseEventsItem struct {
	ID        string                 `json:"id"`
	Topic     string                 `json:"topic"`
	OrgID     string                 `json:"orgId"`
	Livemode  bool                   `json:"livemode"`
	CreatedAt string                 `json:"createdAt"`
	Data      GetOpenapiJsonResponse `json:"data"` // The topic's payload (see the webhooks section).
}

// GetEventsResponse is generated from the OpenAPI component of the same name.
type GetEventsResponse struct {
	Events     []GetEventsResponseEventsItem `json:"events"`
	HasMore    bool                          `json:"hasMore"`
	NextCursor *string                       `json:"nextCursor"` // Pass as `after` for the next page; null on the last.
}

// GetSupportContextResponsePartnerOption1 is generated from the OpenAPI component of the same name.
type GetSupportContextResponsePartnerOption1 struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	Status        string   `json:"status"`
	Tier          *string  `json:"tier"`
	Programs      []string `json:"programs"`
	ReferralLink  *string  `json:"referralLink"`
	PendingMinor  int64    `json:"pendingMinor"`  // Integer minor units (cents).
	ApprovedMinor int64    `json:"approvedMinor"` // Integer minor units (cents).
	PaidMinor     int64    `json:"paidMinor"`     // Integer minor units (cents).
	Currency      string   `json:"currency"`
	NextPayoutAt  *string  `json:"nextPayoutAt"`
	Flags         []string `json:"flags"`
}

// GetSupportContextResponseCustomerOption1ReferredByOption1 is generated from the OpenAPI component of the same name.
type GetSupportContextResponseCustomerOption1ReferredByOption1 struct {
	PartnerID string `json:"partnerId"`
	Name      string `json:"name"`
}

// GetSupportContextResponseCustomerOption1RewardStateOption1 is generated from the OpenAPI component of the same name.
type GetSupportContextResponseCustomerOption1RewardStateOption1 struct {
	Kind           string `json:"kind"`
	Status         string `json:"status"`
	AmountMinor    int64  `json:"amountMinor"` // Integer minor units (cents).
	Currency       string `json:"currency"`
	RemainingMinor int64  `json:"remainingMinor"` // Integer minor units (cents).
}

// GetSupportContextResponseCustomerOption1 is generated from the OpenAPI component of the same name.
type GetSupportContextResponseCustomerOption1 struct {
	ReferredBy       *GetSupportContextResponseCustomerOption1ReferredByOption1  `json:"referredBy"`
	AttributedAt     *string                                                     `json:"attributedAt"`
	RewardState      *GetSupportContextResponseCustomerOption1RewardStateOption1 `json:"rewardState"`
	CommissionsMinor int64                                                       `json:"commissionsMinor"` // Integer minor units (cents).
	Currency         *string                                                     `json:"currency"`
}

// GetSupportContextResponseLinks is generated from the OpenAPI component of the same name.
type GetSupportContextResponseLinks struct {
	Dashboard string  `json:"dashboard"`
	Partner   *string `json:"partner"`
	Referral  *string `json:"referral"`
}

// GetSupportContextResponse is generated from the OpenAPI component of the same name.
type GetSupportContextResponse struct {
	Partner  *GetSupportContextResponsePartnerOption1  `json:"partner"`
	Customer *GetSupportContextResponseCustomerOption1 `json:"customer"`
	Links    GetSupportContextResponseLinks            `json:"links"`
}

// GetSupportAppsResponseValueItemInstall is generated from the OpenAPI component of the same name.
type GetSupportAppsResponseValueItemInstall struct {
	Kind         *string  `json:"kind,omitempty"`
	Steps        []string `json:"steps"`
	ManifestPath *string  `json:"manifestPath,omitempty"`
}

// GetSupportAppsResponseValueItemConfigOption1 is generated from the OpenAPI component of the same name.
type GetSupportAppsResponseValueItemConfigOption1 struct {
	Subdomain    any  `json:"subdomain"`
	CardKey      any  `json:"cardKey"`
	HasPublicKey bool `json:"hasPublicKey"`
}

// GetSupportAppsResponseValueItem is generated from the OpenAPI component of the same name.
type GetSupportAppsResponseValueItem struct {
	ID            string                                        `json:"id"`
	Name          string                                        `json:"name"`
	Status        string                                        `json:"status"`
	DocsURL       string                                        `json:"docsUrl"`
	VendorDocsURL string                                        `json:"vendorDocsUrl"`
	SecretSource  string                                        `json:"secretSource"`
	Install       GetSupportAppsResponseValueItemInstall        `json:"install"`
	CardURL       string                                        `json:"cardUrl"`
	Enabled       bool                                          `json:"enabled"`
	Configured    bool                                          `json:"configured"`
	Config        *GetSupportAppsResponseValueItemConfigOption1 `json:"config"`
	LastOkAt      *string                                       `json:"lastOkAt"`
	LastError     *string                                       `json:"lastError"`
}

// GetSupportAppsResponse is generated from the OpenAPI component of the same name.
type GetSupportAppsResponse = []GetSupportAppsResponseValueItem

// PutSupportAppsAdapterResponse is generated from the OpenAPI component of the same name.
type PutSupportAppsAdapterResponse struct {
	Adapter string  `json:"adapter"`
	Enabled bool    `json:"enabled"`
	CardURL string  `json:"cardUrl"`
	Secret  *string `json:"secret,omitempty"`
}

// DeleteSupportAppsAdapterResponse is generated from the OpenAPI component of the same name.
type DeleteSupportAppsAdapterResponse struct {
	Ok bool `json:"ok"`
}

// PostSupportAppsAdapterTestResponse is generated from the OpenAPI component of the same name.
type PostSupportAppsAdapterTestResponse struct {
	Found    bool                       `json:"found"`
	Context  *GetSupportContextResponse `json:"context"`
	Rendered any                        `json:"rendered"`
}

// PostSupportCardsOrgIdAdapterResponse is generated from the OpenAPI component of the same name.
type PostSupportCardsOrgIdAdapterResponse = map[string]any

// GetPrivacyExecutorsResponseValueItemInstall is generated from the OpenAPI component of the same name.
type GetPrivacyExecutorsResponseValueItemInstall struct {
	Kind         *string  `json:"kind,omitempty"`
	Steps        []string `json:"steps"`
	ManifestPath *string  `json:"manifestPath,omitempty"`
}

// GetPrivacyExecutorsResponseValueItemConfigOption1 is generated from the OpenAPI component of the same name.
type GetPrivacyExecutorsResponseValueItemConfigOption1 struct {
	HasPublicKey bool    `json:"hasPublicKey"`
	Algorithm    *string `json:"algorithm"`
	SombraURL    *string `json:"sombraUrl"`
}

// GetPrivacyExecutorsResponseValueItemHealth is generated from the OpenAPI component of the same name.
type GetPrivacyExecutorsResponseValueItemHealth struct {
	Ok     bool    `json:"ok"`
	Detail string  `json:"detail"`
	Fix    *string `json:"fix"`
}

// GetPrivacyExecutorsResponseValueItem is generated from the OpenAPI component of the same name.
type GetPrivacyExecutorsResponseValueItem struct {
	ID              string                                             `json:"id"`
	Name            string                                             `json:"name"`
	Status          string                                             `json:"status"`
	DocsURL         string                                             `json:"docsUrl"`
	VendorDocsURL   string                                             `json:"vendorDocsUrl"`
	KeyKind         string                                             `json:"keyKind"`
	Install         GetPrivacyExecutorsResponseValueItemInstall        `json:"install"`
	URL             string                                             `json:"url"`
	Enabled         bool                                               `json:"enabled"`
	HasKey          bool                                               `json:"hasKey"`
	HasVendorSecret bool                                               `json:"hasVendorSecret"`
	Config          *GetPrivacyExecutorsResponseValueItemConfigOption1 `json:"config"`
	Health          GetPrivacyExecutorsResponseValueItemHealth         `json:"health"`
	LastOkAt        *string                                            `json:"lastOkAt"`
	LastError       *string                                            `json:"lastError"`
}

// GetPrivacyExecutorsResponse is generated from the OpenAPI component of the same name.
type GetPrivacyExecutorsResponse = []GetPrivacyExecutorsResponseValueItem

// PostPrivacyExecutorKeysResponse is generated from the OpenAPI component of the same name.
type PostPrivacyExecutorKeysResponse struct {
	Adapter string  `json:"adapter"`
	URL     string  `json:"url"`
	Enabled bool    `json:"enabled"`
	Key     *string `json:"key,omitempty"`
}

// DeletePrivacyExecutorsAdapterResponse is generated from the OpenAPI component of the same name.
type DeletePrivacyExecutorsAdapterResponse struct {
	Ok bool `json:"ok"`
}

// PostPrivacyExecutorsAdapterCheckResponse is generated from the OpenAPI component of the same name.
type PostPrivacyExecutorsAdapterCheckResponse struct {
	Adapter string  `json:"adapter"`
	Ok      bool    `json:"ok"`
	Detail  string  `json:"detail"`
	Fix     *string `json:"fix"`
}

// GetPrivacyExecutorRequestsResponseValueItem is generated from the OpenAPI component of the same name.
type GetPrivacyExecutorRequestsResponseValueItem struct {
	ID              string  `json:"id"`
	Adapter         string  `json:"adapter"`
	VendorRequestID *string `json:"vendorRequestId"`
	Op              string  `json:"op"`
	Status          string  `json:"status"`
	CreatedAt       string  `json:"createdAt"`
	CompletedAt     *string `json:"completedAt"`
}

// GetPrivacyExecutorRequestsResponse is generated from the OpenAPI component of the same name.
type GetPrivacyExecutorRequestsResponse = []GetPrivacyExecutorRequestsResponseValueItem

// PostAccountCloseResponse is generated from the OpenAPI component of the same name.
type PostAccountCloseResponse struct {
	Closed bool `json:"closed"`
}

// GetClicksConsentCoverageResponse is generated from the OpenAPI component of the same name.
type GetClicksConsentCoverageResponse struct {
	Days         int64    `json:"days"`
	Since        string   `json:"since"`
	Attributed   int64    `json:"attributed"`   // Attributed (non-void) sign-ups in the window.
	WithEvidence int64    `json:"withEvidence"` // Of those, how many have consent evidence on their click or association.
	Share        *float64 `json:"share"`        // withEvidence / attributed, 4 decimals; null with no sign-ups.
}

// GetClicksIdResponseConsentOption1 is generated from the OpenAPI component of the same name.
type GetClicksIdResponseConsentOption1 struct {
	Provider  string   `json:"provider"`
	State     string   `json:"state"`
	At        *string  `json:"at,omitempty"`
	Stamp     *string  `json:"stamp,omitempty"`
	TcString  *string  `json:"tcString,omitempty"`
	GppString *string  `json:"gppString,omitempty"`
	Purposes  []string `json:"purposes,omitempty"`
}

// GetClicksIdResponse is generated from the OpenAPI component of the same name.
type GetClicksIdResponse struct {
	ID         string                             `json:"id"`
	ProgramID  string                             `json:"programId"`
	PartnerID  string                             `json:"partnerId"`
	LinkID     *string                            `json:"linkId"`
	Ts         string                             `json:"ts"`
	Source     string                             `json:"source"`
	Eligible   bool                               `json:"eligible"`
	FraudScore int64                              `json:"fraudScore"`
	AnonID     *string                            `json:"anonId"`
	Consent    *GetClicksIdResponseConsentOption1 `json:"consent"` // The visitor's consent when the click was stored (a redirect click: from the first consented visit).
}

// GetOrgTrackingConsentResponseAdaptersItem is generated from the OpenAPI component of the same name.
type GetOrgTrackingConsentResponseAdaptersItem struct {
	ID              string  `json:"id"`
	Name            string  `json:"name"`
	DefaultCategory *string `json:"defaultCategory"`
	Status          string  `json:"status"`
	DocsURL         string  `json:"docsUrl"`
}

// GetOrgTrackingConsentResponse is generated from the OpenAPI component of the same name.
type GetOrgTrackingConsentResponse struct {
	Mode        string                                      `json:"mode"` // `auto`, `none`, `pending` or a consent adapter id.
	NoCmp       *string                                     `json:"noCmp,omitempty"`
	Category    *string                                     `json:"category,omitempty"`
	TcfVendorID *int64                                      `json:"tcfVendorId,omitempty"`
	GcmSignal   *string                                     `json:"gcmSignal,omitempty"`
	Configured  bool                                        `json:"configured"` // False for an org that never chose (its tag behaves as before: `none`).
	Adapters    []GetOrgTrackingConsentResponseAdaptersItem `json:"adapters"`
}

// GetOrgHostedPagesResponseConsent is generated from the OpenAPI component of the same name.
type GetOrgHostedPagesResponseConsent struct {
	Adapter string         `json:"adapter"`
	Config  map[string]any `json:"config"`
}

// GetOrgHostedPagesResponseChatOption1 is generated from the OpenAPI component of the same name.
type GetOrgHostedPagesResponseChatOption1 struct {
	Adapter string         `json:"adapter"`
	Config  map[string]any `json:"config"`
}

// GetOrgHostedPagesResponseServedConsentTagsItem is generated from the OpenAPI component of the same name.
type GetOrgHostedPagesResponseServedConsentTagsItem struct {
	Tag    string                     `json:"tag"`
	Attrs  map[string]json.RawMessage `json:"attrs"`
	Inline *string                    `json:"inline,omitempty"`
}

// GetOrgHostedPagesResponseServedConsentSettingsLinkOption1 is generated from the OpenAPI component of the same name.
type GetOrgHostedPagesResponseServedConsentSettingsLinkOption1 struct {
	Call  *string           `json:"call,omitempty"`
	Attrs map[string]string `json:"attrs,omitempty"`
}

// GetOrgHostedPagesResponseServedConsentBuiltin is generated from the OpenAPI component of the same name.
type GetOrgHostedPagesResponseServedConsentBuiltin struct {
	PolicyURL string            `json:"policyUrl"`
	Text      map[string]string `json:"text,omitempty"`
}

// GetOrgHostedPagesResponseServedConsent is generated from the OpenAPI component of the same name.
type GetOrgHostedPagesResponseServedConsent struct {
	Mode         string                                                     `json:"mode"`
	Provider     string                                                     `json:"provider"`
	Tags         []GetOrgHostedPagesResponseServedConsentTagsItem           `json:"tags"`
	SettingsLink *GetOrgHostedPagesResponseServedConsentSettingsLinkOption1 `json:"settingsLink"`
	Builtin      *GetOrgHostedPagesResponseServedConsentBuiltin             `json:"builtin,omitempty"`
}

// GetOrgHostedPagesResponseServedChatOption1TagsItem is generated from the OpenAPI component of the same name.
type GetOrgHostedPagesResponseServedChatOption1TagsItem struct {
	Tag    string                     `json:"tag"`
	Attrs  map[string]json.RawMessage `json:"attrs"`
	Inline *string                    `json:"inline,omitempty"`
}

// GetOrgHostedPagesResponseServedChatOption1 is generated from the OpenAPI component of the same name.
type GetOrgHostedPagesResponseServedChatOption1 struct {
	Provider string                                               `json:"provider"`
	Tags     []GetOrgHostedPagesResponseServedChatOption1TagsItem `json:"tags"`
}

// GetOrgHostedPagesResponseServed is generated from the OpenAPI component of the same name.
type GetOrgHostedPagesResponseServed struct {
	Consent GetOrgHostedPagesResponseServedConsent      `json:"consent"`
	Chat    *GetOrgHostedPagesResponseServedChatOption1 `json:"chat"`
	Csp     string                                      `json:"csp"` // The Content-Security-Policy for this host's pages, built from the enabled adapters.
}

// GetOrgHostedPagesResponseAdaptersConsentItem is generated from the OpenAPI component of the same name.
type GetOrgHostedPagesResponseAdaptersConsentItem struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Status        string `json:"status"`
	DocsURL       string `json:"docsUrl"`
	VendorDocsURL string `json:"vendorDocsUrl"`
}

// GetOrgHostedPagesResponseAdaptersChatItem is generated from the OpenAPI component of the same name.
type GetOrgHostedPagesResponseAdaptersChatItem struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Status        string `json:"status"`
	DocsURL       string `json:"docsUrl"`
	VendorDocsURL string `json:"vendorDocsUrl"`
}

// GetOrgHostedPagesResponseAdapters is generated from the OpenAPI component of the same name.
type GetOrgHostedPagesResponseAdapters struct {
	Consent []GetOrgHostedPagesResponseAdaptersConsentItem `json:"consent"`
	Chat    []GetOrgHostedPagesResponseAdaptersChatItem    `json:"chat"`
}

// GetOrgHostedPagesResponse is generated from the OpenAPI component of the same name.
type GetOrgHostedPagesResponse struct {
	Consent       GetOrgHostedPagesResponseConsent      `json:"consent"`
	Chat          *GetOrgHostedPagesResponseChatOption1 `json:"chat"`
	CustomOrigins map[string][]string                   `json:"customOrigins"`
	Status        string                                `json:"status"`
	LastOkAt      *string                               `json:"lastOkAt"`
	LastError     *string                               `json:"lastError"`
	Served        GetOrgHostedPagesResponseServed       `json:"served"`
	Adapters      GetOrgHostedPagesResponseAdapters     `json:"adapters"`
}

// PostOrgHostedPagesCheckResponseChecksItem is generated from the OpenAPI component of the same name.
type PostOrgHostedPagesCheckResponseChecksItem struct {
	ID     string  `json:"id"`
	Ok     bool    `json:"ok"`
	Detail string  `json:"detail"`
	Fix    *string `json:"fix,omitempty"`
}

// PostOrgHostedPagesCheckResponse is generated from the OpenAPI component of the same name.
type PostOrgHostedPagesCheckResponse struct {
	Ok     bool                                        `json:"ok"`
	Checks []PostOrgHostedPagesCheckResponseChecksItem `json:"checks"`
}

// PostWebhooksBillingConnectionIdResponseResultsItem is generated from the OpenAPI component of the same name.
type PostWebhooksBillingConnectionIdResponseResultsItem struct {
	Duplicate      bool    `json:"duplicate"`
	BillingEventID *string `json:"billingEventId"`
	Conversions    int64   `json:"conversions"`
	LinesCreated   int64   `json:"linesCreated"`
	Reversals      int64   `json:"reversals"`
	Ignored        *string `json:"ignored,omitempty"` // Recorded as ignored_by_filter: earned nothing.
}

// PostWebhooksBillingConnectionIdResponse is generated from the OpenAPI component of the same name.
type PostWebhooksBillingConnectionIdResponse struct {
	Received int64                                                `json:"received"`
	Results  []PostWebhooksBillingConnectionIdResponseResultsItem `json:"results"`
}

// PostWebhooksAuthConnectionIdResponseResultsItem is generated from the OpenAPI component of the same name.
type PostWebhooksAuthConnectionIdResponseResultsItem struct {
	CustomerRef string  `json:"customerRef"`
	Attributed  bool    `json:"attributed"`
	Reason      *string `json:"reason,omitempty"`
	PartnerID   *string `json:"partnerId,omitempty"`
	MatchedVia  *string `json:"matchedVia,omitempty"`
}

// PostWebhooksAuthConnectionIdResponse is generated from the OpenAPI component of the same name.
type PostWebhooksAuthConnectionIdResponse struct {
	Received int64                                             `json:"received"`
	Results  []PostWebhooksAuthConnectionIdResponseResultsItem `json:"results"`
}

// PostWebhooksSsoProviderResponse is generated from the OpenAPI component of the same name.
type PostWebhooksSsoProviderResponse struct {
	Received  bool     `json:"received"`
	Duplicate bool     `json:"duplicate"`
	Handled   *bool    `json:"handled,omitempty"`
	Effects   []string `json:"effects,omitempty"`
}

// PostWebhooksOnchainConnectionIdResponseResultsItem is generated from the OpenAPI component of the same name.
type PostWebhooksOnchainConnectionIdResponseResultsItem struct {
	TxHash       string `json:"txHash"`
	Wallet       string `json:"wallet"`
	AmountMinor  int64  `json:"amountMinor"` // Integer minor units (cents).
	Attributed   bool   `json:"attributed"`
	LinesCreated int64  `json:"linesCreated"`
	Duplicate    bool   `json:"duplicate"`
}

// PostWebhooksOnchainConnectionIdResponse is generated from the OpenAPI component of the same name.
type PostWebhooksOnchainConnectionIdResponse struct {
	Received int64                                                `json:"received"`
	Results  []PostWebhooksOnchainConnectionIdResponseResultsItem `json:"results"`
}

// PostWebhooksPayoutsProviderResponse is generated from the OpenAPI component of the same name.
type PostWebhooksPayoutsProviderResponse struct {
	Ignored   *bool   `json:"ignored,omitempty"`
	Reason    *string `json:"reason,omitempty"`
	Finalized *bool   `json:"finalized,omitempty"`
	Status    *string `json:"status,omitempty"`
}

// PostWebhooksStripeBillingResponse is generated from the OpenAPI component of the same name.
type PostWebhooksStripeBillingResponse struct {
	Received   bool    `json:"received"`
	Duplicate  *bool   `json:"duplicate,omitempty"`
	Ignored    *bool   `json:"ignored,omitempty"`
	Reason     *string `json:"reason,omitempty"`
	OrgID      *string `json:"orgId,omitempty"`
	Plan       *string `json:"plan,omitempty"`
	PlanStatus *string `json:"planStatus,omitempty"`
}

// PostApiPrivacyLookupResponse is generated from the OpenAPI component of the same name.
type PostApiPrivacyLookupResponse struct {
	SubjectID *string `json:"subjectId"`
	Status    string  `json:"status"`
}

// PostApiPrivacyRequestsResponse is generated from the OpenAPI component of the same name.
type PostApiPrivacyRequestsResponse struct {
	ID     string  `json:"id"`
	Status string  `json:"status"`
	Eta    *string `json:"eta,omitempty"`
	Detail *string `json:"detail,omitempty"`
}

// GetApiPrivacyRequestsIdResponse is generated from the OpenAPI component of the same name.
type GetApiPrivacyRequestsIdResponse struct {
	ID     string                  `json:"id"`
	Type   string                  `json:"type"`
	Status string                  `json:"status"`
	Eta    *string                 `json:"eta,omitempty"`
	Detail *string                 `json:"detail,omitempty"`
	Result *GetOpenapiJsonResponse `json:"result,omitempty"`
}

// PostApiPrivacyOrgsOrgIdLookupResponse is generated from the OpenAPI component of the same name.
type PostApiPrivacyOrgsOrgIdLookupResponse struct {
	SubjectID *string `json:"subjectId"`
	Status    string  `json:"status"`
}

// PostApiPrivacyOrgsOrgIdRequestsResponse is generated from the OpenAPI component of the same name.
type PostApiPrivacyOrgsOrgIdRequestsResponse struct {
	ID     string  `json:"id"`
	Status string  `json:"status"`
	Detail *string `json:"detail,omitempty"`
}

// GetApiPrivacyOrgsOrgIdRequestsIdResponse is generated from the OpenAPI component of the same name.
type GetApiPrivacyOrgsOrgIdRequestsIdResponse struct {
	ID     string                                `json:"id"`
	Type   string                                `json:"type"`
	Status string                                `json:"status"`
	Detail *string                               `json:"detail,omitempty"`
	Result *PostSupportCardsOrgIdAdapterResponse `json:"result,omitempty"`
}

// GetGraphqlResponse is generated from the OpenAPI component of the same name.
type GetGraphqlResponse struct {
	Data   *GetOpenapiJsonResponse  `json:"data,omitempty"`
	Errors []GetOpenapiJsonResponse `json:"errors,omitempty"`
}

// PostGraphqlResponse is generated from the OpenAPI component of the same name.
type PostGraphqlResponse struct {
	Data   *GetOpenapiJsonResponse  `json:"data,omitempty"`
	Errors []GetOpenapiJsonResponse `json:"errors,omitempty"`
}

// PostAgentsRegisterBody is generated from the OpenAPI component of the same name.
type PostAgentsRegisterBody struct {
	WalletAddress string  `json:"walletAddress"`
	Nonce         string  `json:"nonce"`
	Signature     string  `json:"signature"`
	Name          *string `json:"name,omitempty"`
	Model         *string `json:"model,omitempty"`
	Operator      *string `json:"operator,omitempty"`
	InstantPayout *bool   `json:"instantPayout,omitempty"`
}

// PostAuthBackchannelLogoutBody is generated from the OpenAPI component of the same name.
type PostAuthBackchannelLogoutBody struct {
	LogoutToken string `json:"logout_token"`
}

// PostOrgsBody is generated from the OpenAPI component of the same name.
type PostOrgsBody struct {
	Name string `json:"name"`
}

// PutOrgPortalAuthBodyOption1 is generated from the OpenAPI component of the same name.
type PutOrgPortalAuthBodyOption1 struct {
	Adapter string `json:"adapter"`
}

// PutOrgPortalAuthBodyOption2 is generated from the OpenAPI component of the same name.
type PutOrgPortalAuthBodyOption2 struct {
	Issuer              string  `json:"issuer"`
	ClientID            string  `json:"clientId"`
	ClientSecret        *string `json:"clientSecret,omitempty"`
	Scopes              *string `json:"scopes,omitempty"`
	Name                *string `json:"name,omitempty"`
	TrustEmailVerified  *bool   `json:"trustEmailVerified,omitempty"`
	Adapter             string  `json:"adapter"`
	AutoEnrollProgramID *string `json:"autoEnrollProgramId,omitempty"`
}

// PutOrgPortalAuthBodyOption3 is generated from the OpenAPI component of the same name.
type PutOrgPortalAuthBodyOption3 struct {
	Adapter             string  `json:"adapter"`
	Issuer              string  `json:"issuer"`
	JwksURL             string  `json:"jwksUrl"`
	TrustEmailVerified  *bool   `json:"trustEmailVerified,omitempty"`
	AutoEnrollProgramID *string `json:"autoEnrollProgramId,omitempty"`
	Name                *string `json:"name,omitempty"`
	LoginURL            *string `json:"loginUrl,omitempty"`
}

// PostMobilePurchasesBody is generated from the OpenAPI component of the same name.
type PostMobilePurchasesBody struct {
	EventID     string  `json:"eventId"`
	CustomerRef string  `json:"customerRef"`
	Type        *string `json:"type,omitempty"`
	AmountMinor int64   `json:"amountMinor"`
	Currency    *string `json:"currency,omitempty"`
	ProductRef  *string `json:"productRef,omitempty"`
}

// PatchOrgBody is generated from the OpenAPI component of the same name.
type PatchOrgBody struct {
	Name              *string        `json:"name,omitempty"`
	Settings          map[string]any `json:"settings,omitempty"`
	PayoutKYCRequired *bool          `json:"payoutKycRequired,omitempty"`
}

// PostOrgMembersBody is generated from the OpenAPI component of the same name.
type PostOrgMembersBody struct {
	Email string `json:"email"`
	Role  string `json:"role"`
}

// PatchOrgMembersUserIdBody is generated from the OpenAPI component of the same name.
type PatchOrgMembersUserIdBody struct {
	Role string `json:"role"`
}

// PutOrgPayoutScheduleBodyScheduleOption1Option1 is generated from the OpenAPI component of the same name.
type PutOrgPayoutScheduleBodyScheduleOption1Option1 struct {
	Kind       string `json:"kind"`
	DayOfMonth int64  `json:"dayOfMonth"`
}

// PutOrgPayoutScheduleBodyScheduleOption1Option2 is generated from the OpenAPI component of the same name.
type PutOrgPayoutScheduleBodyScheduleOption1Option2 struct {
	Kind      string `json:"kind"`
	DayOfWeek int64  `json:"dayOfWeek"`
}

// PutOrgPayoutScheduleBodyScheduleOption1Option3 is generated from the OpenAPI component of the same name.
type PutOrgPayoutScheduleBodyScheduleOption1Option3 struct {
	Kind string  `json:"kind"`
	Days []int64 `json:"days"`
}

// PutOrgPayoutScheduleBodyScheduleOption1Option4 is generated from the OpenAPI component of the same name.
type PutOrgPayoutScheduleBodyScheduleOption1Option4 struct {
	Kind  string   `json:"kind"`
	Dates []string `json:"dates"`
}

// PutOrgPayoutScheduleBody is generated from the OpenAPI component of the same name.
type PutOrgPayoutScheduleBody struct {
	DayOfMonth *int64          `json:"dayOfMonth,omitempty"`
	Schedule   json.RawMessage `json:"schedule,omitempty"`
}

// PutProgramsIdPayoutScheduleBodyScheduleOption1Option1Option1 is generated from the OpenAPI component of the same name.
type PutProgramsIdPayoutScheduleBodyScheduleOption1Option1Option1 struct {
	Kind       string `json:"kind"`
	DayOfMonth int64  `json:"dayOfMonth"`
}

// PutProgramsIdPayoutScheduleBodyScheduleOption1Option1Option2 is generated from the OpenAPI component of the same name.
type PutProgramsIdPayoutScheduleBodyScheduleOption1Option1Option2 struct {
	Kind      string `json:"kind"`
	DayOfWeek int64  `json:"dayOfWeek"`
}

// PutProgramsIdPayoutScheduleBodyScheduleOption1Option1Option3 is generated from the OpenAPI component of the same name.
type PutProgramsIdPayoutScheduleBodyScheduleOption1Option1Option3 struct {
	Kind string  `json:"kind"`
	Days []int64 `json:"days"`
}

// PutProgramsIdPayoutScheduleBodyScheduleOption1Option1Option4 is generated from the OpenAPI component of the same name.
type PutProgramsIdPayoutScheduleBodyScheduleOption1Option1Option4 struct {
	Kind  string   `json:"kind"`
	Dates []string `json:"dates"`
}

// PutProgramsIdPayoutScheduleBodyScheduleOption1Option2 is generated from the OpenAPI component of the same name.
type PutProgramsIdPayoutScheduleBodyScheduleOption1Option2 struct {
	Kind string `json:"kind"`
}

// PutProgramsIdPayoutScheduleBody is generated from the OpenAPI component of the same name.
type PutProgramsIdPayoutScheduleBody struct {
	Schedule json.RawMessage `json:"schedule"`
}

// PutOrgTaxProfileBody is generated from the OpenAPI component of the same name.
type PutOrgTaxProfileBody struct {
	LegalName      *string `json:"legalName,omitempty"`
	Address        *string `json:"address,omitempty"`
	Country        *string `json:"country,omitempty"`
	VATID          *string `json:"vatId,omitempty"`
	Ein            *string `json:"ein,omitempty"`
	DomesticVATBps *int64  `json:"domesticVatBps,omitempty"`
	ViesCheck      *bool   `json:"viesCheck,omitempty"`
}

// PutPartnersIdVatBody is generated from the OpenAPI component of the same name.
type PutPartnersIdVatBody struct {
	VATID *string `json:"vatId"`
}

// PostPartnerOrgsOrgIdSelfBillingBody is generated from the OpenAPI component of the same name.
type PostPartnerOrgsOrgIdSelfBillingBody struct {
	Accept bool `json:"accept"`
}

// PutPartnerOrgsOrgIdVatBody is generated from the OpenAPI component of the same name.
type PutPartnerOrgsOrgIdVatBody struct {
	VATID *string `json:"vatId"`
}

// PutOrgTaxFilingBodyCredentials is generated from the OpenAPI component of the same name.
type PutOrgTaxFilingBodyCredentials struct {
	APIKey  string  `json:"apiKey"`
	BaseURL *string `json:"baseUrl,omitempty"`
}

// PutOrgTaxFilingBody is generated from the OpenAPI component of the same name.
type PutOrgTaxFilingBody struct {
	Provider    string                          `json:"provider"`
	Credentials *PutOrgTaxFilingBodyCredentials `json:"credentials,omitempty"`
}

// PostTax1099FilingsBody is generated from the OpenAPI component of the same name.
type PostTax1099FilingsBody struct {
	Year int64 `json:"year"`
}

// PostAdminStatusIncidentsBody is generated from the OpenAPI component of the same name.
type PostAdminStatusIncidentsBody struct {
	Title      string   `json:"title"`
	Severity   *string  `json:"severity,omitempty"`
	State      *string  `json:"state,omitempty"`
	Body       *string  `json:"body,omitempty"`
	Components []string `json:"components,omitempty"`
	StartedAt  *string  `json:"startedAt,omitempty"`
}

// PatchAdminStatusIncidentsIdBody is generated from the OpenAPI component of the same name.
type PatchAdminStatusIncidentsIdBody struct {
	Title      *string  `json:"title,omitempty"`
	Severity   *string  `json:"severity,omitempty"`
	State      *string  `json:"state,omitempty"`
	Body       *string  `json:"body,omitempty"`
	Components []string `json:"components,omitempty"`
}

// PostAdminStatusIncidentsIdUpdatesBody is generated from the OpenAPI component of the same name.
type PostAdminStatusIncidentsIdUpdatesBody struct {
	Body  string `json:"body"`
	State string `json:"state"`
}

// PutOrgRetentionBody is generated from the OpenAPI component of the same name.
type PutOrgRetentionBody struct {
	Clicks            *int64 `json:"clicks,omitempty"`
	Visits            *int64 `json:"visits,omitempty"`
	BillingPayloads   *int64 `json:"billingPayloads,omitempty"`
	WebhookDeliveries *int64 `json:"webhookDeliveries,omitempty"`
	Notifications     *int64 `json:"notifications,omitempty"`
	AuditLog          *int64 `json:"auditLog,omitempty"`
}

// PutOrgReferralParamsBody is generated from the OpenAPI component of the same name.
type PutOrgReferralParamsBody struct {
	Params []string `json:"params"`
}

// PostProgramsBodyPublicPage is generated from the OpenAPI component of the same name.
type PostProgramsBodyPublicPage struct {
	Pitch       *string `json:"pitch,omitempty"`
	AutoApprove *bool   `json:"autoApprove,omitempty"`
}

// PostProgramsBodyRiskPolicy is generated from the OpenAPI component of the same name.
type PostProgramsBodyRiskPolicy struct {
	HoldBotScore        *float64 `json:"holdBotScore,omitempty"`
	HoldRiskScore       *float64 `json:"holdRiskScore,omitempty"`
	HoldHighRiskLevel   *bool    `json:"holdHighRiskLevel,omitempty"`
	HoldDisposableEmail *bool    `json:"holdDisposableEmail,omitempty"`
	HoldVelocity        *int64   `json:"holdVelocity,omitempty"`
	HoldUnverifiedEmail *bool    `json:"holdUnverifiedEmail,omitempty"`
	HoldSelfReferral    *bool    `json:"holdSelfReferral,omitempty"`
}

// PostProgramsBodyWidget is generated from the OpenAPI component of the same name.
type PostProgramsBodyWidget struct {
	AutoEnroll     *bool   `json:"autoEnroll,omitempty"`
	DestinationURL *string `json:"destinationUrl,omitempty"`
}

// PostProgramsBodyMobileAppOption1IosOption1 is generated from the OpenAPI component of the same name.
type PostProgramsBodyMobileAppOption1IosOption1 struct {
	TeamID   string `json:"teamId"`
	BundleID string `json:"bundleId"`
}

// PostProgramsBodyMobileAppOption1AndroidOption1 is generated from the OpenAPI component of the same name.
type PostProgramsBodyMobileAppOption1AndroidOption1 struct {
	PackageName            string   `json:"packageName"`
	Sha256CertFingerprints []string `json:"sha256CertFingerprints"`
}

// PostProgramsBodyMobileAppOption1 is generated from the OpenAPI component of the same name.
type PostProgramsBodyMobileAppOption1 struct {
	Ios     *PostProgramsBodyMobileAppOption1IosOption1     `json:"ios,omitempty"`
	Android *PostProgramsBodyMobileAppOption1AndroidOption1 `json:"android,omitempty"`
}

// PostProgramsBody is generated from the OpenAPI component of the same name.
type PostProgramsBody struct {
	Name                   string                            `json:"name"`
	Slug                   string                            `json:"slug"`
	AttributionModel       *string                           `json:"attributionModel,omitempty"`
	CookieWindowDays       *int64                            `json:"cookieWindowDays,omitempty"`
	SignupToConversionDays *int64                            `json:"signupToConversionDays,omitempty"`
	TermsMd                *string                           `json:"termsMd,omitempty"`
	Status                 *string                           `json:"status,omitempty"`
	PublicHandle           *string                           `json:"publicHandle,omitempty"`
	PublicPage             *PostProgramsBodyPublicPage       `json:"publicPage,omitempty"`
	ProbabilisticPolicy    *string                           `json:"probabilisticPolicy,omitempty"`
	SelfReferralPolicy     *string                           `json:"selfReferralPolicy,omitempty"`
	SplitCommissions       *bool                             `json:"splitCommissions,omitempty"`
	UTMTemplate            map[string]string                 `json:"utmTemplate,omitempty"`
	AllowPaidSearch        *bool                             `json:"allowPaidSearch,omitempty"`
	ChallengeLinks         *bool                             `json:"challengeLinks,omitempty"`
	Settlement             *string                           `json:"settlement,omitempty"`
	RiskPolicy             *PostProgramsBodyRiskPolicy       `json:"riskPolicy,omitempty"`
	Widget                 *PostProgramsBodyWidget           `json:"widget,omitempty"`
	MobileApp              *PostProgramsBodyMobileAppOption1 `json:"mobileApp,omitempty"`
}

// PatchProgramsIdBodyPublicPage is generated from the OpenAPI component of the same name.
type PatchProgramsIdBodyPublicPage struct {
	Pitch       *string `json:"pitch,omitempty"`
	AutoApprove *bool   `json:"autoApprove,omitempty"`
}

// PatchProgramsIdBodyRiskPolicy is generated from the OpenAPI component of the same name.
type PatchProgramsIdBodyRiskPolicy struct {
	HoldBotScore        *float64 `json:"holdBotScore,omitempty"`
	HoldRiskScore       *float64 `json:"holdRiskScore,omitempty"`
	HoldHighRiskLevel   *bool    `json:"holdHighRiskLevel,omitempty"`
	HoldDisposableEmail *bool    `json:"holdDisposableEmail,omitempty"`
	HoldVelocity        *int64   `json:"holdVelocity,omitempty"`
	HoldUnverifiedEmail *bool    `json:"holdUnverifiedEmail,omitempty"`
	HoldSelfReferral    *bool    `json:"holdSelfReferral,omitempty"`
}

// PatchProgramsIdBodyWidget is generated from the OpenAPI component of the same name.
type PatchProgramsIdBodyWidget struct {
	AutoEnroll     *bool   `json:"autoEnroll,omitempty"`
	DestinationURL *string `json:"destinationUrl,omitempty"`
}

// PatchProgramsIdBodyMobileAppOption1IosOption1 is generated from the OpenAPI component of the same name.
type PatchProgramsIdBodyMobileAppOption1IosOption1 struct {
	TeamID   string `json:"teamId"`
	BundleID string `json:"bundleId"`
}

// PatchProgramsIdBodyMobileAppOption1AndroidOption1 is generated from the OpenAPI component of the same name.
type PatchProgramsIdBodyMobileAppOption1AndroidOption1 struct {
	PackageName            string   `json:"packageName"`
	Sha256CertFingerprints []string `json:"sha256CertFingerprints"`
}

// PatchProgramsIdBodyMobileAppOption1 is generated from the OpenAPI component of the same name.
type PatchProgramsIdBodyMobileAppOption1 struct {
	Ios     *PatchProgramsIdBodyMobileAppOption1IosOption1     `json:"ios,omitempty"`
	Android *PatchProgramsIdBodyMobileAppOption1AndroidOption1 `json:"android,omitempty"`
}

// PatchProgramsIdBody is generated from the OpenAPI component of the same name.
type PatchProgramsIdBody struct {
	Name                   *string                              `json:"name,omitempty"`
	Slug                   *string                              `json:"slug,omitempty"`
	AttributionModel       *string                              `json:"attributionModel,omitempty"`
	CookieWindowDays       *int64                               `json:"cookieWindowDays,omitempty"`
	SignupToConversionDays *int64                               `json:"signupToConversionDays,omitempty"`
	TermsMd                *string                              `json:"termsMd,omitempty"`
	Status                 *string                              `json:"status,omitempty"`
	PublicHandle           *string                              `json:"publicHandle,omitempty"`
	PublicPage             *PatchProgramsIdBodyPublicPage       `json:"publicPage,omitempty"`
	ProbabilisticPolicy    *string                              `json:"probabilisticPolicy,omitempty"`
	SelfReferralPolicy     *string                              `json:"selfReferralPolicy,omitempty"`
	SplitCommissions       *bool                                `json:"splitCommissions,omitempty"`
	UTMTemplate            map[string]string                    `json:"utmTemplate,omitempty"`
	AllowPaidSearch        *bool                                `json:"allowPaidSearch,omitempty"`
	ChallengeLinks         *bool                                `json:"challengeLinks,omitempty"`
	Settlement             *string                              `json:"settlement,omitempty"`
	RiskPolicy             *PatchProgramsIdBodyRiskPolicy       `json:"riskPolicy,omitempty"`
	Widget                 *PatchProgramsIdBodyWidget           `json:"widget,omitempty"`
	MobileApp              *PatchProgramsIdBodyMobileAppOption1 `json:"mobileApp,omitempty"`
}

// PostProgramsIdEnrollmentsBody is generated from the OpenAPI component of the same name.
type PostProgramsIdEnrollmentsBody struct {
	PartnerID string `json:"partnerId"`
}

// PostPartnersBody is generated from the OpenAPI component of the same name.
type PostPartnersBody struct {
	Email       string   `json:"email"`
	DisplayName *string  `json:"displayName,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Notes       *string  `json:"notes,omitempty"`
	ExternalRef *string  `json:"externalRef,omitempty"`
	Upsert      *bool    `json:"upsert,omitempty"`
}

// PatchPartnersIdBody is generated from the OpenAPI component of the same name.
type PatchPartnersIdBody struct {
	DisplayName *string  `json:"displayName,omitempty"`
	Status      *string  `json:"status,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Notes       *string  `json:"notes,omitempty"`
	ExternalRef *string  `json:"externalRef,omitempty"`
}

// PostApiKeysBody is generated from the OpenAPI component of the same name.
type PostApiKeysBody struct {
	Name           string   `json:"name"`
	Scopes         []string `json:"scopes,omitempty"`
	Publishable    *bool    `json:"publishable,omitempty"`
	AllowedOrigins []string `json:"allowedOrigins,omitempty"`
	APIVersion     *string  `json:"apiVersion,omitempty"`
	Mode           *string  `json:"mode,omitempty"`
	Owner          *string  `json:"owner,omitempty"`
	ExpiresInDays  *int64   `json:"expiresInDays,omitempty"`
}

// PatchApiKeysIdBody is generated from the OpenAPI component of the same name.
type PatchApiKeysIdBody struct {
	APIVersion    *string `json:"apiVersion,omitempty"`
	Owner         *string `json:"owner,omitempty"`
	ExpiresInDays *int64  `json:"expiresInDays,omitempty"`
}

// PostTestClockAdvanceBody is generated from the OpenAPI component of the same name.
type PostTestClockAdvanceBody struct {
	Seconds int64 `json:"seconds"`
}

// PostCliDeviceCodesBody is generated from the OpenAPI component of the same name.
type PostCliDeviceCodesBody struct {
	ClientName *string `json:"clientName,omitempty"`
}

// PostCliDeviceCodesTokenBody is generated from the OpenAPI component of the same name.
type PostCliDeviceCodesTokenBody struct {
	DeviceCode string `json:"deviceCode"`
}

// PostCliDeviceCodesUserCodeApproveBody is generated from the OpenAPI component of the same name.
type PostCliDeviceCodesUserCodeApproveBody struct {
	Access *string `json:"access,omitempty"`
	Name   *string `json:"name,omitempty"`
}

// PostCliListenersBody is generated from the OpenAPI component of the same name.
type PostCliListenersBody struct {
	Topics    []string `json:"topics,omitempty"`
	ForwardTo *string  `json:"forwardTo,omitempty"`
}

// PostCliTriggerBody is generated from the OpenAPI component of the same name.
type PostCliTriggerBody struct {
	Topic string         `json:"topic"`
	Data  map[string]any `json:"data,omitempty"`
}

// PostLinksBody is generated from the OpenAPI component of the same name.
type PostLinksBody struct {
	ProgramID      string  `json:"programId"`
	PartnerID      string  `json:"partnerId"`
	Code           *string `json:"code,omitempty"`
	DestinationURL string  `json:"destinationUrl"`
}

// PostCouponsBody is generated from the OpenAPI component of the same name.
type PostCouponsBody struct {
	ProgramID string `json:"programId"`
	PartnerID string `json:"partnerId"`
	Code      string `json:"code"`
}

// PostWidgetTokensBody is generated from the OpenAPI component of the same name.
type PostWidgetTokensBody struct {
	ProgramID          string  `json:"programId"`
	CustomerRef        *string `json:"customerRef,omitempty"`
	PartnerExternalRef *string `json:"partnerExternalRef,omitempty"`
	Email              *string `json:"email,omitempty"`
	DisplayName        *string `json:"displayName,omitempty"`
	TtlSeconds         *int64  `json:"ttlSeconds,omitempty"`
}

// PostReferralsAssociateBodyConsent is generated from the OpenAPI component of the same name.
type PostReferralsAssociateBodyConsent struct {
	Provider  string   `json:"provider"` // The adapter that read it (`onetrust`, `tcf`, `cookiemunch`, …), `custom` for `aff.consent()`, `none`, or `no-cmp`.
	State     string   `json:"state"`
	At        *string  `json:"at,omitempty"`    // When the tracker read it.
	Stamp     *string  `json:"stamp,omitempty"` // The CMP's consent record id, when it has one.
	TcString  *string  `json:"tcString,omitempty"`
	GppString *string  `json:"gppString,omitempty"`
	Purposes  []string `json:"purposes,omitempty"`
}

// PostReferralsAssociateBody is generated from the OpenAPI component of the same name.
type PostReferralsAssociateBody struct {
	ProgramID      string                             `json:"programId"`
	CustomerRef    string                             `json:"customerRef"`
	AnonID         *string                            `json:"anonId,omitempty"`
	ReferralToken  *string                            `json:"referralToken,omitempty"`
	CouponCode     *string                            `json:"couponCode,omitempty"`
	EmailHash      *string                            `json:"emailHash,omitempty"`
	IP             *string                            `json:"ip,omitempty"`
	UA             *string                            `json:"ua,omitempty"`
	IdempotencyKey *string                            `json:"idempotencyKey,omitempty"`
	OccurredAt     *string                            `json:"occurredAt,omitempty"`
	EmailDomain    *string                            `json:"emailDomain,omitempty"`
	Consent        *PostReferralsAssociateBodyConsent `json:"consent,omitempty"`
}

// PatchReferralsIdBody is generated from the OpenAPI component of the same name.
type PatchReferralsIdBody struct {
	Status string `json:"status"`
}

// PostReferralsIdReassignBody is generated from the OpenAPI component of the same name.
type PostReferralsIdReassignBody struct {
	PartnerID string `json:"partnerId"`
}

// PostAttributionReplayBody is generated from the OpenAPI component of the same name.
type PostAttributionReplayBody struct {
	ProgramID string `json:"programId"`
}

// PostTrackEventsBodyOption1 is generated from the OpenAPI component of the same name.
type PostTrackEventsBodyOption1 struct {
	ProgramID string  `json:"programId"`
	PartnerID string  `json:"partnerId"`
	LinkID    *string `json:"linkId,omitempty"`
	URL       *string `json:"url,omitempty"`
	AnonID    *string `json:"anonId,omitempty"`
	IP        *string `json:"ip,omitempty"`
	UA        *string `json:"ua,omitempty"`
	Type      string  `json:"type"`
}

// PostTrackEventsBodyOption2 is generated from the OpenAPI component of the same name.
type PostTrackEventsBodyOption2 struct {
	ProgramID string  `json:"programId"`
	PartnerID string  `json:"partnerId"`
	LinkID    *string `json:"linkId,omitempty"`
	URL       *string `json:"url,omitempty"`
	AnonID    *string `json:"anonId,omitempty"`
	IP        *string `json:"ip,omitempty"`
	UA        *string `json:"ua,omitempty"`
	Type      string  `json:"type"`
}

// PostTrackEventsBodyOption3 is generated from the OpenAPI component of the same name.
type PostTrackEventsBodyOption3 struct {
	Type           string  `json:"type"`
	ProgramID      string  `json:"programId"`
	CustomerRef    string  `json:"customerRef"`
	Event          string  `json:"event"`
	OccurredAt     *string `json:"occurredAt,omitempty"`
	IdempotencyKey *string `json:"idempotencyKey,omitempty"`
}

// PostCommissionRulesBodyConfigBaseOption1 is generated from the OpenAPI component of the same name.
type PostCommissionRulesBodyConfigBaseOption1 struct {
	Kind string `json:"kind"`
	Bps  int64  `json:"bps"`
}

// PostCommissionRulesBodyConfigBaseOption2 is generated from the OpenAPI component of the same name.
type PostCommissionRulesBodyConfigBaseOption2 struct {
	Kind        string `json:"kind"`
	AmountMinor int64  `json:"amountMinor"`
	Currency    string `json:"currency"`
}

// PostCommissionRulesBodyConfigRecurringOption1 is generated from the OpenAPI component of the same name.
type PostCommissionRulesBodyConfigRecurringOption1 struct {
	Type string `json:"type"`
}

// PostCommissionRulesBodyConfigRecurringOption2 is generated from the OpenAPI component of the same name.
type PostCommissionRulesBodyConfigRecurringOption2 struct {
	Type string `json:"type"`
}

// PostCommissionRulesBodyConfigRecurringOption3 is generated from the OpenAPI component of the same name.
type PostCommissionRulesBodyConfigRecurringOption3 struct {
	Type   string `json:"type"`
	Months int64  `json:"months"`
}

// PostCommissionRulesBodyConfigConditionsMetadataItem is generated from the OpenAPI component of the same name.
type PostCommissionRulesBodyConfigConditionsMetadataItem struct {
	Key    string            `json:"key"`
	Equals json.RawMessage   `json:"equals,omitempty"`
	In     []json.RawMessage `json:"in,omitempty"`
}

// PostCommissionRulesBodyConfigConditions is generated from the OpenAPI component of the same name.
type PostCommissionRulesBodyConfigConditions struct {
	Countries        []string                                              `json:"countries,omitempty"`
	ExcludeCountries []string                                              `json:"excludeCountries,omitempty"`
	Metadata         []PostCommissionRulesBodyConfigConditionsMetadataItem `json:"metadata,omitempty"`
}

// PostCommissionRulesBodyConfigOverridesItemMatchMetadataItem is generated from the OpenAPI component of the same name.
type PostCommissionRulesBodyConfigOverridesItemMatchMetadataItem struct {
	Key    string            `json:"key"`
	Equals json.RawMessage   `json:"equals,omitempty"`
	In     []json.RawMessage `json:"in,omitempty"`
}

// PostCommissionRulesBodyConfigOverridesItemMatch is generated from the OpenAPI component of the same name.
type PostCommissionRulesBodyConfigOverridesItemMatch struct {
	PartnerID  *string                                                       `json:"partnerId,omitempty"`
	LinkID     *string                                                       `json:"linkId,omitempty"`
	ProductRef *string                                                       `json:"productRef,omitempty"`
	Countries  []string                                                      `json:"countries,omitempty"`
	Metadata   []PostCommissionRulesBodyConfigOverridesItemMatchMetadataItem `json:"metadata,omitempty"`
}

// PostCommissionRulesBodyConfigOverridesItemBaseOption1 is generated from the OpenAPI component of the same name.
type PostCommissionRulesBodyConfigOverridesItemBaseOption1 struct {
	Kind string `json:"kind"`
	Bps  int64  `json:"bps"`
}

// PostCommissionRulesBodyConfigOverridesItemBaseOption2 is generated from the OpenAPI component of the same name.
type PostCommissionRulesBodyConfigOverridesItemBaseOption2 struct {
	Kind        string `json:"kind"`
	AmountMinor int64  `json:"amountMinor"`
	Currency    string `json:"currency"`
}

// PostCommissionRulesBodyConfigOverridesItem is generated from the OpenAPI component of the same name.
type PostCommissionRulesBodyConfigOverridesItem struct {
	Match PostCommissionRulesBodyConfigOverridesItemMatch `json:"match"`
	Base  json.RawMessage                                 `json:"base"`
}

// PostCommissionRulesBodyConfigTiersItem is generated from the OpenAPI component of the same name.
type PostCommissionRulesBodyConfigTiersItem struct {
	MinCustomers int64 `json:"minCustomers"`
	Bps          int64 `json:"bps"`
}

// PostCommissionRulesBodyConfigBountiesItem is generated from the OpenAPI component of the same name.
type PostCommissionRulesBodyConfigBountiesItem struct {
	Event       string `json:"event"`
	AmountMinor int64  `json:"amountMinor"`
	Currency    string `json:"currency"`
}

// PostCommissionRulesBodyConfigSubAffiliate is generated from the OpenAPI component of the same name.
type PostCommissionRulesBodyConfigSubAffiliate struct {
	Levels   int64   `json:"levels"`
	RatesBps []int64 `json:"ratesBps"`
}

// PostCommissionRulesBodyConfigNewVsExpansion is generated from the OpenAPI component of the same name.
type PostCommissionRulesBodyConfigNewVsExpansion struct {
	NewBps       int64 `json:"newBps"`
	ExpansionBps int64 `json:"expansionBps"`
}

// PostCommissionRulesBodyConfig is generated from the OpenAPI component of the same name.
type PostCommissionRulesBodyConfig struct {
	Base           json.RawMessage                              `json:"base"`
	AppliesTo      *string                                      `json:"appliesTo,omitempty"`
	Recurring      json.RawMessage                              `json:"recurring,omitempty"`
	HoldDays       *int64                                       `json:"holdDays,omitempty"`
	Rounding       *string                                      `json:"rounding,omitempty"`
	Conditions     *PostCommissionRulesBodyConfigConditions     `json:"conditions,omitempty"`
	Overrides      []PostCommissionRulesBodyConfigOverridesItem `json:"overrides,omitempty"`
	Tiers          []PostCommissionRulesBodyConfigTiersItem     `json:"tiers,omitempty"`
	Bounties       []PostCommissionRulesBodyConfigBountiesItem  `json:"bounties,omitempty"`
	SubAffiliate   *PostCommissionRulesBodyConfigSubAffiliate   `json:"subAffiliate,omitempty"`
	NewVsExpansion *PostCommissionRulesBodyConfigNewVsExpansion `json:"newVsExpansion,omitempty"`
}

// PostCommissionRulesBody is generated from the OpenAPI component of the same name.
type PostCommissionRulesBody struct {
	ProgramID string                        `json:"programId"`
	Config    PostCommissionRulesBodyConfig `json:"config"`
}

// PostCommissionsIdAdjustBody is generated from the OpenAPI component of the same name.
type PostCommissionsIdAdjustBody struct {
	AmountMinor   int64  `json:"amountMinor"`
	Memo          string `json:"memo"`
	AllowNegative *bool  `json:"allowNegative,omitempty"`
}

// PostCommissionsIdSettleBody is generated from the OpenAPI component of the same name.
type PostCommissionsIdSettleBody struct {
	Method    *string `json:"method,omitempty"`
	Reference string  `json:"reference"`
	Memo      *string `json:"memo,omitempty"`
}

// PostBillingConnectionsBodyFilter is generated from the OpenAPI component of the same name.
type PostBillingConnectionsBodyFilter struct {
	Metadata   map[string]string `json:"metadata,omitempty"`
	ProductIds []string          `json:"productIds,omitempty"`
	PriceIds   []string          `json:"priceIds,omitempty"`
	Skus       []string          `json:"skus,omitempty"`
}

// PostBillingConnectionsBody is generated from the OpenAPI component of the same name.
type PostBillingConnectionsBody struct {
	Provider      string                            `json:"provider"`
	Name          *string                           `json:"name,omitempty"`
	WebhookSecret *string                           `json:"webhookSecret,omitempty"`
	Filter        *PostBillingConnectionsBodyFilter `json:"filter,omitempty"`
}

// PutBillingConnectionsIdSecretBody is generated from the OpenAPI component of the same name.
type PutBillingConnectionsIdSecretBody struct {
	WebhookSecret string `json:"webhookSecret"`
}

// PutBillingConnectionsIdCredentialsBody is generated from the OpenAPI component of the same name.
type PutBillingConnectionsIdCredentialsBody struct {
	APIKey    *string `json:"apiKey,omitempty"`
	APIToken  *string `json:"apiToken,omitempty"`
	StoreHash *string `json:"storeHash,omitempty"`
	Region    *string `json:"region,omitempty"`
}

// PatchBillingConnectionsIdBodyFilterOption1 is generated from the OpenAPI component of the same name.
type PatchBillingConnectionsIdBodyFilterOption1 struct {
	Metadata   map[string]string `json:"metadata,omitempty"`
	ProductIds []string          `json:"productIds,omitempty"`
	PriceIds   []string          `json:"priceIds,omitempty"`
	Skus       []string          `json:"skus,omitempty"`
}

// PatchBillingConnectionsIdBody is generated from the OpenAPI component of the same name.
type PatchBillingConnectionsIdBody struct {
	Status *string                                     `json:"status,omitempty"`
	Name   *string                                     `json:"name,omitempty"`
	Filter *PatchBillingConnectionsIdBodyFilterOption1 `json:"filter,omitempty"`
}

// PutBillingConnectionsIdValueMappingBody is generated from the OpenAPI component of the same name.
type PutBillingConnectionsIdValueMappingBody struct {
	Cash          *string `json:"cash,omitempty"`
	CreditApplied *string `json:"creditApplied,omitempty"`
	Discount      *string `json:"discount,omitempty"`
	Cost          *string `json:"cost,omitempty"`
	ProductRef    *string `json:"productRef,omitempty"`
	Unit          string  `json:"unit"`
}

// PostBillingConnectionsIdValueMappingPreviewBodyMapping is generated from the OpenAPI component of the same name.
type PostBillingConnectionsIdValueMappingPreviewBodyMapping struct {
	Cash          *string `json:"cash,omitempty"`
	CreditApplied *string `json:"creditApplied,omitempty"`
	Discount      *string `json:"discount,omitempty"`
	Cost          *string `json:"cost,omitempty"`
	ProductRef    *string `json:"productRef,omitempty"`
	Unit          string  `json:"unit"`
}

// PostBillingConnectionsIdValueMappingPreviewBody is generated from the OpenAPI component of the same name.
type PostBillingConnectionsIdValueMappingPreviewBody struct {
	Payload  any                                                     `json:"payload"`
	Currency *string                                                 `json:"currency,omitempty"`
	Mapping  *PostBillingConnectionsIdValueMappingPreviewBodyMapping `json:"mapping,omitempty"`
}

// PostBillingInternalCustomersBody is generated from the OpenAPI component of the same name.
type PostBillingInternalCustomersBody struct {
	CustomerRef string  `json:"customerRef"`
	Note        *string `json:"note,omitempty"`
}

// PutPartnersIdPayoutProfileBody is generated from the OpenAPI component of the same name.
type PutPartnersIdPayoutProfileBody struct {
	Method            string         `json:"method"`
	Details           map[string]any `json:"details,omitempty"`
	PayoutCurrency    *string        `json:"payoutCurrency,omitempty"`
	MinThresholdMinor *int64         `json:"minThresholdMinor,omitempty"`
}

// PatchPartnersIdPayoutProfileBody is generated from the OpenAPI component of the same name.
type PatchPartnersIdPayoutProfileBody struct {
	MinThresholdMinor int64 `json:"minThresholdMinor"`
}

// PostPartnersIdTaxFormBody is generated from the OpenAPI component of the same name.
type PostPartnersIdTaxFormBody struct {
	Type string `json:"type"`
	Data any    `json:"data"`
}

// PostFxRatesBody is generated from the OpenAPI component of the same name.
type PostFxRatesBody struct {
	Base        string          `json:"base"`
	Quote       string          `json:"quote"`
	Rate        json.RawMessage `json:"rate"`
	EffectiveAt *string         `json:"effectiveAt,omitempty"`
}

// PutOrgFxFeedBody is generated from the OpenAPI component of the same name.
type PutOrgFxFeedBody struct {
	Enabled bool `json:"enabled"`
}

// PostFxRatesClearOverrideBody is generated from the OpenAPI component of the same name.
type PostFxRatesClearOverrideBody struct {
	Base  string `json:"base"`
	Quote string `json:"quote"`
}

// PutPayoutConnectionsProviderBodyOption1 is generated from the OpenAPI component of the same name.
type PutPayoutConnectionsProviderBodyOption1 struct {
	ClientID string  `json:"clientId"`
	Secret   string  `json:"secret"`
	Env      *string `json:"env,omitempty"`
}

// PutPayoutConnectionsProviderBodyOption2 is generated from the OpenAPI component of the same name.
type PutPayoutConnectionsProviderBodyOption2 struct {
	SecretKey string  `json:"secretKey"`
	Env       *string `json:"env,omitempty"`
}

// PutPayoutConnectionsProviderBodyOption3 is generated from the OpenAPI component of the same name.
type PutPayoutConnectionsProviderBodyOption3 struct {
	APIKey       string  `json:"apiKey"`
	WalletID     string  `json:"walletId"`
	DefaultChain *string `json:"defaultChain,omitempty"`
	Env          *string `json:"env,omitempty"`
}

// PutPayoutConnectionsProviderBodyOption4 is generated from the OpenAPI component of the same name.
type PutPayoutConnectionsProviderBodyOption4 struct {
	APIToken  string  `json:"apiToken"`
	ProfileID string  `json:"profileId"`
	Env       *string `json:"env,omitempty"`
}

// PutPayoutConnectionsProviderBodyOption5 is generated from the OpenAPI component of the same name.
type PutPayoutConnectionsProviderBodyOption5 struct {
	ClientID  string  `json:"clientId"`
	Secret    string  `json:"secret"`
	ProgramID string  `json:"programId"`
	Env       *string `json:"env,omitempty"`
}

// PutPayoutConnectionsProviderBodyOption6 is generated from the OpenAPI component of the same name.
type PutPayoutConnectionsProviderBodyOption6 struct {
	APIKey          string  `json:"apiKey"`
	FundingSourceID string  `json:"fundingSourceId"`
	CampaignID      *string `json:"campaignId,omitempty"`
	Env             *string `json:"env,omitempty"`
}

// PostFraudInvalidTrafficScanBody is generated from the OpenAPI component of the same name.
type PostFraudInvalidTrafficScanBody struct {
	WindowHours *int64 `json:"windowHours,omitempty"`
}

// PostFraudCasesIdResolveBody is generated from the OpenAPI component of the same name.
type PostFraudCasesIdResolveBody struct {
	Action string `json:"action"`
}

// PutProfitabilitySettingsBody is generated from the OpenAPI component of the same name.
type PutProfitabilitySettingsBody struct {
	ReportingCurrency        *string  `json:"reportingCurrency,omitempty"`
	DefaultMarginBps         *int64   `json:"defaultMarginBps,omitempty"`
	CreditCostModel          *string  `json:"creditCostModel,omitempty"`
	CreditCostRatioBps       *int64   `json:"creditCostRatioBps,omitempty"`
	Inference                *string  `json:"inference,omitempty"`
	CreditDraw               *string  `json:"creditDraw,omitempty"`
	WindowDays               *int64   `json:"windowDays,omitempty"`
	ChurnWindowDays          *int64   `json:"churnWindowDays,omitempty"`
	CountReferrerSpend       *bool    `json:"countReferrerSpend,omitempty"`
	CostToServeMinorPerMonth *int64   `json:"costToServeMinorPerMonth,omitempty"`
	TaxFeesReduceProfit      *bool    `json:"taxFeesReduceProfit,omitempty"`
	ValuePatternWeight       *float64 `json:"valuePatternWeight,omitempty"`
	TargetRoiBps             *int64   `json:"targetRoiBps,omitempty"`
	ProgramID                *string  `json:"programId,omitempty"`
}

// PutProfitabilityProductsProductRefBody is generated from the OpenAPI component of the same name.
type PutProfitabilityProductsProductRefBody struct {
	MarginBps     *int64  `json:"marginBps,omitempty"`
	UnitCostMinor *int64  `json:"unitCostMinor,omitempty"`
	Currency      *string `json:"currency,omitempty"`
}

// PostRoiPatternsPatternReviewBody is generated from the OpenAPI component of the same name.
type PostRoiPatternsPatternReviewBody struct {
	CustomerRef *string `json:"customerRef,omitempty"`
	PartnerID   *string `json:"partnerId,omitempty"`
	ProgramID   *string `json:"programId,omitempty"`
}

// PostValueEventsBodyEventsItem is generated from the OpenAPI component of the same name.
type PostValueEventsBodyEventsItem struct {
	Kind               string         `json:"kind"`
	CustomerRef        *string        `json:"customerRef,omitempty"`        // Your id for the customer. One of customerRef, partnerId or partnerExternalRef is required.
	PartnerID          *string        `json:"partnerId,omitempty"`          // A Referalli partner id.
	PartnerExternalRef *string        `json:"partnerExternalRef,omitempty"` // Your id for the partner (the partner's externalRef).
	CreditGrantID      *string        `json:"creditGrantId,omitempty"`      // The credit grant this event draws on.
	ReferralID         *string        `json:"referralId,omitempty"`         // The referral it belongs to.
	ProgramID          *string        `json:"programId,omitempty"`          // The program it belongs to.
	ProductRef         *string        `json:"productRef,omitempty"`         // Your product or plan id.
	Currency           string         `json:"currency"`                     // ISO 4217 code.
	CashMinor          *int64         `json:"cashMinor,omitempty"`          // Cash paid (a refund is stored negative). Integer minor units.
	CreditAppliedMinor *int64         `json:"creditAppliedMinor,omitempty"` // Credit applied to this payment. Integer minor units.
	DiscountMinor      *int64         `json:"discountMinor,omitempty"`      // Discount given. Integer minor units.
	TaxMinor           *int64         `json:"taxMinor,omitempty"`           // Tax included. Integer minor units.
	FeesMinor          *int64         `json:"feesMinor,omitempty"`          // Payment fees. Integer minor units.
	CostMinor          *int64         `json:"costMinor,omitempty"`          // Cost to you (cost events). Integer minor units.
	FaceMinor          *int64         `json:"faceMinor,omitempty"`          // Face value of the credit (required for credit_granted). Integer minor units.
	ExternalID         string         `json:"externalId"`                   // Your idempotency id for this fact: the same externalId is recorded once.
	InvoiceRef         *string        `json:"invoiceRef,omitempty"`         // Your invoice id.
	Provider           *string        `json:"provider,omitempty"`           // Where the fact came from (stripe, your billing system…).
	OccurredAt         string         `json:"occurredAt"`                   // ISO 8601 date or date-time.
	Role               *string        `json:"role,omitempty"`               // credit_granted only: who received the credit.
	Metadata           map[string]any `json:"metadata,omitempty"`
}

// PostValueEventsBody: Up to 500 value events; each is validated and recorded on its own, so one bad event is reported in `errors` and the rest are recorded.
type PostValueEventsBody struct {
	Events []PostValueEventsBodyEventsItem `json:"events"`
}

// PostValueImportBody: CSV text (up to 10000 rows / 5 MB) plus a column-to-field mapping. "kind" applies to every row unless a column maps to kind. Up to 500 rows are recorded before the response; a larger file answers 202 with a jobId to poll at GET /v1/value/import/{jobId}. A dry run is always immediate.
type PostValueImportBody struct {
	CSV             string            `json:"csv"`
	Mapping         map[string]string `json:"mapping"`
	Kind            *string           `json:"kind,omitempty"`
	DryRun          *bool             `json:"dryRun,omitempty"`
	DefaultCurrency *string           `json:"defaultCurrency,omitempty"`
}

// PostPartnerKycSubmitBody is generated from the OpenAPI component of the same name.
type PostPartnerKycSubmitBody struct {
	LegalName   string  `json:"legalName"`
	Country     string  `json:"country"`
	DateOfBirth *string `json:"dateOfBirth,omitempty"`
	DocumentRef *string `json:"documentRef,omitempty"`
}

// PostPartnerLinksBody is generated from the OpenAPI component of the same name.
type PostPartnerLinksBody struct {
	OrgID          string  `json:"orgId"`
	ProgramID      string  `json:"programId"`
	DestinationURL string  `json:"destinationUrl"`
	AssetID        *string `json:"assetId,omitempty"`
	Code           *string `json:"code,omitempty"`
}

// PostPartnerDisputesBody is generated from the OpenAPI component of the same name.
type PostPartnerDisputesBody struct {
	OrgID      string  `json:"orgId"`
	ReferralID *string `json:"referralId,omitempty"`
	Kind       string  `json:"kind"`
	Detail     string  `json:"detail"`
}

// PutPartnerOrgsOrgIdPayoutProfileBody is generated from the OpenAPI component of the same name.
type PutPartnerOrgsOrgIdPayoutProfileBody struct {
	Method         string         `json:"method"`
	Details        map[string]any `json:"details,omitempty"`
	PayoutCurrency *string        `json:"payoutCurrency,omitempty"`
}

// PostPartnerOrgsOrgIdTaxFormBody is generated from the OpenAPI component of the same name.
type PostPartnerOrgsOrgIdTaxFormBody struct {
	Type string `json:"type"`
	Data any    `json:"data"`
}

// PostWebhookEndpointsBody is generated from the OpenAPI component of the same name.
type PostWebhookEndpointsBody struct {
	URL             string   `json:"url"`
	Topics          []string `json:"topics,omitempty"`
	LegacySignature *bool    `json:"legacySignature,omitempty"`
	ProgramIds      []string `json:"programIds,omitempty"`
	AlertContacts   []string `json:"alertContacts,omitempty"`
}

// PatchWebhookEndpointsIdBody is generated from the OpenAPI component of the same name.
type PatchWebhookEndpointsIdBody struct {
	Status          *string  `json:"status,omitempty"`
	URL             *string  `json:"url,omitempty"`
	Topics          []string `json:"topics,omitempty"`
	LegacySignature *bool    `json:"legacySignature,omitempty"`
	ProgramIds      []string `json:"programIds,omitempty"`
	AlertContacts   []string `json:"alertContacts,omitempty"`
}

// PostSandboxSimulateBody is generated from the OpenAPI component of the same name.
type PostSandboxSimulateBody struct {
	ProgramID  string  `json:"programId"`
	PartnerID  string  `json:"partnerId"`
	Scenario   *string `json:"scenario,omitempty"`
	PriceMinor *int64  `json:"priceMinor,omitempty"`
	Months     *int64  `json:"months,omitempty"`
}

// PostImportProviderBody is generated from the OpenAPI component of the same name.
type PostImportProviderBody struct {
	CSV string `json:"csv"` // The tool's partner export, as CSV text.
}

// PutBrandingBody is generated from the OpenAPI component of the same name.
type PutBrandingBody struct {
	LogoURL   *string           `json:"logoUrl,omitempty"`
	Theme     map[string]string `json:"theme,omitempty"`
	EmailFrom *string           `json:"emailFrom,omitempty"`
}

// PostBrandingLinkDomainBody is generated from the OpenAPI component of the same name.
type PostBrandingLinkDomainBody struct {
	Domain string `json:"domain"`
}

// PostBrandingDomainBody is generated from the OpenAPI component of the same name.
type PostBrandingDomainBody struct {
	Domain string `json:"domain"`
}

// PostEmbedTokensBody is generated from the OpenAPI component of the same name.
type PostEmbedTokensBody struct {
	PartnerEmail string `json:"partnerEmail"`
}

// PostDiscoverySearchBody is generated from the OpenAPI component of the same name.
type PostDiscoverySearchBody struct {
	Keywords []string `json:"keywords"`
	Category *string  `json:"category,omitempty"`
}

// PostDiscoveryCorpusBodyEntriesItem is generated from the OpenAPI component of the same name.
type PostDiscoveryCorpusBodyEntriesItem struct {
	Name            string   `json:"name"`
	URL             string   `json:"url"`
	Channel         *string  `json:"channel,omitempty"`
	AudienceSize    *int64   `json:"audienceSize,omitempty"`
	Language        *string  `json:"language,omitempty"`
	Categories      []string `json:"categories,omitempty"`
	Keywords        []string `json:"keywords,omitempty"`
	ContactEmail    *string  `json:"contactEmail,omitempty"`
	HTML            *string  `json:"html,omitempty"`
	EvidenceSnippet *string  `json:"evidenceSnippet,omitempty"`
	EvidenceSource  *string  `json:"evidenceSource,omitempty"`
}

// PostDiscoveryCorpusBody is generated from the OpenAPI component of the same name.
type PostDiscoveryCorpusBody struct {
	Keywords []string                             `json:"keywords,omitempty"`
	Entries  []PostDiscoveryCorpusBodyEntriesItem `json:"entries"`
}

// PostOutreachDraftBody is generated from the OpenAPI component of the same name.
type PostOutreachDraftBody struct {
	ProspectID  string  `json:"prospectId"`
	ProgramName *string `json:"programName,omitempty"`
}

// PutMarketplaceListingBody is generated from the OpenAPI component of the same name.
type PutMarketplaceListingBody struct {
	ProgramID   string  `json:"programId"`
	Visibility  *string `json:"visibility,omitempty"`
	Category    *string `json:"category,omitempty"`
	Headline    *string `json:"headline,omitempty"`
	Description *string `json:"description,omitempty"`
	Sponsored   *bool   `json:"sponsored,omitempty"`
}

// PostDisputesIdResolveBody is generated from the OpenAPI component of the same name.
type PostDisputesIdResolveBody struct {
	Resolution string `json:"resolution"`
}

// PutPartnerProfileBodyChannelsItem is generated from the OpenAPI component of the same name.
type PutPartnerProfileBodyChannelsItem struct {
	Type     string  `json:"type"`
	Handle   *string `json:"handle,omitempty"`
	URL      *string `json:"url,omitempty"`
	Audience *int64  `json:"audience,omitempty"`
}

// PutPartnerProfileBodyAudienceCountriesItem is generated from the OpenAPI component of the same name.
type PutPartnerProfileBodyAudienceCountriesItem struct {
	Code string  `json:"code"`
	Pct  float64 `json:"pct"`
}

// PutPartnerProfileBodyAudience is generated from the OpenAPI component of the same name.
type PutPartnerProfileBodyAudience struct {
	Countries []PutPartnerProfileBodyAudienceCountriesItem `json:"countries,omitempty"`
	Age       map[string]float64                           `json:"age,omitempty"`
	Gender    map[string]float64                           `json:"gender,omitempty"`
}

// PutPartnerProfileBody is generated from the OpenAPI component of the same name.
type PutPartnerProfileBody struct {
	Published      *bool                               `json:"published,omitempty"`
	Headline       *string                             `json:"headline,omitempty"`
	Bio            *string                             `json:"bio,omitempty"`
	Categories     []string                            `json:"categories,omitempty"`
	Languages      []string                            `json:"languages,omitempty"`
	Country        *string                             `json:"country,omitempty"`
	Channels       []PutPartnerProfileBodyChannelsItem `json:"channels,omitempty"`
	Website        *string                             `json:"website,omitempty"`
	AcceptsInvites *bool                               `json:"acceptsInvites,omitempty"`
	Audience       *PutPartnerProfileBodyAudience      `json:"audience,omitempty"`
	EngagementBps  *int64                              `json:"engagementBps,omitempty"`
}

// PostCreatorsAccountIdInviteBody is generated from the OpenAPI component of the same name.
type PostCreatorsAccountIdInviteBody struct {
	ProgramID string  `json:"programId"`
	Message   *string `json:"message,omitempty"`
}

// PostBillingCheckoutBody is generated from the OpenAPI component of the same name.
type PostBillingCheckoutBody struct {
	Plan string `json:"plan"`
}

// PostBillingActivateBody is generated from the OpenAPI component of the same name.
type PostBillingActivateBody struct {
	Plan string `json:"plan"`
}

// PutProgramsIdRefereeRewardBody is generated from the OpenAPI component of the same name.
type PutProgramsIdRefereeRewardBody struct {
	Enabled     *bool   `json:"enabled,omitempty"`
	Kind        *string `json:"kind,omitempty"`
	AmountMinor *int64  `json:"amountMinor,omitempty"`
	Bps         *int64  `json:"bps,omitempty"`
	CapMinor    *int64  `json:"capMinor,omitempty"`
	On          *string `json:"on,omitempty"`
	Rounding    *string `json:"rounding,omitempty"`
	Currency    *string `json:"currency,omitempty"`
}

// PostRewardsBody is generated from the OpenAPI component of the same name.
type PostRewardsBody struct {
	PartnerID          *string `json:"partnerId,omitempty"`
	PartnerExternalRef *string `json:"partnerExternalRef,omitempty"`
	Reason             string  `json:"reason"`
	AmountMinor        *int64  `json:"amountMinor,omitempty"`
	Currency           *string `json:"currency,omitempty"`
	RuleKey            *string `json:"ruleKey,omitempty"`
	Event              *string `json:"event,omitempty"`
	ProgramID          *string `json:"programId,omitempty"`
	IdempotencyKey     string  `json:"idempotencyKey"`
}

// PostRewardsRedeemByCodeBody is generated from the OpenAPI component of the same name.
type PostRewardsRedeemByCodeBody struct {
	AmountMinor    *int64  `json:"amountMinor,omitempty"`
	BillingEventID *string `json:"billingEventId,omitempty"`
	InvoiceRef     *string `json:"invoiceRef,omitempty"`
	ExternalID     *string `json:"externalId,omitempty"`
	OccurredAt     any     `json:"occurredAt,omitempty"`
	Code           string  `json:"code"`
}

// PostRewardsIdRedeemBody is generated from the OpenAPI component of the same name.
type PostRewardsIdRedeemBody struct {
	AmountMinor    *int64  `json:"amountMinor,omitempty"`
	BillingEventID *string `json:"billingEventId,omitempty"`
	InvoiceRef     *string `json:"invoiceRef,omitempty"`
	ExternalID     *string `json:"externalId,omitempty"`
	OccurredAt     any     `json:"occurredAt,omitempty"`
}

// PutRewardsRulesKeyBody is generated from the OpenAPI component of the same name.
type PutRewardsRulesKeyBody struct {
	Event       *string `json:"event,omitempty"` // An inbound event this rule rewards, e.g. community.answer_accepted. Unique per org.
	ProgramID   *string `json:"programId,omitempty"`
	AmountMinor int64   `json:"amountMinor"`
	Currency    *string `json:"currency,omitempty"`
	Reason      string  `json:"reason"`
	Enabled     *bool   `json:"enabled,omitempty"`
}

// PostProgramsIdGroupsBodyBaseOption1Option1 is generated from the OpenAPI component of the same name.
type PostProgramsIdGroupsBodyBaseOption1Option1 struct {
	Kind string `json:"kind"`
	Bps  int64  `json:"bps"`
}

// PostProgramsIdGroupsBodyBaseOption1Option2 is generated from the OpenAPI component of the same name.
type PostProgramsIdGroupsBodyBaseOption1Option2 struct {
	Kind        string `json:"kind"`
	AmountMinor int64  `json:"amountMinor"`
	Currency    string `json:"currency"`
}

// PostProgramsIdGroupsBody is generated from the OpenAPI component of the same name.
type PostProgramsIdGroupsBody struct {
	Name                string          `json:"name"`
	Base                json.RawMessage `json:"base,omitempty"`
	HoldDays            *int64          `json:"holdDays,omitempty"`
	MinThresholdMinor   *int64          `json:"minThresholdMinor,omitempty"`
	PayoutCurrency      *string         `json:"payoutCurrency,omitempty"`
	OnboardingMessage   *string         `json:"onboardingMessage,omitempty"`
	OnboardingChecklist []string        `json:"onboardingChecklist,omitempty"`
}

// PatchPartnerGroupsIdBodyBaseOption1Option1 is generated from the OpenAPI component of the same name.
type PatchPartnerGroupsIdBodyBaseOption1Option1 struct {
	Kind string `json:"kind"`
	Bps  int64  `json:"bps"`
}

// PatchPartnerGroupsIdBodyBaseOption1Option2 is generated from the OpenAPI component of the same name.
type PatchPartnerGroupsIdBodyBaseOption1Option2 struct {
	Kind        string `json:"kind"`
	AmountMinor int64  `json:"amountMinor"`
	Currency    string `json:"currency"`
}

// PatchPartnerGroupsIdBody is generated from the OpenAPI component of the same name.
type PatchPartnerGroupsIdBody struct {
	Name                *string         `json:"name,omitempty"`
	Base                json.RawMessage `json:"base,omitempty"`
	HoldDays            *int64          `json:"holdDays,omitempty"`
	MinThresholdMinor   *int64          `json:"minThresholdMinor,omitempty"`
	PayoutCurrency      *string         `json:"payoutCurrency,omitempty"`
	OnboardingMessage   *string         `json:"onboardingMessage,omitempty"`
	OnboardingChecklist []string        `json:"onboardingChecklist,omitempty"`
	Archived            *bool           `json:"archived,omitempty"`
}

// PutProgramsIdPartnersPartnerIdGroupBody is generated from the OpenAPI component of the same name.
type PutProgramsIdPartnersPartnerIdGroupBody struct {
	GroupID *string `json:"groupId"`
}

// PostProgramsIdPromotionsBodyEffectOption1BaseOption1 is generated from the OpenAPI component of the same name.
type PostProgramsIdPromotionsBodyEffectOption1BaseOption1 struct {
	Kind string `json:"kind"`
	Bps  int64  `json:"bps"`
}

// PostProgramsIdPromotionsBodyEffectOption1BaseOption2 is generated from the OpenAPI component of the same name.
type PostProgramsIdPromotionsBodyEffectOption1BaseOption2 struct {
	Kind        string `json:"kind"`
	AmountMinor int64  `json:"amountMinor"`
	Currency    string `json:"currency"`
}

// PostProgramsIdPromotionsBodyEffectOption1 is generated from the OpenAPI component of the same name.
type PostProgramsIdPromotionsBodyEffectOption1 struct {
	Kind string          `json:"kind"`
	Base json.RawMessage `json:"base"`
}

// PostProgramsIdPromotionsBodyEffectOption2 is generated from the OpenAPI component of the same name.
type PostProgramsIdPromotionsBodyEffectOption2 struct {
	Kind string `json:"kind"`
	Bps  int64  `json:"bps"`
}

// PostProgramsIdPromotionsBody is generated from the OpenAPI component of the same name.
type PostProgramsIdPromotionsBody struct {
	Name     string          `json:"name"`
	GroupID  *string         `json:"groupId,omitempty"`
	Effect   json.RawMessage `json:"effect"`
	StartsAt string          `json:"startsAt"`
	EndsAt   string          `json:"endsAt"`
}

// PostContractsBodyCompensationOption1 is generated from the OpenAPI component of the same name.
type PostContractsBodyCompensationOption1 struct {
	Kind string `json:"kind"`
	Bps  int64  `json:"bps"`
}

// PostContractsBodyCompensationOption2 is generated from the OpenAPI component of the same name.
type PostContractsBodyCompensationOption2 struct {
	Kind        string `json:"kind"`
	AmountMinor int64  `json:"amountMinor"`
	Currency    string `json:"currency"`
}

// PostContractsBody is generated from the OpenAPI component of the same name.
type PostContractsBody struct {
	ProgramID    string          `json:"programId"`
	PartnerID    string          `json:"partnerId"`
	Title        *string         `json:"title,omitempty"`
	TermsMd      *string         `json:"termsMd,omitempty"`
	Compensation json.RawMessage `json:"compensation"`
	Send         *bool           `json:"send,omitempty"`
}

// PutProgramsIdIncrementalityBody is generated from the OpenAPI component of the same name.
type PutProgramsIdIncrementalityBody struct {
	HoldoutBps int64 `json:"holdoutBps"`
}

// PutProgramsIdAutopricingBody is generated from the OpenAPI component of the same name.
type PutProgramsIdAutopricingBody struct {
	Enabled    *bool    `json:"enabled,omitempty"`
	TargetRoas *float64 `json:"targetRoas,omitempty"`
	MinBps     *int64   `json:"minBps,omitempty"`
	MaxBps     *int64   `json:"maxBps,omitempty"`
	Objective  *string  `json:"objective,omitempty"`
	TargetRoi  *float64 `json:"targetRoi,omitempty"`
}

// PostProgramsIdAutopriceRunBody is generated from the OpenAPI component of the same name.
type PostProgramsIdAutopriceRunBody struct {
	Objective *string `json:"objective,omitempty"`
}

// PostPassportVerifyBody is generated from the OpenAPI component of the same name.
type PostPassportVerifyBody struct {
	Passport any `json:"passport"`
}

// PostPartnerAdvanceAcceptBody is generated from the OpenAPI component of the same name.
type PostPartnerAdvanceAcceptBody struct {
	AmountMinor int64 `json:"amountMinor"`
}

// PutNetworkSettingsBody is generated from the OpenAPI component of the same name.
type PutNetworkSettingsBody struct {
	Enabled bool `json:"enabled"`
}

// PutProgramsIdNegotiationPolicyBody is generated from the OpenAPI component of the same name.
type PutProgramsIdNegotiationPolicyBody struct {
	Enabled          *bool `json:"enabled,omitempty"`
	MinBps           int64 `json:"minBps"`
	MaxBps           int64 `json:"maxBps"`
	AutoAcceptMaxBps int64 `json:"autoAcceptMaxBps"`
}

// PostAdminOrgsIdSuspendBody is generated from the OpenAPI component of the same name.
type PostAdminOrgsIdSuspendBody struct {
	Reason *string `json:"reason,omitempty"`
}

// PostAdminOrgsIdPlanBody is generated from the OpenAPI component of the same name.
type PostAdminOrgsIdPlanBody struct {
	Plan string `json:"plan"`
}

// PostAssetsBody is generated from the OpenAPI component of the same name.
type PostAssetsBody struct {
	ProgramID  string  `json:"programId"`
	Kind       string  `json:"kind"`
	Title      string  `json:"title"`
	Content    *string `json:"content,omitempty"`
	Dimensions *string `json:"dimensions,omitempty"`
	GroupID    *string `json:"groupId,omitempty"`
}

// PostPartnerDealsBody is generated from the OpenAPI component of the same name.
type PostPartnerDealsBody struct {
	OrgID        string  `json:"orgId"`
	ProgramID    string  `json:"programId"`
	Company      string  `json:"company"`
	ContactEmail string  `json:"contactEmail"`
	ValueMinor   *int64  `json:"valueMinor,omitempty"`
	Currency     *string `json:"currency,omitempty"`
	Notes        *string `json:"notes,omitempty"`
}

// PostDealsIdApproveBody is generated from the OpenAPI component of the same name.
type PostDealsIdApproveBody struct {
	ProtectionDays *int64 `json:"protectionDays,omitempty"`
}

// PostDealsIdWonBody is generated from the OpenAPI component of the same name.
type PostDealsIdWonBody struct {
	CustomerRef string `json:"customerRef"`
}

// PutPartnerOrgsOrgIdPostbackBody is generated from the OpenAPI component of the same name.
type PutPartnerOrgsOrgIdPostbackBody struct {
	URL string `json:"url"`
}

// PostPrivacyForgetBody is generated from the OpenAPI component of the same name.
type PostPrivacyForgetBody struct {
	Email       *string `json:"email,omitempty"`
	CustomerRef *string `json:"customerRef,omitempty"`
	PartnerID   *string `json:"partnerId,omitempty"`
}

// PostPrivacyLookupBody is generated from the OpenAPI component of the same name.
type PostPrivacyLookupBody struct {
	Email string `json:"email"`
}

// PostPrivacyRequestsBodySubject is generated from the OpenAPI component of the same name.
type PostPrivacyRequestsBodySubject struct {
	Email       *string `json:"email,omitempty"`
	CustomerRef *string `json:"customerRef,omitempty"`
	PartnerID   *string `json:"partnerId,omitempty"`
}

// PostPrivacyRequestsBody is generated from the OpenAPI component of the same name.
type PostPrivacyRequestsBody struct {
	Subject PostPrivacyRequestsBodySubject `json:"subject"`
	Type    string                         `json:"type"`
}

// PutEmailSettingsBody is generated from the OpenAPI component of the same name.
type PutEmailSettingsBody struct {
	Provider      *string `json:"provider,omitempty"`
	APIKey        *string `json:"apiKey,omitempty"`
	FromAddress   *string `json:"fromAddress,omitempty"`
	FromName      *string `json:"fromName,omitempty"`
	ReplyTo       *string `json:"replyTo,omitempty"`
	DefaultLocale *string `json:"defaultLocale,omitempty"`
	LogoURL       *string `json:"logoUrl,omitempty"`
	PrimaryColor  *string `json:"primaryColor,omitempty"`
}

// PutEmailTemplatesEventLocaleBody is generated from the OpenAPI component of the same name.
type PutEmailTemplatesEventLocaleBody struct {
	Subject string  `json:"subject"`
	Text    string  `json:"text"`
	HTML    *string `json:"html,omitempty"`
}

// PostEmailPreviewBodyDraft is generated from the OpenAPI component of the same name.
type PostEmailPreviewBodyDraft struct {
	Subject string  `json:"subject"`
	Text    string  `json:"text"`
	HTML    *string `json:"html,omitempty"`
}

// PostEmailPreviewBody is generated from the OpenAPI component of the same name.
type PostEmailPreviewBody struct {
	Event  string                     `json:"event"`
	Locale string                     `json:"locale"`
	Draft  *PostEmailPreviewBodyDraft `json:"draft,omitempty"`
}

// PostEmailTestBody is generated from the OpenAPI component of the same name.
type PostEmailTestBody struct {
	Event  string `json:"event"`
	Locale string `json:"locale"`
	To     string `json:"to"`
}

// PutMeLocaleBody is generated from the OpenAPI component of the same name.
type PutMeLocaleBody struct {
	Locale string `json:"locale"`
}

// PutMeRoiDigestBody is generated from the OpenAPI component of the same name.
type PutMeRoiDigestBody struct {
	Enabled bool `json:"enabled"`
}

// PutPartnerLocaleBody is generated from the OpenAPI component of the same name.
type PutPartnerLocaleBody struct {
	Locale string `json:"locale"`
}

// PutIntegrationsHubspotBody is generated from the OpenAPI component of the same name.
type PutIntegrationsHubspotBody struct {
	Token         *string `json:"token,omitempty"`
	Pipeline      *string `json:"pipeline,omitempty"`
	NewDealStage  *string `json:"newDealStage,omitempty"`
	PaidDealStage *string `json:"paidDealStage,omitempty"`
	Status        *string `json:"status,omitempty"`
}

// PutIntegrationsSegmentBody is generated from the OpenAPI component of the same name.
type PutIntegrationsSegmentBody struct {
	WriteKey *string  `json:"writeKey,omitempty"`
	Events   []string `json:"events,omitempty"`
	Status   *string  `json:"status,omitempty"`
}

// PostIntegrationsHubspotTestBody is generated from the OpenAPI component of the same name.
type PostIntegrationsHubspotTestBody struct {
	SetupProperties *bool `json:"setupProperties,omitempty"`
}

// PostZapierHooksBody is generated from the OpenAPI component of the same name.
type PostZapierHooksBody struct {
	TargetURL string `json:"target_url"`
	Event     string `json:"event"`
}

// PostZapierConversionsBody is generated from the OpenAPI component of the same name.
type PostZapierConversionsBody struct {
	ProviderEventID string  `json:"providerEventId"`
	Type            *string `json:"type,omitempty"`
	CustomerRef     string  `json:"customerRef"`
	AmountMinor     int64   `json:"amountMinor"`
	Currency        *string `json:"currency,omitempty"`
	InvoiceRef      *string `json:"invoiceRef,omitempty"`
	OccurredAt      *string `json:"occurredAt,omitempty"`
	ConnectionID    *string `json:"connectionId,omitempty"`
}

// PostPartnerSecurityTotpEnableBody is generated from the OpenAPI component of the same name.
type PostPartnerSecurityTotpEnableBody struct {
	Code string `json:"code"`
}

// PostPartnerSecurityTotpDisableBody is generated from the OpenAPI component of the same name.
type PostPartnerSecurityTotpDisableBody struct {
	Code string `json:"code"`
}

// PostAuthConnectionsBodyKeys is generated from the OpenAPI component of the same name.
type PostAuthConnectionsBodyKeys struct {
	ReferralToken *string `json:"referralToken,omitempty"`
	AnonID        *string `json:"anonId,omitempty"`
	CouponCode    *string `json:"couponCode,omitempty"`
	ProgramID     *string `json:"programId,omitempty"`
	AutoEnroll    *string `json:"autoEnroll,omitempty"`
}

// PostAuthConnectionsBody is generated from the OpenAPI component of the same name.
type PostAuthConnectionsBody struct {
	Provider        string                       `json:"provider"`
	ProgramID       string                       `json:"programId"`
	InstanceID      *string                      `json:"instanceId,omitempty"`
	AutoEnrollUsers *bool                        `json:"autoEnrollUsers,omitempty"`
	DestinationURL  *string                      `json:"destinationUrl,omitempty"`
	Name            *string                      `json:"name,omitempty"`
	Secret          *string                      `json:"secret,omitempty"`
	Keys            *PostAuthConnectionsBodyKeys `json:"keys,omitempty"`
}

// PutAuthConnectionsIdSecretBody is generated from the OpenAPI component of the same name.
type PutAuthConnectionsIdSecretBody struct {
	Secret string `json:"secret"`
}

// PatchAuthConnectionsIdBody is generated from the OpenAPI component of the same name.
type PatchAuthConnectionsIdBody struct {
	Status           *string `json:"status,omitempty"`
	Name             *string `json:"name,omitempty"`
	AutoEnrollUsers  *bool   `json:"autoEnrollUsers,omitempty"`
	DestinationURL   *string `json:"destinationUrl,omitempty"`
	SignatureVersion *string `json:"signatureVersion,omitempty"`
}

// PostAuthConnectionsIdSignupsBodyReferral is generated from the OpenAPI component of the same name.
type PostAuthConnectionsIdSignupsBodyReferral struct {
	Code       *string `json:"code,omitempty"`
	ClickID    *string `json:"clickId,omitempty"`
	AnonID     *string `json:"anonId,omitempty"`
	CapturedAt *string `json:"capturedAt,omitempty"`
	Consent    *string `json:"consent,omitempty"`
}

// PostAuthConnectionsIdSignupsBodyRisk is generated from the OpenAPI component of the same name.
type PostAuthConnectionsIdSignupsBodyRisk struct {
	BotScore        *float64 `json:"botScore,omitempty"`
	DisposableEmail *bool    `json:"disposableEmail,omitempty"`
	EmailVerified   *bool    `json:"emailVerified,omitempty"`
	SelfReferral    *bool    `json:"selfReferral,omitempty"`
	Velocity        *int64   `json:"velocity,omitempty"`
	Reasons         []string `json:"reasons,omitempty"`
	Score           *float64 `json:"score,omitempty"`
	Level           *string  `json:"level,omitempty"`
	Dropped         *bool    `json:"dropped,omitempty"`
}

// PostAuthConnectionsIdSignupsBody is generated from the OpenAPI component of the same name.
type PostAuthConnectionsIdSignupsBody struct {
	CustomerRef  string                                    `json:"customerRef"`
	Email        *string                                   `json:"email,omitempty"`
	EmailHash    *string                                   `json:"emailHash,omitempty"`
	Name         *string                                   `json:"name,omitempty"`
	IP           *string                                   `json:"ip,omitempty"`
	UA           *string                                   `json:"ua,omitempty"`
	SignedUpAt   string                                    `json:"signedUpAt"`
	Referral     *PostAuthConnectionsIdSignupsBodyReferral `json:"referral,omitempty"`
	Risk         *PostAuthConnectionsIdSignupsBodyRisk     `json:"risk,omitempty"`
	EmailDomain  *string                                   `json:"emailDomain,omitempty"`
	ReferralCode *string                                   `json:"referralCode,omitempty"`
	AutoEnroll   *bool                                     `json:"autoEnroll,omitempty"`
	Test         *bool                                     `json:"test,omitempty"`
}

// PostAuthConnectionsIdPartnersBody is generated from the OpenAPI component of the same name.
type PostAuthConnectionsIdPartnersBody struct {
	CustomerRef string  `json:"customerRef"`
	Email       string  `json:"email"`
	Name        *string `json:"name,omitempty"`
	Code        *string `json:"code,omitempty"`
}

// PostAuthConnectionsIdEraseBody is generated from the OpenAPI component of the same name.
type PostAuthConnectionsIdEraseBody struct {
	CustomerRef string `json:"customerRef"`
}

// PostAuthConnectionsIdAliasesBody is generated from the OpenAPI component of the same name.
type PostAuthConnectionsIdAliasesBody struct {
	FromCustomerRef string `json:"fromCustomerRef"`
	ToCustomerRef   string `json:"toCustomerRef"`
}

// PostAuthConnectionsIdUsersBody is generated from the OpenAPI component of the same name.
type PostAuthConnectionsIdUsersBody struct {
	CustomerRef string  `json:"customerRef"`
	Email       *string `json:"email,omitempty"`
	Name        *string `json:"name,omitempty"`
	Status      *string `json:"status,omitempty"`
}

// PostAuthConnectionsIdSignalsBody is generated from the OpenAPI component of the same name.
type PostAuthConnectionsIdSignalsBody struct {
	CustomerRef string         `json:"customerRef"`
	Signal      string         `json:"signal"`
	At          *string        `json:"at,omitempty"`
	Details     map[string]any `json:"details,omitempty"`
}

// PostAuthConnectionsInstallBody is generated from the OpenAPI component of the same name.
type PostAuthConnectionsInstallBody struct {
	ClientID      string  `json:"clientId"`
	Scope         string  `json:"scope"`
	RedirectURI   string  `json:"redirectUri"`
	State         *string `json:"state,omitempty"`
	CodeChallenge string  `json:"codeChallenge"`
	InstanceID    string  `json:"instanceId"`
	ProgramID     *string `json:"programId,omitempty"`
	Decision      string  `json:"decision"`
}

// PostOauthTokenBody is generated from the OpenAPI component of the same name.
type PostOauthTokenBody struct {
	GrantType    string `json:"grant_type"`
	Code         string `json:"code"`
	CodeVerifier string `json:"code_verifier"`
	ClientID     string `json:"client_id"`
	RedirectURI  string `json:"redirect_uri"`
}

// PatchOrgSsoBody is generated from the OpenAPI component of the same name.
type PatchOrgSsoBody struct {
	RequireSSO  *bool             `json:"requireSso,omitempty"`
	RoleMapping map[string]string `json:"roleMapping,omitempty"`
}

// PostOrgSsoTicketsBody is generated from the OpenAPI component of the same name.
type PostOrgSsoTicketsBody struct {
	ExpiresInDays *int64 `json:"expiresInDays,omitempty"`
}

// PutOrgSsoDirectBodyOption1 is generated from the OpenAPI component of the same name.
type PutOrgSsoDirectBodyOption1 struct {
	Adapter            string            `json:"adapter"`
	Issuer             string            `json:"issuer"`
	ClientID           string            `json:"clientId"`
	ClientSecret       *string           `json:"clientSecret,omitempty"`
	Scopes             *string           `json:"scopes,omitempty"`
	RoleClaim          *string           `json:"roleClaim,omitempty"`
	Name               *string           `json:"name,omitempty"`
	Domains            []string          `json:"domains,omitempty"`
	RoleMapping        map[string]string `json:"roleMapping,omitempty"`
	DefaultRole        *string           `json:"defaultRole,omitempty"`
	TrustEmailVerified *bool             `json:"trustEmailVerified,omitempty"`
}

// PutOrgSsoDirectBodyOption2 is generated from the OpenAPI component of the same name.
type PutOrgSsoDirectBodyOption2 struct {
	Adapter            string            `json:"adapter"`
	ClientID           string            `json:"clientId"`
	APIKey             *string           `json:"apiKey,omitempty"`
	OrganizationID     string            `json:"organizationId"`
	Domains            []string          `json:"domains,omitempty"`
	RoleMapping        map[string]string `json:"roleMapping,omitempty"`
	DefaultRole        *string           `json:"defaultRole,omitempty"`
	TrustEmailVerified *bool             `json:"trustEmailVerified,omitempty"`
}

// PostPartnerAuthWalletNonceBody is generated from the OpenAPI component of the same name.
type PostPartnerAuthWalletNonceBody struct {
	Address string `json:"address"`
}

// PostPartnerAuthWalletVerifyBody is generated from the OpenAPI component of the same name.
type PostPartnerAuthWalletVerifyBody struct {
	Address   string  `json:"address"`
	Nonce     string  `json:"nonce"`
	Signature string  `json:"signature"`
	Name      *string `json:"name,omitempty"`
}

// PostOnchainConnectionsBody is generated from the OpenAPI component of the same name.
type PostOnchainConnectionsBody struct {
	ProgramID string           `json:"programId"`
	Name      *string          `json:"name,omitempty"`
	Secret    *string          `json:"secret,omitempty"`
	Decimals  map[string]int64 `json:"decimals,omitempty"`
}

// PostPublicProgramsHandleJoinBody is generated from the OpenAPI component of the same name.
type PostPublicProgramsHandleJoinBody struct {
	DestinationURL *string `json:"destinationUrl,omitempty"`
}

// PutSupportAppsAdapterBodyConfig is generated from the OpenAPI component of the same name.
type PutSupportAppsAdapterBodyConfig struct {
	PublicKeyPem *string `json:"publicKeyPem,omitempty"`
	Subdomain    *string `json:"subdomain,omitempty"`
	CardKey      *string `json:"cardKey,omitempty"`
}

// PutSupportAppsAdapterBody is generated from the OpenAPI component of the same name.
type PutSupportAppsAdapterBody struct {
	Enabled *bool                            `json:"enabled,omitempty"`
	Secret  *string                          `json:"secret,omitempty"` // The vendor-issued signing secret (Chatly, Intercom, HubSpot, Front, Help Scout). Write-only.
	Rotate  *bool                            `json:"rotate,omitempty"` // custom: mint a new secret (shown once).
	Config  *PutSupportAppsAdapterBodyConfig `json:"config,omitempty"`
}

// PostSupportAppsAdapterTestBody is generated from the OpenAPI component of the same name.
type PostSupportAppsAdapterTestBody struct {
	Email              *string `json:"email,omitempty"` // Used only when no stronger key is given.
	CustomerRef        *string `json:"customerRef,omitempty"`
	PartnerExternalRef *string `json:"partnerExternalRef,omitempty"`
	Ext                *string `json:"ext,omitempty"` // A support tool's verified external id: matched as a partner externalRef, then as a customer ref.
}

// PostPrivacyExecutorKeysBodyConfig is generated from the OpenAPI component of the same name.
type PostPrivacyExecutorKeysBodyConfig struct {
	PublicKeyPem *string `json:"publicKeyPem,omitempty"`
	Algorithm    *string `json:"algorithm,omitempty"`
	SombraURL    *string `json:"sombraUrl,omitempty"`
}

// PostPrivacyExecutorKeysBody is generated from the OpenAPI component of the same name.
type PostPrivacyExecutorKeysBody struct {
	Adapter      string                             `json:"adapter"`
	VendorSecret *string                            `json:"vendorSecret,omitempty"` // A credential from the vendor (Transcend API key, Osano webhook secret). Write-only.
	Config       *PostPrivacyExecutorKeysBodyConfig `json:"config,omitempty"`
	Rotate       *bool                              `json:"rotate,omitempty"` // Mint a new executor key (the old one stops working). false keeps the current key and only updates the settings.
	Enabled      *bool                              `json:"enabled,omitempty"`
}

// PostAccountCloseBody is generated from the OpenAPI component of the same name.
type PostAccountCloseBody struct {
	Confirm string `json:"confirm"` // The account's email address, typed out, to confirm.
}

// PutOrgTrackingConsentBody is generated from the OpenAPI component of the same name.
type PutOrgTrackingConsentBody struct {
	Mode        string  `json:"mode"`               // `auto` detects the CMP, `none` always stores, `pending` waits for `aff.consent()`, or an adapter id.
	Category    *string `json:"category,omitempty"` // The CMP category, group, purpose or service to require (each adapter has a default).
	TcfVendorID *int64  `json:"tcfVendorId,omitempty"`
	GcmSignal   *string `json:"gcmSignal,omitempty"`
	NoCmp       *string `json:"noCmp,omitempty"` // `auto` with no CMP on the page: `persist` stores, `memory` never does.
}

// PutOrgHostedPagesBodyConsent is generated from the OpenAPI component of the same name.
type PutOrgHostedPagesBodyConsent struct {
	Adapter string         `json:"adapter"`
	Config  map[string]any `json:"config,omitempty"`
}

// PutOrgHostedPagesBodyChatOption1 is generated from the OpenAPI component of the same name.
type PutOrgHostedPagesBodyChatOption1 struct {
	Adapter string         `json:"adapter"`
	Config  map[string]any `json:"config,omitempty"`
}

// PutOrgHostedPagesBody is generated from the OpenAPI component of the same name.
type PutOrgHostedPagesBody struct {
	Consent       PutOrgHostedPagesBodyConsent      `json:"consent"`
	Chat          *PutOrgHostedPagesBodyChatOption1 `json:"chat,omitempty"`
	CustomOrigins map[string][]string               `json:"customOrigins,omitempty"`
}

// PostGraphqlBody is generated from the OpenAPI component of the same name.
type PostGraphqlBody struct {
	Query         string         `json:"query"`
	Variables     map[string]any `json:"variables,omitempty"`
	OperationName *string        `json:"operationName,omitempty"`
}
