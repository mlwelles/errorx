// normal go errors go here, the ones that we don't expect to be shown in things like text field validation output in the UI.
// put those in "external.go"
package errorx

import (
	"fmt"
)

var ErrNilMVRViolation = New("nil mvr violation")
var ErrNilGroup = New("nil group")
var ErrAccountNotFound = AccountNotFound()

func AccountNotFound(errs ...error) error {
	errs = append([]error{NotFound()}, errs...)
	return New("account not found", errs...)
}

var ErrInvalidUserInfo = New("invalid user info")
var ErrNotImplemented = New("not implemented")
var ErrInvalidArguments = InvalidArguments()

func InvalidArguments(errs ...error) error {
	return New("invalid arguments", errs...)
}

var ErrNilArguments = Wrap(ErrInvalidArguments, "nil arguments not allowed")
var ErrAgeLimitExceeded = New("age exceeds policy requirement")
var ErrBeneficiaryAgeLimitExceeded = New("beneficiary age exceeds policy requirement")
var ErrInvalidCacheItem = New("cache item invalid")
var ErrInvalidAccountCacheable = New("invalid account cacheable")
var ErrIdentityNil = New("identity nil")
var ErrIdentityNoName = New("identity without name")
var ErrIdentityNoExternalID = New("identity without external ID")
var ErrIdentityNoAccountID = New("identity without account ID")

func NilAccount(errs ...error) error { return New("nil account", errs...) }

var ErrNilAccount = NilAccount()
var ErrNilPhoneNumber = New("nil phone number")
var ErrNilProvider = New("nil provider")
var ErrNoPhoneNumber = New("no phone number")
var ErrNilMessage = New("nil message")
var ErrNoAccountID = New("account has no ID")
var ErrNeedPersonID = New("need person ID", ErrInvalidArguments)
var ErrNeedEmail = New("need email address", ErrInvalidArguments)
var ErrIDChangeNotAllowedInUpdate = New("id field cannot be changed as part of an update", ErrInvalidArguments)
var ErrNoMVRProvider = New("no mvr provider enabled")
var ErrDuplicateIdentity = DuplicateIdentity()

func DuplicateIdentity(errs ...error) error {
	return New("duplicate identity", errs...)
}

var ErrIdentityNotFound = New("account identity not found")
var ErrUnableToDeterminePremiumRate = New("failed to determine rate for premium")
var ErrInterviewNotBound = New("interview is not bound to an account")
var ErrInterviewAlreadyBound = InterviewAlreadyBound()

func InterviewAlreadyBound(errs ...error) error {
	return New("interview already bound", errs...)
}

var ErrInterviewBind = InterviewBind()

func InterviewBind(errs ...error) error {
	return New("interview binding error", errs...)
}

var ErrNoInterviews = New("no interviews")
var ErrInterviewNotFound = New("interview not found", NotFound())
var ErrInterviewService = New("interview service error")
var ErrRetryable = New("retryable")
var ErrTemporarilyUnavailable = New("temporarily unavailable")
var ErrInterviewNotSubmittable = New("interview is not ready for policy creation")
var ErrInterviewIncomplete = New("interview is incomplete")
var ErrInterviewReconciliationMaxAttempts = New("interview reconciliation max attempts limit reached")
var ErrInvalidAnswer = InvalidAnswer()

func InvalidAnswer(errs ...error) error {
	return New("invalid answer provided", errs...)
}

var ErrNilAddress = New("nil address")

func NilAddress(errs ...error) error {
	return New("nil address", errs...)
}

var ErrInvalidAddress = New("invalid address", ErrInvalidAnswer)
var ErrInvalidHomeAddress = New("invalid home address")
var ErrNilHomeAddresser = New("home address source is nil")
var ErrInvalidAddressLine1 = New("invalid addresss line 1")
var ErrInvalidAddressLine2 = New("invalid address line 2")
var ErrInvalidCity = New("invalid city")
var ErrInvalidState = New("invalid state")
var ErrNoPostalCode = New("invalid postal code")
var ErrInvalidExternalID = ExternalIDInvalid()
var ErrDealNotFound = New("pipedrive deal not found")

func ExternalIDInvalid(errs ...error) error {
	return New("external id invalid", errs...)
}

var ErrEmptyExternalID = ExternalIDEmpty()

func ExternalIDEmpty(errs ...error) error {
	return New("empty external id", ExternalIDInvalid())
}

var ErrInvalidRateTable = New("invalid rate table for quote")
var ErrNoTermPeriod = New("interview is missing term period")
var ErrNilReciever = New("method called on nil reciever")
var ErrNilAuthUser = New("nil auth user")
var ErrNilAuth0User = New("nil auth0 user")
var ErrNilAuthUserInfo = New("nil auth user info")
var ErrNilAuthIdentity = New("nil auth identity")
var ErrNilInterview = NilInterview()

func NilInterview(errs ...error) error {
	return New("nil interview", errs...)
}

var ErrNilSummary = New("nil summary")
var ErrEmptyInterviewID = New("interview ID is empty string")
var ErrDuplicateInterviewFlag = New("duplicate interview flag")
var ErrNilClient = New("nil client")
var ErrNilCondition = New("nil conditional")
var ErrMaxAttemptsReached = New("max attempts reached")
var ErrConditionCopy = New("error copying conditional")
var ErrNilSection = New("error nil section")
var ErrSectionCopy = New("error copying section")
var ErrNilFlag = New("nil flag")
var ErrFlagCopy = New("error copying flag")
var ErrNilEvent = New("nil event")
var ErrEventCopy = New("error copying event")
var ErrQuoteCopy = New("error copying quote")
var ErrQuoteEventCopy = New("error copying quote event")
var ErrQuoteEventCreate = New("creating quote event")
var ErrNilQuote = New("nil quote")
var ErrNilStatusEvent = New("nil status event")
var ErrStatusEventCopy = New("error copying status event")
var ErrNilAnswerEvent = New("nil answer event")
var ErrAnswerEventCopy = New("error copying answer event")
var ErrInvalidOrder = New("invalid order")
var ErrNilOrder = New("nil order")
var ErrNilSettings = New("nil settings")
var ErrInternal = New("internal error")
var ErrNotAnswered = New("no answer for question", NotFound())
var ErrNilAnswer = NilAnswer()

func NilAnswer(errs ...error) error {
	wrap := InvalidAnswer(errs...)
	return New("answer is nil", wrap)
}

var ErrNoQuestion = New("question is not part of interview")
var ErrInvalidResource = New("invalid resource")
var ErrEmptyResourceURN = Wrap(ErrInvalidResource, "resource URN is empty string")
var ErrAlreadyExists = New("already exists")
var ErrPermissionDenied = PermissionDenied()

func PermissionDenied(errs ...error) error {
	return New("permission denied", errs...)
}

var ErrUnauthorized = Unauthorized()

func Unauthorized(errs ...error) error {
	return New("unauthorized", errs...)
}

var ErrPermissionNotUnderwriter = PermissionNotUnderwriter()

func PermissionNotUnderwriter(errs ...error) error {
	errs = append([]error{PermissionDenied()}, errs...)
	return New("account does not have underwriter privileges", errs...)
}

var ErrNotOwner = Wrap(ErrPermissionDenied, "not owner")
var ErrEmptyPermissionSetRequested = New("empty permission set requested")
var ErrNilPermissionSet = New("nil permission set")
var ErrCannotModify = Wrap(ErrPermissionDenied, "cannot modify")
var ErrCannotModifyNotOwner = WrapError(ErrCannotModify, ErrNotOwner)
var ErrCannotModifyNeedsReview = Wrap(ErrCannotModify, "interview needs review by staff")
var ErrCannotModifySubmitted = Wrap(ErrCannotModify, "cant modify a submitted interview")
var ErrCannotModifySubmissionInProgress = Wrap(ErrCannotModifySubmitted, "submission in progress")
var ErrCannotModifyAccepted = Wrap(ErrCannotModify, "interview is accepted")
var ErrCannotModifyHasPolicy = Wrap(ErrCannotModify, "interview has policy")
var ErrCannotModifyExpired = Wrap(ErrCannotModify, "interview has expired")
var ErrCannotModifyInactive = Wrap(ErrCannotModify, "inactive interview")
var ErrCannotModifyRejected = Wrap(ErrCannotModify, "rejected interview")
var ErrCannotModifyError = Wrap(ErrCannotModify, "interview has error status")
var ErrCannotModifyFrozen = Wrap(ErrCannotModify, "interview frozen")
var ErrCannotModifyAnonymously = Wrap(ErrCannotModify, "interview cannot be modified anonymously")
var ErrCannotRead = Wrap(ErrPermissionDenied, "cannot read")
var ErrCannotReadAnonymously = Wrap(ErrCannotRead, "cannot be read anonymously")
var ErrCannotReadNotOwner = WrapError(ErrCannotRead, ErrNotOwner)
var ErrInvalidInterviewStatus = New("invalid interview status")
var ErrInvalidInterview = New("invalid interview")
var ErrPolicyInvalidStatus = New("invalid policy status")
var ErrPolicyCreate = New("policy creation failure")
var ErrPolicyReconcile = New("policy reconciliation error")
var ErrPolicyUpdate = New("policy update error")
var ErrPolicyList = New("error listing policies")
var ErrPolicyInvoiceCreating = New("error creating policy invoice")
var ErrPolicyInvoicePay = New("error paying policy invoice")
var ErrPolicyFlagResolve = New("error resolving policy flag")
var ErrInterviewInvalidForPolicyCreation = New("interview invalid for policy creation")
var ErrInterviewPolicyIDAlreadyExists = New("interview already has policy id assigned")
var ErrInvalidBeneficiary = New("invalid beneficiary")

func NilContext(errs ...error) error {
	return New("nil context", errs...)
}

func InvalidTx(errs ...error) error {
	return New("invalid transaction", errs...)
}

func InvalidCtx(errs ...error) error {
	return New("invalid context", errs...)
}

func InvalidTxContext(errs ...error) error {
	return New("invalid transaction context", append([]error{InvalidTx()}, errs...)...)
}

var ErrInvalidTxContext = InvalidTxContext()

var ErrEmptyChangeSet = New("empty change set")
var ErrRpcSqlNoRows = New("rpc error: code = Unknown desc = sql: no rows in result set")
var ErrSqlNoRows = New("sql: no rows in result set")
var ErrNoTimes = New("no times")

// MQ errors

func Newf(msg string, args ...interface{}) error {
	return New(fmt.Sprintf(msg, args...))
}

func ErrFailedToConnectToBroker(addr string) error {
	return Newf("failed connecting to message broker at %s", addr)
}

var ErrFailedToDisconnectFromBroker = New("failed disconnecting message broker")
var ErrInvalidChannel = New("invalid message broker channel")

func ErrFailedToOpenChannel(addr string) error {
	return Newf("failed to open message broker channel to %s", addr)
}

var ErrFailedToCloseChannel = New("failed to close message broker channel")
var ErrFailedToPublishMessage = New("failed to publish a message to queue")
var ErrFailedToRequeueMessage = New("failed to requeue a message to queue")
var ErrRetryMaxExceeded = New("allowed retries have been exhausted")
var ErrNilQueueHandler = New("invalid queue handler")
var ErrInvalidQueueHandler = New("invalid queue handler")
var ErrFailBatch = New("failed to connect to queue for batch, returning nil")
var ErrFailedToCreateExchange = New("failed to create an exchange")
var ErrFailedToCreateQueue = New("failed to create a queue")
var ErrFailedToBindQueue = New("failed to bind queue")
var ErrFailedToSetQOS = New("failed to set QOS/prefetch for queue")
var ErrFailedToConsume = New("error consuming message from queue")
var ErrInstantiatingLogger = New("error instantiating logger for publisher, tried falling back to shared")

// Not found errors
var ErrNotFound = NotFound()

func NotFound(errs ...error) error { return New("not found", errs...) }

var ErrFileNotFound = New("file not found", NotFound())
var ErrPrivilegeNotFound = New("privilege not found", NotFound())

func AnswerValidation(errs ...error) error {
	return New("error validating answer", errs...)
}

func QuestionKeyNotFound(key Stringer, errs ...error) error {
	msg := fmt.Sprintf("question key %s not found", key.String())
	return New(msg, errs...)
}

//var ErrQuestionKeyNotFound = New("question key not found", NotFound())
var ErrPersonNotFound = New("person not found", NotFound())
var ErrMedicationNotFound = New("medication not found", NotFound())

func OccupationNotFound(errs ...error) error {
	errs = append([]error{NotFound()}, errs...)
	return New("occupation not found", errs...)
}

var ErrOccupationNotFound = OccupationNotFound()

func OccupationParse(key string) error {
	return OccupationNotFound(fmt.Errorf("can't parse occupation from %q", key))
}

func ErrOccupationParseKey(key string) error {
	return OccupationNotFound(fmt.Errorf("invalid occupation key %q", key))
}

func OccupationParseName(name string) error {
	return Combine(ErrOccupationNotFound, fmt.Errorf("invalid occupation name %q", name))
}

var ErrAccountFromContext = New("error getting context account")
var ErrRiskScoring = New("error calculating risk score")
var ErrUndeterminedRate = New("rate could not be determined")
var ErrInterviewCopy = New("error copying interview")
var ErrExampleCopy = New("error copying example")
var ErrQuestionCopy = New("error copying question")
var ErrAnswerCopy = New("error copying answer")
var ErrScoreCopy = New("error copying score")
var ErrOccupation = New("occupation error")
var ErrMarshal = New("marshaling error")
var ErrEventMarshal = Wrap(ErrMarshal, "marshaling event")

func ErrDiagnosisNotFound(str string) error {
	return Combine(NotFound(), fmt.Errorf("no diagnosis found matching '%s'", str))
}

var ErrRepository = New("repository error")
var ErrInvalidNicotineFrequency = New("invalid nicotine frequency")
var ErrInvalidIrixOrderable = New("invalid irix orderable")
var ErrNilDiagnosis = New("nil diagnosis")
var ErrNilDiagnosisCategory = New("nil diagnosis category")
var ErrNilExample = New("nil example")
var ErrNilOccupation = New("occupation is nil")
var ErrInvalidOccupation = New("invalid occupation")
var ErrPreviousOccupation = New("previous occupation")
var ErrNilPreviousOccupation = WrapError(ErrNilOccupation, ErrPreviousOccupation)
var ErrUndefinedInteractionType = New("unhandled interaction type")
var ErrNilInteraction = New("nil interaction")
var ErrUnmarshal = New("unmarshaling error")
var ErrEventUnmarshal = Wrap(ErrUnmarshal, "unmarshalling event")
var ErrReflexiveCopy = New("error copying reflexive")
var ErrInterviewCreate = New("error creating interview")
var ErrInterviewUpdate = New("error updating interview")
var ErrCachegroup = New("cachegroup error")
var ErrInterviewTemplateCache = New("interview template cache error")
var ErrNilPolicy = New("nil policy")
var ErrNilPerson = New("nil person")
var ErrNilProduct = New("nil product")
var ErrNilScore = New("nil score")
var ErrNeedMoreInformation = New("more information needed to perform operation")
var ErrNilInvoice = New("invoice unexpectedly nil")
var ErrInvoiceAlreadyPaid = New("invoice has already been paid")
var ErrInvoiceNotFound = Wrapf(ErrNotFound, "invoice not found")
var ErrNilBeneficiary = New("beneficiary unexpectedly nil")
var ErrPaymentFailed = New("payment failed")
var ErrNoUserInfo = New("no userinfo for token")

var ErrInvalidToken = InvalidToken()

func InvalidToken(errs ...error) error {
	return New("invalid token", errs...)
}

var ErrNilClaims = NilClaims()

func NilClaims(errs ...error) error {
	return New("nil claims", errs...)
}

var ErrInvalidClaims = InvalidClaims()

func InvalidClaims(errs ...error) error {
	return New("invalid claims", errs...)
}

var ErrClient = Client()

func Client(errs ...error) error {
	return New("client error", errs...)
}

func NoToken(errs ...error) error {
	return Wrap(InvalidToken(), "no token", append([]error{NotFound()}, errs...)...)
}

var ErrExpiredToken = ExpiredToken()

func ExpiredToken(errs ...error) error {
	return New("expired token", errs...)
}

var ErrCreatingPerson = New("error creating person")
var ErrInvalidGender = New("invalid gender")
var ErrInvalidTrustType = New("invalid trust type")
var ErrInvalidRelationship = New("invalid relationship")
var ErrInvalidVersion = New("invalid version")
var ErrInvalidSelection = New("invalid selection")
var ErrInvalidIdentity = New("invalid identity")
var ErrCountryCodeInvalid = New("invalid country code", ErrInvalidAnswer)
var ErrPhoneNumberInvalid = New("Please enter a valid phone number.", ErrInvalidAnswer)
var ErrNilQuestion = New("question cannot be nil")
var ErrNilQuery = new("nil query")
var ErrInvalidReflexiveSource = New("invalid source for reflexive")
var ErrIncorrectReflexiveSource = New("question key and reflexive source key do not match", ErrInvalidReflexiveSource)
var ErrInvalidQuestionKey = New("invalid question key")
var ErrValueIsNegative = New("value cannot be negative")
var ErrValueTooHigh = New("value is too high")
var ErrValueTooLow = New("value is too low")
var ErrAuth0UserUpdateFailed = New("auth0 user update failed")
var ErrInvalidKey = New("invalid key")
var ErrAnswerRequired = New("Please enter a value.")
var ErrNilRequest = New("unexpectedly received a nil request")
var ErrNilResponse = New("unexpectedly received a nil response")
var ErrNilReflexive = New("reflexive is nil")
var ErrInvalidResponse = New("invalid response")

var ErrUnknownIncome = New("unable to determine covered income", ErrInvalidAnswer)

var ErrAssertionFailed = New("assertion failed")
var ErrAssertStatusFailed = New("interview status assertion failed", ErrAssertionFailed)
var ErrAssertRiskScoreFailed = New("interview risk score assertion failed", ErrAssertionFailed)

var ErrUnknownAddressState = New("state not understood")

var ErrExamOneTransaction = New("ExamOne Transaction error")
var ErrExamOneInvalidTransactionID = New("mismatched transaction ID", ErrExamOneTransaction)

var ErrInsufficientIncome = New("income too low to generate quote")
var ErrIncomeTooHigh = New("income too high to generate quote")
var ErrIllogicalTermPeriod = New("unable to determine a sensible term period")
var ErrNoFirstName = New("no first name")
var ErrNoLastName = New("no last name")
var ErrNoDateOfBirth = New("no date of birth")
var ErrNoHomeAddress = New("no home address")
var ErrWeightZeroPounds = New("weight is zero pounds")
var ErrInvalidWeight = New("invalid weight")
var ErrInvalidHeight = New("invalid height")
var ErrHeightZeroInches = New("height is zero inches")
var ErrNoHeight = New("no height")
var ErrNoGender = New("no gender")
var ErrNilDate = New("nil date")
var ErrNoOccupation = New("no occupation")
var ErrNoNicotineFrequency = New("no nicotine frequency")
var ErrUnknown = New("unknown")
var ErrInvalidZIPCode = New("invalid ZIP code")

var ErrInvalidInstantIDOrderable = New("invalid instant id orderable")
var ErrInvalidLifeDataPrefillOrderable = New("invalid life data prefill orderable")
var ErrInvalidMVROrderable = New("invalid mvr orderable")
var ErrInvalidDate = New("invalid date")
var ErrInvalidDateOfBirth = New("invalid date of birth")
var ErrNoDriversLicenseNumber = New("no drivers license number")
var ErrInvalidHealthPiqtureOrderable = New("invalid health piqture orderable")

var ErrInvalidCurrentlyWorking = New("invalid curently working")
var ErrInvalidNotCurrentlyWorkingReason = New("invalid not currently working reason")
var ErrEmptyString = New("empty string")
var ErrIDMismatch = fmt.Errorf("id mismatch")
var ErrInvalidAccountID = New("invalid account id")
var ErrNilSession = New("nil session")
var ErrEmptyTrackingEvent = New("empty tracking event")
var ErrTrackingRequestError = New("error making tracking request to external tracking api")

var ErrNilOrderResult = New("nil order result")
var ErrOrderRetryableNilResult = WrapError(ErrOrderRetryable, New("nil order result"))
var ErrNilProductResults = WrapError(ErrOrder, New("nil product results"))
var ErrProductResultsNilMotorVehicleReport = New("product results motor vehicle report field nil")
var ErrProductResultsNilInstantID = WrapError(ErrOrder, New("product results instant id field nil"))
var ErrProductResultsInstantIDReportNil = WrapError(ErrOrder, New("product results instant id report field nil"))
var ErrInstantIDResponseNilResponse = WrapError(ErrOrder, New("instant id response nil response"))
var ErrInstantIDResponseNilResult = WrapError(ErrOrder, New("instant id response nil result"))
var ErrOrder = New("order error")
var ErrOrderRetryable = Combine(ErrOrder, ErrRetryable)
var ErrMismatchIndicator = New("mismatch indicator")
var ErrMessageType = New("message type error received")
var ErrOrderDecode = Combine(ErrOrder, New("decoding order"))
var ErrNilWorker = New("nil worker")
var ErrNilPublisher = New("nil publisher")
var ErrInterviewStale = New("stale interview", ErrRetryable)

func InvalidSectionDisplayOption(errs ...error) error {
	return New("invalid section display option", errs...)
}

func HealthPiqture(errs ...error) error {
	return new("health piqture order error", errs...)
}

func Irix(errs ...error) error {
	return new("irix order error", errs...)
}

func LifeDataPrefill(errs ...error) error {
	return new("life data prefill order error", errs...)
}

func MVR(errs ...error) error {
	return new("mvr order error", errs...)
}

func InstantID(errs ...error) error {
	return new("instant id order error", errs...)
}

var ErrNilUpdatedTimestamp = NilUpdatedTimestamp()

func NilUpdatedTimestamp(errs ...error) error {
	return new("nil updated time", errs...)
}
