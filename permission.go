package rbac

import "fmt"

type Permission struct {
	ID       int64
	Position int
	Code     string
	Val      int64
}

func (p *Permission) String() string {
	return fmt.Sprintf("Permission {\n\tID: %v\n\tPosition: %v\n\tCode: %v\n\tValInt: %v\n\tValBinary: %064b\n}\n", p.ID, p.Position, p.Code, p.Val, p.Val)
}
