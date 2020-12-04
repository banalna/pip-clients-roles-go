package version1

import (
	"strings"
	"time"

	"github.com/pip-services3-go/pip-services3-commons-go/data"
	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
)

type RolesMemoryClientV1 struct {
	roles []UserRolesV1
}

func NewRolesMemoryClientV1() *RolesMemoryClientV1 {

	c := RolesMemoryClientV1{
		roles: make([]UserRolesV1, 0),
	}
	return &c
}

func (c *RolesMemoryClientV1) GetRolesByFilter(correlationId string, filter *data.FilterParams,
	paging *data.PagingParams) (result *data.DataPage, err error) {

	var total int64 = (int64)(len(c.roles))
	items := make([]interface{}, 0)
	for _, v := range c.roles {
		item := v
		items = append(items, &item)
	}
	return cdata.NewDataPage(&total, items), nil

}

func (c *RolesMemoryClientV1) GetRolesById(correlationId string, userId string) (result []string, err error) {

	result = make([]string, 0)
	for _, v := range c.roles {
		if v.Id == userId {
			result = v.Roles
			break
		}
	}
	return result, nil

}

func (c *RolesMemoryClientV1) SetRoles(correlationId string, userId string, roles []string) (result []string, err error) {

	var userRoles UserRolesV1
	var index = -1
	for i, v := range c.roles {
		if v.Id == userId {
			index = i
			userRoles = v
			break
		}
	}

	if index >= 0 {
		c.roles[index].Roles = roles
		c.roles[index].UpdateTime = time.Now().UTC()
	} else {
		userRoles = *NewUserRolesV1(userId, roles)
		c.roles = append(c.roles, userRoles)
	}

	return roles, nil
}

func (c *RolesMemoryClientV1) GrantRoles(correlationId string, userId string, roles []string) (result []string, err error) {
	// If there are no roles then skip processing
	if len(roles) == 0 {
		return nil, nil
	}

	existingRoles, err := c.GetRolesById(
		correlationId,
		userId)

	if err != nil {
		return nil, err
	}

	newRoles := c.union(roles, existingRoles)

	return c.SetRoles(
		correlationId,
		userId,
		newRoles)

}

func (c *RolesMemoryClientV1) RevokeRoles(correlationId string, userId string, roles []string) (result []string, err error) {
	// If there are no roles then skip processing
	if len(roles) == 0 {
		return nil, nil
	}

	existingRoles, err := c.GetRolesById(
		correlationId,
		userId)

	if err != nil {
		return nil, err
	}

	newRoles := c.difference(existingRoles, roles)

	return c.SetRoles(
		correlationId,
		userId,
		newRoles,
	)

}

func (c *RolesMemoryClientV1) Authorize(correlationId string, userId string, roles []string) (result bool, err error) {
	// If there are no roles then skip processing
	if len(roles) == 0 {
		return false, nil
	}

	existingRoles, err := c.GetRolesById(
		correlationId,
		userId)

	if err != nil {
		return false, err
	}

	authorized := len(c.difference(roles, existingRoles)) == 0

	return authorized, nil

}

func (c *RolesMemoryClientV1) contains(array1 []string, array2 []string) bool {
	if array1 == nil || array2 == nil {
		return false
	}

	for i1 := 0; i1 < len(array1); i1++ {
		for i2 := 0; i2 < len(array2); i2++ {
			if array1[i1] == array2[i2] {
				return true
			}
		}
	}

	return false
}

func (c *RolesMemoryClientV1) composeFilter(filter *cdata.FilterParams) func(item UserRolesV1) bool {
	if filter == nil {
		filter = cdata.NewEmptyFilterParams()
	}

	id := filter.GetAsString("id")
	ids := filter.GetAsString("ids")
	exceptIds := filter.GetAsString("except_ids")
	roles := filter.GetAsString("roles")
	exceptRoles := filter.GetAsString("except_roles")

	// Process ids filter
	var idsValues []string
	if ids != "" {
		idsValues = strings.Split(ids, ",")
	}

	// Process except ids filter
	var exceptIdsValues []string
	if exceptIds != "" {
		exceptIdsValues = strings.Split(exceptIds, ",")
	}

	// Process roles filter
	var rolesValues []string
	if roles != "" {
		rolesValues = strings.Split(roles, ",")
	}

	// Process except roles filter
	var exceptRolesValues []string
	if exceptRoles != "" {
		exceptRolesValues = strings.Split(exceptRoles, ",")
	}

	return func(item UserRolesV1) bool {
		if id != "" && item.Id != id {
			return false
		}
		if len(idsValues) > 0 && strings.Index(ids, item.Id) < 0 {
			return false
		}
		if len(exceptIdsValues) > 0 && strings.Index(exceptIds, item.Id) >= 0 {
			return false
		}
		if len(rolesValues) > 0 && !c.contains(rolesValues, item.Roles) {
			return false
		}
		if len(exceptRolesValues) > 0 && c.contains(exceptRolesValues, item.Roles) {
			return false
		}
		return true
	}

}

func (c *RolesMemoryClientV1) union(a, b []string) []string {
	m := make(map[string]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; !ok {
			a = append(a, item)
		}
	}
	return a
}

// Set Difference: A - B
func (c *RolesMemoryClientV1) difference(a, b []string) (diff []string) {
	m := make(map[string]bool)

	for _, item := range b {
		m[item] = true
	}

	for _, item := range a {
		if _, ok := m[item]; !ok {
			diff = append(diff, item)
		}
	}
	return
}
