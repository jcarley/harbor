package actions

import "github.com/gobuffalo/buffalo"

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	return c.Render(200, r.JSON(map[string]string{"message": "Welcome to Buffalo!"}))
}

func CloneRepoHandler(c buffalo.Context) error {

	cmd := NewCloneCommand(repoUrl, localPath)
	err := cmd.Run()
	if err != nil {
		return c.Error(400, err)
	}

	return nil
}
