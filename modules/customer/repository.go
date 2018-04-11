package customer

import (
	"golang-firestore/helper"
	"golang-firestore/services/firebase"
)

type customerrepository struct {
	Response helper.Response
}

/**
 * @brief      Saves a customer in fire store.
 * @param      requestData  The request data
 * @return     object
 */
func (this *customerrepository) SaveCustomerInFireStore(requestData RequestDataCustomer) error {

	params := map[string]interface{}{
		"id":             requestData.ID,
		"fullname":       requestData.Fullname,
		"email":          requestData.Email,
		"mobile":         requestData.Mobile,
		"gender":         requestData.Gender,
		"id_card_number": requestData.IdNumber,
		"location":       requestData.Location,
		"date_of_birth":  requestData.DateOfBirth,
		"note":           requestData.Note,
		"image_file":     requestData.ImageFile,
		"images":         requestData.Images,
		"is_deleted":     requestData.IsDeleted,
	}

	// send data to firebase
	return firebase.SendDataToFirestore(requestData.Subdomain, requestData.ID, params)
}

/**
 * @brief      declare customer firestore
 * @return     object
 */
func CustomerRepository() *customerrepository {
	var repository customerrepository

	return &repository
}
