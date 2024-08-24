package cmd

import (
	"errors"
	"fmt"
	"github.com/motivewc/wowpatch/internal/binary"
	"github.com/motivewc/wowpatch/internal/patterns"
	"github.com/motivewc/wowpatch/internal/platform"
	"os"

	"github.com/spf13/cobra"
)

type RootOptions struct {
	StripCodesignAttributes bool
	WarcraftExeLocation     string
	OutputFile              string
}

var (
	options RootOptions
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wowpatch",
	Short: "modifies WoW binary to enable connecting to private servers",
	Long: `This application takes as input a retail World of Warcraft client and will generate a modified executable
from it by using binary patching. The resulting executable can be run safely and connect to private servers.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		data, err := os.ReadFile(options.WarcraftExeLocation)
		if err != nil {
			return fmt.Errorf("unable to read the WoW executable file: %w", err)
		}
		binary.Patch(&data, patterns.PortalPattern, patterns.PortalReplace)
		binary.Patch(&data, patterns.ConnectToModulusPattern, patterns.TrinityRsaModulus)
		binary.Patch(&data, patterns.CryptoEdPublicKeyPattern, patterns.TrinityCryptoEdPublicKey)

		if err = os.Remove(options.OutputFile); !errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("unable to remove file that already exists at %v: %w", options.OutputFile, err)
		}

		err = os.WriteFile(options.OutputFile, data, 0777)
		if err != nil {
			return fmt.Errorf("unable to write to file %v: %w", options.OutputFile, err)
		}

		fmt.Printf("Client has been successfully patched and saved to %v\n", options.OutputFile)

		return nil

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println("An error has occurred, the client has not been patched.")
		fmt.Println()
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&options.WarcraftExeLocation, "warcraft-exe", platform.FindWarcraftClientExecutable(), "the location of the WoW executable")
	rootCmd.PersistentFlags().BoolVar(&options.StripCodesignAttributes, "strip-binary-codesign", true, "removes macOS codesigning from resulting binary")
}
