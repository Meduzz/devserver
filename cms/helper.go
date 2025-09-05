package cms

func fail(err error) *ErrorDTO {
	return &ErrorDTO{
		Message: err.Error(),
	}
}
