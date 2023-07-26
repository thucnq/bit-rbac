package rbac

import (
	"fmt"
	"strings"
)

func GeneratePermissions(roles Roles, actions Actions, resources Resources) map[string][]Permission {

	permissions := make(map[string][]Permission)

	lenActions := actions.Len()

	// generate permission from roles, actions, resources
	for _, resource := range resources {
		pers := make([]Permission, 0)
		for roleIdx, role := range roles {
			for actionIdx, action := range actions {
				p := roleIdx*lenActions + actionIdx
				pers = append(pers, Permission{
					Position: p,
					Code:     strings.Join([]string{role, action, resource}, "."),
					Val:      1 << (p),
				})
			}
		}
		permissions[resource] = pers
	}

	fmt.Println("permissions after gen:")
	for resource, _ := range permissions {
		for _, permission := range permissions[resource] {
			fmt.Println(permission.String())
		}
	}

	return permissions
}

func GetPermissionByCode(pers map[string][]Permission, code string) *Permission {
	cParts := strings.Split(code, ".")
	if ps, ok := pers[cParts[2]]; ok {
		for _, item := range ps {
			if item.Code == code {
				return &item
			}
		}
	}
	return nil
}

func Rbac() bool {
	roles := Roles{}
	roles.Init()
	actions := Actions{}
	actions.Init()
	resources := Resources{}
	resources.Init()
	pm := GeneratePermissions(roles, actions, resources)

	p1 := GetPermissionByCode(pm, "member.update.shift")
	p2 := GetPermissionByCode(pm, "manager.delete.shift")
	p3 := GetPermissionByCode(pm, "admin.read.shift")

	authorize := Authorize{
		WorkspaceID: 123456,
	}

	authorize.AddPermission(*p1)
	authorize.AddPermission(*p2)
	authorize.AddPermission(*p3)
	fmt.Println("authorize after add permissions:")
	fmt.Println(authorize.String())

	//roles.Add("moderate")
	//fmt.Println(roles)
	//
	//fmt.Println("Permissions after add role:")
	//pm = GeneratePermissions(roles, actions, resources)

	actions.Add("download")
	fmt.Println(actions)

	fmt.Println("Permissions after add action:")
	pm = GeneratePermissions(roles, actions, resources)

	targetAuthorize := Authorize{
		WorkspaceID: 123456,
	}

	p1 = GetPermissionByCode(pm, "member.update.shift")
	p2 = GetPermissionByCode(pm, "manager.delete.shift")
	p3 = GetPermissionByCode(pm, "admin.read.shift")
	targetAuthorize.AddPermission(*p1)
	targetAuthorize.AddPermission(*p2)
	targetAuthorize.AddPermission(*p3)
	fmt.Println("targetAuthorize after add permissions:")
	fmt.Println(targetAuthorize.String())

	authorize.RebuildAfterSystemAddAction(roles.Len(), actions.Len())
	fmt.Println("authorize after add RebuildAfterSystemAddAction:")
	fmt.Println(authorize.String())

	return false
}
