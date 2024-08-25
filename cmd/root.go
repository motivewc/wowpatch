package cmd

import (
	"errors"
	"fmt"
	"github.com/motivewc/wowpatch/internal/binary"
	"github.com/motivewc/wowpatch/internal/patterns"
	"github.com/motivewc/wowpatch/internal/platform"
	"github.com/motivewc/wowpatch/internal/trinity"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
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
		binary.Patch(&data, patterns.PortalPattern, patterns.PortalPattern.Empty())
		binary.Patch(&data, patterns.ConnectToModulusPattern, trinity.RsaModulus)
		binary.Patch(&data, patterns.CryptoEdPublicKeyPattern, trinity.CryptoEd25519PublicKey)

		wd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("unable to determine current working directory: %w", err)
		}

		file, err := filepath.Abs(options.OutputFile)
		if err != nil {
			return fmt.Errorf("invalid output file path specified: %w", err)
		}

		if err = os.Remove(file); err != nil && !errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("unable to remove file that already exists at %v: %w", file, err)
		}

		err = os.WriteFile(file, data, 0777)
		if err != nil {
			return fmt.Errorf("unable to write to file %v: %w", file, err)
		}

		if options.StripCodesignAttributes {
			if err = platform.RemoveCodesigningSignature(file); err != nil {
				return fmt.Errorf("unable to remove codesigning signature from %v: %w", file, err)
			}
		}
		relativePath, _ := filepath.Rel(wd, file)

		fmt.Printf("Client has been successfully patched and saved to \"%v\".\n", relativePath)

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
	rootCmd.Example = "wowpatch -l ./your/wow/exe -o ./patched-exe"
	rootCmd.PersistentFlags().StringVarP(&options.OutputFile, "output-file", "o", "Arctium", "where to output a modified client")
	rootCmd.PersistentFlags().StringVarP(&options.WarcraftExeLocation, "warcraft-exe", "l", platform.FindWarcraftClientExecutable(), "the location of the WoW executable")
	rootCmd.PersistentFlags().BoolVarP(&options.StripCodesignAttributes, "strip-binary-codesign", "s", true, "removes macOS codesigning from resulting binary")
}
