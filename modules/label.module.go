package modules

type Label struct{}

func (ref Label) Route() {
	handler := LabelHandler{}
	handler.Create("label", "green")
	handler.List()
	handler.GetDetail(1)
	handler.Update(1, "body")
	handler.Delete("label")
}

type LabelHandler struct{}

func (ref LabelHandler) Create(label string, color string) bool {
	return true
}

func (ref LabelHandler) List() bool {
	return true
}
func (ref LabelHandler) GetDetail(id int) bool {
	return true
}

func (ref LabelHandler) Update(id int, body interface{}) bool {
	return true
}

func (ref LabelHandler) Delete(label string) bool {
	return true
}