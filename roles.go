package rbac

type Roles []string

func (r *Roles) Init() {
	*r = []string{"admin", "manager", "member"}
}
func (r *Roles) Len() int {
	return len(*r)
}
func (r *Roles) Add(role string) {
	*r = append(*r, role)
}
