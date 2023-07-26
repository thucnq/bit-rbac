package rbac

type Actions []string

func (a *Actions) Init() {
	*a = []string{"create", "read", "update", "delete"}
}
func (a *Actions) Len() int {
	return len(*a)
}
func (a *Actions) Add(action string) {
	*a = append(*a, action)
}
