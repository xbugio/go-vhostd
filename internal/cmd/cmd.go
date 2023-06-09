package cmd

import (
	"io/fs"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/xbugio/go-vhostd/internal/app"
)

type Cmd struct {
	cobra.Command

	config string
	htmlFs fs.FS
}

func NewCmd(htmlFs fs.FS) *Cmd {
	c := &Cmd{
		Command: cobra.Command{
			Use:   filepath.Base(os.Args[0]),
			Short: "vhost management system",
			Args:  cobra.OnlyValidArgs,
		},
		htmlFs: htmlFs,
	}

	c.Flags().StringVarP(&c.config, "config", "c", "config.yaml", "config file path")
	c.MarkFlagFilename("config")

	c.Command.Run = c.Run
	return c
}

func (c *Cmd) Run(cmd *cobra.Command, args []string) {
	if err := app.NewApp(c.config, c.htmlFs).Run(); err != nil {
		c.PrintErrln(err)
		os.Exit(1)
	}
}
