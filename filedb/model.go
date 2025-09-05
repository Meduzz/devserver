package filedb

type (
	DB interface {
		Create(map[string]any) (map[string]any, error)
		Update(string, map[string]any) (map[string]any, error)
		Delete(string) error
		List() ([]map[string]any, error) // TODO filter
		Save() error
	}
)
