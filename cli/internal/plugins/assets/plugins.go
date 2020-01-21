package assets

import (
	"github.com/gobuffalo/buffalo-cli/cli/internal/plugins/assets/build"
	"github.com/gobuffalo/buffalo-cli/cli/internal/plugins/assets/develop"
	"github.com/gobuffalo/buffalo-cli/plugins"
)

// Plugins returns all of the plugins available in this package.
// All plugins use zero values.
func Plugins() []plugins.Plugin {
	return []plugins.Plugin{
		&build.Builder{},
		&develop.Developer{},
	}
}