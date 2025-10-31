// Errors that may end up being shown to the end user go here (validation errors and the like)
package errorx

import (
	"fmt"

	"github.com/dustin/go-humanize"
)

var ErrAgeAnswerInvalid = New("We had trouble understanding the date you entered. Please try again.", ErrInvalidAgeAnswer)
var ErrAnswerNumberNegative = New("Please enter a number that is greater than 0.")
var ErrChildAgeAnswer = New("We had trouble understanding your child's age. Please try again.", ErrInvalidAgeAnswer)
var ErrChildAgeAnswerTooFarInPast = New("Please correct your child's birthdate.", ErrChildAgeAnswer)
var ErrChildAgeMax = New("Your child's age is too old for us to support.", ErrChildAgeAnswer)
var ErrCurrencyAnswerNegative = New("Please enter an amount greater than 0.", ErrInvalidAnswer)
var ErrCurrencyAnswerTooHigh = New("Please enter an amount less than $10,000,000", ErrInvalidAnswer)
var ErrDateAnswer = New("Invalid date provided.", ErrInvalidAnswer)
var ErrDateAnswerInFuture = New("The date you entered is in the future. Please correct it.", ErrDateAnswer)
var ErrDateAnswerInPast = New("The date you entered is in the past. Please correct it.", ErrDateAnswer)
var ErrDateAnswerInvalid = New("Please enter a valid date.", ErrDateAnswer)
var ErrDateAnswerTooFarInFuture = New("The date you entered is too far in the future. Please correct it.", ErrDateAnswer)
var ErrDateAnswerTooFarInPast = New("The date you entered is too far in the past. Please correct it.", ErrDateAnswer)
var ErrDriversLicenseAnswerInvalid = New("The driver's license you entered is incorrect. Please try again.", ErrInvalidAnswer)
var ErrHeightAnswerInvalid = New("Please enter a height.", ErrInvalidAnswer)
var ErrHeightAnswerTooShort = New("Your height seems too short. Please try again.")
var ErrHeightAnswerTooTall = New("Your height seems too tall. Please try again.")

var ErrInsuredTooOld = New("Age of the insured is beyond prescribed limits.")
var ErrInsuredTooYoung = New("The insured is too young to be eligible.")
var ErrInvalidAgeAnswer = New("Please enter your correct birthday.", ErrInvalidAnswer)
var ErrInvalidApproximateDate = New("Please enter the correct month and year", ErrDateAnswer)
var ErrInvalidInsuredBirthDate = New("Invalid birthdate for the insured party.")
var ErrMaxAgeAnswer75 = New("Sorry, we can't insure people older than 75.", ErrInvalidAgeAnswer)
var ErrMiddleInitialAnswerTooLong = New("Please enter your middle initial.", ErrInvalidAnswer)
var ErrMinAgeAnswer18 = New("Sorry, you have to be 18 or older to apply.", ErrInvalidAgeAnswer)
var ErrNameAnswerTooLong = New("Your name seems long.  Please try again.", ErrInvalidAnswer)
var ErrNeedAnswer = New("Please provide an answer.")
var ErrNotWorkingReasonAnswersInvalid = New("The options you selected are incompatible. Please try again.", ErrInvalidAnswer)
var ErrPastDateAnswerNotPastDate = New("Please enter a date from the past.", ErrDateAnswer)
var ErrQuoteOverride = New("Invalid overrides on quote.")
var ErrSSNAnswerInvalid = New("Please enter a valid social security number.", ErrInvalidAnswer)
var ErrSelectedBenefitTooHigh = New("Twice-a-month benefit is above maximum value.", ErrQuoteOverride)
var ErrSelectedBenefitTooLow = New("Twice-a-month benefit is below minimum value.", ErrQuoteOverride)
var ErrSelectedPremiumTooHigh = New("Premium is above maximum value.", ErrQuoteOverride)
var ErrSelectedPremiumTooLow = New("Premium is below minimum value.", ErrQuoteOverride)
var ErrSelectedTermMonthsTooLong = New("Please select a shorter term length.", ErrQuoteOverride)
var ErrSelectedTermMonthsTooShort = New("Please select a longer term length.", ErrQuoteOverride)
var ErrSingleValueAnswerRequired = New("Please select a single answer.", ErrInvalidAnswer)
var ErrStateAnswerInvalid = New("Please enter your state.", ErrInvalidAnswer)
var ErrTotalExisitingInsuranceAnswerImpossible = New("Please tell us the correct total amount of life insurance you have.", ErrInvalidAnswer)
var ErrTrustFormationDateTooEarly = New("Please enter a date of April 20, 1676 or later (the formation date of the oldest active trust in the United States).")
var ErrTrustFormationDateTooLate = New("Your trust needs to be formed at least one day prior to signing your application.")
var ErrUnquoteableRiskScore = New("Manual premium calculation required for provided health class.")
var ErrWeightAnswerTooHigh = New("The weight you entered seems too high.", ErrInvalidAnswer)
var ErrWeightAnswerTooLow = New("The weight you entered seems too low.", ErrInvalidAnswer)
var ErrZipAnswerInvalid = New("Please enter a valid postal code.", ErrInvalidAnswer)
var ErrHouseholdIncomeBelowPersonal = New("Please enter the entire household income, including your income.", ErrInvalidAnswer)
var ErrAccountExistsForEmail = New("Account already exists for the email address entered.")

var ErrDisqualifyingAnswer = New("We are unable to cover you at this time.", ErrInvalidAnswer)
var ErrDisqualifyingBMI = New("BMI unacceptable", ErrDisqualifyingAnswer)
var ErrDisqualifyingAge = New("Age outside of acceptable range", ErrDisqualifyingAnswer)
var ErrDisqualifyingPhone = New("Phone number not valid US number", ErrDisqualifyingAnswer)
var ErrDisqualifyingDUI = New("DUI", ErrDisqualifyingAnswer)
var ErrDisqualifyingNicotineUse = New("Nicotine use", ErrDisqualifyingAnswer)

type AnnualIncomer interface {
	AnnualIncomeInputMax() int64
	AnnualIncomeInputMin() int64
	AnnualIncomeMin() float64
	HouseholdIncomeInputMin() int64
}

func ErrIncomeAnswerTooLow(mi AnnualIncomer) error {
	msg := fmt.Sprintf("Please enter an income greater than $%s.", humanize.Comma(mi.AnnualIncomeInputMin()))
	return New(msg, ErrInvalidAnswer)
}

func ErrIncomeAnswerTooHigh(mi AnnualIncomer) error {
	msg := fmt.Sprintf("Please enter an income less than $%s.", humanize.Comma(mi.AnnualIncomeInputMax()))
	return New(msg, ErrInvalidAnswer)
}

func ErrHouseholdIncomeInputTooLow(mi AnnualIncomer) error {
	msg := fmt.Sprintf("Your household income must be at least $%s.", humanize.Comma(mi.HouseholdIncomeInputMin()))
	return New(msg, ErrInvalidAnswer)
}

func ErrEffectiveIncomeTooLow(mi AnnualIncomer) error {
	msg := fmt.Sprintf("Please enter an income greater than $%s.", humanize.Comma(int64(mi.AnnualIncomeMin())))
	return New(msg, ErrInvalidAnswer)
}
