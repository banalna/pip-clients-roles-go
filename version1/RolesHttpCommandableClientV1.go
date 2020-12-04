package version1

import (
	"reflect"

	"github.com/pip-services3-go/pip-services3-commons-go/data"
	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	cclients "github.com/pip-services3-go/pip-services3-rpc-go/clients"
)

type RolesHttpCommandableClientV1 struct {
	*cclients.CommandableHttpClient
	stringArrayType reflect.Type
	dataPageType    reflect.Type
	boolType        reflect.Type
}

func NewRolesHttpCommandableClientV1() *RolesHttpCommandableClientV1 {
	c := &RolesHttpCommandableClientV1{
		CommandableHttpClient: cclients.NewCommandableHttpClient("v1/roles"),
		dataPageType:          reflect.TypeOf(&cdata.DataPage{}),
		stringArrayType:       reflect.TypeOf(make([]string, 0)),
		boolType:              reflect.TypeOf(true),
	}
	return c
}

func (c *RolesHttpCommandableClientV1) GetRolesByFilter(correlationId string, filter *data.FilterParams,
	paging *data.PagingParams) (result *cdata.DataPage, err error) {

	params := cdata.NewAnyValueMapFromTuples(
		"filter", filter,
		"paging", paging,
	)

	res, err := c.CallCommand(c.dataPageType, "get_roles_by_filter", correlationId, params)
	if err != nil {
		return nil, err
	}

	result, _ = res.(*cdata.DataPage)
	return result, nil
}

func (c *RolesHttpCommandableClientV1) GetRolesById(correlationId string, userId string) (result []string, err error) {
	params := cdata.NewAnyValueMapFromTuples(
		"user_id", userId,
	)

	res, err := c.CallCommand(c.stringArrayType, "get_roles_by_id", correlationId, params)
	if err != nil {
		return nil, err
	}

	result, _ = res.([]string)
	return result, nil
}

func (c *RolesHttpCommandableClientV1) SetRoles(correlationId string, userId string, roles []string) (result []string, err error) {

	params := cdata.NewAnyValueMapFromTuples(
		"user_id", userId,
		"roles", roles,
	)

	res, err := c.CallCommand(c.stringArrayType, "set_roles", correlationId, params)
	if err != nil {
		return nil, err
	}

	result, _ = res.([]string)
	return result, nil
}

func (c *RolesHttpCommandableClientV1) GrantRoles(correlationId string, userId string, roles []string) (result []string, err error) {

	params := cdata.NewAnyValueMapFromTuples(
		"user_id", userId,
		"roles", roles,
	)

	res, err := c.CallCommand(c.stringArrayType, "grant_roles", correlationId, params)
	if err != nil {
		return nil, err
	}

	result, _ = res.([]string)
	return result, nil
}

func (c *RolesHttpCommandableClientV1) RevokeRoles(correlationId string, userId string, roles []string) (result []string, err error) {

	params := cdata.NewAnyValueMapFromTuples(
		"user_id", userId,
		"roles", roles,
	)

	res, err := c.CallCommand(c.stringArrayType, "revoke_roles", correlationId, params)
	if err != nil {
		return nil, err
	}

	result, _ = res.([]string)
	return result, nil
}

func (c *RolesHttpCommandableClientV1) Authorize(correlationId string, userId string, roles []string) (result bool, err error) {

	params := cdata.NewAnyValueMapFromTuples(
		"user_id", userId,
		"roles", roles,
	)

	res, err := c.CallCommand(c.boolType, "authorize", correlationId, params)
	if err != nil {
		return false, err
	}

	result, _ = res.(bool)
	return result, nil
}
