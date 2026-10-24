package referalli

// Code generated from the Referalli OpenAPI document. DO NOT EDIT.
// Changes belong in the generator, which overwrites this file.

import (
	"context"
	"encoding/json"
	"fmt"
	"iter"
	"net/url"
)

var _ = iter.Seq2[int, error](nil)

// Operations lists every operation: id → method and path template.
var Operations = func() map[string][2]string {
	m := map[string][2]string{}
	for _, o := range [][3]string{
		{"deleteApiKeysId", "DELETE", "/v1/api-keys/{id}"},
		{"deleteAssetsId", "DELETE", "/v1/assets/{id}"},
		{"deleteAuthConnectionsId", "DELETE", "/v1/auth/connections/{id}"},
		{"deleteBillingConnectionsId", "DELETE", "/v1/billing/connections/{id}"},
		{"deleteBillingInternalCustomersCustomerRef", "DELETE", "/v1/billing/internal-customers/{customerRef}"},
		{"deleteBrandingDomain", "DELETE", "/v1/branding/domain"},
		{"deleteBrandingLinkDomain", "DELETE", "/v1/branding/link-domain"},
		{"deleteCliListenersId", "DELETE", "/v1/cli/listeners/{id}"},
		{"deleteEmailTemplatesEventLocale", "DELETE", "/v1/email/templates/{event}/{locale}"},
		{"deleteIntegrationsKind", "DELETE", "/v1/integrations/{kind}"},
		{"deleteOrgMembersUserId", "DELETE", "/v1/org/members/{userId}"},
		{"deleteOrgSsoDirect", "DELETE", "/v1/org/sso/direct"},
		{"deletePartnersId", "DELETE", "/v1/partners/{id}"},
		{"deletePayoutConnectionsProvider", "DELETE", "/v1/payout-connections/{provider}"},
		{"deletePrivacyExecutorsAdapter", "DELETE", "/v1/privacy/executors/{adapter}"},
		{"deleteProfitabilityProductsProductRef", "DELETE", "/v1/profitability/products/{productRef}"},
		{"deleteRewardsRulesKey", "DELETE", "/v1/rewards/rules/{key}"},
		{"deleteSupportAppsAdapter", "DELETE", "/v1/support/apps/{adapter}"},
		{"deleteWebhookEndpointsId", "DELETE", "/v1/webhook-endpoints/{id}"},
		{"deleteZapierHooksId", "DELETE", "/v1/zapier/hooks/{id}"},
		{"getAdminMetrics", "GET", "/v1/admin/metrics"},
		{"getAdminOrgs", "GET", "/v1/admin/orgs"},
		{"getAdminPartnerAccounts", "GET", "/v1/admin/partner-accounts"},
		{"getAgents", "GET", "/v1/agents"},
		{"getAgentsRecommendations", "GET", "/v1/agents/recommendations"},
		{"getAnalyticsByChannel", "GET", "/v1/analytics/by-channel"},
		{"getAnalyticsByGeo", "GET", "/v1/analytics/by-geo"},
		{"getAnalyticsCohortLtv", "GET", "/v1/analytics/cohort-ltv"},
		{"getAnalyticsCouponPerformance", "GET", "/v1/analytics/coupon-performance"},
		{"getAnalyticsOverview", "GET", "/v1/analytics/overview"},
		{"getAnalyticsTimeseries", "GET", "/v1/analytics/timeseries"},
		{"getApiKeys", "GET", "/v1/api-keys"},
		{"getApiKeysId", "GET", "/v1/api-keys/{id}"},
		{"getAssets", "GET", "/v1/assets"},
		{"getAssistantSuggestions", "GET", "/v1/assistant/suggestions"},
		{"getAuditLogs", "GET", "/v1/audit-logs"},
		{"getAuthConnections", "GET", "/v1/auth/connections"},
		{"getAuthConnectionsIdExport", "GET", "/v1/auth/connections/{id}/export"},
		{"getAuthConnectionsProviders", "GET", "/v1/auth/connections/providers"},
		{"getAuthProviders", "GET", "/v1/auth/providers"},
		{"getBillingConnections", "GET", "/v1/billing/connections"},
		{"getBillingConnectionsId", "GET", "/v1/billing/connections/{id}"},
		{"getBillingEvents", "GET", "/v1/billing/events"},
		{"getBillingInternalCustomers", "GET", "/v1/billing/internal-customers"},
		{"getBillingPlan", "GET", "/v1/billing/plan"},
		{"getBranding", "GET", "/v1/branding"},
		{"getChatlyIdentity", "GET", "/v1/chatly/identity"},
		{"getCleanroomConsent", "GET", "/v1/cleanroom/consent"},
		{"getCleanroomQuery", "GET", "/v1/cleanroom/query"},
		{"getClicksConsentCoverage", "GET", "/v1/clicks/consent-coverage"},
		{"getClicksId", "GET", "/v1/clicks/{id}"},
		{"getCliDeviceCodesUserCode", "GET", "/v1/cli/device-codes/{userCode}"},
		{"getCliListenersIdEvents", "GET", "/v1/cli/listeners/{id}/events"},
		{"getCliLogs", "GET", "/v1/cli/logs"},
		{"getCliTopics", "GET", "/v1/cli/topics"},
		{"getCommissionRules", "GET", "/v1/commission-rules"},
		{"getCommissions", "GET", "/v1/commissions"},
		{"getCommissionsId", "GET", "/v1/commissions/{id}"},
		{"getContracts", "GET", "/v1/contracts"},
		{"getCoupons", "GET", "/v1/coupons"},
		{"getCreators", "GET", "/v1/creators"},
		{"getCreatorsAccountId", "GET", "/v1/creators/{accountId}"},
		{"getDeals", "GET", "/v1/deals"},
		{"getDisputes", "GET", "/v1/disputes"},
		{"getEmailSettings", "GET", "/v1/email/settings"},
		{"getEmailTemplates", "GET", "/v1/email/templates"},
		{"getEmbed", "GET", "/v1/embed"},
		{"getEmbedSummary", "GET", "/v1/embed/summary"},
		{"getEvents", "GET", "/v1/events"},
		{"getExportAccounting", "GET", "/v1/export/accounting"},
		{"getExportFull", "GET", "/v1/export/full"},
		{"getFraudCases", "GET", "/v1/fraud/cases"},
		{"getFxRates", "GET", "/v1/fx-rates"},
		{"getGraphql", "GET", "/v1/graphql"},
		{"getHealth", "GET", "/v1/health"},
		{"getHostsCurrent", "GET", "/v1/hosts/current"},
		{"getHostsCurrentThemeCss", "GET", "/v1/hosts/current/theme.css"},
		{"getIncrementality", "GET", "/v1/incrementality"},
		{"getIntegrationHealth", "GET", "/v1/integration-health"},
		{"getIntegrations", "GET", "/v1/integrations"},
		{"getIntegrationsDeliveries", "GET", "/v1/integrations/deliveries"},
		{"getInvitations", "GET", "/v1/invitations"},
		{"getLinks", "GET", "/v1/links"},
		{"getLinksIdQr", "GET", "/v1/links/{id}/qr"},
		{"getMarketplaceListing", "GET", "/v1/marketplace/listing"},
		{"getMe", "GET", "/v1/me"},
		{"getNegotiations", "GET", "/v1/negotiations"},
		{"getNetworkJourney", "GET", "/v1/network/journey"},
		{"getNetworkSettings", "GET", "/v1/network/settings"},
		{"getNotifications", "GET", "/v1/notifications"},
		{"getOnchainConnections", "GET", "/v1/onchain/connections"},
		{"getOpenapiJson", "GET", "/v1/openapi.json"},
		{"getOpenapiYaml", "GET", "/v1/openapi.yaml"},
		{"getOrg", "GET", "/v1/org"},
		{"getOrgFxFeed", "GET", "/v1/org/fx-feed"},
		{"getOrgHostedPages", "GET", "/v1/org/hosted-pages"},
		{"getOrgMembers", "GET", "/v1/org/members"},
		{"getOrgPayoutSchedule", "GET", "/v1/org/payout-schedule"},
		{"getOrgPortalAuth", "GET", "/v1/org/portal-auth"},
		{"getOrgReferralParams", "GET", "/v1/org/referral-params"},
		{"getOrgRetention", "GET", "/v1/org/retention"},
		{"getOrgSso", "GET", "/v1/org/sso"},
		{"getOrgSsoConnectionsIdSamlMetadata", "GET", "/v1/org/sso/connections/{id}/saml-metadata"},
		{"getOrgTaxProfile", "GET", "/v1/org/tax-profile"},
		{"getOrgTrackingConsent", "GET", "/v1/org/tracking-consent"},
		{"getOutreach", "GET", "/v1/outreach"},
		{"getPartnerAchievements", "GET", "/v1/partner/achievements"},
		{"getPartnerAdvance", "GET", "/v1/partner/advance"},
		{"getPartnerAdvanceOffer", "GET", "/v1/partner/advance/offer"},
		{"getPartnerAssets", "GET", "/v1/partner/assets"},
		{"getPartnerContracts", "GET", "/v1/partner/contracts"},
		{"getPartnerDeals", "GET", "/v1/partner/deals"},
		{"getPartnerDisputes", "GET", "/v1/partner/disputes"},
		{"getPartnerGroups", "GET", "/v1/partner/groups"},
		{"getPartnerInvitations", "GET", "/v1/partner/invitations"},
		{"getPartnerKycStatus", "GET", "/v1/partner/kyc/status"},
		{"getPartnerLinks", "GET", "/v1/partner/links"},
		{"getPartnerLinksIdQr", "GET", "/v1/partner/links/{id}/qr"},
		{"getPartnerMarketplace", "GET", "/v1/partner/marketplace"},
		{"getPartnerMe", "GET", "/v1/partner/me"},
		{"getPartnerNotifications", "GET", "/v1/partner/notifications"},
		{"getPartnerOrgsOrgIdSelfBilling", "GET", "/v1/partner/orgs/{orgId}/self-billing"},
		{"getPartnerPassport", "GET", "/v1/partner/passport"},
		{"getPartnerPayouts", "GET", "/v1/partner/payouts"},
		{"getPartnerPayoutsIdInvoice", "GET", "/v1/partner/payouts/{id}/invoice"},
		{"getPartnerPostbackLogs", "GET", "/v1/partner/postback-logs"},
		{"getPartnerProfile", "GET", "/v1/partner/profile"},
		{"getPartnerProfileVerified", "GET", "/v1/partner/profile/verified"},
		{"getPartnerProgramsProgramIdLeaderboard", "GET", "/v1/partner/programs/{programId}/leaderboard"},
		{"getPartnerRecruit", "GET", "/v1/partner/recruit"},
		{"getPartnerReferrals", "GET", "/v1/partner/referrals"},
		{"getPartnerReferralsIdTransparency", "GET", "/v1/partner/referrals/{id}/transparency"},
		{"getPartners", "GET", "/v1/partners"},
		{"getPartnerSecurity", "GET", "/v1/partner/security"},
		{"getPartnersId", "GET", "/v1/partners/{id}"},
		{"getPartnersIdBalances", "GET", "/v1/partners/{id}/balances"},
		{"getPartnersIdEnrollments", "GET", "/v1/partners/{id}/enrollments"},
		{"getPartnersIdPayoutProfile", "GET", "/v1/partners/{id}/payout-profile"},
		{"getPartnersIdQuality", "GET", "/v1/partners/{id}/quality"},
		{"getPartnerStats", "GET", "/v1/partner/stats"},
		{"getPartnerTaxDocuments", "GET", "/v1/partner/tax-documents"},
		{"getPartnerTaxDocumentsId", "GET", "/v1/partner/tax-documents/{id}"},
		{"getPartnerTerms", "GET", "/v1/partner/terms"},
		{"getPassportPublicKey", "GET", "/v1/passport/public-key"},
		{"getPayoutConnections", "GET", "/v1/payout-connections"},
		{"getPayoutRuns", "GET", "/v1/payout-runs"},
		{"getPayouts", "GET", "/v1/payouts"},
		{"getPayoutsIdInvoice", "GET", "/v1/payouts/{id}/invoice"},
		{"getPlans", "GET", "/v1/plans"},
		{"getPortalAuthCurrent", "GET", "/v1/portal-auth/current"},
		{"getPrivacyExecutorRequests", "GET", "/v1/privacy/executor-requests"},
		{"getPrivacyExecutors", "GET", "/v1/privacy/executors"},
		{"getPrivacyExport", "GET", "/v1/privacy/export"},
		{"getPrivacyRequests", "GET", "/v1/privacy/requests"},
		{"getPrivacyRequestsId", "GET", "/v1/privacy/requests/{id}"},
		{"getPrivacyRequestsIdResult", "GET", "/v1/privacy/requests/{id}/result"},
		{"getProfitabilityProducts", "GET", "/v1/profitability/products"},
		{"getProfitabilityProductsProductRef", "GET", "/v1/profitability/products/{productRef}"},
		{"getProfitabilityRebuildJobId", "GET", "/v1/profitability/rebuild/{jobId}"},
		{"getProfitabilitySettings", "GET", "/v1/profitability/settings"},
		{"getPrograms", "GET", "/v1/programs"},
		{"getProgramsId", "GET", "/v1/programs/{id}"},
		{"getProgramsIdAutopricing", "GET", "/v1/programs/{id}/autopricing"},
		{"getProgramsIdEnrollments", "GET", "/v1/programs/{id}/enrollments"},
		{"getProgramsIdGraph", "GET", "/v1/programs/{id}/graph"},
		{"getProgramsIdGroups", "GET", "/v1/programs/{id}/groups"},
		{"getProgramsIdNegotiationPolicy", "GET", "/v1/programs/{id}/negotiation-policy"},
		{"getProgramsIdPromotions", "GET", "/v1/programs/{id}/promotions"},
		{"getProgramsIdRefereeReward", "GET", "/v1/programs/{id}/referee-reward"},
		{"getProgramsIdTermsAcceptances", "GET", "/v1/programs/{id}/terms-acceptances"},
		{"getProspects", "GET", "/v1/prospects"},
		{"getPublicProgramsHandle", "GET", "/v1/public/programs/{handle}"},
		{"getReferrals", "GET", "/v1/referrals"},
		{"getReferralsId", "GET", "/v1/referrals/{id}"},
		{"getReferralsIdJourney", "GET", "/v1/referrals/{id}/journey"},
		{"getReportsFilteredClicks", "GET", "/v1/reports/filtered-clicks"},
		{"getReportsPartnerPerformance", "GET", "/v1/reports/partner-performance"},
		{"getRewards", "GET", "/v1/rewards"},
		{"getRewardsGranted", "GET", "/v1/rewards/granted"},
		{"getRewardsLookupCode", "GET", "/v1/rewards/lookup/{code}"},
		{"getRewardsRules", "GET", "/v1/rewards/rules"},
		{"getRoiBreakdown", "GET", "/v1/roi/breakdown"},
		{"getRoiDataQuality", "GET", "/v1/roi/data-quality"},
		{"getRoiGrants", "GET", "/v1/roi/grants"},
		{"getRoiPartnersIdTrace", "GET", "/v1/roi/partners/{id}/trace"},
		{"getRoiPatterns", "GET", "/v1/roi/patterns"},
		{"getRoiPatternsPattern", "GET", "/v1/roi/patterns/{pattern}"},
		{"getRoiSummary", "GET", "/v1/roi/summary"},
		{"getRoiTraceCustomerRef", "GET", "/v1/roi/trace/{customerRef}"},
		{"getSetupStatus", "GET", "/v1/setup/status"},
		{"getStatus", "GET", "/v1/status"},
		{"getStatusIncidents", "GET", "/v1/status/incidents"},
		{"getStatusIncidentsAtom", "GET", "/v1/status/incidents.atom"},
		{"getStatusIncidentsId", "GET", "/v1/status/incidents/{id}"},
		{"getSupportApps", "GET", "/v1/support/apps"},
		{"getSupportCardsOrgIdAdapter", "GET", "/v1/support/cards/{orgId}/{adapter}"},
		{"getSupportContext", "GET", "/v1/support/context"},
		{"getTax1099", "GET", "/v1/tax/1099"},
		{"getTax1099Filings", "GET", "/v1/tax/1099/filings"},
		{"getTax1099Preview", "GET", "/v1/tax/1099/preview"},
		{"getTestClock", "GET", "/v1/test_clock"},
		{"getTestEmails", "GET", "/v1/test/emails"},
		{"getValueImportJobId", "GET", "/v1/value/import/{jobId}"},
		{"getVersions", "GET", "/v1/versions"},
		{"getWebhookDeliveries", "GET", "/v1/webhook-deliveries"},
		{"getWebhookDeliveriesId", "GET", "/v1/webhook-deliveries/{id}"},
		{"getWebhookEndpoints", "GET", "/v1/webhook-endpoints"},
		{"getWebhookEndpointsId", "GET", "/v1/webhook-endpoints/{id}"},
		{"getWebhookEndpointsIdDeliveries", "GET", "/v1/webhook-endpoints/{id}/deliveries"},
		{"getZapierEvents", "GET", "/v1/zapier/events"},
		{"getZapierHooks", "GET", "/v1/zapier/hooks"},
		{"getZapierTriggersEvent", "GET", "/v1/zapier/triggers/{event}"},
		{"patchAdminStatusIncidentsId", "PATCH", "/v1/admin/status/incidents/{id}"},
		{"patchApiKeysId", "PATCH", "/v1/api-keys/{id}"},
		{"patchAuthConnectionsId", "PATCH", "/v1/auth/connections/{id}"},
		{"patchBillingConnectionsId", "PATCH", "/v1/billing/connections/{id}"},
		{"patchOrg", "PATCH", "/v1/org"},
		{"patchOrgMembersUserId", "PATCH", "/v1/org/members/{userId}"},
		{"patchOrgSso", "PATCH", "/v1/org/sso"},
		{"patchPartnerGroupsId", "PATCH", "/v1/partner-groups/{id}"},
		{"patchPartnersId", "PATCH", "/v1/partners/{id}"},
		{"patchPartnersIdPayoutProfile", "PATCH", "/v1/partners/{id}/payout-profile"},
		{"patchProgramsId", "PATCH", "/v1/programs/{id}"},
		{"patchReferralsId", "PATCH", "/v1/referrals/{id}"},
		{"patchWebhookEndpointsId", "PATCH", "/v1/webhook-endpoints/{id}"},
		{"postAccountClose", "POST", "/v1/account/close"},
		{"postAdminOrgsIdPlan", "POST", "/v1/admin/orgs/{id}/plan"},
		{"postAdminOrgsIdSuspend", "POST", "/v1/admin/orgs/{id}/suspend"},
		{"postAdminOrgsIdUnsuspend", "POST", "/v1/admin/orgs/{id}/unsuspend"},
		{"postAdminPartnerAccountsIdBan", "POST", "/v1/admin/partner-accounts/{id}/ban"},
		{"postAdminPartnerAccountsIdUnban", "POST", "/v1/admin/partner-accounts/{id}/unban"},
		{"postAdminStatusIncidents", "POST", "/v1/admin/status/incidents"},
		{"postAdminStatusIncidentsIdUpdates", "POST", "/v1/admin/status/incidents/{id}/updates"},
		{"postAgentsRegister", "POST", "/v1/agents/register"},
		{"postAnomaliesScan", "POST", "/v1/anomalies/scan"},
		{"postApiKeys", "POST", "/v1/api-keys"},
		{"postAssets", "POST", "/v1/assets"},
		{"postAttributionReplay", "POST", "/v1/attribution/replay"},
		{"postAuthBackchannelLogout", "POST", "/v1/auth/backchannel-logout"},
		{"postAuthConnections", "POST", "/v1/auth/connections"},
		{"postAuthConnectionsIdAliases", "POST", "/v1/auth/connections/{id}/aliases"},
		{"postAuthConnectionsIdErase", "POST", "/v1/auth/connections/{id}/erase"},
		{"postAuthConnectionsIdPartners", "POST", "/v1/auth/connections/{id}/partners"},
		{"postAuthConnectionsIdPing", "POST", "/v1/auth/connections/{id}/ping"},
		{"postAuthConnectionsIdRevokePreviousSecret", "POST", "/v1/auth/connections/{id}/revoke-previous-secret"},
		{"postAuthConnectionsIdRotateSecret", "POST", "/v1/auth/connections/{id}/rotate-secret"},
		{"postAuthConnectionsIdSignals", "POST", "/v1/auth/connections/{id}/signals"},
		{"postAuthConnectionsIdSignups", "POST", "/v1/auth/connections/{id}/signups"},
		{"postAuthConnectionsIdUsers", "POST", "/v1/auth/connections/{id}/users"},
		{"postAuthConnectionsInstall", "POST", "/v1/auth/connections/install"},
		{"postAuthLogout", "POST", "/v1/auth/logout"},
		{"postBillingActivate", "POST", "/v1/billing/activate"},
		{"postBillingCheckout", "POST", "/v1/billing/checkout"},
		{"postBillingConnections", "POST", "/v1/billing/connections"},
		{"postBillingConnectionsIdRotateSecret", "POST", "/v1/billing/connections/{id}/rotate-secret"},
		{"postBillingConnectionsIdValueMappingPreview", "POST", "/v1/billing/connections/{id}/value-mapping/preview"},
		{"postBillingInternalCustomers", "POST", "/v1/billing/internal-customers"},
		{"postBillingPortal", "POST", "/v1/billing/portal"},
		{"postBrandingDomain", "POST", "/v1/branding/domain"},
		{"postBrandingDomainVerify", "POST", "/v1/branding/domain/verify"},
		{"postBrandingLinkDomain", "POST", "/v1/branding/link-domain"},
		{"postBrandingLinkDomainVerify", "POST", "/v1/branding/link-domain/verify"},
		{"postCliDeviceCodes", "POST", "/v1/cli/device-codes"},
		{"postCliDeviceCodesToken", "POST", "/v1/cli/device-codes/token"},
		{"postCliDeviceCodesUserCodeApprove", "POST", "/v1/cli/device-codes/{userCode}/approve"},
		{"postCliDeviceCodesUserCodeDeny", "POST", "/v1/cli/device-codes/{userCode}/deny"},
		{"postCliListeners", "POST", "/v1/cli/listeners"},
		{"postCliTrigger", "POST", "/v1/cli/trigger"},
		{"postCommissionRules", "POST", "/v1/commission-rules"},
		{"postCommissionsApproveDue", "POST", "/v1/commissions/approve-due"},
		{"postCommissionsIdAdjust", "POST", "/v1/commissions/{id}/adjust"},
		{"postCommissionsIdSettle", "POST", "/v1/commissions/{id}/settle"},
		{"postContracts", "POST", "/v1/contracts"},
		{"postContractsIdSend", "POST", "/v1/contracts/{id}/send"},
		{"postCoupons", "POST", "/v1/coupons"},
		{"postCreatorsAccountIdInvite", "POST", "/v1/creators/{accountId}/invite"},
		{"postDealsIdApprove", "POST", "/v1/deals/{id}/approve"},
		{"postDealsIdReject", "POST", "/v1/deals/{id}/reject"},
		{"postDealsIdWon", "POST", "/v1/deals/{id}/won"},
		{"postDiscoveryCorpus", "POST", "/v1/discovery/corpus"},
		{"postDiscoverySearch", "POST", "/v1/discovery/search"},
		{"postDisputesIdResolve", "POST", "/v1/disputes/{id}/resolve"},
		{"postEmailPreview", "POST", "/v1/email/preview"},
		{"postEmailSettingsVerifyDomain", "POST", "/v1/email/settings/verify-domain"},
		{"postEmailTest", "POST", "/v1/email/test"},
		{"postEmbedTokens", "POST", "/v1/embed/tokens"},
		{"postFraudCasesIdResolve", "POST", "/v1/fraud/cases/{id}/resolve"},
		{"postFraudDetectRings", "POST", "/v1/fraud/detect-rings"},
		{"postFraudInvalidTrafficScan", "POST", "/v1/fraud/invalid-traffic/scan"},
		{"postFxRates", "POST", "/v1/fx-rates"},
		{"postFxRatesClearOverride", "POST", "/v1/fx-rates/clear-override"},
		{"postFxRatesSync", "POST", "/v1/fx-rates/sync"},
		{"postGraphql", "POST", "/v1/graphql"},
		{"postImportProvider", "POST", "/v1/import/{provider}"},
		{"postIntegrationsDeliveriesIdRetry", "POST", "/v1/integrations/deliveries/{id}/retry"},
		{"postIntegrationsHubspotTest", "POST", "/v1/integrations/hubspot/test"},
		{"postIntegrationsSegmentTest", "POST", "/v1/integrations/segment/test"},
		{"postInvitationsIdWithdraw", "POST", "/v1/invitations/{id}/withdraw"},
		{"postLinks", "POST", "/v1/links"},
		{"postMobilePurchases", "POST", "/v1/mobile/purchases"},
		{"postNotificationsIdRead", "POST", "/v1/notifications/{id}/read"},
		{"postOauthToken", "POST", "/v1/oauth/token"},
		{"postOnchainConnections", "POST", "/v1/onchain/connections"},
		{"postOrgHostedPagesCheck", "POST", "/v1/org/hosted-pages/check"},
		{"postOrgMembers", "POST", "/v1/org/members"},
		{"postOrgPortalAuthTest", "POST", "/v1/org/portal-auth/test"},
		{"postOrgs", "POST", "/v1/orgs"},
		{"postOrgSsoDirectTest", "POST", "/v1/org/sso/direct/test"},
		{"postOrgSsoLink", "POST", "/v1/org/sso/link"},
		{"postOrgSsoTickets", "POST", "/v1/org/sso/tickets"},
		{"postOrgSsoTicketsIdRevoke", "POST", "/v1/org/sso/tickets/{id}/revoke"},
		{"postOutreachDraft", "POST", "/v1/outreach/draft"},
		{"postOutreachIdApprove", "POST", "/v1/outreach/{id}/approve"},
		{"postOutreachSendDue", "POST", "/v1/outreach/send-due"},
		{"postPartnerAdvanceAccept", "POST", "/v1/partner/advance/accept"},
		{"postPartnerAuthWalletNonce", "POST", "/v1/partner/auth/wallet/nonce"},
		{"postPartnerAuthWalletVerify", "POST", "/v1/partner/auth/wallet/verify"},
		{"postPartnerContractsIdAccept", "POST", "/v1/partner/contracts/{id}/accept"},
		{"postPartnerContractsIdDecline", "POST", "/v1/partner/contracts/{id}/decline"},
		{"postPartnerDeals", "POST", "/v1/partner/deals"},
		{"postPartnerDisputes", "POST", "/v1/partner/disputes"},
		{"postPartnerInvitationsIdAccept", "POST", "/v1/partner/invitations/{id}/accept"},
		{"postPartnerInvitationsIdDecline", "POST", "/v1/partner/invitations/{id}/decline"},
		{"postPartnerKycSubmit", "POST", "/v1/partner/kyc/submit"},
		{"postPartnerLinks", "POST", "/v1/partner/links"},
		{"postPartnerNotificationsIdRead", "POST", "/v1/partner/notifications/{id}/read"},
		{"postPartnerOrgsOrgIdSelfBilling", "POST", "/v1/partner/orgs/{orgId}/self-billing"},
		{"postPartnerOrgsOrgIdTaxForm", "POST", "/v1/partner/orgs/{orgId}/tax-form"},
		{"postPartnerProgramsProgramIdApply", "POST", "/v1/partner/programs/{programId}/apply"},
		{"postPartners", "POST", "/v1/partners"},
		{"postPartnerSecurityTotpDisable", "POST", "/v1/partner/security/totp/disable"},
		{"postPartnerSecurityTotpEnable", "POST", "/v1/partner/security/totp/enable"},
		{"postPartnerSecurityTotpSetup", "POST", "/v1/partner/security/totp/setup"},
		{"postPartnersIdTaxForm", "POST", "/v1/partners/{id}/tax-form"},
		{"postPartnerTermsVersionIdAccept", "POST", "/v1/partner/terms/{versionId}/accept"},
		{"postPassportVerify", "POST", "/v1/passport/verify"},
		{"postPayoutRuns", "POST", "/v1/payout-runs"},
		{"postPayoutRunsIdCancel", "POST", "/v1/payout-runs/{id}/cancel"},
		{"postPayoutRunsIdExecute", "POST", "/v1/payout-runs/{id}/execute"},
		{"postPayoutRunsRunScheduled", "POST", "/v1/payout-runs/run-scheduled"},
		{"postPrivacyExecutorKeys", "POST", "/v1/privacy/executor-keys"},
		{"postPrivacyExecutorsAdapterCheck", "POST", "/v1/privacy/executors/{adapter}/check"},
		{"postPrivacyForget", "POST", "/v1/privacy/forget"},
		{"postPrivacyLookup", "POST", "/v1/privacy/lookup"},
		{"postPrivacyRequests", "POST", "/v1/privacy/requests"},
		{"postProfitabilityRebuild", "POST", "/v1/profitability/rebuild"},
		{"postPrograms", "POST", "/v1/programs"},
		{"postProgramsIdApplyLeveling", "POST", "/v1/programs/{id}/apply-leveling"},
		{"postProgramsIdAutopriceRun", "POST", "/v1/programs/{id}/autoprice/run"},
		{"postProgramsIdEnrollments", "POST", "/v1/programs/{id}/enrollments"},
		{"postProgramsIdEnrollmentsEnrollmentIdApprove", "POST", "/v1/programs/{id}/enrollments/{enrollmentId}/approve"},
		{"postProgramsIdEnrollmentsEnrollmentIdReject", "POST", "/v1/programs/{id}/enrollments/{enrollmentId}/reject"},
		{"postProgramsIdGroups", "POST", "/v1/programs/{id}/groups"},
		{"postProgramsIdPromotions", "POST", "/v1/programs/{id}/promotions"},
		{"postPromotionsIdEnd", "POST", "/v1/promotions/{id}/end"},
		{"postPublicProgramsHandleJoin", "POST", "/v1/public/programs/{handle}/join"},
		{"postQualityRecompute", "POST", "/v1/quality/recompute"},
		{"postReferralsAssociate", "POST", "/v1/referrals/associate"},
		{"postReferralsIdApprove", "POST", "/v1/referrals/{id}/approve"},
		{"postReferralsIdReassign", "POST", "/v1/referrals/{id}/reassign"},
		{"postReferralsIdVoid", "POST", "/v1/referrals/{id}/void"},
		{"postRewards", "POST", "/v1/rewards"},
		{"postRewardsIdRedeem", "POST", "/v1/rewards/{id}/redeem"},
		{"postRewardsRedeemByCode", "POST", "/v1/rewards/redeem-by-code"},
		{"postRoiPatternsPatternReview", "POST", "/v1/roi/patterns/{pattern}/review"},
		{"postSandboxSimulate", "POST", "/v1/sandbox/simulate"},
		{"postSupportAppsAdapterTest", "POST", "/v1/support/apps/{adapter}/test"},
		{"postSupportCardsOrgIdAdapter", "POST", "/v1/support/cards/{orgId}/{adapter}"},
		{"postTax1099DocumentsIdCorrect", "POST", "/v1/tax/1099/documents/{id}/correct"},
		{"postTax1099Filings", "POST", "/v1/tax/1099/filings"},
		{"postTax1099FilingsIdRefresh", "POST", "/v1/tax/1099/filings/{id}/refresh"},
		{"postTestClockAdvance", "POST", "/v1/test_clock/advance"},
		{"postTrackEvents", "POST", "/v1/track/events"},
		{"postUnsubscribeToken", "POST", "/v1/unsubscribe/{token}"},
		{"postValueEvents", "POST", "/v1/value/events"},
		{"postValueImport", "POST", "/v1/value/import"},
		{"postWebhookDeliveriesIdReplay", "POST", "/v1/webhook-deliveries/{id}/replay"},
		{"postWebhookEndpoints", "POST", "/v1/webhook-endpoints"},
		{"postWebhookEndpointsIdReplay", "POST", "/v1/webhook-endpoints/{id}/replay"},
		{"postWebhookEndpointsIdRotateSecret", "POST", "/v1/webhook-endpoints/{id}/rotate-secret"},
		{"postWidgetTokens", "POST", "/v1/widget/tokens"},
		{"postZapierConversions", "POST", "/v1/zapier/conversions"},
		{"postZapierHooks", "POST", "/v1/zapier/hooks"},
		{"putAuthConnectionsIdSecret", "PUT", "/v1/auth/connections/{id}/secret"},
		{"putBillingConnectionsIdCredentials", "PUT", "/v1/billing/connections/{id}/credentials"},
		{"putBillingConnectionsIdSecret", "PUT", "/v1/billing/connections/{id}/secret"},
		{"putBillingConnectionsIdValueMapping", "PUT", "/v1/billing/connections/{id}/value-mapping"},
		{"putBranding", "PUT", "/v1/branding"},
		{"putEmailSettings", "PUT", "/v1/email/settings"},
		{"putEmailTemplatesEventLocale", "PUT", "/v1/email/templates/{event}/{locale}"},
		{"putIntegrationsHubspot", "PUT", "/v1/integrations/hubspot"},
		{"putIntegrationsSegment", "PUT", "/v1/integrations/segment"},
		{"putMarketplaceListing", "PUT", "/v1/marketplace/listing"},
		{"putMeLocale", "PUT", "/v1/me/locale"},
		{"putMeRoiDigest", "PUT", "/v1/me/roi-digest"},
		{"putNetworkSettings", "PUT", "/v1/network/settings"},
		{"putOrgFxFeed", "PUT", "/v1/org/fx-feed"},
		{"putOrgHostedPages", "PUT", "/v1/org/hosted-pages"},
		{"putOrgPayoutSchedule", "PUT", "/v1/org/payout-schedule"},
		{"putOrgPortalAuth", "PUT", "/v1/org/portal-auth"},
		{"putOrgReferralParams", "PUT", "/v1/org/referral-params"},
		{"putOrgRetention", "PUT", "/v1/org/retention"},
		{"putOrgSsoDirect", "PUT", "/v1/org/sso/direct"},
		{"putOrgTaxFiling", "PUT", "/v1/org/tax-filing"},
		{"putOrgTaxProfile", "PUT", "/v1/org/tax-profile"},
		{"putOrgTrackingConsent", "PUT", "/v1/org/tracking-consent"},
		{"putPartnerLocale", "PUT", "/v1/partner/locale"},
		{"putPartnerOrgsOrgIdPayoutProfile", "PUT", "/v1/partner/orgs/{orgId}/payout-profile"},
		{"putPartnerOrgsOrgIdPostback", "PUT", "/v1/partner/orgs/{orgId}/postback"},
		{"putPartnerOrgsOrgIdVat", "PUT", "/v1/partner/orgs/{orgId}/vat"},
		{"putPartnerProfile", "PUT", "/v1/partner/profile"},
		{"putPartnersIdPayoutProfile", "PUT", "/v1/partners/{id}/payout-profile"},
		{"putPartnersIdVat", "PUT", "/v1/partners/{id}/vat"},
		{"putPayoutConnectionsProvider", "PUT", "/v1/payout-connections/{provider}"},
		{"putProfitabilityProductsProductRef", "PUT", "/v1/profitability/products/{productRef}"},
		{"putProfitabilitySettings", "PUT", "/v1/profitability/settings"},
		{"putProgramsIdAutopricing", "PUT", "/v1/programs/{id}/autopricing"},
		{"putProgramsIdIncrementality", "PUT", "/v1/programs/{id}/incrementality"},
		{"putProgramsIdNegotiationPolicy", "PUT", "/v1/programs/{id}/negotiation-policy"},
		{"putProgramsIdPartnersPartnerIdGroup", "PUT", "/v1/programs/{id}/partners/{partnerId}/group"},
		{"putProgramsIdPayoutSchedule", "PUT", "/v1/programs/{id}/payout-schedule"},
		{"putProgramsIdRefereeReward", "PUT", "/v1/programs/{id}/referee-reward"},
		{"putRewardsRulesKey", "PUT", "/v1/rewards/rules/{key}"},
		{"putSupportAppsAdapter", "PUT", "/v1/support/apps/{adapter}"},
	} {
		m[o[0]] = [2]string{o[1], o[2]}
	}
	return m
}()

// DeleteAPIKeysIDParams are the query parameters of DeleteAPIKeysID.
type DeleteAPIKeysIDParams struct {
	Cascade *string `json:"cascade,omitempty"`
	Pair    *string `json:"pair,omitempty"`
}

func (p *DeleteAPIKeysIDParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "cascade", p.Cascade)
	setQuery(v, "pair", p.Pair)
	return v
}

// DeleteAPIKeysID calls Revoke an API key. DELETE /v1/api-keys/{id}.
//
// Revokes a key at once. `cascade=true` also revokes every key it minted (and theirs); `pair=true` also revokes its live/test twin. Returns the revoked ids. Needs an owner or admin session, or an API key with the `admin` scope. Audited.
//
// **Auth:** an API key with `team:write` (or `full`), or a session of any role.
//
// cascade: true: also revoke every key this key minted, and theirs.
// pair: true: also revoke its live/test twin.
func (c *Client) DeleteAPIKeysID(ctx context.Context, id string, params *DeleteAPIKeysIDParams, opts ...RequestOption) (*DeleteApiKeysIdResponse, error) {
	var out DeleteApiKeysIdResponse
	if err := c.do(ctx, "DELETE", "/v1/api-keys/"+url.PathEscape(id), params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// DeleteAssetsID calls Delete a marketing asset. DELETE /v1/assets/{id}.
//
// Removes a marketing asset.
//
// **Auth:** an API key with `partners:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) DeleteAssetsID(ctx context.Context, id string, opts ...RequestOption) (*Ok, error) {
	var out Ok
	if err := c.do(ctx, "DELETE", "/v1/assets/"+url.PathEscape(id), nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// DeleteAuthConnectionsID calls Delete an auth connection. DELETE /v1/auth/connections/{id}.
//
// Removes the connection; its webhook URL and signed endpoints stop accepting calls. Referrals it filed stay.
//
// **Auth:** an API key with `integrations:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) DeleteAuthConnectionsID(ctx context.Context, id string, opts ...RequestOption) (*Ok, error) {
	var out Ok
	if err := c.do(ctx, "DELETE", "/v1/auth/connections/"+url.PathEscape(id), nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// DeleteBillingConnectionsID calls Delete a billing connection. DELETE /v1/billing/connections/{id}.
//
// Deletes the connection: its secrets and credentials are wiped and its webhook URL stops accepting events. Recorded billing events stay. Audited.
//
// **Auth:** an API key with `billing:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) DeleteBillingConnectionsID(ctx context.Context, id string, opts ...RequestOption) (*Ok, error) {
	var out Ok
	if err := c.do(ctx, "DELETE", "/v1/billing/connections/"+url.PathEscape(id), nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// DeleteBillingInternalCustomersCustomerRef calls Remove an internal customer. DELETE /v1/billing/internal-customers/{customerRef}.
//
// Their payments earn again from now on (events already ignored stay ignored). Audited.
//
// **Auth:** an API key with `billing:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) DeleteBillingInternalCustomersCustomerRef(ctx context.Context, customerRef string, opts ...RequestOption) (*Ok, error) {
	var out Ok
	if err := c.do(ctx, "DELETE", "/v1/billing/internal-customers/"+url.PathEscape(customerRef), nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// DeleteBrandingDomain calls Remove the portal domain. DELETE /v1/branding/domain.
//
// Stops serving the partner portal on the custom domain.
//
// **Auth:** an API key with `org:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) DeleteBrandingDomain(ctx context.Context, opts ...RequestOption) (*Ok, error) {
	var out Ok
	if err := c.do(ctx, "DELETE", "/v1/branding/domain", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// DeleteBrandingLinkDomain calls Remove the link domain. DELETE /v1/branding/link-domain.
//
// Stops using the custom link domain; links go back to the default host.
//
// **Auth:** an API key with `org:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) DeleteBrandingLinkDomain(ctx context.Context, opts ...RequestOption) (*Ok, error) {
	var out Ok
	if err := c.do(ctx, "DELETE", "/v1/branding/link-domain", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// DeleteCliListenersID calls Close a CLI listener. DELETE /v1/cli/listeners/{id}.
//
// Closes the listener; queued events are dropped.
//
// **Auth:** an API key with `webhooks:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) DeleteCliListenersID(ctx context.Context, id string, opts ...RequestOption) (*Ok, error) {
	var out Ok
	if err := c.do(ctx, "DELETE", "/v1/cli/listeners/"+url.PathEscape(id), nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// DeleteEmailTemplatesEventLocale calls Reset an email template. DELETE /v1/email/templates/{event}/{locale}.
//
// Removes the override; the default template is used again.
//
// **Auth:** an API key with `org:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) DeleteEmailTemplatesEventLocale(ctx context.Context, event string, locale string, opts ...RequestOption) (*DeleteEmailTemplatesEventLocaleResponse, error) {
	var out DeleteEmailTemplatesEventLocaleResponse
	if err := c.do(ctx, "DELETE", "/v1/email/templates/"+url.PathEscape(event)+"/"+url.PathEscape(locale), nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// DeleteIntegrationsKind calls Disconnect an integration. DELETE /v1/integrations/{kind}.
//
// Removes the HubSpot or Segment connection and its stored credential.
//
// **Auth:** an API key with `integrations:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) DeleteIntegrationsKind(ctx context.Context, kind string, opts ...RequestOption) (*Ok, error) {
	var out Ok
	if err := c.do(ctx, "DELETE", "/v1/integrations/"+url.PathEscape(kind), nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// DeleteOrgMembersUserID calls Remove a team member. DELETE /v1/org/members/{userId}.
//
// Removes a member. Members can remove themselves; otherwise owner or admin. The last owner stays (409). Audited. Live mode only.
//
// **Auth:** an API key with `team:write` (or `full`), or a session of any role.
func (c *Client) DeleteOrgMembersUserID(ctx context.Context, userId string, opts ...RequestOption) (*Ok, error) {
	var out Ok
	if err := c.do(ctx, "DELETE", "/v1/org/members/"+url.PathEscape(userId), nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// DeleteOrgSSODirect calls Remove your identity provider. DELETE /v1/org/sso/direct.
//
// Removes the org's own IdP configuration; team SSO and "require SSO" turn off, so nobody is locked out. Owner only.
//
// **Auth:** an API key with `org:write` (or `full`), or a session with role `owner`.
func (c *Client) DeleteOrgSSODirect(ctx context.Context, opts ...RequestOption) (*DeleteOrgSsoDirectResponse, error) {
	var out DeleteOrgSsoDirectResponse
	if err := c.do(ctx, "DELETE", "/v1/org/sso/direct", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// DeletePartnersID calls Delete a partner. DELETE /v1/partners/{id}.
//
// Deletes a partner that never did anything, with its enrollments, links, coupons and payout profile. A partner with any clicks, referrals, commissions, payouts or other history is `409`: archive it instead (`status: "archived"`). Audited.
//
// **Auth:** an API key with `partners:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) DeletePartnersID(ctx context.Context, id string, opts ...RequestOption) (*Ok, error) {
	var out Ok
	if err := c.do(ctx, "DELETE", "/v1/partners/"+url.PathEscape(id), nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// DeletePayoutConnectionsProvider calls Disconnect a payout provider. DELETE /v1/payout-connections/{provider}.
//
// Removes the stored credentials for a provider. Payouts by that method are then simulated until you reconnect. Audited.
//
// **Auth:** an API key with `payouts:write` (or `full`), or a session with role `owner`, `admin`, `finance`.
func (c *Client) DeletePayoutConnectionsProvider(ctx context.Context, provider string, opts ...RequestOption) (*DeletePayoutConnectionsProviderResponse, error) {
	var out DeletePayoutConnectionsProviderResponse
	if err := c.do(ctx, "DELETE", "/v1/payout-connections/"+url.PathEscape(provider), nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// DeletePrivacyExecutorsAdapter calls Disconnect a DSAR tool. DELETE /v1/privacy/executors/{adapter}.
//
// Removes the tool, its key and settings; its URL answers 404 from now on. The request log stays.
//
// **Auth:** an API key with `privacy:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) DeletePrivacyExecutorsAdapter(ctx context.Context, adapter string, opts ...RequestOption) (*DeletePrivacyExecutorsAdapterResponse, error) {
	var out DeletePrivacyExecutorsAdapterResponse
	if err := c.do(ctx, "DELETE", "/v1/privacy/executors/"+url.PathEscape(adapter), nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// DeleteProfitabilityProductsProductRef calls Delete a product margin. DELETE /v1/profitability/products/{productRef}.
//
// Removes a product from the margin catalog; its payments fall back to the program or org margin. Audited.
//
// **Auth:** an API key with `analytics:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) DeleteProfitabilityProductsProductRef(ctx context.Context, productRef string, opts ...RequestOption) (*OkStatus, error) {
	var out OkStatus
	if err := c.do(ctx, "DELETE", "/v1/profitability/products/"+url.PathEscape(productRef), nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// DeleteRewardsRulesKey calls Delete a reward rule. DELETE /v1/rewards/rules/{key}.
//
// Deletes the rule. Rewards already granted through it stay.
//
// **Auth:** an API key with `commissions:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) DeleteRewardsRulesKey(ctx context.Context, key string, opts ...RequestOption) (*DeleteRewardsRulesKeyResponse, error) {
	var out DeleteRewardsRulesKeyResponse
	if err := c.do(ctx, "DELETE", "/v1/rewards/rules/"+url.PathEscape(key), nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// DeleteSupportAppsAdapter calls Disconnect a support tool. DELETE /v1/support/apps/{adapter}.
//
// Removes the tool and its secret; its card URL answers 404 from now on.
//
// **Auth:** an API key with `partners:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) DeleteSupportAppsAdapter(ctx context.Context, adapter string, opts ...RequestOption) (*DeleteSupportAppsAdapterResponse, error) {
	var out DeleteSupportAppsAdapterResponse
	if err := c.do(ctx, "DELETE", "/v1/support/apps/"+url.PathEscape(adapter), nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// DeleteWebhookEndpointsID calls Delete a webhook endpoint. DELETE /v1/webhook-endpoints/{id}.
//
// Removes the endpoint and its delivery history. Audited.
//
// **Auth:** an API key with `webhooks:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) DeleteWebhookEndpointsID(ctx context.Context, id string, opts ...RequestOption) (*Ok, error) {
	var out Ok
	if err := c.do(ctx, "DELETE", "/v1/webhook-endpoints/"+url.PathEscape(id), nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// DeleteZapierHooksID calls Unsubscribe a Zapier hook. DELETE /v1/zapier/hooks/{id}.
//
// REST hook unsubscribe, called by Zapier when a Zap turns off.
//
// **Auth:** an API key with `integrations:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) DeleteZapierHooksID(ctx context.Context, id string, opts ...RequestOption) (*Ok, error) {
	var out Ok
	if err := c.do(ctx, "DELETE", "/v1/zapier/hooks/"+url.PathEscape(id), nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetAdminMetrics calls Get platform metrics. GET /v1/admin/metrics.
//
// Platform admins only. Live orgs (and suspended ones), partner accounts (and banned ones), programs, commission earned, payouts sent and open fraud cases across the platform.
//
// **Auth:** a platform admin session.
func (c *Client) GetAdminMetrics(ctx context.Context, opts ...RequestOption) (*GetAdminMetricsResponse, error) {
	var out GetAdminMetricsResponse
	if err := c.do(ctx, "GET", "/v1/admin/metrics", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetAdminOrgsParams are the query parameters of GetAdminOrgs.
type GetAdminOrgsParams struct {
	Q             *string `json:"q,omitempty"`
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetAdminOrgsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "q", p.Q)
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetAdminOrgs calls List organizations. GET /v1/admin/orgs.
//
// Platform admins only. Organizations, searchable by name with `q`.
//
// **Auth:** a platform admin session.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetAdminOrgs(ctx context.Context, params *GetAdminOrgsParams, opts ...RequestOption) (*Page[GetAdminOrgsResponseDataItem], error) {
	var out Page[GetAdminOrgsResponseDataItem]
	if err := c.do(ctx, "GET", "/v1/admin/orgs", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetAdminOrgsEach yields every item of GetAdminOrgs, following the pages.
func (c *Client) GetAdminOrgsEach(ctx context.Context, params *GetAdminOrgsParams, opts ...RequestOption) iter.Seq2[GetAdminOrgsResponseDataItem, error] {
	return paginate(func(cursor *string) (*Page[GetAdminOrgsResponseDataItem], error) {
		p := GetAdminOrgsParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetAdminOrgs(ctx, &p, opts...)
	})
}

// GetAdminPartnerAccountsParams are the query parameters of GetAdminPartnerAccounts.
type GetAdminPartnerAccountsParams struct {
	Flagged       *bool   `json:"flagged,omitempty"`
	Q             *string `json:"q,omitempty"`
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetAdminPartnerAccountsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "flagged", p.Flagged)
	setQuery(v, "q", p.Q)
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetAdminPartnerAccounts calls List partner accounts. GET /v1/admin/partner-accounts.
//
// Platform admins only. Partner accounts with KYC and sanctions status, filtered by `flagged` (banned) and text `q`.
//
// **Auth:** a platform admin session.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetAdminPartnerAccounts(ctx context.Context, params *GetAdminPartnerAccountsParams, opts ...RequestOption) (*Page[GetAdminPartnerAccountsResponseDataItem], error) {
	var out Page[GetAdminPartnerAccountsResponseDataItem]
	if err := c.do(ctx, "GET", "/v1/admin/partner-accounts", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetAdminPartnerAccountsEach yields every item of GetAdminPartnerAccounts, following the pages.
func (c *Client) GetAdminPartnerAccountsEach(ctx context.Context, params *GetAdminPartnerAccountsParams, opts ...RequestOption) iter.Seq2[GetAdminPartnerAccountsResponseDataItem, error] {
	return paginate(func(cursor *string) (*Page[GetAdminPartnerAccountsResponseDataItem], error) {
		p := GetAdminPartnerAccountsParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetAdminPartnerAccounts(ctx, &p, opts...)
	})
}

// GetAgentsParams are the query parameters of GetAgents.
type GetAgentsParams struct {
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetAgentsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetAgents calls List partnered AI agents. GET /v1/agents.
//
// The AI agents partnered with this org, with their wallet, metadata and how many recommendations each has attributed, most active first.
//
// **Auth:** an API key with `growth:read` (or `full`), or a session with role `owner`, `admin`, `analyst`.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetAgents(ctx context.Context, params *GetAgentsParams, opts ...RequestOption) (*Page[GetAgentsResponseDataItem], error) {
	var out Page[GetAgentsResponseDataItem]
	if err := c.do(ctx, "GET", "/v1/agents", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetAgentsEach yields every item of GetAgents, following the pages.
func (c *Client) GetAgentsEach(ctx context.Context, params *GetAgentsParams, opts ...RequestOption) iter.Seq2[GetAgentsResponseDataItem, error] {
	return paginate(func(cursor *string) (*Page[GetAgentsResponseDataItem], error) {
		p := GetAgentsParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetAgents(ctx, &p, opts...)
	})
}

// GetAgentsRecommendationsParams are the query parameters of GetAgentsRecommendations.
type GetAgentsRecommendationsParams struct {
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetAgentsRecommendationsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetAgentsRecommendations calls List agent recommendations. GET /v1/agents/recommendations.
//
// Signed recommendations AI agents have attributed in this org, newest first.
//
// **Auth:** an API key with `growth:read` (or `full`), or a session with role `owner`, `admin`, `analyst`.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetAgentsRecommendations(ctx context.Context, params *GetAgentsRecommendationsParams, opts ...RequestOption) (*Page[GetAgentsRecommendationsResponseDataItem], error) {
	var out Page[GetAgentsRecommendationsResponseDataItem]
	if err := c.do(ctx, "GET", "/v1/agents/recommendations", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetAgentsRecommendationsEach yields every item of GetAgentsRecommendations, following the pages.
func (c *Client) GetAgentsRecommendationsEach(ctx context.Context, params *GetAgentsRecommendationsParams, opts ...RequestOption) iter.Seq2[GetAgentsRecommendationsResponseDataItem, error] {
	return paginate(func(cursor *string) (*Page[GetAgentsRecommendationsResponseDataItem], error) {
		p := GetAgentsRecommendationsParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetAgentsRecommendations(ctx, &p, opts...)
	})
}

// GetAnalyticsByChannelParams are the query parameters of GetAnalyticsByChannel.
type GetAnalyticsByChannelParams struct {
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetAnalyticsByChannelParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetAnalyticsByChannel calls Get revenue by channel. GET /v1/analytics/by-channel.
//
// Referrals and referred revenue by how they were matched (`coupon`, `deterministic` click, `email_hash`, `probabilistic`, `manual`…).
//
// **Auth:** an API key with `analytics:read` (or `full`), or a session of any role.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetAnalyticsByChannel(ctx context.Context, params *GetAnalyticsByChannelParams, opts ...RequestOption) (*Page[GetAnalyticsByChannelResponseDataItem], error) {
	var out Page[GetAnalyticsByChannelResponseDataItem]
	if err := c.do(ctx, "GET", "/v1/analytics/by-channel", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetAnalyticsByChannelEach yields every item of GetAnalyticsByChannel, following the pages.
func (c *Client) GetAnalyticsByChannelEach(ctx context.Context, params *GetAnalyticsByChannelParams, opts ...RequestOption) iter.Seq2[GetAnalyticsByChannelResponseDataItem, error] {
	return paginate(func(cursor *string) (*Page[GetAnalyticsByChannelResponseDataItem], error) {
		p := GetAnalyticsByChannelParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetAnalyticsByChannel(ctx, &p, opts...)
	})
}

// GetAnalyticsByGeoParams are the query parameters of GetAnalyticsByGeo.
type GetAnalyticsByGeoParams struct {
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetAnalyticsByGeoParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetAnalyticsByGeo calls Get revenue by country. GET /v1/analytics/by-geo.
//
// Referrals and referred revenue by the country of the winning click, highest revenue first.
//
// **Auth:** an API key with `analytics:read` (or `full`), or a session of any role.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetAnalyticsByGeo(ctx context.Context, params *GetAnalyticsByGeoParams, opts ...RequestOption) (*Page[GetAnalyticsByGeoResponseDataItem], error) {
	var out Page[GetAnalyticsByGeoResponseDataItem]
	if err := c.do(ctx, "GET", "/v1/analytics/by-geo", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetAnalyticsByGeoEach yields every item of GetAnalyticsByGeo, following the pages.
func (c *Client) GetAnalyticsByGeoEach(ctx context.Context, params *GetAnalyticsByGeoParams, opts ...RequestOption) iter.Seq2[GetAnalyticsByGeoResponseDataItem, error] {
	return paginate(func(cursor *string) (*Page[GetAnalyticsByGeoResponseDataItem], error) {
		p := GetAnalyticsByGeoParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetAnalyticsByGeo(ctx, &p, opts...)
	})
}

// GetAnalyticsCohortLtvParams are the query parameters of GetAnalyticsCohortLtv.
type GetAnalyticsCohortLtvParams struct {
	Months        *int64  `json:"months,omitempty"`
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetAnalyticsCohortLtvParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "months", p.Months)
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetAnalyticsCohortLtv calls Get cohort LTV. GET /v1/analytics/cohort-ltv.
//
// Referred customers grouped by sign-up month over the last `months` (default 6), with net revenue (payments minus refunds) and revenue per customer.
//
// **Auth:** an API key with `analytics:read` (or `full`), or a session of any role.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetAnalyticsCohortLtv(ctx context.Context, params *GetAnalyticsCohortLtvParams, opts ...RequestOption) (*Page[GetAnalyticsCohortLtvResponseDataItem], error) {
	var out Page[GetAnalyticsCohortLtvResponseDataItem]
	if err := c.do(ctx, "GET", "/v1/analytics/cohort-ltv", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetAnalyticsCohortLtvEach yields every item of GetAnalyticsCohortLtv, following the pages.
func (c *Client) GetAnalyticsCohortLtvEach(ctx context.Context, params *GetAnalyticsCohortLtvParams, opts ...RequestOption) iter.Seq2[GetAnalyticsCohortLtvResponseDataItem, error] {
	return paginate(func(cursor *string) (*Page[GetAnalyticsCohortLtvResponseDataItem], error) {
		p := GetAnalyticsCohortLtvParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetAnalyticsCohortLtv(ctx, &p, opts...)
	})
}

// GetAnalyticsCouponPerformanceParams are the query parameters of GetAnalyticsCouponPerformance.
type GetAnalyticsCouponPerformanceParams struct {
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetAnalyticsCouponPerformanceParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetAnalyticsCouponPerformance calls Get coupon performance. GET /v1/analytics/coupon-performance.
//
// Referrals and revenue per coupon code, highest revenue first.
//
// **Auth:** an API key with `analytics:read` (or `full`), or a session of any role.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetAnalyticsCouponPerformance(ctx context.Context, params *GetAnalyticsCouponPerformanceParams, opts ...RequestOption) (*Page[GetAnalyticsCouponPerformanceResponseDataItem], error) {
	var out Page[GetAnalyticsCouponPerformanceResponseDataItem]
	if err := c.do(ctx, "GET", "/v1/analytics/coupon-performance", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetAnalyticsCouponPerformanceEach yields every item of GetAnalyticsCouponPerformance, following the pages.
func (c *Client) GetAnalyticsCouponPerformanceEach(ctx context.Context, params *GetAnalyticsCouponPerformanceParams, opts ...RequestOption) iter.Seq2[GetAnalyticsCouponPerformanceResponseDataItem, error] {
	return paginate(func(cursor *string) (*Page[GetAnalyticsCouponPerformanceResponseDataItem], error) {
		p := GetAnalyticsCouponPerformanceParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetAnalyticsCouponPerformance(ctx, &p, opts...)
	})
}

// GetAnalyticsOverviewParams are the query parameters of GetAnalyticsOverview.
type GetAnalyticsOverviewParams struct {
	ProgramID *string `json:"programId,omitempty"`
	Days      *int64  `json:"days,omitempty"`
}

func (p *GetAnalyticsOverviewParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "programId", p.ProgramID)
	setQuery(v, "days", p.Days)
	return v
}

// GetAnalyticsOverview calls Get the funnel overview. GET /v1/analytics/overview.
//
// Over the last `days` (default 30), optionally for one program: clicks (and how many were filtered), visitors, sign-ups, payments, referred revenue, commission earned and reversed, active partners and the click-to-sign-up and sign-up-to-paid rates.
//
// **Auth:** an API key with `analytics:read` (or `full`), or a session of any role.
func (c *Client) GetAnalyticsOverview(ctx context.Context, params *GetAnalyticsOverviewParams, opts ...RequestOption) (*GetAnalyticsOverviewResponse, error) {
	var out GetAnalyticsOverviewResponse
	if err := c.do(ctx, "GET", "/v1/analytics/overview", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetAnalyticsTimeseriesParams are the query parameters of GetAnalyticsTimeseries.
type GetAnalyticsTimeseriesParams struct {
	Metric        *string `json:"metric,omitempty"`
	Days          *int64  `json:"days,omitempty"`
	ProgramID     *string `json:"programId,omitempty"`
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetAnalyticsTimeseriesParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "metric", p.Metric)
	setQuery(v, "days", p.Days)
	setQuery(v, "programId", p.ProgramID)
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetAnalyticsTimeseries calls Get a daily time series. GET /v1/analytics/timeseries.
//
// One `metric` (`clicks`, `signups` or `revenue` in minor units) per day over the last `days`, optionally for one program.
//
// **Auth:** an API key with `analytics:read` (or `full`), or a session of any role.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetAnalyticsTimeseries(ctx context.Context, params *GetAnalyticsTimeseriesParams, opts ...RequestOption) (*Page[GetAnalyticsTimeseriesResponseDataItem], error) {
	var out Page[GetAnalyticsTimeseriesResponseDataItem]
	if err := c.do(ctx, "GET", "/v1/analytics/timeseries", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetAnalyticsTimeseriesEach yields every item of GetAnalyticsTimeseries, following the pages.
func (c *Client) GetAnalyticsTimeseriesEach(ctx context.Context, params *GetAnalyticsTimeseriesParams, opts ...RequestOption) iter.Seq2[GetAnalyticsTimeseriesResponseDataItem, error] {
	return paginate(func(cursor *string) (*Page[GetAnalyticsTimeseriesResponseDataItem], error) {
		p := GetAnalyticsTimeseriesParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetAnalyticsTimeseries(ctx, &p, opts...)
	})
}

// GetAPIKeysParams are the query parameters of GetAPIKeys.
type GetAPIKeysParams struct {
	Mode          *string `json:"mode,omitempty"`
	Owner         *string `json:"owner,omitempty"`
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetAPIKeysParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "mode", p.Mode)
	setQuery(v, "owner", p.Owner)
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetAPIKeys calls List API keys. GET /v1/api-keys.
//
// Every key of the org, live and test (a test-mode caller sees test keys only), oldest first, with scopes, pinned version, owner, expiry and last use. Never returns the secret. Filter by `mode` or `owner`.
//
// **Auth:** an API key with `team:read` (or `full`), or a session of any role.
//
// mode: Only live or only test keys; default both.
// owner: Only keys with this owner.
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetAPIKeys(ctx context.Context, params *GetAPIKeysParams, opts ...RequestOption) (*Page[ApiKey], error) {
	var out Page[ApiKey]
	if err := c.do(ctx, "GET", "/v1/api-keys", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetAPIKeysEach yields every item of GetAPIKeys, following the pages.
func (c *Client) GetAPIKeysEach(ctx context.Context, params *GetAPIKeysParams, opts ...RequestOption) iter.Seq2[ApiKey, error] {
	return paginate(func(cursor *string) (*Page[ApiKey], error) {
		p := GetAPIKeysParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetAPIKeys(ctx, &p, opts...)
	})
}

// GetAPIKeysID calls Get an API key. GET /v1/api-keys/{id}.
//
// One key of the org, without its secret.
//
// **Auth:** an API key with `team:read` (or `full`), or a session of any role.
func (c *Client) GetAPIKeysID(ctx context.Context, id string, opts ...RequestOption) (*ApiKey, error) {
	var out ApiKey
	if err := c.do(ctx, "GET", "/v1/api-keys/"+url.PathEscape(id), nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetAssetsParams are the query parameters of GetAssets.
type GetAssetsParams struct {
	ProgramID     *string `json:"programId,omitempty"`
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetAssetsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "programId", p.ProgramID)
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetAssets calls List marketing assets. GET /v1/assets.
//
// Marketing assets, filtered by `programId`.
//
// **Auth:** an API key with `partners:read` (or `full`), or a session of any role.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetAssets(ctx context.Context, params *GetAssetsParams, opts ...RequestOption) (*Page[GetAssetsResponseDataItem], error) {
	var out Page[GetAssetsResponseDataItem]
	if err := c.do(ctx, "GET", "/v1/assets", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetAssetsEach yields every item of GetAssets, following the pages.
func (c *Client) GetAssetsEach(ctx context.Context, params *GetAssetsParams, opts ...RequestOption) iter.Seq2[GetAssetsResponseDataItem, error] {
	return paginate(func(cursor *string) (*Page[GetAssetsResponseDataItem], error) {
		p := GetAssetsParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetAssets(ctx, &p, opts...)
	})
}

// GetAssistantSuggestionsParams are the query parameters of GetAssistantSuggestions.
type GetAssistantSuggestionsParams struct {
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetAssistantSuggestionsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetAssistantSuggestions calls Get program suggestions. GET /v1/assistant/suggestions.
//
// Draft-only suggestions for running the program, each with a suggested action for a person to take: partners who never drove a click, links with clicks but no referrals, programs without a commission rule, the review queue, and credit ROI findings. It never changes anything.
//
// **Auth:** an API key with `analytics:read` (or `full`), or a session of any role.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetAssistantSuggestions(ctx context.Context, params *GetAssistantSuggestionsParams, opts ...RequestOption) (*Page[GetAssistantSuggestionsResponseDataItem], error) {
	var out Page[GetAssistantSuggestionsResponseDataItem]
	if err := c.do(ctx, "GET", "/v1/assistant/suggestions", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetAssistantSuggestionsEach yields every item of GetAssistantSuggestions, following the pages.
func (c *Client) GetAssistantSuggestionsEach(ctx context.Context, params *GetAssistantSuggestionsParams, opts ...RequestOption) iter.Seq2[GetAssistantSuggestionsResponseDataItem, error] {
	return paginate(func(cursor *string) (*Page[GetAssistantSuggestionsResponseDataItem], error) {
		p := GetAssistantSuggestionsParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetAssistantSuggestions(ctx, &p, opts...)
	})
}

// GetAuditLogsParams are the query parameters of GetAuditLogs.
type GetAuditLogsParams struct {
	Actor         *string `json:"actor,omitempty"`
	ActorType     *string `json:"actorType,omitempty"`
	Action        *string `json:"action,omitempty"`
	SubjectType   *string `json:"subjectType,omitempty"`
	SubjectID     *string `json:"subjectId,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
}

func (p *GetAuditLogsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "actor", p.Actor)
	setQuery(v, "actorType", p.ActorType)
	setQuery(v, "action", p.Action)
	setQuery(v, "subjectType", p.SubjectType)
	setQuery(v, "subjectId", p.SubjectID)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	return v
}

// GetAuditLogs calls List audit log entries. GET /v1/audit-logs.
//
// Every change recorded in the org, newest first: who did it (`actor`, `actorType`), what (`action`; a trailing `*` matches a prefix such as `payout.*`) and to what (`subjectType`, `subjectId`). Filter by any of those and by `created[gte]`/`created[lte]`.
//
// **Auth:** an API key with `audit:read` (or `full`), or a session with role `owner`, `admin`.
//
// actor: Actor id: a user, API key or connection id.
// action: Exact action, or a prefix ending in * (e.g. payout.*).
// starting_after: An entry id: return the entries after it.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetAuditLogs(ctx context.Context, params *GetAuditLogsParams, opts ...RequestOption) (*Page[AuditLogEntry], error) {
	var out Page[AuditLogEntry]
	if err := c.do(ctx, "GET", "/v1/audit-logs", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetAuditLogsEach yields every item of GetAuditLogs, following the pages.
func (c *Client) GetAuditLogsEach(ctx context.Context, params *GetAuditLogsParams, opts ...RequestOption) iter.Seq2[AuditLogEntry, error] {
	return paginate(func(cursor *string) (*Page[AuditLogEntry], error) {
		p := GetAuditLogsParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetAuditLogs(ctx, &p, opts...)
	})
}

// GetAuthConnectionsParams are the query parameters of GetAuthConnections.
type GetAuthConnectionsParams struct {
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetAuthConnectionsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetAuthConnections calls List auth connections. GET /v1/auth/connections.
//
// Auth-platform connections with provider, program, status, signature version and rotation state. Secrets are never returned.
//
// **Auth:** an API key with `integrations:read` (or `full`), or a session of any role.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetAuthConnections(ctx context.Context, params *GetAuthConnectionsParams, opts ...RequestOption) (*Page[AuthConnection], error) {
	var out Page[AuthConnection]
	if err := c.do(ctx, "GET", "/v1/auth/connections", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetAuthConnectionsEach yields every item of GetAuthConnections, following the pages.
func (c *Client) GetAuthConnectionsEach(ctx context.Context, params *GetAuthConnectionsParams, opts ...RequestOption) iter.Seq2[AuthConnection, error] {
	return paginate(func(cursor *string) (*Page[AuthConnection], error) {
		p := GetAuthConnectionsParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetAuthConnections(ctx, &p, opts...)
	})
}

// GetAuthConnectionsIDExportParams are the query parameters of GetAuthConnectionsIDExport.
type GetAuthConnectionsIDExportParams struct {
	CustomerRef *string `json:"customerRef,omitempty"`
}

func (p *GetAuthConnectionsIDExportParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "customerRef", p.CustomerRef)
	return v
}

// GetAuthConnectionsIDExport calls Export a user (signed protocol). GET /v1/auth/connections/{id}/export.
//
// An access request the platform forwards: everything filed under the customer and their partner record. Signed with the connection secret in `x-referalli-signature: t=<unix>,v1=…,v2=…` (v2 also signs the request line; once a connection has seen v2 it requires it), with a unique `x-referalli-id` per call. The legacy header names from before the rename are also accepted, for existing integrations. Calls older than 5 minutes, reused ids (`409 replay`), paused connections (`423 connection_paused`) and unknown fields (`422`) are refused.
//
// **Auth:** a signature (see above); no API key.
func (c *Client) GetAuthConnectionsIDExport(ctx context.Context, id string, params *GetAuthConnectionsIDExportParams, opts ...RequestOption) (*GetAuthConnectionsIdExportResponse, error) {
	var out GetAuthConnectionsIdExportResponse
	if err := c.do(ctx, "GET", "/v1/auth/connections/"+url.PathEscape(id)+"/export", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetAuthConnectionsProviders calls List identity connectors. GET /v1/auth/connections/providers.
//
// Every auth provider you can connect (Atlas, Clerk, Auth0, Supabase, Firebase, Cognito, Stytch, Kinde, WorkOS, generic, your own auth), with how it connects, who issues the secret, the setup steps, a copyable template where the vendor has no native webhook, and where the referral metadata lands.
//
// **Auth:** an API key with `integrations:read` (or `full`), or a session of any role.
func (c *Client) GetAuthConnectionsProviders(ctx context.Context, opts ...RequestOption) (*GetAuthConnectionsProvidersResponse, error) {
	var out GetAuthConnectionsProvidersResponse
	if err := c.do(ctx, "GET", "/v1/auth/connections/providers", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetAuthProviders calls List sign-in providers. GET /v1/auth/providers.
//
// The identity providers configured on this deployment, for rendering sign-in buttons.
//
// **Auth:** none.
func (c *Client) GetAuthProviders(ctx context.Context, opts ...RequestOption) (*GetAuthProvidersResponse, error) {
	var out GetAuthProvidersResponse
	if err := c.do(ctx, "GET", "/v1/auth/providers", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetBillingConnectionsParams are the query parameters of GetBillingConnections.
type GetBillingConnectionsParams struct {
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetBillingConnectionsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetBillingConnections calls List billing connections. GET /v1/billing/connections.
//
// Billing connections (deleted ones excluded), with status, whether credentials are needed and set, the value mapping and secret-rotation state.
//
// **Auth:** an API key with `billing:read` (or `full`), or a session of any role.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetBillingConnections(ctx context.Context, params *GetBillingConnectionsParams, opts ...RequestOption) (*Page[BillingConnection], error) {
	var out Page[BillingConnection]
	if err := c.do(ctx, "GET", "/v1/billing/connections", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetBillingConnectionsEach yields every item of GetBillingConnections, following the pages.
func (c *Client) GetBillingConnectionsEach(ctx context.Context, params *GetBillingConnectionsParams, opts ...RequestOption) iter.Seq2[BillingConnection, error] {
	return paginate(func(cursor *string) (*Page[BillingConnection], error) {
		p := GetBillingConnectionsParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetBillingConnections(ctx, &p, opts...)
	})
}

// GetBillingConnectionsID calls Get a billing connection. GET /v1/billing/connections/{id}.
//
// One billing connection, without its secrets.
//
// **Auth:** an API key with `billing:read` (or `full`), or a session of any role.
func (c *Client) GetBillingConnectionsID(ctx context.Context, id string, opts ...RequestOption) (*BillingConnection, error) {
	var out BillingConnection
	if err := c.do(ctx, "GET", "/v1/billing/connections/"+url.PathEscape(id), nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetBillingEventsParams are the query parameters of GetBillingEvents.
type GetBillingEventsParams struct {
	CustomerRef   *string `json:"customerRef,omitempty"`
	Status        *string `json:"status,omitempty"`
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetBillingEventsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "customerRef", p.CustomerRef)
	setQuery(v, "status", p.Status)
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetBillingEvents calls List billing events. GET /v1/billing/events.
//
// Payments, refunds and other billing events received from every connection, newest first. Filter with `customerRef`, or `status=ignored_by_filter` for the events a product filter or the internal-customer list set aside (recorded, never processed).
//
// **Auth:** an API key with `billing:read` (or `full`), or a session of any role.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetBillingEvents(ctx context.Context, params *GetBillingEventsParams, opts ...RequestOption) (*Page[BillingEvent], error) {
	var out Page[BillingEvent]
	if err := c.do(ctx, "GET", "/v1/billing/events", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetBillingEventsEach yields every item of GetBillingEvents, following the pages.
func (c *Client) GetBillingEventsEach(ctx context.Context, params *GetBillingEventsParams, opts ...RequestOption) iter.Seq2[BillingEvent, error] {
	return paginate(func(cursor *string) (*Page[BillingEvent], error) {
		p := GetBillingEventsParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetBillingEvents(ctx, &p, opts...)
	})
}

// GetBillingInternalCustomers calls List internal customers. GET /v1/billing/internal-customers.
//
// The org's own customers (staff, test accounts, your own plans) whose payments never earn on any connection. Metadata `planSource: "internal"` on the provider's customer does the same without a list.
//
// **Auth:** an API key with `billing:read` (or `full`), or a session of any role.
func (c *Client) GetBillingInternalCustomers(ctx context.Context, opts ...RequestOption) ([]InternalCustomer, error) {
	var out []InternalCustomer
	if err := c.do(ctx, "GET", "/v1/billing/internal-customers", nil, nil, &out, "json", opts); err != nil {
		return out, err
	}
	return out, nil
}

// GetBillingPlan calls Get the plan. GET /v1/billing/plan.
//
// The org's Referalli plan and status, its limits and current usage (programs, partners), the plan catalog, whether there is a subscription to manage and whether this caller may activate a plan without checkout.
//
// **Auth:** an API key with `billing:read` (or `full`), or a session of any role.
func (c *Client) GetBillingPlan(ctx context.Context, opts ...RequestOption) (*GetBillingPlanResponse, error) {
	var out GetBillingPlanResponse
	if err := c.do(ctx, "GET", "/v1/billing/plan", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetBranding calls Get branding. GET /v1/branding.
//
// The org's logo, theme tokens, sender address, and its custom portal and link domains with their status (`none`, `pending` DNS, `provisioning` a certificate, `active`), and `warnings` for theme colour pairs below WCAG AA contrast.
//
// **Auth:** an API key with `org:read` (or `full`), or a session of any role.
func (c *Client) GetBranding(ctx context.Context, opts ...RequestOption) (*GetBrandingResponse, error) {
	var out GetBrandingResponse
	if err := c.do(ctx, "GET", "/v1/branding", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetChatlyIdentity calls Get the support-chat identity. GET /v1/chatly/identity.
//
// The signed-in merchant user or partner as a Chatly widget identity: `externalId` (`merchant:<id>` or `partner:<id>`), email, name and `userHash`, the HMAC-SHA256 of `externalId` with Referalli's identity secret. `identity` is null when signed out or when identity verification isn't configured. Never cached.
//
// **Auth:** none.
func (c *Client) GetChatlyIdentity(ctx context.Context, opts ...RequestOption) (*GetChatlyIdentityResponse, error) {
	var out GetChatlyIdentityResponse
	if err := c.do(ctx, "GET", "/v1/chatly/identity", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetCleanroomConsentParams are the query parameters of GetCleanroomConsent.
type GetCleanroomConsentParams struct {
	CustomerRef *string `json:"customerRef,omitempty"`
}

func (p *GetCleanroomConsentParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "customerRef", p.CustomerRef)
	return v
}

// GetCleanroomConsent calls Get a customer's consent. GET /v1/cleanroom/consent.
//
// A customer's recorded consent: whether granted, its purposes, the attributes they declared and when it changed. `consent: null` when none is recorded.
//
// **Auth:** an API key with `growth:read` (or `full`), or a session of any role.
func (c *Client) GetCleanroomConsent(ctx context.Context, params *GetCleanroomConsentParams, opts ...RequestOption) (*GetCleanroomConsentResponse, error) {
	var out GetCleanroomConsentResponse
	if err := c.do(ctx, "GET", "/v1/cleanroom/consent", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetCleanroomQueryParams are the query parameters of GetCleanroomQuery.
type GetCleanroomQueryParams struct {
	Dimension *string `json:"dimension,omitempty"`
}

func (p *GetCleanroomQueryParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "dimension", p.Dimension)
	return v
}

// GetCleanroomQuery calls Query the clean room. GET /v1/cleanroom/query.
//
// Aggregates over consenting customers only, grouped by a zero-party attribute they declared (`dimension`): customers, converted customers and referred revenue per value. Cohorts smaller than 5 are suppressed (k-anonymity); individual rows never leave.
//
// **Auth:** an API key with `growth:read` (or `full`), or a session of any role.
func (c *Client) GetCleanroomQuery(ctx context.Context, params *GetCleanroomQueryParams, opts ...RequestOption) (*GetCleanroomQueryResponse, error) {
	var out GetCleanroomQueryResponse
	if err := c.do(ctx, "GET", "/v1/cleanroom/query", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetClicksConsentCoverageParams are the query parameters of GetClicksConsentCoverage.
type GetClicksConsentCoverageParams struct {
	Days *int64 `json:"days,omitempty"`
}

func (p *GetClicksConsentCoverageParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "days", p.Days)
	return v
}

// GetClicksConsentCoverage calls Get consent coverage. GET /v1/clicks/consent-coverage.
//
// The share of attributed sign-ups in the last `days` (default 30) whose referral carries consent evidence, on its click or on the `associate` call. Low coverage means sign-ups are attributed from tags without a consent setting (`none`) or from server-side calls that don't pass `consent`.
//
// **Auth:** an API key with `links:read` (or `full`), or a session of any role.
//
// days: Look-back window in days.
func (c *Client) GetClicksConsentCoverage(ctx context.Context, params *GetClicksConsentCoverageParams, opts ...RequestOption) (*GetClicksConsentCoverageResponse, error) {
	var out GetClicksConsentCoverageResponse
	if err := c.do(ctx, "GET", "/v1/clicks/consent-coverage", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetClicksID calls Get a click. GET /v1/clicks/{id}.
//
// One click with its consent evidence: what the visitor's CMP said when the tracker stored the referral (`consent.provider`, `state`, `stamp`, `tcString`…). A redirect click (`/go`) is recorded before the landing page reads the CMP, so its evidence comes from the first consented visit beacon naming it. `consent` is null when none arrived.
//
// **Auth:** an API key with `links:read` (or `full`), or a session of any role.
func (c *Client) GetClicksID(ctx context.Context, id string, opts ...RequestOption) (*GetClicksIdResponse, error) {
	var out GetClicksIdResponse
	if err := c.do(ctx, "GET", "/v1/clicks/"+url.PathEscape(id), nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetCliDeviceCodesUserCode calls Get a pending CLI sign-in. GET /v1/cli/device-codes/{userCode}.
//
// For the approval page: the pending code, which CLI asked and when it expires. Needs a signed-in person (owner or admin). 404 when expired or used.
//
// **Auth:** an API key with `team:read` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) GetCliDeviceCodesUserCode(ctx context.Context, userCode string, opts ...RequestOption) (*GetCliDeviceCodesUserCodeResponse, error) {
	var out GetCliDeviceCodesUserCodeResponse
	if err := c.do(ctx, "GET", "/v1/cli/device-codes/"+url.PathEscape(userCode), nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetCliListenersIDEventsParams are the query parameters of GetCliListenersIDEvents.
type GetCliListenersIDEventsParams struct {
	Wait *int64 `json:"wait,omitempty"`
}

func (p *GetCliListenersIDEventsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "wait", p.Wait)
	return v
}

// GetCliListenersIDEvents calls Receive listener events. GET /v1/cli/listeners/{id}/events.
//
// Up to 50 queued events for the listener, oldest first, each with its body and signed headers. `wait` long-polls up to 25 seconds for the first event. Polling keeps the listener alive.
//
// **Auth:** an API key with `webhooks:read` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) GetCliListenersIDEvents(ctx context.Context, id string, params *GetCliListenersIDEventsParams, opts ...RequestOption) (*GetCliListenersIdEventsResponse, error) {
	var out GetCliListenersIdEventsResponse
	if err := c.do(ctx, "GET", "/v1/cli/listeners/"+url.PathEscape(id)+"/events", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetCliLogsParams are the query parameters of GetCliLogs.
type GetCliLogsParams struct {
	After *int64 `json:"after,omitempty"`
	Wait  *int64 `json:"wait,omitempty"`
	Limit *int64 `json:"limit,omitempty"`
}

func (p *GetCliLogsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "after", p.After)
	setQuery(v, "wait", p.Wait)
	setQuery(v, "limit", p.Limit)
	return v
}

// GetCliLogs calls Tail the request log. GET /v1/cli/logs.
//
// The CLI's `logs tail` command: the org's recent API requests (method, path without query string, status, latency, actor). The first call returns the latest 20 and a `cursor`; pass it as `after` (with `wait` to long-poll up to 25 seconds) for newer entries. Kept for a day.
//
// **Auth:** an API key with `audit:read` (or `full`), or a session with role `owner`, `admin`.
//
// after: Return entries after this cursor.
func (c *Client) GetCliLogs(ctx context.Context, params *GetCliLogsParams, opts ...RequestOption) (*GetCliLogsResponse, error) {
	var out GetCliLogsResponse
	if err := c.do(ctx, "GET", "/v1/cli/logs", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetCliTopics calls List triggerable topics. GET /v1/cli/topics.
//
// The event topics `POST /v1/cli/trigger` can send.
//
// **Auth:** an API key with `webhooks:read` (or `full`), or a session of any role.
func (c *Client) GetCliTopics(ctx context.Context, opts ...RequestOption) (*GetCliTopicsResponse, error) {
	var out GetCliTopicsResponse
	if err := c.do(ctx, "GET", "/v1/cli/topics", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetCommissionRulesParams are the query parameters of GetCommissionRules.
type GetCommissionRulesParams struct {
	ProgramID     *string `json:"programId,omitempty"`
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetCommissionRulesParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "programId", p.ProgramID)
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetCommissionRules calls List commission rules. GET /v1/commission-rules.
//
// Every rule version, by program and newest version first. Filter with `programId`.
//
// **Auth:** an API key with `commissions:read` (or `full`), or a session of any role.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetCommissionRules(ctx context.Context, params *GetCommissionRulesParams, opts ...RequestOption) (*Page[CommissionRule], error) {
	var out Page[CommissionRule]
	if err := c.do(ctx, "GET", "/v1/commission-rules", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetCommissionRulesEach yields every item of GetCommissionRules, following the pages.
func (c *Client) GetCommissionRulesEach(ctx context.Context, params *GetCommissionRulesParams, opts ...RequestOption) iter.Seq2[CommissionRule, error] {
	return paginate(func(cursor *string) (*Page[CommissionRule], error) {
		p := GetCommissionRulesParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetCommissionRules(ctx, &p, opts...)
	})
}

// GetCommissionsParams are the query parameters of GetCommissions.
type GetCommissionsParams struct {
	PartnerID     *string `json:"partnerId,omitempty"`
	ProgramID     *string `json:"programId,omitempty"`
	Type          *string `json:"type,omitempty"`
	ReferralID    *string `json:"referralId,omitempty"`
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetCommissionsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "partnerId", p.PartnerID)
	setQuery(v, "programId", p.ProgramID)
	setQuery(v, "type", p.Type)
	setQuery(v, "referralId", p.ReferralID)
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetCommissions calls List commission ledger events. GET /v1/commissions.
//
// The append-only commission ledger, newest first: `earned` lines and the `approved`, `paid`, `reversed` and `adjusted` events recorded against them (`relatedCommissionId`). Filter by `partnerId`, `programId`, `type` or `referralId`.
//
// **Auth:** an API key with `commissions:read` (or `full`), or a session of any role.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetCommissions(ctx context.Context, params *GetCommissionsParams, opts ...RequestOption) (*Page[CommissionEvent], error) {
	var out Page[CommissionEvent]
	if err := c.do(ctx, "GET", "/v1/commissions", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetCommissionsEach yields every item of GetCommissions, following the pages.
func (c *Client) GetCommissionsEach(ctx context.Context, params *GetCommissionsParams, opts ...RequestOption) iter.Seq2[CommissionEvent, error] {
	return paginate(func(cursor *string) (*Page[CommissionEvent], error) {
		p := GetCommissionsParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetCommissions(ctx, &p, opts...)
	})
}

// GetCommissionsID calls Get a commission. GET /v1/commissions/{id}.
//
// One ledger event. An `earned` line also carries its billing event, its net after reversals and adjustments, its state (`pending`, `approved` or `paid`) and every event recorded against it.
//
// **Auth:** an API key with `commissions:read` (or `full`), or a session of any role.
func (c *Client) GetCommissionsID(ctx context.Context, id string, opts ...RequestOption) (*CommissionLine, error) {
	var out CommissionLine
	if err := c.do(ctx, "GET", "/v1/commissions/"+url.PathEscape(id), nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetContractsParams are the query parameters of GetContracts.
type GetContractsParams struct {
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetContractsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetContracts calls List partner contracts. GET /v1/contracts.
//
// The org's contracts and their status (`draft`, `sent`, `accepted`, `declined`, `void`).
//
// **Auth:** an API key with `partners:read` (or `full`), or a session with role `owner`, `admin`.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetContracts(ctx context.Context, params *GetContractsParams, opts ...RequestOption) (*Page[GetContractsResponseDataItem], error) {
	var out Page[GetContractsResponseDataItem]
	if err := c.do(ctx, "GET", "/v1/contracts", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetContractsEach yields every item of GetContracts, following the pages.
func (c *Client) GetContractsEach(ctx context.Context, params *GetContractsParams, opts ...RequestOption) iter.Seq2[GetContractsResponseDataItem, error] {
	return paginate(func(cursor *string) (*Page[GetContractsResponseDataItem], error) {
		p := GetContractsParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetContracts(ctx, &p, opts...)
	})
}

// GetCouponsParams are the query parameters of GetCoupons.
type GetCouponsParams struct {
	ProgramID     *string `json:"programId,omitempty"`
	PartnerID     *string `json:"partnerId,omitempty"`
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetCouponsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "programId", p.ProgramID)
	setQuery(v, "partnerId", p.PartnerID)
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetCoupons calls List coupon codes. GET /v1/coupons.
//
// Coupon codes, filtered by `programId` and/or `partnerId`.
//
// **Auth:** an API key with `links:read` (or `full`), or a session of any role.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetCoupons(ctx context.Context, params *GetCouponsParams, opts ...RequestOption) (*Page[Coupon], error) {
	var out Page[Coupon]
	if err := c.do(ctx, "GET", "/v1/coupons", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetCouponsEach yields every item of GetCoupons, following the pages.
func (c *Client) GetCouponsEach(ctx context.Context, params *GetCouponsParams, opts ...RequestOption) iter.Seq2[Coupon, error] {
	return paginate(func(cursor *string) (*Page[Coupon], error) {
		p := GetCouponsParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetCoupons(ctx, &p, opts...)
	})
}

// GetCreatorsParams are the query parameters of GetCreators.
type GetCreatorsParams struct {
	Q                *string `json:"q,omitempty"`
	Category         *string `json:"category,omitempty"`
	Channel          *string `json:"channel,omitempty"`
	MinAudience      *int64  `json:"minAudience,omitempty"`
	MinEngagementBps *int64  `json:"minEngagementBps,omitempty"`
	AudienceCountry  *string `json:"audienceCountry,omitempty"`
	Limit            *int64  `json:"limit,omitempty"`
	StartingAfter    *string `json:"starting_after,omitempty"`
	EndingBefore     *string `json:"ending_before,omitempty"`
	CreatedGte       *string `json:"created[gte],omitempty"`
	CreatedLte       *string `json:"created[lte],omitempty"`
}

func (p *GetCreatorsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "q", p.Q)
	setQuery(v, "category", p.Category)
	setQuery(v, "channel", p.Channel)
	setQuery(v, "minAudience", p.MinAudience)
	setQuery(v, "minEngagementBps", p.MinEngagementBps)
	setQuery(v, "audienceCountry", p.AudienceCountry)
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetCreators calls Search creators. GET /v1/creators.
//
// Published creator profiles, most engaged first, filtered by text (`q`), `category`, `channel`, minimum audience, minimum engagement and audience country. Each carries verified stats and whether this org already invited them. Emails are never returned.
//
// **Auth:** an API key with `growth:read` (or `full`), or a session with role `owner`, `admin`.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetCreators(ctx context.Context, params *GetCreatorsParams, opts ...RequestOption) (*Page[CreatorProfile], error) {
	var out Page[CreatorProfile]
	if err := c.do(ctx, "GET", "/v1/creators", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetCreatorsEach yields every item of GetCreators, following the pages.
func (c *Client) GetCreatorsEach(ctx context.Context, params *GetCreatorsParams, opts ...RequestOption) iter.Seq2[CreatorProfile, error] {
	return paginate(func(cursor *string) (*Page[CreatorProfile], error) {
		p := GetCreatorsParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetCreators(ctx, &p, opts...)
	})
}

// GetCreatorsAccountID calls Get a creator. GET /v1/creators/{accountId}.
//
// One published creator profile with verified stats. 404 when not listed.
//
// **Auth:** an API key with `growth:read` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) GetCreatorsAccountID(ctx context.Context, accountId string, opts ...RequestOption) (*CreatorProfile, error) {
	var out CreatorProfile
	if err := c.do(ctx, "GET", "/v1/creators/"+url.PathEscape(accountId), nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetDealsParams are the query parameters of GetDeals.
type GetDealsParams struct {
	Status        *string `json:"status,omitempty"`
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetDealsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "status", p.Status)
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetDeals calls List registered deals. GET /v1/deals.
//
// Deals partners registered, newest first. Filter with `status` (`submitted`, `approved`, `rejected`, `won`, `lost`).
//
// **Auth:** an API key with `growth:read` (or `full`), or a session of any role.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetDeals(ctx context.Context, params *GetDealsParams, opts ...RequestOption) (*Page[GetDealsResponseDataItem], error) {
	var out Page[GetDealsResponseDataItem]
	if err := c.do(ctx, "GET", "/v1/deals", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetDealsEach yields every item of GetDeals, following the pages.
func (c *Client) GetDealsEach(ctx context.Context, params *GetDealsParams, opts ...RequestOption) iter.Seq2[GetDealsResponseDataItem, error] {
	return paginate(func(cursor *string) (*Page[GetDealsResponseDataItem], error) {
		p := GetDealsParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetDeals(ctx, &p, opts...)
	})
}

// GetDisputesParams are the query parameters of GetDisputes.
type GetDisputesParams struct {
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetDisputesParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetDisputes calls List partner disputes. GET /v1/disputes.
//
// Disputes partners filed with this org, newest first.
//
// **Auth:** an API key with `commissions:read` (or `full`), or a session of any role.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetDisputes(ctx context.Context, params *GetDisputesParams, opts ...RequestOption) (*Page[GetDisputesResponseDataItem], error) {
	var out Page[GetDisputesResponseDataItem]
	if err := c.do(ctx, "GET", "/v1/disputes", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetDisputesEach yields every item of GetDisputes, following the pages.
func (c *Client) GetDisputesEach(ctx context.Context, params *GetDisputesParams, opts ...RequestOption) iter.Seq2[GetDisputesResponseDataItem, error] {
	return paginate(func(cursor *string) (*Page[GetDisputesResponseDataItem], error) {
		p := GetDisputesParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetDisputes(ctx, &p, opts...)
	})
}

// GetEmailSettings calls Get email settings. GET /v1/email/settings.
//
// The org's own email provider (`resend`, `brevo` or `postmark`), sender identity, branding and sending-domain status. The API key is never returned.
//
// **Auth:** an API key with `org:read` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) GetEmailSettings(ctx context.Context, opts ...RequestOption) (*EmailSettings, error) {
	var out EmailSettings
	if err := c.do(ctx, "GET", "/v1/email/settings", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetEmailTemplatesParams are the query parameters of GetEmailTemplates.
type GetEmailTemplatesParams struct {
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetEmailTemplatesParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetEmailTemplates calls List email templates. GET /v1/email/templates.
//
// Every email event with its placeholders and, per locale, the default template and the org's override if any.
//
// **Auth:** an API key with `org:read` (or `full`), or a session with role `owner`, `admin`.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetEmailTemplates(ctx context.Context, params *GetEmailTemplatesParams, opts ...RequestOption) (*Page[EmailTemplateEvent], error) {
	var out Page[EmailTemplateEvent]
	if err := c.do(ctx, "GET", "/v1/email/templates", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetEmailTemplatesEach yields every item of GetEmailTemplates, following the pages.
func (c *Client) GetEmailTemplatesEach(ctx context.Context, params *GetEmailTemplatesParams, opts ...RequestOption) iter.Seq2[EmailTemplateEvent, error] {
	return paginate(func(cursor *string) (*Page[EmailTemplateEvent], error) {
		p := GetEmailTemplatesParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetEmailTemplates(ctx, &p, opts...)
	})
}

// GetEmbedParams are the query parameters of GetEmbed.
type GetEmbedParams struct {
	Token *string `json:"token,omitempty"`
}

func (p *GetEmbedParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "token", p.Token)
	return v
}

// GetEmbed calls Get one org's embed manifest. GET /v1/embed.
//
// The embed manifest for the org that owns `token` (its public tracking token): its link domain in the CSP origins, its longest program cookie window as the storage lifetime, and its consent default. Public, CORS `*`.
//
// **Auth:** none.
//
// token: The org's public tracking token.
func (c *Client) GetEmbed(ctx context.Context, params *GetEmbedParams, opts ...RequestOption) (*GetWellKnownForgedeckEmbedJsonResponse, error) {
	var out GetWellKnownForgedeckEmbedJsonResponse
	if err := c.do(ctx, "GET", "/v1/embed", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetEmbedSummary calls Read embedded partner stats. GET /v1/embed/summary.
//
// Called by the embed component with `Authorization: Bearer <embed token>`: the partner's name, eligible clicks, active referrals, balances and link codes. 401 for a missing, invalid or expired token.
//
// **Auth:** none.
func (c *Client) GetEmbedSummary(ctx context.Context, opts ...RequestOption) (*GetEmbedSummaryResponse, error) {
	var out GetEmbedSummaryResponse
	if err := c.do(ctx, "GET", "/v1/embed/summary", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetEventsParams are the query parameters of GetEvents.
type GetEventsParams struct {
	Since  *string `json:"since,omitempty"`
	Topics *string `json:"topics,omitempty"`
	After  *string `json:"after,omitempty"`
	Limit  *int64  `json:"limit,omitempty"`
}

func (p *GetEventsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "since", p.Since)
	setQuery(v, "topics", p.Topics)
	setQuery(v, "after", p.After)
	setQuery(v, "limit", p.Limit)
	return v
}

// GetEvents calls List past events. GET /v1/events.
//
// The org's events, oldest first, in the shape webhook deliveries carry: filter with `since` and `topics`, page with `after` (the previous page's `nextCursor`). Every event is listed whether or not an endpoint subscribed to it. `referalli events tail --since 1h` uses it to backfill before tailing live.
//
// **Auth:** an API key with `webhooks:read` (or `full`), or a session of any role.
//
// since: Only events created at or after this time.
// topics: Comma-separated topics (default: every topic).
// after: The `nextCursor` of the previous page: events after this one.
func (c *Client) GetEvents(ctx context.Context, params *GetEventsParams, opts ...RequestOption) (*GetEventsResponse, error) {
	var out GetEventsResponse
	if err := c.do(ctx, "GET", "/v1/events", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetExportAccountingParams are the query parameters of GetExportAccounting.
type GetExportAccountingParams struct {
	Format *string `json:"format,omitempty"`
	Year   *int64  `json:"year,omitempty"`
}

func (p *GetExportAccountingParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "format", p.Format)
	setQuery(v, "year", p.Year)
	return v
}

// GetExportAccounting calls Export payouts for accounting. GET /v1/export/accounting.
//
// A year's sent payouts as a CSV that QuickBooks (`format=quickbooks`, the default) or Xero imports.
//
// **Auth:** an API key with `payouts:read` (or `full`), or a session with role `owner`, `admin`, `finance`.
func (c *Client) GetExportAccounting(ctx context.Context, params *GetExportAccountingParams, opts ...RequestOption) (string, error) {
	var out string
	err := c.do(ctx, "GET", "/v1/export/accounting", params.values(), nil, &out, "text", opts)
	return out, err
}

// GetExportFull calls Export all org data. GET /v1/export/full.
//
// Every tenant table of the org as NDJSON (one JSON object per line, tagged with its table), streamed so any size works. Owner or admin.
//
// **Auth:** an API key with `analytics:read` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) GetExportFull(ctx context.Context, opts ...RequestOption) (string, error) {
	var out string
	err := c.do(ctx, "GET", "/v1/export/full", nil, nil, &out, "text", opts)
	return out, err
}

// GetFraudCasesParams are the query parameters of GetFraudCases.
type GetFraudCasesParams struct {
	Status        *string `json:"status,omitempty"`
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetFraudCasesParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "status", p.Status)
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetFraudCases calls List fraud cases. GET /v1/fraud/cases.
//
// The fraud case queue (self-referrals, quality, rings, value patterns, identity risk), newest first, with the evidence. Filter with `status`.
//
// **Auth:** an API key with `fraud:read` (or `full`), or a session of any role.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetFraudCases(ctx context.Context, params *GetFraudCasesParams, opts ...RequestOption) (*Page[FraudCase], error) {
	var out Page[FraudCase]
	if err := c.do(ctx, "GET", "/v1/fraud/cases", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetFraudCasesEach yields every item of GetFraudCases, following the pages.
func (c *Client) GetFraudCasesEach(ctx context.Context, params *GetFraudCasesParams, opts ...RequestOption) iter.Seq2[FraudCase, error] {
	return paginate(func(cursor *string) (*Page[FraudCase], error) {
		p := GetFraudCasesParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetFraudCases(ctx, &p, opts...)
	})
}

// GetFXRatesParams are the query parameters of GetFXRates.
type GetFXRatesParams struct {
	Base          *string `json:"base,omitempty"`
	Quote         *string `json:"quote,omitempty"`
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetFXRatesParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "base", p.Base)
	setQuery(v, "quote", p.Quote)
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetFXRates calls List FX rates. GET /v1/fx-rates.
//
// The org's rates first, then the platform fallback, newest per pair first. A payout uses the org's latest rate at or before the run, else the platform's. Filter by `base` and `quote`.
//
// **Auth:** an API key with `payouts:read` (or `full`), or a session with role `owner`, `admin`, `finance`.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetFXRates(ctx context.Context, params *GetFXRatesParams, opts ...RequestOption) (*Page[FxRate], error) {
	var out Page[FxRate]
	if err := c.do(ctx, "GET", "/v1/fx-rates", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetFXRatesEach yields every item of GetFXRates, following the pages.
func (c *Client) GetFXRatesEach(ctx context.Context, params *GetFXRatesParams, opts ...RequestOption) iter.Seq2[FxRate, error] {
	return paginate(func(cursor *string) (*Page[FxRate], error) {
		p := GetFXRatesParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetFXRates(ctx, &p, opts...)
	})
}

// GetGraphql calls Query the GraphQL read layer. GET /v1/graphql.
//
// A read-only GraphQL endpoint over the same org-scoped data as REST (programs, partners, referrals, analytics), with the same authentication. Mutations stay in REST.
//
// **Auth:** an API key with `analytics:read` (or `full`), or a session of any role.
func (c *Client) GetGraphql(ctx context.Context, opts ...RequestOption) (*GetGraphqlResponse, error) {
	var out GetGraphqlResponse
	if err := c.do(ctx, "GET", "/v1/graphql", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetHealth calls Check API health. GET /v1/health.
//
// `200 { "ok": true }` when this instance can serve requests: both database pools answer and it is not shutting down. Otherwise `503` with `reason` (`shutting_down`, with `Retry-After`, or `database_unreachable`).
//
// **Auth:** none.
func (c *Client) GetHealth(ctx context.Context, opts ...RequestOption) (*Ok, error) {
	var out Ok
	if err := c.do(ctx, "GET", "/v1/health", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetHostsCurrent calls Describe the current host. GET /v1/hosts/current.
//
// What the request's host is: a merchant's verified portal domain (with the org's name, branding and public programs), a merchant link domain, or Referalli's own app. Used by the web app before anyone signs in.
//
// **Auth:** none.
func (c *Client) GetHostsCurrent(ctx context.Context, opts ...RequestOption) (*GetHostsCurrentResponse, error) {
	var out GetHostsCurrentResponse
	if err := c.do(ctx, "GET", "/v1/hosts/current", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetHostsCurrentThemeCss calls Get the current host's theme stylesheet. GET /v1/hosts/current/theme.css.
//
// On a merchant's portal domain, the org's theme tokens as a `:root` stylesheet; empty elsewhere.
//
// **Auth:** none.
func (c *Client) GetHostsCurrentThemeCss(ctx context.Context, opts ...RequestOption) (string, error) {
	var out string
	err := c.do(ctx, "GET", "/v1/hosts/current/theme.css", nil, nil, &out, "text", opts)
	return out, err
}

// GetIncrementalityParams are the query parameters of GetIncrementality.
type GetIncrementalityParams struct {
	ProgramID *string `json:"programId,omitempty"`
}

func (p *GetIncrementalityParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "programId", p.ProgramID)
	return v
}

// GetIncrementality calls Get incrementality results. GET /v1/incrementality.
//
// For a program's active referrals: treatment and control counts, conversion rates and lift, overall and per partner.
//
// **Auth:** an API key with `analytics:read` (or `full`), or a session of any role.
func (c *Client) GetIncrementality(ctx context.Context, params *GetIncrementalityParams, opts ...RequestOption) (*GetIncrementalityResponse, error) {
	var out GetIncrementalityResponse
	if err := c.do(ctx, "GET", "/v1/incrementality", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetIntegrationHealth calls Check integration health. GET /v1/integration-health.
//
// Every connection the org depends on, each with `ok`, a `detail`, and when it fails a `fix` and `docs` link: billing connections (and when each last received an event), sign-up connections, the tracking snippet (a visit in the last 30 days), payout connections, outbound webhook endpoints, and every adapter of a connected port. `hop` names the leg that failed on a path between systems. `ok` at the top is true only when every check is.
//
// **Auth:** an API key with `integrations:read` (or `full`), or a session of any role.
func (c *Client) GetIntegrationHealth(ctx context.Context, opts ...RequestOption) (*GetIntegrationHealthResponse, error) {
	var out GetIntegrationHealthResponse
	if err := c.do(ctx, "GET", "/v1/integration-health", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetIntegrations calls List integrations. GET /v1/integrations.
//
// The HubSpot and Segment connections with status, settings and the last test result. Credentials are never returned.
//
// **Auth:** an API key with `integrations:read` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) GetIntegrations(ctx context.Context, opts ...RequestOption) (*GetIntegrationsResponse, error) {
	var out GetIntegrationsResponse
	if err := c.do(ctx, "GET", "/v1/integrations", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetIntegrationsDeliveriesParams are the query parameters of GetIntegrationsDeliveries.
type GetIntegrationsDeliveriesParams struct {
	Kind          *string `json:"kind,omitempty"`
	Status        *string `json:"status,omitempty"`
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetIntegrationsDeliveriesParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "kind", p.Kind)
	setQuery(v, "status", p.Status)
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetIntegrationsDeliveries calls List integration deliveries. GET /v1/integrations/deliveries.
//
// What was sent to Zapier, HubSpot and Segment, newest first, filtered by `kind` and `status`.
//
// **Auth:** an API key with `integrations:read` (or `full`), or a session with role `owner`, `admin`.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetIntegrationsDeliveries(ctx context.Context, params *GetIntegrationsDeliveriesParams, opts ...RequestOption) (*Page[IntegrationDelivery], error) {
	var out Page[IntegrationDelivery]
	if err := c.do(ctx, "GET", "/v1/integrations/deliveries", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetIntegrationsDeliveriesEach yields every item of GetIntegrationsDeliveries, following the pages.
func (c *Client) GetIntegrationsDeliveriesEach(ctx context.Context, params *GetIntegrationsDeliveriesParams, opts ...RequestOption) iter.Seq2[IntegrationDelivery, error] {
	return paginate(func(cursor *string) (*Page[IntegrationDelivery], error) {
		p := GetIntegrationsDeliveriesParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetIntegrationsDeliveries(ctx, &p, opts...)
	})
}

// GetInvitationsParams are the query parameters of GetInvitations.
type GetInvitationsParams struct {
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetInvitationsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetInvitations calls List creator invitations. GET /v1/invitations.
//
// The org's creator invitations and their status (`pending`, `accepted`, `declined`, `withdrawn`).
//
// **Auth:** an API key with `partners:read` (or `full`), or a session with role `owner`, `admin`.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetInvitations(ctx context.Context, params *GetInvitationsParams, opts ...RequestOption) (*Page[GetInvitationsResponseDataItem], error) {
	var out Page[GetInvitationsResponseDataItem]
	if err := c.do(ctx, "GET", "/v1/invitations", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetInvitationsEach yields every item of GetInvitations, following the pages.
func (c *Client) GetInvitationsEach(ctx context.Context, params *GetInvitationsParams, opts ...RequestOption) iter.Seq2[GetInvitationsResponseDataItem, error] {
	return paginate(func(cursor *string) (*Page[GetInvitationsResponseDataItem], error) {
		p := GetInvitationsParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetInvitations(ctx, &p, opts...)
	})
}

// GetLinksParams are the query parameters of GetLinks.
type GetLinksParams struct {
	ProgramID     *string `json:"programId,omitempty"`
	PartnerID     *string `json:"partnerId,omitempty"`
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetLinksParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "programId", p.ProgramID)
	setQuery(v, "partnerId", p.PartnerID)
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetLinks calls List links. GET /v1/links.
//
// Tracked links with their short URL and click count, filtered by `programId` and/or `partnerId`.
//
// **Auth:** an API key with `links:read` (or `full`), or a session of any role.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetLinks(ctx context.Context, params *GetLinksParams, opts ...RequestOption) (*Page[Link], error) {
	var out Page[Link]
	if err := c.do(ctx, "GET", "/v1/links", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetLinksEach yields every item of GetLinks, following the pages.
func (c *Client) GetLinksEach(ctx context.Context, params *GetLinksParams, opts ...RequestOption) iter.Seq2[Link, error] {
	return paginate(func(cursor *string) (*Page[Link], error) {
		p := GetLinksParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetLinks(ctx, &p, opts...)
	})
}

// GetLinksIDQR calls Get a link's QR code. GET /v1/links/{id}/qr.
//
// An SVG QR code for the link's short URL.
//
// **Auth:** an API key with `links:read` (or `full`), or a session of any role.
func (c *Client) GetLinksIDQR(ctx context.Context, id string, opts ...RequestOption) (string, error) {
	var out string
	err := c.do(ctx, "GET", "/v1/links/"+url.PathEscape(id)+"/qr", nil, nil, &out, "text", opts)
	return out, err
}

// GetMarketplaceListingParams are the query parameters of GetMarketplaceListing.
type GetMarketplaceListingParams struct {
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetMarketplaceListingParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetMarketplaceListing calls List the org's marketplace listings. GET /v1/marketplace/listing.
//
// The org's listings with their platform-computed stats: eligible clicks, active referrals, earnings per click, conversion rate, payout reliability and approval latency.
//
// **Auth:** an API key with `growth:read` (or `full`), or a session of any role.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetMarketplaceListing(ctx context.Context, params *GetMarketplaceListingParams, opts ...RequestOption) (*Page[GetMarketplaceListingResponseDataItem], error) {
	var out Page[GetMarketplaceListingResponseDataItem]
	if err := c.do(ctx, "GET", "/v1/marketplace/listing", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetMarketplaceListingEach yields every item of GetMarketplaceListing, following the pages.
func (c *Client) GetMarketplaceListingEach(ctx context.Context, params *GetMarketplaceListingParams, opts ...RequestOption) iter.Seq2[GetMarketplaceListingResponseDataItem, error] {
	return paginate(func(cursor *string) (*Page[GetMarketplaceListingResponseDataItem], error) {
		p := GetMarketplaceListingParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetMarketplaceListing(ctx, &p, opts...)
	})
}

// GetMe calls Get the current caller. GET /v1/me.
//
// Who is calling. With a user session: the user and their org memberships (no `x-org-id` needed). With an API key: the key (name, prefix, scopes, pinned version, mode, owner, expiry) and its org.
//
// **Auth:** an API key with `team:read` (or `full`), or a session of any role.
func (c *Client) GetMe(ctx context.Context, opts ...RequestOption) (*GetMeResponse, error) {
	var out GetMeResponse
	if err := c.do(ctx, "GET", "/v1/me", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetNegotiationsParams are the query parameters of GetNegotiations.
type GetNegotiationsParams struct {
	ProgramID     *string `json:"programId,omitempty"`
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetNegotiationsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "programId", p.ProgramID)
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetNegotiations calls List negotiation rounds. GET /v1/negotiations.
//
// Agent negotiation rounds and their decisions, newest first, optionally for one program.
//
// **Auth:** an API key with `growth:read` (or `full`), or a session of any role.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetNegotiations(ctx context.Context, params *GetNegotiationsParams, opts ...RequestOption) (*Page[GetNegotiationsResponseDataItem], error) {
	var out Page[GetNegotiationsResponseDataItem]
	if err := c.do(ctx, "GET", "/v1/negotiations", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetNegotiationsEach yields every item of GetNegotiations, following the pages.
func (c *Client) GetNegotiationsEach(ctx context.Context, params *GetNegotiationsParams, opts ...RequestOption) iter.Seq2[GetNegotiationsResponseDataItem, error] {
	return paginate(func(cursor *string) (*Page[GetNegotiationsResponseDataItem], error) {
		p := GetNegotiationsParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetNegotiations(ctx, &p, opts...)
	})
}

// GetNetworkJourneyParams are the query parameters of GetNetworkJourney.
type GetNetworkJourneyParams struct {
	EmailHash *string `json:"emailHash,omitempty"`
	Model     *string `json:"model,omitempty"`
}

func (p *GetNetworkJourneyParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "emailHash", p.EmailHash)
	setQuery(v, "model", p.Model)
	return v
}

// GetNetworkJourney calls Get a customer's network journey. GET /v1/network/journey.
//
// For a hashed customer email, the partner accounts that touched them at any opted-in merchant (never which merchant) and each one's credit under `model`. 403 unless the org has joined the network.
//
// **Auth:** an API key with `growth:read` (or `full`), or a session of any role.
func (c *Client) GetNetworkJourney(ctx context.Context, params *GetNetworkJourneyParams, opts ...RequestOption) (*GetNetworkJourneyResponse, error) {
	var out GetNetworkJourneyResponse
	if err := c.do(ctx, "GET", "/v1/network/journey", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetNetworkSettings calls Get network membership. GET /v1/network/settings.
//
// Whether the org contributes to (and can read) the attribution network.
//
// **Auth:** an API key with `growth:read` (or `full`), or a session of any role.
func (c *Client) GetNetworkSettings(ctx context.Context, opts ...RequestOption) (*GetNetworkSettingsResponse, error) {
	var out GetNetworkSettingsResponse
	if err := c.do(ctx, "GET", "/v1/network/settings", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetNotificationsParams are the query parameters of GetNotifications.
type GetNotificationsParams struct {
	Unread        *string `json:"unread,omitempty"`
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetNotificationsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "unread", p.Unread)
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetNotifications calls List notifications. GET /v1/notifications.
//
// The org's in-app notifications (failing webhooks, anomalies, answered invitations…), newest first. `unread=true` for unread only.
//
// **Auth:** an API key with `org:read` (or `full`), or a session of any role.
//
// unread: true: unread notifications only.
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetNotifications(ctx context.Context, params *GetNotificationsParams, opts ...RequestOption) (*Page[Notification], error) {
	var out Page[Notification]
	if err := c.do(ctx, "GET", "/v1/notifications", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetNotificationsEach yields every item of GetNotifications, following the pages.
func (c *Client) GetNotificationsEach(ctx context.Context, params *GetNotificationsParams, opts ...RequestOption) iter.Seq2[Notification, error] {
	return paginate(func(cursor *string) (*Page[Notification], error) {
		p := GetNotificationsParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetNotifications(ctx, &p, opts...)
	})
}

// GetOnchainConnectionsParams are the query parameters of GetOnchainConnections.
type GetOnchainConnectionsParams struct {
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetOnchainConnectionsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetOnchainConnections calls List on-chain connections. GET /v1/onchain/connections.
//
// The org's on-chain conversion connections.
//
// **Auth:** an API key with `billing:read` (or `full`), or a session of any role.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetOnchainConnections(ctx context.Context, params *GetOnchainConnectionsParams, opts ...RequestOption) (*Page[GetOnchainConnectionsResponseDataItem], error) {
	var out Page[GetOnchainConnectionsResponseDataItem]
	if err := c.do(ctx, "GET", "/v1/onchain/connections", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetOnchainConnectionsEach yields every item of GetOnchainConnections, following the pages.
func (c *Client) GetOnchainConnectionsEach(ctx context.Context, params *GetOnchainConnectionsParams, opts ...RequestOption) iter.Seq2[GetOnchainConnectionsResponseDataItem, error] {
	return paginate(func(cursor *string) (*Page[GetOnchainConnectionsResponseDataItem], error) {
		p := GetOnchainConnectionsParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetOnchainConnections(ctx, &p, opts...)
	})
}

// GetOpenapiJSON calls Get the OpenAPI document. GET /v1/openapi.json.
//
// This reference as an OpenAPI 3.1 document (JSON).
//
// **Auth:** none.
func (c *Client) GetOpenapiJSON(ctx context.Context, opts ...RequestOption) (*GetOpenapiJsonResponse, error) {
	var out GetOpenapiJsonResponse
	if err := c.do(ctx, "GET", "/v1/openapi.json", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetOpenapiYaml calls Get the OpenAPI document as YAML. GET /v1/openapi.yaml.
//
// The same OpenAPI 3.1 document as `GET /v1/openapi.json`, serialized as YAML.
//
// **Auth:** none.
func (c *Client) GetOpenapiYaml(ctx context.Context, opts ...RequestOption) (string, error) {
	var out string
	err := c.do(ctx, "GET", "/v1/openapi.yaml", nil, nil, &out, "text", opts)
	return out, err
}

// GetOrg calls Get the organization. GET /v1/org.
//
// The current org: name, plan, settings, public `trackingToken`, payout KYC requirement and mode (`livemode`; test orgs also carry `liveOrgId`).
//
// **Auth:** an API key with `org:read` (or `full`), or a session of any role.
func (c *Client) GetOrg(ctx context.Context, opts ...RequestOption) (*Org, error) {
	var out Org
	if err := c.do(ctx, "GET", "/v1/org", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetOrgFXFeed calls Get the FX feed setting. GET /v1/org/fx-feed.
//
// Whether the daily ECB reference-rate feed is on for this org, and the date of its latest rate.
//
// **Auth:** an API key with `org:read` (or `full`), or a session with role `owner`, `admin`, `finance`.
func (c *Client) GetOrgFXFeed(ctx context.Context, opts ...RequestOption) (*GetOrgFxFeedResponse, error) {
	var out GetOrgFxFeedResponse
	if err := c.do(ctx, "GET", "/v1/org/fx-feed", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetOrgHostedPages calls Get the hosted-page consent and chat. GET /v1/org/hosted-pages.
//
// What runs on the pages Referalli hosts for you on your own portal domain (the partner portal and `/join`): your consent banner (Cookie Munch, OneTrust, Cookiebot, Usercentrics, Didomi, Osano, Termly, iubenda, a custom CMP, a built-in necessary-only notice, or none) and your support chat (Chatly, Intercom, Crisp or custom). `served` is exactly what visitors get: the head tags, the "Cookie settings" link and the Content-Security-Policy built from these adapters. Referalli's own banner and chat never run on your domain.
//
// **Auth:** an API key with `org:read` (or `full`), or a session of any role.
func (c *Client) GetOrgHostedPages(ctx context.Context, opts ...RequestOption) (*GetOrgHostedPagesResponse, error) {
	var out GetOrgHostedPagesResponse
	if err := c.do(ctx, "GET", "/v1/org/hosted-pages", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetOrgMembersParams are the query parameters of GetOrgMembers.
type GetOrgMembersParams struct {
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetOrgMembersParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetOrgMembers calls List team members. GET /v1/org/members.
//
// Members of the org with their role and email, oldest first.
//
// **Auth:** an API key with `team:read` (or `full`), or a session of any role.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetOrgMembers(ctx context.Context, params *GetOrgMembersParams, opts ...RequestOption) (*Page[Member], error) {
	var out Page[Member]
	if err := c.do(ctx, "GET", "/v1/org/members", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetOrgMembersEach yields every item of GetOrgMembers, following the pages.
func (c *Client) GetOrgMembersEach(ctx context.Context, params *GetOrgMembersParams, opts ...RequestOption) iter.Seq2[Member, error] {
	return paginate(func(cursor *string) (*Page[Member], error) {
		p := GetOrgMembersParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetOrgMembers(ctx, &p, opts...)
	})
}

// GetOrgPayoutSchedule calls Get payout schedules. GET /v1/org/payout-schedule.
//
// The org schedule and its next run date, plus each program's override (`{ kind: 'off' }`, its own schedule, or `null` to follow the org) and next run date.
//
// **Auth:** an API key with `payouts:read` (or `full`), or a session of any role.
func (c *Client) GetOrgPayoutSchedule(ctx context.Context, opts ...RequestOption) (*GetOrgPayoutScheduleResponse, error) {
	var out GetOrgPayoutScheduleResponse
	if err := c.do(ctx, "GET", "/v1/org/payout-schedule", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetOrgPortalAuth calls Get partner sign-in settings. GET /v1/org/portal-auth.
//
// How partners sign in to the org's portal: `referalli` (partner accounts, the default), `oidc` (the merchant's IdP) or `custom` (the merchant's own session, as a short-lived JWT). Includes the audience tokens must carry, the URL to send partners to, the redirect URI to register, and the last check. Never the secret.
//
// **Auth:** an API key with `org:read` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) GetOrgPortalAuth(ctx context.Context, opts ...RequestOption) (*GetOrgPortalAuthResponse, error) {
	var out GetOrgPortalAuthResponse
	if err := c.do(ctx, "GET", "/v1/org/portal-auth", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetOrgReferralParams calls Get referral query parameters. GET /v1/org/referral-params.
//
// The query parameters aff.js reads a referral code from: the built-in ones (`via`, `ref`, `aff`) and the org's extra ones.
//
// **Auth:** an API key with `org:read` (or `full`), or a session of any role.
func (c *Client) GetOrgReferralParams(ctx context.Context, opts ...RequestOption) (*GetOrgReferralParamsResponse, error) {
	var out GetOrgReferralParamsResponse
	if err := c.do(ctx, "GET", "/v1/org/referral-params", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetOrgRetention calls Get the data retention policy. GET /v1/org/retention.
//
// How many days each kind of raw data is kept before the nightly purge, the allowed minimum and default per kind, and the last purge run.
//
// **Auth:** an API key with `org:read` (or `full`), or a session of any role.
func (c *Client) GetOrgRetention(ctx context.Context, opts ...RequestOption) (*Retention, error) {
	var out Retention
	if err := c.do(ctx, "GET", "/v1/org/retention", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetOrgSSO calls Get SSO settings. GET /v1/org/sso.
//
// Single sign-on for the org: whether SSO is configured on this deployment and linked to a directory tenant, its connections and domains, whether SSO is required, the role mapping, recent setup tickets and any provider error.
//
// **Auth:** an API key with `org:read` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) GetOrgSSO(ctx context.Context, opts ...RequestOption) (*GetOrgSsoResponse, error) {
	var out GetOrgSsoResponse
	if err := c.do(ctx, "GET", "/v1/org/sso", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetOrgSSOConnectionsIDSamlMetadata calls Get SAML metadata. GET /v1/org/sso/connections/{id}/saml-metadata.
//
// The SAML service-provider metadata for one of the org's SSO connections, to give the identity provider.
//
// **Auth:** an API key with `org:read` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) GetOrgSSOConnectionsIDSamlMetadata(ctx context.Context, id string, opts ...RequestOption) (*GetOrgSsoConnectionsIdSamlMetadataResponse, error) {
	var out GetOrgSsoConnectionsIdSamlMetadataResponse
	if err := c.do(ctx, "GET", "/v1/org/sso/connections/"+url.PathEscape(id)+"/saml-metadata", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetOrgTaxProfile calls Get the tax profile. GET /v1/org/tax-profile.
//
// The org's legal name, address, country, VAT id, EIN, domestic VAT rate, whether partner VAT ids are checked against EU VIES, and the 1099 filing provider.
//
// **Auth:** an API key with `org:read` (or `full`), or a session with role `owner`, `admin`, `finance`.
func (c *Client) GetOrgTaxProfile(ctx context.Context, opts ...RequestOption) (*TaxProfile, error) {
	var out TaxProfile
	if err := c.do(ctx, "GET", "/v1/org/tax-profile", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetOrgTrackingConsent calls Get the tracker's consent default. GET /v1/org/tracking-consent.
//
// How `aff.js` reads visitor consent unless the tag says otherwise (`data-consent`): `auto` detects the site's CMP (Cookie Munch, OneTrust, Cookiebot, Usercentrics, Didomi, Osano, Termly, IAB TCF and GPP, Google Consent Mode, or `window.referalliConsent`), an adapter id pins one, `pending` waits for `aff.consent(true|false)`, and `none` always stores. Until consent is granted nothing is stored or sent; the referral stays in memory for the sign-up. `noCmp` is what `auto` does on a site without a CMP. `configured: false` is an org that never chose: it serves `none`, so existing tags behave as before. Lists the adapters with their default category.
//
// **Auth:** an API key with `org:read` (or `full`), or a session of any role.
func (c *Client) GetOrgTrackingConsent(ctx context.Context, opts ...RequestOption) (*GetOrgTrackingConsentResponse, error) {
	var out GetOrgTrackingConsentResponse
	if err := c.do(ctx, "GET", "/v1/org/tracking-consent", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetOutreachParams are the query parameters of GetOutreach.
type GetOutreachParams struct {
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetOutreachParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetOutreach calls List outreach messages. GET /v1/outreach.
//
// Drafted, approved and sent outreach messages, newest first.
//
// **Auth:** an API key with `growth:read` (or `full`), or a session of any role.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetOutreach(ctx context.Context, params *GetOutreachParams, opts ...RequestOption) (*Page[GetOutreachResponseDataItem], error) {
	var out Page[GetOutreachResponseDataItem]
	if err := c.do(ctx, "GET", "/v1/outreach", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetOutreachEach yields every item of GetOutreach, following the pages.
func (c *Client) GetOutreachEach(ctx context.Context, params *GetOutreachParams, opts ...RequestOption) iter.Seq2[GetOutreachResponseDataItem, error] {
	return paginate(func(cursor *string) (*Page[GetOutreachResponseDataItem], error) {
		p := GetOutreachParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetOutreach(ctx, &p, opts...)
	})
}

// GetPartnerAchievementsParams are the query parameters of GetPartnerAchievements.
type GetPartnerAchievementsParams struct {
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetPartnerAchievementsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetPartnerAchievements calls List my achievements. GET /v1/partner/achievements.
//
// Badges computed from the ledger (first click, referral milestones, first payout, $1,000 earned), each with progress toward its target.
//
// **Auth:** a partner session.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetPartnerAchievements(ctx context.Context, params *GetPartnerAchievementsParams, opts ...RequestOption) (*Page[GetPartnerAchievementsResponseDataItem], error) {
	var out Page[GetPartnerAchievementsResponseDataItem]
	if err := c.do(ctx, "GET", "/v1/partner/achievements", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetPartnerAchievementsEach yields every item of GetPartnerAchievements, following the pages.
func (c *Client) GetPartnerAchievementsEach(ctx context.Context, params *GetPartnerAchievementsParams, opts ...RequestOption) iter.Seq2[GetPartnerAchievementsResponseDataItem, error] {
	return paginate(func(cursor *string) (*Page[GetPartnerAchievementsResponseDataItem], error) {
		p := GetPartnerAchievementsParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetPartnerAchievements(ctx, &p, opts...)
	})
}

// GetPartnerAdvanceParams are the query parameters of GetPartnerAdvance.
type GetPartnerAdvanceParams struct {
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetPartnerAdvanceParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetPartnerAdvance calls List my advances. GET /v1/partner/advance.
//
// The partner's advances with what remains to repay.
//
// **Auth:** a partner session.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetPartnerAdvance(ctx context.Context, params *GetPartnerAdvanceParams, opts ...RequestOption) (*Page[GetPartnerAdvanceResponseDataItem], error) {
	var out Page[GetPartnerAdvanceResponseDataItem]
	if err := c.do(ctx, "GET", "/v1/partner/advance", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetPartnerAdvanceEach yields every item of GetPartnerAdvance, following the pages.
func (c *Client) GetPartnerAdvanceEach(ctx context.Context, params *GetPartnerAdvanceParams, opts ...RequestOption) iter.Seq2[GetPartnerAdvanceResponseDataItem, error] {
	return paginate(func(cursor *string) (*Page[GetPartnerAdvanceResponseDataItem], error) {
		p := GetPartnerAdvanceParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetPartnerAdvance(ctx, &p, opts...)
	})
}

// GetPartnerAdvanceOffer calls Get my advance offer. GET /v1/partner/advance/offer.
//
// How much the partner could receive now against future commissions (`maxAdvanceMinor`), from their recent run rate and quality score, and the fee.
//
// **Auth:** a partner session.
func (c *Client) GetPartnerAdvanceOffer(ctx context.Context, opts ...RequestOption) (*GetPartnerAdvanceOfferResponse, error) {
	var out GetPartnerAdvanceOfferResponse
	if err := c.do(ctx, "GET", "/v1/partner/advance/offer", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetPartnerAssetsParams are the query parameters of GetPartnerAssets.
type GetPartnerAssetsParams struct {
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetPartnerAssetsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetPartnerAssets calls List marketing assets. GET /v1/partner/assets.
//
// Marketing assets (copy, images, banners) of the programs the partner is approved in; group-only assets show only to that group's members.
//
// **Auth:** a partner session.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetPartnerAssets(ctx context.Context, params *GetPartnerAssetsParams, opts ...RequestOption) (*Page[GetPartnerAssetsResponseDataItem], error) {
	var out Page[GetPartnerAssetsResponseDataItem]
	if err := c.do(ctx, "GET", "/v1/partner/assets", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetPartnerAssetsEach yields every item of GetPartnerAssets, following the pages.
func (c *Client) GetPartnerAssetsEach(ctx context.Context, params *GetPartnerAssetsParams, opts ...RequestOption) iter.Seq2[GetPartnerAssetsResponseDataItem, error] {
	return paginate(func(cursor *string) (*Page[GetPartnerAssetsResponseDataItem], error) {
		p := GetPartnerAssetsParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetPartnerAssets(ctx, &p, opts...)
	})
}

// GetPartnerContractsParams are the query parameters of GetPartnerContracts.
type GetPartnerContractsParams struct {
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetPartnerContractsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetPartnerContracts calls List my contracts. GET /v1/partner/contracts.
//
// Contracts merchants sent the partner.
//
// **Auth:** a partner session.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetPartnerContracts(ctx context.Context, params *GetPartnerContractsParams, opts ...RequestOption) (*Page[GetPartnerContractsResponseDataItem], error) {
	var out Page[GetPartnerContractsResponseDataItem]
	if err := c.do(ctx, "GET", "/v1/partner/contracts", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetPartnerContractsEach yields every item of GetPartnerContracts, following the pages.
func (c *Client) GetPartnerContractsEach(ctx context.Context, params *GetPartnerContractsParams, opts ...RequestOption) iter.Seq2[GetPartnerContractsResponseDataItem, error] {
	return paginate(func(cursor *string) (*Page[GetPartnerContractsResponseDataItem], error) {
		p := GetPartnerContractsParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetPartnerContracts(ctx, &p, opts...)
	})
}

// GetPartnerDealsParams are the query parameters of GetPartnerDeals.
type GetPartnerDealsParams struct {
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetPartnerDealsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetPartnerDeals calls List my deals. GET /v1/partner/deals.
//
// The partner's registered deals across merchants, newest first.
//
// **Auth:** a partner session.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetPartnerDeals(ctx context.Context, params *GetPartnerDealsParams, opts ...RequestOption) (*Page[GetPartnerDealsResponseDataItem], error) {
	var out Page[GetPartnerDealsResponseDataItem]
	if err := c.do(ctx, "GET", "/v1/partner/deals", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetPartnerDealsEach yields every item of GetPartnerDeals, following the pages.
func (c *Client) GetPartnerDealsEach(ctx context.Context, params *GetPartnerDealsParams, opts ...RequestOption) iter.Seq2[GetPartnerDealsResponseDataItem, error] {
	return paginate(func(cursor *string) (*Page[GetPartnerDealsResponseDataItem], error) {
		p := GetPartnerDealsParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetPartnerDeals(ctx, &p, opts...)
	})
}

// GetPartnerDisputesParams are the query parameters of GetPartnerDisputes.
type GetPartnerDisputesParams struct {
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetPartnerDisputesParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetPartnerDisputes calls List my disputes. GET /v1/partner/disputes.
//
// Disputes the partner filed about missing or wrong credit, newest first, with the merchant's resolution.
//
// **Auth:** a partner session.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetPartnerDisputes(ctx context.Context, params *GetPartnerDisputesParams, opts ...RequestOption) (*Page[GetPartnerDisputesResponseDataItem], error) {
	var out Page[GetPartnerDisputesResponseDataItem]
	if err := c.do(ctx, "GET", "/v1/partner/disputes", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetPartnerDisputesEach yields every item of GetPartnerDisputes, following the pages.
func (c *Client) GetPartnerDisputesEach(ctx context.Context, params *GetPartnerDisputesParams, opts ...RequestOption) iter.Seq2[GetPartnerDisputesResponseDataItem, error] {
	return paginate(func(cursor *string) (*Page[GetPartnerDisputesResponseDataItem], error) {
		p := GetPartnerDisputesParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetPartnerDisputes(ctx, &p, opts...)
	})
}

// GetPartnerGroupsParams are the query parameters of GetPartnerGroups.
type GetPartnerGroupsParams struct {
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetPartnerGroupsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetPartnerGroups calls List my groups. GET /v1/partner/groups.
//
// The partner's groups across programs, with their onboarding and payout terms.
//
// **Auth:** a partner session.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetPartnerGroups(ctx context.Context, params *GetPartnerGroupsParams, opts ...RequestOption) (*Page[GetPartnerGroupsResponseDataItem], error) {
	var out Page[GetPartnerGroupsResponseDataItem]
	if err := c.do(ctx, "GET", "/v1/partner/groups", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetPartnerGroupsEach yields every item of GetPartnerGroups, following the pages.
func (c *Client) GetPartnerGroupsEach(ctx context.Context, params *GetPartnerGroupsParams, opts ...RequestOption) iter.Seq2[GetPartnerGroupsResponseDataItem, error] {
	return paginate(func(cursor *string) (*Page[GetPartnerGroupsResponseDataItem], error) {
		p := GetPartnerGroupsParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetPartnerGroups(ctx, &p, opts...)
	})
}

// GetPartnerInvitationsParams are the query parameters of GetPartnerInvitations.
type GetPartnerInvitationsParams struct {
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetPartnerInvitationsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetPartnerInvitations calls List my invitations. GET /v1/partner/invitations.
//
// Program invitations merchants sent the partner.
//
// **Auth:** a partner session.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetPartnerInvitations(ctx context.Context, params *GetPartnerInvitationsParams, opts ...RequestOption) (*Page[GetPartnerInvitationsResponseDataItem], error) {
	var out Page[GetPartnerInvitationsResponseDataItem]
	if err := c.do(ctx, "GET", "/v1/partner/invitations", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetPartnerInvitationsEach yields every item of GetPartnerInvitations, following the pages.
func (c *Client) GetPartnerInvitationsEach(ctx context.Context, params *GetPartnerInvitationsParams, opts ...RequestOption) iter.Seq2[GetPartnerInvitationsResponseDataItem, error] {
	return paginate(func(cursor *string) (*Page[GetPartnerInvitationsResponseDataItem], error) {
		p := GetPartnerInvitationsParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetPartnerInvitations(ctx, &p, opts...)
	})
}

// GetPartnerKYCStatus calls Get my identity verification status. GET /v1/partner/kyc/status.
//
// The partner account's KYC status (`unverified`, `pending`, `verified`, `rejected`), sanctions screening status and when it was verified.
//
// **Auth:** a partner session.
func (c *Client) GetPartnerKYCStatus(ctx context.Context, opts ...RequestOption) (*GetPartnerKycStatusResponse, error) {
	var out GetPartnerKycStatusResponse
	if err := c.do(ctx, "GET", "/v1/partner/kyc/status", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetPartnerLinksParams are the query parameters of GetPartnerLinks.
type GetPartnerLinksParams struct {
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetPartnerLinksParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetPartnerLinks calls List my links. GET /v1/partner/links.
//
// The partner's tracked links across merchants, newest first, with the short URL, the `?via=` parameter for direct links and eligible clicks.
//
// **Auth:** a partner session.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetPartnerLinks(ctx context.Context, params *GetPartnerLinksParams, opts ...RequestOption) (*Page[GetPartnerLinksResponseDataItem], error) {
	var out Page[GetPartnerLinksResponseDataItem]
	if err := c.do(ctx, "GET", "/v1/partner/links", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetPartnerLinksEach yields every item of GetPartnerLinks, following the pages.
func (c *Client) GetPartnerLinksEach(ctx context.Context, params *GetPartnerLinksParams, opts ...RequestOption) iter.Seq2[GetPartnerLinksResponseDataItem, error] {
	return paginate(func(cursor *string) (*Page[GetPartnerLinksResponseDataItem], error) {
		p := GetPartnerLinksParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetPartnerLinks(ctx, &p, opts...)
	})
}

// GetPartnerLinksIDQR calls Get my link's QR code. GET /v1/partner/links/{id}/qr.
//
// An SVG QR code for one of the partner's links.
//
// **Auth:** a partner session.
func (c *Client) GetPartnerLinksIDQR(ctx context.Context, id string, opts ...RequestOption) (string, error) {
	var out string
	err := c.do(ctx, "GET", "/v1/partner/links/"+url.PathEscape(id)+"/qr", nil, nil, &out, "text", opts)
	return out, err
}

// GetPartnerMarketplaceParams are the query parameters of GetPartnerMarketplace.
type GetPartnerMarketplaceParams struct {
	Category      *string `json:"category,omitempty"`
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetPartnerMarketplaceParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "category", p.Category)
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetPartnerMarketplace calls Browse the marketplace. GET /v1/partner/marketplace.
//
// Public listings of active live programs across every merchant, sponsored first, with commission terms and platform-computed stats. Filter with `category`. Partner session only.
//
// **Auth:** a partner session.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetPartnerMarketplace(ctx context.Context, params *GetPartnerMarketplaceParams, opts ...RequestOption) (*Page[GetPartnerMarketplaceResponseDataItem], error) {
	var out Page[GetPartnerMarketplaceResponseDataItem]
	if err := c.do(ctx, "GET", "/v1/partner/marketplace", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetPartnerMarketplaceEach yields every item of GetPartnerMarketplace, following the pages.
func (c *Client) GetPartnerMarketplaceEach(ctx context.Context, params *GetPartnerMarketplaceParams, opts ...RequestOption) iter.Seq2[GetPartnerMarketplaceResponseDataItem, error] {
	return paginate(func(cursor *string) (*Page[GetPartnerMarketplaceResponseDataItem], error) {
		p := GetPartnerMarketplaceParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetPartnerMarketplace(ctx, &p, opts...)
	})
}

// GetPartnerMe calls Get my account and memberships. GET /v1/partner/me.
//
// The partner account and every merchant it partners with, each with its programs and enrollment status.
//
// **Auth:** a partner session.
func (c *Client) GetPartnerMe(ctx context.Context, opts ...RequestOption) (*GetPartnerMeResponse, error) {
	var out GetPartnerMeResponse
	if err := c.do(ctx, "GET", "/v1/partner/me", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetPartnerNotificationsParams are the query parameters of GetPartnerNotifications.
type GetPartnerNotificationsParams struct {
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetPartnerNotificationsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetPartnerNotifications calls List my notifications. GET /v1/partner/notifications.
//
// The partner account's notifications, newest first.
//
// **Auth:** a partner session.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetPartnerNotifications(ctx context.Context, params *GetPartnerNotificationsParams, opts ...RequestOption) (*Page[Notification], error) {
	var out Page[Notification]
	if err := c.do(ctx, "GET", "/v1/partner/notifications", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetPartnerNotificationsEach yields every item of GetPartnerNotifications, following the pages.
func (c *Client) GetPartnerNotificationsEach(ctx context.Context, params *GetPartnerNotificationsParams, opts ...RequestOption) iter.Seq2[Notification, error] {
	return paginate(func(cursor *string) (*Page[Notification], error) {
		p := GetPartnerNotificationsParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetPartnerNotifications(ctx, &p, opts...)
	})
}

// GetPartnerOrgsOrgIDSelfBilling calls Get self-billing status. GET /v1/partner/orgs/{orgId}/self-billing.
//
// The partner's VAT details with this merchant and whether they accepted the current self-billing agreement (`currentVersion`).
//
// **Auth:** a partner session.
func (c *Client) GetPartnerOrgsOrgIDSelfBilling(ctx context.Context, orgId string, opts ...RequestOption) (*GetPartnerOrgsOrgIdSelfBillingResponse, error) {
	var out GetPartnerOrgsOrgIdSelfBillingResponse
	if err := c.do(ctx, "GET", "/v1/partner/orgs/"+url.PathEscape(orgId)+"/self-billing", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetPartnerPassport calls Get my reputation passport. GET /v1/partner/passport.
//
// A signed, portable passport of the partner's verified performance across every program (programs, eligible clicks, referrals, repeat-payer rate, average quality), valid for 90 days.
//
// **Auth:** a partner session.
func (c *Client) GetPartnerPassport(ctx context.Context, opts ...RequestOption) (*Passport, error) {
	var out Passport
	if err := c.do(ctx, "GET", "/v1/partner/passport", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetPartnerPayoutsParams are the query parameters of GetPartnerPayouts.
type GetPartnerPayoutsParams struct {
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetPartnerPayoutsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetPartnerPayouts calls List my payouts. GET /v1/partner/payouts.
//
// The partner's payouts across merchants, newest first.
//
// **Auth:** a partner session.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetPartnerPayouts(ctx context.Context, params *GetPartnerPayoutsParams, opts ...RequestOption) (*Page[GetPartnerPayoutsResponseDataItem], error) {
	var out Page[GetPartnerPayoutsResponseDataItem]
	if err := c.do(ctx, "GET", "/v1/partner/payouts", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetPartnerPayoutsEach yields every item of GetPartnerPayouts, following the pages.
func (c *Client) GetPartnerPayoutsEach(ctx context.Context, params *GetPartnerPayoutsParams, opts ...RequestOption) iter.Seq2[GetPartnerPayoutsResponseDataItem, error] {
	return paginate(func(cursor *string) (*Page[GetPartnerPayoutsResponseDataItem], error) {
		p := GetPartnerPayoutsParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetPartnerPayouts(ctx, &p, opts...)
	})
}

// GetPartnerPayoutsIDInvoiceParams are the query parameters of GetPartnerPayoutsIDInvoice.
type GetPartnerPayoutsIDInvoiceParams struct {
	Format *string `json:"format,omitempty"`
}

func (p *GetPartnerPayoutsIDInvoiceParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "format", p.Format)
	return v
}

// GetPartnerPayoutsIDInvoice calls Get my payout invoice. GET /v1/partner/payouts/{id}/invoice.
//
// The invoice for one of the partner's payouts, as JSON or (`format=html`) a printable page.
//
// **Auth:** a partner session.
func (c *Client) GetPartnerPayoutsIDInvoice(ctx context.Context, id string, params *GetPartnerPayoutsIDInvoiceParams, opts ...RequestOption) (*Invoice, error) {
	var out Invoice
	if err := c.do(ctx, "GET", "/v1/partner/payouts/"+url.PathEscape(id)+"/invoice", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetPartnerPostbackLogsParams are the query parameters of GetPartnerPostbackLogs.
type GetPartnerPostbackLogsParams struct {
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetPartnerPostbackLogsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetPartnerPostbackLogs calls List my postback calls. GET /v1/partner/postback-logs.
//
// Recent postback calls with the topic, URL and the status code the partner's server answered.
//
// **Auth:** a partner session.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetPartnerPostbackLogs(ctx context.Context, params *GetPartnerPostbackLogsParams, opts ...RequestOption) (*Page[GetPartnerPostbackLogsResponseDataItem], error) {
	var out Page[GetPartnerPostbackLogsResponseDataItem]
	if err := c.do(ctx, "GET", "/v1/partner/postback-logs", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetPartnerPostbackLogsEach yields every item of GetPartnerPostbackLogs, following the pages.
func (c *Client) GetPartnerPostbackLogsEach(ctx context.Context, params *GetPartnerPostbackLogsParams, opts ...RequestOption) iter.Seq2[GetPartnerPostbackLogsResponseDataItem, error] {
	return paginate(func(cursor *string) (*Page[GetPartnerPostbackLogsResponseDataItem], error) {
		p := GetPartnerPostbackLogsParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetPartnerPostbackLogs(ctx, &p, opts...)
	})
}

// GetPartnerProfile calls Get my creator profile. GET /v1/partner/profile.
//
// The partner's creator profile (headline, bio, channels, audience, engagement) and verified stats.
//
// **Auth:** a partner session.
func (c *Client) GetPartnerProfile(ctx context.Context, opts ...RequestOption) (*CreatorProfile, error) {
	var out CreatorProfile
	if err := c.do(ctx, "GET", "/v1/partner/profile", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetPartnerProfileVerified calls Get my verified performance. GET /v1/partner/profile/verified.
//
// The partner's performance across all programs, computed by the platform from the ledger rather than self-reported: programs, eligible clicks, active referrals and repeat-payer rate.
//
// **Auth:** a partner session.
func (c *Client) GetPartnerProfileVerified(ctx context.Context, opts ...RequestOption) (*GetPartnerProfileVerifiedResponse, error) {
	var out GetPartnerProfileVerifiedResponse
	if err := c.do(ctx, "GET", "/v1/partner/profile/verified", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetPartnerProgramsProgramIDLeaderboard calls Get a program leaderboard. GET /v1/partner/programs/{programId}/leaderboard.
//
// Top 25 partners by approved commission in the program (during the active contest, if any). Other partners' names are masked; the partner's own row is marked `you`. 403 when the program has no leaderboard.
//
// **Auth:** a partner session.
func (c *Client) GetPartnerProgramsProgramIDLeaderboard(ctx context.Context, programId string, opts ...RequestOption) (*GetPartnerProgramsProgramIdLeaderboardResponse, error) {
	var out GetPartnerProgramsProgramIdLeaderboardResponse
	if err := c.do(ctx, "GET", "/v1/partner/programs/"+url.PathEscape(programId)+"/leaderboard", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetPartnerRecruit calls Get my recruiting link. GET /v1/partner/recruit.
//
// The partner's recruit code and sign-up URL. Partners who sign up through it become their sub-affiliates where both partner the same merchant.
//
// **Auth:** a partner session.
func (c *Client) GetPartnerRecruit(ctx context.Context, opts ...RequestOption) (*GetPartnerRecruitResponse, error) {
	var out GetPartnerRecruitResponse
	if err := c.do(ctx, "GET", "/v1/partner/recruit", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetPartnerReferralsParams are the query parameters of GetPartnerReferrals.
type GetPartnerReferralsParams struct {
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetPartnerReferralsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetPartnerReferrals calls List my referrals. GET /v1/partner/referrals.
//
// The partner's referrals across merchants, newest first. Customer ids are redacted to their last 4 characters.
//
// **Auth:** a partner session.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetPartnerReferrals(ctx context.Context, params *GetPartnerReferralsParams, opts ...RequestOption) (*Page[GetPartnerReferralsResponseDataItem], error) {
	var out Page[GetPartnerReferralsResponseDataItem]
	if err := c.do(ctx, "GET", "/v1/partner/referrals", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetPartnerReferralsEach yields every item of GetPartnerReferrals, following the pages.
func (c *Client) GetPartnerReferralsEach(ctx context.Context, params *GetPartnerReferralsParams, opts ...RequestOption) iter.Seq2[GetPartnerReferralsResponseDataItem, error] {
	return paginate(func(cursor *string) (*Page[GetPartnerReferralsResponseDataItem], error) {
		p := GetPartnerReferralsParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetPartnerReferrals(ctx, &p, opts...)
	})
}

// GetPartnerReferralsIDTransparency calls Explain a referral. GET /v1/partner/referrals/{id}/transparency.
//
// Why a referral credited the partner (or did not fully): the model, window, touchpoints considered, the partner's share of the credit and the reasoning notes. No customer data.
//
// **Auth:** a partner session.
func (c *Client) GetPartnerReferralsIDTransparency(ctx context.Context, id string, opts ...RequestOption) (*GetPartnerReferralsIdTransparencyResponse, error) {
	var out GetPartnerReferralsIdTransparencyResponse
	if err := c.do(ctx, "GET", "/v1/partner/referrals/"+url.PathEscape(id)+"/transparency", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetPartnersParams are the query parameters of GetPartners.
type GetPartnersParams struct {
	Email           *string `json:"email,omitempty"`
	ExternalRef     *string `json:"externalRef,omitempty"`
	Status          *string `json:"status,omitempty"`
	IncludeArchived *string `json:"includeArchived,omitempty"`
	Limit           *int64  `json:"limit,omitempty"`
	StartingAfter   *string `json:"starting_after,omitempty"`
	EndingBefore    *string `json:"ending_before,omitempty"`
	CreatedGte      *string `json:"created[gte],omitempty"`
	CreatedLte      *string `json:"created[lte],omitempty"`
}

func (p *GetPartnersParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "email", p.Email)
	setQuery(v, "externalRef", p.ExternalRef)
	setQuery(v, "status", p.Status)
	setQuery(v, "includeArchived", p.IncludeArchived)
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetPartners calls List partners. GET /v1/partners.
//
// Partners, newest first. Filter by `email`, `externalRef` or `status`. Archived partners are hidden unless you ask for them with `status=archived` or `includeArchived=true`.
//
// **Auth:** an API key with `partners:read` (or `full`), or a session of any role.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetPartners(ctx context.Context, params *GetPartnersParams, opts ...RequestOption) (*Page[Partner], error) {
	var out Page[Partner]
	if err := c.do(ctx, "GET", "/v1/partners", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetPartnersEach yields every item of GetPartners, following the pages.
func (c *Client) GetPartnersEach(ctx context.Context, params *GetPartnersParams, opts ...RequestOption) iter.Seq2[Partner, error] {
	return paginate(func(cursor *string) (*Page[Partner], error) {
		p := GetPartnersParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetPartners(ctx, &p, opts...)
	})
}

// GetPartnerSecurity calls Get my security settings. GET /v1/partner/security.
//
// Whether two-factor sign-in is on for the partner account.
//
// **Auth:** a partner session.
func (c *Client) GetPartnerSecurity(ctx context.Context, opts ...RequestOption) (*GetPartnerSecurityResponse, error) {
	var out GetPartnerSecurityResponse
	if err := c.do(ctx, "GET", "/v1/partner/security", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetPartnersID calls Get a partner. GET /v1/partners/{id}.
//
// One partner, with a cross-network warning when the same account is flagged or banned by other merchants (it never says which).
//
// **Auth:** an API key with `partners:read` (or `full`), or a session of any role.
func (c *Client) GetPartnersID(ctx context.Context, id string, opts ...RequestOption) (*GetPartnersIdResponse, error) {
	var out GetPartnersIdResponse
	if err := c.do(ctx, "GET", "/v1/partners/"+url.PathEscape(id), nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetPartnersIDBalancesParams are the query parameters of GetPartnersIDBalances.
type GetPartnersIDBalancesParams struct {
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetPartnersIDBalancesParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetPartnersIDBalances calls Get a partner's balances. GET /v1/partners/{id}/balances.
//
// The partner's commission balances per currency: pending (in hold), approved (payable), paid and reversed.
//
// **Auth:** an API key with `commissions:read` (or `full`), or a session of any role.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetPartnersIDBalances(ctx context.Context, id string, params *GetPartnersIDBalancesParams, opts ...RequestOption) (*Page[Balance], error) {
	var out Page[Balance]
	if err := c.do(ctx, "GET", "/v1/partners/"+url.PathEscape(id)+"/balances", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetPartnersIDBalancesEach yields every item of GetPartnersIDBalances, following the pages.
func (c *Client) GetPartnersIDBalancesEach(ctx context.Context, id string, params *GetPartnersIDBalancesParams, opts ...RequestOption) iter.Seq2[Balance, error] {
	return paginate(func(cursor *string) (*Page[Balance], error) {
		p := GetPartnersIDBalancesParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetPartnersIDBalances(ctx, id, &p, opts...)
	})
}

// GetPartnersIDEnrollmentsParams are the query parameters of GetPartnersIDEnrollments.
type GetPartnersIDEnrollmentsParams struct {
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetPartnersIDEnrollmentsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetPartnersIDEnrollments calls List a partner's enrollments. GET /v1/partners/{id}/enrollments.
//
// The programs a partner is enrolled in or applied to, with status and group.
//
// **Auth:** an API key with `partners:read` (or `full`), or a session of any role.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetPartnersIDEnrollments(ctx context.Context, id string, params *GetPartnersIDEnrollmentsParams, opts ...RequestOption) (*Page[GetPartnersIdEnrollmentsResponseDataItem], error) {
	var out Page[GetPartnersIdEnrollmentsResponseDataItem]
	if err := c.do(ctx, "GET", "/v1/partners/"+url.PathEscape(id)+"/enrollments", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetPartnersIDEnrollmentsEach yields every item of GetPartnersIDEnrollments, following the pages.
func (c *Client) GetPartnersIDEnrollmentsEach(ctx context.Context, id string, params *GetPartnersIDEnrollmentsParams, opts ...RequestOption) iter.Seq2[GetPartnersIdEnrollmentsResponseDataItem, error] {
	return paginate(func(cursor *string) (*Page[GetPartnersIdEnrollmentsResponseDataItem], error) {
		p := GetPartnersIDEnrollmentsParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetPartnersIDEnrollments(ctx, id, &p, opts...)
	})
}

// GetPartnersIDPayoutProfile calls Get a partner's payout profile. GET /v1/partners/{id}/payout-profile.
//
// The partner's payout method, currency, threshold and tax form status. Method details are not returned. 404 when none is set.
//
// **Auth:** an API key with `payouts:read` (or `full`), or a session of any role.
func (c *Client) GetPartnersIDPayoutProfile(ctx context.Context, id string, opts ...RequestOption) (*PayoutProfile, error) {
	var out PayoutProfile
	if err := c.do(ctx, "GET", "/v1/partners/"+url.PathEscape(id)+"/payout-profile", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetPartnersIDQuality calls Get a partner's quality score. GET /v1/partners/{id}/quality.
//
// The partner's quality score computed now (with its metrics) and the last stored snapshot.
//
// **Auth:** an API key with `partners:read` (or `full`), or a session of any role.
func (c *Client) GetPartnersIDQuality(ctx context.Context, id string, opts ...RequestOption) (*GetPartnersIdQualityResponse, error) {
	var out GetPartnersIdQualityResponse
	if err := c.do(ctx, "GET", "/v1/partners/"+url.PathEscape(id)+"/quality", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetPartnerStatsParams are the query parameters of GetPartnerStats.
type GetPartnerStatsParams struct {
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetPartnerStatsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetPartnerStats calls Get my stats. GET /v1/partner/stats.
//
// Per merchant: eligible clicks, active referrals, payments and commission balances.
//
// **Auth:** a partner session.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetPartnerStats(ctx context.Context, params *GetPartnerStatsParams, opts ...RequestOption) (*Page[GetPartnerStatsResponseDataItem], error) {
	var out Page[GetPartnerStatsResponseDataItem]
	if err := c.do(ctx, "GET", "/v1/partner/stats", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetPartnerStatsEach yields every item of GetPartnerStats, following the pages.
func (c *Client) GetPartnerStatsEach(ctx context.Context, params *GetPartnerStatsParams, opts ...RequestOption) iter.Seq2[GetPartnerStatsResponseDataItem, error] {
	return paginate(func(cursor *string) (*Page[GetPartnerStatsResponseDataItem], error) {
		p := GetPartnerStatsParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetPartnerStats(ctx, &p, opts...)
	})
}

// GetPartnerTaxDocumentsParams are the query parameters of GetPartnerTaxDocuments.
type GetPartnerTaxDocumentsParams struct {
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetPartnerTaxDocumentsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetPartnerTaxDocuments calls List my tax documents. GET /v1/partner/tax-documents.
//
// Tax forms delivered to the partner by every merchant they work with, newest year first.
//
// **Auth:** a partner session.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetPartnerTaxDocuments(ctx context.Context, params *GetPartnerTaxDocumentsParams, opts ...RequestOption) (*Page[GetPartnerTaxDocumentsResponseDataItem], error) {
	var out Page[GetPartnerTaxDocumentsResponseDataItem]
	if err := c.do(ctx, "GET", "/v1/partner/tax-documents", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetPartnerTaxDocumentsEach yields every item of GetPartnerTaxDocuments, following the pages.
func (c *Client) GetPartnerTaxDocumentsEach(ctx context.Context, params *GetPartnerTaxDocumentsParams, opts ...RequestOption) iter.Seq2[GetPartnerTaxDocumentsResponseDataItem, error] {
	return paginate(func(cursor *string) (*Page[GetPartnerTaxDocumentsResponseDataItem], error) {
		p := GetPartnerTaxDocumentsParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetPartnerTaxDocuments(ctx, &p, opts...)
	})
}

// GetPartnerTaxDocumentsID calls Get my tax document. GET /v1/partner/tax-documents/{id}.
//
// The recipient copy of a delivered tax form, as an HTML page.
//
// **Auth:** a partner session.
func (c *Client) GetPartnerTaxDocumentsID(ctx context.Context, id string, opts ...RequestOption) (string, error) {
	var out string
	err := c.do(ctx, "GET", "/v1/partner/tax-documents/"+url.PathEscape(id), nil, nil, &out, "text", opts)
	return out, err
}

// GetPartnerTermsParams are the query parameters of GetPartnerTerms.
type GetPartnerTermsParams struct {
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetPartnerTermsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetPartnerTerms calls List program terms. GET /v1/partner/terms.
//
// The latest terms version of each program the partner is approved in, and whether they accepted it.
//
// **Auth:** a partner session.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetPartnerTerms(ctx context.Context, params *GetPartnerTermsParams, opts ...RequestOption) (*Page[GetPartnerTermsResponseDataItem], error) {
	var out Page[GetPartnerTermsResponseDataItem]
	if err := c.do(ctx, "GET", "/v1/partner/terms", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetPartnerTermsEach yields every item of GetPartnerTerms, following the pages.
func (c *Client) GetPartnerTermsEach(ctx context.Context, params *GetPartnerTermsParams, opts ...RequestOption) iter.Seq2[GetPartnerTermsResponseDataItem, error] {
	return paginate(func(cursor *string) (*Page[GetPartnerTermsResponseDataItem], error) {
		p := GetPartnerTermsParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetPartnerTerms(ctx, &p, opts...)
	})
}

// GetPassportPublicKey calls Get the passport signing key. GET /v1/passport/public-key.
//
// The ed25519 public key partner passports are signed with, so anyone can verify one offline.
//
// **Auth:** none.
func (c *Client) GetPassportPublicKey(ctx context.Context, opts ...RequestOption) (*GetPassportPublicKeyResponse, error) {
	var out GetPassportPublicKeyResponse
	if err := c.do(ctx, "GET", "/v1/passport/public-key", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetPayoutConnectionsParams are the query parameters of GetPayoutConnections.
type GetPayoutConnectionsParams struct {
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetPayoutConnectionsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetPayoutConnections calls List payout provider connections. GET /v1/payout-connections.
//
// The payout providers the org connected with its own credentials (PayPal, Stripe, Circle, Wise, Payoneer, Tremendous), with environment and status. Secrets are never returned.
//
// **Auth:** an API key with `payouts:read` (or `full`), or a session with role `owner`, `admin`, `finance`.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetPayoutConnections(ctx context.Context, params *GetPayoutConnectionsParams, opts ...RequestOption) (*Page[GetPayoutConnectionsResponseDataItem], error) {
	var out Page[GetPayoutConnectionsResponseDataItem]
	if err := c.do(ctx, "GET", "/v1/payout-connections", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetPayoutConnectionsEach yields every item of GetPayoutConnections, following the pages.
func (c *Client) GetPayoutConnectionsEach(ctx context.Context, params *GetPayoutConnectionsParams, opts ...RequestOption) iter.Seq2[GetPayoutConnectionsResponseDataItem, error] {
	return paginate(func(cursor *string) (*Page[GetPayoutConnectionsResponseDataItem], error) {
		p := GetPayoutConnectionsParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetPayoutConnections(ctx, &p, opts...)
	})
}

// GetPayoutRunsParams are the query parameters of GetPayoutRuns.
type GetPayoutRunsParams struct {
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetPayoutRunsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetPayoutRuns calls List payout runs. GET /v1/payout-runs.
//
// Payout runs, newest first, with their status and summary.
//
// **Auth:** an API key with `payouts:read` (or `full`), or a session of any role.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetPayoutRuns(ctx context.Context, params *GetPayoutRunsParams, opts ...RequestOption) (*Page[PayoutRun], error) {
	var out Page[PayoutRun]
	if err := c.do(ctx, "GET", "/v1/payout-runs", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetPayoutRunsEach yields every item of GetPayoutRuns, following the pages.
func (c *Client) GetPayoutRunsEach(ctx context.Context, params *GetPayoutRunsParams, opts ...RequestOption) iter.Seq2[PayoutRun, error] {
	return paginate(func(cursor *string) (*Page[PayoutRun], error) {
		p := GetPayoutRunsParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetPayoutRuns(ctx, &p, opts...)
	})
}

// GetPayoutsParams are the query parameters of GetPayouts.
type GetPayoutsParams struct {
	PartnerID     *string `json:"partnerId,omitempty"`
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetPayoutsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "partnerId", p.PartnerID)
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetPayouts calls List payouts. GET /v1/payouts.
//
// Individual payouts, newest first, with rail, provider reference and failure reason. Filter with `partnerId`.
//
// **Auth:** an API key with `payouts:read` (or `full`), or a session of any role.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetPayouts(ctx context.Context, params *GetPayoutsParams, opts ...RequestOption) (*Page[Payout], error) {
	var out Page[Payout]
	if err := c.do(ctx, "GET", "/v1/payouts", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetPayoutsEach yields every item of GetPayouts, following the pages.
func (c *Client) GetPayoutsEach(ctx context.Context, params *GetPayoutsParams, opts ...RequestOption) iter.Seq2[Payout, error] {
	return paginate(func(cursor *string) (*Page[Payout], error) {
		p := GetPayoutsParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetPayouts(ctx, &p, opts...)
	})
}

// GetPayoutsIDInvoiceParams are the query parameters of GetPayoutsIDInvoice.
type GetPayoutsIDInvoiceParams struct {
	Format *string `json:"format,omitempty"`
}

func (p *GetPayoutsIDInvoiceParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "format", p.Format)
	return v
}

// GetPayoutsIDInvoice calls Get a payout's invoice. GET /v1/payouts/{id}/invoice.
//
// The self-billing invoice issued for a payout, as JSON or (`format=html`) a printable page. 404 when the payout has no invoice.
//
// **Auth:** an API key with `payouts:read` (or `full`), or a session of any role.
func (c *Client) GetPayoutsIDInvoice(ctx context.Context, id string, params *GetPayoutsIDInvoiceParams, opts ...RequestOption) (*Invoice, error) {
	var out Invoice
	if err := c.do(ctx, "GET", "/v1/payouts/"+url.PathEscape(id)+"/invoice", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetPlans calls List public plans. GET /v1/plans.
//
// The public plan catalog in display order: monthly price in minor units (USD, billed monthly in advance), limits and features. No authentication; cacheable for an hour. No plan charges a transaction or payout fee.
//
// **Auth:** none.
func (c *Client) GetPlans(ctx context.Context, opts ...RequestOption) (*GetPlansResponse, error) {
	var out GetPlansResponse
	if err := c.do(ctx, "GET", "/v1/plans", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetPortalAuthCurrentParams are the query parameters of GetPortalAuthCurrent.
type GetPortalAuthCurrentParams struct {
	Org *string `json:"org,omitempty"`
}

func (p *GetPortalAuthCurrentParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "org", p.Org)
	return v
}

// GetPortalAuthCurrent calls What a merchant's portal sign-in offers. GET /v1/portal-auth/current.
//
// For the portal's sign-in page: the org (from `org`, or the merchant's portal domain this is called on) and its partner sign-in adapter, with the button label and where it starts. `referalli` with no start URL when the org uses Referalli partner accounts.
//
// **Auth:** none.
func (c *Client) GetPortalAuthCurrent(ctx context.Context, params *GetPortalAuthCurrentParams, opts ...RequestOption) (*GetPortalAuthCurrentResponse, error) {
	var out GetPortalAuthCurrentResponse
	if err := c.do(ctx, "GET", "/v1/portal-auth/current", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetPrivacyExecutorRequests calls List DSAR tool requests. GET /v1/privacy/executor-requests.
//
// The last 200 requests connected DSAR tools sent: operation, outcome and the vendor request id. Subjects are kept only as hashes.
//
// **Auth:** an API key with `privacy:read` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) GetPrivacyExecutorRequests(ctx context.Context, opts ...RequestOption) (*GetPrivacyExecutorRequestsResponse, error) {
	var out GetPrivacyExecutorRequestsResponse
	if err := c.do(ctx, "GET", "/v1/privacy/executor-requests", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetPrivacyExecutors calls List connected DSAR tools. GET /v1/privacy/executors.
//
// Every DSAR tool that can reach Referalli as a connected system (Cookie Munch cm-standard-v1, OneTrust, Transcend, DataGrail, Ketch, Osano, `custom` for any tool speaking Standard Webhooks), with the per-org URL to register, install steps and a health check. Keys are never returned.
//
// **Auth:** an API key with `privacy:read` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) GetPrivacyExecutors(ctx context.Context, opts ...RequestOption) (*GetPrivacyExecutorsResponse, error) {
	var out GetPrivacyExecutorsResponse
	if err := c.do(ctx, "GET", "/v1/privacy/executors", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetPrivacyExportParams are the query parameters of GetPrivacyExport.
type GetPrivacyExportParams struct {
	CustomerRef *string `json:"customerRef,omitempty"`
}

func (p *GetPrivacyExportParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "customerRef", p.CustomerRef)
	return v
}

// GetPrivacyExport calls Export a customer's data. GET /v1/privacy/export.
//
// Every stored row that references the customer, for a data-subject access request.
//
// **Auth:** an API key with `privacy:read` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) GetPrivacyExport(ctx context.Context, params *GetPrivacyExportParams, opts ...RequestOption) (*CustomerExport, error) {
	var out CustomerExport
	if err := c.do(ctx, "GET", "/v1/privacy/export", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetPrivacyRequestsParams are the query parameters of GetPrivacyRequests.
type GetPrivacyRequestsParams struct {
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetPrivacyRequestsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetPrivacyRequests calls List data-subject requests. GET /v1/privacy/requests.
//
// Access and erasure requests, newest first, with their status.
//
// **Auth:** an API key with `privacy:read` (or `full`), or a session with role `owner`, `admin`.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetPrivacyRequests(ctx context.Context, params *GetPrivacyRequestsParams, opts ...RequestOption) (*Page[PrivacyRequest], error) {
	var out Page[PrivacyRequest]
	if err := c.do(ctx, "GET", "/v1/privacy/requests", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetPrivacyRequestsEach yields every item of GetPrivacyRequests, following the pages.
func (c *Client) GetPrivacyRequestsEach(ctx context.Context, params *GetPrivacyRequestsParams, opts ...RequestOption) iter.Seq2[PrivacyRequest, error] {
	return paginate(func(cursor *string) (*Page[PrivacyRequest], error) {
		p := GetPrivacyRequestsParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetPrivacyRequests(ctx, &p, opts...)
	})
}

// GetPrivacyRequestsID calls Get a data-subject request. GET /v1/privacy/requests/{id}.
//
// Poll one request: `pending`, `completed` or `failed`, and whether data was found.
//
// **Auth:** an API key with `privacy:read` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) GetPrivacyRequestsID(ctx context.Context, id string, opts ...RequestOption) (*PrivacyRequest, error) {
	var out PrivacyRequest
	if err := c.do(ctx, "GET", "/v1/privacy/requests/"+url.PathEscape(id), nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetPrivacyRequestsIDResult calls Get a data-subject request result. GET /v1/privacy/requests/{id}/result.
//
// The full result of a completed request; for access requests, the exported data. 400 until it is completed.
//
// **Auth:** an API key with `privacy:read` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) GetPrivacyRequestsIDResult(ctx context.Context, id string, opts ...RequestOption) (*GetPrivacyRequestsIdResultResponse, error) {
	var out GetPrivacyRequestsIdResultResponse
	if err := c.do(ctx, "GET", "/v1/privacy/requests/"+url.PathEscape(id)+"/result", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetProfitabilityProducts calls List product margins. GET /v1/profitability/products.
//
// The margin catalog: per product, a margin (bps) or a unit cost, used to turn revenue into profit.
//
// **Auth:** an API key with `analytics:read` (or `full`), or a session of any role.
func (c *Client) GetProfitabilityProducts(ctx context.Context, opts ...RequestOption) (*GetProfitabilityProductsResponse, error) {
	var out GetProfitabilityProductsResponse
	if err := c.do(ctx, "GET", "/v1/profitability/products", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetProfitabilityProductsProductRef calls Get a product margin. GET /v1/profitability/products/{productRef}.
//
// One product of the margin catalog.
//
// **Auth:** an API key with `analytics:read` (or `full`), or a session of any role.
func (c *Client) GetProfitabilityProductsProductRef(ctx context.Context, productRef string, opts ...RequestOption) (*GetProfitabilityProductsProductRefResponse, error) {
	var out GetProfitabilityProductsProductRefResponse
	if err := c.do(ctx, "GET", "/v1/profitability/products/"+url.PathEscape(productRef), nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetProfitabilityRebuildJobID calls Get a rebuild job. GET /v1/profitability/rebuild/{jobId}.
//
// Progress of a rebuild: `queued`, `running`, `done` or `failed` (a job untouched for an hour reports `failed` with `error: "stale"`), with processed and total counts.
//
// **Auth:** an API key with `analytics:read` (or `full`), or a session of any role.
func (c *Client) GetProfitabilityRebuildJobID(ctx context.Context, jobId string, opts ...RequestOption) (*GetProfitabilityRebuildJobIdResponse, error) {
	var out GetProfitabilityRebuildJobIdResponse
	if err := c.do(ctx, "GET", "/v1/profitability/rebuild/"+url.PathEscape(jobId), nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetProfitabilitySettingsParams are the query parameters of GetProfitabilitySettings.
type GetProfitabilitySettingsParams struct {
	ProgramID *string `json:"programId,omitempty"`
}

func (p *GetProfitabilitySettingsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "programId", p.ProgramID)
	return v
}

// GetProfitabilitySettings calls Get profitability settings. GET /v1/profitability/settings.
//
// The economics credit ROI is computed with: reporting currency, default margin, how a credit's cost is modelled, how credit use is inferred, attribution and churn windows, target ROI and more. With `programId`, the program's overrides merged over the org's.
//
// **Auth:** an API key with `analytics:read` (or `full`), or a session of any role.
func (c *Client) GetProfitabilitySettings(ctx context.Context, params *GetProfitabilitySettingsParams, opts ...RequestOption) (*GetProfitabilitySettingsResponse, error) {
	var out GetProfitabilitySettingsResponse
	if err := c.do(ctx, "GET", "/v1/profitability/settings", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetProgramsParams are the query parameters of GetPrograms.
type GetProgramsParams struct {
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetProgramsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetPrograms calls List programs. GET /v1/programs.
//
// The org's programs, oldest first, with their settings.
//
// **Auth:** an API key with `programs:read` (or `full`), or a session of any role.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetPrograms(ctx context.Context, params *GetProgramsParams, opts ...RequestOption) (*Page[Program], error) {
	var out Page[Program]
	if err := c.do(ctx, "GET", "/v1/programs", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetProgramsEach yields every item of GetPrograms, following the pages.
func (c *Client) GetProgramsEach(ctx context.Context, params *GetProgramsParams, opts ...RequestOption) iter.Seq2[Program, error] {
	return paginate(func(cursor *string) (*Page[Program], error) {
		p := GetProgramsParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetPrograms(ctx, &p, opts...)
	})
}

// GetProgramsID calls Get a program. GET /v1/programs/{id}.
//
// One program with its settings (defaults filled in) and public page.
//
// **Auth:** an API key with `programs:read` (or `full`), or a session of any role.
func (c *Client) GetProgramsID(ctx context.Context, id string, opts ...RequestOption) (*Program, error) {
	var out Program
	if err := c.do(ctx, "GET", "/v1/programs/"+url.PathEscape(id), nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetProgramsIDAutopricing calls Get autopricing settings. GET /v1/programs/{id}/autopricing.
//
// The program's autopricing configuration, or `null` when it's off.
//
// **Auth:** an API key with `programs:read` (or `full`), or a session of any role.
func (c *Client) GetProgramsIDAutopricing(ctx context.Context, id string, opts ...RequestOption) (*GetProgramsIdAutopricingResponse, error) {
	var out GetProgramsIdAutopricingResponse
	if err := c.do(ctx, "GET", "/v1/programs/"+url.PathEscape(id)+"/autopricing", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetProgramsIDEnrollmentsParams are the query parameters of GetProgramsIDEnrollments.
type GetProgramsIDEnrollmentsParams struct {
	Status        *string `json:"status,omitempty"`
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetProgramsIDEnrollmentsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "status", p.Status)
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetProgramsIDEnrollments calls List program enrollments. GET /v1/programs/{id}/enrollments.
//
// The program's enrollments and applications, newest first, each with the partner and a cross-network warning when other merchants banned the same account. Filter with `status`.
//
// **Auth:** an API key with `programs:read` (or `full`), or a session of any role.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetProgramsIDEnrollments(ctx context.Context, id string, params *GetProgramsIDEnrollmentsParams, opts ...RequestOption) (*Page[GetProgramsIdEnrollmentsResponseDataItem], error) {
	var out Page[GetProgramsIdEnrollmentsResponseDataItem]
	if err := c.do(ctx, "GET", "/v1/programs/"+url.PathEscape(id)+"/enrollments", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetProgramsIDEnrollmentsEach yields every item of GetProgramsIDEnrollments, following the pages.
func (c *Client) GetProgramsIDEnrollmentsEach(ctx context.Context, id string, params *GetProgramsIDEnrollmentsParams, opts ...RequestOption) iter.Seq2[GetProgramsIdEnrollmentsResponseDataItem, error] {
	return paginate(func(cursor *string) (*Page[GetProgramsIdEnrollmentsResponseDataItem], error) {
		p := GetProgramsIDEnrollmentsParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetProgramsIDEnrollments(ctx, id, &p, opts...)
	})
}

// GetProgramsIDGraphParams are the query parameters of GetProgramsIDGraph.
type GetProgramsIDGraphParams struct {
	Root  *string `json:"root,omitempty"`
	Depth *int64  `json:"depth,omitempty"`
}

func (p *GetProgramsIDGraphParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "root", p.Root)
	setQuery(v, "depth", p.Depth)
	return v
}

// GetProgramsIDGraph calls Get the referral tree. GET /v1/programs/{id}/graph.
//
// Who referred whom from one partner (`root`) down, with revenue and net commission per node. Where customers are partners too (customerRef equals a partner's externalRef), the tree continues up to `depth` partner levels. Stops at 5,000 referrals and says so.
//
// **Auth:** an API key with `programs:read` (or `full`), or a session of any role.
func (c *Client) GetProgramsIDGraph(ctx context.Context, id string, params *GetProgramsIDGraphParams, opts ...RequestOption) (*GetProgramsIdGraphResponse, error) {
	var out GetProgramsIdGraphResponse
	if err := c.do(ctx, "GET", "/v1/programs/"+url.PathEscape(id)+"/graph", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetProgramsIDGroupsParams are the query parameters of GetProgramsIDGroups.
type GetProgramsIDGroupsParams struct {
	IncludeArchived *string `json:"includeArchived,omitempty"`
	Limit           *int64  `json:"limit,omitempty"`
	StartingAfter   *string `json:"starting_after,omitempty"`
	EndingBefore    *string `json:"ending_before,omitempty"`
	CreatedGte      *string `json:"created[gte],omitempty"`
	CreatedLte      *string `json:"created[lte],omitempty"`
}

func (p *GetProgramsIDGroupsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "includeArchived", p.IncludeArchived)
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetProgramsIDGroups calls List partner groups. GET /v1/programs/{id}/groups.
//
// The program's partner groups with member counts. Archived groups are hidden unless `includeArchived=true`.
//
// **Auth:** an API key with `programs:read` (or `full`), or a session of any role.
//
// includeArchived: true: include archived groups.
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetProgramsIDGroups(ctx context.Context, id string, params *GetProgramsIDGroupsParams, opts ...RequestOption) (*Page[PartnerGroup], error) {
	var out Page[PartnerGroup]
	if err := c.do(ctx, "GET", "/v1/programs/"+url.PathEscape(id)+"/groups", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetProgramsIDGroupsEach yields every item of GetProgramsIDGroups, following the pages.
func (c *Client) GetProgramsIDGroupsEach(ctx context.Context, id string, params *GetProgramsIDGroupsParams, opts ...RequestOption) iter.Seq2[PartnerGroup, error] {
	return paginate(func(cursor *string) (*Page[PartnerGroup], error) {
		p := GetProgramsIDGroupsParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetProgramsIDGroups(ctx, id, &p, opts...)
	})
}

// GetProgramsIDNegotiationPolicy calls Get the agent negotiation policy. GET /v1/programs/{id}/negotiation-policy.
//
// The program's negotiation policy, or `null` when it isn't open to negotiation.
//
// **Auth:** an API key with `programs:read` (or `full`), or a session of any role.
func (c *Client) GetProgramsIDNegotiationPolicy(ctx context.Context, id string, opts ...RequestOption) (*GetProgramsIdNegotiationPolicyResponse, error) {
	var out GetProgramsIdNegotiationPolicyResponse
	if err := c.do(ctx, "GET", "/v1/programs/"+url.PathEscape(id)+"/negotiation-policy", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetProgramsIDPromotionsParams are the query parameters of GetProgramsIDPromotions.
type GetProgramsIDPromotionsParams struct {
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetProgramsIDPromotionsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetProgramsIDPromotions calls List promotions. GET /v1/programs/{id}/promotions.
//
// The program's promotions, latest start first.
//
// **Auth:** an API key with `programs:read` (or `full`), or a session of any role.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetProgramsIDPromotions(ctx context.Context, id string, params *GetProgramsIDPromotionsParams, opts ...RequestOption) (*Page[Promotion], error) {
	var out Page[Promotion]
	if err := c.do(ctx, "GET", "/v1/programs/"+url.PathEscape(id)+"/promotions", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetProgramsIDPromotionsEach yields every item of GetProgramsIDPromotions, following the pages.
func (c *Client) GetProgramsIDPromotionsEach(ctx context.Context, id string, params *GetProgramsIDPromotionsParams, opts ...RequestOption) iter.Seq2[Promotion, error] {
	return paginate(func(cursor *string) (*Page[Promotion], error) {
		p := GetProgramsIDPromotionsParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetProgramsIDPromotions(ctx, id, &p, opts...)
	})
}

// GetProgramsIDRefereeReward calls Get the referee reward. GET /v1/programs/{id}/referee-reward.
//
// The program's reward for referred customers (the friend's side of a double-sided program), or `null` when it has none.
//
// **Auth:** an API key with `programs:read` (or `full`), or a session of any role.
func (c *Client) GetProgramsIDRefereeReward(ctx context.Context, id string, opts ...RequestOption) (*GetProgramsIdRefereeRewardResponse, error) {
	var out GetProgramsIdRefereeRewardResponse
	if err := c.do(ctx, "GET", "/v1/programs/"+url.PathEscape(id)+"/referee-reward", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetProgramsIDTermsAcceptancesParams are the query parameters of GetProgramsIDTermsAcceptances.
type GetProgramsIDTermsAcceptancesParams struct {
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetProgramsIDTermsAcceptancesParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetProgramsIDTermsAcceptances calls List terms acceptances. GET /v1/programs/{id}/terms-acceptances.
//
// Which partner accepted which version of the program's terms, newest first.
//
// **Auth:** an API key with `programs:read` (or `full`), or a session of any role.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetProgramsIDTermsAcceptances(ctx context.Context, id string, params *GetProgramsIDTermsAcceptancesParams, opts ...RequestOption) (*Page[GetProgramsIdTermsAcceptancesResponseDataItem], error) {
	var out Page[GetProgramsIdTermsAcceptancesResponseDataItem]
	if err := c.do(ctx, "GET", "/v1/programs/"+url.PathEscape(id)+"/terms-acceptances", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetProgramsIDTermsAcceptancesEach yields every item of GetProgramsIDTermsAcceptances, following the pages.
func (c *Client) GetProgramsIDTermsAcceptancesEach(ctx context.Context, id string, params *GetProgramsIDTermsAcceptancesParams, opts ...RequestOption) iter.Seq2[GetProgramsIdTermsAcceptancesResponseDataItem, error] {
	return paginate(func(cursor *string) (*Page[GetProgramsIdTermsAcceptancesResponseDataItem], error) {
		p := GetProgramsIDTermsAcceptancesParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetProgramsIDTermsAcceptances(ctx, id, &p, opts...)
	})
}

// GetProspectsParams are the query parameters of GetProspects.
type GetProspectsParams struct {
	Status        *string `json:"status,omitempty"`
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetProspectsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "status", p.Status)
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetProspects calls List prospects. GET /v1/prospects.
//
// Prospects, best fit first, with channel, audience, evidence and outreach status. Filter with `status`.
//
// **Auth:** an API key with `growth:read` (or `full`), or a session of any role.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetProspects(ctx context.Context, params *GetProspectsParams, opts ...RequestOption) (*Page[GetProspectsResponseDataItem], error) {
	var out Page[GetProspectsResponseDataItem]
	if err := c.do(ctx, "GET", "/v1/prospects", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetProspectsEach yields every item of GetProspects, following the pages.
func (c *Client) GetProspectsEach(ctx context.Context, params *GetProspectsParams, opts ...RequestOption) iter.Seq2[GetProspectsResponseDataItem, error] {
	return paginate(func(cursor *string) (*Page[GetProspectsResponseDataItem], error) {
		p := GetProspectsParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetProspects(ctx, &p, opts...)
	})
}

// GetPublicProgramsHandle calls Get a program's public page. GET /v1/public/programs/{handle}.
//
// What a program's public join page shows: the program and merchant names, pitch, cookie window, terms, base commission, branding and whether joining is approved at once. No authentication; only active live programs with a `publicHandle` are found.
//
// **Auth:** none.
func (c *Client) GetPublicProgramsHandle(ctx context.Context, handle string, opts ...RequestOption) (*GetPublicProgramsHandleResponse, error) {
	var out GetPublicProgramsHandleResponse
	if err := c.do(ctx, "GET", "/v1/public/programs/"+url.PathEscape(handle), nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetReferralsParams are the query parameters of GetReferrals.
type GetReferralsParams struct {
	ProgramID     *string `json:"programId,omitempty"`
	PartnerID     *string `json:"partnerId,omitempty"`
	Status        *string `json:"status,omitempty"`
	CustomerRef   *string `json:"customerRef,omitempty"`
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetReferralsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "programId", p.ProgramID)
	setQuery(v, "partnerId", p.PartnerID)
	setQuery(v, "status", p.Status)
	setQuery(v, "customerRef", p.CustomerRef)
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetReferrals calls List referrals. GET /v1/referrals.
//
// Referrals, newest first, filtered by `programId`, `partnerId`, `status` or `customerRef`.
//
// **Auth:** an API key with `referrals:read` (or `full`), or a session of any role.
//
// customerRef: Your id for the referred customer.
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetReferrals(ctx context.Context, params *GetReferralsParams, opts ...RequestOption) (*Page[Referral], error) {
	var out Page[Referral]
	if err := c.do(ctx, "GET", "/v1/referrals", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetReferralsEach yields every item of GetReferrals, following the pages.
func (c *Client) GetReferralsEach(ctx context.Context, params *GetReferralsParams, opts ...RequestOption) iter.Seq2[Referral, error] {
	return paginate(func(cursor *string) (*Page[Referral], error) {
		p := GetReferralsParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetReferrals(ctx, &p, opts...)
	})
}

// GetReferralsID calls Get a referral. GET /v1/referrals/{id}.
//
// One referral with its touch credits and attribution reasoning.
//
// **Auth:** an API key with `referrals:read` (or `full`), or a session of any role.
func (c *Client) GetReferralsID(ctx context.Context, id string, opts ...RequestOption) (*Referral, error) {
	var out Referral
	if err := c.do(ctx, "GET", "/v1/referrals/"+url.PathEscape(id), nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetReferralsIDJourney calls Get the attribution journey. GET /v1/referrals/{id}/journey.
//
// The referral with every touchpoint considered: each click's time, URL, fraud score, eligibility, credit weight and which one won.
//
// **Auth:** an API key with `referrals:read` (or `full`), or a session of any role.
func (c *Client) GetReferralsIDJourney(ctx context.Context, id string, opts ...RequestOption) (*GetReferralsIdJourneyResponse, error) {
	var out GetReferralsIdJourneyResponse
	if err := c.do(ctx, "GET", "/v1/referrals/"+url.PathEscape(id)+"/journey", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetReportsFilteredClicksParams are the query parameters of GetReportsFilteredClicks.
type GetReportsFilteredClicksParams struct {
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetReportsFilteredClicksParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetReportsFilteredClicks calls Report filtered clicks. GET /v1/reports/filtered-clicks.
//
// Per day, newest first: total clicks, clicks excluded from attribution (bots, rate limits, invalid traffic) and the average fraud score.
//
// **Auth:** an API key with `analytics:read` (or `full`), or a session of any role.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetReportsFilteredClicks(ctx context.Context, params *GetReportsFilteredClicksParams, opts ...RequestOption) (*Page[GetReportsFilteredClicksResponseDataItem], error) {
	var out Page[GetReportsFilteredClicksResponseDataItem]
	if err := c.do(ctx, "GET", "/v1/reports/filtered-clicks", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetReportsFilteredClicksEach yields every item of GetReportsFilteredClicks, following the pages.
func (c *Client) GetReportsFilteredClicksEach(ctx context.Context, params *GetReportsFilteredClicksParams, opts ...RequestOption) iter.Seq2[GetReportsFilteredClicksResponseDataItem, error] {
	return paginate(func(cursor *string) (*Page[GetReportsFilteredClicksResponseDataItem], error) {
		p := GetReportsFilteredClicksParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetReportsFilteredClicks(ctx, &p, opts...)
	})
}

// GetReportsPartnerPerformanceParams are the query parameters of GetReportsPartnerPerformance.
type GetReportsPartnerPerformanceParams struct {
	Format *string `json:"format,omitempty"`
}

func (p *GetReportsPartnerPerformanceParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "format", p.Format)
	return v
}

// GetReportsPartnerPerformance calls Get partner performance. GET /v1/reports/partner-performance.
//
// Per partner, highest revenue first: clicks, referrals, payments, referred revenue, commission earned and reversed. `format=csv` returns the same as a CSV file.
//
// **Auth:** an API key with `analytics:read` (or `full`), or a session of any role.
//
// format: csv: the same rows as a CSV file.
func (c *Client) GetReportsPartnerPerformance(ctx context.Context, params *GetReportsPartnerPerformanceParams, opts ...RequestOption) (*GetReportsPartnerPerformanceResponse, error) {
	var out GetReportsPartnerPerformanceResponse
	if err := c.do(ctx, "GET", "/v1/reports/partner-performance", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetRewardsParams are the query parameters of GetRewards.
type GetRewardsParams struct {
	Status        *string `json:"status,omitempty"`
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetRewardsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "status", p.Status)
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetRewards calls List issued referee rewards. GET /v1/rewards.
//
// Rewards issued to referred customers, newest first, with code and status (`issued`, `redeemed`, `void`). Filter with `status`.
//
// **Auth:** an API key with `commissions:read` (or `full`), or a session of any role.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetRewards(ctx context.Context, params *GetRewardsParams, opts ...RequestOption) (*Page[GetRewardsResponseDataItem], error) {
	var out Page[GetRewardsResponseDataItem]
	if err := c.do(ctx, "GET", "/v1/rewards", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetRewardsEach yields every item of GetRewards, following the pages.
func (c *Client) GetRewardsEach(ctx context.Context, params *GetRewardsParams, opts ...RequestOption) iter.Seq2[GetRewardsResponseDataItem, error] {
	return paginate(func(cursor *string) (*Page[GetRewardsResponseDataItem], error) {
		p := GetRewardsParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetRewards(ctx, &p, opts...)
	})
}

// GetRewardsGranted calls List partner rewards. GET /v1/rewards/granted.
//
// The last 200 rewards granted through `POST /v1/rewards`, with the commission line each one posted.
//
// **Auth:** an API key with `commissions:read` (or `full`), or a session of any role.
func (c *Client) GetRewardsGranted(ctx context.Context, opts ...RequestOption) (*GetRewardsGrantedResponse, error) {
	var out GetRewardsGrantedResponse
	if err := c.do(ctx, "GET", "/v1/rewards/granted", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetRewardsLookupCode calls Look up a reward code. GET /v1/rewards/lookup/{code}.
//
// Whether a reward code is valid (issued and unredeemed), with its kind and amount, so your checkout can apply it. Unknown codes answer `{ "valid": false }`.
//
// **Auth:** an API key with `commissions:read` (or `full`), or a session of any role.
func (c *Client) GetRewardsLookupCode(ctx context.Context, code string, opts ...RequestOption) (*GetRewardsLookupCodeResponse, error) {
	var out GetRewardsLookupCodeResponse
	if err := c.do(ctx, "GET", "/v1/rewards/lookup/"+url.PathEscape(code), nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetRewardsRules calls List reward rules. GET /v1/rewards/rules.
//
// Your rules that turn an inbound event (or a rule key) into a fixed partner reward.
//
// **Auth:** an API key with `commissions:read` (or `full`), or a session of any role.
func (c *Client) GetRewardsRules(ctx context.Context, opts ...RequestOption) ([]PutRewardsRulesKeyResponse, error) {
	var out []PutRewardsRulesKeyResponse
	if err := c.do(ctx, "GET", "/v1/rewards/rules", nil, nil, &out, "json", opts); err != nil {
		return out, err
	}
	return out, nil
}

// GetRoiBreakdownParams are the query parameters of GetRoiBreakdown.
type GetRoiBreakdownParams struct {
	From          *string `json:"from,omitempty"`
	To            *string `json:"to,omitempty"`
	ProgramID     *string `json:"programId,omitempty"`
	PartnerID     *string `json:"partnerId,omitempty"`
	WindowDays    *int64  `json:"windowDays,omitempty"`
	Currency      *string `json:"currency,omitempty"`
	Format        *string `json:"format,omitempty"`
	By            *string `json:"by,omitempty"`
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetRoiBreakdownParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "from", p.From)
	setQuery(v, "to", p.To)
	setQuery(v, "programId", p.ProgramID)
	setQuery(v, "partnerId", p.PartnerID)
	setQuery(v, "windowDays", p.WindowDays)
	setQuery(v, "currency", p.Currency)
	setQuery(v, "format", p.Format)
	setQuery(v, "by", p.By)
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetRoiBreakdown calls Get the ROI breakdown. GET /v1/roi/breakdown.
//
// The summary figures grouped `by` program, partner, cohort (grant month), source or role, best ROI first (cohorts in date order). Partner and program rows also carry their referrals' combined figures. `from`/`to` select grants by grant date (a date covers its whole UTC day), `programId` and `partnerId` narrow them, `windowDays` overrides each program's attribution window (30, 90, 180, 365, or 0 for lifetime), and `currency` views one native currency instead of the reporting currency. `format=csv` returns the same as a CSV file.
//
// **Auth:** an API key with `analytics:read` (or `full`), or a session of any role.
//
// from: Grants (or referrals) from this instant or date, inclusive.
// to: Up to this instant, or through the end of this date (UTC), inclusive.
// partnerId: The grant's partner: the referrer, or the partner of the referee's referral.
// windowDays: Attribution window; default each program's setting. 0 = lifetime.
// currency: Default the reporting currency (converted); another code gives that currency only.
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetRoiBreakdown(ctx context.Context, params *GetRoiBreakdownParams, opts ...RequestOption) (*Page[GetRoiBreakdownResponseDataItem], error) {
	var out Page[GetRoiBreakdownResponseDataItem]
	if err := c.do(ctx, "GET", "/v1/roi/breakdown", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetRoiBreakdownEach yields every item of GetRoiBreakdown, following the pages.
func (c *Client) GetRoiBreakdownEach(ctx context.Context, params *GetRoiBreakdownParams, opts ...RequestOption) iter.Seq2[GetRoiBreakdownResponseDataItem, error] {
	return paginate(func(cursor *string) (*Page[GetRoiBreakdownResponseDataItem], error) {
		p := GetRoiBreakdownParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetRoiBreakdown(ctx, &p, opts...)
	})
}

// GetRoiDataQualityParams are the query parameters of GetRoiDataQuality.
type GetRoiDataQualityParams struct {
	From   *string `json:"from,omitempty"`
	To     *string `json:"to,omitempty"`
	Format *string `json:"format,omitempty"`
}

func (p *GetRoiDataQualityParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "from", p.From)
	setQuery(v, "to", p.To)
	setQuery(v, "format", p.Format)
	return v
}

// GetRoiDataQuality calls Get ROI data quality. GET /v1/roi/data-quality.
//
// How complete the inputs to ROI are over `from`/`to`: payments with credit-applied info, products with a margin, partners linked to their own customer id, events excluded for a missing FX rate, and a list of gaps with how to fix each.
//
// **Auth:** an API key with `analytics:read` (or `full`), or a session of any role.
//
// from: Grants (or referrals) from this instant or date, inclusive.
// to: Up to this instant, or through the end of this date (UTC), inclusive.
func (c *Client) GetRoiDataQuality(ctx context.Context, params *GetRoiDataQualityParams, opts ...RequestOption) (*GetRoiDataQualityResponse, error) {
	var out GetRoiDataQualityResponse
	if err := c.do(ctx, "GET", "/v1/roi/data-quality", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetRoiGrantsParams are the query parameters of GetRoiGrants.
type GetRoiGrantsParams struct {
	From          *string `json:"from,omitempty"`
	To            *string `json:"to,omitempty"`
	ProgramID     *string `json:"programId,omitempty"`
	PartnerID     *string `json:"partnerId,omitempty"`
	WindowDays    *int64  `json:"windowDays,omitempty"`
	Currency      *string `json:"currency,omitempty"`
	Format        *string `json:"format,omitempty"`
	Status        *string `json:"status,omitempty"`
	Role          *string `json:"role,omitempty"`
	PaidBack      *string `json:"paidBack,omitempty"`
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetRoiGrantsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "from", p.From)
	setQuery(v, "to", p.To)
	setQuery(v, "programId", p.ProgramID)
	setQuery(v, "partnerId", p.PartnerID)
	setQuery(v, "windowDays", p.WindowDays)
	setQuery(v, "currency", p.Currency)
	setQuery(v, "format", p.Format)
	setQuery(v, "status", p.Status)
	setQuery(v, "role", p.Role)
	setQuery(v, "paidBack", p.PaidBack)
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetRoiGrants calls List credit grants. GET /v1/roi/grants.
//
// Each credit grant in view with its status (`open`, `used`, `expired`, `revoked`) and figures, newest first. Filter by `status`, `role` (`referee` or `referrer`) and `paidBack`. `from`/`to` select grants by grant date (a date covers its whole UTC day), `programId` and `partnerId` narrow them, `windowDays` overrides each program's attribution window (30, 90, 180, 365, or 0 for lifetime), and `currency` views one native currency instead of the reporting currency. `format=csv` returns the same as a CSV file. CSV returns every grant, unpaged.
//
// **Auth:** an API key with `analytics:read` (or `full`), or a session of any role.
//
// from: Grants (or referrals) from this instant or date, inclusive.
// to: Up to this instant, or through the end of this date (UTC), inclusive.
// partnerId: The grant's partner: the referrer, or the partner of the referee's referral.
// windowDays: Attribution window; default each program's setting. 0 = lifetime.
// currency: Default the reporting currency (converted); another code gives that currency only.
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetRoiGrants(ctx context.Context, params *GetRoiGrantsParams, opts ...RequestOption) (*Page[RoiGrant], error) {
	var out Page[RoiGrant]
	if err := c.do(ctx, "GET", "/v1/roi/grants", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetRoiGrantsEach yields every item of GetRoiGrants, following the pages.
func (c *Client) GetRoiGrantsEach(ctx context.Context, params *GetRoiGrantsParams, opts ...RequestOption) iter.Seq2[RoiGrant, error] {
	return paginate(func(cursor *string) (*Page[RoiGrant], error) {
		p := GetRoiGrantsParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetRoiGrants(ctx, &p, opts...)
	})
}

// GetRoiPartnersIDTraceParams are the query parameters of GetRoiPartnersIDTrace.
type GetRoiPartnersIDTraceParams struct {
	WindowDays *int64  `json:"windowDays,omitempty"`
	Format     *string `json:"format,omitempty"`
}

func (p *GetRoiPartnersIDTraceParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "windowDays", p.WindowDays)
	setQuery(v, "format", p.Format)
	return v
}

// GetRoiPartnersIDTrace calls Trace a partner's value. GET /v1/roi/partners/{id}/trace.
//
// The same trace for a partner's own customer identity (through their externalRef).
//
// **Auth:** an API key with `analytics:read` (or `full`), or a session of any role.
//
// windowDays: Attribution window; default each program's setting. 0 = lifetime.
func (c *Client) GetRoiPartnersIDTrace(ctx context.Context, id string, params *GetRoiPartnersIDTraceParams, opts ...RequestOption) (*RoiTrace, error) {
	var out RoiTrace
	if err := c.do(ctx, "GET", "/v1/roi/partners/"+url.PathEscape(id)+"/trace", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetRoiPatterns calls List value patterns. GET /v1/roi/patterns.
//
// Every value pattern from one detection run, with the matching subjects: `credit_only` (a referee used their credit and paid no cash after within the window), `churn_at_exhaustion` (used it up, paid, then stopped), `bonus_collector` (a partner holding commission credit who never spends their own cash) and `negative_partner` (a partner whose referrals net below zero). Matches wait until the evidence is in, so young customers are never flagged.
//
// **Auth:** an API key with `analytics:read` (or `full`), or a session of any role.
func (c *Client) GetRoiPatterns(ctx context.Context, opts ...RequestOption) (*GetRoiPatternsResponse, error) {
	var out GetRoiPatternsResponse
	if err := c.do(ctx, "GET", "/v1/roi/patterns", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetRoiPatternsPatternParams are the query parameters of GetRoiPatternsPattern.
type GetRoiPatternsPatternParams struct {
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetRoiPatternsPatternParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetRoiPatternsPattern calls Get a value pattern. GET /v1/roi/patterns/{pattern}.
//
// The subjects currently matching one pattern, with the detail of each match.
//
// **Auth:** an API key with `analytics:read` (or `full`), or a session of any role.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetRoiPatternsPattern(ctx context.Context, pattern string, params *GetRoiPatternsPatternParams, opts ...RequestOption) (*Page[RoiPatternEntry], error) {
	var out Page[RoiPatternEntry]
	if err := c.do(ctx, "GET", "/v1/roi/patterns/"+url.PathEscape(pattern), params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetRoiPatternsPatternEach yields every item of GetRoiPatternsPattern, following the pages.
func (c *Client) GetRoiPatternsPatternEach(ctx context.Context, pattern string, params *GetRoiPatternsPatternParams, opts ...RequestOption) iter.Seq2[RoiPatternEntry, error] {
	return paginate(func(cursor *string) (*Page[RoiPatternEntry], error) {
		p := GetRoiPatternsPatternParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetRoiPatternsPattern(ctx, pattern, &p, opts...)
	})
}

// GetRoiSummaryParams are the query parameters of GetRoiSummary.
type GetRoiSummaryParams struct {
	From              *string `json:"from,omitempty"`
	To                *string `json:"to,omitempty"`
	ProgramID         *string `json:"programId,omitempty"`
	PartnerID         *string `json:"partnerId,omitempty"`
	WindowDays        *int64  `json:"windowDays,omitempty"`
	Currency          *string `json:"currency,omitempty"`
	Format            *string `json:"format,omitempty"`
	Model             *string `json:"model,omitempty"`
	IncludeProjection *string `json:"includeProjection,omitempty"`
}

func (p *GetRoiSummaryParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "from", p.From)
	setQuery(v, "to", p.To)
	setQuery(v, "programId", p.ProgramID)
	setQuery(v, "partnerId", p.PartnerID)
	setQuery(v, "windowDays", p.WindowDays)
	setQuery(v, "currency", p.Currency)
	setQuery(v, "format", p.Format)
	setQuery(v, "model", p.Model)
	setQuery(v, "includeProjection", p.IncludeProjection)
	return v
}

// GetRoiSummary calls Get the ROI summary. GET /v1/roi/summary.
//
// Credit ROI for the grants in view: credit cost, cash and profit from the recipients after each grant, net, ROI, payback share and median days, outstanding credit and how much of it is measured rather than inferred. With `programId`, also status against the program's target ROI. `model=incremental` adds the lift against the program's holdout group; `includeProjection=true` adds the expected payback of open grants. `from`/`to` select grants by grant date (a date covers its whole UTC day), `programId` and `partnerId` narrow them, `windowDays` overrides each program's attribution window (30, 90, 180, 365, or 0 for lifetime), and `currency` views one native currency instead of the reporting currency. `format=csv` returns the same as a CSV file.
//
// **Auth:** an API key with `analytics:read` (or `full`), or a session of any role.
//
// from: Grants (or referrals) from this instant or date, inclusive.
// to: Up to this instant, or through the end of this date (UTC), inclusive.
// partnerId: The grant's partner: the referrer, or the partner of the referee's referral.
// windowDays: Attribution window; default each program's setting. 0 = lifetime.
// currency: Default the reporting currency (converted); another code gives that currency only.
// model: naive (default): ROI of every grant. incremental: also `incremental`, the program's credit-caused lift against its recorded holdout (requires programId; reporting currency; from/to select referrals by first qualifying payment).
// includeProjection: true: also `projection`, the expected payback of the open grants in view from their program-and-role cohort's median curve (labelled 'projection'; needs 50 closed grants per cohort).
func (c *Client) GetRoiSummary(ctx context.Context, params *GetRoiSummaryParams, opts ...RequestOption) (*RoiSummary, error) {
	var out RoiSummary
	if err := c.do(ctx, "GET", "/v1/roi/summary", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetRoiTraceCustomerRefParams are the query parameters of GetRoiTraceCustomerRef.
type GetRoiTraceCustomerRefParams struct {
	WindowDays *int64  `json:"windowDays,omitempty"`
	Format     *string `json:"format,omitempty"`
}

func (p *GetRoiTraceCustomerRefParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "windowDays", p.WindowDays)
	setQuery(v, "format", p.Format)
	return v
}

// GetRoiTraceCustomerRef calls Trace a customer's value. GET /v1/roi/trace/{customerRef}.
//
// Every value event of a customer (and the identities linked to them) in order, with the running net, to see exactly how their ROI was computed. `format=csv` for a file.
//
// **Auth:** an API key with `analytics:read` (or `full`), or a session of any role.
//
// windowDays: Attribution window; default each program's setting. 0 = lifetime.
func (c *Client) GetRoiTraceCustomerRef(ctx context.Context, customerRef string, params *GetRoiTraceCustomerRefParams, opts ...RequestOption) (*RoiTrace, error) {
	var out RoiTrace
	if err := c.do(ctx, "GET", "/v1/roi/trace/"+url.PathEscape(customerRef), params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetSetupStatus calls Get setup progress. GET /v1/setup/status.
//
// The install checker: which integration steps are done (auth or associate, program and rule, partner, snippet or click, billing, a payable partner, webhooks) and the next steps as hints.
//
// **Auth:** an API key with `analytics:read` (or `full`), or a session of any role.
func (c *Client) GetSetupStatus(ctx context.Context, opts ...RequestOption) (*GetSetupStatusResponse, error) {
	var out GetSetupStatusResponse
	if err := c.do(ctx, "GET", "/v1/setup/status", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetStatus calls Get platform status. GET /v1/status.
//
// Overall state (`operational`, `degraded` or `major_outage`), each component's health and the open incidents. Cached 30 seconds.
//
// **Auth:** none.
func (c *Client) GetStatus(ctx context.Context, opts ...RequestOption) (*GetStatusResponse, error) {
	var out GetStatusResponse
	if err := c.do(ctx, "GET", "/v1/status", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetStatusIncidentsParams are the query parameters of GetStatusIncidents.
type GetStatusIncidentsParams struct {
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetStatusIncidentsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetStatusIncidents calls List status incidents. GET /v1/status/incidents.
//
// Recent incidents, newest first, with their updates.
//
// **Auth:** none.
//
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetStatusIncidents(ctx context.Context, params *GetStatusIncidentsParams, opts ...RequestOption) (*Page[Incident], error) {
	var out Page[Incident]
	if err := c.do(ctx, "GET", "/v1/status/incidents", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetStatusIncidentsEach yields every item of GetStatusIncidents, following the pages.
func (c *Client) GetStatusIncidentsEach(ctx context.Context, params *GetStatusIncidentsParams, opts ...RequestOption) iter.Seq2[Incident, error] {
	return paginate(func(cursor *string) (*Page[Incident], error) {
		p := GetStatusIncidentsParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetStatusIncidents(ctx, &p, opts...)
	})
}

// GetStatusIncidentsAtom calls Get the incidents feed. GET /v1/status/incidents.atom.
//
// The 50 most recent incidents as an Atom feed.
//
// **Auth:** none.
func (c *Client) GetStatusIncidentsAtom(ctx context.Context, opts ...RequestOption) (string, error) {
	var out string
	err := c.do(ctx, "GET", "/v1/status/incidents.atom", nil, nil, &out, "text", opts)
	return out, err
}

// GetStatusIncidentsID calls Get a status incident. GET /v1/status/incidents/{id}.
//
// One incident with its updates.
//
// **Auth:** none.
func (c *Client) GetStatusIncidentsID(ctx context.Context, id string, opts ...RequestOption) (*Incident, error) {
	var out Incident
	if err := c.do(ctx, "GET", "/v1/status/incidents/"+url.PathEscape(id), nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetSupportApps calls List support tool connections. GET /v1/support/apps.
//
// Every support tool Referalli can show referral context in (Chatly cards and AI tools, Intercom, Zendesk, HubSpot, Front, Help Scout, custom), with its card URL, install steps and whether it is set up. Secrets are never returned.
//
// **Auth:** an API key with `partners:read` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) GetSupportApps(ctx context.Context, opts ...RequestOption) (*GetSupportAppsResponse, error) {
	var out GetSupportAppsResponse
	if err := c.do(ctx, "GET", "/v1/support/apps", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetSupportCardsOrgIDAdapterParams are the query parameters of GetSupportCardsOrgIDAdapter.
type GetSupportCardsOrgIDAdapterParams struct {
	Email      *string `json:"email,omitempty"`
	ExternalID *string `json:"externalId,omitempty"`
}

func (p *GetSupportCardsOrgIDAdapterParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "email", p.Email)
	setQuery(v, "externalId", p.ExternalID)
	return v
}

// GetSupportCardsOrgIDAdapter calls Render a support card for a tool (GET). GET /v1/support/cards/{orgId}/{adapter}.
//
// The same card for tools that fetch with GET (HubSpot legacy CRM cards pass the contact as query parameters). Signed as for the POST form; 404 otherwise.
//
// **Auth:** none.
func (c *Client) GetSupportCardsOrgIDAdapter(ctx context.Context, orgId string, adapter string, params *GetSupportCardsOrgIDAdapterParams, opts ...RequestOption) (*PostSupportCardsOrgIdAdapterResponse, error) {
	var out PostSupportCardsOrgIdAdapterResponse
	if err := c.do(ctx, "GET", "/v1/support/cards/"+url.PathEscape(orgId)+"/"+url.PathEscape(adapter), params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetSupportContextParams are the query parameters of GetSupportContext.
type GetSupportContextParams struct {
	Email              *string `json:"email,omitempty"`
	CustomerRef        *string `json:"customerRef,omitempty"`
	PartnerExternalRef *string `json:"partnerExternalRef,omitempty"`
	Ext                *string `json:"ext,omitempty"`
}

func (p *GetSupportContextParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "email", p.Email)
	setQuery(v, "customerRef", p.CustomerRef)
	setQuery(v, "partnerExternalRef", p.PartnerExternalRef)
	setQuery(v, "ext", p.Ext)
	return v
}

// GetSupportContext calls Get referral context for a support contact. GET /v1/support/context.
//
// What a support agent should know about a contact: the partner (status, tier, programs, referral link, pending/approved/paid balances, next payout, open flags) and/or the referred customer (who referred them, when, the referee reward and what is left of it). Name the contact by `partnerExternalRef`, `ext` (a tool's verified external id), `customerRef` or `email`; the email is used only when no stronger key is given. Your org's data only; `partner` and `customer` are null when nothing matches.
//
// **Auth:** an API key with `partners:read` (or `full`), or a session with role `owner`, `admin`, `support`, `finance`, `analyst`.
//
// email: Used only when no stronger key is given.
// ext: A support tool's verified external id: matched as a partner externalRef, then as a customer ref.
func (c *Client) GetSupportContext(ctx context.Context, params *GetSupportContextParams, opts ...RequestOption) (*GetSupportContextResponse, error) {
	var out GetSupportContextResponse
	if err := c.do(ctx, "GET", "/v1/support/context", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetTax1099Params are the query parameters of GetTax1099.
type GetTax1099Params struct {
	Year          *int64  `json:"year,omitempty"`
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetTax1099Params) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "year", p.Year)
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetTax1099 calls Export 1099-NEC data. GET /v1/tax/1099.
//
// For each partner with a W-9 on file: the total actually paid in the tax year, in USD at recorded FX rates, and whether a form is required (US$600 or more, or a currency that couldn't be converted, listed in `fxMissingCurrencies` for review).
//
// **Auth:** an API key with `payouts:read` (or `full`), or a session with role `owner`, `admin`, `finance`.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetTax1099(ctx context.Context, params *GetTax1099Params, opts ...RequestOption) (*Page[GetTax1099ResponseDataItem], error) {
	var out Page[GetTax1099ResponseDataItem]
	if err := c.do(ctx, "GET", "/v1/tax/1099", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetTax1099Each yields every item of GetTax1099, following the pages.
func (c *Client) GetTax1099Each(ctx context.Context, params *GetTax1099Params, opts ...RequestOption) iter.Seq2[GetTax1099ResponseDataItem, error] {
	return paginate(func(cursor *string) (*Page[GetTax1099ResponseDataItem], error) {
		p := GetTax1099Params{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetTax1099(ctx, &p, opts...)
	})
}

// GetTax1099FilingsParams are the query parameters of GetTax1099Filings.
type GetTax1099FilingsParams struct {
	Year          *int64  `json:"year,omitempty"`
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetTax1099FilingsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "year", p.Year)
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetTax1099Filings calls List 1099 filings. GET /v1/tax/1099/filings.
//
// A tax year's filings with their provider status and the documents in each.
//
// **Auth:** an API key with `payouts:read` (or `full`), or a session with role `owner`, `admin`, `finance`.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetTax1099Filings(ctx context.Context, params *GetTax1099FilingsParams, opts ...RequestOption) (*Page[GetTax1099FilingsResponseDataItem], error) {
	var out Page[GetTax1099FilingsResponseDataItem]
	if err := c.do(ctx, "GET", "/v1/tax/1099/filings", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetTax1099FilingsEach yields every item of GetTax1099Filings, following the pages.
func (c *Client) GetTax1099FilingsEach(ctx context.Context, params *GetTax1099FilingsParams, opts ...RequestOption) iter.Seq2[GetTax1099FilingsResponseDataItem, error] {
	return paginate(func(cursor *string) (*Page[GetTax1099FilingsResponseDataItem], error) {
		p := GetTax1099FilingsParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetTax1099Filings(ctx, &p, opts...)
	})
}

// GetTax1099PreviewParams are the query parameters of GetTax1099Preview.
type GetTax1099PreviewParams struct {
	Year *int64 `json:"year,omitempty"`
}

func (p *GetTax1099PreviewParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "year", p.Year)
	return v
}

// GetTax1099Preview calls Preview 1099-NEC filings. GET /v1/tax/1099/preview.
//
// Which partners would get a 1099-NEC for a tax year, with amounts and anything missing (W-9s, TINs), without filing.
//
// **Auth:** an API key with `payouts:read` (or `full`), or a session with role `owner`, `admin`, `finance`.
func (c *Client) GetTax1099Preview(ctx context.Context, params *GetTax1099PreviewParams, opts ...RequestOption) (*NecPreview, error) {
	var out NecPreview
	if err := c.do(ctx, "GET", "/v1/tax/1099/preview", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetTestClock calls Get the test clock. GET /v1/test_clock.
//
// Test mode only (`409 test_mode_only` otherwise). The test org's current time and how far it has been advanced.
//
// **Auth:** an API key with `org:read` (or `full`), or a session of any role.
func (c *Client) GetTestClock(ctx context.Context, opts ...RequestOption) (*GetTestClockResponse, error) {
	var out GetTestClockResponse
	if err := c.do(ctx, "GET", "/v1/test_clock", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetTestEmailsParams are the query parameters of GetTestEmails.
type GetTestEmailsParams struct {
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetTestEmailsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetTestEmails calls List captured test emails. GET /v1/test/emails.
//
// Test mode only. Emails the test org would have sent, captured instead of delivered, newest first.
//
// **Auth:** an API key with `org:read` (or `full`), or a session of any role.
//
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetTestEmails(ctx context.Context, params *GetTestEmailsParams, opts ...RequestOption) (*Page[GetTestEmailsResponseDataItem], error) {
	var out Page[GetTestEmailsResponseDataItem]
	if err := c.do(ctx, "GET", "/v1/test/emails", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetTestEmailsEach yields every item of GetTestEmails, following the pages.
func (c *Client) GetTestEmailsEach(ctx context.Context, params *GetTestEmailsParams, opts ...RequestOption) iter.Seq2[GetTestEmailsResponseDataItem, error] {
	return paginate(func(cursor *string) (*Page[GetTestEmailsResponseDataItem], error) {
		p := GetTestEmailsParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetTestEmails(ctx, &p, opts...)
	})
}

// GetValueImportJobID calls Get a value import job. GET /v1/value/import/{jobId}.
//
// Progress and result of a background import: status, processed and total rows, recorded and duplicate counts and row errors.
//
// **Auth:** an API key with `analytics:read` (or `full`), or a session of any role.
func (c *Client) GetValueImportJobID(ctx context.Context, jobId string, opts ...RequestOption) (*GetValueImportJobIdResponse, error) {
	var out GetValueImportJobIdResponse
	if err := c.do(ctx, "GET", "/v1/value/import/"+url.PathEscape(jobId), nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetVersions calls List API versions. GET /v1/versions.
//
// Every API version, newest first, with what changed in it, the latest version and the deprecation policy. Pin a version with the `Referalli-Version` header or on the API key.
//
// **Auth:** none.
func (c *Client) GetVersions(ctx context.Context, opts ...RequestOption) (*GetVersionsResponse, error) {
	var out GetVersionsResponse
	if err := c.do(ctx, "GET", "/v1/versions", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetWebhookDeliveriesParams are the query parameters of GetWebhookDeliveries.
type GetWebhookDeliveriesParams struct {
	EndpointID    *string `json:"endpointId,omitempty"`
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetWebhookDeliveriesParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "endpointId", p.EndpointID)
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetWebhookDeliveries calls List webhook deliveries. GET /v1/webhook-deliveries.
//
// Deliveries to every endpoint, newest first: topic, status (`pending`, `delivered`, `failed`), attempts, the last response code and, while pending, when the next attempt is. Filter with `endpointId`.
//
// **Auth:** an API key with `webhooks:read` (or `full`), or a session of any role.
//
// endpointId: Only deliveries to this endpoint.
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetWebhookDeliveries(ctx context.Context, params *GetWebhookDeliveriesParams, opts ...RequestOption) (*Page[WebhookDelivery], error) {
	var out Page[WebhookDelivery]
	if err := c.do(ctx, "GET", "/v1/webhook-deliveries", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetWebhookDeliveriesEach yields every item of GetWebhookDeliveries, following the pages.
func (c *Client) GetWebhookDeliveriesEach(ctx context.Context, params *GetWebhookDeliveriesParams, opts ...RequestOption) iter.Seq2[WebhookDelivery, error] {
	return paginate(func(cursor *string) (*Page[WebhookDelivery], error) {
		p := GetWebhookDeliveriesParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetWebhookDeliveries(ctx, &p, opts...)
	})
}

// GetWebhookDeliveriesID calls Get a webhook delivery. GET /v1/webhook-deliveries/{id}.
//
// One delivery in full: the exact `payload` sent, its `eventId`, and what the endpoint answered last (`lastStatusCode`, the first 2,000 characters of `lastResponseBody`, or `lastError` when there was no response).
//
// **Auth:** an API key with `webhooks:read` (or `full`), or a session of any role.
func (c *Client) GetWebhookDeliveriesID(ctx context.Context, id string, opts ...RequestOption) (*GetWebhookDeliveriesIdResponse, error) {
	var out GetWebhookDeliveriesIdResponse
	if err := c.do(ctx, "GET", "/v1/webhook-deliveries/"+url.PathEscape(id), nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetWebhookEndpointsParams are the query parameters of GetWebhookEndpoints.
type GetWebhookEndpointsParams struct {
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetWebhookEndpointsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetWebhookEndpoints calls List webhook endpoints. GET /v1/webhook-endpoints.
//
// Endpoints with their topics, programs, status (`active`, `paused`, `disabled`), secret-rotation state and failure state. Secrets are never returned.
//
// **Auth:** an API key with `webhooks:read` (or `full`), or a session of any role.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetWebhookEndpoints(ctx context.Context, params *GetWebhookEndpointsParams, opts ...RequestOption) (*Page[WebhookEndpoint], error) {
	var out Page[WebhookEndpoint]
	if err := c.do(ctx, "GET", "/v1/webhook-endpoints", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetWebhookEndpointsEach yields every item of GetWebhookEndpoints, following the pages.
func (c *Client) GetWebhookEndpointsEach(ctx context.Context, params *GetWebhookEndpointsParams, opts ...RequestOption) iter.Seq2[WebhookEndpoint, error] {
	return paginate(func(cursor *string) (*Page[WebhookEndpoint], error) {
		p := GetWebhookEndpointsParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetWebhookEndpoints(ctx, &p, opts...)
	})
}

// GetWebhookEndpointsID calls Get a webhook endpoint. GET /v1/webhook-endpoints/{id}.
//
// One endpoint with its status, rotation and failure state.
//
// **Auth:** an API key with `webhooks:read` (or `full`), or a session of any role.
func (c *Client) GetWebhookEndpointsID(ctx context.Context, id string, opts ...RequestOption) (*WebhookEndpoint, error) {
	var out WebhookEndpoint
	if err := c.do(ctx, "GET", "/v1/webhook-endpoints/"+url.PathEscape(id), nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetWebhookEndpointsIDDeliveriesParams are the query parameters of GetWebhookEndpointsIDDeliveries.
type GetWebhookEndpointsIDDeliveriesParams struct {
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetWebhookEndpointsIDDeliveriesParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetWebhookEndpointsIDDeliveries calls List an endpoint's deliveries. GET /v1/webhook-endpoints/{id}/deliveries.
//
// What one endpoint was sent, newest first.
//
// **Auth:** an API key with `webhooks:read` (or `full`), or a session of any role.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetWebhookEndpointsIDDeliveries(ctx context.Context, id string, params *GetWebhookEndpointsIDDeliveriesParams, opts ...RequestOption) (*Page[WebhookDelivery], error) {
	var out Page[WebhookDelivery]
	if err := c.do(ctx, "GET", "/v1/webhook-endpoints/"+url.PathEscape(id)+"/deliveries", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetWebhookEndpointsIDDeliveriesEach yields every item of GetWebhookEndpointsIDDeliveries, following the pages.
func (c *Client) GetWebhookEndpointsIDDeliveriesEach(ctx context.Context, id string, params *GetWebhookEndpointsIDDeliveriesParams, opts ...RequestOption) iter.Seq2[WebhookDelivery, error] {
	return paginate(func(cursor *string) (*Page[WebhookDelivery], error) {
		p := GetWebhookEndpointsIDDeliveriesParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetWebhookEndpointsIDDeliveries(ctx, id, &p, opts...)
	})
}

// GetZapierEvents calls List Zapier trigger events. GET /v1/zapier/events.
//
// The event keys Zapier triggers can subscribe to.
//
// **Auth:** an API key with `integrations:read` (or `full`), or a session of any role.
func (c *Client) GetZapierEvents(ctx context.Context, opts ...RequestOption) (*GetZapierEventsResponse, error) {
	var out GetZapierEventsResponse
	if err := c.do(ctx, "GET", "/v1/zapier/events", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetZapierHooksParams are the query parameters of GetZapierHooks.
type GetZapierHooksParams struct {
	Limit         *int64  `json:"limit,omitempty"`
	StartingAfter *string `json:"starting_after,omitempty"`
	EndingBefore  *string `json:"ending_before,omitempty"`
	CreatedGte    *string `json:"created[gte],omitempty"`
	CreatedLte    *string `json:"created[lte],omitempty"`
}

func (p *GetZapierHooksParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "limit", p.Limit)
	setQuery(v, "starting_after", p.StartingAfter)
	setQuery(v, "ending_before", p.EndingBefore)
	setQuery(v, "created[gte]", p.CreatedGte)
	setQuery(v, "created[lte]", p.CreatedLte)
	return v
}

// GetZapierHooks calls List Zapier hooks. GET /v1/zapier/hooks.
//
// Active Zapier REST hook subscriptions, newest first.
//
// **Auth:** an API key with `integrations:read` (or `full`), or a session with role `owner`, `admin`.
//
// limit: Items per page, 1 to 100 (default 50). From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// starting_after: A cursor: the id of an item (the previous page's `next_cursor`); returns the items after it. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// ending_before: A cursor: returns the items before this item id (the previous page). Not with `starting_after`. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[gte]: Only items created at or after this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
// created[lte]: Only items created at or before this time (ISO 8601), on lists of timestamped records. From API version 2026-11-01; ignored under 2026-10-21, where lists are a capped plain array.
func (c *Client) GetZapierHooks(ctx context.Context, params *GetZapierHooksParams, opts ...RequestOption) (*Page[ZapierHook], error) {
	var out Page[ZapierHook]
	if err := c.do(ctx, "GET", "/v1/zapier/hooks", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetZapierHooksEach yields every item of GetZapierHooks, following the pages.
func (c *Client) GetZapierHooksEach(ctx context.Context, params *GetZapierHooksParams, opts ...RequestOption) iter.Seq2[ZapierHook, error] {
	return paginate(func(cursor *string) (*Page[ZapierHook], error) {
		p := GetZapierHooksParams{}
		if params != nil {
			p = *params
		}
		p.StartingAfter = cursor
		return c.GetZapierHooks(ctx, &p, opts...)
	})
}

// GetZapierTriggersEventParams are the query parameters of GetZapierTriggersEvent.
type GetZapierTriggersEventParams struct {
	Limit *int64 `json:"limit,omitempty"`
}

func (p *GetZapierTriggersEventParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "limit", p.Limit)
	return v
}

// GetZapierTriggersEvent calls Poll a Zapier trigger. GET /v1/zapier/triggers/{event}.
//
// Zapier polling and sample data: the latest `limit` events of that kind, newest first, in the shape the hook delivers. Always a plain array (Zapier's contract), whatever the API version.
//
// **Auth:** an API key with `integrations:read` (or `full`), or a session of any role.
func (c *Client) GetZapierTriggersEvent(ctx context.Context, event string, params *GetZapierTriggersEventParams, opts ...RequestOption) (*GetZapierTriggersEventResponse, error) {
	var out GetZapierTriggersEventResponse
	if err := c.do(ctx, "GET", "/v1/zapier/triggers/"+url.PathEscape(event), params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PatchAdminStatusIncidentsID calls Edit a status incident. PATCH /v1/admin/status/incidents/{id}.
//
// Platform admins only. Changes an incident's title, severity, state, body or affected components. Setting `state: "resolved"` stamps `resolvedAt`.
//
// **Auth:** a platform admin session.
func (c *Client) PatchAdminStatusIncidentsID(ctx context.Context, id string, body PatchAdminStatusIncidentsIdBody, opts ...RequestOption) (*Incident, error) {
	var out Incident
	if err := c.do(ctx, "PATCH", "/v1/admin/status/incidents/"+url.PathEscape(id), nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PatchAPIKeysID calls Update an API key. PATCH /v1/api-keys/{id}.
//
// Changes a key's pinned `apiVersion`, `owner` or expiry (`expiresInDays` from now; `null` removes it). Revoked keys are 404. Needs an owner or admin session, or an API key with the `admin` scope. Audited.
//
// **Auth:** an API key with `team:write` (or `full`), or a session of any role.
func (c *Client) PatchAPIKeysID(ctx context.Context, id string, body PatchApiKeysIdBody, opts ...RequestOption) (*PatchApiKeysIdResponse, error) {
	var out PatchApiKeysIdResponse
	if err := c.do(ctx, "PATCH", "/v1/api-keys/"+url.PathEscape(id), nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PatchAuthConnectionsID calls Update an auth connection. PATCH /v1/auth/connections/{id}.
//
// Pauses or resumes a connection, renames it, changes `autoEnrollUsers` or the destination of auto-created links, or (signed-protocol connections) sets `signatureVersion`: `v2` to require request-line signatures, `v1+v2` to accept both again.
//
// **Auth:** an API key with `integrations:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PatchAuthConnectionsID(ctx context.Context, id string, body PatchAuthConnectionsIdBody, opts ...RequestOption) (*PatchAuthConnectionsIdResponse, error) {
	var out PatchAuthConnectionsIdResponse
	if err := c.do(ctx, "PATCH", "/v1/auth/connections/"+url.PathEscape(id), nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PatchBillingConnectionsID calls Update a billing connection. PATCH /v1/billing/connections/{id}.
//
// Renames a connection, sets `status` (`paused` makes its webhook URL refuse events until set back to `active`), or sets its product `filter` (`null` removes it). A filter applies to events received from then on. Audited.
//
// **Auth:** an API key with `billing:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PatchBillingConnectionsID(ctx context.Context, id string, body PatchBillingConnectionsIdBody, opts ...RequestOption) (*BillingConnection, error) {
	var out BillingConnection
	if err := c.do(ctx, "PATCH", "/v1/billing/connections/"+url.PathEscape(id), nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PatchOrg calls Update the organization. PATCH /v1/org.
//
// Changes the org name, merges keys into `settings` (a free-form map) and turns `payoutKycRequired` on or off. Audited.
//
// **Auth:** an API key with `org:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PatchOrg(ctx context.Context, body PatchOrgBody, opts ...RequestOption) (*PatchOrgResponse, error) {
	var out PatchOrgResponse
	if err := c.do(ctx, "PATCH", "/v1/org", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PatchOrgMembersUserID calls Change a member's role. PATCH /v1/org/members/{userId}.
//
// Changes a member's role. Only owners change owner roles, and the last owner can't be demoted (409). Audited. Live mode only.
//
// **Auth:** an API key with `team:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PatchOrgMembersUserID(ctx context.Context, userId string, body PatchOrgMembersUserIdBody, opts ...RequestOption) (*PatchOrgMembersUserIdResponse, error) {
	var out PatchOrgMembersUserIdResponse
	if err := c.do(ctx, "PATCH", "/v1/org/members/"+url.PathEscape(userId), nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PatchOrgSSO calls Update SSO settings. PATCH /v1/org/sso.
//
// Requires SSO for members on the org's SSO domains (`requireSso`; needs an active connection and the owner signed in through SSO, so nobody is locked out) and maps identity-provider roles to Referalli roles. Owner only. Live mode only.
//
// **Auth:** an API key with `org:write` (or `full`), or a session with role `owner`.
func (c *Client) PatchOrgSSO(ctx context.Context, body PatchOrgSsoBody, opts ...RequestOption) (*PatchOrgSsoResponse, error) {
	var out PatchOrgSsoResponse
	if err := c.do(ctx, "PATCH", "/v1/org/sso", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PatchPartnerGroupsID calls Update a partner group. PATCH /v1/partner-groups/{id}.
//
// Changes a group's terms, or archives it (`archived: true`); archived groups take no new members.
//
// **Auth:** an API key with `programs:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PatchPartnerGroupsID(ctx context.Context, id string, body PatchPartnerGroupsIdBody, opts ...RequestOption) (*PartnerGroup, error) {
	var out PartnerGroup
	if err := c.do(ctx, "PATCH", "/v1/partner-groups/"+url.PathEscape(id), nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PatchPartnersID calls Update a partner. PATCH /v1/partners/{id}.
//
// Changes the display name, status (`active`, `suspended`, `banned`, `archived`), tags, notes or externalRef (`null` clears it). Suspended partners keep earning but are not paid; banned ones stop attributing. 409 when another partner has the externalRef. Audited.
//
// **Auth:** an API key with `partners:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PatchPartnersID(ctx context.Context, id string, body PatchPartnersIdBody, opts ...RequestOption) (*Partner, error) {
	var out Partner
	if err := c.do(ctx, "PATCH", "/v1/partners/"+url.PathEscape(id), nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PatchPartnersIDPayoutProfile calls Change a partner's payout threshold. PATCH /v1/partners/{id}/payout-profile.
//
// Changes only the minimum balance before a payout (`minThresholdMinor`). 404 when the partner has no profile. Audited.
//
// **Auth:** an API key with `payouts:write` (or `full`), or a session with role `owner`, `admin`, `finance`.
func (c *Client) PatchPartnersIDPayoutProfile(ctx context.Context, id string, body PatchPartnersIdPayoutProfileBody, opts ...RequestOption) (*PatchPartnersIdPayoutProfileResponse, error) {
	var out PatchPartnersIdPayoutProfileResponse
	if err := c.do(ctx, "PATCH", "/v1/partners/"+url.PathEscape(id)+"/payout-profile", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PatchProgramsID calls Update a program. PATCH /v1/programs/{id}.
//
// Changes any program field or setting; omitted fields keep their value, and `publicPage` is merged. Changing `termsMd` records a new terms version. Setting `status` to `paused` or `archived` stops new attribution. `mobileApp` (iOS Team ID and bundle id, Android package and signing-certificate fingerprints) publishes the program's apps in the universal-links and App Links files on the org's link hosts. Audited.
//
// **Auth:** an API key with `programs:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PatchProgramsID(ctx context.Context, id string, body PatchProgramsIdBody, opts ...RequestOption) (*Program, error) {
	var out Program
	if err := c.do(ctx, "PATCH", "/v1/programs/"+url.PathEscape(id), nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PatchReferralsID calls Pause or resume a referral. PATCH /v1/referrals/{id}.
//
// Sets `status` to `paused` (no new commissions) or back to `active`. Only active referrals can be paused and only paused ones resumed (409 otherwise); setting the current status is a no-op. Publishes `referral.paused` or `referral.resumed`. Audited.
//
// **Auth:** an API key with `referrals:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PatchReferralsID(ctx context.Context, id string, body PatchReferralsIdBody, opts ...RequestOption) (*Referral, error) {
	var out Referral
	if err := c.do(ctx, "PATCH", "/v1/referrals/"+url.PathEscape(id), nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PatchWebhookEndpointsID calls Update a webhook endpoint. PATCH /v1/webhook-endpoints/{id}.
//
// Changes the URL, topics, programs (`null` or `[]` for every program) or legacy signature, or pauses it (`status: "paused"`, deliveries held). Setting `active` re-enables a paused or auto-disabled endpoint: held deliveries resume and the failure clock starts over. `warnings` has `domain_changed` when the URL moves to another host. Audited.
//
// **Auth:** an API key with `webhooks:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PatchWebhookEndpointsID(ctx context.Context, id string, body PatchWebhookEndpointsIdBody, opts ...RequestOption) (*PatchWebhookEndpointsIdResponse, error) {
	var out PatchWebhookEndpointsIdResponse
	if err := c.do(ctx, "PATCH", "/v1/webhook-endpoints/"+url.PathEscape(id), nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostAccountClose calls Close your own account. POST /v1/account/close.
//
// Closes the signed-in merchant user or partner account: every session ends, identity is redacted and commission records stay under a pseudonym. Send `confirm` with the account's email. Refused (409) for an organization's only owner or a partner still owed money. Step-up: `x-totp-code` when two-factor is on, otherwise a sign-in less than 15 minutes old (else 401 `step_up_required`).
//
// **Auth:** a partner session.
func (c *Client) PostAccountClose(ctx context.Context, body PostAccountCloseBody, opts ...RequestOption) (*PostAccountCloseResponse, error) {
	var out PostAccountCloseResponse
	if err := c.do(ctx, "POST", "/v1/account/close", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostAdminOrgsIDPlan calls Set an organization's plan. POST /v1/admin/orgs/{id}/plan.
//
// Platform admins only. Moves an org to a plan without checkout.
//
// **Auth:** a platform admin session.
func (c *Client) PostAdminOrgsIDPlan(ctx context.Context, id string, body PostAdminOrgsIdPlanBody, opts ...RequestOption) (*PostAdminOrgsIdPlanResponse, error) {
	var out PostAdminOrgsIdPlanResponse
	if err := c.do(ctx, "POST", "/v1/admin/orgs/"+url.PathEscape(id)+"/plan", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostAdminOrgsIDSuspend calls Suspend an organization. POST /v1/admin/orgs/{id}/suspend.
//
// Platform admins only. Suspends an org with a reason; its members and keys get 403 until it's unsuspended.
//
// **Auth:** a platform admin session.
func (c *Client) PostAdminOrgsIDSuspend(ctx context.Context, id string, body PostAdminOrgsIdSuspendBody, opts ...RequestOption) (*PostAdminOrgsIdSuspendResponse, error) {
	var out PostAdminOrgsIdSuspendResponse
	if err := c.do(ctx, "POST", "/v1/admin/orgs/"+url.PathEscape(id)+"/suspend", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostAdminOrgsIDUnsuspend calls Unsuspend an organization. POST /v1/admin/orgs/{id}/unsuspend.
//
// Platform admins only. Lifts a suspension.
//
// **Auth:** a platform admin session.
func (c *Client) PostAdminOrgsIDUnsuspend(ctx context.Context, id string, opts ...RequestOption) (*PostAdminOrgsIdUnsuspendResponse, error) {
	var out PostAdminOrgsIdUnsuspendResponse
	if err := c.do(ctx, "POST", "/v1/admin/orgs/"+url.PathEscape(id)+"/unsuspend", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostAdminPartnerAccountsIDBan calls Ban a partner account. POST /v1/admin/partner-accounts/{id}/ban.
//
// Platform admins only. Bans a partner account across the platform and signs it out everywhere.
//
// **Auth:** a platform admin session.
func (c *Client) PostAdminPartnerAccountsIDBan(ctx context.Context, id string, opts ...RequestOption) (*PostAdminPartnerAccountsIdBanResponse, error) {
	var out PostAdminPartnerAccountsIdBanResponse
	if err := c.do(ctx, "POST", "/v1/admin/partner-accounts/"+url.PathEscape(id)+"/ban", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostAdminPartnerAccountsIDUnban calls Unban a partner account. POST /v1/admin/partner-accounts/{id}/unban.
//
// Platform admins only. Lifts a platform ban.
//
// **Auth:** a platform admin session.
func (c *Client) PostAdminPartnerAccountsIDUnban(ctx context.Context, id string, opts ...RequestOption) (*PostAdminPartnerAccountsIdUnbanResponse, error) {
	var out PostAdminPartnerAccountsIdUnbanResponse
	if err := c.do(ctx, "POST", "/v1/admin/partner-accounts/"+url.PathEscape(id)+"/unban", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostAdminStatusIncidents calls Open a status incident. POST /v1/admin/status/incidents.
//
// Platform admins only. Opens an incident on the status page (and mirrors it to the hosted status space when configured).
//
// **Auth:** a platform admin session.
func (c *Client) PostAdminStatusIncidents(ctx context.Context, body PostAdminStatusIncidentsBody, opts ...RequestOption) (*Incident, error) {
	var out Incident
	if err := c.do(ctx, "POST", "/v1/admin/status/incidents", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostAdminStatusIncidentsIDUpdates calls Post an incident update. POST /v1/admin/status/incidents/{id}/updates.
//
// Platform admins only. Adds a timestamped update to an incident and moves it to the given state.
//
// **Auth:** a platform admin session.
func (c *Client) PostAdminStatusIncidentsIDUpdates(ctx context.Context, id string, body PostAdminStatusIncidentsIdUpdatesBody, opts ...RequestOption) (*Incident, error) {
	var out Incident
	if err := c.do(ctx, "POST", "/v1/admin/status/incidents/"+url.PathEscape(id)+"/updates", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostAgentsRegister calls Register an AI agent. POST /v1/agents/register.
//
// Registers (or updates) an AI agent identity, proven by a wallet signature over a nonce from `POST /v1/partner/auth/wallet/nonce`. The nonce is single use and valid for 10 minutes. `instantPayout` pays the agent per outcome on the crypto rail. No API key: the signature is the credential.
//
// **Auth:** none.
func (c *Client) PostAgentsRegister(ctx context.Context, body PostAgentsRegisterBody, opts ...RequestOption) (*PostAgentsRegisterResponse, error) {
	var out PostAgentsRegisterResponse
	if err := c.do(ctx, "POST", "/v1/agents/register", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostAnomaliesScan calls Scan for anomalies. POST /v1/anomalies/scan.
//
// Compares the last 24 hours of clicks, sign-ups and revenue with the trailing 7-day average and returns this org's spikes (3x or more) and drops (under 20%). Findings also become notifications; nothing is changed automatically. Also runs hourly.
//
// **Auth:** an API key with `analytics:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostAnomaliesScan(ctx context.Context, opts ...RequestOption) (*PostAnomaliesScanResponse, error) {
	var out PostAnomaliesScanResponse
	if err := c.do(ctx, "POST", "/v1/anomalies/scan", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostAPIKeys calls Create an API key. POST /v1/api-keys.
//
// Mints a secret key (`rk_live_…` / `rk_test_…`), shown only in this response. `scopes` defaults to `full`; a key can't grant scopes it doesn't hold. `mode` picks live, test or a matching pair of both (`both` returns `{ pairId, live, test }`); a test-mode caller can mint test keys only. `publishable: true` mints a browser-safe `rk_pub_` key limited to `programs:read` and `events:write` from `allowedOrigins`. `apiVersion` pins the key (default: latest). Needs an owner or admin session, or an API key with the `admin` scope. Audited.
//
// **Auth:** an API key with `team:write` (or `full`), or a session of any role.
func (c *Client) PostAPIKeys(ctx context.Context, body PostApiKeysBody, opts ...RequestOption) (*PostApiKeysResponse, error) {
	var out PostApiKeysResponse
	if err := c.do(ctx, "POST", "/v1/api-keys", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostAssets calls Create a marketing asset. POST /v1/assets.
//
// Adds a marketing asset (`banner`, `copy`, `video` or `logo`) to a program for its approved partners, or only for one group's members with `groupId`.
//
// **Auth:** an API key with `partners:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostAssets(ctx context.Context, body PostAssetsBody, opts ...RequestOption) (*PostAssetsResponse, error) {
	var out PostAssetsResponse
	if err := c.do(ctx, "POST", "/v1/assets", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostAttributionReplay calls Replay attribution. POST /v1/attribution/replay.
//
// Deterministically recomputes attribution for a program's referrals from the recorded clicks and identity links, under its current model. Returns what changed. Audited.
//
// **Auth:** an API key with `referrals:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostAttributionReplay(ctx context.Context, body PostAttributionReplayBody, opts ...RequestOption) (*PostAttributionReplayResponse, error) {
	var out PostAttributionReplayResponse
	if err := c.do(ctx, "POST", "/v1/attribution/replay", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostAuthBackchannelLogout calls Receive an OpenID Connect back-channel logout. POST /v1/auth/backchannel-logout.
//
// The `backchannel_logout_uri` of Referalli's own sign-in client. Takes a form-encoded `logout_token` (OpenID Connect Back-Channel Logout 1.0): verified against the issuer's keys, `aud`, `events`, a fresh `iat`, no `nonce`, and a single-use `jti`; then ends the Referalli sessions for its `sid` (or every session of its `sub`). Answers 503 until the deployment sets IDP_LOGOUT_SYNC=true.
//
// **Auth:** none.
func (c *Client) PostAuthBackchannelLogout(ctx context.Context, body PostAuthBackchannelLogoutBody, opts ...RequestOption) (*PostAuthBackchannelLogoutResponse, error) {
	var out PostAuthBackchannelLogoutResponse
	if err := c.do(ctx, "POST", "/v1/auth/backchannel-logout", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostAuthConnections calls Connect an auth platform. POST /v1/auth/connections.
//
// Connects your sign-up system so new users are attributed server-side without calling associate: `clerk`, `supabase`, `workos` or `generic` receive the platform's user webhooks at the returned URL; `atlas` and `custom` call the signed protocol endpoints under `/v1/auth/connections/{id}/`. `autoEnrollUsers` makes every new user a partner too. Returns the signing secret once. 409 when an Atlas instance already has a connection.
//
// **Auth:** an API key with `integrations:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostAuthConnections(ctx context.Context, body PostAuthConnectionsBody, opts ...RequestOption) (*CreatedAuthConnection, error) {
	var out CreatedAuthConnection
	if err := c.do(ctx, "POST", "/v1/auth/connections", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostAuthConnectionsIDAliases calls Change a user's id (signed protocol). POST /v1/auth/connections/{id}/aliases.
//
// The app stamped or changed a user's id after the sign-up was filed: referrals and identity move from `fromCustomerRef` to `toCustomerRef`. Publishes `customer.aliased`. Signed with the connection secret in `x-referalli-signature: t=<unix>,v1=…,v2=…` (v2 also signs the request line; once a connection has seen v2 it requires it), with a unique `x-referalli-id` per call. The legacy header names from before the rename are also accepted, for existing integrations. Calls older than 5 minutes, reused ids (`409 replay`), paused connections (`423 connection_paused`) and unknown fields (`422`) are refused.
//
// **Auth:** a signature (see above); no API key.
func (c *Client) PostAuthConnectionsIDAliases(ctx context.Context, id string, body PostAuthConnectionsIdAliasesBody, opts ...RequestOption) (*PostAuthConnectionsIdAliasesResponse, error) {
	var out PostAuthConnectionsIdAliasesResponse
	if err := c.do(ctx, "POST", "/v1/auth/connections/"+url.PathEscape(id)+"/aliases", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostAuthConnectionsIDErase calls Erase a user (signed protocol). POST /v1/auth/connections/{id}/erase.
//
// The platform deleted a user or fulfilled an erasure request: the customer's referral data and their partner record are erased. An unknown user succeeds with nothing erased. Signed with the connection secret in `x-referalli-signature: t=<unix>,v1=…,v2=…` (v2 also signs the request line; once a connection has seen v2 it requires it), with a unique `x-referalli-id` per call. The legacy header names from before the rename are also accepted, for existing integrations. Calls older than 5 minutes, reused ids (`409 replay`), paused connections (`423 connection_paused`) and unknown fields (`422`) are refused.
//
// **Auth:** a signature (see above); no API key.
func (c *Client) PostAuthConnectionsIDErase(ctx context.Context, id string, body PostAuthConnectionsIdEraseBody, opts ...RequestOption) (*PostAuthConnectionsIdEraseResponse, error) {
	var out PostAuthConnectionsIdEraseResponse
	if err := c.do(ctx, "POST", "/v1/auth/connections/"+url.PathEscape(id)+"/erase", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostAuthConnectionsIDPartners calls Enrol a user as a partner (signed protocol). POST /v1/auth/connections/{id}/partners.
//
// A user asked to become a partner from the platform, with the code they chose: they are enrolled now (whatever `autoEnrollUsers` says), with a coupon and, when the connection has a destination, a link. 409 when the code is held by someone else. Signed like the other protocol calls (`x-referalli-signature`, `x-referalli-id`).
//
// **Auth:** none.
func (c *Client) PostAuthConnectionsIDPartners(ctx context.Context, id string, body PostAuthConnectionsIdPartnersBody, opts ...RequestOption) (*PostAuthConnectionsIdPartnersResponse, error) {
	var out PostAuthConnectionsIdPartnersResponse
	if err := c.do(ctx, "POST", "/v1/auth/connections/"+url.PathEscape(id)+"/partners", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostAuthConnectionsIDPing calls Test an auth connection (signed protocol). POST /v1/auth/connections/{id}/ping.
//
// Proves the secret works: returns the connection, its program and org and the server time, to check clock skew. Signed with the connection secret in `x-referalli-signature: t=<unix>,v1=…,v2=…` (v2 also signs the request line; once a connection has seen v2 it requires it), with a unique `x-referalli-id` per call. The legacy header names from before the rename are also accepted, for existing integrations. Calls older than 5 minutes, reused ids (`409 replay`), paused connections (`423 connection_paused`) and unknown fields (`422`) are refused.
//
// **Auth:** a signature (see above); no API key.
func (c *Client) PostAuthConnectionsIDPing(ctx context.Context, id string, body map[string]any, opts ...RequestOption) (*PostAuthConnectionsIdPingResponse, error) {
	var out PostAuthConnectionsIdPingResponse
	if err := c.do(ctx, "POST", "/v1/auth/connections/"+url.PathEscape(id)+"/ping", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostAuthConnectionsIDRevokePreviousSecret calls Revoke the previous auth secret. POST /v1/auth/connections/{id}/revoke-previous-secret.
//
// Ends the rotation overlap at once: only the newest secret verifies from now on.
//
// **Auth:** an API key with `integrations:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostAuthConnectionsIDRevokePreviousSecret(ctx context.Context, id string, opts ...RequestOption) (*PostAuthConnectionsIdRevokePreviousSecretResponse, error) {
	var out PostAuthConnectionsIdRevokePreviousSecretResponse
	if err := c.do(ctx, "POST", "/v1/auth/connections/"+url.PathEscape(id)+"/revoke-previous-secret", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostAuthConnectionsIDRotateSecret calls Rotate an auth connection secret. POST /v1/auth/connections/{id}/rotate-secret.
//
// Issues a new signing secret, returned once. The previous one keeps verifying for 24 hours, or until revoked.
//
// **Auth:** an API key with `integrations:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostAuthConnectionsIDRotateSecret(ctx context.Context, id string, opts ...RequestOption) (*PostAuthConnectionsIdRotateSecretResponse, error) {
	var out PostAuthConnectionsIdRotateSecretResponse
	if err := c.do(ctx, "POST", "/v1/auth/connections/"+url.PathEscape(id)+"/rotate-secret", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostAuthConnectionsIDSignals calls Report a user signal (signed protocol). POST /v1/auth/connections/{id}/signals.
//
// A later verdict about a user: `email_verified` releases a hold for an unverified email, `fraud_confirmed` or `self_referral` voids their referral, `banned`/`unbanned` follow on their partner record. `details` is kept as evidence. Signed with the connection secret in `x-referalli-signature: t=<unix>,v1=…,v2=…` (v2 also signs the request line; once a connection has seen v2 it requires it), with a unique `x-referalli-id` per call. The legacy header names from before the rename are also accepted, for existing integrations. Calls older than 5 minutes, reused ids (`409 replay`), paused connections (`423 connection_paused`) and unknown fields (`422`) are refused.
//
// **Auth:** a signature (see above); no API key.
func (c *Client) PostAuthConnectionsIDSignals(ctx context.Context, id string, body PostAuthConnectionsIdSignalsBody, opts ...RequestOption) (*PostAuthConnectionsIdSignalsResponse, error) {
	var out PostAuthConnectionsIdSignalsResponse
	if err := c.do(ctx, "POST", "/v1/auth/connections/"+url.PathEscape(id)+"/signals", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostAuthConnectionsIDSignups calls Attribute a new user (signed protocol). POST /v1/auth/connections/{id}/signups.
//
// The auth platform reports a new user with the referral context captured at sign-up (`referral.code`, or with cookie consent `clickId`/`anonId`) and optional risk signals. The user is attributed (risky sign-ups are held for review), optionally made a partner with their own link, and the answer says why with reason codes. `test: true` or `atlas-test: 1` simulates without storing. Repeating the same user is idempotent. Signed with the connection secret in `x-referalli-signature: t=<unix>,v1=…,v2=…` (v2 also signs the request line; once a connection has seen v2 it requires it), with a unique `x-referalli-id` per call. The legacy header names from before the rename are also accepted, for existing integrations. Calls older than 5 minutes, reused ids (`409 replay`), paused connections (`423 connection_paused`) and unknown fields (`422`) are refused.
//
// **Auth:** a signature (see above); no API key.
func (c *Client) PostAuthConnectionsIDSignups(ctx context.Context, id string, body PostAuthConnectionsIdSignupsBody, opts ...RequestOption) (*PostAuthConnectionsIdSignupsResponse, error) {
	var out PostAuthConnectionsIdSignupsResponse
	if err := c.do(ctx, "POST", "/v1/auth/connections/"+url.PathEscape(id)+"/signups", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostAuthConnectionsIDUsers calls Update a user (signed protocol). POST /v1/auth/connections/{id}/users.
//
// The user changed: email and name follow to their partner record, `suspended` suspends it, `deleted` erases. Signed with the connection secret in `x-referalli-signature: t=<unix>,v1=…,v2=…` (v2 also signs the request line; once a connection has seen v2 it requires it), with a unique `x-referalli-id` per call. The legacy header names from before the rename are also accepted, for existing integrations. Calls older than 5 minutes, reused ids (`409 replay`), paused connections (`423 connection_paused`) and unknown fields (`422`) are refused.
//
// **Auth:** a signature (see above); no API key.
func (c *Client) PostAuthConnectionsIDUsers(ctx context.Context, id string, body PostAuthConnectionsIdUsersBody, opts ...RequestOption) (*PostAuthConnectionsIdUsersResponse, error) {
	var out PostAuthConnectionsIdUsersResponse
	if err := c.do(ctx, "POST", "/v1/auth/connections/"+url.PathEscape(id)+"/users", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostAuthConnectionsInstall calls Approve a one-click install. POST /v1/auth/connections/install.
//
// The dashboard's consent step for `GET /v1/oauth/authorize`: `allow` returns the redirect back to the platform with a single-use code (10 minutes) for the chosen program; `deny` returns it with `error=access_denied`. Owner or admin. Audited.
//
// **Auth:** an API key with `integrations:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostAuthConnectionsInstall(ctx context.Context, body PostAuthConnectionsInstallBody, opts ...RequestOption) (*PostAuthConnectionsInstallResponse, error) {
	var out PostAuthConnectionsInstallResponse
	if err := c.do(ctx, "POST", "/v1/auth/connections/install", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostAuthLogout calls Sign out. POST /v1/auth/logout.
//
// Ends the browser session (and any portal-domain sessions handed off from it) and clears the cookies. Returns the identity provider's `logoutUrl` to end the provider session too, or `null`.
//
// **Auth:** none.
func (c *Client) PostAuthLogout(ctx context.Context, opts ...RequestOption) (*PostAuthLogoutResponse, error) {
	var out PostAuthLogoutResponse
	if err := c.do(ctx, "POST", "/v1/auth/logout", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostBillingActivate calls Activate a plan directly. POST /v1/billing/activate.
//
// Moves the org to a plan without checkout: anyone may move to Free (after cancelling a subscription); paid plans only for platform admins or self-hosted deployments (403 otherwise). Owner only. Audited.
//
// **Auth:** an API key with `billing:write` (or `full`), or a session with role `owner`.
func (c *Client) PostBillingActivate(ctx context.Context, body PostBillingActivateBody, opts ...RequestOption) (*PostBillingActivateResponse, error) {
	var out PostBillingActivateResponse
	if err := c.do(ctx, "POST", "/v1/billing/activate", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostBillingCheckout calls Start a plan checkout. POST /v1/billing/checkout.
//
// Creates a checkout session for a paid plan and returns its `url` to redirect to. 400 when billing is not configured on this deployment.
//
// **Auth:** an API key with `billing:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostBillingCheckout(ctx context.Context, body PostBillingCheckoutBody, opts ...RequestOption) (*PostBillingCheckoutResponse, error) {
	var out PostBillingCheckoutResponse
	if err := c.do(ctx, "POST", "/v1/billing/checkout", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostBillingConnections calls Create a billing connection. POST /v1/billing/connections.
//
// Connects a billing provider (`stripe`, `paddle`, `lemonsqueezy`, `polar`, `chargebee`, `shopify`, `recurly`, `braintree`, `woocommerce`, `bigcommerce` or `generic`) so its payment webhooks become commissions. Returns the `webhookUrl` path to register with the provider and its signing secret (generated unless you pass the provider's own `webhookSecret`); the secret is returned only here. A `filter` (`productIds`, `priceIds`, `skus` and/or exact `metadata`) limits which payments and subscription events earn when one provider account sells several products; the rest are recorded as `ignored_by_filter`. Audited.
//
// **Auth:** an API key with `billing:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostBillingConnections(ctx context.Context, body PostBillingConnectionsBody, opts ...RequestOption) (*CreatedBillingConnection, error) {
	var out CreatedBillingConnection
	if err := c.do(ctx, "POST", "/v1/billing/connections", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostBillingConnectionsIDRotateSecret calls Rotate the webhook signing secret. POST /v1/billing/connections/{id}/rotate-secret.
//
// Generates a new signing secret, returned once. The previous one keeps verifying until `previousSecretExpiresAt` (24 hours). Audited.
//
// **Auth:** an API key with `billing:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostBillingConnectionsIDRotateSecret(ctx context.Context, id string, opts ...RequestOption) (*PostBillingConnectionsIdRotateSecretResponse, error) {
	var out PostBillingConnectionsIdRotateSecretResponse
	if err := c.do(ctx, "POST", "/v1/billing/connections/"+url.PathEscape(id)+"/rotate-secret", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostBillingConnectionsIDValueMappingPreview calls Preview the value mapping. POST /v1/billing/connections/{id}/value-mapping/preview.
//
// Applies the saved mapping (or a draft `mapping`) to a sample `payload` and returns what would be recorded. Writes nothing. A generic-connector body with `events` gets one result per event.
//
// **Auth:** an API key with `billing:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostBillingConnectionsIDValueMappingPreview(ctx context.Context, id string, body PostBillingConnectionsIdValueMappingPreviewBody, opts ...RequestOption) (*ValueMappingPreview, error) {
	var out ValueMappingPreview
	if err := c.do(ctx, "POST", "/v1/billing/connections/"+url.PathEscape(id)+"/value-mapping/preview", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostBillingInternalCustomers calls Add an internal customer. POST /v1/billing/internal-customers.
//
// Marks a customer ref (the provider customer id, or the `referalli_customer_ref` it is linked to) as internal: its payments are recorded as `ignored_by_filter` from now on. Idempotent; a repeat updates the note. Audited.
//
// **Auth:** an API key with `billing:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostBillingInternalCustomers(ctx context.Context, body PostBillingInternalCustomersBody, opts ...RequestOption) (*PostBillingInternalCustomersResponse, error) {
	var out PostBillingInternalCustomersResponse
	if err := c.do(ctx, "POST", "/v1/billing/internal-customers", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostBillingPortal calls Open the subscription portal. POST /v1/billing/portal.
//
// Returns the `url` of the billing portal where the owner manages or cancels the subscription. 400 before there is a subscription.
//
// **Auth:** an API key with `billing:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostBillingPortal(ctx context.Context, opts ...RequestOption) (*PostBillingPortalResponse, error) {
	var out PostBillingPortalResponse
	if err := c.do(ctx, "POST", "/v1/billing/portal", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostBrandingDomain calls Set a custom portal domain. POST /v1/branding/domain.
//
// Starts serving the partner portal on your own domain (such as `partners.example.com`), in your branding: returns the CNAME and TXT records to add. Needs a plan with custom domains (402 otherwise).
//
// **Auth:** an API key with `org:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostBrandingDomain(ctx context.Context, body PostBrandingDomainBody, opts ...RequestOption) (*PostBrandingDomainResponse, error) {
	var out PostBrandingDomainResponse
	if err := c.do(ctx, "POST", "/v1/branding/domain", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostBrandingDomainVerify calls Verify the portal domain. POST /v1/branding/domain/verify.
//
// Checks the portal domain's TXT record and CNAME, routes it and reports `provisioning` or `active` as for the link domain.
//
// **Auth:** an API key with `org:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostBrandingDomainVerify(ctx context.Context, opts ...RequestOption) (*PostBrandingLinkDomainVerifyResponse, error) {
	var out PostBrandingLinkDomainVerifyResponse
	if err := c.do(ctx, "POST", "/v1/branding/domain/verify", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostBrandingLinkDomain calls Set a custom link domain. POST /v1/branding/link-domain.
//
// Starts using your own domain (such as `go.example.com`) for partner links: returns the CNAME target and the TXT record that proves you control it. Links move onto it once verified. Needs a plan with custom domains (402 otherwise); 409 when the domain is in use.
//
// **Auth:** an API key with `org:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostBrandingLinkDomain(ctx context.Context, body PostBrandingLinkDomainBody, opts ...RequestOption) (*PostBrandingLinkDomainResponse, error) {
	var out PostBrandingLinkDomainResponse
	if err := c.do(ctx, "POST", "/v1/branding/link-domain", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostBrandingLinkDomainVerify calls Verify the link domain. POST /v1/branding/link-domain/verify.
//
// Checks the TXT record and the CNAME. When both are right the domain is routed here; it is `active` once HTTPS answers on it (`provisioning` while the certificate is issued, which finishes on its own). 400 with what is missing otherwise.
//
// **Auth:** an API key with `org:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostBrandingLinkDomainVerify(ctx context.Context, opts ...RequestOption) (*PostBrandingLinkDomainVerifyResponse, error) {
	var out PostBrandingLinkDomainVerifyResponse
	if err := c.do(ctx, "POST", "/v1/branding/link-domain/verify", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostCliDeviceCodes calls Start a CLI sign-in. POST /v1/cli/device-codes.
//
// The CLI `login` command calls this to start a device sign-in: returns a `deviceCode` for polling and a `userCode` for the person to approve at `verificationUri`. Codes expire after 10 minutes. No authentication; limited per IP.
//
// **Auth:** none.
func (c *Client) PostCliDeviceCodes(ctx context.Context, body PostCliDeviceCodesBody, opts ...RequestOption) (*PostCliDeviceCodesResponse, error) {
	var out PostCliDeviceCodesResponse
	if err := c.do(ctx, "POST", "/v1/cli/device-codes", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostCliDeviceCodesToken calls Poll a CLI sign-in. POST /v1/cli/device-codes/token.
//
// Polled by the CLI every `interval` seconds with its `deviceCode`. Answers `pending`, `expired`, `denied` or `consumed`, or once approved, `approved` with the new API key (returned once). Polling faster than allowed is `429`.
//
// **Auth:** none.
func (c *Client) PostCliDeviceCodesToken(ctx context.Context, body PostCliDeviceCodesTokenBody, opts ...RequestOption) (*PostCliDeviceCodesTokenResponse, error) {
	var out PostCliDeviceCodesTokenResponse
	if err := c.do(ctx, "POST", "/v1/cli/device-codes/token", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostCliDeviceCodesUserCodeApprove calls Approve a CLI sign-in. POST /v1/cli/device-codes/{userCode}/approve.
//
// Approves the sign-in and mints its API key in this org. `access` picks the scopes: `cli` (webhooks write and read access to the rest, the default), `read` or `full`. Needs a signed-in owner or admin. Audited.
//
// **Auth:** an API key with `team:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostCliDeviceCodesUserCodeApprove(ctx context.Context, userCode string, body PostCliDeviceCodesUserCodeApproveBody, opts ...RequestOption) (*PostCliDeviceCodesUserCodeApproveResponse, error) {
	var out PostCliDeviceCodesUserCodeApproveResponse
	if err := c.do(ctx, "POST", "/v1/cli/device-codes/"+url.PathEscape(userCode)+"/approve", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostCliDeviceCodesUserCodeDeny calls Deny a CLI sign-in. POST /v1/cli/device-codes/{userCode}/deny.
//
// Refuses a pending sign-in; the CLI then sees `denied`.
//
// **Auth:** an API key with `team:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostCliDeviceCodesUserCodeDeny(ctx context.Context, userCode string, opts ...RequestOption) (*Ok, error) {
	var out Ok
	if err := c.do(ctx, "POST", "/v1/cli/device-codes/"+url.PathEscape(userCode)+"/deny", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostCliListeners calls Open a CLI event listener. POST /v1/cli/listeners.
//
// The CLI `listen` command opens a listener for some `topics` (default every topic). Events are queued for it while it polls at least every 2 minutes, and each comes with Standard Webhooks headers signed with the returned `secret`, so the CLI can forward them to a local server.
//
// **Auth:** an API key with `webhooks:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostCliListeners(ctx context.Context, body PostCliListenersBody, opts ...RequestOption) (*PostCliListenersResponse, error) {
	var out PostCliListenersResponse
	if err := c.do(ctx, "POST", "/v1/cli/listeners", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostCliTrigger calls Send a test event. POST /v1/cli/trigger.
//
// The CLI `trigger` command sends a realistic sample event (ids `*_test_…`, `test: true`) to the org's listeners and webhook endpoints. `data` overrides fields of the sample. 400 for a topic with no sample; `GET /v1/cli/topics` lists them.
//
// **Auth:** an API key with `webhooks:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostCliTrigger(ctx context.Context, body PostCliTriggerBody, opts ...RequestOption) (*PostCliTriggerResponse, error) {
	var out PostCliTriggerResponse
	if err := c.do(ctx, "POST", "/v1/cli/trigger", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostCommissionRules calls Publish a commission rule. POST /v1/commission-rules.
//
// Publishes a new version of a program's commission rule; the latest version applies to new payments, and past lines keep the version that earned them. A rule has a `base` (percent in basis points, or a fixed amount), what it applies to (gross or net of tax or fees), how long it recurs, a hold before lines can be paid (`holdDays`, default 30), and optional conditions, per-partner/link/product overrides, volume tiers, event bounties, sub-affiliate levels and new-vs-expansion rates. Unknown fields at any depth are rejected. Audited.
//
// **Auth:** an API key with `commissions:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostCommissionRules(ctx context.Context, body PostCommissionRulesBody, opts ...RequestOption) (*PostCommissionRulesResponse, error) {
	var out PostCommissionRulesResponse
	if err := c.do(ctx, "POST", "/v1/commission-rules", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostCommissionsApproveDue calls Approve due commissions. POST /v1/commissions/approve-due.
//
// Approves every pending line whose hold has elapsed (this also runs on a schedule). Returns how many were approved. Audited.
//
// **Auth:** an API key with `commissions:write` (or `full`), or a session with role `owner`, `admin`, `finance`.
func (c *Client) PostCommissionsApproveDue(ctx context.Context, opts ...RequestOption) (*PostCommissionsApproveDueResponse, error) {
	var out PostCommissionsApproveDueResponse
	if err := c.do(ctx, "POST", "/v1/commissions/approve-due", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostCommissionsIDAdjust calls Adjust a commission. POST /v1/commissions/{id}/adjust.
//
// Records an `adjusted` event against an earned line (a correction; the ledger is append-only). `amountMinor` is signed. An adjustment that would take the line below zero is refused unless `allowNegative`. Audited.
//
// **Auth:** an API key with `commissions:write` (or `full`), or a session with role `owner`, `admin`, `finance`.
func (c *Client) PostCommissionsIDAdjust(ctx context.Context, id string, body PostCommissionsIdAdjustBody, opts ...RequestOption) (*PostCommissionsIdAdjustResponse, error) {
	var out PostCommissionsIdAdjustResponse
	if err := c.do(ctx, "POST", "/v1/commissions/"+url.PathEscape(id)+"/adjust", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostCommissionsIDSettle calls Settle a commission yourself. POST /v1/commissions/{id}/settle.
//
// Marks an earned line paid because you credited the partner yourself (for example as store credit), with your `reference`: records `approved` (if it isn't yet) and `paid`, and publishes `commission.settled`. The line never reaches a payout run or the 1099 export. Repeating the same reference returns the original settlement; a line already paid otherwise is 409. Audited.
//
// **Auth:** an API key with `commissions:write` (or `full`), or a session with role `owner`, `admin`, `finance`.
func (c *Client) PostCommissionsIDSettle(ctx context.Context, id string, body PostCommissionsIdSettleBody, opts ...RequestOption) (*Settlement, error) {
	var out Settlement
	if err := c.do(ctx, "POST", "/v1/commissions/"+url.PathEscape(id)+"/settle", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostContracts calls Create a partner contract. POST /v1/contracts.
//
// A contract offering one partner custom terms (`termsMd`) and `compensation` (percent or fixed) in a program, as a draft or sent at once (`send: true`). When accepted, the compensation becomes a partner override on a new rule version. Audited.
//
// **Auth:** an API key with `partners:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostContracts(ctx context.Context, body PostContractsBody, opts ...RequestOption) (*PostContractsResponse, error) {
	var out PostContractsResponse
	if err := c.do(ctx, "POST", "/v1/contracts", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostContractsIDSend calls Send a contract. POST /v1/contracts/{id}/send.
//
// Sends a draft contract to the partner. 400 when it is not a draft.
//
// **Auth:** an API key with `partners:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostContractsIDSend(ctx context.Context, id string, opts ...RequestOption) (*OkStatus, error) {
	var out OkStatus
	if err := c.do(ctx, "POST", "/v1/contracts/"+url.PathEscape(id)+"/send", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostCoupons calls Create a coupon code. POST /v1/coupons.
//
// Registers a coupon code for a partner in a program, so checkouts that use it attribute to them (coupon codes win over clicks). Codes match case-insensitively and are unique per org (409).
//
// **Auth:** an API key with `links:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostCoupons(ctx context.Context, body PostCouponsBody, opts ...RequestOption) (*PostCouponsResponse, error) {
	var out PostCouponsResponse
	if err := c.do(ctx, "POST", "/v1/coupons", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostCreatorsAccountIDInvite calls Invite a creator. POST /v1/creators/{accountId}/invite.
//
// Invites a listed creator to a program with an optional message. 400 when they do not accept invitations or were already invited to that program.
//
// **Auth:** an API key with `growth:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostCreatorsAccountIDInvite(ctx context.Context, accountId string, body PostCreatorsAccountIdInviteBody, opts ...RequestOption) (*PostCreatorsAccountIdInviteResponse, error) {
	var out PostCreatorsAccountIdInviteResponse
	if err := c.do(ctx, "POST", "/v1/creators/"+url.PathEscape(accountId)+"/invite", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostDealsIDApprove calls Approve a deal. POST /v1/deals/{id}/approve.
//
// Approves a submitted deal and protects it for the partner for `protectionDays` (default 90). 404 when there's no submitted deal with that id.
//
// **Auth:** an API key with `growth:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostDealsIDApprove(ctx context.Context, id string, body PostDealsIdApproveBody, opts ...RequestOption) (*PostDealsIdApproveResponse, error) {
	var out PostDealsIdApproveResponse
	if err := c.do(ctx, "POST", "/v1/deals/"+url.PathEscape(id)+"/approve", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostDealsIDReject calls Reject a deal. POST /v1/deals/{id}/reject.
//
// Rejects a submitted deal.
//
// **Auth:** an API key with `growth:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostDealsIDReject(ctx context.Context, id string, opts ...RequestOption) (*PostDealsIdRejectResponse, error) {
	var out PostDealsIdRejectResponse
	if err := c.do(ctx, "POST", "/v1/deals/"+url.PathEscape(id)+"/reject", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostDealsIDWon calls Mark a deal won. POST /v1/deals/{id}/won.
//
// Marks an approved deal won, within its protection window: creates the referral for `customerRef` so the customer's payments earn the partner commission. Publishes `referral.created`.
//
// **Auth:** an API key with `growth:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostDealsIDWon(ctx context.Context, id string, body PostDealsIdWonBody, opts ...RequestOption) (*PostDealsIdWonResponse, error) {
	var out PostDealsIdWonResponse
	if err := c.do(ctx, "POST", "/v1/deals/"+url.PathEscape(id)+"/won", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostDiscoveryCorpus calls Ingest prospects. POST /v1/discovery/corpus.
//
// Adds up to 500 prospects from your own crawl or data source. Raw page `html` is scanned for affiliate footprints (the competitor programs it recognizes) to enrich the evidence; `keywords` scores fit. Duplicate URLs are skipped. Needs a plan with discovery (402 otherwise).
//
// **Auth:** an API key with `growth:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostDiscoveryCorpus(ctx context.Context, body PostDiscoveryCorpusBody, opts ...RequestOption) (*PostDiscoveryCorpusResponse, error) {
	var out PostDiscoveryCorpusResponse
	if err := c.do(ctx, "POST", "/v1/discovery/corpus", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostDiscoverySearch calls Search for prospective partners. POST /v1/discovery/search.
//
// Searches the configured corpus for creators and sites matching `keywords` (and `category`), scores their fit and saves them as prospects. Needs a plan with discovery (402 otherwise).
//
// **Auth:** an API key with `growth:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostDiscoverySearch(ctx context.Context, body PostDiscoverySearchBody, opts ...RequestOption) (*PostDiscoverySearchResponse, error) {
	var out PostDiscoverySearchResponse
	if err := c.do(ctx, "POST", "/v1/discovery/search", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostDisputesIDResolve calls Resolve a dispute. POST /v1/disputes/{id}/resolve.
//
// Resolves an open dispute with a `resolution` note, which the partner sees. Publishes `dispute.resolved`. 404 when not open. Audited.
//
// **Auth:** an API key with `commissions:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostDisputesIDResolve(ctx context.Context, id string, body PostDisputesIdResolveBody, opts ...RequestOption) (*Ok, error) {
	var out Ok
	if err := c.do(ctx, "POST", "/v1/disputes/"+url.PathEscape(id)+"/resolve", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostEmailPreview calls Preview an email. POST /v1/email/preview.
//
// Renders an event's email with sample values, from the stored template or an unsaved draft.
//
// **Auth:** an API key with `org:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostEmailPreview(ctx context.Context, body PostEmailPreviewBody, opts ...RequestOption) (*EmailRendered, error) {
	var out EmailRendered
	if err := c.do(ctx, "POST", "/v1/email/preview", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostEmailSettingsVerifyDomain calls Check the sending domain. POST /v1/email/settings/verify-domain.
//
// Asks the email provider whether the from address's domain is verified for sending, and stores the result. 400 until a provider, key and from address are set.
//
// **Auth:** an API key with `org:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostEmailSettingsVerifyDomain(ctx context.Context, opts ...RequestOption) (*PostEmailSettingsVerifyDomainResponse, error) {
	var out PostEmailSettingsVerifyDomainResponse
	if err := c.do(ctx, "POST", "/v1/email/settings/verify-domain", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostEmailTest calls Send a test email. POST /v1/email/test.
//
// Sends one email with sample values to `to` through the org's provider. 400 without a provider or when the provider refuses it.
//
// **Auth:** an API key with `org:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostEmailTest(ctx context.Context, body PostEmailTestBody, opts ...RequestOption) (*PostEmailTestResponse, error) {
	var out PostEmailTestResponse
	if err := c.do(ctx, "POST", "/v1/email/test", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostEmbedTokens calls Mint a partner embed token. POST /v1/embed/tokens.
//
// A 15-minute token for the partner with this email, for the partner stats web component (`/embed.js`) in your app.
//
// **Auth:** an API key with `links:write` (or `full`), or a session of any role.
func (c *Client) PostEmbedTokens(ctx context.Context, body PostEmbedTokensBody, opts ...RequestOption) (*PostEmbedTokensResponse, error) {
	var out PostEmbedTokensResponse
	if err := c.do(ctx, "POST", "/v1/embed/tokens", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostFraudCasesIDResolve calls Resolve a fraud case. POST /v1/fraud/cases/{id}/resolve.
//
// Closes a case with an `action`: `dismiss` (a held referral is approved), `void_referral`, `suspend_partner` or `ban_partner`. The matching `referral.*` events are published. Audited.
//
// **Auth:** an API key with `fraud:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostFraudCasesIDResolve(ctx context.Context, id string, body PostFraudCasesIdResolveBody, opts ...RequestOption) (*Ok, error) {
	var out Ok
	if err := c.do(ctx, "POST", "/v1/fraud/cases/"+url.PathEscape(id)+"/resolve", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostFraudDetectRings calls Detect referral rings. POST /v1/fraud/detect-rings.
//
// Looks for partners who are likely one operator with many accounts (a shared payout destination or a shared non-freemail sign-up domain) and opens a `ring` case for each group.
//
// **Auth:** an API key with `fraud:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostFraudDetectRings(ctx context.Context, opts ...RequestOption) (*PostFraudDetectRingsResponse, error) {
	var out PostFraudDetectRingsResponse
	if err := c.do(ctx, "POST", "/v1/fraud/detect-rings", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostFraudInvalidTrafficScan calls Scan for invalid traffic. POST /v1/fraud/invalid-traffic/scan.
//
// Scores each partner's clicks over the last `windowHours` (1 to 168, default 24) for invalid traffic and returns the probability, label and contributing signals per partner, plus how many were flagged.
//
// **Auth:** an API key with `fraud:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostFraudInvalidTrafficScan(ctx context.Context, body PostFraudInvalidTrafficScanBody, opts ...RequestOption) (*PostFraudInvalidTrafficScanResponse, error) {
	var out PostFraudInvalidTrafficScanResponse
	if err := c.do(ctx, "POST", "/v1/fraud/invalid-traffic/scan", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostFXRates calls Record an FX rate. POST /v1/fx-rates.
//
// Records the rate cross-currency payouts convert at, for a currency pair, as a decimal string (`"1.0825"`), effective now or at `effectiveAt`. A manual rate pins its pair over the daily feed until cleared. Audited.
//
// **Auth:** an API key with `payouts:write` (or `full`), or a session with role `owner`, `admin`, `finance`.
func (c *Client) PostFXRates(ctx context.Context, body PostFxRatesBody, opts ...RequestOption) (*FxRate, error) {
	var out FxRate
	if err := c.do(ctx, "POST", "/v1/fx-rates", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostFXRatesClearOverride calls Clear a manual FX rate. POST /v1/fx-rates/clear-override.
//
// Retires the manual rates for a pair, handing it back to the feed. Returns how many were cleared. Audited.
//
// **Auth:** an API key with `payouts:write` (or `full`), or a session with role `owner`, `admin`, `finance`.
func (c *Client) PostFXRatesClearOverride(ctx context.Context, body PostFxRatesClearOverrideBody, opts ...RequestOption) (*PostFxRatesClearOverrideResponse, error) {
	var out PostFxRatesClearOverrideResponse
	if err := c.do(ctx, "POST", "/v1/fx-rates/clear-override", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostFXRatesSync calls Sync the FX feed now. POST /v1/fx-rates/sync.
//
// Fetches the latest ECB reference rates for this org now, instead of waiting for the daily sync.
//
// **Auth:** an API key with `payouts:write` (or `full`), or a session with role `owner`, `admin`, `finance`.
func (c *Client) PostFXRatesSync(ctx context.Context, opts ...RequestOption) (*PostFxRatesSyncResponse, error) {
	var out PostFxRatesSyncResponse
	if err := c.do(ctx, "POST", "/v1/fx-rates/sync", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostGraphql calls Query the GraphQL read layer (POST). POST /v1/graphql.
//
// The GraphQL endpoint with the query in a JSON body. Read-only, so a read-scoped API key may call it.
//
// **Auth:** an API key with `analytics:read` (or `full`), or a session of any role.
func (c *Client) PostGraphql(ctx context.Context, body PostGraphqlBody, opts ...RequestOption) (*PostGraphqlResponse, error) {
	var out PostGraphqlResponse
	if err := c.do(ctx, "POST", "/v1/graphql", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostImportProviderParams are the query parameters of PostImportProvider.
type PostImportProviderParams struct {
	ProgramID      *string `json:"programId,omitempty"`
	DryRun         *bool   `json:"dryRun,omitempty"`
	DestinationURL *string `json:"destinationUrl,omitempty"`
}

func (p *PostImportProviderParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "programId", p.ProgramID)
	setQuery(v, "dryRun", p.DryRun)
	setQuery(v, "destinationUrl", p.DestinationURL)
	return v
}

// PostImportProvider calls Import partners from another tool. POST /v1/import/{provider}.
//
// Imports partners, links and codes from a Rewardful, FirstPromoter, Tolt or PartnerStack CSV export (as a `text/csv` body or `{ "csv": "…" }`) into a program. `dryRun=true` reports what would happen; `destinationUrl` is used when the export has no link column. The tool's referral parameter is added to aff.js so shared links keep crediting. Audited.
//
// **Auth:** an API key with `integrations:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostImportProvider(ctx context.Context, provider string, body PostImportProviderBody, params *PostImportProviderParams, opts ...RequestOption) (*PostImportProviderResponse, error) {
	var out PostImportProviderResponse
	if err := c.do(ctx, "POST", "/v1/import/"+url.PathEscape(provider), params.values(), body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostIntegrationsDeliveriesIDRetry calls Retry an integration delivery. POST /v1/integrations/deliveries/{id}/retry.
//
// Queues a failed or skipped delivery again. 404 when there is no failed or skipped delivery with that id.
//
// **Auth:** an API key with `integrations:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostIntegrationsDeliveriesIDRetry(ctx context.Context, id string, opts ...RequestOption) (*IntegrationDelivery, error) {
	var out IntegrationDelivery
	if err := c.do(ctx, "POST", "/v1/integrations/deliveries/"+url.PathEscape(id)+"/retry", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostIntegrationsHubspotTest calls Test the HubSpot connection. POST /v1/integrations/hubspot/test.
//
// Checks the token against HubSpot and (unless `setupProperties: false`) creates the custom properties the sync uses. Returns `ok` and the error, if any, rather than failing.
//
// **Auth:** an API key with `integrations:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostIntegrationsHubspotTest(ctx context.Context, body PostIntegrationsHubspotTestBody, opts ...RequestOption) (*PostIntegrationsHubspotTestResponse, error) {
	var out PostIntegrationsHubspotTestResponse
	if err := c.do(ctx, "POST", "/v1/integrations/hubspot/test", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostIntegrationsSegmentTest calls Test the Segment connection. POST /v1/integrations/segment/test.
//
// Sends one test track event with the write key and reports whether Segment accepted it.
//
// **Auth:** an API key with `integrations:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostIntegrationsSegmentTest(ctx context.Context, opts ...RequestOption) (*PostIntegrationsSegmentTestResponse, error) {
	var out PostIntegrationsSegmentTestResponse
	if err := c.do(ctx, "POST", "/v1/integrations/segment/test", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostInvitationsIDWithdraw calls Withdraw an invitation. POST /v1/invitations/{id}/withdraw.
//
// Withdraws a pending creator invitation. 400 when it was already answered or withdrawn.
//
// **Auth:** an API key with `partners:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostInvitationsIDWithdraw(ctx context.Context, id string, opts ...RequestOption) (*OkStatus, error) {
	var out OkStatus
	if err := c.do(ctx, "POST", "/v1/invitations/"+url.PathEscape(id)+"/withdraw", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostLinks calls Create a tracked link. POST /v1/links.
//
// Creates a partner's tracked link in a program. `code` (letters, digits, `-`; stored lowercase) is unique per org and generated when omitted. Returns the `shortUrl` to share, on the org's verified link domain when it has one. 409 when the code is taken. Audited.
//
// **Auth:** an API key with `links:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostLinks(ctx context.Context, body PostLinksBody, opts ...RequestOption) (*CreatedLink, error) {
	var out CreatedLink
	if err := c.do(ctx, "POST", "/v1/links", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostMobilePurchases calls Report an in-app purchase. POST /v1/mobile/purchases.
//
// Your backend reports a validated App Store or Play purchase (or refund), after checking the receipt; it flows through the same commission engine as web billing. Idempotent per `eventId`. Never call this from the app itself.
//
// **Auth:** an API key with `billing:write` (or `full`), or a session with role `owner`, `admin`, `finance`.
func (c *Client) PostMobilePurchases(ctx context.Context, body PostMobilePurchasesBody, opts ...RequestOption) (*PostMobilePurchasesResponse, error) {
	var out PostMobilePurchasesResponse
	if err := c.do(ctx, "POST", "/v1/mobile/purchases", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostNotificationsIDRead calls Mark a notification read. POST /v1/notifications/{id}/read.
//
// Marks one notification read. Unknown or already-read ids are a no-op.
//
// **Auth:** an API key with `org:write` (or `full`), or a session of any role.
func (c *Client) PostNotificationsIDRead(ctx context.Context, id string, opts ...RequestOption) (*Ok, error) {
	var out Ok
	if err := c.do(ctx, "POST", "/v1/notifications/"+url.PathEscape(id)+"/read", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostOauthToken calls Exchange an install code. POST /v1/oauth/token.
//
// Server to server: the platform trades the single-use code and its PKCE `code_verifier` for a new signed-protocol connection and its signing secret (returned only here). Installing again for the same instance keeps the connection and issues a new secret, the old one verifying for 24 hours. 401 for an unknown, used, expired or mismatched code. No authentication.
//
// **Auth:** none.
func (c *Client) PostOauthToken(ctx context.Context, body PostOauthTokenBody, opts ...RequestOption) (*PostOauthTokenResponse, error) {
	var out PostOauthTokenResponse
	if err := c.do(ctx, "POST", "/v1/oauth/token", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostOnchainConnections calls Create an on-chain connection. POST /v1/onchain/connections.
//
// Connects an on-chain indexer: on-chain conversions posted to its webhook URL become billing events. Returns the webhook URL and signing secret.
//
// **Auth:** an API key with `billing:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostOnchainConnections(ctx context.Context, body PostOnchainConnectionsBody, opts ...RequestOption) (*PostOnchainConnectionsResponse, error) {
	var out PostOnchainConnectionsResponse
	if err := c.do(ctx, "POST", "/v1/onchain/connections", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostOrgHostedPagesCheck calls Check the hosted-page adapters. POST /v1/org/hosted-pages/check.
//
// Read-only checks against the vendors: does the Cookie Munch site exist (and a reminder to add your portal domain as its alias), does each script URL answer. Updates `status` and `lastError`. Owner or admin.
//
// **Auth:** an API key with `org:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostOrgHostedPagesCheck(ctx context.Context, opts ...RequestOption) (*PostOrgHostedPagesCheckResponse, error) {
	var out PostOrgHostedPagesCheckResponse
	if err := c.do(ctx, "POST", "/v1/org/hosted-pages/check", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostOrgMembers calls Add a team member. POST /v1/org/members.
//
// Adds an existing Referalli user (by email) with a role: `owner`, `admin`, `finance`, `analyst` or `support`. Only an owner can add an owner. 404 when no user has that email; 409 when already a member. Live mode only.
//
// **Auth:** an API key with `team:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostOrgMembers(ctx context.Context, body PostOrgMembersBody, opts ...RequestOption) (*PostOrgMembersResponse, error) {
	var out PostOrgMembersResponse
	if err := c.do(ctx, "POST", "/v1/org/members", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostOrgPortalAuthTest calls Test partner sign-in. POST /v1/org/portal-auth/test.
//
// Checks the partner sign-in configuration (JWKS reachable, or OIDC discovery) and records the result.
//
// **Auth:** an API key with `org:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostOrgPortalAuthTest(ctx context.Context, opts ...RequestOption) (*PostOrgPortalAuthTestResponse, error) {
	var out PostOrgPortalAuthTestResponse
	if err := c.do(ctx, "POST", "/v1/org/portal-auth/test", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostOrgs calls Create an organization. POST /v1/orgs.
//
// Creates an organization with the signed-in user as its owner. Needs a user session (not an API key) but no `x-org-id`. Limited to 20 per hour per user. Live mode only.
//
// **Auth:** an API key with `team:write` (or `full`), or a session of any role.
func (c *Client) PostOrgs(ctx context.Context, body PostOrgsBody, opts ...RequestOption) (*PostOrgsResponse, error) {
	var out PostOrgsResponse
	if err := c.do(ctx, "POST", "/v1/orgs", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostOrgSSODirectTest calls Test your identity provider. POST /v1/org/sso/direct/test.
//
// Checks the configuration against the IdP (OIDC discovery, or the WorkOS organization) and records the result as the configuration's last success or last error.
//
// **Auth:** an API key with `org:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostOrgSSODirectTest(ctx context.Context, opts ...RequestOption) (*PostOrgPortalAuthTestResponse, error) {
	var out PostOrgPortalAuthTestResponse
	if err := c.do(ctx, "POST", "/v1/org/sso/direct/test", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostOrgSSOLink calls Link the org to SSO. POST /v1/org/sso/link.
//
// Creates the org's directory tenant at the SSO provider, with the calling owner (who must have signed in with the provider) as its first admin. Owner session only. 409 when already linked elsewhere. Live mode only.
//
// **Auth:** an API key with `org:write` (or `full`), or a session with role `owner`.
func (c *Client) PostOrgSSOLink(ctx context.Context, opts ...RequestOption) (*PostOrgSsoLinkResponse, error) {
	var out PostOrgSsoLinkResponse
	if err := c.do(ctx, "POST", "/v1/org/sso/link", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostOrgSSOTickets calls Create an SSO setup link. POST /v1/org/sso/tickets.
//
// A one-time link (valid `expiresInDays`, default 3) for the customer's IT admin to configure SAML or OIDC (and SCIM) on the provider's hosted page. The URL is returned once and never stored. Live mode only.
//
// **Auth:** an API key with `org:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostOrgSSOTickets(ctx context.Context, body PostOrgSsoTicketsBody, opts ...RequestOption) (*PostOrgSsoTicketsResponse, error) {
	var out PostOrgSsoTicketsResponse
	if err := c.do(ctx, "POST", "/v1/org/sso/tickets", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostOrgSSOTicketsIDRevoke calls Revoke an SSO setup link. POST /v1/org/sso/tickets/{id}/revoke.
//
// Revokes a pending setup link. Live mode only.
//
// **Auth:** an API key with `org:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostOrgSSOTicketsIDRevoke(ctx context.Context, id string, opts ...RequestOption) (*PostOrgSsoTicketsIdRevokeResponse, error) {
	var out PostOrgSsoTicketsIdRevokeResponse
	if err := c.do(ctx, "POST", "/v1/org/sso/tickets/"+url.PathEscape(id)+"/revoke", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostOutreachDraft calls Draft an outreach sequence. POST /v1/outreach/draft.
//
// Drafts a recruiting email sequence for a prospect. Drafts are only sent after you approve them.
//
// **Auth:** an API key with `growth:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostOutreachDraft(ctx context.Context, body PostOutreachDraftBody, opts ...RequestOption) (*PostOutreachDraftResponse, error) {
	var out PostOutreachDraftResponse
	if err := c.do(ctx, "POST", "/v1/outreach/draft", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostOutreachIDApprove calls Approve an outreach message. POST /v1/outreach/{id}/approve.
//
// Approves a drafted message so `POST /v1/outreach/send-due` sends it when due. Audited.
//
// **Auth:** an API key with `growth:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostOutreachIDApprove(ctx context.Context, id string, opts ...RequestOption) (*Ok, error) {
	var out Ok
	if err := c.do(ctx, "POST", "/v1/outreach/"+url.PathEscape(id)+"/approve", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostOutreachSendDue calls Send due outreach. POST /v1/outreach/send-due.
//
// Sends approved outreach messages that are due, skipping prospects without an email and addresses on the do-not-contact list (unsubscribed).
//
// **Auth:** an API key with `growth:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostOutreachSendDue(ctx context.Context, opts ...RequestOption) (*PostOutreachSendDueResponse, error) {
	var out PostOutreachSendDueResponse
	if err := c.do(ctx, "POST", "/v1/outreach/send-due", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostPartnerAdvanceAccept calls Take a commission advance. POST /v1/partner/advance/accept.
//
// Takes an advance of `amountMinor` (up to the offer). The principal plus fee is repaid by withholding future commissions from payouts.
//
// **Auth:** a partner session.
func (c *Client) PostPartnerAdvanceAccept(ctx context.Context, body PostPartnerAdvanceAcceptBody, opts ...RequestOption) (*PostPartnerAdvanceAcceptResponse, error) {
	var out PostPartnerAdvanceAcceptResponse
	if err := c.do(ctx, "POST", "/v1/partner/advance/accept", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostPartnerAuthWalletNonce calls Get a wallet sign-in nonce. POST /v1/partner/auth/wallet/nonce.
//
// Starts wallet sign-in (and AI agent registration): returns a single-use nonce, valid 10 minutes, and the message to sign with the wallet.
//
// **Auth:** none.
func (c *Client) PostPartnerAuthWalletNonce(ctx context.Context, body PostPartnerAuthWalletNonceBody, opts ...RequestOption) (*PostPartnerAuthWalletNonceResponse, error) {
	var out PostPartnerAuthWalletNonceResponse
	if err := c.do(ctx, "POST", "/v1/partner/auth/wallet/nonce", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostPartnerAuthWalletVerify calls Sign in with a wallet. POST /v1/partner/auth/wallet/verify.
//
// Verifies the wallet's signature over the nonce message and signs the partner in (creating the account on first use), setting the session cookie. 401 for an expired nonce or a signature that does not match.
//
// **Auth:** none.
func (c *Client) PostPartnerAuthWalletVerify(ctx context.Context, body PostPartnerAuthWalletVerifyBody, opts ...RequestOption) (*PostPartnerAuthWalletVerifyResponse, error) {
	var out PostPartnerAuthWalletVerifyResponse
	if err := c.do(ctx, "POST", "/v1/partner/auth/wallet/verify", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostPartnerContractsIDAccept calls Accept a contract. POST /v1/partner/contracts/{id}/accept.
//
// Accepts a sent contract; its compensation then applies to the partner's commissions in that program.
//
// **Auth:** a partner session.
func (c *Client) PostPartnerContractsIDAccept(ctx context.Context, id string, opts ...RequestOption) (*OkStatus, error) {
	var out OkStatus
	if err := c.do(ctx, "POST", "/v1/partner/contracts/"+url.PathEscape(id)+"/accept", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostPartnerContractsIDDecline calls Decline a contract. POST /v1/partner/contracts/{id}/decline.
//
// Declines a sent contract.
//
// **Auth:** a partner session.
func (c *Client) PostPartnerContractsIDDecline(ctx context.Context, id string, opts ...RequestOption) (*OkStatus, error) {
	var out OkStatus
	if err := c.do(ctx, "POST", "/v1/partner/contracts/"+url.PathEscape(id)+"/decline", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostPartnerDeals calls Register a deal. POST /v1/partner/deals.
//
// A partner registers a B2B deal (company, contact, expected value) in a program they are approved in, for the merchant to approve. Publishes `deal.submitted`.
//
// **Auth:** a partner session.
func (c *Client) PostPartnerDeals(ctx context.Context, body PostPartnerDealsBody, opts ...RequestOption) (*PostPartnerDealsResponse, error) {
	var out PostPartnerDealsResponse
	if err := c.do(ctx, "POST", "/v1/partner/deals", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostPartnerDisputes calls File a dispute. POST /v1/partner/disputes.
//
// Files a dispute with a merchant about a missing attribution or an unpaid commission, optionally for a referral. The merchant resolves it.
//
// **Auth:** a partner session.
func (c *Client) PostPartnerDisputes(ctx context.Context, body PostPartnerDisputesBody, opts ...RequestOption) (*PostPartnerDisputesResponse, error) {
	var out PostPartnerDisputesResponse
	if err := c.do(ctx, "POST", "/v1/partner/disputes", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostPartnerInvitationsIDAccept calls Accept an invitation. POST /v1/partner/invitations/{id}/accept.
//
// Accepts a pending invitation: the partner is created in the merchant's org if needed and enrolled, approved. The merchant is notified.
//
// **Auth:** a partner session.
func (c *Client) PostPartnerInvitationsIDAccept(ctx context.Context, id string, opts ...RequestOption) (*OkStatus, error) {
	var out OkStatus
	if err := c.do(ctx, "POST", "/v1/partner/invitations/"+url.PathEscape(id)+"/accept", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostPartnerInvitationsIDDecline calls Decline an invitation. POST /v1/partner/invitations/{id}/decline.
//
// Declines a pending invitation. The merchant is notified.
//
// **Auth:** a partner session.
func (c *Client) PostPartnerInvitationsIDDecline(ctx context.Context, id string, opts ...RequestOption) (*OkStatus, error) {
	var out OkStatus
	if err := c.do(ctx, "POST", "/v1/partner/invitations/"+url.PathEscape(id)+"/decline", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostPartnerKYCSubmit calls Submit identity verification. POST /v1/partner/kyc/submit.
//
// Submits the partner's legal name, country and optional date of birth and document reference for identity verification and sanctions screening. Merchants that require KYC pay only verified partners. Limited to 5 per hour per IP.
//
// **Auth:** a partner session.
func (c *Client) PostPartnerKYCSubmit(ctx context.Context, body PostPartnerKycSubmitBody, opts ...RequestOption) (*PostPartnerKycSubmitResponse, error) {
	var out PostPartnerKycSubmitResponse
	if err := c.do(ctx, "POST", "/v1/partner/kyc/submit", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostPartnerLinks calls Create a link. POST /v1/partner/links.
//
// Creates a tracked link in a program the partner is approved in (403 otherwise), optionally for one of the program's marketing assets. `code` is generated when omitted; 409 when taken.
//
// **Auth:** a partner session.
func (c *Client) PostPartnerLinks(ctx context.Context, body PostPartnerLinksBody, opts ...RequestOption) (*CreatedLink, error) {
	var out CreatedLink
	if err := c.do(ctx, "POST", "/v1/partner/links", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostPartnerNotificationsIDRead calls Mark my notification read. POST /v1/partner/notifications/{id}/read.
//
// Marks one of the partner's notifications read.
//
// **Auth:** a partner session.
func (c *Client) PostPartnerNotificationsIDRead(ctx context.Context, id string, opts ...RequestOption) (*Ok, error) {
	var out Ok
	if err := c.do(ctx, "POST", "/v1/partner/notifications/"+url.PathEscape(id)+"/read", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostPartnerOrgsOrgIDSelfBilling calls Accept self-billing. POST /v1/partner/orgs/{orgId}/self-billing.
//
// Accepts the current self-billing agreement with this merchant (`{ "accept": true }`), so the merchant issues the invoices for the partner's payouts. 404 until payout details exist.
//
// **Auth:** a partner session.
func (c *Client) PostPartnerOrgsOrgIDSelfBilling(ctx context.Context, orgId string, body PostPartnerOrgsOrgIdSelfBillingBody, opts ...RequestOption) (*PostPartnerOrgsOrgIdSelfBillingResponse, error) {
	var out PostPartnerOrgsOrgIdSelfBillingResponse
	if err := c.do(ctx, "POST", "/v1/partner/orgs/"+url.PathEscape(orgId)+"/self-billing", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostPartnerOrgsOrgIDTaxForm calls Submit my tax form. POST /v1/partner/orgs/{orgId}/tax-form.
//
// Submits a W-9, W-8BEN or W-8BEN-E to this merchant. A tax form is required before the partner is included in payouts.
//
// **Auth:** a partner session.
func (c *Client) PostPartnerOrgsOrgIDTaxForm(ctx context.Context, orgId string, body PostPartnerOrgsOrgIdTaxFormBody, opts ...RequestOption) (*Ok, error) {
	var out Ok
	if err := c.do(ctx, "POST", "/v1/partner/orgs/"+url.PathEscape(orgId)+"/tax-form", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostPartnerProgramsProgramIDApply calls Apply to a program. POST /v1/partner/programs/{programId}/apply.
//
// Applies to an active program (for example from the marketplace), creating the partner in that merchant's org if needed. The merchant approves or rejects it. Publishes `partner.applied`. 403 for blocked accounts.
//
// **Auth:** a partner session.
func (c *Client) PostPartnerProgramsProgramIDApply(ctx context.Context, programId string, opts ...RequestOption) (*PostPartnerProgramsProgramIdApplyResponse, error) {
	var out PostPartnerProgramsProgramIdApplyResponse
	if err := c.do(ctx, "POST", "/v1/partner/programs/"+url.PathEscape(programId)+"/apply", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostPartners calls Create a partner. POST /v1/partners.
//
// Creates a partner. `email` is unique per org, and so is `externalRef` (your own id for them, such as a user id). With `upsert: true` an existing partner matching the email or externalRef is returned instead of a 409 (and gets the externalRef if it had none). Over the plan's partner limit this is `402 payment_required`. Audited.
//
// **Auth:** an API key with `partners:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostPartners(ctx context.Context, body PostPartnersBody, opts ...RequestOption) (*Partner, error) {
	var out Partner
	if err := c.do(ctx, "POST", "/v1/partners", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostPartnerSecurityTotpDisable calls Turn off two-factor. POST /v1/partner/security/totp/disable.
//
// Turns two-factor off; needs a valid current `code`. 401 for a wrong code.
//
// **Auth:** a partner session.
func (c *Client) PostPartnerSecurityTotpDisable(ctx context.Context, body PostPartnerSecurityTotpDisableBody, opts ...RequestOption) (*PostPartnerSecurityTotpDisableResponse, error) {
	var out PostPartnerSecurityTotpDisableResponse
	if err := c.do(ctx, "POST", "/v1/partner/security/totp/disable", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostPartnerSecurityTotpEnable calls Turn on two-factor. POST /v1/partner/security/totp/enable.
//
// Confirms setup with a current 6-digit `code`. From then on, changing payout details needs a code in `x-totp-code`. 401 for a wrong code.
//
// **Auth:** a partner session.
func (c *Client) PostPartnerSecurityTotpEnable(ctx context.Context, body PostPartnerSecurityTotpEnableBody, opts ...RequestOption) (*PostPartnerSecurityTotpEnableResponse, error) {
	var out PostPartnerSecurityTotpEnableResponse
	if err := c.do(ctx, "POST", "/v1/partner/security/totp/enable", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostPartnerSecurityTotpSetup calls Start two-factor setup. POST /v1/partner/security/totp/setup.
//
// Generates a TOTP secret and its `otpauth://` URI for an authenticator app. 400 when two-factor is already on.
//
// **Auth:** a partner session.
func (c *Client) PostPartnerSecurityTotpSetup(ctx context.Context, opts ...RequestOption) (*PostPartnerSecurityTotpSetupResponse, error) {
	var out PostPartnerSecurityTotpSetupResponse
	if err := c.do(ctx, "POST", "/v1/partner/security/totp/setup", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostPartnersIDTaxForm calls Record a partner's tax form. POST /v1/partners/{id}/tax-form.
//
// Records a W-9, W-8BEN or W-8BEN-E for the partner (`data` is the form's fields). A tax form on file is required before a partner is included in a payout run.
//
// **Auth:** an API key with `payouts:write` (or `full`), or a session with role `owner`, `admin`, `finance`.
func (c *Client) PostPartnersIDTaxForm(ctx context.Context, id string, body PostPartnersIdTaxFormBody, opts ...RequestOption) (*Ok, error) {
	var out Ok
	if err := c.do(ctx, "POST", "/v1/partners/"+url.PathEscape(id)+"/tax-form", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostPartnerTermsVersionIDAccept calls Accept program terms. POST /v1/partner/terms/{versionId}/accept.
//
// Records the partner's acceptance of a terms version. Repeating it is a no-op.
//
// **Auth:** a partner session.
func (c *Client) PostPartnerTermsVersionIDAccept(ctx context.Context, versionId string, opts ...RequestOption) (*Ok, error) {
	var out Ok
	if err := c.do(ctx, "POST", "/v1/partner/terms/"+url.PathEscape(versionId)+"/accept", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostPassportVerify calls Verify a partner passport. POST /v1/passport/verify.
//
// Checks a partner's signed reputation passport: `valid: true` with its claims, or `valid: false` with the reason (bad signature, expired). No authentication.
//
// **Auth:** none.
func (c *Client) PostPassportVerify(ctx context.Context, body PostPassportVerifyBody, opts ...RequestOption) (*PostPassportVerifyResponse, error) {
	var out PostPassportVerifyResponse
	if err := c.do(ctx, "POST", "/v1/passport/verify", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostPayoutRuns calls Create a payout run. POST /v1/payout-runs.
//
// Builds a draft run: one payout per partner who has a payout profile, a tax form and an approved, unpaid balance over their threshold, converted into their payout currency at recorded FX, each with a self-billing invoice. Partners left out are listed in `skipped` with the reason. Nothing is paid until you execute it. Audited.
//
// **Auth:** an API key with `payouts:write` (or `full`), or a session with role `owner`, `admin`, `finance`.
func (c *Client) PostPayoutRuns(ctx context.Context, opts ...RequestOption) (*PayoutRunCreated, error) {
	var out PayoutRunCreated
	if err := c.do(ctx, "POST", "/v1/payout-runs", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostPayoutRunsIDCancel calls Cancel a payout run. POST /v1/payout-runs/{id}/cancel.
//
// Cancels a run that hasn't paid out and releases its lines back to the partners' approved balances. A run with money in flight is 409. Publishes `payout.cancelled` per payout. Audited.
//
// **Auth:** an API key with `payouts:write` (or `full`), or a session with role `owner`, `admin`, `finance`.
func (c *Client) PostPayoutRunsIDCancel(ctx context.Context, id string, opts ...RequestOption) (*PostPayoutRunsIdCancelResponse, error) {
	var out PostPayoutRunsIdCancelResponse
	if err := c.do(ctx, "POST", "/v1/payout-runs/"+url.PathEscape(id)+"/cancel", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostPayoutRunsIDExecute calls Execute a payout run. POST /v1/payout-runs/{id}/execute.
//
// Pays a draft run through the org's own connected payout providers (methods with no connection are simulated and logged; test mode uses the test rail). Returns how many were sent, submitted (settling asynchronously) and failed; each outcome publishes `payout.sent`, `payout.submitted` or `payout.failed`. Audited.
//
// **Auth:** an API key with `payouts:write` (or `full`), or a session with role `owner`, `admin`, `finance`.
func (c *Client) PostPayoutRunsIDExecute(ctx context.Context, id string, opts ...RequestOption) (*PostPayoutRunsIdExecuteResponse, error) {
	var out PostPayoutRunsIdExecuteResponse
	if err := c.do(ctx, "POST", "/v1/payout-runs/"+url.PathEscape(id)+"/execute", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostPayoutRunsRunScheduled calls Run due payout schedules. POST /v1/payout-runs/run-scheduled.
//
// Creates and executes payout runs for every schedule (org or program) that is due now. This also runs hourly on its own.
//
// **Auth:** an API key with `payouts:write` (or `full`), or a session with role `owner`, `admin`, `finance`.
func (c *Client) PostPayoutRunsRunScheduled(ctx context.Context, opts ...RequestOption) (*PostPayoutRunsRunScheduledResponse, error) {
	var out PostPayoutRunsRunScheduledResponse
	if err := c.do(ctx, "POST", "/v1/payout-runs/run-scheduled", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostPrivacyExecutorKeys calls Mint a DSAR executor key. POST /v1/privacy/executor-keys.
//
// Enables a DSAR tool for your org and mints its executor key (a bearer key, or a `whsec_` secret for Standard Webhooks), returned once. Rotating stops the old key at once. Also stores vendor settings: a Transcend public key and API key for completion callbacks, an Osano webhook secret.
//
// **Auth:** an API key with `privacy:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostPrivacyExecutorKeys(ctx context.Context, body PostPrivacyExecutorKeysBody, opts ...RequestOption) (*PostPrivacyExecutorKeysResponse, error) {
	var out PostPrivacyExecutorKeysResponse
	if err := c.do(ctx, "POST", "/v1/privacy/executor-keys", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostPrivacyExecutorsAdapterCheck calls Check a DSAR tool connection. POST /v1/privacy/executors/{adapter}/check.
//
// Whether the tool is ready: a key or the vendor credentials it needs, and no failing requests since it was set up. Answers `{ok, detail, fix}`.
//
// **Auth:** an API key with `privacy:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostPrivacyExecutorsAdapterCheck(ctx context.Context, adapter string, opts ...RequestOption) (*PostPrivacyExecutorsAdapterCheckResponse, error) {
	var out PostPrivacyExecutorsAdapterCheckResponse
	if err := c.do(ctx, "POST", "/v1/privacy/executors/"+url.PathEscape(adapter)+"/check", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostPrivacyForget calls Erase a data subject. POST /v1/privacy/forget.
//
// Erases (pseudonymizes, keeping the ledger balanced) a customer by `customerRef`, a partner by `partnerId`, or by `email` both the partner with that email and every customer stored under its hash. Send exactly one. 400 when nothing is found. Audited.
//
// **Auth:** an API key with `privacy:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostPrivacyForget(ctx context.Context, body PostPrivacyForgetBody, opts ...RequestOption) (*PostPrivacyForgetResponse, error) {
	var out PostPrivacyForgetResponse
	if err := c.do(ctx, "POST", "/v1/privacy/forget", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostPrivacyLookup calls Find a data subject by email. POST /v1/privacy/lookup.
//
// Which customers and partners an email refers to. The email is hashed on the server and never stored.
//
// **Auth:** an API key with `privacy:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostPrivacyLookup(ctx context.Context, body PostPrivacyLookupBody, opts ...RequestOption) (*PostPrivacyLookupResponse, error) {
	var out PostPrivacyLookupResponse
	if err := c.do(ctx, "POST", "/v1/privacy/lookup", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostPrivacyRequests calls Open a data-subject request. POST /v1/privacy/requests.
//
// Opens an `access` or `erasure` request for one subject (`email`, `customerRef` or `partnerId`) and answers `202` with the request to poll. Publishes `privacy.request_completed` when done.
//
// **Auth:** an API key with `privacy:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostPrivacyRequests(ctx context.Context, body PostPrivacyRequestsBody, opts ...RequestOption) (*PostPrivacyRequestsResponse, error) {
	var out PostPrivacyRequestsResponse
	if err := c.do(ctx, "POST", "/v1/privacy/requests", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostProfitabilityRebuild calls Rebuild ROI figures. POST /v1/profitability/rebuild.
//
// Recomputes every stored ROI figure from the value events with the current settings and mappings, in the background. Returns a `jobId` to poll; while one rebuild is running, the same job is returned. Audited.
//
// **Auth:** an API key with `analytics:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostProfitabilityRebuild(ctx context.Context, opts ...RequestOption) (*PostProfitabilityRebuildResponse, error) {
	var out PostProfitabilityRebuildResponse
	if err := c.do(ctx, "POST", "/v1/profitability/rebuild", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostPrograms calls Create a program. POST /v1/programs.
//
// Creates an affiliate program: its attribution model, cookie window, how long after sign-up a conversion still counts, terms, public page and settings. `slug` is unique per org. Terms are versioned so partners can be asked to accept new ones. Over the plan's program limit this is `402 payment_required`.
//
// **Auth:** an API key with `programs:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostPrograms(ctx context.Context, body PostProgramsBody, opts ...RequestOption) (*Program, error) {
	var out Program
	if err := c.do(ctx, "POST", "/v1/programs", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostProgramsIDApplyLeveling calls Apply partner levels. POST /v1/programs/{id}/apply-leveling.
//
// Promotes the program's partners to the highest level their referral count qualifies for (never demotes). Publishes `partner.level_up` per promotion.
//
// **Auth:** an API key with `programs:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostProgramsIDApplyLeveling(ctx context.Context, id string, opts ...RequestOption) (*PostProgramsIdApplyLevelingResponse, error) {
	var out PostProgramsIdApplyLevelingResponse
	if err := c.do(ctx, "POST", "/v1/programs/"+url.PathEscape(id)+"/apply-leveling", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostProgramsIDAutopriceRun calls Run autopricing. POST /v1/programs/{id}/autoprice/run.
//
// Moves each partner's rate toward the target and publishes the new rates as partner overrides on a new rule version. Objective `roas` compares attributed revenue to commission; `roi` uses each partner's matured referral ROI, holding partners with too few matured referrals or an unknown ROI (listed in `held`). 400 when autopricing is off, the program has no percent-based rule, or no partner has a known ROI under `roi`.
//
// **Auth:** an API key with `programs:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostProgramsIDAutopriceRun(ctx context.Context, id string, body PostProgramsIdAutopriceRunBody, opts ...RequestOption) (*PostProgramsIdAutopriceRunResponse, error) {
	var out PostProgramsIdAutopriceRunResponse
	if err := c.do(ctx, "POST", "/v1/programs/"+url.PathEscape(id)+"/autoprice/run", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostProgramsIDEnrollments calls Enroll a partner. POST /v1/programs/{id}/enrollments.
//
// Enrolls an existing partner in the program, approved at once. Publishes `partner.enrolled`. 409 when already enrolled.
//
// **Auth:** an API key with `programs:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostProgramsIDEnrollments(ctx context.Context, id string, body PostProgramsIdEnrollmentsBody, opts ...RequestOption) (*Enrollment, error) {
	var out Enrollment
	if err := c.do(ctx, "POST", "/v1/programs/"+url.PathEscape(id)+"/enrollments", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostProgramsIDEnrollmentsEnrollmentIDApprove calls Approve an application. POST /v1/programs/{id}/enrollments/{enrollmentId}/approve.
//
// Approves a pending application (`status: "applied"`). Publishes `partner.approved`. 404 when there is no pending application with that id. Audited.
//
// **Auth:** an API key with `programs:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostProgramsIDEnrollmentsEnrollmentIDApprove(ctx context.Context, id string, enrollmentId string, opts ...RequestOption) (*PostProgramsIdEnrollmentsEnrollmentIdApproveResponse, error) {
	var out PostProgramsIdEnrollmentsEnrollmentIdApproveResponse
	if err := c.do(ctx, "POST", "/v1/programs/"+url.PathEscape(id)+"/enrollments/"+url.PathEscape(enrollmentId)+"/approve", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostProgramsIDEnrollmentsEnrollmentIDReject calls Reject an application. POST /v1/programs/{id}/enrollments/{enrollmentId}/reject.
//
// Rejects a pending application. 404 when there is no pending application with that id. Audited.
//
// **Auth:** an API key with `programs:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostProgramsIDEnrollmentsEnrollmentIDReject(ctx context.Context, id string, enrollmentId string, opts ...RequestOption) (*PostProgramsIdEnrollmentsEnrollmentIdRejectResponse, error) {
	var out PostProgramsIdEnrollmentsEnrollmentIdRejectResponse
	if err := c.do(ctx, "POST", "/v1/programs/"+url.PathEscape(id)+"/enrollments/"+url.PathEscape(enrollmentId)+"/reject", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostProgramsIDGroups calls Create a partner group. POST /v1/programs/{id}/groups.
//
// A group of partners in a program with its own commission base (or `null` to keep the program rule's), hold days, payout threshold, payout currency, onboarding message and checklist. Names are unique per program (409).
//
// **Auth:** an API key with `programs:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostProgramsIDGroups(ctx context.Context, id string, body PostProgramsIdGroupsBody, opts ...RequestOption) (*PartnerGroup, error) {
	var out PartnerGroup
	if err := c.do(ctx, "POST", "/v1/programs/"+url.PathEscape(id)+"/groups", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostProgramsIDPromotions calls Create a promotion. POST /v1/programs/{id}/promotions.
//
// A time-boxed commission change for the whole program or one group: `replace` the base rate, or `boost` it by `bps`, between `startsAt` and `endsAt`. Payments in that window get it.
//
// **Auth:** an API key with `programs:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostProgramsIDPromotions(ctx context.Context, id string, body PostProgramsIdPromotionsBody, opts ...RequestOption) (*Promotion, error) {
	var out Promotion
	if err := c.do(ctx, "POST", "/v1/programs/"+url.PathEscape(id)+"/promotions", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostPromotionsIDEnd calls End a promotion. POST /v1/promotions/{id}/end.
//
// Ends a promotion now: payments from this moment on no longer get it.
//
// **Auth:** an API key with `programs:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostPromotionsIDEnd(ctx context.Context, id string, opts ...RequestOption) (*Promotion, error) {
	var out Promotion
	if err := c.do(ctx, "POST", "/v1/promotions/"+url.PathEscape(id)+"/end", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostPublicProgramsHandleJoin calls Join a program. POST /v1/public/programs/{handle}/join.
//
// One-step join from the public page for a signed-in partner: creates (or claims, with a verified email) the partner in the merchant's org and enrolls them, approved with a first link when the program auto-approves, otherwise as an application. Publishes `partner.approved` or `partner.applied`. 409 when the merchant's partner with that email belongs to another account.
//
// **Auth:** a partner session.
func (c *Client) PostPublicProgramsHandleJoin(ctx context.Context, handle string, body PostPublicProgramsHandleJoinBody, opts ...RequestOption) (*PostPublicProgramsHandleJoinResponse, error) {
	var out PostPublicProgramsHandleJoinResponse
	if err := c.do(ctx, "POST", "/v1/public/programs/"+url.PathEscape(handle)+"/join", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostQualityRecompute calls Recompute quality scores. POST /v1/quality/recompute.
//
// Snapshots every partner's quality score now; per the org's policy, low scores open a `quality` case or suspend the partner. Audited.
//
// **Auth:** an API key with `partners:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostQualityRecompute(ctx context.Context, opts ...RequestOption) (*PostQualityRecomputeResponse, error) {
	var out PostQualityRecomputeResponse
	if err := c.do(ctx, "POST", "/v1/quality/recompute", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostReferralsAssociate calls Attribute a new customer. POST /v1/referrals/associate.
//
// The key server-side call: tell Referalli a customer signed up, with whatever you know about where they came from (`referralToken`/`ref_click_id`, `anonId`, `couponCode`, `emailHash`, IP and user agent). The attribution engine picks the partner under the program's model and windows. `referral` is `null` when no partner matched (not an error); `reasoning.notes` says why. Repeating `idempotencyKey` returns the existing referral. `occurredAt` records a late call at the real sign-up time (up to 5 minutes of clock skew is tolerated) and backfills payments since. `emailDomain` lets burner-email sign-ups be held for review. Publishes `referral.created` (or `referral.held`).
//
// **Auth:** an API key with `referrals:write` (or `full`), or a session of any role.
func (c *Client) PostReferralsAssociate(ctx context.Context, body PostReferralsAssociateBody, opts ...RequestOption) (*AssociateResponse, error) {
	var out AssociateResponse
	if err := c.do(ctx, "POST", "/v1/referrals/associate", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostReferralsIDApprove calls Approve a held referral. POST /v1/referrals/{id}/approve.
//
// Releases a referral held for review (`pending_review`) back to `active`. 404 when it isn't pending review. Publishes `referral.approved`. Audited.
//
// **Auth:** an API key with `referrals:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostReferralsIDApprove(ctx context.Context, id string, opts ...RequestOption) (*Referral, error) {
	var out Referral
	if err := c.do(ctx, "POST", "/v1/referrals/"+url.PathEscape(id)+"/approve", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostReferralsIDReassign calls Reassign a referral. POST /v1/referrals/{id}/reassign.
//
// Moves the referral to another partner (recorded as a manual match). Publishes `referral.reassigned` with the previous partner. Audited.
//
// **Auth:** an API key with `referrals:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostReferralsIDReassign(ctx context.Context, id string, body PostReferralsIdReassignBody, opts ...RequestOption) (*Referral, error) {
	var out Referral
	if err := c.do(ctx, "POST", "/v1/referrals/"+url.PathEscape(id)+"/reassign", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostReferralsIDVoid calls Void a referral. POST /v1/referrals/{id}/void.
//
// Voids the referral so it earns nothing more. Publishes `referral.voided`. Audited.
//
// **Auth:** an API key with `referrals:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostReferralsIDVoid(ctx context.Context, id string, opts ...RequestOption) (*Referral, error) {
	var out Referral
	if err := c.do(ctx, "POST", "/v1/referrals/"+url.PathEscape(id)+"/void", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostRewards calls Reward a partner. POST /v1/rewards.
//
// Rewards a partner for something outside a sale (a community answer, an event): one `earned` commission line that is approved and paid like any other. Name the partner by `partnerId` or `partnerExternalRef`; the amount is `amountMinor` (+ `currency`) or comes from a rule (`ruleKey`, or the `event` a rule maps). Idempotent on `idempotencyKey`: a retry returns the first grant with 200; the same key with a different request is 409.
//
// **Auth:** an API key with `commissions:write` (or `full`), or a session with role `owner`, `admin`, `finance`.
func (c *Client) PostRewards(ctx context.Context, body PostRewardsBody, opts ...RequestOption) (*PostRewardsResponse, error) {
	var out PostRewardsResponse
	if err := c.do(ctx, "POST", "/v1/rewards", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostRewardsIDRedeem calls Redeem a reward. POST /v1/rewards/{id}/redeem.
//
// Records a (possibly partial) redemption: `amountMinor` (default the remaining balance) is capped at what's left, and the reward becomes `redeemed` once it is used up. `externalId` makes it idempotent (409 if reused for another reward). Publishes `credit.redeemed`.
//
// **Auth:** an API key with `commissions:write` (or `full`), or a session with role `owner`, `admin`, `finance`.
func (c *Client) PostRewardsIDRedeem(ctx context.Context, id string, body PostRewardsIdRedeemBody, opts ...RequestOption) (*PostRewardsRedeemByCodeResponse, error) {
	var out PostRewardsRedeemByCodeResponse
	if err := c.do(ctx, "POST", "/v1/rewards/"+url.PathEscape(id)+"/redeem", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostRewardsRedeemByCode calls Redeem a reward by code. POST /v1/rewards/redeem-by-code.
//
// Records a redemption of the reward with this code; see `POST /v1/rewards/{id}/redeem`. 404 for an unknown code.
//
// **Auth:** an API key with `commissions:write` (or `full`), or a session with role `owner`, `admin`, `finance`.
func (c *Client) PostRewardsRedeemByCode(ctx context.Context, body PostRewardsRedeemByCodeBody, opts ...RequestOption) (*PostRewardsRedeemByCodeResponse, error) {
	var out PostRewardsRedeemByCodeResponse
	if err := c.do(ctx, "POST", "/v1/rewards/redeem-by-code", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostRoiPatternsPatternReview calls Open a review for a pattern match. POST /v1/roi/patterns/{pattern}/review.
//
// Turns one live pattern match (by `customerRef` and/or `partnerId`, optionally `programId`) into a `value_pattern` fraud case against its partner, or returns the open case for the same match. Publishes `fraud.case_opened`. 404 when no such match exists. Audited.
//
// **Auth:** an API key with `analytics:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostRoiPatternsPatternReview(ctx context.Context, pattern string, body PostRoiPatternsPatternReviewBody, opts ...RequestOption) (*FraudCase, error) {
	var out FraudCase
	if err := c.do(ctx, "POST", "/v1/roi/patterns/"+url.PathEscape(pattern)+"/review", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostSandboxSimulate calls Simulate billing history. POST /v1/sandbox/simulate.
//
// Generates deterministic billing history through the real commission pipeline, so you can see commissions flow without a provider: a new sandbox customer referred to the partner, `months` (default 3) of subscription payments at `priceMinor`, then for `scenario: "refund"` a half refund of the first invoice or for `churn` a cancellation. Returns the customer and what was created. Audited.
//
// **Auth:** an API key with `billing:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostSandboxSimulate(ctx context.Context, body PostSandboxSimulateBody, opts ...RequestOption) (*PostSandboxSimulateResponse, error) {
	var out PostSandboxSimulateResponse
	if err := c.do(ctx, "POST", "/v1/sandbox/simulate", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostSupportAppsAdapterTest calls Preview a support card. POST /v1/support/apps/{adapter}/test.
//
// Renders what the tool would show for a contact, from your own session: the context and the tool-format card. Nothing is sent to the tool.
//
// **Auth:** an API key with `partners:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostSupportAppsAdapterTest(ctx context.Context, adapter string, body PostSupportAppsAdapterTestBody, opts ...RequestOption) (*PostSupportAppsAdapterTestResponse, error) {
	var out PostSupportAppsAdapterTestResponse
	if err := c.do(ctx, "POST", "/v1/support/apps/"+url.PathEscape(adapter)+"/test", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostSupportCardsOrgIDAdapter calls Render a support card for a tool. POST /v1/support/cards/{orgId}/{adapter}.
//
// Called by the support tool, signed in its own scheme (Chatly `x-chatly-signature`, Intercom `X-Body-Signature`, HubSpot v3, Front, Help Scout, a Zendesk app JWT, or the custom `X-Referalli-Signature`). Answers the tool's card format. An unknown org, a tool that is not connected, or a bad or stale signature all answer 404.
//
// **Auth:** none.
func (c *Client) PostSupportCardsOrgIDAdapter(ctx context.Context, orgId string, adapter string, body map[string]any, opts ...RequestOption) (*PostSupportCardsOrgIdAdapterResponse, error) {
	var out PostSupportCardsOrgIdAdapterResponse
	if err := c.do(ctx, "POST", "/v1/support/cards/"+url.PathEscape(orgId)+"/"+url.PathEscape(adapter), nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostTax1099DocumentsIDCorrect calls File a corrected 1099. POST /v1/tax/1099/documents/{id}/correct.
//
// Files a corrected form for one document with the partner's current amount. Audited.
//
// **Auth:** an API key with `payouts:write` (or `full`), or a session with role `owner`, `admin`, `finance`.
func (c *Client) PostTax1099DocumentsIDCorrect(ctx context.Context, id string, opts ...RequestOption) (*TaxFiling, error) {
	var out TaxFiling
	if err := c.do(ctx, "POST", "/v1/tax/1099/documents/"+url.PathEscape(id)+"/correct", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostTax1099Filings calls File 1099-NEC forms. POST /v1/tax/1099/filings.
//
// E-files the year's 1099-NEC forms through the connected provider and records one document per partner. Audited.
//
// **Auth:** an API key with `payouts:write` (or `full`), or a session with role `owner`, `admin`, `finance`.
func (c *Client) PostTax1099Filings(ctx context.Context, body PostTax1099FilingsBody, opts ...RequestOption) (*TaxFiling, error) {
	var out TaxFiling
	if err := c.do(ctx, "POST", "/v1/tax/1099/filings", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostTax1099FilingsIDRefresh calls Refresh a filing status. POST /v1/tax/1099/filings/{id}/refresh.
//
// Asks the filing provider for the latest status of a filing and its documents; accepted forms are then delivered to partners (`tax.form_delivered`).
//
// **Auth:** an API key with `payouts:write` (or `full`), or a session with role `owner`, `admin`, `finance`.
func (c *Client) PostTax1099FilingsIDRefresh(ctx context.Context, id string, opts ...RequestOption) (*PostTax1099FilingsIdRefreshResponse, error) {
	var out PostTax1099FilingsIdRefreshResponse
	if err := c.do(ctx, "POST", "/v1/tax/1099/filings/"+url.PathEscape(id)+"/refresh", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostTestClockAdvance calls Advance the test clock. POST /v1/test_clock/advance.
//
// Test mode only. Moves the test org's clock forward by `seconds` (never back), then runs what came due: held commissions are approved, payout schedules fire and retention purges run. Returns the new time and what ran. Audited.
//
// **Auth:** an API key with `org:write` (or `full`), or a session with role `owner`, `admin`, `finance`.
func (c *Client) PostTestClockAdvance(ctx context.Context, body PostTestClockAdvanceBody, opts ...RequestOption) (*PostTestClockAdvanceResponse, error) {
	var out PostTestClockAdvanceResponse
	if err := c.do(ctx, "POST", "/v1/test_clock/advance", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostTrackEvents calls Track an event. POST /v1/track/events.
//
// Records one event. `type: "event"` is a product event for a known customer (`event`, such as `trial_started`) that can pay a bounty under the program's commission rule; it returns whether it was recorded and the commission. `idempotencyKey` deduplicates. `type: "click"` records a server-side click for a partner; `type: "custom"` records an audited custom event. Publishable keys may send only custom events.
//
// **Auth:** an API key with `events:write` (or `full`), or a session of any role.
func (c *Client) PostTrackEvents(ctx context.Context, body json.RawMessage, opts ...RequestOption) (*PostTrackEventsResponse, error) {
	var out PostTrackEventsResponse
	if err := c.do(ctx, "POST", "/v1/track/events", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostUnsubscribeToken calls Unsubscribe from outreach. POST /v1/unsubscribe/{token}.
//
// The one-click unsubscribe link in outreach emails. No authentication; the token identifies the recipient.
//
// **Auth:** none.
func (c *Client) PostUnsubscribeToken(ctx context.Context, token string, opts ...RequestOption) (*PostUnsubscribeTokenResponse, error) {
	var out PostUnsubscribeTokenResponse
	if err := c.do(ctx, "POST", "/v1/unsubscribe/"+url.PathEscape(token), nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostValueEvents calls Record value events. POST /v1/value/events.
//
// Reports up to 500 value facts from your platform: `payment`, `refund`, `credit_granted`, `credit_redeemed`, `credit_expired`, `credit_revoked`, `cost` or `adjustment`. Each needs `kind`, `currency`, `externalId` (your stable id, which makes it idempotent) and `occurredAt`, plus the customer, partner (`partnerId` or `partnerExternalRef`), grant, program, product and amount fields (`cashMinor`, `creditAppliedMinor`, `discountMinor`, `taxMinor`, `feesMinor`, `costMinor`, `faceMinor`) that apply. Each event is validated and recorded on its own: `results` and `errors` are by index. A repeat of an `externalId` with new figures supersedes the earlier one.
//
// **Auth:** an API key with `analytics:write` (or `full`), or a session of any role.
func (c *Client) PostValueEvents(ctx context.Context, body PostValueEventsBody, opts ...RequestOption) (*PostValueEventsResponse, error) {
	var out PostValueEventsResponse
	if err := c.do(ctx, "POST", "/v1/value/events", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostValueImport calls Import value events from CSV. POST /v1/value/import.
//
// Imports value events from CSV text (up to 10,000 rows and 5 MB, else `413 payload_too_large`) with a `mapping` from column names to event fields (money as `<field>Minor` or `<field>Major`). `kind` applies to every row unless a column maps to it. Up to 500 rows are recorded before the response; a larger file answers `202` with a `jobId` to poll. `dryRun` returns the first 20 parsed rows and every row error without recording.
//
// **Auth:** an API key with `analytics:write` (or `full`), or a session of any role.
func (c *Client) PostValueImport(ctx context.Context, body PostValueImportBody, opts ...RequestOption) (*PostValueImportResponse, error) {
	var out PostValueImportResponse
	if err := c.do(ctx, "POST", "/v1/value/import", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostWebhookDeliveriesIDReplay calls Replay a webhook delivery. POST /v1/webhook-deliveries/{id}/replay.
//
// Queues the delivery again at once with its attempt count reset. It keeps its `webhook-id`, so receivers can deduplicate.
//
// **Auth:** an API key with `webhooks:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostWebhookDeliveriesIDReplay(ctx context.Context, id string, opts ...RequestOption) (*Ok, error) {
	var out Ok
	if err := c.do(ctx, "POST", "/v1/webhook-deliveries/"+url.PathEscape(id)+"/replay", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostWebhookEndpoints calls Create a webhook endpoint. POST /v1/webhook-endpoints.
//
// Subscribes an HTTPS URL to event topics (default `["*"]`, every event; see the `webhooks` section of this document for each topic's payload). `programIds` limits it to events about those programs (org-wide events such as payouts are then not sent). Deliveries are signed the Standard Webhooks way with the returned `secret` (`whsec_…`), shown only here. `legacySignature: true` also sends the deprecated body-only signature. `alertContacts` (up to 10 emails) are emailed, with the org owner, when the endpoint has failed for 1 hour and again at 24 hours; it is disabled after 72 hours of failures or 1,000 failed attempts in a row.
//
// **Auth:** an API key with `webhooks:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostWebhookEndpoints(ctx context.Context, body PostWebhookEndpointsBody, opts ...RequestOption) (*PostWebhookEndpointsResponse, error) {
	var out PostWebhookEndpointsResponse
	if err := c.do(ctx, "POST", "/v1/webhook-endpoints", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostWebhookEndpointsIDReplayParams are the query parameters of PostWebhookEndpointsIDReplay.
type PostWebhookEndpointsIDReplayParams struct {
	Since *string `json:"since,omitempty"`
}

func (p *PostWebhookEndpointsIDReplayParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	setQuery(v, "since", p.Since)
	return v
}

// PostWebhookEndpointsIDReplay calls Replay failed deliveries. POST /v1/webhook-endpoints/{id}/replay.
//
// Requeues the endpoint's failed deliveries created at or after `since` (ISO 8601), for after an outage. Each keeps its `webhook-id`, so receivers can dedupe. Audited.
//
// **Auth:** an API key with `webhooks:write` (or `full`), or a session with role `owner`, `admin`.
//
// since: Requeue failed deliveries created at or after this time.
func (c *Client) PostWebhookEndpointsIDReplay(ctx context.Context, id string, params *PostWebhookEndpointsIDReplayParams, opts ...RequestOption) (*PostWebhookEndpointsIdReplayResponse, error) {
	var out PostWebhookEndpointsIdReplayResponse
	if err := c.do(ctx, "POST", "/v1/webhook-endpoints/"+url.PathEscape(id)+"/replay", params.values(), nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostWebhookEndpointsIDRotateSecret calls Rotate an endpoint secret. POST /v1/webhook-endpoints/{id}/rotate-secret.
//
// Issues a new signing secret, returned once. For 24 hours (`previousSecretExpiresAt`) every delivery carries a signature from each secret, so receivers can switch without dropping any. Audited.
//
// **Auth:** an API key with `webhooks:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostWebhookEndpointsIDRotateSecret(ctx context.Context, id string, opts ...RequestOption) (*PostWebhookEndpointsIdRotateSecretResponse, error) {
	var out PostWebhookEndpointsIdRotateSecretResponse
	if err := c.do(ctx, "POST", "/v1/webhook-endpoints/"+url.PathEscape(id)+"/rotate-secret", nil, nil, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostWidgetTokens calls Mint a referral widget token. POST /v1/widget/tokens.
//
// Your backend mints a short-lived token (60 to 3600 seconds, default 900) for a signed-in customer, identified by `customerRef` or `partnerExternalRef`, for the in-app referral widget. When the program has `settings.widget.autoEnroll`, a customer who isn't a partner yet is created on demand (send `email`). The token only reads that partner's own data.
//
// **Auth:** an API key with `partners:write` (or `full`), or a session of any role.
func (c *Client) PostWidgetTokens(ctx context.Context, body PostWidgetTokensBody, opts ...RequestOption) (*PostWidgetTokensResponse, error) {
	var out PostWidgetTokensResponse
	if err := c.do(ctx, "POST", "/v1/widget/tokens", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostZapierConversions calls Record a conversion from Zapier. POST /v1/zapier/conversions.
//
// Zapier's "Record conversion" action: a payment (or refund) for a customer through a generic billing connection (the org's first active one unless `connectionId`), with no signature needed. Idempotent by `providerEventId`. 400 without a generic billing connection.
//
// **Auth:** an API key with `billing:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostZapierConversions(ctx context.Context, body PostZapierConversionsBody, opts ...RequestOption) (*PostZapierConversionsResponse, error) {
	var out PostZapierConversionsResponse
	if err := c.do(ctx, "POST", "/v1/zapier/conversions", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PostZapierHooks calls Subscribe a Zapier hook. POST /v1/zapier/hooks.
//
// REST hook subscribe, called by Zapier when a Zap turns on: events of `event` are posted to the HTTPS `target_url`.
//
// **Auth:** an API key with `integrations:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PostZapierHooks(ctx context.Context, body PostZapierHooksBody, opts ...RequestOption) (*ZapierHook, error) {
	var out ZapierHook
	if err := c.do(ctx, "POST", "/v1/zapier/hooks", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PutAuthConnectionsIDSecret calls Set a vendor-issued secret. PUT /v1/auth/connections/{id}/secret.
//
// For webhook connectors whose vendor issues the signing secret (Clerk, Stytch, WorkOS) or key set (Kinde's JWKS document): replaces the connection's secret. The previous one keeps verifying for 24 hours. Signed-protocol connections rotate instead. Audited.
//
// **Auth:** an API key with `integrations:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PutAuthConnectionsIDSecret(ctx context.Context, id string, body PutAuthConnectionsIdSecretBody, opts ...RequestOption) (*PutAuthConnectionsIdSecretResponse, error) {
	var out PutAuthConnectionsIdSecretResponse
	if err := c.do(ctx, "PUT", "/v1/auth/connections/"+url.PathEscape(id)+"/secret", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PutBillingConnectionsIDCredentials calls Set provider API credentials. PUT /v1/billing/connections/{id}/credentials.
//
// For providers that send thin webhooks (Recurly, BigCommerce), the API credentials used to fetch the full event: `apiKey`, or `apiToken` plus `storeHash`, and `region`. Stored encrypted. 400 for providers that take no credentials or when a required field is missing. Audited.
//
// **Auth:** an API key with `billing:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PutBillingConnectionsIDCredentials(ctx context.Context, id string, body PutBillingConnectionsIdCredentialsBody, opts ...RequestOption) (*PutBillingConnectionsIdCredentialsResponse, error) {
	var out PutBillingConnectionsIdCredentialsResponse
	if err := c.do(ctx, "PUT", "/v1/billing/connections/"+url.PathEscape(id)+"/credentials", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PutBillingConnectionsIDSecret calls Set the webhook signing secret. PUT /v1/billing/connections/{id}/secret.
//
// Replaces the connection's signing secret with the provider's own (at least 8 characters). The previous secret keeps verifying for 24 hours. Audited.
//
// **Auth:** an API key with `billing:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PutBillingConnectionsIDSecret(ctx context.Context, id string, body PutBillingConnectionsIdSecretBody, opts ...RequestOption) (*PutBillingConnectionsIdSecretResponse, error) {
	var out PutBillingConnectionsIdSecretResponse
	if err := c.do(ctx, "PUT", "/v1/billing/connections/"+url.PathEscape(id)+"/secret", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PutBillingConnectionsIDValueMapping calls Set the value mapping. PUT /v1/billing/connections/{id}/value-mapping.
//
// Where cash, credit applied, discount, cost and product live in this connection's payloads, as dot paths (`[]` sums an array, for example `data.discounts[].amount`), and whether amounts are in minor or major units. Each save bumps the mapping version; a rebuild re-applies it to stored events. Audited.
//
// **Auth:** an API key with `billing:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PutBillingConnectionsIDValueMapping(ctx context.Context, id string, body PutBillingConnectionsIdValueMappingBody, opts ...RequestOption) (*PutBillingConnectionsIdValueMappingResponse, error) {
	var out PutBillingConnectionsIdValueMappingResponse
	if err := c.do(ctx, "PUT", "/v1/billing/connections/"+url.PathEscape(id)+"/value-mapping", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PutBranding calls Update branding. PUT /v1/branding.
//
// Sets the logo URL, theme (only the documented `--r-*` tokens, with plain CSS values) and email sender. Omitted fields keep their value. `warnings` lists theme text/background pairs (`--r-accent-ink` on `--r-accent`, `--r-ink` on `--r-paper`) below the 4.5:1 WCAG AA contrast; the theme is saved anyway.
//
// **Auth:** an API key with `org:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PutBranding(ctx context.Context, body PutBrandingBody, opts ...RequestOption) (*PutBrandingResponse, error) {
	var out PutBrandingResponse
	if err := c.do(ctx, "PUT", "/v1/branding", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PutEmailSettings calls Update email settings. PUT /v1/email/settings.
//
// Sets the email provider and its API key (write-only: omit to keep it, `null` to remove), the from address and name, reply-to, logo and primary color for partner emails. Without a provider, no email is sent.
//
// **Auth:** an API key with `org:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PutEmailSettings(ctx context.Context, body PutEmailSettingsBody, opts ...RequestOption) (*EmailSettings, error) {
	var out EmailSettings
	if err := c.do(ctx, "PUT", "/v1/email/settings", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PutEmailTemplatesEventLocale calls Override an email template. PUT /v1/email/templates/{event}/{locale}.
//
// Replaces the subject, text and optional HTML of one event in one locale. 404 for an unknown event or locale.
//
// **Auth:** an API key with `org:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PutEmailTemplatesEventLocale(ctx context.Context, event string, locale string, body PutEmailTemplatesEventLocaleBody, opts ...RequestOption) (*PutEmailTemplatesEventLocaleResponse, error) {
	var out PutEmailTemplatesEventLocaleResponse
	if err := c.do(ctx, "PUT", "/v1/email/templates/"+url.PathEscape(event)+"/"+url.PathEscape(locale), nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PutIntegrationsHubspot calls Connect HubSpot. PUT /v1/integrations/hubspot.
//
// Connects HubSpot with your private-app `token` (required the first time; omit to keep it) to sync partners as contacts and referrals as deals in `pipeline`, moving them to `paidDealStage` when they pay. `status: "paused"` stops syncing.
//
// **Auth:** an API key with `integrations:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PutIntegrationsHubspot(ctx context.Context, body PutIntegrationsHubspotBody, opts ...RequestOption) (*IntegrationConnection, error) {
	var out IntegrationConnection
	if err := c.do(ctx, "PUT", "/v1/integrations/hubspot", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PutIntegrationsSegment calls Connect Segment. PUT /v1/integrations/segment.
//
// Connects a Segment source with its `writeKey` (required the first time) and sends referral, commission, payout and partner events as track calls; `events` limits which topics (`null` for all).
//
// **Auth:** an API key with `integrations:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PutIntegrationsSegment(ctx context.Context, body PutIntegrationsSegmentBody, opts ...RequestOption) (*IntegrationConnection, error) {
	var out IntegrationConnection
	if err := c.do(ctx, "PUT", "/v1/integrations/segment", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PutMarketplaceListing calls List a program in the marketplace. PUT /v1/marketplace/listing.
//
// Creates or updates the marketplace listing of a program: visibility (`public` or `unlisted`), category, headline, description and sponsored placement. Audited.
//
// **Auth:** an API key with `growth:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PutMarketplaceListing(ctx context.Context, body PutMarketplaceListingBody, opts ...RequestOption) (*Ok, error) {
	var out Ok
	if err := c.do(ctx, "PUT", "/v1/marketplace/listing", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PutMeLocale calls Set my email language. PUT /v1/me/locale.
//
// The signed-in user's preferred locale for emails. User session only.
//
// **Auth:** an API key with `team:write` (or `full`), or a session of any role.
func (c *Client) PutMeLocale(ctx context.Context, body PutMeLocaleBody, opts ...RequestOption) (*PutMeLocaleResponse, error) {
	var out PutMeLocaleResponse
	if err := c.do(ctx, "PUT", "/v1/me/locale", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PutMeRoiDigest calls Turn my ROI digest on or off. PUT /v1/me/roi-digest.
//
// Whether the signed-in member receives this org's periodic credit ROI digest email. User session only.
//
// **Auth:** an API key with `team:write` (or `full`), or a session of any role.
func (c *Client) PutMeRoiDigest(ctx context.Context, body PutMeRoiDigestBody, opts ...RequestOption) (*PutMeRoiDigestResponse, error) {
	var out PutMeRoiDigestResponse
	if err := c.do(ctx, "PUT", "/v1/me/roi-digest", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PutNetworkSettings calls Join or leave the attribution network. PUT /v1/network/settings.
//
// Opts the org in or out of the cross-merchant attribution network: contributing its referral touchpoints (as hashed customer identities and partner accounts, never merchant identity or PII) and, in return, reading the network journey.
//
// **Auth:** an API key with `growth:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PutNetworkSettings(ctx context.Context, body PutNetworkSettingsBody, opts ...RequestOption) (*PutNetworkSettingsResponse, error) {
	var out PutNetworkSettingsResponse
	if err := c.do(ctx, "PUT", "/v1/network/settings", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PutOrgFXFeed calls Turn the FX feed on or off. PUT /v1/org/fx-feed.
//
// Turns the daily ECB reference-rate feed on or off. Turning it on syncs once at once and returns how many rates were inserted. Audited.
//
// **Auth:** an API key with `org:write` (or `full`), or a session with role `owner`, `admin`, `finance`.
func (c *Client) PutOrgFXFeed(ctx context.Context, body PutOrgFxFeedBody, opts ...RequestOption) (*PutOrgFxFeedResponse, error) {
	var out PutOrgFxFeedResponse
	if err := c.do(ctx, "PUT", "/v1/org/fx-feed", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PutOrgHostedPages calls Set the hosted-page consent and chat. PUT /v1/org/hosted-pages.
//
// Validates and saves the adapters. A CMP's `scriptUrl` must be https on that vendor's own hosts; `custom` needs every origin listed. `customOrigins` adds origins per CSP directive (https or wss, no wildcards, no keywords). Owner or admin. Audited.
//
// **Auth:** an API key with `org:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PutOrgHostedPages(ctx context.Context, body PutOrgHostedPagesBody, opts ...RequestOption) (*GetOrgHostedPagesResponse, error) {
	var out GetOrgHostedPagesResponse
	if err := c.do(ctx, "PUT", "/v1/org/hosted-pages", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PutOrgPayoutSchedule calls Set the org payout schedule. PUT /v1/org/payout-schedule.
//
// When payouts run automatically for every program without its own schedule: a `schedule` (`monthly`, `weekly`, `twice_monthly` or `custom` dates), the older `dayOfMonth` form, or `null` to turn automatic payouts off. Send exactly one of `schedule` and `dayOfMonth`. Audited.
//
// **Auth:** an API key with `payouts:write` (or `full`), or a session with role `owner`, `admin`, `finance`.
func (c *Client) PutOrgPayoutSchedule(ctx context.Context, body PutOrgPayoutScheduleBody, opts ...RequestOption) (*PutOrgPayoutScheduleResponse, error) {
	var out PutOrgPayoutScheduleResponse
	if err := c.do(ctx, "PUT", "/v1/org/payout-schedule", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PutOrgPortalAuth calls Set partner sign-in. PUT /v1/org/portal-auth.
//
// Chooses the portal sign-in adapter and its settings. `custom`: `issuer`, `jwksUrl` (https), optional `loginUrl` (your page that mints the token) and `autoEnrollProgramId`. `oidc`: `issuer`, `clientId`, `clientSecret`. `referalli` removes any configuration. Owner only.
//
// **Auth:** an API key with `org:write` (or `full`), or a session with role `owner`.
func (c *Client) PutOrgPortalAuth(ctx context.Context, body json.RawMessage, opts ...RequestOption) (*PutOrgPortalAuthResponse, error) {
	var out PutOrgPortalAuthResponse
	if err := c.do(ctx, "PUT", "/v1/org/portal-auth", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PutOrgReferralParams calls Set referral query parameters. PUT /v1/org/referral-params.
//
// Replaces the extra query parameters aff.js reads a referral code from (up to 10; lowercase letters, digits and `_`). Built-in names are dropped. Imports add the previous tool's parameter here so already-shared links keep crediting. Audited.
//
// **Auth:** an API key with `org:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PutOrgReferralParams(ctx context.Context, body PutOrgReferralParamsBody, opts ...RequestOption) (*GetOrgReferralParamsResponse, error) {
	var out GetOrgReferralParamsResponse
	if err := c.do(ctx, "PUT", "/v1/org/referral-params", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PutOrgRetention calls Set the data retention policy. PUT /v1/org/retention.
//
// Changes retention days for one or more kinds of raw data (others keep their value). Values below a kind's minimum are rejected. Audited.
//
// **Auth:** an API key with `org:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PutOrgRetention(ctx context.Context, body PutOrgRetentionBody, opts ...RequestOption) (*Retention, error) {
	var out Retention
	if err := c.do(ctx, "PUT", "/v1/org/retention", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PutOrgSSODirect calls Connect your own identity provider. PUT /v1/org/sso/direct.
//
// Team SSO with the org's own IdP, no broker: `adapter` `oidc` (issuer, client id and secret, the claim that holds roles) or `workos` (client id, organization id, the org's own API key). `roleMapping` maps IdP groups or roles to `admin`, `finance`, `analyst` or `support` (never owner); `defaultRole` covers everyone else, or null to keep them out; `domains` are the email domains "require SSO" applies to. The secret is encrypted and never returned; leave it out to keep the stored one. 409 while the org uses the Atlas broker. Owner only.
//
// **Auth:** an API key with `org:write` (or `full`), or a session with role `owner`.
func (c *Client) PutOrgSSODirect(ctx context.Context, body json.RawMessage, opts ...RequestOption) (*PutOrgSsoDirectResponse, error) {
	var out PutOrgSsoDirectResponse
	if err := c.do(ctx, "PUT", "/v1/org/sso/direct", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PutOrgTaxFiling calls Connect a 1099 filing provider. PUT /v1/org/tax-filing.
//
// Chooses who e-files 1099-NEC forms: `test` (simulated) or `tax1099` with your own API key (stored encrypted). Audited.
//
// **Auth:** an API key with `org:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PutOrgTaxFiling(ctx context.Context, body PutOrgTaxFilingBody, opts ...RequestOption) (*PutOrgTaxFilingResponse, error) {
	var out PutOrgTaxFilingResponse
	if err := c.do(ctx, "PUT", "/v1/org/tax-filing", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PutOrgTaxProfile calls Update the tax profile. PUT /v1/org/tax-profile.
//
// Sets any of the tax profile fields; `null` clears one. A VAT id must be a valid EU VAT number; an EIN looks like `12-3456789`. Audited.
//
// **Auth:** an API key with `org:write` (or `full`), or a session with role `owner`, `admin`, `finance`.
func (c *Client) PutOrgTaxProfile(ctx context.Context, body PutOrgTaxProfileBody, opts ...RequestOption) (*TaxProfile, error) {
	var out TaxProfile
	if err := c.do(ctx, "PUT", "/v1/org/tax-profile", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PutOrgTrackingConsent calls Set the tracker's consent default. PUT /v1/org/tracking-consent.
//
// Sets the consent default served inside `/t/{token}/aff.js` (cached up to 5 minutes). Owner or admin. Audited.
//
// **Auth:** an API key with `org:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PutOrgTrackingConsent(ctx context.Context, body PutOrgTrackingConsentBody, opts ...RequestOption) (*GetOrgTrackingConsentResponse, error) {
	var out GetOrgTrackingConsentResponse
	if err := c.do(ctx, "PUT", "/v1/org/tracking-consent", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PutPartnerLocale calls Set my email language (partner). PUT /v1/partner/locale.
//
// The partner account's preferred locale for emails.
//
// **Auth:** a partner session.
func (c *Client) PutPartnerLocale(ctx context.Context, body PutPartnerLocaleBody, opts ...RequestOption) (*PutPartnerLocaleResponse, error) {
	var out PutPartnerLocaleResponse
	if err := c.do(ctx, "PUT", "/v1/partner/locale", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PutPartnerOrgsOrgIDPayoutProfile calls Set my payout details. PUT /v1/partner/orgs/{orgId}/payout-profile.
//
// How this merchant pays the partner: the method, its details (checked for what the rail needs, such as a PayPal email or wallet address) and currency. When two-factor sign-in is on, send the current code in `x-totp-code` (`401 step_up_required` otherwise).
//
// **Auth:** a partner session.
func (c *Client) PutPartnerOrgsOrgIDPayoutProfile(ctx context.Context, orgId string, body PutPartnerOrgsOrgIdPayoutProfileBody, opts ...RequestOption) (*PutPartnerOrgsOrgIdPayoutProfileResponse, error) {
	var out PutPartnerOrgsOrgIdPayoutProfileResponse
	if err := c.do(ctx, "PUT", "/v1/partner/orgs/"+url.PathEscape(orgId)+"/payout-profile", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PutPartnerOrgsOrgIDPostback calls Set my conversion postback. PUT /v1/partner/orgs/{orgId}/postback.
//
// Registers a server-to-server postback URL the merchant calls on the partner's events, with macros `{topic}`, `{partner_id}`, `{referral_id}`, `{customer}`, `{amount}`, `{currency}`. Returns a new signing secret.
//
// **Auth:** a partner session.
func (c *Client) PutPartnerOrgsOrgIDPostback(ctx context.Context, orgId string, body PutPartnerOrgsOrgIdPostbackBody, opts ...RequestOption) (*PutPartnerOrgsOrgIdPostbackResponse, error) {
	var out PutPartnerOrgsOrgIdPostbackResponse
	if err := c.do(ctx, "PUT", "/v1/partner/orgs/"+url.PathEscape(orgId)+"/postback", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PutPartnerOrgsOrgIDVAT calls Set my VAT number. PUT /v1/partner/orgs/{orgId}/vat.
//
// Sets the partner's own EU VAT number with this merchant, checked against VIES when the merchant requires it.
//
// **Auth:** a partner session.
func (c *Client) PutPartnerOrgsOrgIDVAT(ctx context.Context, orgId string, body PutPartnerOrgsOrgIdVatBody, opts ...RequestOption) (*PutPartnersIdVatResponse, error) {
	var out PutPartnersIdVatResponse
	if err := c.do(ctx, "PUT", "/v1/partner/orgs/"+url.PathEscape(orgId)+"/vat", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PutPartnerProfile calls Update my creator profile. PUT /v1/partner/profile.
//
// Updates the creator profile merchants search. `published: true` lists it in creator search; `acceptsInvites` controls whether merchants can invite the partner.
//
// **Auth:** a partner session.
func (c *Client) PutPartnerProfile(ctx context.Context, body PutPartnerProfileBody, opts ...RequestOption) (*Ok, error) {
	var out Ok
	if err := c.do(ctx, "PUT", "/v1/partner/profile", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PutPartnersIDPayoutProfile calls Set a partner's payout profile. PUT /v1/partners/{id}/payout-profile.
//
// How the partner is paid: `method` (`paypal`, `wise`, `stripe_connect`, `ach`, `crypto`, `payoneer`, `tremendous` or `csv`), method `details` (a free-form map such as `{ "paypalEmail": … }`), the currency to pay in and the minimum balance before a payout. Partners can also set this themselves in the portal.
//
// **Auth:** an API key with `payouts:write` (or `full`), or a session with role `owner`, `admin`, `finance`.
func (c *Client) PutPartnersIDPayoutProfile(ctx context.Context, id string, body PutPartnersIdPayoutProfileBody, opts ...RequestOption) (*PutPartnersIdPayoutProfileResponse, error) {
	var out PutPartnersIdPayoutProfileResponse
	if err := c.do(ctx, "PUT", "/v1/partners/"+url.PathEscape(id)+"/payout-profile", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PutPartnersIDVAT calls Set a partner's VAT number. PUT /v1/partners/{id}/vat.
//
// Sets (or clears with `null`) the partner's EU VAT number on their payout profile. With `viesCheck` on, VIES must list it as registered. 404 until the partner has payout details. Audited.
//
// **Auth:** an API key with `partners:write` (or `full`), or a session with role `owner`, `admin`, `finance`.
func (c *Client) PutPartnersIDVAT(ctx context.Context, id string, body PutPartnersIdVatBody, opts ...RequestOption) (*PutPartnersIdVatResponse, error) {
	var out PutPartnersIdVatResponse
	if err := c.do(ctx, "PUT", "/v1/partners/"+url.PathEscape(id)+"/vat", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PutPayoutConnectionsProvider calls Connect a payout provider. PUT /v1/payout-connections/{provider}.
//
// Connects (or replaces) your own account with a payout provider, so payouts move money from your account: `paypal` (`clientId`, `secret`), `stripe` (`secretKey`), `circle` (`apiKey`, `walletId`), `wise` (`apiToken`, `profileId`), `payoneer` (`clientId`, `secret`, `programId`) or `tremendous` (`apiKey`, `fundingSourceId`), each with `env`. Stored encrypted and never echoed. Live mode only (`409 test_mode_unsupported` in test mode). Audited.
//
// **Auth:** an API key with `payouts:write` (or `full`), or a session with role `owner`, `admin`, `finance`.
func (c *Client) PutPayoutConnectionsProvider(ctx context.Context, provider string, body json.RawMessage, opts ...RequestOption) (*PutPayoutConnectionsProviderResponse, error) {
	var out PutPayoutConnectionsProviderResponse
	if err := c.do(ctx, "PUT", "/v1/payout-connections/"+url.PathEscape(provider), nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PutProfitabilityProductsProductRef calls Set a product margin. PUT /v1/profitability/products/{productRef}.
//
// Sets a product's margin (`marginBps`) or unit cost (`unitCostMinor`, in `currency`, default the reporting currency); exactly one of the two. Audited.
//
// **Auth:** an API key with `analytics:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PutProfitabilityProductsProductRef(ctx context.Context, productRef string, body PutProfitabilityProductsProductRefBody, opts ...RequestOption) (*GetProfitabilityProductsProductRefResponse, error) {
	var out GetProfitabilityProductsProductRefResponse
	if err := c.do(ctx, "PUT", "/v1/profitability/products/"+url.PathEscape(productRef), nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PutProfitabilitySettings calls Update profitability settings. PUT /v1/profitability/settings.
//
// Changes the org's economics settings, or with `programId` that program's overrides (a program can't override org-only fields such as the reporting currency). Returns the resolved settings. Re-run `POST /v1/profitability/rebuild` to apply changes to past figures. Audited.
//
// **Auth:** an API key with `analytics:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PutProfitabilitySettings(ctx context.Context, body PutProfitabilitySettingsBody, opts ...RequestOption) (*GetProfitabilitySettingsResponse, error) {
	var out GetProfitabilitySettingsResponse
	if err := c.do(ctx, "PUT", "/v1/profitability/settings", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PutProgramsIDAutopricing calls Configure autopricing. PUT /v1/programs/{id}/autopricing.
//
// Self-pricing commissions: a target (`targetRoas`, or with `objective: "roi"` a `targetRoi`) and the `minBps`/`maxBps` band partner rates may move within. `enabled: false` turns it off.
//
// **Auth:** an API key with `programs:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PutProgramsIDAutopricing(ctx context.Context, id string, body PutProgramsIdAutopricingBody, opts ...RequestOption) (*PutProgramsIdAutopricingResponse, error) {
	var out PutProgramsIdAutopricingResponse
	if err := c.do(ctx, "PUT", "/v1/programs/"+url.PathEscape(id)+"/autopricing", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PutProgramsIDIncrementality calls Set the incrementality holdout. PUT /v1/programs/{id}/incrementality.
//
// Holds out a random control share of referrals (`holdoutBps`, up to 5000 = 50%) from being commissioned, so comparing control and treatment conversion measures the true lift partners cause. 0 turns it off.
//
// **Auth:** an API key with `programs:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PutProgramsIDIncrementality(ctx context.Context, id string, body PutProgramsIdIncrementalityBody, opts ...RequestOption) (*PutProgramsIdIncrementalityResponse, error) {
	var out PutProgramsIdIncrementalityResponse
	if err := c.do(ctx, "PUT", "/v1/programs/"+url.PathEscape(id)+"/incrementality", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PutProgramsIDNegotiationPolicy calls Set the agent negotiation policy. PUT /v1/programs/{id}/negotiation-policy.
//
// What an agent may negotiate to in this program: rates between `minBps` and `maxBps`, accepted automatically up to `autoAcceptMaxBps`. Requires `minBps <= autoAcceptMaxBps <= maxBps`. `enabled: false` closes the program to negotiation.
//
// **Auth:** an API key with `programs:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PutProgramsIDNegotiationPolicy(ctx context.Context, id string, body PutProgramsIdNegotiationPolicyBody, opts ...RequestOption) (*PutProgramsIdNegotiationPolicyResponse, error) {
	var out PutProgramsIdNegotiationPolicyResponse
	if err := c.do(ctx, "PUT", "/v1/programs/"+url.PathEscape(id)+"/negotiation-policy", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PutProgramsIDPartnersPartnerIDGroup calls Set a partner's group. PUT /v1/programs/{id}/partners/{partnerId}/group.
//
// Puts an enrolled partner in one of the program's groups, or none (`groupId: null`). A partner is in at most one group per program.
//
// **Auth:** an API key with `programs:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PutProgramsIDPartnersPartnerIDGroup(ctx context.Context, id string, partnerId string, body PutProgramsIdPartnersPartnerIdGroupBody, opts ...RequestOption) (*PutProgramsIdPartnersPartnerIdGroupResponse, error) {
	var out PutProgramsIdPartnersPartnerIdGroupResponse
	if err := c.do(ctx, "PUT", "/v1/programs/"+url.PathEscape(id)+"/partners/"+url.PathEscape(partnerId)+"/group", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PutProgramsIDPayoutSchedule calls Set a program's payout schedule. PUT /v1/programs/{id}/payout-schedule.
//
// Overrides the org payout schedule for this program: a schedule, `{ "kind": "off" }` to never pay it automatically, or `null` to follow the org schedule again. Returns the next run date. Audited.
//
// **Auth:** an API key with `programs:write` (or `full`), or a session with role `owner`, `admin`, `finance`.
func (c *Client) PutProgramsIDPayoutSchedule(ctx context.Context, id string, body PutProgramsIdPayoutScheduleBody, opts ...RequestOption) (*PutProgramsIdPayoutScheduleResponse, error) {
	var out PutProgramsIdPayoutScheduleResponse
	if err := c.do(ctx, "PUT", "/v1/programs/"+url.PathEscape(id)+"/payout-schedule", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PutProgramsIDRefereeReward calls Set the referee reward. PUT /v1/programs/{id}/referee-reward.
//
// What a referred customer gets: `store_credit`, a `coupon` or a `gift_card`, either a flat `amountMinor` or a percent (`bps`, optionally capped with `capMinor`) of their first payment, issued on their first commission or first payment. A percent reward needs `on: "first_payment"`. `enabled: false` removes it. Each reward issued publishes `referee_reward.earned`.
//
// **Auth:** an API key with `programs:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PutProgramsIDRefereeReward(ctx context.Context, id string, body PutProgramsIdRefereeRewardBody, opts ...RequestOption) (*PutProgramsIdRefereeRewardResponse, error) {
	var out PutProgramsIdRefereeRewardResponse
	if err := c.do(ctx, "PUT", "/v1/programs/"+url.PathEscape(id)+"/referee-reward", nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PutRewardsRulesKey calls Save a reward rule. PUT /v1/rewards/rules/{key}.
//
// Creates or replaces a rule: the amount, the reason, an optional program and an optional `event` it rewards (one rule per event).
//
// **Auth:** an API key with `commissions:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PutRewardsRulesKey(ctx context.Context, key string, body PutRewardsRulesKeyBody, opts ...RequestOption) (*PutRewardsRulesKeyResponse, error) {
	var out PutRewardsRulesKeyResponse
	if err := c.do(ctx, "PUT", "/v1/rewards/rules/"+url.PathEscape(key), nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PutSupportAppsAdapter calls Connect a support tool. PUT /v1/support/apps/{adapter}.
//
// Enable a support tool and store its signing secret (vendor-issued; write-only) or, for Zendesk, the app public key in `config.publicKeyPem`. For `custom` Referalli mints the secret: it is returned once, and `rotate: true` mints a new one.
//
// **Auth:** an API key with `partners:write` (or `full`), or a session with role `owner`, `admin`.
func (c *Client) PutSupportAppsAdapter(ctx context.Context, adapter string, body PutSupportAppsAdapterBody, opts ...RequestOption) (*PutSupportAppsAdapterResponse, error) {
	var out PutSupportAppsAdapterResponse
	if err := c.do(ctx, "PUT", "/v1/support/apps/"+url.PathEscape(adapter), nil, body, &out, "json", opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// Call invokes any operation by its id (see Operations) through its typed method:
// path holds path params, query is decoded into the operation's Params, body into its body type.
func (c *Client) Call(ctx context.Context, operationID string, path map[string]string, query map[string]any, body json.RawMessage, opts ...RequestOption) (any, error) {
	switch operationID {
	case "deleteApiKeysId":
		var p DeleteAPIKeysIDParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.DeleteAPIKeysID(ctx, path["id"], &p, opts...)
	case "deleteAssetsId":
		return c.DeleteAssetsID(ctx, path["id"], opts...)
	case "deleteAuthConnectionsId":
		return c.DeleteAuthConnectionsID(ctx, path["id"], opts...)
	case "deleteBillingConnectionsId":
		return c.DeleteBillingConnectionsID(ctx, path["id"], opts...)
	case "deleteBillingInternalCustomersCustomerRef":
		return c.DeleteBillingInternalCustomersCustomerRef(ctx, path["customerRef"], opts...)
	case "deleteBrandingDomain":
		return c.DeleteBrandingDomain(ctx, opts...)
	case "deleteBrandingLinkDomain":
		return c.DeleteBrandingLinkDomain(ctx, opts...)
	case "deleteCliListenersId":
		return c.DeleteCliListenersID(ctx, path["id"], opts...)
	case "deleteEmailTemplatesEventLocale":
		return c.DeleteEmailTemplatesEventLocale(ctx, path["event"], path["locale"], opts...)
	case "deleteIntegrationsKind":
		return c.DeleteIntegrationsKind(ctx, path["kind"], opts...)
	case "deleteOrgMembersUserId":
		return c.DeleteOrgMembersUserID(ctx, path["userId"], opts...)
	case "deleteOrgSsoDirect":
		return c.DeleteOrgSSODirect(ctx, opts...)
	case "deletePartnersId":
		return c.DeletePartnersID(ctx, path["id"], opts...)
	case "deletePayoutConnectionsProvider":
		return c.DeletePayoutConnectionsProvider(ctx, path["provider"], opts...)
	case "deletePrivacyExecutorsAdapter":
		return c.DeletePrivacyExecutorsAdapter(ctx, path["adapter"], opts...)
	case "deleteProfitabilityProductsProductRef":
		return c.DeleteProfitabilityProductsProductRef(ctx, path["productRef"], opts...)
	case "deleteRewardsRulesKey":
		return c.DeleteRewardsRulesKey(ctx, path["key"], opts...)
	case "deleteSupportAppsAdapter":
		return c.DeleteSupportAppsAdapter(ctx, path["adapter"], opts...)
	case "deleteWebhookEndpointsId":
		return c.DeleteWebhookEndpointsID(ctx, path["id"], opts...)
	case "deleteZapierHooksId":
		return c.DeleteZapierHooksID(ctx, path["id"], opts...)
	case "getAdminMetrics":
		return c.GetAdminMetrics(ctx, opts...)
	case "getAdminOrgs":
		var p GetAdminOrgsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetAdminOrgs(ctx, &p, opts...)
	case "getAdminPartnerAccounts":
		var p GetAdminPartnerAccountsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetAdminPartnerAccounts(ctx, &p, opts...)
	case "getAgents":
		var p GetAgentsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetAgents(ctx, &p, opts...)
	case "getAgentsRecommendations":
		var p GetAgentsRecommendationsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetAgentsRecommendations(ctx, &p, opts...)
	case "getAnalyticsByChannel":
		var p GetAnalyticsByChannelParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetAnalyticsByChannel(ctx, &p, opts...)
	case "getAnalyticsByGeo":
		var p GetAnalyticsByGeoParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetAnalyticsByGeo(ctx, &p, opts...)
	case "getAnalyticsCohortLtv":
		var p GetAnalyticsCohortLtvParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetAnalyticsCohortLtv(ctx, &p, opts...)
	case "getAnalyticsCouponPerformance":
		var p GetAnalyticsCouponPerformanceParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetAnalyticsCouponPerformance(ctx, &p, opts...)
	case "getAnalyticsOverview":
		var p GetAnalyticsOverviewParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetAnalyticsOverview(ctx, &p, opts...)
	case "getAnalyticsTimeseries":
		var p GetAnalyticsTimeseriesParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetAnalyticsTimeseries(ctx, &p, opts...)
	case "getApiKeys":
		var p GetAPIKeysParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetAPIKeys(ctx, &p, opts...)
	case "getApiKeysId":
		return c.GetAPIKeysID(ctx, path["id"], opts...)
	case "getAssets":
		var p GetAssetsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetAssets(ctx, &p, opts...)
	case "getAssistantSuggestions":
		var p GetAssistantSuggestionsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetAssistantSuggestions(ctx, &p, opts...)
	case "getAuditLogs":
		var p GetAuditLogsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetAuditLogs(ctx, &p, opts...)
	case "getAuthConnections":
		var p GetAuthConnectionsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetAuthConnections(ctx, &p, opts...)
	case "getAuthConnectionsIdExport":
		var p GetAuthConnectionsIDExportParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetAuthConnectionsIDExport(ctx, path["id"], &p, opts...)
	case "getAuthConnectionsProviders":
		return c.GetAuthConnectionsProviders(ctx, opts...)
	case "getAuthProviders":
		return c.GetAuthProviders(ctx, opts...)
	case "getBillingConnections":
		var p GetBillingConnectionsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetBillingConnections(ctx, &p, opts...)
	case "getBillingConnectionsId":
		return c.GetBillingConnectionsID(ctx, path["id"], opts...)
	case "getBillingEvents":
		var p GetBillingEventsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetBillingEvents(ctx, &p, opts...)
	case "getBillingInternalCustomers":
		return c.GetBillingInternalCustomers(ctx, opts...)
	case "getBillingPlan":
		return c.GetBillingPlan(ctx, opts...)
	case "getBranding":
		return c.GetBranding(ctx, opts...)
	case "getChatlyIdentity":
		return c.GetChatlyIdentity(ctx, opts...)
	case "getCleanroomConsent":
		var p GetCleanroomConsentParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetCleanroomConsent(ctx, &p, opts...)
	case "getCleanroomQuery":
		var p GetCleanroomQueryParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetCleanroomQuery(ctx, &p, opts...)
	case "getClicksConsentCoverage":
		var p GetClicksConsentCoverageParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetClicksConsentCoverage(ctx, &p, opts...)
	case "getClicksId":
		return c.GetClicksID(ctx, path["id"], opts...)
	case "getCliDeviceCodesUserCode":
		return c.GetCliDeviceCodesUserCode(ctx, path["userCode"], opts...)
	case "getCliListenersIdEvents":
		var p GetCliListenersIDEventsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetCliListenersIDEvents(ctx, path["id"], &p, opts...)
	case "getCliLogs":
		var p GetCliLogsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetCliLogs(ctx, &p, opts...)
	case "getCliTopics":
		return c.GetCliTopics(ctx, opts...)
	case "getCommissionRules":
		var p GetCommissionRulesParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetCommissionRules(ctx, &p, opts...)
	case "getCommissions":
		var p GetCommissionsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetCommissions(ctx, &p, opts...)
	case "getCommissionsId":
		return c.GetCommissionsID(ctx, path["id"], opts...)
	case "getContracts":
		var p GetContractsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetContracts(ctx, &p, opts...)
	case "getCoupons":
		var p GetCouponsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetCoupons(ctx, &p, opts...)
	case "getCreators":
		var p GetCreatorsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetCreators(ctx, &p, opts...)
	case "getCreatorsAccountId":
		return c.GetCreatorsAccountID(ctx, path["accountId"], opts...)
	case "getDeals":
		var p GetDealsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetDeals(ctx, &p, opts...)
	case "getDisputes":
		var p GetDisputesParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetDisputes(ctx, &p, opts...)
	case "getEmailSettings":
		return c.GetEmailSettings(ctx, opts...)
	case "getEmailTemplates":
		var p GetEmailTemplatesParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetEmailTemplates(ctx, &p, opts...)
	case "getEmbed":
		var p GetEmbedParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetEmbed(ctx, &p, opts...)
	case "getEmbedSummary":
		return c.GetEmbedSummary(ctx, opts...)
	case "getEvents":
		var p GetEventsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetEvents(ctx, &p, opts...)
	case "getExportAccounting":
		var p GetExportAccountingParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetExportAccounting(ctx, &p, opts...)
	case "getExportFull":
		return c.GetExportFull(ctx, opts...)
	case "getFraudCases":
		var p GetFraudCasesParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetFraudCases(ctx, &p, opts...)
	case "getFxRates":
		var p GetFXRatesParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetFXRates(ctx, &p, opts...)
	case "getGraphql":
		return c.GetGraphql(ctx, opts...)
	case "getHealth":
		return c.GetHealth(ctx, opts...)
	case "getHostsCurrent":
		return c.GetHostsCurrent(ctx, opts...)
	case "getHostsCurrentThemeCss":
		return c.GetHostsCurrentThemeCss(ctx, opts...)
	case "getIncrementality":
		var p GetIncrementalityParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetIncrementality(ctx, &p, opts...)
	case "getIntegrationHealth":
		return c.GetIntegrationHealth(ctx, opts...)
	case "getIntegrations":
		return c.GetIntegrations(ctx, opts...)
	case "getIntegrationsDeliveries":
		var p GetIntegrationsDeliveriesParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetIntegrationsDeliveries(ctx, &p, opts...)
	case "getInvitations":
		var p GetInvitationsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetInvitations(ctx, &p, opts...)
	case "getLinks":
		var p GetLinksParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetLinks(ctx, &p, opts...)
	case "getLinksIdQr":
		return c.GetLinksIDQR(ctx, path["id"], opts...)
	case "getMarketplaceListing":
		var p GetMarketplaceListingParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetMarketplaceListing(ctx, &p, opts...)
	case "getMe":
		return c.GetMe(ctx, opts...)
	case "getNegotiations":
		var p GetNegotiationsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetNegotiations(ctx, &p, opts...)
	case "getNetworkJourney":
		var p GetNetworkJourneyParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetNetworkJourney(ctx, &p, opts...)
	case "getNetworkSettings":
		return c.GetNetworkSettings(ctx, opts...)
	case "getNotifications":
		var p GetNotificationsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetNotifications(ctx, &p, opts...)
	case "getOnchainConnections":
		var p GetOnchainConnectionsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetOnchainConnections(ctx, &p, opts...)
	case "getOpenapiJson":
		return c.GetOpenapiJSON(ctx, opts...)
	case "getOpenapiYaml":
		return c.GetOpenapiYaml(ctx, opts...)
	case "getOrg":
		return c.GetOrg(ctx, opts...)
	case "getOrgFxFeed":
		return c.GetOrgFXFeed(ctx, opts...)
	case "getOrgHostedPages":
		return c.GetOrgHostedPages(ctx, opts...)
	case "getOrgMembers":
		var p GetOrgMembersParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetOrgMembers(ctx, &p, opts...)
	case "getOrgPayoutSchedule":
		return c.GetOrgPayoutSchedule(ctx, opts...)
	case "getOrgPortalAuth":
		return c.GetOrgPortalAuth(ctx, opts...)
	case "getOrgReferralParams":
		return c.GetOrgReferralParams(ctx, opts...)
	case "getOrgRetention":
		return c.GetOrgRetention(ctx, opts...)
	case "getOrgSso":
		return c.GetOrgSSO(ctx, opts...)
	case "getOrgSsoConnectionsIdSamlMetadata":
		return c.GetOrgSSOConnectionsIDSamlMetadata(ctx, path["id"], opts...)
	case "getOrgTaxProfile":
		return c.GetOrgTaxProfile(ctx, opts...)
	case "getOrgTrackingConsent":
		return c.GetOrgTrackingConsent(ctx, opts...)
	case "getOutreach":
		var p GetOutreachParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetOutreach(ctx, &p, opts...)
	case "getPartnerAchievements":
		var p GetPartnerAchievementsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetPartnerAchievements(ctx, &p, opts...)
	case "getPartnerAdvance":
		var p GetPartnerAdvanceParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetPartnerAdvance(ctx, &p, opts...)
	case "getPartnerAdvanceOffer":
		return c.GetPartnerAdvanceOffer(ctx, opts...)
	case "getPartnerAssets":
		var p GetPartnerAssetsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetPartnerAssets(ctx, &p, opts...)
	case "getPartnerContracts":
		var p GetPartnerContractsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetPartnerContracts(ctx, &p, opts...)
	case "getPartnerDeals":
		var p GetPartnerDealsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetPartnerDeals(ctx, &p, opts...)
	case "getPartnerDisputes":
		var p GetPartnerDisputesParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetPartnerDisputes(ctx, &p, opts...)
	case "getPartnerGroups":
		var p GetPartnerGroupsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetPartnerGroups(ctx, &p, opts...)
	case "getPartnerInvitations":
		var p GetPartnerInvitationsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetPartnerInvitations(ctx, &p, opts...)
	case "getPartnerKycStatus":
		return c.GetPartnerKYCStatus(ctx, opts...)
	case "getPartnerLinks":
		var p GetPartnerLinksParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetPartnerLinks(ctx, &p, opts...)
	case "getPartnerLinksIdQr":
		return c.GetPartnerLinksIDQR(ctx, path["id"], opts...)
	case "getPartnerMarketplace":
		var p GetPartnerMarketplaceParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetPartnerMarketplace(ctx, &p, opts...)
	case "getPartnerMe":
		return c.GetPartnerMe(ctx, opts...)
	case "getPartnerNotifications":
		var p GetPartnerNotificationsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetPartnerNotifications(ctx, &p, opts...)
	case "getPartnerOrgsOrgIdSelfBilling":
		return c.GetPartnerOrgsOrgIDSelfBilling(ctx, path["orgId"], opts...)
	case "getPartnerPassport":
		return c.GetPartnerPassport(ctx, opts...)
	case "getPartnerPayouts":
		var p GetPartnerPayoutsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetPartnerPayouts(ctx, &p, opts...)
	case "getPartnerPayoutsIdInvoice":
		var p GetPartnerPayoutsIDInvoiceParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetPartnerPayoutsIDInvoice(ctx, path["id"], &p, opts...)
	case "getPartnerPostbackLogs":
		var p GetPartnerPostbackLogsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetPartnerPostbackLogs(ctx, &p, opts...)
	case "getPartnerProfile":
		return c.GetPartnerProfile(ctx, opts...)
	case "getPartnerProfileVerified":
		return c.GetPartnerProfileVerified(ctx, opts...)
	case "getPartnerProgramsProgramIdLeaderboard":
		return c.GetPartnerProgramsProgramIDLeaderboard(ctx, path["programId"], opts...)
	case "getPartnerRecruit":
		return c.GetPartnerRecruit(ctx, opts...)
	case "getPartnerReferrals":
		var p GetPartnerReferralsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetPartnerReferrals(ctx, &p, opts...)
	case "getPartnerReferralsIdTransparency":
		return c.GetPartnerReferralsIDTransparency(ctx, path["id"], opts...)
	case "getPartners":
		var p GetPartnersParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetPartners(ctx, &p, opts...)
	case "getPartnerSecurity":
		return c.GetPartnerSecurity(ctx, opts...)
	case "getPartnersId":
		return c.GetPartnersID(ctx, path["id"], opts...)
	case "getPartnersIdBalances":
		var p GetPartnersIDBalancesParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetPartnersIDBalances(ctx, path["id"], &p, opts...)
	case "getPartnersIdEnrollments":
		var p GetPartnersIDEnrollmentsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetPartnersIDEnrollments(ctx, path["id"], &p, opts...)
	case "getPartnersIdPayoutProfile":
		return c.GetPartnersIDPayoutProfile(ctx, path["id"], opts...)
	case "getPartnersIdQuality":
		return c.GetPartnersIDQuality(ctx, path["id"], opts...)
	case "getPartnerStats":
		var p GetPartnerStatsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetPartnerStats(ctx, &p, opts...)
	case "getPartnerTaxDocuments":
		var p GetPartnerTaxDocumentsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetPartnerTaxDocuments(ctx, &p, opts...)
	case "getPartnerTaxDocumentsId":
		return c.GetPartnerTaxDocumentsID(ctx, path["id"], opts...)
	case "getPartnerTerms":
		var p GetPartnerTermsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetPartnerTerms(ctx, &p, opts...)
	case "getPassportPublicKey":
		return c.GetPassportPublicKey(ctx, opts...)
	case "getPayoutConnections":
		var p GetPayoutConnectionsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetPayoutConnections(ctx, &p, opts...)
	case "getPayoutRuns":
		var p GetPayoutRunsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetPayoutRuns(ctx, &p, opts...)
	case "getPayouts":
		var p GetPayoutsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetPayouts(ctx, &p, opts...)
	case "getPayoutsIdInvoice":
		var p GetPayoutsIDInvoiceParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetPayoutsIDInvoice(ctx, path["id"], &p, opts...)
	case "getPlans":
		return c.GetPlans(ctx, opts...)
	case "getPortalAuthCurrent":
		var p GetPortalAuthCurrentParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetPortalAuthCurrent(ctx, &p, opts...)
	case "getPrivacyExecutorRequests":
		return c.GetPrivacyExecutorRequests(ctx, opts...)
	case "getPrivacyExecutors":
		return c.GetPrivacyExecutors(ctx, opts...)
	case "getPrivacyExport":
		var p GetPrivacyExportParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetPrivacyExport(ctx, &p, opts...)
	case "getPrivacyRequests":
		var p GetPrivacyRequestsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetPrivacyRequests(ctx, &p, opts...)
	case "getPrivacyRequestsId":
		return c.GetPrivacyRequestsID(ctx, path["id"], opts...)
	case "getPrivacyRequestsIdResult":
		return c.GetPrivacyRequestsIDResult(ctx, path["id"], opts...)
	case "getProfitabilityProducts":
		return c.GetProfitabilityProducts(ctx, opts...)
	case "getProfitabilityProductsProductRef":
		return c.GetProfitabilityProductsProductRef(ctx, path["productRef"], opts...)
	case "getProfitabilityRebuildJobId":
		return c.GetProfitabilityRebuildJobID(ctx, path["jobId"], opts...)
	case "getProfitabilitySettings":
		var p GetProfitabilitySettingsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetProfitabilitySettings(ctx, &p, opts...)
	case "getPrograms":
		var p GetProgramsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetPrograms(ctx, &p, opts...)
	case "getProgramsId":
		return c.GetProgramsID(ctx, path["id"], opts...)
	case "getProgramsIdAutopricing":
		return c.GetProgramsIDAutopricing(ctx, path["id"], opts...)
	case "getProgramsIdEnrollments":
		var p GetProgramsIDEnrollmentsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetProgramsIDEnrollments(ctx, path["id"], &p, opts...)
	case "getProgramsIdGraph":
		var p GetProgramsIDGraphParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetProgramsIDGraph(ctx, path["id"], &p, opts...)
	case "getProgramsIdGroups":
		var p GetProgramsIDGroupsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetProgramsIDGroups(ctx, path["id"], &p, opts...)
	case "getProgramsIdNegotiationPolicy":
		return c.GetProgramsIDNegotiationPolicy(ctx, path["id"], opts...)
	case "getProgramsIdPromotions":
		var p GetProgramsIDPromotionsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetProgramsIDPromotions(ctx, path["id"], &p, opts...)
	case "getProgramsIdRefereeReward":
		return c.GetProgramsIDRefereeReward(ctx, path["id"], opts...)
	case "getProgramsIdTermsAcceptances":
		var p GetProgramsIDTermsAcceptancesParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetProgramsIDTermsAcceptances(ctx, path["id"], &p, opts...)
	case "getProspects":
		var p GetProspectsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetProspects(ctx, &p, opts...)
	case "getPublicProgramsHandle":
		return c.GetPublicProgramsHandle(ctx, path["handle"], opts...)
	case "getReferrals":
		var p GetReferralsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetReferrals(ctx, &p, opts...)
	case "getReferralsId":
		return c.GetReferralsID(ctx, path["id"], opts...)
	case "getReferralsIdJourney":
		return c.GetReferralsIDJourney(ctx, path["id"], opts...)
	case "getReportsFilteredClicks":
		var p GetReportsFilteredClicksParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetReportsFilteredClicks(ctx, &p, opts...)
	case "getReportsPartnerPerformance":
		var p GetReportsPartnerPerformanceParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetReportsPartnerPerformance(ctx, &p, opts...)
	case "getRewards":
		var p GetRewardsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetRewards(ctx, &p, opts...)
	case "getRewardsGranted":
		return c.GetRewardsGranted(ctx, opts...)
	case "getRewardsLookupCode":
		return c.GetRewardsLookupCode(ctx, path["code"], opts...)
	case "getRewardsRules":
		return c.GetRewardsRules(ctx, opts...)
	case "getRoiBreakdown":
		var p GetRoiBreakdownParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetRoiBreakdown(ctx, &p, opts...)
	case "getRoiDataQuality":
		var p GetRoiDataQualityParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetRoiDataQuality(ctx, &p, opts...)
	case "getRoiGrants":
		var p GetRoiGrantsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetRoiGrants(ctx, &p, opts...)
	case "getRoiPartnersIdTrace":
		var p GetRoiPartnersIDTraceParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetRoiPartnersIDTrace(ctx, path["id"], &p, opts...)
	case "getRoiPatterns":
		return c.GetRoiPatterns(ctx, opts...)
	case "getRoiPatternsPattern":
		var p GetRoiPatternsPatternParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetRoiPatternsPattern(ctx, path["pattern"], &p, opts...)
	case "getRoiSummary":
		var p GetRoiSummaryParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetRoiSummary(ctx, &p, opts...)
	case "getRoiTraceCustomerRef":
		var p GetRoiTraceCustomerRefParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetRoiTraceCustomerRef(ctx, path["customerRef"], &p, opts...)
	case "getSetupStatus":
		return c.GetSetupStatus(ctx, opts...)
	case "getStatus":
		return c.GetStatus(ctx, opts...)
	case "getStatusIncidents":
		var p GetStatusIncidentsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetStatusIncidents(ctx, &p, opts...)
	case "getStatusIncidentsAtom":
		return c.GetStatusIncidentsAtom(ctx, opts...)
	case "getStatusIncidentsId":
		return c.GetStatusIncidentsID(ctx, path["id"], opts...)
	case "getSupportApps":
		return c.GetSupportApps(ctx, opts...)
	case "getSupportCardsOrgIdAdapter":
		var p GetSupportCardsOrgIDAdapterParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetSupportCardsOrgIDAdapter(ctx, path["orgId"], path["adapter"], &p, opts...)
	case "getSupportContext":
		var p GetSupportContextParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetSupportContext(ctx, &p, opts...)
	case "getTax1099":
		var p GetTax1099Params
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetTax1099(ctx, &p, opts...)
	case "getTax1099Filings":
		var p GetTax1099FilingsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetTax1099Filings(ctx, &p, opts...)
	case "getTax1099Preview":
		var p GetTax1099PreviewParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetTax1099Preview(ctx, &p, opts...)
	case "getTestClock":
		return c.GetTestClock(ctx, opts...)
	case "getTestEmails":
		var p GetTestEmailsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetTestEmails(ctx, &p, opts...)
	case "getValueImportJobId":
		return c.GetValueImportJobID(ctx, path["jobId"], opts...)
	case "getVersions":
		return c.GetVersions(ctx, opts...)
	case "getWebhookDeliveries":
		var p GetWebhookDeliveriesParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetWebhookDeliveries(ctx, &p, opts...)
	case "getWebhookDeliveriesId":
		return c.GetWebhookDeliveriesID(ctx, path["id"], opts...)
	case "getWebhookEndpoints":
		var p GetWebhookEndpointsParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetWebhookEndpoints(ctx, &p, opts...)
	case "getWebhookEndpointsId":
		return c.GetWebhookEndpointsID(ctx, path["id"], opts...)
	case "getWebhookEndpointsIdDeliveries":
		var p GetWebhookEndpointsIDDeliveriesParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetWebhookEndpointsIDDeliveries(ctx, path["id"], &p, opts...)
	case "getZapierEvents":
		return c.GetZapierEvents(ctx, opts...)
	case "getZapierHooks":
		var p GetZapierHooksParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetZapierHooks(ctx, &p, opts...)
	case "getZapierTriggersEvent":
		var p GetZapierTriggersEventParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.GetZapierTriggersEvent(ctx, path["event"], &p, opts...)
	case "patchAdminStatusIncidentsId":
		var b PatchAdminStatusIncidentsIdBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PatchAdminStatusIncidentsID(ctx, path["id"], b, opts...)
	case "patchApiKeysId":
		var b PatchApiKeysIdBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PatchAPIKeysID(ctx, path["id"], b, opts...)
	case "patchAuthConnectionsId":
		var b PatchAuthConnectionsIdBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PatchAuthConnectionsID(ctx, path["id"], b, opts...)
	case "patchBillingConnectionsId":
		var b PatchBillingConnectionsIdBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PatchBillingConnectionsID(ctx, path["id"], b, opts...)
	case "patchOrg":
		var b PatchOrgBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PatchOrg(ctx, b, opts...)
	case "patchOrgMembersUserId":
		var b PatchOrgMembersUserIdBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PatchOrgMembersUserID(ctx, path["userId"], b, opts...)
	case "patchOrgSso":
		var b PatchOrgSsoBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PatchOrgSSO(ctx, b, opts...)
	case "patchPartnerGroupsId":
		var b PatchPartnerGroupsIdBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PatchPartnerGroupsID(ctx, path["id"], b, opts...)
	case "patchPartnersId":
		var b PatchPartnersIdBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PatchPartnersID(ctx, path["id"], b, opts...)
	case "patchPartnersIdPayoutProfile":
		var b PatchPartnersIdPayoutProfileBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PatchPartnersIDPayoutProfile(ctx, path["id"], b, opts...)
	case "patchProgramsId":
		var b PatchProgramsIdBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PatchProgramsID(ctx, path["id"], b, opts...)
	case "patchReferralsId":
		var b PatchReferralsIdBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PatchReferralsID(ctx, path["id"], b, opts...)
	case "patchWebhookEndpointsId":
		var b PatchWebhookEndpointsIdBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PatchWebhookEndpointsID(ctx, path["id"], b, opts...)
	case "postAccountClose":
		var b PostAccountCloseBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostAccountClose(ctx, b, opts...)
	case "postAdminOrgsIdPlan":
		var b PostAdminOrgsIdPlanBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostAdminOrgsIDPlan(ctx, path["id"], b, opts...)
	case "postAdminOrgsIdSuspend":
		var b PostAdminOrgsIdSuspendBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostAdminOrgsIDSuspend(ctx, path["id"], b, opts...)
	case "postAdminOrgsIdUnsuspend":
		return c.PostAdminOrgsIDUnsuspend(ctx, path["id"], opts...)
	case "postAdminPartnerAccountsIdBan":
		return c.PostAdminPartnerAccountsIDBan(ctx, path["id"], opts...)
	case "postAdminPartnerAccountsIdUnban":
		return c.PostAdminPartnerAccountsIDUnban(ctx, path["id"], opts...)
	case "postAdminStatusIncidents":
		var b PostAdminStatusIncidentsBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostAdminStatusIncidents(ctx, b, opts...)
	case "postAdminStatusIncidentsIdUpdates":
		var b PostAdminStatusIncidentsIdUpdatesBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostAdminStatusIncidentsIDUpdates(ctx, path["id"], b, opts...)
	case "postAgentsRegister":
		var b PostAgentsRegisterBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostAgentsRegister(ctx, b, opts...)
	case "postAnomaliesScan":
		return c.PostAnomaliesScan(ctx, opts...)
	case "postApiKeys":
		var b PostApiKeysBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostAPIKeys(ctx, b, opts...)
	case "postAssets":
		var b PostAssetsBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostAssets(ctx, b, opts...)
	case "postAttributionReplay":
		var b PostAttributionReplayBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostAttributionReplay(ctx, b, opts...)
	case "postAuthBackchannelLogout":
		var b PostAuthBackchannelLogoutBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostAuthBackchannelLogout(ctx, b, opts...)
	case "postAuthConnections":
		var b PostAuthConnectionsBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostAuthConnections(ctx, b, opts...)
	case "postAuthConnectionsIdAliases":
		var b PostAuthConnectionsIdAliasesBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostAuthConnectionsIDAliases(ctx, path["id"], b, opts...)
	case "postAuthConnectionsIdErase":
		var b PostAuthConnectionsIdEraseBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostAuthConnectionsIDErase(ctx, path["id"], b, opts...)
	case "postAuthConnectionsIdPartners":
		var b PostAuthConnectionsIdPartnersBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostAuthConnectionsIDPartners(ctx, path["id"], b, opts...)
	case "postAuthConnectionsIdPing":
		var b map[string]any
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostAuthConnectionsIDPing(ctx, path["id"], b, opts...)
	case "postAuthConnectionsIdRevokePreviousSecret":
		return c.PostAuthConnectionsIDRevokePreviousSecret(ctx, path["id"], opts...)
	case "postAuthConnectionsIdRotateSecret":
		return c.PostAuthConnectionsIDRotateSecret(ctx, path["id"], opts...)
	case "postAuthConnectionsIdSignals":
		var b PostAuthConnectionsIdSignalsBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostAuthConnectionsIDSignals(ctx, path["id"], b, opts...)
	case "postAuthConnectionsIdSignups":
		var b PostAuthConnectionsIdSignupsBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostAuthConnectionsIDSignups(ctx, path["id"], b, opts...)
	case "postAuthConnectionsIdUsers":
		var b PostAuthConnectionsIdUsersBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostAuthConnectionsIDUsers(ctx, path["id"], b, opts...)
	case "postAuthConnectionsInstall":
		var b PostAuthConnectionsInstallBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostAuthConnectionsInstall(ctx, b, opts...)
	case "postAuthLogout":
		return c.PostAuthLogout(ctx, opts...)
	case "postBillingActivate":
		var b PostBillingActivateBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostBillingActivate(ctx, b, opts...)
	case "postBillingCheckout":
		var b PostBillingCheckoutBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostBillingCheckout(ctx, b, opts...)
	case "postBillingConnections":
		var b PostBillingConnectionsBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostBillingConnections(ctx, b, opts...)
	case "postBillingConnectionsIdRotateSecret":
		return c.PostBillingConnectionsIDRotateSecret(ctx, path["id"], opts...)
	case "postBillingConnectionsIdValueMappingPreview":
		var b PostBillingConnectionsIdValueMappingPreviewBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostBillingConnectionsIDValueMappingPreview(ctx, path["id"], b, opts...)
	case "postBillingInternalCustomers":
		var b PostBillingInternalCustomersBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostBillingInternalCustomers(ctx, b, opts...)
	case "postBillingPortal":
		return c.PostBillingPortal(ctx, opts...)
	case "postBrandingDomain":
		var b PostBrandingDomainBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostBrandingDomain(ctx, b, opts...)
	case "postBrandingDomainVerify":
		return c.PostBrandingDomainVerify(ctx, opts...)
	case "postBrandingLinkDomain":
		var b PostBrandingLinkDomainBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostBrandingLinkDomain(ctx, b, opts...)
	case "postBrandingLinkDomainVerify":
		return c.PostBrandingLinkDomainVerify(ctx, opts...)
	case "postCliDeviceCodes":
		var b PostCliDeviceCodesBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostCliDeviceCodes(ctx, b, opts...)
	case "postCliDeviceCodesToken":
		var b PostCliDeviceCodesTokenBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostCliDeviceCodesToken(ctx, b, opts...)
	case "postCliDeviceCodesUserCodeApprove":
		var b PostCliDeviceCodesUserCodeApproveBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostCliDeviceCodesUserCodeApprove(ctx, path["userCode"], b, opts...)
	case "postCliDeviceCodesUserCodeDeny":
		return c.PostCliDeviceCodesUserCodeDeny(ctx, path["userCode"], opts...)
	case "postCliListeners":
		var b PostCliListenersBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostCliListeners(ctx, b, opts...)
	case "postCliTrigger":
		var b PostCliTriggerBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostCliTrigger(ctx, b, opts...)
	case "postCommissionRules":
		var b PostCommissionRulesBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostCommissionRules(ctx, b, opts...)
	case "postCommissionsApproveDue":
		return c.PostCommissionsApproveDue(ctx, opts...)
	case "postCommissionsIdAdjust":
		var b PostCommissionsIdAdjustBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostCommissionsIDAdjust(ctx, path["id"], b, opts...)
	case "postCommissionsIdSettle":
		var b PostCommissionsIdSettleBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostCommissionsIDSettle(ctx, path["id"], b, opts...)
	case "postContracts":
		var b PostContractsBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostContracts(ctx, b, opts...)
	case "postContractsIdSend":
		return c.PostContractsIDSend(ctx, path["id"], opts...)
	case "postCoupons":
		var b PostCouponsBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostCoupons(ctx, b, opts...)
	case "postCreatorsAccountIdInvite":
		var b PostCreatorsAccountIdInviteBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostCreatorsAccountIDInvite(ctx, path["accountId"], b, opts...)
	case "postDealsIdApprove":
		var b PostDealsIdApproveBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostDealsIDApprove(ctx, path["id"], b, opts...)
	case "postDealsIdReject":
		return c.PostDealsIDReject(ctx, path["id"], opts...)
	case "postDealsIdWon":
		var b PostDealsIdWonBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostDealsIDWon(ctx, path["id"], b, opts...)
	case "postDiscoveryCorpus":
		var b PostDiscoveryCorpusBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostDiscoveryCorpus(ctx, b, opts...)
	case "postDiscoverySearch":
		var b PostDiscoverySearchBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostDiscoverySearch(ctx, b, opts...)
	case "postDisputesIdResolve":
		var b PostDisputesIdResolveBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostDisputesIDResolve(ctx, path["id"], b, opts...)
	case "postEmailPreview":
		var b PostEmailPreviewBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostEmailPreview(ctx, b, opts...)
	case "postEmailSettingsVerifyDomain":
		return c.PostEmailSettingsVerifyDomain(ctx, opts...)
	case "postEmailTest":
		var b PostEmailTestBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostEmailTest(ctx, b, opts...)
	case "postEmbedTokens":
		var b PostEmbedTokensBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostEmbedTokens(ctx, b, opts...)
	case "postFraudCasesIdResolve":
		var b PostFraudCasesIdResolveBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostFraudCasesIDResolve(ctx, path["id"], b, opts...)
	case "postFraudDetectRings":
		return c.PostFraudDetectRings(ctx, opts...)
	case "postFraudInvalidTrafficScan":
		var b PostFraudInvalidTrafficScanBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostFraudInvalidTrafficScan(ctx, b, opts...)
	case "postFxRates":
		var b PostFxRatesBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostFXRates(ctx, b, opts...)
	case "postFxRatesClearOverride":
		var b PostFxRatesClearOverrideBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostFXRatesClearOverride(ctx, b, opts...)
	case "postFxRatesSync":
		return c.PostFXRatesSync(ctx, opts...)
	case "postGraphql":
		var b PostGraphqlBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostGraphql(ctx, b, opts...)
	case "postImportProvider":
		var b PostImportProviderBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		var p PostImportProviderParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.PostImportProvider(ctx, path["provider"], b, &p, opts...)
	case "postIntegrationsDeliveriesIdRetry":
		return c.PostIntegrationsDeliveriesIDRetry(ctx, path["id"], opts...)
	case "postIntegrationsHubspotTest":
		var b PostIntegrationsHubspotTestBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostIntegrationsHubspotTest(ctx, b, opts...)
	case "postIntegrationsSegmentTest":
		return c.PostIntegrationsSegmentTest(ctx, opts...)
	case "postInvitationsIdWithdraw":
		return c.PostInvitationsIDWithdraw(ctx, path["id"], opts...)
	case "postLinks":
		var b PostLinksBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostLinks(ctx, b, opts...)
	case "postMobilePurchases":
		var b PostMobilePurchasesBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostMobilePurchases(ctx, b, opts...)
	case "postNotificationsIdRead":
		return c.PostNotificationsIDRead(ctx, path["id"], opts...)
	case "postOauthToken":
		var b PostOauthTokenBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostOauthToken(ctx, b, opts...)
	case "postOnchainConnections":
		var b PostOnchainConnectionsBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostOnchainConnections(ctx, b, opts...)
	case "postOrgHostedPagesCheck":
		return c.PostOrgHostedPagesCheck(ctx, opts...)
	case "postOrgMembers":
		var b PostOrgMembersBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostOrgMembers(ctx, b, opts...)
	case "postOrgPortalAuthTest":
		return c.PostOrgPortalAuthTest(ctx, opts...)
	case "postOrgs":
		var b PostOrgsBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostOrgs(ctx, b, opts...)
	case "postOrgSsoDirectTest":
		return c.PostOrgSSODirectTest(ctx, opts...)
	case "postOrgSsoLink":
		return c.PostOrgSSOLink(ctx, opts...)
	case "postOrgSsoTickets":
		var b PostOrgSsoTicketsBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostOrgSSOTickets(ctx, b, opts...)
	case "postOrgSsoTicketsIdRevoke":
		return c.PostOrgSSOTicketsIDRevoke(ctx, path["id"], opts...)
	case "postOutreachDraft":
		var b PostOutreachDraftBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostOutreachDraft(ctx, b, opts...)
	case "postOutreachIdApprove":
		return c.PostOutreachIDApprove(ctx, path["id"], opts...)
	case "postOutreachSendDue":
		return c.PostOutreachSendDue(ctx, opts...)
	case "postPartnerAdvanceAccept":
		var b PostPartnerAdvanceAcceptBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostPartnerAdvanceAccept(ctx, b, opts...)
	case "postPartnerAuthWalletNonce":
		var b PostPartnerAuthWalletNonceBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostPartnerAuthWalletNonce(ctx, b, opts...)
	case "postPartnerAuthWalletVerify":
		var b PostPartnerAuthWalletVerifyBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostPartnerAuthWalletVerify(ctx, b, opts...)
	case "postPartnerContractsIdAccept":
		return c.PostPartnerContractsIDAccept(ctx, path["id"], opts...)
	case "postPartnerContractsIdDecline":
		return c.PostPartnerContractsIDDecline(ctx, path["id"], opts...)
	case "postPartnerDeals":
		var b PostPartnerDealsBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostPartnerDeals(ctx, b, opts...)
	case "postPartnerDisputes":
		var b PostPartnerDisputesBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostPartnerDisputes(ctx, b, opts...)
	case "postPartnerInvitationsIdAccept":
		return c.PostPartnerInvitationsIDAccept(ctx, path["id"], opts...)
	case "postPartnerInvitationsIdDecline":
		return c.PostPartnerInvitationsIDDecline(ctx, path["id"], opts...)
	case "postPartnerKycSubmit":
		var b PostPartnerKycSubmitBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostPartnerKYCSubmit(ctx, b, opts...)
	case "postPartnerLinks":
		var b PostPartnerLinksBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostPartnerLinks(ctx, b, opts...)
	case "postPartnerNotificationsIdRead":
		return c.PostPartnerNotificationsIDRead(ctx, path["id"], opts...)
	case "postPartnerOrgsOrgIdSelfBilling":
		var b PostPartnerOrgsOrgIdSelfBillingBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostPartnerOrgsOrgIDSelfBilling(ctx, path["orgId"], b, opts...)
	case "postPartnerOrgsOrgIdTaxForm":
		var b PostPartnerOrgsOrgIdTaxFormBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostPartnerOrgsOrgIDTaxForm(ctx, path["orgId"], b, opts...)
	case "postPartnerProgramsProgramIdApply":
		return c.PostPartnerProgramsProgramIDApply(ctx, path["programId"], opts...)
	case "postPartners":
		var b PostPartnersBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostPartners(ctx, b, opts...)
	case "postPartnerSecurityTotpDisable":
		var b PostPartnerSecurityTotpDisableBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostPartnerSecurityTotpDisable(ctx, b, opts...)
	case "postPartnerSecurityTotpEnable":
		var b PostPartnerSecurityTotpEnableBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostPartnerSecurityTotpEnable(ctx, b, opts...)
	case "postPartnerSecurityTotpSetup":
		return c.PostPartnerSecurityTotpSetup(ctx, opts...)
	case "postPartnersIdTaxForm":
		var b PostPartnersIdTaxFormBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostPartnersIDTaxForm(ctx, path["id"], b, opts...)
	case "postPartnerTermsVersionIdAccept":
		return c.PostPartnerTermsVersionIDAccept(ctx, path["versionId"], opts...)
	case "postPassportVerify":
		var b PostPassportVerifyBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostPassportVerify(ctx, b, opts...)
	case "postPayoutRuns":
		return c.PostPayoutRuns(ctx, opts...)
	case "postPayoutRunsIdCancel":
		return c.PostPayoutRunsIDCancel(ctx, path["id"], opts...)
	case "postPayoutRunsIdExecute":
		return c.PostPayoutRunsIDExecute(ctx, path["id"], opts...)
	case "postPayoutRunsRunScheduled":
		return c.PostPayoutRunsRunScheduled(ctx, opts...)
	case "postPrivacyExecutorKeys":
		var b PostPrivacyExecutorKeysBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostPrivacyExecutorKeys(ctx, b, opts...)
	case "postPrivacyExecutorsAdapterCheck":
		return c.PostPrivacyExecutorsAdapterCheck(ctx, path["adapter"], opts...)
	case "postPrivacyForget":
		var b PostPrivacyForgetBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostPrivacyForget(ctx, b, opts...)
	case "postPrivacyLookup":
		var b PostPrivacyLookupBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostPrivacyLookup(ctx, b, opts...)
	case "postPrivacyRequests":
		var b PostPrivacyRequestsBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostPrivacyRequests(ctx, b, opts...)
	case "postProfitabilityRebuild":
		return c.PostProfitabilityRebuild(ctx, opts...)
	case "postPrograms":
		var b PostProgramsBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostPrograms(ctx, b, opts...)
	case "postProgramsIdApplyLeveling":
		return c.PostProgramsIDApplyLeveling(ctx, path["id"], opts...)
	case "postProgramsIdAutopriceRun":
		var b PostProgramsIdAutopriceRunBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostProgramsIDAutopriceRun(ctx, path["id"], b, opts...)
	case "postProgramsIdEnrollments":
		var b PostProgramsIdEnrollmentsBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostProgramsIDEnrollments(ctx, path["id"], b, opts...)
	case "postProgramsIdEnrollmentsEnrollmentIdApprove":
		return c.PostProgramsIDEnrollmentsEnrollmentIDApprove(ctx, path["id"], path["enrollmentId"], opts...)
	case "postProgramsIdEnrollmentsEnrollmentIdReject":
		return c.PostProgramsIDEnrollmentsEnrollmentIDReject(ctx, path["id"], path["enrollmentId"], opts...)
	case "postProgramsIdGroups":
		var b PostProgramsIdGroupsBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostProgramsIDGroups(ctx, path["id"], b, opts...)
	case "postProgramsIdPromotions":
		var b PostProgramsIdPromotionsBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostProgramsIDPromotions(ctx, path["id"], b, opts...)
	case "postPromotionsIdEnd":
		return c.PostPromotionsIDEnd(ctx, path["id"], opts...)
	case "postPublicProgramsHandleJoin":
		var b PostPublicProgramsHandleJoinBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostPublicProgramsHandleJoin(ctx, path["handle"], b, opts...)
	case "postQualityRecompute":
		return c.PostQualityRecompute(ctx, opts...)
	case "postReferralsAssociate":
		var b PostReferralsAssociateBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostReferralsAssociate(ctx, b, opts...)
	case "postReferralsIdApprove":
		return c.PostReferralsIDApprove(ctx, path["id"], opts...)
	case "postReferralsIdReassign":
		var b PostReferralsIdReassignBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostReferralsIDReassign(ctx, path["id"], b, opts...)
	case "postReferralsIdVoid":
		return c.PostReferralsIDVoid(ctx, path["id"], opts...)
	case "postRewards":
		var b PostRewardsBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostRewards(ctx, b, opts...)
	case "postRewardsIdRedeem":
		var b PostRewardsIdRedeemBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostRewardsIDRedeem(ctx, path["id"], b, opts...)
	case "postRewardsRedeemByCode":
		var b PostRewardsRedeemByCodeBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostRewardsRedeemByCode(ctx, b, opts...)
	case "postRoiPatternsPatternReview":
		var b PostRoiPatternsPatternReviewBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostRoiPatternsPatternReview(ctx, path["pattern"], b, opts...)
	case "postSandboxSimulate":
		var b PostSandboxSimulateBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostSandboxSimulate(ctx, b, opts...)
	case "postSupportAppsAdapterTest":
		var b PostSupportAppsAdapterTestBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostSupportAppsAdapterTest(ctx, path["adapter"], b, opts...)
	case "postSupportCardsOrgIdAdapter":
		var b map[string]any
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostSupportCardsOrgIDAdapter(ctx, path["orgId"], path["adapter"], b, opts...)
	case "postTax1099DocumentsIdCorrect":
		return c.PostTax1099DocumentsIDCorrect(ctx, path["id"], opts...)
	case "postTax1099Filings":
		var b PostTax1099FilingsBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostTax1099Filings(ctx, b, opts...)
	case "postTax1099FilingsIdRefresh":
		return c.PostTax1099FilingsIDRefresh(ctx, path["id"], opts...)
	case "postTestClockAdvance":
		var b PostTestClockAdvanceBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostTestClockAdvance(ctx, b, opts...)
	case "postTrackEvents":
		var b json.RawMessage
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostTrackEvents(ctx, b, opts...)
	case "postUnsubscribeToken":
		return c.PostUnsubscribeToken(ctx, path["token"], opts...)
	case "postValueEvents":
		var b PostValueEventsBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostValueEvents(ctx, b, opts...)
	case "postValueImport":
		var b PostValueImportBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostValueImport(ctx, b, opts...)
	case "postWebhookDeliveriesIdReplay":
		return c.PostWebhookDeliveriesIDReplay(ctx, path["id"], opts...)
	case "postWebhookEndpoints":
		var b PostWebhookEndpointsBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostWebhookEndpoints(ctx, b, opts...)
	case "postWebhookEndpointsIdReplay":
		var p PostWebhookEndpointsIDReplayParams
		if err := decodeQuery(query, &p); err != nil {
			return nil, err
		}
		return c.PostWebhookEndpointsIDReplay(ctx, path["id"], &p, opts...)
	case "postWebhookEndpointsIdRotateSecret":
		return c.PostWebhookEndpointsIDRotateSecret(ctx, path["id"], opts...)
	case "postWidgetTokens":
		var b PostWidgetTokensBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostWidgetTokens(ctx, b, opts...)
	case "postZapierConversions":
		var b PostZapierConversionsBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostZapierConversions(ctx, b, opts...)
	case "postZapierHooks":
		var b PostZapierHooksBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PostZapierHooks(ctx, b, opts...)
	case "putAuthConnectionsIdSecret":
		var b PutAuthConnectionsIdSecretBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PutAuthConnectionsIDSecret(ctx, path["id"], b, opts...)
	case "putBillingConnectionsIdCredentials":
		var b PutBillingConnectionsIdCredentialsBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PutBillingConnectionsIDCredentials(ctx, path["id"], b, opts...)
	case "putBillingConnectionsIdSecret":
		var b PutBillingConnectionsIdSecretBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PutBillingConnectionsIDSecret(ctx, path["id"], b, opts...)
	case "putBillingConnectionsIdValueMapping":
		var b PutBillingConnectionsIdValueMappingBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PutBillingConnectionsIDValueMapping(ctx, path["id"], b, opts...)
	case "putBranding":
		var b PutBrandingBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PutBranding(ctx, b, opts...)
	case "putEmailSettings":
		var b PutEmailSettingsBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PutEmailSettings(ctx, b, opts...)
	case "putEmailTemplatesEventLocale":
		var b PutEmailTemplatesEventLocaleBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PutEmailTemplatesEventLocale(ctx, path["event"], path["locale"], b, opts...)
	case "putIntegrationsHubspot":
		var b PutIntegrationsHubspotBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PutIntegrationsHubspot(ctx, b, opts...)
	case "putIntegrationsSegment":
		var b PutIntegrationsSegmentBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PutIntegrationsSegment(ctx, b, opts...)
	case "putMarketplaceListing":
		var b PutMarketplaceListingBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PutMarketplaceListing(ctx, b, opts...)
	case "putMeLocale":
		var b PutMeLocaleBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PutMeLocale(ctx, b, opts...)
	case "putMeRoiDigest":
		var b PutMeRoiDigestBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PutMeRoiDigest(ctx, b, opts...)
	case "putNetworkSettings":
		var b PutNetworkSettingsBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PutNetworkSettings(ctx, b, opts...)
	case "putOrgFxFeed":
		var b PutOrgFxFeedBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PutOrgFXFeed(ctx, b, opts...)
	case "putOrgHostedPages":
		var b PutOrgHostedPagesBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PutOrgHostedPages(ctx, b, opts...)
	case "putOrgPayoutSchedule":
		var b PutOrgPayoutScheduleBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PutOrgPayoutSchedule(ctx, b, opts...)
	case "putOrgPortalAuth":
		var b json.RawMessage
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PutOrgPortalAuth(ctx, b, opts...)
	case "putOrgReferralParams":
		var b PutOrgReferralParamsBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PutOrgReferralParams(ctx, b, opts...)
	case "putOrgRetention":
		var b PutOrgRetentionBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PutOrgRetention(ctx, b, opts...)
	case "putOrgSsoDirect":
		var b json.RawMessage
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PutOrgSSODirect(ctx, b, opts...)
	case "putOrgTaxFiling":
		var b PutOrgTaxFilingBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PutOrgTaxFiling(ctx, b, opts...)
	case "putOrgTaxProfile":
		var b PutOrgTaxProfileBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PutOrgTaxProfile(ctx, b, opts...)
	case "putOrgTrackingConsent":
		var b PutOrgTrackingConsentBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PutOrgTrackingConsent(ctx, b, opts...)
	case "putPartnerLocale":
		var b PutPartnerLocaleBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PutPartnerLocale(ctx, b, opts...)
	case "putPartnerOrgsOrgIdPayoutProfile":
		var b PutPartnerOrgsOrgIdPayoutProfileBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PutPartnerOrgsOrgIDPayoutProfile(ctx, path["orgId"], b, opts...)
	case "putPartnerOrgsOrgIdPostback":
		var b PutPartnerOrgsOrgIdPostbackBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PutPartnerOrgsOrgIDPostback(ctx, path["orgId"], b, opts...)
	case "putPartnerOrgsOrgIdVat":
		var b PutPartnerOrgsOrgIdVatBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PutPartnerOrgsOrgIDVAT(ctx, path["orgId"], b, opts...)
	case "putPartnerProfile":
		var b PutPartnerProfileBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PutPartnerProfile(ctx, b, opts...)
	case "putPartnersIdPayoutProfile":
		var b PutPartnersIdPayoutProfileBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PutPartnersIDPayoutProfile(ctx, path["id"], b, opts...)
	case "putPartnersIdVat":
		var b PutPartnersIdVatBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PutPartnersIDVAT(ctx, path["id"], b, opts...)
	case "putPayoutConnectionsProvider":
		var b json.RawMessage
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PutPayoutConnectionsProvider(ctx, path["provider"], b, opts...)
	case "putProfitabilityProductsProductRef":
		var b PutProfitabilityProductsProductRefBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PutProfitabilityProductsProductRef(ctx, path["productRef"], b, opts...)
	case "putProfitabilitySettings":
		var b PutProfitabilitySettingsBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PutProfitabilitySettings(ctx, b, opts...)
	case "putProgramsIdAutopricing":
		var b PutProgramsIdAutopricingBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PutProgramsIDAutopricing(ctx, path["id"], b, opts...)
	case "putProgramsIdIncrementality":
		var b PutProgramsIdIncrementalityBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PutProgramsIDIncrementality(ctx, path["id"], b, opts...)
	case "putProgramsIdNegotiationPolicy":
		var b PutProgramsIdNegotiationPolicyBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PutProgramsIDNegotiationPolicy(ctx, path["id"], b, opts...)
	case "putProgramsIdPartnersPartnerIdGroup":
		var b PutProgramsIdPartnersPartnerIdGroupBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PutProgramsIDPartnersPartnerIDGroup(ctx, path["id"], path["partnerId"], b, opts...)
	case "putProgramsIdPayoutSchedule":
		var b PutProgramsIdPayoutScheduleBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PutProgramsIDPayoutSchedule(ctx, path["id"], b, opts...)
	case "putProgramsIdRefereeReward":
		var b PutProgramsIdRefereeRewardBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PutProgramsIDRefereeReward(ctx, path["id"], b, opts...)
	case "putRewardsRulesKey":
		var b PutRewardsRulesKeyBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PutRewardsRulesKey(ctx, path["key"], b, opts...)
	case "putSupportAppsAdapter":
		var b PutSupportAppsAdapterBody
		if err := decodeInto(body, &b); err != nil {
			return nil, err
		}
		return c.PutSupportAppsAdapter(ctx, path["adapter"], b, opts...)
	}
	return nil, fmt.Errorf("referalli: unknown operation %q", operationID)
}
