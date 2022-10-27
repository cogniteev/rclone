// Package copyfiles provides the copyfiles command.
package copyfiles

import (
	"context"
	"github.com/rclone/rclone/fs/sync"
	"strings"

	"github.com/rclone/rclone/cmd"
	"github.com/rclone/rclone/fs/config/flags"
	"github.com/spf13/cobra"
)

var (
	filesList = ""
)

func init() {
	cmd.Root.AddCommand(commandDefinition)
	cmdFlags := commandDefinition.Flags()
	flags.StringVarP(cmdFlags, &filesList, "files-list", "", "", "Read list of source-file names from file without any processing of lines")
}

var commandDefinition = &cobra.Command{
	Use:   "copyfiles source:path dest:path --files-list files.txt",
	Short: `Copy files from source to dest, skipping identical files.`,
	// Note: "|" will be replaced by backticks below
	Long: strings.ReplaceAll(`
Copy the source to the destination.  Does not transfer files that are
identical on source and destination, testing by size and modification
time or MD5SUM.  Doesn't delete files from the destination. If you
want to also delete files from destination, to make it match source,
use the [sync](/commands/rclone_sync/) command instead.

Only files contained in files-list will be copied.
`, "|", "`"),
	Run: func(command *cobra.Command, args []string) {
		cmd.CheckArgs(2, 2, command, args)
		fsrc, _, fdst := cmd.NewFsSrcFileDst(args)
		cmd.Run(true, true, command, func() error {
			return sync.CopyFiles(context.Background(), fdst, fsrc, filesList)
		})
	},
}
