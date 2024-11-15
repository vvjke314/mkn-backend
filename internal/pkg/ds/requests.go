package ds

type ChangeEmailRequest struct {
	Email string `json:"email" binding:"required"`
}

type CreateProjectRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Color       string `json:"color" binding:"required"`
}

type UpdateProjectRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Color       string `json:"color"`
}

type CreateSectionRequest struct {
	Title string `json:"title" binding:"required"`
	Color string `json:"color" binding:"required"`
}

type UpdateSectionRequest struct {
	Title string `json:"title"`
	Color string `json:"color"`
}

type ResendNotificationRequest struct {
	Deadline string `json:"deadline"`
}

type CreateNotificationRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Deadline    string `json:"deadline" binding:"required"`
}

type UpdateNotificationRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	DeadLine    string `json:"deadline"`
}
