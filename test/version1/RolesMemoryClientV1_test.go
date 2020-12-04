package test_version1

import (
	"testing"

	"github.com/pip-services-users/pip-clients-roles-go/version1"
)

type rolesMemoryClientV1Test struct {
	client  *version1.RolesMemoryClientV1
	fixture *RolesClientFixtureV1
}

func newRolesMemoryClientV1Test() *rolesMemoryClientV1Test {
	return &rolesMemoryClientV1Test{}
}

func (c *rolesMemoryClientV1Test) setup(t *testing.T) *RolesClientFixtureV1 {

	c.client = version1.NewRolesMemoryClientV1()
	c.fixture = NewRolesClientFixtureV1(c.client)
	return c.fixture
}

func (c *rolesMemoryClientV1Test) teardown(t *testing.T) {
	c.client = nil
	c.fixture = nil
}

func TestMemoryGetAndSetRoles(t *testing.T) {
	c := newRolesMemoryClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestGetAndSetRoles(t)
}

func TestMemoryGrantAndRevokeRoles(t *testing.T) {
	c := newRolesMemoryClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestGrantAndRevokeRoles(t)
}

func TestMemoryAuthorize(t *testing.T) {
	c := newRolesMemoryClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestAuthorize(t)
}
