package rbac

type Resources []string

func (r *Resources) Init() {
	*r = []string{"assigment", "shift"}
}
func (r *Resources) Len() int {
	return len(*r)
}
func (r *Resources) Add(resource string) {
	*r = append(*r, resource)
}
