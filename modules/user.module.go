package modules

type User struct{}

func (ref User) Route() {
	handler := UserHandler{}
	handler.CreateClient("username", "password", "role_code", "name", "email")
	handler.Paginate(1, 10)
	handler.List("role_code")
	handler.UpdateClient("username", "body")
	handler.BlockClient("username")
}

type UserHandler struct{}

func (ref UserHandler) CreateClient(username string, password string, role_code string, name string, email string) bool {
	return true
}

func (ref UserHandler) Paginate(page int, limit int) bool {
	return true
}
func (ref UserHandler) List(role_code string) bool {
	return true
}

func (ref UserHandler) UpdateClient(username string, body interface{}) bool {
	return true
}

func (ref UserHandler) BlockClient(username string) bool {
	return true
}
