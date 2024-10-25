package apperror

import (
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/sidiqPratomo/DJKI-Pengaduan/appconstant"
)

type AppError struct {
	Code    int
	Err     error
	Message string
	stack   []byte
}

func (ae AppError) Error() string {
	return fmt.Sprintf("Error %d %s from: %s", ae.Code, ae.Message, ae.Err.Error())
}

func (ae *AppError) GetStackTrace() []byte {
	return ae.stack
}

func NewAppError(code int, err error, message string) *AppError {
	return &AppError{
		Code:    code,
		Err:     err,
		Message: message,
		stack:   debug.Stack(),
	}
}

func UnauthorizedError() *AppError {
	err := errors.New(appconstant.MsgUnauthorized)
	return NewAppError(http.StatusUnauthorized, err, appconstant.MsgUnauthorized)
}

func InvalidCartItemError() *AppError {
	err := errors.New(appconstant.MsgInvalidCartItem)
	return NewAppError(http.StatusBadRequest, err, appconstant.MsgInvalidCartItem)
}

func NotFoundError() *AppError {
	err := errors.New(appconstant.MsgNotFound)
	return NewAppError(http.StatusNotFound, err, appconstant.MsgNotFound)
}

func EmailNotFoundError() *AppError {
	err := errors.New(appconstant.MsgEmailNotFound)
	return NewAppError(http.StatusUnauthorized, err, appconstant.MsgEmailNotFound)
}

func AccountNotFoundError() *AppError {
	err := errors.New(appconstant.MsgAccountNotFound)
	return NewAppError(http.StatusUnauthorized, err, appconstant.MsgAccountNotFound)
}

func PharmacyManagerNotFoundError() *AppError {
	err := errors.New(appconstant.MsgPharmacyManagerNotFound)
	return NewAppError(http.StatusUnauthorized, err, appconstant.MsgPharmacyManagerNotFound)
}

func UserNotFoundError() *AppError {
	err := errors.New(appconstant.MsgUserNotFound)
	return NewAppError(http.StatusUnauthorized, err, appconstant.MsgUserNotFound)
}

func UserAddressNotFoundError() *AppError {
	err := errors.New(appconstant.MsgUserAddressNotFound)
	return NewAppError(http.StatusUnauthorized, err, appconstant.MsgUserAddressNotFound)
}

func LocationNotFoundError() *AppError {
	err := errors.New(appconstant.MsgLocationNotFound)
	return NewAppError(http.StatusNotFound, err, appconstant.MsgLocationNotFound)
}

func PartnerNotFoundError() *AppError {
	err := errors.New(appconstant.MsgPartnerNotFound)
	return NewAppError(http.StatusNotFound, err, appconstant.MsgPartnerNotFound)
}

func CartItemNotFoundError() *AppError {
	err := errors.New(appconstant.MsgCartItemNotFound)
	return NewAppError(http.StatusBadRequest, err, appconstant.MsgCartItemNotFound)
}

func CourierNotFoundError() *AppError {
	err := errors.New(appconstant.MsgCourierNotFound)
	return NewAppError(http.StatusBadRequest, err, appconstant.MsgCourierNotFound)
}

func UnauthorizedUserCartAccessError() *AppError {
	err := errors.New(appconstant.MsgUnauthorizedUserCartAccess)
	return NewAppError(http.StatusBadRequest, err, appconstant.MsgUnauthorizedUserCartAccess)
}

func AccountNotVerifiedError() *AppError {
	err := errors.New(appconstant.MsgAccountNotVerified)
	return NewAppError(http.StatusUnauthorized, err, appconstant.MsgAccountNotVerified)
}

func ForbiddenAction() *AppError {
	err := errors.New(appconstant.MsgForbiddenAction)
	return NewAppError(http.StatusForbidden, err, appconstant.MsgForbiddenAction)
}

func WrongPasswordError(err error) *AppError {
	return NewAppError(http.StatusUnauthorized, err, appconstant.MsgWrongPassword)
}

func RefreshTokenExpiredError() *AppError {
	err := errors.New(appconstant.MsgRefreshTokenExpired)
	return NewAppError(http.StatusUnauthorized, err, appconstant.MsgRefreshTokenExpired)
}

func ExpiredTokenError() *AppError {
	err := errors.New(appconstant.MsgExpiredToken)
	return NewAppError(http.StatusUnauthorized, err, appconstant.MsgExpiredToken)
}

func InvalidTokenError() *AppError {
	err := errors.New(appconstant.MsgInvalidToken)
	return NewAppError(http.StatusUnauthorized, err, appconstant.MsgInvalidToken)
}

func ExpiredCodeError() *AppError {
	err := errors.New(appconstant.MsgExpiredCode)
	return NewAppError(http.StatusUnauthorized, err, appconstant.MsgExpiredCode)
}

func InvalidCodeError() *AppError {
	err := errors.New(appconstant.MsgInvalidCode)
	return NewAppError(http.StatusUnauthorized, err, appconstant.MsgInvalidCode)
}

func EmailTakenError() *AppError {
	err := errors.New(appconstant.MsgEmailTaken)
	return NewAppError(http.StatusConflict, err, appconstant.MsgEmailTaken)
}

func InvalidNameError(err error) *AppError {
	return NewAppError(http.StatusBadRequest, err, appconstant.MsgInvalidName)
}

func OldPasswordReusedError() *AppError {
	err := errors.New(appconstant.MsgOldPasswordReused)
	return NewAppError(http.StatusBadRequest, err, appconstant.MsgOldPasswordReused)
}

func EmptyDataError() *AppError {
	err := errors.New(appconstant.MsgEmptyData)
	return NewAppError(http.StatusBadRequest, err, appconstant.MsgEmptyData)
}

func FileNotAttachedError() *AppError {
	err := errors.New(appconstant.MsgFileNotAttached)
	return NewAppError(http.StatusBadRequest, err, appconstant.MsgFileNotAttached)
}

func BadRequestError(err error) *AppError {
	return NewAppError(http.StatusBadRequest, err, err.Error())
}

func InternalServerError(err error) *AppError {
	return NewAppError(http.StatusInternalServerError, err, appconstant.MsgInternalServerError)
}

func AddressIdInvalidError() *AppError {
	err := errors.New(appconstant.MsgAddressIdInvalid)
	return NewAppError(http.StatusBadRequest, err, appconstant.MsgAddressIdInvalid)
}

func DrugIdInvalidError() *AppError {
	err := errors.New(appconstant.MsgDrugIdInvalid)
	return NewAppError(http.StatusBadRequest, err, appconstant.MsgDrugIdInvalid)
}

func CoordinateInvalidError() *AppError {
	err := errors.New(appconstant.MsgInvalidCoordinate)
	return NewAppError(http.StatusBadRequest, err, appconstant.MsgInvalidCoordinate)
}

func InvalidSortByError() *AppError {
	err := errors.New(appconstant.MsgInvalidSortBy)
	return NewAppError(http.StatusBadRequest, err, appconstant.MsgInvalidSortBy)
}

func InvalidSortError() *AppError {
	err := errors.New(appconstant.MsgInvalidSort)
	return NewAppError(http.StatusBadRequest, err, appconstant.MsgInvalidSort)
}

func InvalidSortPairError() *AppError {
	err := errors.New(appconstant.MsgInvalidSortPair)
	return NewAppError(http.StatusBadRequest, err, appconstant.MsgInvalidSortPair)
}

func InvalidCategoryError() *AppError {
	err := errors.New(appconstant.MsgInvalidCategory)
	return NewAppError(http.StatusBadRequest, err, appconstant.MsgInvalidCategory)
}

func InvalidMinPriceError() *AppError {
	err := errors.New(appconstant.MsgInvalidMinPrice)
	return NewAppError(http.StatusBadRequest, err, appconstant.MsgInvalidMinPrice)
}

func InvalidMaxPriceError() *AppError {
	err := errors.New(appconstant.MsgInvalidMaxPrice)
	return NewAppError(http.StatusBadRequest, err, appconstant.MsgInvalidMaxPrice)
}

func InvalidPriceRangeError() *AppError {
	err := errors.New(appconstant.MsgInvalidPriceRange)
	return NewAppError(http.StatusBadRequest, err, appconstant.MsgInvalidPriceRange)
}

func InvalidPageError() *AppError {
	err := errors.New(appconstant.MsgInvalidPage)
	return NewAppError(http.StatusBadRequest, err, appconstant.MsgInvalidPage)
}

func InvalidLimitError() *AppError {
	err := errors.New(appconstant.MsgInvalidLimit)
	return NewAppError(http.StatusBadRequest, err, appconstant.MsgInvalidLimit)
}

func DrugNotFoundError() *AppError {
	err := errors.New(appconstant.MsgDrugNotFound)
	return NewAppError(http.StatusNotFound, err, appconstant.MsgDrugNotFound)
}

func CategoryNotFoundError() *AppError {
	err := errors.New(appconstant.MsgCategoryNotFound)
	return NewAppError(http.StatusBadRequest, err, appconstant.MsgCategoryNotFound)
}

func CategoryNotUniqueError() *AppError {
	err := errors.New(appconstant.MsgCategoryNotUnique)
	return NewAppError(http.StatusBadRequest, err, appconstant.MsgCategoryNotUnique)
}

func DrugNameAlreadyExistError() *AppError {
	err := errors.New(appconstant.MsgDrugNameAlreadyExist)
	return NewAppError(http.StatusBadRequest, err, appconstant.MsgDrugNameAlreadyExist)
}

func ClassificationNotFoundError() *AppError {
	err := errors.New(appconstant.MsgClassificationNotFound)
	return NewAppError(http.StatusBadRequest, err, appconstant.MsgClassificationNotFound)
}

func DrugFormNotFoundError() *AppError {
	err := errors.New(appconstant.MsgDrugFormNotFound)
	return NewAppError(http.StatusBadRequest, err, appconstant.MsgDrugFormNotFound)
}

func DoctorNotFoundError() *AppError {
	err := errors.New(appconstant.MsgDoctorNotFound)
	return NewAppError(http.StatusBadRequest, err, appconstant.MsgDoctorNotFound)
}

func ChatRoomNotFoundError() *AppError {
	err := errors.New(appconstant.MsgChatRoomNotFound)
	return NewAppError(http.StatusForbidden, err, appconstant.MsgChatRoomNotFound)
}

func OnGoingChatExistError() *AppError {
	err := errors.New(appconstant.MsgOnGoingChatExists)
	return NewAppError(http.StatusForbidden, err, appconstant.MsgOnGoingChatExists)
}

func AbortPreviousListenRequestError() *AppError {
	err := errors.New(appconstant.MsgAbortPreviousListenRequestError)
	return NewAppError(http.StatusOK, err, appconstant.MsgAbortPreviousListenRequestError)
}

func EmptyCartSelectionError() *AppError {
	err := errors.New(appconstant.MsgEmptyCartSelection)
	return NewAppError(http.StatusBadRequest, err, appconstant.MsgEmptyCartSelection)
}

func RoomIsNowExpiredError() *AppError {
	err := errors.New(appconstant.MsgRoomIsNowExpired)
	return NewAppError(http.StatusForbidden, err, appconstant.MsgRoomIsNowExpired)
}

func InvalidOrderError() *AppError {
	err := errors.New(appconstant.MsgInvalidOrder)
	return NewAppError(http.StatusBadRequest, err, appconstant.MsgInvalidOrder)
}

func InsufficientStockDuringCheckoutError(list []int64) *AppError {
	err := errors.New(appconstant.MsgInsufficientStock + ":" + strings.Trim(strings.Join(strings.Fields(fmt.Sprint(list)), ","), ""))
	return NewAppError(http.StatusBadRequest, err, err.Error())
}

func OrderNotFoundError() *AppError {
	err := errors.New(appconstant.MsgOrderNotFound)
	return NewAppError(http.StatusBadRequest, err, appconstant.MsgOrderNotFound)
}

func InvalidOrderStatusError() *AppError {
	err := errors.New(appconstant.MsgInvalidOrderStatus)
	return NewAppError(http.StatusBadRequest, err, appconstant.MsgInvalidOrderStatus)
}

func PaymentProofIsEmptyError() *AppError {
	err := errors.New(appconstant.MsgPaymentProofIsEmpty)
	return NewAppError(http.StatusForbidden, err, appconstant.MsgPaymentProofIsEmpty)
}

func PrescriptionIdNotANumberError() *AppError {
	err := errors.New(appconstant.MsgPrescriptionIdNotANumber)
	return NewAppError(http.StatusBadRequest, err, appconstant.MsgPrescriptionIdNotANumber)
}

func InvalidPrescriptionIdError() *AppError {
	err := errors.New(appconstant.MsgPrescriptionIdInvalid)
	return NewAppError(http.StatusForbidden, err, appconstant.MsgPrescriptionIdInvalid)
}

func PrescriptionHasBeenRedeemedError() *AppError {
	err := errors.New(appconstant.MsgPrescriptionHasBeenRedeemed)
	return NewAppError(http.StatusForbidden, err, appconstant.MsgPrescriptionHasBeenRedeemed)
}

func DrugIsInactiveError() *AppError {
	err := errors.New(appconstant.MsgDrugIsInactive)
	return NewAppError(http.StatusForbidden, err, appconstant.MsgDrugIsInactive)
}

func PrescriptionHasBeenUsedError() *AppError {
	err := errors.New(appconstant.MsgPrescriptionHasBeenUsed)
	return NewAppError(http.StatusBadRequest, err, appconstant.MsgPrescriptionHasBeenUsed)
}

func NoDrugNearby() *AppError {
	err := errors.New(appconstant.MsgDrugNotAvailableInNearby)
	return NewAppError(http.StatusBadRequest, err, appconstant.MsgDrugNotAvailableInNearby)
}

func PharmacyOrderNotFoundError() *AppError {
	err := errors.New(appconstant.MsgPharmacyOrderNotFound)
	return NewAppError(http.StatusNotFound, err, appconstant.MsgPharmacyOrderNotFound)
}

func ChatRoomAlreadyClosedError() *AppError {
	err := errors.New(appconstant.MsgChatRoomAlreadyClosed)
	return NewAppError(http.StatusOK, err, appconstant.MsgChatRoomAlreadyClosed)
}

func PharmacyNotFoundError() *AppError {
	err := errors.New(appconstant.MsgPharmacyNotFound)
	return NewAppError(http.StatusNotFound, err, appconstant.MsgPharmacyNotFound)
}

func InvalidMaxDateError() *AppError {
	err := errors.New(appconstant.MsgInvalidMaxDate)
	return NewAppError(http.StatusBadRequest, err, appconstant.MsgInvalidMaxDate)
}

func InvalidMinDateError() *AppError {
	err := errors.New(appconstant.MsgInvalidMinDate)
	return NewAppError(http.StatusBadRequest, err, appconstant.MsgInvalidMinDate)
}

func InsufficientStockError() *AppError {
	err := errors.New(appconstant.MsgInsufficientStock)
	return NewAppError(http.StatusBadRequest, err, appconstant.MsgInsufficientStock)
}

func DuplicatePharmacyDrugIdError() *AppError {
	err := errors.New(appconstant.MsgDuplicateDrugId)
	return NewAppError(http.StatusBadRequest, err, appconstant.MsgDuplicateDrugId)
}

func InvalidStockMutationRequestError() *AppError {
	err := errors.New(appconstant.MsgInvalidStockMutationRequest)
	return NewAppError(http.StatusBadRequest, err, appconstant.MsgInvalidStockMutationRequest)
}

func InvalidPharmacyOperationalError() *AppError {
	err := errors.New(appconstant.MsgInvalidPharmacyOperational)
	return NewAppError(http.StatusBadRequest, err, appconstant.MsgInvalidPharmacyOperational)
}

func InvalidPharmacyCourierError() *AppError {
	err := errors.New(appconstant.MsgInvalidPharmacyCourier)
	return NewAppError(http.StatusBadRequest, err, appconstant.MsgInvalidPharmacyCourier)
}

func OngoingOrderExistsError() *AppError {
	err := errors.New(appconstant.MsgOngoingOrderExists)
	return NewAppError(http.StatusForbidden, err, appconstant.MsgOngoingOrderExists)
}
