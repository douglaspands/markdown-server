# Tarefas de Implementação: Fundação Arquitetural do md_server (project-foundation)

## 1. Scaffolding, Módulo Go e Versionamento Dinâmico

- [x] 1.1 Inicializar módulo Go 1.22+ (`go.mod`) com nome canônico e adicionar dependências fundamentais (`cobra`, `goldmark`, `chroma/v2`, `testify`).
- [x] 1.2 Criar pacote `internal/version/version.go` com variáveis públicas (`Version`, `Commit`, `Date`) e suporte a injeção em tempo de compilação via `-ldflags`.
- [x] 1.3 Criar pacote `internal/testutils/` com utilitários auxiliares de teste, fixtures de Markdown e gerador de arquivos temporários.

## 2. Núcleo de Domínio e Contratos de Portas (Core Domain & Ports)

- [x] 2.1 Criar modelos de domínio em `internal/core/domain/` (`MarkdownDocument`, `FileNode`, `HealthStatus`, `ServerConfig` e erros de domínio sentinela como `ErrFileNotFound`, `ErrPathEscape`).
- [x] 2.2 Definir interfaces formais de portas de entrada em `internal/core/ports/input_ports.go` (`MarkdownServicePort`, `HealthServicePort`).
- [x] 2.3 Definir interfaces formais de portas de saída em `internal/core/ports/output_ports.go` (`FileScannerPort`, `MarkdownRendererPort`, `BrowserLauncherPort`).

## 3. Implementação dos Serviços de Domínio e Testes BDD

- [x] 3.1 Implementar serviço de referência `HealthCheckService` em `internal/core/services/health_service.go`.
- [x] 3.2 Criar suíte de testes BDD para `HealthCheckService` em `internal/core/services/health_service_test.go` utilizando `testify` (cobertura >= 80%).
- [x] 3.3 Implementar `MarkdownService` em `internal/core/services/markdown_service.go` orquestrando o scanner de arquivos, o renderizador Markdown e a construção da árvore de diretórios.
- [x] 3.4 Criar suíte de testes BDD para `MarkdownService` em `internal/core/services/markdown_service_test.go` com mocks determinísticos e validação rigorosa de cenários de sucesso e erro.

## 4. Adaptadores de Saída (Output Adapters)

- [x] 4.1 Implementar adaptador `FSScanner` em `internal/adapters/out/fs/scanner.go` com sanitização estrita de caminhos contra Directory Traversal e leitura recursiva de arquivos `.md`.
- [x] 4.2 Criar suíte de testes BDD para `FSScanner` em `internal/adapters/out/fs/scanner_test.go` cobrindo cenários válidos, arquivos ausentes e tentativas de path escape.
- [x] 4.3 Implementar adaptador `GoldmarkRenderer` em `internal/adapters/out/renderer/goldmark.go` com extensões GFM, syntax highlighting via `chroma/v2` e AST transformer para reescrita de links relativos `.md`.
- [x] 4.4 Criar suíte de testes BDD para `GoldmarkRenderer` em `internal/adapters/out/renderer/goldmark_test.go` validando parsing de Markdown, tabelas GFM, blocos de código coloridos e transformação de links internos.
- [x] 4.5 Implementar adaptador `BrowserLauncher` em `internal/adapters/out/browser/launcher.go` com suporte multiplataforma para abertura do navegador padrão no Windows (`cmd.exe /c start`) e Linux (`xdg-open`).

## 5. Assets Web Embutidos e Interface do Usuário (Frontend)

- [x] 5.1 Criar ícone vetorial SVG em `web/static/img/icon.svg` e atualizar templates HTML (`base.html`, `document.html`, `404.html`) com o branding "Markdown Viewer", ícone SVG, tag de favicon, botão retrátil de toggle de sidebar e modal de QR Code.
- [x] 5.2 Refatorar estilos CSS (`style.css`, `chroma.css`) em `web/static/css/` com design minimalista estilo *md-reader*, altura total de 100vh na barra lateral, animação de colapso, modal de QR Code e media queries para smartphones e tablets.
- [x] 5.3 Atualizar scripts JS (`app.js` e embutir gerador de QR Code) em `web/static/js/` para controlar o menu retrátil com persistência no `localStorage`, renderização de QR Code da URL atual e suporte a toque/drawer no mobile.
- [x] 5.4 Mapear rota de `/favicon.ico` no servidor HTTP (`internal/adapters/in/http/server.go`), atualizar `web/embed.go` e testes unitários em `web/embed_test.go`.

## 6. Adaptadores de Entrada e Ponto de Entrada (Composition Root)

- [x] 6.1 Implementar servidor HTTP em `internal/adapters/in/http/server.go` e handlers em `handlers.go` com rotas para renderização de páginas, assets estáticos, endpoint `/api/health` e encerramento gracioso via `context.Context`.
- [x] 6.2 Criar suíte de testes HTTP em `internal/adapters/in/http/server_test.go` com `net/http/httptest` validando rotas, cabeçalhos, status 200, 404 e 403.
- [x] 6.3 Implementar adaptador de comando CLI via Cobra em `internal/adapters/in/cli/root.go` com suporte a flags (`--dir`, `--port`, `--open`, `--version`) e subcomando `version`.
- [x] 6.4 Implementar Composition Root em `cmd/md_server/main.go` instanciando dependências manualmente, detectando portas livres, disparando a abertura de navegador e tratando inicialização por duplo clique no Windows.

## 7. Automação via Makefile, Script de Cobertura e Linter

- [x] 7.1 Criar script `scripts/coverage.sh` com permissão de execução para cálculo global de cobertura de código e falha mandatória se cobertura for inferior a 80%.
- [x] 7.2 Criar arquivo de configuração `.golangci.yml` com regras estritas de análise estática e qualidade de código Go.
- [x] 7.3 Criar `Makefile` universal autodocumentado com alvos obrigatórios (`help`, `setup`, `dev`, `run`, `test`, `test-unit`, `test-coverage`, `lint`, `fmt`, `check`, `build`, `build-all`, `clean`) e injeção de ldflags.

## 8. CI/CD Multiplataforma, Documentação e Validação Final

- [x] 8.1 Criar pipeline de Integração Contínua em `.github/workflows/ci.yml` para validação em Pull Requests e commits na branch `main` (lint, govulncheck, testes com cobertura >= 80% e validação OpenSpec).
- [x] 8.2 Criar pipeline de Release em `.github/workflows/release.yml` com matriz de cross-compilação para Windows/AMD64 (`.zip` com `.exe`) e Linux/AMD64 (`.tar.gz`) e publicação automática no GitHub Releases com SHA256.
- [x] 8.3 Criar documentação viva e exaustiva no `README.md` raiz com badges, visão geral, guia de instalação/uso, guia do desenvolvedor, arquitetura e convenções Git/IA.
- [x] 8.4 Executar validação final do Quality Gate com `make check` e `openspec validate --all` para assegurar 100% de conformidade com todos os 10 pilares.
