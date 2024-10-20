package modules

type Auth struct{}

func (ref Auth) Route() {
	handler := AuthHandler{}
	handler.login("username", "password")
	handler.logout("username")
	handler.TokenValidation("token")
}

type AuthHandler struct{}

func (ref AuthHandler) login(username string, password string) bool {
	return true
}

func (ref AuthHandler) logout(username string) bool {
	return true
}

func (ref AuthHandler) TokenValidation(token string) bool {
	return true
}