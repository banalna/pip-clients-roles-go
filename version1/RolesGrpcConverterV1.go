package version1

import (
	"encoding/json"

	"github.com/pip-services-users/pip-clients-roles-go/protos"
	"github.com/pip-services3-go/pip-services3-commons-go/convert"
	"github.com/pip-services3-go/pip-services3-commons-go/data"
	"github.com/pip-services3-go/pip-services3-commons-go/errors"
)

func fromError(err error) *protos.ErrorDescription {
	if err == nil {
		return nil
	}

	desc := errors.ErrorDescriptionFactory.Create(err)
	obj := &protos.ErrorDescription{
		Type:          desc.Type,
		Category:      desc.Category,
		Code:          desc.Code,
		CorrelationId: desc.CorrelationId,
		Status:        convert.StringConverter.ToString(desc.Status),
		Message:       desc.Message,
		Cause:         desc.Cause,
		StackTrace:    desc.StackTrace,
		Details:       fromMap(desc.Details),
	}

	return obj
}

func toError(obj *protos.ErrorDescription) error {
	if obj == nil || (obj.Category == "" && obj.Message == "") {
		return nil
	}

	description := &errors.ErrorDescription{
		Type:          obj.Type,
		Category:      obj.Category,
		Code:          obj.Code,
		CorrelationId: obj.CorrelationId,
		Status:        convert.IntegerConverter.ToInteger(obj.Status),
		Message:       obj.Message,
		Cause:         obj.Cause,
		StackTrace:    obj.StackTrace,
		Details:       toMap(obj.Details),
	}

	return errors.ApplicationErrorFactory.Create(description)
}

func fromMap(val map[string]interface{}) map[string]string {
	r := map[string]string{}

	for k, v := range val {
		r[k] = convert.ToString(v)
	}

	return r
}

func toMap(val map[string]string) map[string]interface{} {
	var r map[string]interface{}

	for k, v := range val {
		r[k] = v
	}

	return r
}

func toJson(value interface{}) string {
	if value == nil {
		return ""
	}

	b, err := json.Marshal(value)
	if err != nil {
		return ""
	}
	return string(b[:])
}

func fromJson(value string) interface{} {
	if value == "" {
		return nil
	}

	var m interface{}
	json.Unmarshal([]byte(value), &m)
	return m
}

func fromUserRoles(role *UserRolesV1) *protos.UserRoles {
	if role == nil {
		return nil
	}

	obj := &protos.UserRoles{
		Id:         role.Id,
		Roles:      role.Roles,
		UpdateTime: convert.StringConverter.ToString(role.UpdateTime),
	}

	return obj
}

func toUserRoles(obj *protos.UserRoles) *UserRolesV1 {
	if obj == nil {
		return nil
	}

	role := &UserRolesV1{
		Id:         obj.Id,
		Roles:      obj.Roles,
		UpdateTime: convert.DateTimeConverter.ToDateTime(obj.UpdateTime),
	}

	return role
}

func fromUserRolesPage(page *data.DataPage) *protos.UserRolesPage {
	if page == nil {
		return nil
	}

	obj := &protos.UserRolesPage{
		Total: *page.Total,
		Data:  make([]*protos.UserRoles, len(page.Data)),
	}

	for i, v := range page.Data {
		role := v.(*UserRolesV1)
		obj.Data[i] = fromUserRoles(role)
	}

	return obj
}

func toUserRolesPage(obj *protos.UserRolesPage) *data.DataPage {
	if obj == nil {
		return nil
	}

	roles := make([]interface{}, len(obj.Data))

	for i, v := range obj.Data {
		roles[i] = toUserRoles(v)
	}

	page := data.NewDataPage(&obj.Total, roles)

	return page
}
