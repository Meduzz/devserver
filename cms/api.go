package cms

type (
	InitRequest struct {
		Name string `json:"name"`
	}

	ErrorDTO struct {
		Message string `json:"message"`
	}
)
