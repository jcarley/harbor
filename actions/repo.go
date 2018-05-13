package actions

import "github.com/gobuffalo/buffalo"

// RepoClone default implementation.
func RepoClone(c buffalo.Context) error {
	return c.Render(200, r.JSON(map[string]string{"message": "Welcome to Buffalo!"}))
}
