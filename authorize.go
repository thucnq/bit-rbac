package rbac

import "fmt"

type Authorize struct {
	WorkspaceID int64
	Permissions int64
}

func (a *Authorize) AddPermission(p Permission) {
	a.Permissions |= p.Val
}
func (a *Authorize) RebuildAfterSystemAddAction(lenRoles, lenActions int) {
	var filter int64 = 1
	var newPers int64 = 0
	for i := 0; i < lenActions-2; i++ {
		filter = filter<<1 + 1
	}
	fmt.Printf("filter %064b - lenActions %v\n", filter, lenActions)
	for i := 0; i < lenRoles*lenActions; i += lenActions {
		newPers |= (filter & a.Permissions) << (i / lenActions)
		filter <<= lenActions - 1
	}
	a.Permissions = newPers
}
func (a *Authorize) String() string {
	return fmt.Sprintf("Authorize {\n\tWorkspaceID: %v\n\tPermissionsInt: %v\n\tPermissionsBinary: %064b\n}\n", a.WorkspaceID, a.Permissions, a.Permissions)
}
