package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/jcarley/harbor/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
