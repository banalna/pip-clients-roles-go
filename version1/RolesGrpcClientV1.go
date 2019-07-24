package version1

import (
	"github.com/pip-services-users/pip-clients-roles-go/protos"
	"github.com/pip-services3-go/pip-services3-commons-go/data"
	"github.com/pip-services3-go/pip-services3-grpc-go/clients"
)

type RoleGrpcClientV1 struct {
	clients.GrpcClient
}

func NewRoleGrpcClientV1() *RoleGrpcClientV1 {
	return &RoleGrpcClientV1{
		GrpcClient: *clients.NewGrpcClient("roles_v1.Roles"),
	}
}

func (c *RoleGrpcClientV1) GetRolesByFilter(correlationId string, filter *data.FilterParams,
	paging *data.PagingParams) (result *data.DataPage, err error) {
	req := &protos.RolesPageRequest{
		CorrelationId: correlationId,
	}
	if filter != nil {
		req.Filter = filter.Value()
	}
	if paging != nil {
		req.Paging = &protos.PagingParams{
			Skip:  paging.GetSkip(0),
			Take:  (int32)(paging.GetTake(100)),
			Total: paging.Total,
		}
	}

	reply := new(protos.RolesPageReply)
	err = c.Call("get_roles_by_filter", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toUserRolesPage(reply.Page)

	return result, nil
}

func (c *RoleGrpcClientV1) GetRolesById(correlationId string, userId string) (result []string, err error) {
	req := &protos.RoleIdRequest{
		CorrelationId: correlationId,
		UserId:        userId,
	}

	reply := new(protos.RolesReply)
	err = c.Call("get_roles_by_id", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = reply.Roles

	return result, nil
}

func (c *RoleGrpcClientV1) SetRoles(correlationId string, userId string, roles []string) (result []string, err error) {
	req := &protos.RolesRequest{
		CorrelationId: correlationId,
		UserId:        userId,
		Roles:         roles,
	}

	reply := new(protos.RolesReply)
	err = c.Call("set_roles", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = reply.Roles

	return result, nil
}

func (c *RoleGrpcClientV1) GrantRoles(correlationId string, userId string, roles []string) (result []string, err error) {
	req := &protos.RolesRequest{
		CorrelationId: correlationId,
		UserId:        userId,
		Roles:         roles,
	}

	reply := new(protos.RolesReply)
	err = c.Call("grant_roles", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = reply.Roles

	return result, nil
}

func (c *RoleGrpcClientV1) RevokeRoles(correlationId string, userId string, roles []string) (result []string, err error) {
	req := &protos.RolesRequest{
		CorrelationId: correlationId,
		UserId:        userId,
		Roles:         roles,
	}

	reply := new(protos.RolesReply)
	err = c.Call("revoke_roles", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = reply.Roles

	return result, nil
}

func (c *RoleGrpcClientV1) Authorize(correlationId string, userId string, roles []string) (result bool, err error) {
	req := &protos.RolesRequest{
		CorrelationId: correlationId,
		UserId:        userId,
		Roles:         roles,
	}

	reply := new(protos.AuthorizeReply)
	err = c.Call("authorize", correlationId, req, reply)
	if err != nil {
		return false, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return false, err
	}

	result = reply.Authorized

	return result, nil
}
