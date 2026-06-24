package dto

type CreateCustomerDTO struct {
	Name         string `json:"name"`
	Phone        string `json:"phone"`
	VehiclePlate string `json:"vehiclePlate"`
	VehicleBrand string `json:"vehicleBrand"`
	VehicleModel string `json:"vehicleModel"`
}

type UpdateCustomerDTO struct {
	Name         string `json:"name"`
	Phone        string `json:"phone"`
	VehiclePlate string `json:"vehiclePlate"`
	VehicleBrand string `json:"vehicleBrand"`
	VehicleModel string `json:"vehicleModel"`
}