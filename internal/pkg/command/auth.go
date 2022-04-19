package command

import (
	"fmt"
	"github.com/itzoo-space/go-flagenv"
	inrConst "github.com/michael-kalashnikov-dev/gringotts/internal/pkg/constant"
	"github.com/michael-kalashnikov-dev/gringotts/pkg/constant"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
)

// authCmd entry point to application API
var authCmd = &cobra.Command{
	Use:   inrConst.AuthCLIName,
	Short: "You shall not pass!",
	Long: `
=============================================
== YOOOOOOOU SHAAAAALL NOOOOOOOT PAAAAAASS ==
=============================================
`,
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().Lookup("init").Value.String() == "true" {
			fmt.Println("Start exporting...")
			file := fmt.Sprintf(
				"./%s",
				inrConst.AuthConfigFileName,
			)
			fmt.Printf("Attempt to write to %s\n", file)
			cobra.CheckErr(viper.WriteConfigAs(file))
			fmt.Printf("Exported successfully to %s\n", file)
			os.Exit(0)
		} else {
			cobra.CheckErr(cmd.Help())
		}
	},
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile := viper.GetString(constant.AuthAppEnvConfigFile); cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory.
		viper.AddConfigPath(home)
		viper.SetConfigType(inrConst.AuthConfigFileType)
		viper.SetConfigName(inrConst.AuthConfigFileName)
	}

	// Read in environment variables that match
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		_, err := fmt.Fprintln(
			os.Stderr,
			"Using config file: ",
			viper.ConfigFileUsed(),
		)
		cobra.CheckErr(err)
	}
}

func authNormalizer(flag string) string {
	return strings.Replace(flag, constant.AuthAppEnvPrefix, "", 1)
}

func init() {
	cobra.OnInitialize(initConfig)

	flagenv.New(
		authCmd.Flags(),
		flagenv.String(),
		flagenv.WithNormalizers(authNormalizer),
		flagenv.WithEnvName(constant.AuthAppEnvConfigFile),
		flagenv.WithShorthand("-^"),
		flagenv.WithUsage(
			fmt.Sprintf(
				"config file (default is specified in \"%s\" env or in $HOME/%s)",
				constant.AuthAppEnvConfigFile, inrConst.AuthConfigFileName,
			),
		),
	)

	flagenv.New(
		authCmd.Flags(),
		flagenv.Bool(),
		flagenv.WithFlagName("init"),
		flagenv.WithShorthand("-"),
		flagenv.WithUsage("generates config file with all envs in current work directory and exits"),
	)
}

func ExecuteAuth() {
	cobra.CheckErr(authCmd.Execute())
}
