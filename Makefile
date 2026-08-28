# Makefile Universal e Autodocumentado para md_server

APP_NAME := md_server
MODULE   := github.com/douglas/markdown-server
CMD_DIR  := ./cmd/md_server
DIST_DIR := ./dist

# Versionamento dinâmico via Git e ldflags
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
COMMIT  ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo "none")
DATE    ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

LDFLAGS := -s -w \
	-X '$(MODULE)/internal/version.Version=$(VERSION)' \
	-X '$(MODULE)/internal/version.Commit=$(COMMIT)' \
	-X '$(MODULE)/internal/version.Date=$(DATE)'

.DEFAULT_GOAL := help

.PHONY: help
help: ## Exibe este menu de ajuda autodocumentado com todos os comandos disponíveis
	@echo "========================================================================"
	@echo " Makefile - $(APP_NAME) (Versão: $(VERSION))"
	@echo "========================================================================"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-18s\033[0m %s\n", $$1, $$2}'
	@echo "========================================================================"

.PHONY: setup
setup: ## Instala e valida ferramentas de desenvolvimento locais (linters, etc.)
	@echo "==> Verificando dependências e ferramentas de desenvolvimento..."
	@go mod download
	@which golangci-lint > /dev/null 2>&1 || (echo "Instalando golangci-lint..." && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest)
	@which govulncheck > /dev/null 2>&1 || (echo "Instalando govulncheck..." && go install golang.org/x/vuln/cmd/govulncheck@latest)
	@echo "==> Setup de ferramentas concluído com sucesso!"

.PHONY: dev
dev: ## Inicia o servidor em modo de desenvolvimento local
	@echo "==> Iniciando $(APP_NAME) em modo dev..."
	@go run $(CMD_DIR)/main.go --dir . --port 8080 --open=true

.PHONY: run
run: ## Executa a aplicação diretamente a partir do código-fonte
	@go run $(CMD_DIR)/main.go

.PHONY: test
test: ## Executa a suíte completa de testes com detecção de concorrência (-race)
	@echo "==> Executando testes automatizados..."
	@go test -race ./...

.PHONY: test-unit
test-unit: ## Executa testes rápidos
	@echo "==> Executando testes rápidos..."
	@go test ./...

.PHONY: test-coverage
test-coverage: ## Executa testes calculando cobertura global e validando a barreira >= 80%
	@echo "==> Executando testes com barreira de cobertura (scripts/coverage.sh)..."
	@bash scripts/coverage.sh

.PHONY: lint
lint: ## Executa o linter estrito golangci-lint
	@echo "==> Executando golangci-lint..."
	@golangci-lint run ./... || true

.PHONY: fmt
fmt: ## Formata todo o código-fonte Go (gofmt)
	@echo "==> Formatando arquivos Go..."
	@gofmt -s -w .

.PHONY: check
check: fmt test-coverage ## Quality Gate unificado (fmt + test-coverage >= 80%)
	@echo "==> [SUCESSO] Quality Gate aprovado com sucesso!"

.PHONY: build
build: ## Compila o binário de produção otimizado com stripping de símbolos para a plataforma corrente
	@echo "==> Compilando $(APP_NAME) para o sistema local (Versão: $(VERSION))..."
	@mkdir -p $(DIST_DIR)
	@CGO_ENABLED=0 go build -ldflags="$(LDFLAGS)" -o $(DIST_DIR)/$(APP_NAME) $(CMD_DIR)/main.go
	@echo "==> Binário compilado com sucesso em $(DIST_DIR)/$(APP_NAME)"

.PHONY: build-all
build-all: ## Compila binários para a matriz multiplataforma oficial (Linux e Windows AMD64)
	@echo "==> Compilando matriz multiplataforma oficial (Linux e Windows/AMD64)..."
	@mkdir -p $(DIST_DIR)
	@echo "  -> Compilando Linux/AMD64..."
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="$(LDFLAGS)" -o $(DIST_DIR)/$(APP_NAME)-linux-amd64 $(CMD_DIR)/main.go
	@echo "  -> Compilando Windows/AMD64..."
	@GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="$(LDFLAGS)" -o $(DIST_DIR)/$(APP_NAME)-windows-amd64.exe $(CMD_DIR)/main.go
	@echo "==> Compilação multiplataforma concluída com sucesso em $(DIST_DIR)!"

.PHONY: clean
clean: ## Limpa binários compilados em dist/, relatórios de cobertura e arquivos temporários
	@echo "==> Limpando artefatos temporários e de compilação..."
	@rm -rf $(DIST_DIR) coverage.out coverage.html
	@echo "==> Limpeza concluída!"
