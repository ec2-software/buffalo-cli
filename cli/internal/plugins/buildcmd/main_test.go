package buildcmd

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/gobuffalo/buffalo-cli/plugins"
	"github.com/stretchr/testify/require"
)

func Test_BuildCmd_Main(t *testing.T) {
	r := require.New(t)

	ref := newRef(t, "ref")
	defer ref.Close()
	os.Chdir(ref.Root)

	bc := &BuildCmd{}

	exp := []string{"go", "build", "-o", "bin/coke"}
	var act []string
	ctx := context.Background()
	ctx = WithBuilderContext(ctx, func(cmd *exec.Cmd) error {
		act = make([]string, len(cmd.Args))
		copy(act, cmd.Args)
		return nil
	})

	var args []string
	err := bc.Main(ctx, args)
	r.NoError(err)
	r.Equal(exp, act)
}

func Test_BuildCmd_Main_SubCommand(t *testing.T) {
	r := require.New(t)

	ref := newRef(t, "ref")
	defer ref.Close()
	os.Chdir(ref.Root)

	p := &builder{name: "foo"}
	plugs := plugins.Plugins{p}

	bc := &BuildCmd{
		pluginsFn: plugs.ScopedPlugins,
	}

	ctx := WithBuilderContext(context.Background(), nil)
	args := []string{p.name, "a", "b", "c"}

	err := bc.Main(ctx, args)
	r.NoError(err)
	r.Equal([]string{"a", "b", "c"}, p.args)
}

func Test_BuildCmd_Main_SubCommand_err(t *testing.T) {
	r := require.New(t)

	ref := newRef(t, "ref")
	defer ref.Close()
	os.Chdir(ref.Root)

	p := &builder{name: "foo", err: fmt.Errorf("error")}
	plugs := plugins.Plugins{p}

	bc := &BuildCmd{
		pluginsFn: plugs.ScopedPlugins,
	}

	ctx := WithBuilderContext(context.Background(), nil)
	args := []string{p.name, "a", "b", "c"}

	err := bc.Main(ctx, args)
	r.Error(err)
}

func Test_BuildCmd_Main_ValidateTemplates(t *testing.T) {
	r := require.New(t)

	ref := newRef(t, "ref")
	defer ref.Close()
	os.Chdir(ref.Root)

	p := &templatesValidator{}
	plugs := plugins.Plugins{p}

	bc := &BuildCmd{
		pluginsFn: plugs.ScopedPlugins,
	}

	ctx := WithBuilderContext(context.Background(), nil)
	args := []string{}

	err := bc.Main(ctx, args)
	r.NoError(err)
	r.Equal(ref.Root, p.root)
}

func Test_BuildCmd_Main_ValidateTemplates_err(t *testing.T) {
	r := require.New(t)

	ref := newRef(t, "ref")
	defer ref.Close()
	os.Chdir(ref.Root)

	p := &templatesValidator{err: fmt.Errorf("error")}
	plugs := plugins.Plugins{p}

	bc := &BuildCmd{
		pluginsFn: plugs.ScopedPlugins,
	}

	ctx := WithBuilderContext(context.Background(), nil)
	args := []string{}

	err := bc.Main(ctx, args)
	r.Error(err)
}

func Test_BuildCmd_Main_BeforeBuilders(t *testing.T) {
	r := require.New(t)

	ref := newRef(t, "ref")
	defer ref.Close()
	os.Chdir(ref.Root)

	p := &beforeBuilder{}
	plugs := plugins.Plugins{p}

	bc := &BuildCmd{
		pluginsFn: plugs.ScopedPlugins,
	}

	ctx := WithBuilderContext(context.Background(), nil)
	var args []string

	err := bc.Main(ctx, args)
	r.NoError(err)
}

func Test_BuildCmd_Main_BeforeBuilders_err(t *testing.T) {
	r := require.New(t)

	ref := newRef(t, "ref")
	defer ref.Close()
	os.Chdir(ref.Root)

	p := &beforeBuilder{err: fmt.Errorf("error")}
	plugs := plugins.Plugins{p}

	bc := &BuildCmd{
		pluginsFn: plugs.ScopedPlugins,
	}

	ctx := WithBuilderContext(context.Background(), nil)
	var args []string

	err := bc.Main(ctx, args)
	r.Error(err)
}

func Test_BuildCmd_Main_AfterBuilders(t *testing.T) {
	r := require.New(t)

	ref := newRef(t, "ref")
	defer ref.Close()
	os.Chdir(ref.Root)

	p := &afterBuilder{}
	plugs := plugins.Plugins{p}

	bc := &BuildCmd{
		pluginsFn: plugs.ScopedPlugins,
	}

	ctx := WithBuilderContext(context.Background(), nil)
	var args []string

	err := bc.Main(ctx, args)
	r.NoError(err)
}

func Test_BuildCmd_Main_AfterBuilders_err(t *testing.T) {
	r := require.New(t)

	ref := newRef(t, "ref")
	defer ref.Close()
	os.Chdir(ref.Root)

	b := &beforeBuilder{err: fmt.Errorf("error")}
	a := &afterBuilder{}
	plugs := plugins.Plugins{a, b}

	bc := &BuildCmd{
		pluginsFn: plugs.ScopedPlugins,
	}

	ctx := WithBuilderContext(context.Background(), nil)
	var args []string

	err := bc.Main(ctx, args)
	r.Error(err)
	r.Equal(err, a.err)
}
