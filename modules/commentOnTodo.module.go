package modules

type CommentOnTodo struct {
}

func (ref CommentOnTodo) Route() {
	handler := CommentOnTodoHandler{}

	handler.Create(1, "comment")

	handler.List(1)

	handler.Update(1, "body")

	handler.Delete(1)
}

type CommentOnTodoHandler struct{}

func (ref CommentOnTodoHandler) Create(id int, comment string) bool {
	return true
}

func (ref CommentOnTodoHandler) List(id int) bool {
	return true
}

func (ref CommentOnTodoHandler) Update(id int, body interface{}) bool {
	return true
}

func (ref CommentOnTodoHandler) Delete(id int) bool {
	return true
}
