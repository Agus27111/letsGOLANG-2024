package modules

type AssignmentTask struct{}

func (ref AssignmentTask) Route() {
	handler := AssignmentTaskHandler{}

	handler.Change(1, "body")
}

type AssignmentTaskHandler struct{}

func (ref AssignmentTaskHandler) Change(id int, body interface{}) bool {
	return true
}