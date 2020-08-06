package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "cerulean",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
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
				//println(err.Error())
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

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cerulean.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		viper.AddConfigPath(home)
		viper.SetConfigName(".cerulean")
	}
	viper.AutomaticEnv() // read in environment variables that match
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
