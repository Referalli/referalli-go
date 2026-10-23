package referalli

// Code generated from the Referalli OpenAPI document. DO NOT EDIT.
// Changes belong in the generator, which overwrites this file.

// ValidationError is the `validation_error` error (HTTP 400). Match it with errors.As.
type ValidationError struct{ *APIError }

func (e *ValidationError) Unwrap() error { return e.APIError }

// ModeMismatchError is the `mode_mismatch` error (HTTP 400). Match it with errors.As.
type ModeMismatchError struct{ *APIError }

func (e *ModeMismatchError) Unwrap() error { return e.APIError }

// BadRequestError is the `bad_request` error (HTTP 400). Match it with errors.As.
type BadRequestError struct{ *APIError }

func (e *BadRequestError) Unwrap() error { return e.APIError }

// AuthenticationError is the `unauthorized` error (HTTP 401). Match it with errors.As.
type AuthenticationError struct{ *APIError }

func (e *AuthenticationError) Unwrap() error { return e.APIError }

// StepUpRequiredError is the `step_up_required` error (HTTP 401). Match it with errors.As.
type StepUpRequiredError struct{ *APIError }

func (e *StepUpRequiredError) Unwrap() error { return e.APIError }

// PaymentRequiredError is the `payment_required` error (HTTP 402). Match it with errors.As.
type PaymentRequiredError struct{ *APIError }

func (e *PaymentRequiredError) Unwrap() error { return e.APIError }

// PermissionDeniedError is the `forbidden` error (HTTP 403). Match it with errors.As.
type PermissionDeniedError struct{ *APIError }

func (e *PermissionDeniedError) Unwrap() error { return e.APIError }

// WebhooksDisabledError is the `webhooks_disabled` error (HTTP 403). Match it with errors.As.
type WebhooksDisabledError struct{ *APIError }

func (e *WebhooksDisabledError) Unwrap() error { return e.APIError }

// SsoRequiredError is the `sso_required` error (HTTP 403). Match it with errors.As.
type SsoRequiredError struct{ *APIError }

func (e *SsoRequiredError) Unwrap() error { return e.APIError }

// AccessSuspendedError is the `access_suspended` error (HTTP 403). Match it with errors.As.
type AccessSuspendedError struct{ *APIError }

func (e *AccessSuspendedError) Unwrap() error { return e.APIError }

// NotFoundError is the `not_found` error (HTTP 404). Match it with errors.As.
type NotFoundError struct{ *APIError }

func (e *NotFoundError) Unwrap() error { return e.APIError }

// ConflictError is the `conflict` error (HTTP 409). Match it with errors.As.
type ConflictError struct{ *APIError }

func (e *ConflictError) Unwrap() error { return e.APIError }

// IdempotencyInProgressError is the `idempotency_in_progress` error (HTTP 409). Match it with errors.As.
type IdempotencyInProgressError struct{ *APIError }

func (e *IdempotencyInProgressError) Unwrap() error { return e.APIError }

// ReplayError is the `replay` error (HTTP 409). Match it with errors.As.
type ReplayError struct{ *APIError }

func (e *ReplayError) Unwrap() error { return e.APIError }

// LiveModeOnlyError is the `live_mode_only` error (HTTP 409). Match it with errors.As.
type LiveModeOnlyError struct{ *APIError }

func (e *LiveModeOnlyError) Unwrap() error { return e.APIError }

// TestModeOnlyError is the `test_mode_only` error (HTTP 409). Match it with errors.As.
type TestModeOnlyError struct{ *APIError }

func (e *TestModeOnlyError) Unwrap() error { return e.APIError }

// TestModeUnsupportedError is the `test_mode_unsupported` error (HTTP 409). Match it with errors.As.
type TestModeUnsupportedError struct{ *APIError }

func (e *TestModeUnsupportedError) Unwrap() error { return e.APIError }

// GoneError is the `gone` error (HTTP 410). Match it with errors.As.
type GoneError struct{ *APIError }

func (e *GoneError) Unwrap() error { return e.APIError }

// IdempotencyKeyReusedError is the `idempotency_key_reused` error (HTTP 422). Match it with errors.As.
type IdempotencyKeyReusedError struct{ *APIError }

func (e *IdempotencyKeyReusedError) Unwrap() error { return e.APIError }

// PayloadTooLargeError is the `payload_too_large` error (HTTP 413). Match it with errors.As.
type PayloadTooLargeError struct{ *APIError }

func (e *PayloadTooLargeError) Unwrap() error { return e.APIError }

// ConnectionPausedError is the `connection_paused` error (HTTP 423). Match it with errors.As.
type ConnectionPausedError struct{ *APIError }

func (e *ConnectionPausedError) Unwrap() error { return e.APIError }

// RateLimitedError is the `rate_limited` error (HTTP 429). Match it with errors.As.
type RateLimitedError struct{ *APIError }

func (e *RateLimitedError) Unwrap() error { return e.APIError }

// InternalServerError is the `internal_error` error (HTTP 500). Match it with errors.As.
type InternalServerError struct{ *APIError }

func (e *InternalServerError) Unwrap() error { return e.APIError }

// AtlasError is the `atlas_error` error (HTTP 502). Match it with errors.As.
type AtlasError struct{ *APIError }

func (e *AtlasError) Unwrap() error { return e.APIError }

// SsoProviderError is the `sso_provider_error` error (HTTP 502). Match it with errors.As.
type SsoProviderError struct{ *APIError }

func (e *SsoProviderError) Unwrap() error { return e.APIError }

// SsoNotConfiguredError is the `sso_not_configured` error (HTTP 503). Match it with errors.As.
type SsoNotConfiguredError struct{ *APIError }

func (e *SsoNotConfiguredError) Unwrap() error { return e.APIError }

func typedError(base *APIError) error {
	switch base.Code {
	case "validation_error":
		return &ValidationError{base}
	case "mode_mismatch":
		return &ModeMismatchError{base}
	case "bad_request":
		return &BadRequestError{base}
	case "unauthorized":
		return &AuthenticationError{base}
	case "step_up_required":
		return &StepUpRequiredError{base}
	case "payment_required":
		return &PaymentRequiredError{base}
	case "forbidden":
		return &PermissionDeniedError{base}
	case "webhooks_disabled":
		return &WebhooksDisabledError{base}
	case "sso_required":
		return &SsoRequiredError{base}
	case "access_suspended":
		return &AccessSuspendedError{base}
	case "not_found":
		return &NotFoundError{base}
	case "conflict":
		return &ConflictError{base}
	case "idempotency_in_progress":
		return &IdempotencyInProgressError{base}
	case "replay":
		return &ReplayError{base}
	case "live_mode_only":
		return &LiveModeOnlyError{base}
	case "test_mode_only":
		return &TestModeOnlyError{base}
	case "test_mode_unsupported":
		return &TestModeUnsupportedError{base}
	case "gone":
		return &GoneError{base}
	case "idempotency_key_reused":
		return &IdempotencyKeyReusedError{base}
	case "payload_too_large":
		return &PayloadTooLargeError{base}
	case "connection_paused":
		return &ConnectionPausedError{base}
	case "rate_limited":
		return &RateLimitedError{base}
	case "internal_error":
		return &InternalServerError{base}
	case "atlas_error":
		return &AtlasError{base}
	case "sso_provider_error":
		return &SsoProviderError{base}
	case "sso_not_configured":
		return &SsoNotConfiguredError{base}
	}
	return base
}
