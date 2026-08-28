# Proposta: Fundação Arquitetural e de Engenharia do md_server (project-foundation)

## Why

Desenvolvedores, engenheiros, analistas técnicos e equipes multidisciplinares necessitam frequentemente de uma forma simples, rápida e visualmente agradável de visualizar, navegar e validar documentações locais em formato Markdown com diagramas Mermaid e blocos de código com destaque de sintaxe (syntax highlighting), sem a complexidade de configurar servidores Node.js pesados, serviços em nuvem ou ferramentas web externas.

O objetivo do **md_server** (apresentado visualmente como **Markdown Viewer**) é fornecer um executável único, estático, multiplataforma (Windows e Linux/AMD64), ultrarrápido e sem dependências externas. O usuário pode simplesmente dar um **duplo clique no executável** (ou executá-lo via linha de comando informando a pasta raiz) para que a aplicação sirva imediatamente todos os arquivos Markdown daquele diretório, abra o navegador padrão automaticamente e permita a navegação fluida entre documentos através de links internos, renderização de diagramas Mermaid em tempo real e um design clean, moderno e minimalista inspirado no [md-reader](https://md-reader.github.io/).

A interface visual conta com uma identidade visual moderna (ícone vetorial estilizado tanto no cabeçalho quanto no favicon da aba do navegador), menu esquerdo retrátil com altura total de 100vh, responsividade total para smartphones e tablets, e gerador de **QR Code** para compartilhamento rápido na rede local.

## What Changes

- **Identidade Visual, Branding e Favicon (Markdown Viewer)**:
  - Marca **"Markdown Viewer"** com ícone vetorial SVG exclusivo e moderno, servido tanto como elemento visual no cabeçalho quanto como **Favicon** (`/favicon.ico` e SVG icon) para a aba do navegador.
- **Navegação e Menu Lateral Retrátil (Full-Height 100vh)**:
  - Menu esquerdo retrátil/colapsável com botão de alternância suave e atalho de teclado, permitindo maximizar a área de leitura.
  - Layout com barra lateral contínua que estende seu fundo e linha divisória até a base da viewport (`100vh`), mantendo consistência visual mesmo em documentos curtos ou sem conteúdo.
- **Responsividade Total para Mobile e Tablets**:
  - Layout flexível e adaptativo para smartphones e tablets, com barra lateral estilo *drawer* (gaveta deslizante) e overlay de fundo para navegação por toque.
- **Compartilhamento Rápido via QR Code**:
  - Botão integrado no cabeçalho que abre um modal com o **QR Code** da página atual e da URL do servidor local, facilitando o acesso direto a partir de smartphones e tablets.
- **Arquitetura Canônica em Go (Clean Architecture)**:
  - Estrutura estrita de pastas: `cmd/md_server/` (Composition Root), `internal/core/domain/`, `internal/core/ports/`, `internal/core/services/` e `internal/adapters/`.
  - Injeção explícita de dependências via construtores e propagação de `context.Context` com encerramento gracioso (*graceful shutdown*).
- **Ponto de Entrada, CLI e Suporte a Duplo Clique (Windows / Linux)**:
  - CLI via `spf13/cobra` com suporte a flags (`--dir`, `--port`, `--open`, `--version`) e subcomando `version`.
  - Comportamento de inicialização por duplo clique: ao ser executado sem argumentos, serve a pasta corrente, aloca porta livre e abre o navegador padrão automaticamente.
- **Renderização de Markdown, Mermaid e Syntax Highlighting**:
  - Processamento de Markdown (GFM) via `goldmark` com tabelas e autolinks.
  - Destaque de sintaxe de alto contraste com `chroma/v2`.
  - Renderização client-side assíncrona de diagramas `mermaid.js`.
  - Resolução inteligente de links Markdown internos (`./documento.md` para rotas web).
- **Frontend Embutido via `//go:embed`**:
  - Templates HTML, CSS minimalista, ícone SVG e scripts JS empacotados diretamente no binário final.
- **Infraestrutura de Testes TDD/BDD e Barreira de Cobertura >= 80%**:
  - Testes com subtestes BDD declarativos (`Given / When / Then`), asserções `testify` e validação via `scripts/coverage.sh`.
- **Automação via Makefile e CI/CD**:
  - Makefile universal e pipelines GitHub Actions para compilação multiplataforma (Linux `.tar.gz` e Windows `.zip` com `.exe`).

## Capabilities

### New Capabilities
- `project-foundation`: Estabelece a fundação arquitetural em Go (Clean Architecture), o visualizador web "Markdown Viewer" com ícone SVG/Favicon moderno, menu lateral retrátil full-height (100vh), interface minimalista responsiva para smartphones/tablets, compartilhamento facilitado via QR Code, renderização de Markdown, Mermaid e Chroma, inicialização por duplo clique no Windows/Linux, suíte TDD/BDD com cobertura >= 80% e automação universal via Makefile.

### Modified Capabilities
<!-- Nenhuma capacidade existente foi modificada (projeto novo). -->

## Impact

- **Interface e Experiência do Usuário (UI/UX)**: Ícone estilizado no cabeçalho e na aba do navegador (Favicon), interface limpa estilo md-reader, menu retrátil que aproveita 100% da tela, responsividade em smartphones e compartilhamento via QR Code.
- **Compatibilidade**: Desktop (Windows/Linux/macOS), tablets e smartphones.
- **Performance**: Zero dependências pesadas no backend, tempo de inicialização < 100ms e consumo de memória < 25MB.
