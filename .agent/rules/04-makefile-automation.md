# Regra 04: Interface Universal de Comandos via Makefile Autodocumentado

O `Makefile` é a interface universal de automação de desenvolvimento, compilação, validação de qualidade e release do projeto. O agente deve usar os comandos do Makefile para todas as operações de ciclo de vida.

## 1. Alvos Padronizados Obrigatórios

- `make help`: (Alvo padrão) Lista todos os alvos disponíveis com suas respectivas descrições extraídas de comentários inline `##`.
- `make setup`: Instala e valida ferramentas locais de desenvolvimento (`golangci-lint`, `govulncheck`).
- `make dev`: Executa a aplicação em modo de desenvolvimento local.
- `make run`: Executa a aplicação diretamente do código-fonte Go (`go run cmd/md_server/main.go`).
- `make test`: Executa todos os testes do projeto com flag de concorrência e `-race`.
- `make test-unit`: Executa testes unitários rápidos.
- `make test-coverage`: Executa a suíte de testes calculando a cobertura via `scripts/coverage.sh` e validando o limiar de 80%.
- `make lint`: Executa a análise estática via `golangci-lint` com as regras estritas configuradas em `.golangci.yml`.
- `make fmt`: Aplica formatação automática (`gofmt -s` e `goimports`).
- `make check`: Quality Gate unificado e obrigatório antes de qualquer conclusão de tarefa (`fmt` + `lint` + `govulncheck` + `test-coverage`).
- `make build`: Compila o binário de produção otimizado com stripping de símbolos (`-ldflags="-s -w ..."`) para o sistema operacional corrente em `dist/`.
- `make build-all`: Compila a matriz multiplataforma oficial (Linux/AMD64 e Windows/AMD64 com terminação `.exe`).
- `make clean`: Remove binários compilados em `dist/`, arquivos `coverage.out`, `coverage.html` e caches temporários.

## 2. Injeção de Versão via ldflags
Todos os alvos de compilação devem injetar as variáveis do pacote `internal/version`:
```makefile
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
COMMIT  ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo "none")
DATE    ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

LDFLAGS = -s -w \
  -X 'github.com/douglas/markdown-server/internal/version.Version=$(VERSION)' \
  -X 'github.com/douglas/markdown-server/internal/version.Commit=$(COMMIT)' \
  -X 'github.com/douglas/markdown-server/internal/version.Date=$(DATE)'
```
