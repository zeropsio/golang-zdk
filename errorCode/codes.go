// Generated ZEROPS sdk

package errorCode

type ErrorCode string

const (
	InternalServerError                                      ErrorCode = "internalServerError"
	NotFound                                                 ErrorCode = "notFound"
	InvalidUserInput                                         ErrorCode = "invalidUserInput"
	InvalidUserInputWithText                                 ErrorCode = "invalidUserInputWithText"
	NotAuthorized                                            ErrorCode = "notAuthorized"
	InsufficientPrivileges                                   ErrorCode = "insufficientPrivileges"
	RefreshTokenInvalid                                      ErrorCode = "refreshTokenInvalid"
	EndpointNotFound                                         ErrorCode = "endpointNotFound"
	PotentiallyFraudBehaviorDetected                         ErrorCode = "potentiallyFraudBehaviorDetected"
	ReCaptchaResponseInvalid                                 ErrorCode = "reCaptchaResponseInvalid"
	DownloadTokenInvalidOrExpired                            ErrorCode = "downloadTokenInvalidOrExpired"
	IncorrectUserCredentials                                 ErrorCode = "incorrectUserCredentials"
	NoRightsToSelectedAction                                 ErrorCode = "noRightsToSelectedAction"
	FileSizeLimitExceeded                                    ErrorCode = "fileSizeLimitExceeded"
	AdditionalVerificationFailed                             ErrorCode = "additionalVerificationFailed"
	AtLeastOneManager                                        ErrorCode = "atLeastOneManager"
	ActiveUserMustBeConnectedToAtLeastOneClient              ErrorCode = "activeUserMustBeConnectedToAtLeastOneClient"
	InvalidCredentials                                       ErrorCode = "invalidCredentials"
	PasswordTooShort                                         ErrorCode = "passwordTooShort"
	PasswordTooLong                                          ErrorCode = "passwordTooLong"
	NoActiveConnectionToAnyClient                            ErrorCode = "noActiveConnectionToAnyClient"
	NoActiveConnectionToCurrentClient                        ErrorCode = "noActiveConnectionToCurrentClient"
	TokenExpired                                             ErrorCode = "tokenExpired"
	ClientNotFound                                           ErrorCode = "clientNotFound"
	NoAccessToZerops                                         ErrorCode = "noAccessToZerops"
	UserNotFound                                             ErrorCode = "userNotFound"
	InvalidDocumentRoot                                      ErrorCode = "invalidDocumentRoot"
	InvalidEntrypoint                                        ErrorCode = "invalidEntrypoint"
	InvalidInitCommands                                      ErrorCode = "invalidInitCommands"
	DownloadExpired                                          ErrorCode = "downloadExpired"
	MultipleAccountsConnected                                ErrorCode = "multipleAccountsConnected"
	UserDataNotFound                                         ErrorCode = "userDataNotFound"
	UserDataKeyInvalid                                       ErrorCode = "userDataKeyInvalid"
	UserDataContentInvalid                                   ErrorCode = "userDataContentInvalid"
	UserDataUseOfSystemKey                                   ErrorCode = "userDataUseOfSystemKey"
	UserDataDeleteForbidden                                  ErrorCode = "userDataDeleteForbidden"
	UserDataDuplicateKey                                     ErrorCode = "userDataDuplicateKey"
	UserDataSyncRunning                                      ErrorCode = "userDataSyncRunning"
	UserDataVersionMismatch                                  ErrorCode = "userDataVersionMismatch"
	UserTokenNotFound                                        ErrorCode = "userTokenNotFound"
	UserAccountNotFound                                      ErrorCode = "userAccountNotFound"
	TemplateNotFound                                         ErrorCode = "templateNotFound"
	CountryCallingCodeDoesNotMatchPhoneNumber                ErrorCode = "countryCallingCodeDoesNotMatchPhoneNumber"
	InvalidPhoneNumber                                       ErrorCode = "invalidPhoneNumber"
	InvalidUserRole                                          ErrorCode = "invalidUserRole"
	InvalidClientUserGroupCombination                        ErrorCode = "invalidClientUserGroupCombination"
	EmailExists                                              ErrorCode = "emailExists"
	UserAlreadyConnectedToCompany                            ErrorCode = "userAlreadyConnectedToCompany"
	CountryNotFound                                          ErrorCode = "countryNotFound"
	RecordNotFound                                           ErrorCode = "recordNotFound"
	MaximumNumberOfClientExceeded                            ErrorCode = "maximumNumberOfClientExceeded"
	UserIsNotInBeingDeletedStatus                            ErrorCode = "userIsNotInBeingDeletedStatus"
	UnableToDeleteYourOwnAccount                             ErrorCode = "unableToDeleteYourOwnAccount"
	ClientUserConnectionNotFound                             ErrorCode = "clientUserConnectionNotFound"
	ClientUserConnectionHasBeenDeleted                       ErrorCode = "clientUserConnectionHasBeenDeleted"
	BadOrExpiredActivationLink                               ErrorCode = "badOrExpiredActivationLink"
	UnableToEditDeletedUser                                  ErrorCode = "unableToEditDeletedUser"
	ClientUserAlreadyBeenDeleted                             ErrorCode = "clientUserAlreadyBeenDeleted"
	LanguageNotFound                                         ErrorCode = "languageNotFound"
	LanguageNotSupported                                     ErrorCode = "languageNotSupported"
	DuplicatedUser                                           ErrorCode = "duplicatedUser"
	RecordIsReadOnly                                         ErrorCode = "recordIsReadOnly"
	RecordIsDeleted                                          ErrorCode = "recordIsDeleted"
	RecordCommonKeyEdit                                      ErrorCode = "recordCommonKeyEdit"
	ServiceStackNotFound                                     ErrorCode = "serviceStackNotFound"
	ServiceStackTypeNotFound                                 ErrorCode = "serviceStackTypeNotFound"
	ServiceStackTypeVersionNotFound                          ErrorCode = "serviceStackTypeVersionNotFound"
	ServiceStackTypeVersionIsNotActive                       ErrorCode = "serviceStackTypeVersionIsNotActive"
	ServiceStackTypeVersionRemoveDefault                     ErrorCode = "serviceStackTypeVersionRemoveDefault"
	ServiceStackTypeInvalidVersion                           ErrorCode = "serviceStackTypeInvalidVersion"
	ServiceStackIsFailed                                     ErrorCode = "serviceStackIsFailed"
	ServiceStackIsSystemNotSupported                         ErrorCode = "serviceStackIsSystemNotSupported"
	ServiceStackCustomPortsEnabledRequired                   ErrorCode = "serviceStackCustomPortsEnabledRequired"
	ServiceStackTypeVersionNotMatchWithStackType             ErrorCode = "serviceStackTypeVersionNotMatchWithStackType"
	ServiceStackStorageDiskMaxLimit                          ErrorCode = "serviceStackStorageDiskMaxLimit"
	ServiceStackStorageDiskMinLimit                          ErrorCode = "serviceStackStorageDiskMinLimit"
	ServiceStackObjectStorageQuotaNotFound                   ErrorCode = "serviceStackObjectStorageQuotaNotFound"
	ServiceStackInitCommandsNotFound                         ErrorCode = "serviceStackInitCommandsNotFound"
	ServiceStackTypeObjectStorageRequired                    ErrorCode = "serviceStackTypeObjectStorageRequired"
	ServiceStackTypeObjectStorageForbidden                   ErrorCode = "serviceStackTypeObjectStorageForbidden"
	ServiceStackTypeSharedStorageForbidden                   ErrorCode = "serviceStackTypeSharedStorageForbidden"
	ServiceStackTypeStandardForbidden                        ErrorCode = "serviceStackTypeStandardForbidden"
	ServiceStackIsReadyToDeploy                              ErrorCode = "serviceStackIsReadyToDeploy"
	ServiceStackIsNotHttp                                    ErrorCode = "serviceStackIsNotHttp"
	ServiceStackReloadNotAvailable                           ErrorCode = "serviceStackReloadNotAvailable"
	ServiceStackNginxConfig                                  ErrorCode = "serviceStackNginxConfig"
	ServiceStackSubdomainAccessAlreadyDisabled               ErrorCode = "serviceStackSubdomainAccessAlreadyDisabled"
	ServiceStackCreating                                     ErrorCode = "serviceStackCreating"
	ServiceStackUpgrading                                    ErrorCode = "serviceStackUpgrading"
	ServiceStackInvalidCategory                              ErrorCode = "serviceStackInvalidCategory"
	ServiceStackStatusNotSupported                           ErrorCode = "serviceStackStatusNotSupported"
	ServiceStackUserNameInvalid                              ErrorCode = "serviceStackUserNameInvalid"
	ActiveServiceStackRequired                               ErrorCode = "activeServiceStackRequired"
	UserApplicationRuntimeServiceStackRequired               ErrorCode = "userApplicationRuntimeServiceStackRequired"
	SharedStorageNonemptyConnectedStacksForbidden            ErrorCode = "sharedStorageNonemptyConnectedStacksForbidden"
	SharedStorageServiceStackInactive                        ErrorCode = "sharedStorageServiceStackInactive"
	ProcessNotFound                                          ErrorCode = "processNotFound"
	ProcessIsAlreadyOver                                     ErrorCode = "processIsAlreadyOver"
	ProcessIsAlreadyRunning                                  ErrorCode = "processIsAlreadyRunning"
	UnableToCreateProcess                                    ErrorCode = "unableToCreateProcess"
	CreateProcessQueuedProjectProcessCountExceeded           ErrorCode = "createProcessQueuedProjectProcessCountExceeded"
	ManualRepairRequired                                     ErrorCode = "manualRepairRequired"
	NothingToSync                                            ErrorCode = "nothingToSync"
	IsNotPossibleToEditAnotherUser                           ErrorCode = "isNotPossibleToEditAnotherUser"
	ActiveServiceStackExistsInProject                        ErrorCode = "activeServiceStackExistsInProject"
	InvalidOldPassword                                       ErrorCode = "invalidOldPassword"
	ProjectStatusNotSupported                                ErrorCode = "projectStatusNotSupported"
	ProjectNotFound                                          ErrorCode = "projectNotFound"
	ProjectIsNotActive                                       ErrorCode = "projectIsNotActive"
	ServiceStackNameUnavailable                              ErrorCode = "serviceStackNameUnavailable"
	ServiceStackNameInvalid                                  ErrorCode = "serviceStackNameInvalid"
	CoreServiceStackTypeVersionNotFound                      ErrorCode = "coreServiceStackTypeVersionNotFound"
	L7HttpBalancerServiceStackTypeVersionNotFound            ErrorCode = "l7HttpBalancerServiceStackTypeVersionNotFound"
	ProjectWillBeDeleted                                     ErrorCode = "projectWillBeDeleted"
	ProjectWillBeStarted                                     ErrorCode = "projectWillBeStarted"
	ProjectWillBeStopped                                     ErrorCode = "projectWillBeStopped"
	StackWillBeDeleted                                       ErrorCode = "stackWillBeDeleted"
	StackOfProjectWillBeDeleted                              ErrorCode = "stackOfProjectWillBeDeleted"
	StackOfProjectWillBeCreated                              ErrorCode = "stackOfProjectWillBeCreated"
	UserDataOfProjectWillBeSynced                            ErrorCode = "userDataOfProjectWillBeSynced"
	PublicHttpRoutingNotFound                                ErrorCode = "publicHttpRoutingNotFound"
	PublicHttpRoutingAtLeastOneDomainRequired                ErrorCode = "publicHttpRoutingAtLeastOneDomainRequired"
	PublicHttpRoutingDuplicateDomainName                     ErrorCode = "publicHttpRoutingDuplicateDomainName"
	PublicHttpRoutingDuplicateFallback                       ErrorCode = "publicHttpRoutingDuplicateFallback"
	PublicHttpRoutingDuplicateDomainNameInAnotherRouting     ErrorCode = "publicHttpRoutingDuplicateDomainNameInAnotherRouting"
	PublicHttpRoutingDuplicateFallbackInAnotherRouting       ErrorCode = "publicHttpRoutingDuplicateFallbackInAnotherRouting"
	PublicHttpRoutingDuplicateLocationPaths                  ErrorCode = "publicHttpRoutingDuplicateLocationPaths"
	PublicHttpRoutingItemsLimit                              ErrorCode = "publicHttpRoutingItemsLimit"
	PublicHttpRoutingLocationsLimit                          ErrorCode = "publicHttpRoutingLocationsLimit"
	PublicHttpRoutingDomainsLimit                            ErrorCode = "publicHttpRoutingDomainsLimit"
	PublicHttpRoutingInvalidDomainNameFQDN                   ErrorCode = "publicHttpRoutingInvalidDomainNameFQDN"
	PublicHttpRoutingSslFallback                             ErrorCode = "publicHttpRoutingSslFallback"
	PublicHttpRoutingLocationDoesNotFound                    ErrorCode = "publicHttpRoutingLocationDoesNotFound"
	PublicHttpRoutingInvalidLocationsPrefix                  ErrorCode = "publicHttpRoutingInvalidLocationsPrefix"
	PublicHttpRoutingMultipleServiceStackAccessPoints        ErrorCode = "publicHttpRoutingMultipleServiceStackAccessPoints"
	PublicHttpRoutingServiceStackAccessPointNotFound         ErrorCode = "publicHttpRoutingServiceStackAccessPointNotFound"
	PublicHttpRoutingRequestOnlyTcpPort                      ErrorCode = "publicHttpRoutingRequestOnlyTcpPort"
	PublicHttpRoutingServiceStackDoesNotSupportHttp          ErrorCode = "publicHttpRoutingServiceStackDoesNotSupportHttp"
	PublicHttpRoutingForbiddenZeropsSubdomainSuffix          ErrorCode = "publicHttpRoutingForbiddenZeropsSubdomainSuffix"
	NoPublicHttpRoutingInL7HttpBalancing                     ErrorCode = "noPublicHttpRoutingInL7HttpBalancing"
	SslNotSupported                                          ErrorCode = "sslNotSupported"
	RegisteredDomainHostLimitReached                         ErrorCode = "registeredDomainHostLimitReached"
	PublicSuffixListControlFailed                            ErrorCode = "publicSuffixListControlFailed"
	PortNotSupported                                         ErrorCode = "portNotSupported"
	PublicHttpRoutingSyncRunning                             ErrorCode = "publicHttpRoutingSyncRunning"
	PublicHttpRoutingVersionMismatch                         ErrorCode = "publicHttpRoutingVersionMismatch"
	PortOutOfRange                                           ErrorCode = "portOutOfRange"
	PortProtocolInvalid                                      ErrorCode = "portProtocolInvalid"
	PortSchemeInvalid                                        ErrorCode = "portSchemeInvalid"
	UdpProtocolSchemeMismatch                                ErrorCode = "udpProtocolSchemeMismatch"
	PublicIpRequestNotFound                                  ErrorCode = "publicIpRequestNotFound"
	PublicIpRequestOnlyOne                                   ErrorCode = "publicIpRequestOnlyOne"
	PublicIpRequestReleaseRunning                            ErrorCode = "publicIpRequestReleaseRunning"
	UserNotificationNotFound                                 ErrorCode = "userNotificationNotFound"
	PublicPortRoutingNotFound                                ErrorCode = "publicPortRoutingNotFound"
	PublicIpTypeNotSupported                                 ErrorCode = "publicIpTypeNotSupported"
	PortAndProtocolCombinationNotSupported                   ErrorCode = "portAndProtocolCombinationNotSupported"
	PublicPortRoutingDuplicateRule                           ErrorCode = "publicPortRoutingDuplicateRule"
	InternalPortDuplicateRule                                ErrorCode = "internalPortDuplicateRule"
	PublicPortRoutingSyncRunning                             ErrorCode = "publicPortRoutingSyncRunning"
	PublicPortRoutingInvalidIpFormat                         ErrorCode = "publicPortRoutingInvalidIpFormat"
	PublicPortRoutingFirewallPolicyConflict                  ErrorCode = "publicPortRoutingFirewallPolicyConflict"
	PublicPortRoutingVersionMismatch                         ErrorCode = "publicPortRoutingVersionMismatch"
	PublicPortRoutingUnableToEditDeletedRecord               ErrorCode = "publicPortRoutingUnableToEditDeletedRecord"
	PublicPortRoutingUnableToEditPortsOnSystemService        ErrorCode = "publicPortRoutingUnableToEditPortsOnSystemService"
	DriverNotFound                                           ErrorCode = "driverNotFound"
	DriverEntityIsNotServiceStack                            ErrorCode = "driverEntityIsNotServiceStack"
	DriverAlreadySet                                         ErrorCode = "driverAlreadySet"
	ClientUserConnectionAlreadyExists                        ErrorCode = "clientUserConnectionAlreadyExists"
	UserAlreadyAssigned                                      ErrorCode = "userAlreadyAssigned"
	ServiceStackTypeUserStorageRequired                      ErrorCode = "serviceStackTypeUserStorageRequired"
	AppVersionNotFound                                       ErrorCode = "appVersionNotFound"
	AppVersionInvalidStatus                                  ErrorCode = "appVersionInvalidStatus"
	DeployProcessAlreadyCreated                              ErrorCode = "deployProcessAlreadyCreated"
	ObjectContentLengthZero                                  ErrorCode = "objectContentLengthZero"
	ObjectMimeTypeNotSupported                               ErrorCode = "objectMimeTypeNotSupported"
	ObjectNotFound                                           ErrorCode = "objectNotFound"
	MaxObjectDiskSizeEstimateExceeded                        ErrorCode = "maxObjectDiskSizeEstimateExceeded"
	AppVersionIsActive                                       ErrorCode = "appVersionIsActive"
	AppVersionNoActiveFound                                  ErrorCode = "appVersionNoActiveFound"
	AppVersionIsBeingDeployed                                ErrorCode = "appVersionIsBeingDeployed"
	AppVersionIsBeingBuilt                                   ErrorCode = "appVersionIsBeingBuilt"
	AppVersionIsNotBeingUploaded                             ErrorCode = "appVersionIsNotBeingUploaded"
	AppVersionDeployingFailed                                ErrorCode = "appVersionDeployingFailed"
	AppVersionBuildingFailed                                 ErrorCode = "appVersionBuildingFailed"
	AppVersionYamlParseFailed                                ErrorCode = "appVersionYamlParseFailed"
	AppVersionYamlConfigRequired                             ErrorCode = "appVersionYamlConfigRequired"
	AppVersionDataNotAvailable                               ErrorCode = "appVersionDataNotAvailable"
	AppVersionActiveHasDefaultRuntime                        ErrorCode = "appVersionActiveHasDefaultRuntime"
	SharedStorageNotFound                                    ErrorCode = "sharedStorageNotFound"
	SharedStorageInvalidProject                              ErrorCode = "sharedStorageInvalidProject"
	ServiceStackConnectionNotFound                           ErrorCode = "serviceStackConnectionNotFound"
	SharedStorageInvalidCategory                             ErrorCode = "sharedStorageInvalidCategory"
	UserServiceStackInvalidCategory                          ErrorCode = "userServiceStackInvalidCategory"
	IdenticalSharedStorageAndServiceStack                    ErrorCode = "identicalSharedStorageAndServiceStack"
	ClientIdProjectServiceStackMissing                       ErrorCode = "clientIdProjectServiceStackMissing"
	DifferentProjectIds                                      ErrorCode = "differentProjectIds"
	StatsHistorySearchLimitInvalid                           ErrorCode = "statsHistorySearchLimitInvalid"
	StatsHistorySearchTimeGroupByInvalid                     ErrorCode = "statsHistorySearchTimeGroupByInvalid"
	StatsHistorySearchIdFilterInvalid                        ErrorCode = "statsHistorySearchIdFilterInvalid"
	StatsHistorySearchProjectIdRequired                      ErrorCode = "statsHistorySearchProjectIdRequired"
	StatsHistorySearchSubscriptionNotSupported               ErrorCode = "statsHistorySearchSubscriptionNotSupported"
	StatsHistorySearchHistoryLimitReached                    ErrorCode = "statsHistorySearchHistoryLimitReached"
	StatsHistorySearchRangeInvalid                           ErrorCode = "statsHistorySearchRangeInvalid"
	StatsHistorySearchEitherRangeOrLimitRequired             ErrorCode = "statsHistorySearchEitherRangeOrLimitRequired"
	StatsHistorySearchFullRangeRequired                      ErrorCode = "statsHistorySearchFullRangeRequired"
	TransactionDebitSearchLimitInvalid                       ErrorCode = "transactionDebitSearchLimitInvalid"
	TransactionDebitSearchTimeGroupByInvalid                 ErrorCode = "transactionDebitSearchTimeGroupByInvalid"
	TransactionDebitSearchRangeInvalid                       ErrorCode = "transactionDebitSearchRangeInvalid"
	TransactionDebitSearchClientIdRequired                   ErrorCode = "transactionDebitSearchClientIdRequired"
	TransactionDebitSearchIdFilterInvalid                    ErrorCode = "transactionDebitSearchIdFilterInvalid"
	TransactionDebitSearchEitherRangeOrLimitRequired         ErrorCode = "transactionDebitSearchEitherRangeOrLimitRequired"
	TransactionDebitSearchFullRangeRequired                  ErrorCode = "transactionDebitSearchFullRangeRequired"
	TransactionDebitSearchTillAndFromDateTimezonesDiffer     ErrorCode = "transactionDebitSearchTillAndFromDateTimezonesDiffer"
	TransactionDebitSearchUnknownTimezone                    ErrorCode = "transactionDebitSearchUnknownTimezone"
	TransactionDebitSearchInvalidFilter                      ErrorCode = "transactionDebitSearchInvalidFilter"
	GithubVerificationExpired                                ErrorCode = "githubVerificationExpired"
	InvalidGithubTokenAction                                 ErrorCode = "invalidGithubTokenAction"
	GithubRequestFailed                                      ErrorCode = "githubRequestFailed"
	GithubAuthorizationRequired                              ErrorCode = "githubAuthorizationRequired"
	GithubEmailAccessRequired                                ErrorCode = "githubEmailAccessRequired"
	GithubNoVerifiedEmailFound                               ErrorCode = "githubNoVerifiedEmailFound"
	GithubRepositoryAccessRequired                           ErrorCode = "githubRepositoryAccessRequired"
	GithubRateLimitExceeded                                  ErrorCode = "githubRateLimitExceeded"
	ServiceStackNoGithubIntegration                          ErrorCode = "serviceStackNoGithubIntegration"
	BranchNameRequired                                       ErrorCode = "branchNameRequired"
	TriggerBuildRequiresBranchEventType                      ErrorCode = "triggerBuildRequiresBranchEventType"
	UserServiceStackRequired                                 ErrorCode = "userServiceStackRequired"
	NoExternalRepositoryIntegration                          ErrorCode = "noExternalRepositoryIntegration"
	GithubBranchDeleted                                      ErrorCode = "githubBranchDeleted"
	GithubWebhookDeleted                                     ErrorCode = "githubWebhookDeleted"
	GithubWebhookInvalidParameters                           ErrorCode = "githubWebhookInvalidParameters"
	UnexpectedGithubFileEncoding                             ErrorCode = "unexpectedGithubFileEncoding"
	ZeropsBuildFileNotFound                                  ErrorCode = "zeropsBuildFileNotFound"
	IncorrectWebhookSecret                                   ErrorCode = "incorrectWebhookSecret"
	UnsupportedWebhookEventType                              ErrorCode = "unsupportedWebhookEventType"
	RepositoryDoesNotMatchServiceStackSettings               ErrorCode = "repositoryDoesNotMatchServiceStackSettings"
	ProjectStatusNotInNewCreatingActive                      ErrorCode = "projectStatusNotInNewCreatingActive"
	GithubPushIgnored                                        ErrorCode = "githubPushIgnored"
	ZeropsBuildFileTooLarge                                  ErrorCode = "zeropsBuildFileTooLarge"
	SameServiceStackTypeUsed                                 ErrorCode = "sameServiceStackTypeUsed"
	GithubWrongRepositoryFullNameFormat                      ErrorCode = "githubWrongRepositoryFullNameFormat"
	GithubWebhooksLimitExceeded                              ErrorCode = "githubWebhooksLimitExceeded"
	AvatarUploadForbiddenImageType                           ErrorCode = "avatarUploadForbiddenImageType"
	AvatarUploadMaxSize                                      ErrorCode = "avatarUploadMaxSize"
	AvatarUploadFileNotFound                                 ErrorCode = "avatarUploadFileNotFound"
	GitlabRateLimitExceeded                                  ErrorCode = "gitlabRateLimitExceeded"
	GitlabVerificationExpired                                ErrorCode = "gitlabVerificationExpired"
	InvalidGitlabTokenAction                                 ErrorCode = "invalidGitlabTokenAction"
	GitlabRequestFailed                                      ErrorCode = "gitlabRequestFailed"
	GitlabAuthorizationRequired                              ErrorCode = "gitlabAuthorizationRequired"
	GitlabFullApiAccessRequired                              ErrorCode = "gitlabFullApiAccessRequired"
	MultipleIntegrationsAreForbidden                         ErrorCode = "multipleIntegrationsAreForbidden"
	GitlabBranchDeleted                                      ErrorCode = "gitlabBranchDeleted"
	GitlabWebhookDeleted                                     ErrorCode = "gitlabWebhookDeleted"
	GitlabWebhookInvalidParameters                           ErrorCode = "gitlabWebhookInvalidParameters"
	ServiceStackNoGitlabIntegration                          ErrorCode = "serviceStackNoGitlabIntegration"
	GitlabPushIgnored                                        ErrorCode = "gitlabPushIgnored"
	YamlFieldLengthError                                     ErrorCode = "yamlFieldLengthError"
	YamlWrongUseFieldType                                    ErrorCode = "yamlWrongUseFieldType"
	VatNumberNotInEUCountry                                  ErrorCode = "vatNumberNotInEUCountry"
	CompanyNotFoundInEU                                      ErrorCode = "companyNotFoundInEU"
	VatNumberMissing                                         ErrorCode = "vatNumberMissing"
	PaymentInProgress                                        ErrorCode = "paymentInProgress"
	DuplicateVatNumber                                       ErrorCode = "duplicateVatNumber"
	DuplicateCompanyNumber                                   ErrorCode = "duplicateCompanyNumber"
	CurrencyChange                                           ErrorCode = "currencyChange"
	AdminUserAccountExists                                   ErrorCode = "adminUserAccountExists"
	CreditLimitReached                                       ErrorCode = "creditLimitReached"
	InsufficientCredit                                       ErrorCode = "insufficientCredit"
	ProjectLockFailed                                        ErrorCode = "projectLockFailed"
	ProjectUpdateFailed                                      ErrorCode = "projectUpdateFailed"
	MaximumAmountExceeded                                    ErrorCode = "maximumAmountExceeded"
	PaymentFailedError                                       ErrorCode = "paymentFailedError"
	PaymentNotFound                                          ErrorCode = "paymentNotFound"
	AddonNotFound                                            ErrorCode = "addonNotFound"
	InvoiceNotFound                                          ErrorCode = "invoiceNotFound"
	PriceListNotFound                                        ErrorCode = "priceListNotFound"
	CostLimitTooSmall                                        ErrorCode = "costLimitTooSmall"
	ProjectImportInvalidYaml                                 ErrorCode = "projectImportInvalidYaml"
	ProjectImportInvalidConfigStructure                      ErrorCode = "projectImportInvalidConfigStructure"
	ProjectImportInvalidParameter                            ErrorCode = "projectImportInvalidParameter"
	ProjectImportInvalidTypeVersion                          ErrorCode = "projectImportInvalidTypeVersion"
	ProjectImportMissingParameter                            ErrorCode = "projectImportMissingParameter"
	ProjectImportUnableToCreateServiceStack                  ErrorCode = "projectImportUnableToCreateServiceStack"
	ProjectImportProjectIncluded                             ErrorCode = "projectImportProjectIncluded"
	ProjectImportProjectMissing                              ErrorCode = "projectImportProjectMissing"
	ProjectImportMissingService                              ErrorCode = "projectImportMissingService"
	InvalidCustomAutoscalingValue                            ErrorCode = "invalidCustomAutoscalingValue"
	InvalidCustomAutoscalingCpuRamRatio                      ErrorCode = "invalidCustomAutoscalingCpuRamRatio"
	CustomVerticalAutoscalingForbiddenForThisStackCategory   ErrorCode = "customVerticalAutoscalingForbiddenForThisStackCategory"
	CustomHorizontalAutoscalingForbiddenForThisStackCategory ErrorCode = "customHorizontalAutoscalingForbiddenForThisStackCategory"
	CustomHorizontalAutoscalingForbiddenInNonHaMode          ErrorCode = "customHorizontalAutoscalingForbiddenInNonHaMode"
	ActionForbiddenForThisStackCategory                      ErrorCode = "actionForbiddenForThisStackCategory"
	ModeUpdateForbiddenForThisStackCategory                  ErrorCode = "modeUpdateForbiddenForThisStackCategory"
	YamlValidationInvalidYaml                                ErrorCode = "yamlValidationInvalidYaml"
	UnexpectedYamlEncoding                                   ErrorCode = "unexpectedYamlEncoding"
	BucketAlreadyExists                                      ErrorCode = "bucketAlreadyExists"
	S3ApiRequestFailed                                       ErrorCode = "s3ApiRequestFailed"
	UserIdWasNotFound                                        ErrorCode = "userIdWasNotFound"
	AuthorIdWasNotFound                                      ErrorCode = "authorIdWasNotFound"
	UserDoesNotHaveActiveConnectionToAnyClient               ErrorCode = "userDoesNotHaveActiveConnectionToAnyClient"
	UserHasNotAccessToZerops                                 ErrorCode = "userHasNotAccessToZerops"
	InvalidEmailToken                                        ErrorCode = "invalidEmailToken"
	ClientCreateFailed                                       ErrorCode = "clientCreateFailed"
	ClientIdTooLong                                          ErrorCode = "clientIdTooLong"
	TransferIdTooLong                                        ErrorCode = "transferIdTooLong"
	EmailIsInvalid                                           ErrorCode = "emailIsInvalid"
	UserAccountDeleted                                       ErrorCode = "userAccountDeleted"
	NoContaboClientConnected                                 ErrorCode = "noContaboClientConnected"
	AmbiguousUser                                            ErrorCode = "ambiguousUser"
	AmbiguousUserSignIn                                      ErrorCode = "ambiguousUserSignIn"
	UnsupportedCurrency                                      ErrorCode = "unsupportedCurrency"
	InvalidCurrencyCode                                      ErrorCode = "invalidCurrencyCode"
	TransferParamMismatch                                    ErrorCode = "transferParamMismatch"
	InvalidAmount                                            ErrorCode = "invalidAmount"
)
