package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/douglas/markdown-server/internal/version"
)

// NewVersionCommand cria o subcomando 'version' para exibição dos metadados de compilação.
func NewVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Exibe a versão e metadados de compilação do md_server",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version.Formatted("md_server"))
		},
	}
}
