package router

import (
	delivery "backend-challenge/api/shopping/delivery"
	repositories "backend-challenge/api/shopping/repositories"
	usecase "backend-challenge/api/shopping/usecase"

	"database/sql"

	"github.com/gin-gonic/gin"
)

func InitRoute(db *sql.DB, router *gin.Engine) {
	customerRepository := repositories.NewCustomerRepository(db)
	customerUsecase := usecase.NewCustomerUsecase(customerRepository)

	delivery.NewReportingController(router, customerUsecase)
}
