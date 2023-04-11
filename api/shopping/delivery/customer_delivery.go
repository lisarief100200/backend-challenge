package delivery

import (
	usecase "backend-challenge/api/shopping/usecase"
	req "backend-challenge/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewReportingController(router *gin.Engine, customerUsecase usecase.CustomerUsecase) {
	customerController := NewCustomerController(customerUsecase)

	router.POST("/customers", customerController.CreateCustomer)
	router.GET("/customers", customerController.GetCustomers)
	router.PUT("/customers/:id", customerController.UpdatedCustomer)
	router.DELETE("/customers/:id", customerController.DeleteMovie)
}

type customerController struct {
	customerUsecase usecase.CustomerUsecase
}

func NewCustomerController(customerUsecase usecase.CustomerUsecase) *customerController {
	return &customerController{customerUsecase}
}

func (c *customerController) GetCustomers(cnx *gin.Context) {
	customers, err := c.customerUsecase.FindAll()
	if err != nil {
		cnx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var customersResponse []req.CustomerResponse

	for _, cus := range customers {
		customerResponse := req.CustomerResponse{
			ID:              cus.ID,
			CustomerName:    cus.CustomerName,
			CustomerContNo:  cus.CustomerContNo,
			CustomerAddress: cus.CustomerAddress,
			TotalBuy:        cus.TotalBuy,
		}

		customersResponse = append(customersResponse, customerResponse)
	}

	cnx.JSON(http.StatusOK, gin.H{
		"data": customersResponse,
	})
}

func (c *customerController) CreateCustomer(cnx *gin.Context) {
	var customerRequest req.CustomerRequest

	err := cnx.ShouldBindJSON(&customerRequest)
	if err != nil {
		cnx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	customer, err := c.customerUsecase.Create(customerRequest)

	if err != nil {
		cnx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	cnx.JSON(http.StatusOK, gin.H{
		"data": customer,
	})
}

func (c *customerController) UpdatedCustomer(cnx *gin.Context) {
	//title, price
	var customerRequest req.CustomerRequest

	// Agak beda sama tutor YouTube.
	err := cnx.ShouldBindJSON(&customerRequest)
	if err != nil {
		errorMessages := "Invalid Input"
		cnx.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	idString := cnx.Param("id")
	id, _ := strconv.Atoi(idString)

	customer, err := c.customerUsecase.Update(id, customerRequest)

	if err != nil {
		cnx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	customerResponse := req.CustomerResponse{
		ID:              customer.ID,
		CustomerName:    customer.CustomerName,
		CustomerContNo:  customer.CustomerContNo,
		CustomerAddress: customer.CustomerAddress,
		TotalBuy:        customer.TotalBuy,
	}
	cnx.JSON(http.StatusOK, gin.H{
		"data": customerResponse,
	})
}

func (c *customerController) DeleteMovie(cnx *gin.Context) {
	idString := cnx.Param("id")
	id, _ := strconv.Atoi(idString)

	customer, err := c.customerUsecase.Delete(id)

	if err != nil {
		cnx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	customerResponse := req.CustomerResponse{
		ID:              customer.ID,
		CustomerName:    customer.CustomerName,
		CustomerContNo:  customer.CustomerContNo,
		CustomerAddress: customer.CustomerAddress,
		TotalBuy:        customer.TotalBuy,
	}
	cnx.JSON(http.StatusOK, gin.H{
		"data": customerResponse,
	})
}
