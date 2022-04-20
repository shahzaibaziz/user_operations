package main

import (
	"github.com/go-openapi/loads"
	"github.com/sirupsen/logrus"

	runtime "github.com/shahzaibaziz/user_operations"
	"github.com/shahzaibaziz/user_operations/gen/restapi"
	"github.com/shahzaibaziz/user_operations/gen/restapi/operations"
	"github.com/shahzaibaziz/user_operations/handler"
)

func main() {
	log := logrus.WithField("pkg", "main")

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}
	rt, err := runtime.NewRuntime()
	if err != nil {
		log.Fatalln(err)
	}

	var server *restapi.Server

	api := operations.NewUserOperationsAPI(swaggerSpec)
	api.Logger = log.Infof

	handler.NewCustomHandler(api, rt)

	server = restapi.NewServer(api)
	server.Port = 8080

	defer server.Shutdown()

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
