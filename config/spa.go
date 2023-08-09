package config

type (
	// Define a spa page
	Spa struct {
		Page *Component
	}

	Component struct {
		Component string         `json:"component"`
		Settings  map[string]any `json:"settings"`
		Children  []*Component   `json:"children"`
	}
)
