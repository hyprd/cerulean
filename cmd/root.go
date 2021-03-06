package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			println("Fetching mp3...")
			root := "youtube-dl"
			r_ignore_error := "-i"
			r_metadata := "--add-metadata"
			r_extract_audio := "-x"
			r_audio_format := "--audio-format"
			r_audio_type := "mp3"
			r_manifest := "--youtube-skip-dash-manifest"

			track := args[0]
			run := exec.Command(root, r_ignore_error, r_metadata, r_extract_audio, r_audio_format, r_audio_type, r_manifest, track)
			stdout, err := run.Output()
			if err != nil {
				println("Failure - check your YouTube URL (everything after ?v=)")
				return
			}
			println(string(stdout))
		} else {
			println("Specify a valid YouTube video URL (everything after `v?=`)")
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {}
