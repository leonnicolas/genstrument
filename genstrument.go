package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"

	"github.com/oklog/run"
	"github.com/spf13/cobra"

	"github.com/leonnicolas/genstrument/generator"
)

var root = &cobra.Command{
	Use:   "genstrument",
	Short: "instrument an interface with prometheus",
	Long:  "generate code for a provided interface name, that provided an instrumented wrapper around the original interface",
	RunE: func(cmd *cobra.Command, _ []string) (err error) {
		c := &generator.Config{}
		c.FilePath, err = cmd.PersistentFlags().GetString("file-path")
		if err != nil {
			return err
		}
		c.Pattern, err = cmd.PersistentFlags().GetString("pattern")
		if err != nil {
			return err
		}
		c.MetricHelp, err = cmd.PersistentFlags().GetString("metric-help")
		if err != nil {
			return err
		}
		c.MetricName, err = cmd.PersistentFlags().GetString("metric-name")
		if err != nil {
			return err
		}

		// We must save the output file in a buffer to avoid parsing this file when we load the go package.
		buf := bytes.NewBuffer(nil)
		c.Out = buf

		g := run.Group{}
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		g.Add(func() error { return c.Generate(ctx) }, func(err error) { cancel() })
		g.Add(run.SignalHandler(ctx, os.Interrupt))

		if err = g.Run(); err != nil {
			return err
		}
		out, err := cmd.PersistentFlags().GetString("out")
		if err != nil {
			return err
		}
		w := os.Stdout
		if out != "-" {
			w, err = os.OpenFile(out, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0600)
			if err != nil {
				return fmt.Errorf("couldn't open %s for writing: %w", out, err)
			}
		}
		_, err = io.Copy(w, buf)
		return err
	},
}

func init() {
	root.PersistentFlags().StringP("pattern", "p", "", "pattern of Interface name to match")
	root.MarkPersistentFlagRequired("pattern")

	root.PersistentFlags().StringP("file-path", "f", "", "path of the go file that contains the interface to be instrumented")
	root.MarkPersistentFlagRequired("file-path")

	root.PersistentFlags().String("metric-name", "", "metric name")
	root.MarkPersistentFlagRequired("metric-name")

	root.PersistentFlags().String("metric-help", "", "metric help text")
	root.MarkPersistentFlagRequired("metric-help")

	root.PersistentFlags().StringP("out", "o", "-", "metric help text")
}

func main() {
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
