package modules

type Todo struct{}

func (ref Todo) Route() {
	handler := TodoHandler{}
	handler.Create("name", "description", true, "due_date", "create_at")
	handler.Paginate(1, 10)
	handler.GetDetail(1)
	handler.Update(1, "body")
	handler.Delete(1)
}

type TodoHandler struct{}

func (ref TodoHandler) Create(name string, description string, status bool, due_date string, create_at string) bool {
	return true
}

func (ref TodoHandler) Paginate(page int, limit int) bool {
	return true
}

func (ref TodoHandler) GetDetail(id int) bool {
	return true
}

func (ref TodoHandler) Update(id int, body interface{}) bool {
	return true
}

func (ref TodoHandler) Delete(id int) bool {
	return true
}