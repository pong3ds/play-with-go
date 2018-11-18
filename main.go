package main

import (
	"github.com/labstack/echo"
)

// IService is interfaceo for any services
type IService interface {
	RegisterServices(e *echo.Echo) error
}

func allServices() []IService {
	return []IService{
		NewBNKServiceHTTP(),
		NewAKBServiceHTTP(),
	}
}

func main() {
	e := echo.New()
	services := allServices()
	for _, service := range services {
		service.RegisterServices(e)
	}
	e.Start(":8080")
}
