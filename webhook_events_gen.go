package referalli

// Code generated from the Referalli OpenAPI document. DO NOT EDIT.
// Changes belong in the generator, which overwrites this file.

import "encoding/json"

// Every webhook topic the API publishes.
const (
	TopicBillingEventRecorded    = "billing_event.recorded"
	TopicClickRecorded           = "click.recorded"
	TopicCommissionEarned        = "commission.earned"
	TopicCommissionReversed      = "commission.reversed"
	TopicCommissionSettled       = "commission.settled"
	TopicCreditRedeemed          = "credit.redeemed"
	TopicCustomerAliased         = "customer.aliased"
	TopicDealSubmitted           = "deal.submitted"
	TopicDisputeResolved         = "dispute.resolved"
	TopicFraudCaseOpened         = "fraud.case_opened"
	TopicPartnerApplied          = "partner.applied"
	TopicPartnerApproved         = "partner.approved"
	TopicPartnerEnrolled         = "partner.enrolled"
	TopicPartnerLevelUp          = "partner.level_up"
	TopicPayoutCancelled         = "payout.cancelled"
	TopicPayoutFailed            = "payout.failed"
	TopicPayoutScheduled         = "payout.scheduled"
	TopicPayoutSent              = "payout.sent"
	TopicPayoutSubmitted         = "payout.submitted"
	TopicPrivacyRequestCompleted = "privacy.request_completed"
	TopicRefereeRewardEarned     = "referee_reward.earned"
	TopicReferralApproved        = "referral.approved"
	TopicReferralCreated         = "referral.created"
	TopicReferralHeld            = "referral.held"
	TopicReferralPaused          = "referral.paused"
	TopicReferralReassigned      = "referral.reassigned"
	TopicReferralRejected        = "referral.rejected"
	TopicReferralResumed         = "referral.resumed"
	TopicReferralVoided          = "referral.voided"
	TopicRoiPatternDetected      = "roi.pattern_detected"
	TopicRoiPaybackReached       = "roi.payback_reached"
	TopicTaxFormDelivered        = "tax.form_delivered"
	TopicWebhookEndpointDisabled = "webhook_endpoint.disabled"
)

// Topics lists every webhook topic.
var Topics = []string{
	TopicBillingEventRecorded,
	TopicClickRecorded,
	TopicCommissionEarned,
	TopicCommissionReversed,
	TopicCommissionSettled,
	TopicCreditRedeemed,
	TopicCustomerAliased,
	TopicDealSubmitted,
	TopicDisputeResolved,
	TopicFraudCaseOpened,
	TopicPartnerApplied,
	TopicPartnerApproved,
	TopicPartnerEnrolled,
	TopicPartnerLevelUp,
	TopicPayoutCancelled,
	TopicPayoutFailed,
	TopicPayoutScheduled,
	TopicPayoutSent,
	TopicPayoutSubmitted,
	TopicPrivacyRequestCompleted,
	TopicRefereeRewardEarned,
	TopicReferralApproved,
	TopicReferralCreated,
	TopicReferralHeld,
	TopicReferralPaused,
	TopicReferralReassigned,
	TopicReferralRejected,
	TopicReferralResumed,
	TopicReferralVoided,
	TopicRoiPatternDetected,
	TopicRoiPaybackReached,
	TopicTaxFormDelivered,
	TopicWebhookEndpointDisabled,
}

// AsBillingEventRecorded decodes the data of a `billing_event.recorded` event (A billing event was recorded).
// It returns a *WrongTopicError for an event of another topic.
func (e *WebhookEvent) AsBillingEventRecorded() (BillingEventRecordedEventData, error) {
	var out BillingEventRecordedEventData
	if e.Topic != TopicBillingEventRecorded {
		return out, &WrongTopicError{Want: TopicBillingEventRecorded, Got: e.Topic}
	}
	err := json.Unmarshal(e.Data, &out)
	return out, err
}

// AsClickRecorded decodes the data of a `click.recorded` event (A tracked click was recorded).
// It returns a *WrongTopicError for an event of another topic.
func (e *WebhookEvent) AsClickRecorded() (ClickRecordedEventData, error) {
	var out ClickRecordedEventData
	if e.Topic != TopicClickRecorded {
		return out, &WrongTopicError{Want: TopicClickRecorded, Got: e.Topic}
	}
	err := json.Unmarshal(e.Data, &out)
	return out, err
}

// AsCommissionEarned decodes the data of a `commission.earned` event (A commission line was earned).
// It returns a *WrongTopicError for an event of another topic.
func (e *WebhookEvent) AsCommissionEarned() (CommissionEarnedEventData, error) {
	var out CommissionEarnedEventData
	if e.Topic != TopicCommissionEarned {
		return out, &WrongTopicError{Want: TopicCommissionEarned, Got: e.Topic}
	}
	err := json.Unmarshal(e.Data, &out)
	return out, err
}

// AsCommissionReversed decodes the data of a `commission.reversed` event (A commission line was reversed).
// It returns a *WrongTopicError for an event of another topic.
func (e *WebhookEvent) AsCommissionReversed() (CommissionReversedEventData, error) {
	var out CommissionReversedEventData
	if e.Topic != TopicCommissionReversed {
		return out, &WrongTopicError{Want: TopicCommissionReversed, Got: e.Topic}
	}
	err := json.Unmarshal(e.Data, &out)
	return out, err
}

// AsCommissionSettled decodes the data of a `commission.settled` event (A commission line was settled by the merchant).
// It returns a *WrongTopicError for an event of another topic.
func (e *WebhookEvent) AsCommissionSettled() (CommissionSettledEventData, error) {
	var out CommissionSettledEventData
	if e.Topic != TopicCommissionSettled {
		return out, &WrongTopicError{Want: TopicCommissionSettled, Got: e.Topic}
	}
	err := json.Unmarshal(e.Data, &out)
	return out, err
}

// AsCreditRedeemed decodes the data of a `credit.redeemed` event (A referee reward was redeemed).
// It returns a *WrongTopicError for an event of another topic.
func (e *WebhookEvent) AsCreditRedeemed() (CreditRedeemedEventData, error) {
	var out CreditRedeemedEventData
	if e.Topic != TopicCreditRedeemed {
		return out, &WrongTopicError{Want: TopicCreditRedeemed, Got: e.Topic}
	}
	err := json.Unmarshal(e.Data, &out)
	return out, err
}

// AsCustomerAliased decodes the data of a `customer.aliased` event (A customer's id changed).
// It returns a *WrongTopicError for an event of another topic.
func (e *WebhookEvent) AsCustomerAliased() (CustomerAliasedEventData, error) {
	var out CustomerAliasedEventData
	if e.Topic != TopicCustomerAliased {
		return out, &WrongTopicError{Want: TopicCustomerAliased, Got: e.Topic}
	}
	err := json.Unmarshal(e.Data, &out)
	return out, err
}

// AsDealSubmitted decodes the data of a `deal.submitted` event (A partner registered a deal).
// It returns a *WrongTopicError for an event of another topic.
func (e *WebhookEvent) AsDealSubmitted() (DealSubmittedEventData, error) {
	var out DealSubmittedEventData
	if e.Topic != TopicDealSubmitted {
		return out, &WrongTopicError{Want: TopicDealSubmitted, Got: e.Topic}
	}
	err := json.Unmarshal(e.Data, &out)
	return out, err
}

// AsDisputeResolved decodes the data of a `dispute.resolved` event (A partner dispute was resolved).
// It returns a *WrongTopicError for an event of another topic.
func (e *WebhookEvent) AsDisputeResolved() (DisputeResolvedEventData, error) {
	var out DisputeResolvedEventData
	if e.Topic != TopicDisputeResolved {
		return out, &WrongTopicError{Want: TopicDisputeResolved, Got: e.Topic}
	}
	err := json.Unmarshal(e.Data, &out)
	return out, err
}

// AsFraudCaseOpened decodes the data of a `fraud.case_opened` event (A fraud case was opened).
// It returns a *WrongTopicError for an event of another topic.
func (e *WebhookEvent) AsFraudCaseOpened() (FraudCaseOpenedEventData, error) {
	var out FraudCaseOpenedEventData
	if e.Topic != TopicFraudCaseOpened {
		return out, &WrongTopicError{Want: TopicFraudCaseOpened, Got: e.Topic}
	}
	err := json.Unmarshal(e.Data, &out)
	return out, err
}

// AsPartnerApplied decodes the data of a `partner.applied` event (A partner applied to a program).
// It returns a *WrongTopicError for an event of another topic.
func (e *WebhookEvent) AsPartnerApplied() (PartnerAppliedEventData, error) {
	var out PartnerAppliedEventData
	if e.Topic != TopicPartnerApplied {
		return out, &WrongTopicError{Want: TopicPartnerApplied, Got: e.Topic}
	}
	err := json.Unmarshal(e.Data, &out)
	return out, err
}

// AsPartnerApproved decodes the data of a `partner.approved` event (A partner was approved).
// It returns a *WrongTopicError for an event of another topic.
func (e *WebhookEvent) AsPartnerApproved() (PartnerApprovedEventData, error) {
	var out PartnerApprovedEventData
	if e.Topic != TopicPartnerApproved {
		return out, &WrongTopicError{Want: TopicPartnerApproved, Got: e.Topic}
	}
	err := json.Unmarshal(e.Data, &out)
	return out, err
}

// AsPartnerEnrolled decodes the data of a `partner.enrolled` event (A partner was enrolled).
// It returns a *WrongTopicError for an event of another topic.
func (e *WebhookEvent) AsPartnerEnrolled() (PartnerEnrolledEventData, error) {
	var out PartnerEnrolledEventData
	if e.Topic != TopicPartnerEnrolled {
		return out, &WrongTopicError{Want: TopicPartnerEnrolled, Got: e.Topic}
	}
	err := json.Unmarshal(e.Data, &out)
	return out, err
}

// AsPartnerLevelUp decodes the data of a `partner.level_up` event (A partner reached a new level).
// It returns a *WrongTopicError for an event of another topic.
func (e *WebhookEvent) AsPartnerLevelUp() (PartnerLevelUpEventData, error) {
	var out PartnerLevelUpEventData
	if e.Topic != TopicPartnerLevelUp {
		return out, &WrongTopicError{Want: TopicPartnerLevelUp, Got: e.Topic}
	}
	err := json.Unmarshal(e.Data, &out)
	return out, err
}

// AsPayoutCancelled decodes the data of a `payout.cancelled` event (A payout was cancelled).
// It returns a *WrongTopicError for an event of another topic.
func (e *WebhookEvent) AsPayoutCancelled() (PayoutCancelledEventData, error) {
	var out PayoutCancelledEventData
	if e.Topic != TopicPayoutCancelled {
		return out, &WrongTopicError{Want: TopicPayoutCancelled, Got: e.Topic}
	}
	err := json.Unmarshal(e.Data, &out)
	return out, err
}

// AsPayoutFailed decodes the data of a `payout.failed` event (A payout failed).
// It returns a *WrongTopicError for an event of another topic.
func (e *WebhookEvent) AsPayoutFailed() (PayoutFailedEventData, error) {
	var out PayoutFailedEventData
	if e.Topic != TopicPayoutFailed {
		return out, &WrongTopicError{Want: TopicPayoutFailed, Got: e.Topic}
	}
	err := json.Unmarshal(e.Data, &out)
	return out, err
}

// AsPayoutScheduled decodes the data of a `payout.scheduled` event (A payout was created in a run).
// It returns a *WrongTopicError for an event of another topic.
func (e *WebhookEvent) AsPayoutScheduled() (PayoutScheduledEventData, error) {
	var out PayoutScheduledEventData
	if e.Topic != TopicPayoutScheduled {
		return out, &WrongTopicError{Want: TopicPayoutScheduled, Got: e.Topic}
	}
	err := json.Unmarshal(e.Data, &out)
	return out, err
}

// AsPayoutSent decodes the data of a `payout.sent` event (A payout was paid).
// It returns a *WrongTopicError for an event of another topic.
func (e *WebhookEvent) AsPayoutSent() (PayoutSentEventData, error) {
	var out PayoutSentEventData
	if e.Topic != TopicPayoutSent {
		return out, &WrongTopicError{Want: TopicPayoutSent, Got: e.Topic}
	}
	err := json.Unmarshal(e.Data, &out)
	return out, err
}

// AsPayoutSubmitted decodes the data of a `payout.submitted` event (A payout was submitted to the provider).
// It returns a *WrongTopicError for an event of another topic.
func (e *WebhookEvent) AsPayoutSubmitted() (PayoutSubmittedEventData, error) {
	var out PayoutSubmittedEventData
	if e.Topic != TopicPayoutSubmitted {
		return out, &WrongTopicError{Want: TopicPayoutSubmitted, Got: e.Topic}
	}
	err := json.Unmarshal(e.Data, &out)
	return out, err
}

// AsPrivacyRequestCompleted decodes the data of a `privacy.request_completed` event (A data-subject request finished).
// It returns a *WrongTopicError for an event of another topic.
func (e *WebhookEvent) AsPrivacyRequestCompleted() (PrivacyRequestCompletedEventData, error) {
	var out PrivacyRequestCompletedEventData
	if e.Topic != TopicPrivacyRequestCompleted {
		return out, &WrongTopicError{Want: TopicPrivacyRequestCompleted, Got: e.Topic}
	}
	err := json.Unmarshal(e.Data, &out)
	return out, err
}

// AsRefereeRewardEarned decodes the data of a `referee_reward.earned` event (A referred customer earned their reward).
// It returns a *WrongTopicError for an event of another topic.
func (e *WebhookEvent) AsRefereeRewardEarned() (RefereeRewardEarnedEventData, error) {
	var out RefereeRewardEarnedEventData
	if e.Topic != TopicRefereeRewardEarned {
		return out, &WrongTopicError{Want: TopicRefereeRewardEarned, Got: e.Topic}
	}
	err := json.Unmarshal(e.Data, &out)
	return out, err
}

// AsReferralApproved decodes the data of a `referral.approved` event (A held referral was approved).
// It returns a *WrongTopicError for an event of another topic.
func (e *WebhookEvent) AsReferralApproved() (ReferralApprovedEventData, error) {
	var out ReferralApprovedEventData
	if e.Topic != TopicReferralApproved {
		return out, &WrongTopicError{Want: TopicReferralApproved, Got: e.Topic}
	}
	err := json.Unmarshal(e.Data, &out)
	return out, err
}

// AsReferralCreated decodes the data of a `referral.created` event (A customer was attributed to a partner).
// It returns a *WrongTopicError for an event of another topic.
func (e *WebhookEvent) AsReferralCreated() (ReferralCreatedEventData, error) {
	var out ReferralCreatedEventData
	if e.Topic != TopicReferralCreated {
		return out, &WrongTopicError{Want: TopicReferralCreated, Got: e.Topic}
	}
	err := json.Unmarshal(e.Data, &out)
	return out, err
}

// AsReferralHeld decodes the data of a `referral.held` event (A referral was held for review).
// It returns a *WrongTopicError for an event of another topic.
func (e *WebhookEvent) AsReferralHeld() (ReferralHeldEventData, error) {
	var out ReferralHeldEventData
	if e.Topic != TopicReferralHeld {
		return out, &WrongTopicError{Want: TopicReferralHeld, Got: e.Topic}
	}
	err := json.Unmarshal(e.Data, &out)
	return out, err
}

// AsReferralPaused decodes the data of a `referral.paused` event (A referral was paused).
// It returns a *WrongTopicError for an event of another topic.
func (e *WebhookEvent) AsReferralPaused() (ReferralPausedEventData, error) {
	var out ReferralPausedEventData
	if e.Topic != TopicReferralPaused {
		return out, &WrongTopicError{Want: TopicReferralPaused, Got: e.Topic}
	}
	err := json.Unmarshal(e.Data, &out)
	return out, err
}

// AsReferralReassigned decodes the data of a `referral.reassigned` event (A referral moved to another partner).
// It returns a *WrongTopicError for an event of another topic.
func (e *WebhookEvent) AsReferralReassigned() (ReferralReassignedEventData, error) {
	var out ReferralReassignedEventData
	if e.Topic != TopicReferralReassigned {
		return out, &WrongTopicError{Want: TopicReferralReassigned, Got: e.Topic}
	}
	err := json.Unmarshal(e.Data, &out)
	return out, err
}

// AsReferralRejected decodes the data of a `referral.rejected` event (A self-referral was refused).
// It returns a *WrongTopicError for an event of another topic.
func (e *WebhookEvent) AsReferralRejected() (ReferralRejectedEventData, error) {
	var out ReferralRejectedEventData
	if e.Topic != TopicReferralRejected {
		return out, &WrongTopicError{Want: TopicReferralRejected, Got: e.Topic}
	}
	err := json.Unmarshal(e.Data, &out)
	return out, err
}

// AsReferralResumed decodes the data of a `referral.resumed` event (A paused referral was resumed).
// It returns a *WrongTopicError for an event of another topic.
func (e *WebhookEvent) AsReferralResumed() (ReferralResumedEventData, error) {
	var out ReferralResumedEventData
	if e.Topic != TopicReferralResumed {
		return out, &WrongTopicError{Want: TopicReferralResumed, Got: e.Topic}
	}
	err := json.Unmarshal(e.Data, &out)
	return out, err
}

// AsReferralVoided decodes the data of a `referral.voided` event (A referral was voided).
// It returns a *WrongTopicError for an event of another topic.
func (e *WebhookEvent) AsReferralVoided() (ReferralVoidedEventData, error) {
	var out ReferralVoidedEventData
	if e.Topic != TopicReferralVoided {
		return out, &WrongTopicError{Want: TopicReferralVoided, Got: e.Topic}
	}
	err := json.Unmarshal(e.Data, &out)
	return out, err
}

// AsRoiPatternDetected decodes the data of a `roi.pattern_detected` event (A value pattern changed).
// It returns a *WrongTopicError for an event of another topic.
func (e *WebhookEvent) AsRoiPatternDetected() (RoiPatternDetectedEventData, error) {
	var out RoiPatternDetectedEventData
	if e.Topic != TopicRoiPatternDetected {
		return out, &WrongTopicError{Want: TopicRoiPatternDetected, Got: e.Topic}
	}
	err := json.Unmarshal(e.Data, &out)
	return out, err
}

// AsRoiPaybackReached decodes the data of a `roi.payback_reached` event (A credit grant paid back).
// It returns a *WrongTopicError for an event of another topic.
func (e *WebhookEvent) AsRoiPaybackReached() (RoiPaybackReachedEventData, error) {
	var out RoiPaybackReachedEventData
	if e.Topic != TopicRoiPaybackReached {
		return out, &WrongTopicError{Want: TopicRoiPaybackReached, Got: e.Topic}
	}
	err := json.Unmarshal(e.Data, &out)
	return out, err
}

// AsTaxFormDelivered decodes the data of a `tax.form_delivered` event (A tax form was delivered to a partner).
// It returns a *WrongTopicError for an event of another topic.
func (e *WebhookEvent) AsTaxFormDelivered() (TaxFormDeliveredEventData, error) {
	var out TaxFormDeliveredEventData
	if e.Topic != TopicTaxFormDelivered {
		return out, &WrongTopicError{Want: TopicTaxFormDelivered, Got: e.Topic}
	}
	err := json.Unmarshal(e.Data, &out)
	return out, err
}

// AsWebhookEndpointDisabled decodes the data of a `webhook_endpoint.disabled` event (A failing webhook endpoint was disabled).
// It returns a *WrongTopicError for an event of another topic.
func (e *WebhookEvent) AsWebhookEndpointDisabled() (WebhookEndpointDisabledEventData, error) {
	var out WebhookEndpointDisabledEventData
	if e.Topic != TopicWebhookEndpointDisabled {
		return out, &WrongTopicError{Want: TopicWebhookEndpointDisabled, Got: e.Topic}
	}
	err := json.Unmarshal(e.Data, &out)
	return out, err
}
