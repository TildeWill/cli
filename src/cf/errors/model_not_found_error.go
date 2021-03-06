package errors

type ModelNotFoundError struct {
	ModelType string
	ModelName string
}

func NewModelNotFoundError(modelType, name string) ModelNotFoundError {
	return ModelNotFoundError{
		ModelType: modelType,
		ModelName: name,
	}
}

func (err ModelNotFoundError) Error() string {
	return err.ModelType + " " + err.ModelName + " not found"
}

func (err ModelNotFoundError) ErrorCode() string {
	return "model-not-found"
}
