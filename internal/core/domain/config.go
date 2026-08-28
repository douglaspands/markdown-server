package domain

// ServerConfig encapsula as configurações de inicialização do servidor HTTP e da aplicação.
type ServerConfig struct {
	// RootDir é o diretório raiz local servido pelo md_server.
	RootDir string

	// Port é a porta TCP utilizada para atender requisições HTTP (padrão: 8080).
	Port int

	// AutoOpenBrowser indica se o navegador padrão do sistema operacional deve ser aberto na inicialização.
	AutoOpenBrowser bool

	// ExposeLAN indica se o servidor deve escutar em 0.0.0.0 e expor o QR Code para a rede local (padrão: false / escuta em 127.0.0.1).
	ExposeLAN bool
}
