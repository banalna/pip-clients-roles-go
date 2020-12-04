package test_version1

import (
	"testing"

	"github.com/pip-services-users/pip-clients-roles-go/version1"
	"github.com/pip-services3-go/pip-services3-commons-go/data"
	"github.com/stretchr/testify/assert"
)

type RolesClientFixtureV1 struct {
	Client version1.IRolesClientV1
}

var ROLES = []string{"Role 1", "Role 2", "Role 3"}

func NewRolesClientFixtureV1(client version1.IRolesClientV1) *RolesClientFixtureV1 {
	return &RolesClientFixtureV1{
		Client: client,
	}
}

func (c *RolesClientFixtureV1) clear() {
	page, _ := c.Client.GetRolesByFilter("", nil, nil)
	if page != nil {
		for _, v := range page.Data {
			roles := v.(*version1.UserRolesV1)
			c.Client.RevokeRoles("", roles.Id, roles.Roles)
		}
	}
}

func (c *RolesClientFixtureV1) TestGetAndSetRoles(t *testing.T) {
	c.clear()
	defer c.clear()

	// Update party roles
	roles, err := c.Client.SetRoles("", "1", ROLES)
	assert.Nil(t, err)

	assert.True(t, len(roles) == 3)

	// Read and check party roles
	roles, err = c.Client.GetRolesById("", "1")
	assert.Nil(t, err)

	assert.True(t, len(roles) == 3)

	// Get roles by filter
	page, err1 := c.Client.GetRolesByFilter("", data.NewFilterParamsFromTuples("roles", ROLES), nil)
	assert.Nil(t, err1)

	assert.NotNil(t, page)
	assert.True(t, len(page.Data) >= 1)
}

func (c *RolesClientFixtureV1) TestGrantAndRevokeRoles(t *testing.T) {
	c.clear()
	defer c.clear()

	// Grant roles first time
	roles, err := c.Client.GrantRoles("", "1", []string{"Role 1"})
	assert.Nil(t, err)

	assert.Len(t, roles, 1)
	assert.Contains(t, roles, "Role 1")

	// Grant roles second time
	roles, err = c.Client.GrantRoles("", "1", []string{"Role 1", "Role 2", "Role 3"})
	assert.Nil(t, err)

	assert.Len(t, roles, 3)
	assert.Contains(t, roles, "Role 1")
	assert.Contains(t, roles, "Role 2")
	assert.Contains(t, roles, "Role 3")

	// Revoke roles first time
	roles, err = c.Client.RevokeRoles("", "1", []string{"Role 1"})
	assert.Nil(t, err)

	assert.Len(t, roles, 2)
	assert.Contains(t, roles, "Role 2")
	assert.Contains(t, roles, "Role 3")

	// Get roles
	roles, err = c.Client.GetRolesById("", "1")
	assert.Nil(t, err)

	assert.True(t, len(roles) == 2)
	assert.Contains(t, roles, "Role 2")
	assert.Contains(t, roles, "Role 3")

	// Revoke roles second time
	roles, err = c.Client.RevokeRoles("", "1", []string{"Role 1", "Role 2"})
	assert.Nil(t, err)

	assert.Len(t, roles, 1)
	assert.Contains(t, roles, "Role 3")
}

func (c *RolesClientFixtureV1) TestAuthorize(t *testing.T) {
	c.clear()
	defer c.clear()

	// Grant roles
	roles, err := c.Client.GrantRoles("", "1", []string{"Role 1", "Role 2"})
	assert.Nil(t, err)

	assert.Len(t, roles, 2)

	// Authorize positively
	auth, err1 := c.Client.Authorize("", "1", []string{"Role 1"})
	assert.Nil(t, err1)

	assert.True(t, auth)

	// Authorize negatively
	auth, err1 = c.Client.Authorize("", "1", []string{"Role 2", "Role 3"})
	assert.Nil(t, err1)

	assert.False(t, auth)
}
