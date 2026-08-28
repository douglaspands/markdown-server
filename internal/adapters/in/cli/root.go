package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/douglas/markdown-server/internal/core/domain"
	"github.com/douglas/markdown-server/internal/version"
)

// RunnerFunc é o callback executado pelo comando raiz com as configurações parsed.
type RunnerFunc func(cmd *cobra.Command, args []string, cfg domain.ServerConfig) error

// NewRootCommand instancia o comando raiz do CLI do md_server configurando flags e subcomandos.
func NewRootCommand(runner RunnerFunc) *cobra.Command {
	var (
		rootDir     string
		port        int
		autoOpen    bool
		exposeLAN   bool
		showVersion bool
	)

	rootCmd := &cobra.Command{
		Use:   "md_server [caminho_opcional]",
		Short: "md_server - Servidor e visualizador local de arquivos Markdown com Mermaid e Chroma",
		Long: `md_server é uma ferramenta de alta performance para visualização local de documentos Markdown.
Permite navegar entre arquivos por links internos, renderiza diagramas Mermaid e código colorido.
Suporta duplo clique no Windows com abertura automática do navegador padrão.`,
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if showVersion {
				fmt.Println(version.Formatted("md_server"))
				return nil
			}

			// Se um argumento de caminho for passado como argumento posicional, usa como rootDir
			if len(args) == 1 && rootDir == "." {
				rootDir = args[0]
			}

			cfg := domain.ServerConfig{
				RootDir:         rootDir,
				Port:            port,
				AutoOpenBrowser: autoOpen,
				ExposeLAN:       exposeLAN,
			}

			if runner != nil {
				return runner(cmd, args, cfg)
			}
			return nil
		},
	}

	rootCmd.Flags().StringVarP(&rootDir, "dir", "d", ".", "Diretório raiz contendo os arquivos Markdown a serem servidos")
	rootCmd.Flags().IntVarP(&port, "port", "p", 8080, "Porta TCP para o servidor HTTP")
	rootCmd.Flags().BoolVarP(&autoOpen, "open", "o", true, "Abre automaticamente o navegador padrão do sistema operacional")
	rootCmd.Flags().BoolVarP(&exposeLAN, "lan", "l", false, "Habilita escuta em todas as interfaces de rede local (0.0.0.0) e exibe o QR Code")
	rootCmd.Flags().BoolVarP(&showVersion, "version", "v", false, "Exibe os metadados de versão da aplicação")

	rootCmd.AddCommand(NewVersionCommand())

	return rootCmd
}

// Execute executa o CLI principal capturando erros para a saída padrão.
func Execute(runner RunnerFunc) {
	rootCmd := NewRootCommand(runner)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Erro: %v\n", err)
		os.Exit(1)
	}
}
