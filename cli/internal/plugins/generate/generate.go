package generate

import "github.com/gobuffalo/buffalo-cli/plugins"

func Plugins() []plugins.Plugin {
	return []plugins.Plugin{
		&Cmd{},
	}
}
