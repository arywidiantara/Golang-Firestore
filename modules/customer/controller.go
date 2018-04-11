package customer

import (
	"golang-firestore/core"
	"golang-firestore/helper"

	"github.com/gin-gonic/gin"
)

type customercontroller struct {
	core.Controller
	CustomerRepository *customerrepository
}

/**
 * @brief      Saves a data customer.
 * @param      c
 * @return     object
 */
func (this *customercontroller) SaveDataCustomer(c *gin.Context) {
	var bindingCustomer BindingDataCustomer
	var requestCustomer RequestDataCustomer

	// check error binding data
	errorBinding := c.BindJSON(&bindingCustomer)

	if errorBinding != nil {
		// logging data error
		helper.LoggingError(errorBinding, "customer.controller.SaveDataCustomer.27")

		// return error data
		this.Response.SetStatus(core.STATUS_INVALID_PARAMETERS)
		err := map[string]interface{}{
			"errors": []string{
				"Invalid Parameters",
			},
		}
		this.Response.SetData(err)
		c.JSON(400, this.Response.ToResponse())

	} else {

		// set request data
		requestCustomer.ID = bindingCustomer.ID
		requestCustomer.Subdomain = bindingCustomer.Subdomain
		requestCustomer.Fullname = bindingCustomer.Fullname
		requestCustomer.Email = bindingCustomer.Email
		requestCustomer.Mobile = bindingCustomer.Mobile
		requestCustomer.Gender = bindingCustomer.Gender
		requestCustomer.IdNumber = bindingCustomer.IdNumber
		requestCustomer.Location = bindingCustomer.Location
		requestCustomer.DateOfBirth = bindingCustomer.DateOfBirth
		requestCustomer.Note = bindingCustomer.Note
		requestCustomer.ImageFile = bindingCustomer.ImageFile
		requestCustomer.Images = bindingCustomer.Images
		requestCustomer.IsDeleted = bindingCustomer.IsDeleted

		// processing data request
		error := this.CustomerRepository.SaveCustomerInFireStore(requestCustomer)

		// check error data
		if error != nil {
			return_data := map[string]interface{}{
				"errors": error,
			}
			// return data
			this.Response.SetData(return_data)
			this.Response.SetStatus(core.STATUS_INVALID_PARAMETERS)
			c.JSON(400, this.Response.ToResponse())
		} else {
			return_data := map[string]interface{}{
				"message": "success send data",
			}
			// return data
			this.Response.SetData(return_data)
			this.Response.SetStatus(core.STATUS_SUCCESS)
			c.JSON(200, this.Response.ToResponse())
		}
	}
}

/**
 * @brief      this function like constract
 * @return     object
 */
func CustomerController() customercontroller {
	var controller customercontroller

	controller.CustomerRepository = CustomerRepository()

	return controller
}
