package version1

import (
	"github.com/pip-services3-go/pip-services3-commons-go/data"
)

type IRolesClientV1 interface {
	GetRolesByFilter(correlationId string, filter *data.FilterParams,
		paging *data.PagingParams) (result *data.DataPage, err error)

	GetRolesById(correlationId string, userId string) (result []string, err error)

	SetRoles(correlationId string, userId string, roles []string) (result []string, err error)

	GrantRoles(correlationId string, userId string, roles []string) (result []string, err error)

	RevokeRoles(correlationId string, userId string, roles []string) (result []string, err error)

	Authorize(correlationId string, userId string, roles []string) (result bool, err error)
}
