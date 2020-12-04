package build

import (
	clients1 "github.com/pip-services-users/pip-clients-roles-go/version1"
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	cbuild "github.com/pip-services3-go/pip-services3-components-go/build"
)

type RolesClientFactory struct {
	cbuild.Factory
}

func NewRolesClientFactory() *RolesClientFactory {
	c := &RolesClientFactory{
		Factory: *cbuild.NewFactory(),
	}

	// nullClientDescriptor := cref.NewDescriptor("pip-services-roles", "client", "null", "*", "1.0")
	// directClientDescriptor := cref.NewDescriptor("pip-services-roles", "client", "direct", "*", "1.0")
	cmdHttpClientDescriptor := cref.NewDescriptor("pip-services-roles", "client", "commandable-http", "*", "1.0")
	grpcClientDescriptor := cref.NewDescriptor("pip-services-roles", "client", "grpc", "*", "1.0")
	memoryClientDescriptor := cref.NewDescriptor("pip-services-roles", "client", "memory", "*", "1.0")

	// c.RegisterType(nullClientDescriptor, clients1.NewRolesNullClientV1)
	// c.RegisterType(directClientDescriptor, clients1.NewRolesDirectClientV1)
	c.RegisterType(cmdHttpClientDescriptor, clients1.NewRolesHttpCommandableClientV1)
	c.RegisterType(grpcClientDescriptor, clients1.NewRoleGrpcClientV1)
	c.RegisterType(memoryClientDescriptor, clients1.NewRolesMemoryClientV1)

	return c
}
