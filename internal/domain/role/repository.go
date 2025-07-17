package role

type IRoleRepository interface {
	Count() (int, error)
	Create(role *Role) error
	RoleExists(name string) bool
	Find(id int) (interface{}, error)
	Update(uuid string, role *Role) error
	FindByName(name string) (RoleResponse, error)
}
