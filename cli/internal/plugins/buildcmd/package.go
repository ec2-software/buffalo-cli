package buildcmd

import (
	"context"

	"github.com/gobuffalo/buffalo-cli/internal/plugins"
	"github.com/gobuffalo/here"
)

func (bc *BuildCmd) pack(ctx context.Context, info here.Info, plugs []plugins.Plugin) error {
	var files []string
	for _, p := range plugs {
		pkg, ok := p.(PackFiler)
		if !ok {
			continue
		}
		res, err := pkg.PackageFiles(ctx, info.Dir)
		if err != nil {
			return err
		}
		files = append(files, res...)
	}

	for _, p := range plugs {
		pkg, ok := p.(Packager)
		if !ok {
			continue
		}
		if err := pkg.Package(ctx, info.Dir, files); err != nil {
			return err
		}
	}
	return nil
}
