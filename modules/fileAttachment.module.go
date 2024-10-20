package modules

type FileAttachment struct{}

func (ref FileAttachment) Route() {

	handler := FileAttachmentHandler{}

	handler.Upload(1, "fileName", "create_at")

	handler.Delete(1)

	handler.List(1)

}

type FileAttachmentHandler struct{}

func (ref FileAttachmentHandler) Upload(task_id int, fileName string, create_at string) bool {
	return true
}

func (ref FileAttachmentHandler) Delete(task_id int) bool {
	return true
}

func (ref FileAttachmentHandler) List(task_id int) bool {
	return true
}
