package interfaces

import "github.com/yamyy123/netxd-dal/models"

type Icustomer interface{
	CreateCustomer(customer *models.CustomerRequest)(*models.CustomerResponse,error)
}