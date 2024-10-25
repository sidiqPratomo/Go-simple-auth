package appconstant

type contextKey string

const (
	AddressIdString  = "address_id"
	CategoryIdString = "category_id"

	AddressIdKey = contextKey(AddressIdString)

	AccountIdKey = contextKey(AccountId)

	DrugIdString            = "drug_id"
	CartIdString            = "cart_id"
	RoomIdString            = "room_id"
	PharmacyManagerIdString = "pharmacy_manager_id"
	PharmacyIdString        = "pharmacy_id"
	OrderIdString           = "order_id"
	PrescriptionIdString    = "prescription_id"
	OrderPharmacyIdString   = "order_pharmacy_id"
	PharmacyDrugIdString    = "pharmacy_drug_id"
	DoctorIdString          = "doctor_id"
)
