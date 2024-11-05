package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/tools/cover"
	"os"
)

func Main() {
	var (
		files []string
		out   string
	)
	var cmd = &cobra.Command{
		Use: "merge",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			var merged []*cover.Profile

			for _, file := range files {
				var profiles []*cover.Profile
				if profiles, err = cover.ParseProfiles(file); err != nil {
					return fmt.Errorf("failed to parse profiles: %v", err)
				}

				for _, p := range profiles {
					if merged, err = addProfile(merged, p); err != nil {
						return err
					}
				}
			}

			var f *os.File
			if f, err = os.Create(out); err != nil {
				return err
			}

			dumpProfiles(merged, f)
			return nil
		},
	}

	_ = cmd.MarkFlagRequired("out")
	_ = cmd.MarkFlagRequired("file")

	cmd.Flags().StringArrayVarP(&files, "file", "f", []string{}, "")
	cmd.Flags().StringVarP(&out, "out", "o", "", "")

	var ctx = context.Background()
	if err := cmd.ExecuteContext(ctx); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
