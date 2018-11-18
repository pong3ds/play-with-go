package main

// IService is interfaceo for any services
type IService interface {
	RegisterServices(e *Engine) error
}

func allServices() []IService {
	return []IService{
		NewBNKServiceHTTP(),
		NewAKBServiceHTTP(),
	}
}

func main() {
	e := NewEngine()
	services := allServices()
	for _, service := range services {
		service.RegisterServices(e)
	}
	e.Start(":8080")
}
