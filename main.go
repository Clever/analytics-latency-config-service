package main

import (
	"context"
	"flag"
	"log"

	"github.com/Clever/analytics-latency-config-service/gen-go/models"
	"github.com/Clever/analytics-latency-config-service/gen-go/server"
	"github.com/Clever/wag/swagger"
)

// MyController implements server.Controller
type MyController struct{}

var _ server.Controller = MyController{}

// HealthCheck handles GET requests to /_health
func (mc MyController) HealthCheck(ctx context.Context) error {
	return nil
}

// GetThings handles GET requests to /v2/things
func (mc MyController) GetThings(ctx context.Context) ([]models.Thing, error) {
	return nil, models.InternalError{Message: "TODO, this is just an example"}
}

// DeleteThing handles DELETE requests to /v2/things/{id}
func (mc MyController) DeleteThing(ctx context.Context, id string) error {
	return models.InternalError{Message: "TODO, this is just an example"}
}

// GetThing handles GET requests to /v2/things/{id}
func (mc MyController) GetThing(ctx context.Context, id string) (*models.Thing, error) {
	return nil, models.NotFound{Message: "TODO, this is just an example"}
}

// CreateOrUpdateThing handles PUT requests to /v2/things/{id}
func (mc MyController) CreateOrUpdateThing(ctx context.Context, i *models.CreateOrUpdateThingInput) (*models.Thing, error) {
	return nil, models.InternalError{Message: "TODO, this is just an example"}
}

func main() {
	addr := flag.String("addr", ":6723", "Address to listen at")
	flag.Parse()

	swagger.InitCustomFormats()

	myController := MyController{}
	s := server.New(myController, *addr)

	if err := s.Serve(); err != nil {
		log.Fatal(err)
	}

	log.Println("analytics-latency-config-service exited without error")
}
