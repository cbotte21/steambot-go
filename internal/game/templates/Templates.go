package templates

const (
	Default string = "myexampletemplate"
)

// GetTemplate matches a string pattern, returning desired template.
func GetTemplate(name string) string {
	switch name {
	default:
		return Default
	}
}
