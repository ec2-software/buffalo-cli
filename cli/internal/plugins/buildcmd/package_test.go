package buildcmd

import (
	"context"
	"os"
	"testing"

	"github.com/gobuffalo/buffalo-cli/plugins"
	"github.com/stretchr/testify/require"
)

func Test_BuildCmd_Package(t *testing.T) {
	r := require.New(t)

	ref := newRef(t, "ref")
	defer ref.Close()
	os.Chdir(ref.Root)

	pkg := &packager{
		files: []string{"A"},
	}
	pf := &packFiler{
		files: []string{"B"},
	}

	plugs := plugins.Plugins{
		pkg,
		pf,
	}

	bc := &BuildCmd{}
	bc.WithPlugins(plugs.ScopedPlugins)

	ctx := WithBuilderContext(context.Background(), nil)
	err := bc.Main(ctx, nil)
	r.NoError(err)

	r.Len(pkg.files, 2)
	r.Equal([]string{"A", "B"}, pkg.files)
}
