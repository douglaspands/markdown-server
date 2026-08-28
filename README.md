# md_server 📄⚡

[![Go Version](https://img.shields.io/badge/Go-1.22%2B-blue.svg)](https://golang.org)
[![CI Quality Gate](https://img.shields.io/badge/CI-Quality%20Gate%20Passing-success.svg)](.github/workflows/ci.yml)
[![Coverage](https://img.shields.io/badge/Coverage-%E2%89%A580%25%20(85.6%25)-brightgreen.svg)](scripts/coverage.sh)
[![Platform](https://img.shields.io/badge/Platform-Windows%20%7C%20Linux%20(amd64)-orange.svg)](#plataformas-alvo)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

**md_server** é um visualizador e servidor web local de arquivos Markdown ultrarrápido, minimalista e moderno, projetado para renderizar documentações locais com suporte completo a diagramas **Mermaid**, destaque de sintaxe (**Chroma**), árvore lateral de arquivos e navegação transparente por links internos.

Foi construído em **Go 1.22+** com **Clean Architecture (Ports & Adapters)**, gerando um binário estático e independente sem nenhuma dependência externa de runtime (zero Node.js, zero Python, zero browsers headless no servidor).

---

## ✨ Principais Recursos

- 🖱️ **Zero-Config & Duplo Clique**: No Windows (ou Linux), dê um duplo clique no executável na pasta desejada para abrir a aplicação e o navegador padrão instantaneamente.
- 📊 **Diagramas Mermaid Integrados**: Renderização assíncrona no cliente de blocos ` ```mermaid ` em diagramas SVG vetoriais e interativos.
- 🎨 **Destaque de Sintaxe Colorido**: Suporte a dezenas de linguagens de programação (Go, Python, JS, TS, JSON, YAML, SQL, Shell, etc.) via Chroma com temas de alto contraste.
- 🧭 **Navegação Contínua por Links**: Links Markdown relativos (`[Guia](./guia.md)`) são automaticamente interceptados e convertidos em rotas web fluidas.
- 📁 **Árvore Lateral de Arquivos (Sidebar Tree)**: Navegue por todas as pastas e arquivos `.md` do diretório raiz através de uma barra lateral expansível.
- 🌓 **Tema Claro / Escuro (Dark/Light Mode)**: Alternância de tema com persistência local e detecção automática da preferência do sistema operacional.
- 🔒 **Segurança Nativa**: Prevenção estrita contra Directory Traversal (`..` escapes) garantindo que nenhum arquivo fora da pasta raiz seja acessado.

---

## 🚀 Instalação e Uso

### 1. Download dos Binários Pré-Compilados
Baixe a versão mais recente para o seu sistema operacional na aba [Releases](../../releases):
- **Windows (amd64)**: `md_server-windows-amd64.zip` (contém `md_server.exe`)
- **Linux (amd64)**: `md_server-linux-amd64.tar.gz` (contém `md_server`)

### 2. Execução via Linha de Comando (CLI)

```bash
# Executa servindo a pasta atual na porta padrão 8080 e abre o navegador
./md_server

# Servindo uma pasta específica em outra porta sem abrir o navegador
./md_server --dir ./minha-documentacao --port 9090 --open=false

# Exibe a versão oficial e metadados de compilação
./md_server version
```

### Tabela de Parâmetros e Flags

| Flag | Shorthand | Padrão | Descrição |
| :--- | :--- | :--- | :--- |
| `--dir` | `-d` | `.` | Diretório raiz local contendo os arquivos Markdown a serem servidos |
| `--port` | `-p` | `8080` | Porta TCP para escuta HTTP (aloca porta livre se ocupada) |
| `--open` | `-o` | `true` | Dispara automaticamente a abertura da URL no navegador padrão |
| `--version` | `-v` | `false` | Exibe versão semântica, commit e data de compilação |

---

## 🏗️ Arquitetura de Software (Clean Architecture)

O código segue rigorosamente os princípios de Clean Architecture (Ports & Adapters / Arquitetura Hexagonal):

```
md_server/
├── cmd/
│   └── md_server/           # Ponto de entrada (Composition Root e CLI)
├── internal/
│   ├── core/
│   │   ├── domain/          # Entidades de negócio (MarkdownDocument, FileNode, HealthStatus)
│   │   ├── ports/           # Interfaces formais de entrada e saída
│   │   └── services/        # Casos de uso (MarkdownService, HealthCheckService)
│   ├── adapters/
│   │   ├── in/              # Adaptadores de entrada (HTTP server, CLI Cobra)
│   │   └── out/             # Adaptadores de saída (FS scanner, Goldmark/Chroma renderer, Browser launcher)
│   ├── version/             # Injeção dinâmica de versão via ldflags
│   └── testutils/           # Mocks determinísticos e fixtures para testes
├── web/                     # Assets embutidos via //go:embed (Templates HTML, CSS, JS Mermaid)
├── scripts/                 # Automação de cobertura (scripts/coverage.sh >= 80%)
└── .github/workflows/       # CI e Release Multiplataforma
```

---

## 🛠️ Guia do Desenvolvedor

### Pré-Requisitos
- Go 1.22 ou superior
- GNU Make (opcional, recomendado)

### Comandos do Makefile

```bash
# Exibe todos os comandos disponíveis
make help

# Instala ferramentas locais (golangci-lint, govulncheck)
make setup

# Executa todos os testes com verificação de cobertura (barreira >= 80%)
make test-coverage

# Valida o Quality Gate completo (formatação + testes)
make check

# Compila o binário de produção local
make build

# Compila a matriz multiplataforma (Linux e Windows/AMD64)
make build-all
```

---

## 🤖 Governança de Agentes de IA (Antigravity CLI)

Este projeto opera sob governança de agentes autônomos com o arcabouço **Antigravity (AGY)**:
- **Prioridade de Ferramentas Nativas**: Uso mandatório de `write_to_file`, `replace_file_content`, `view_file` e `grep_search` (ver [AGENTS.md](file:///home/douglas/Workspace/gemini/markdown-server/AGENTS.md) e [.agent/rules/](file:///home/douglas/Workspace/gemini/markdown-server/.agent/rules/)).
- **Alinhamento PO & QA**: Especificações gerenciadas via OpenSpec em [openspec/](file:///home/douglas/Workspace/gemini/markdown-server/openspec/).
- **Higiene Git**: Padrão Conventional Commits, branches isoladas e integração obrigatória via **Squash Merge**.

---

## 📄 Licença

Distribuído sob a licença MIT. Consulte `LICENSE` para mais detalhes.
