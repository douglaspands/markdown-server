## 1. Governança Git e Isolamento da Feature Branch

- [x] 1.1 Inicializar o repositório Git (se ainda não inicializado) e criar a branch de trabalho `feature/extended-markdown-plugins`
- [x] 1.2 Garantir `.gitignore` íntegro e alinhado aos padrões do projeto (ocultando binários, relatórios de cobertura e builds temporários)

## 2. Extensões Nativas do Goldmark e Alertas GitHub (Backend Go)

- [x] 2.1 Habilitar `extension.Footnote`, `extension.DefinitionList` e `extension.Typographer` no construtor `GoldmarkRenderer`
- [x] 2.2 Implementar transformador/processador para blocos de aviso no padrão GitHub Alerts (`[!NOTE]`, `[!TIP]`, `[!IMPORTANT]`, `[!WARNING]`, `[!CAUTION]`) com ícones SVG
- [x] 2.3 Adicionar suporte à preservação de delimitadores de fórmulas matemáticas LaTeX (`$...$` e `$$...$$`) na pipeline de renderização

## 3. Estilização e Componentes Visuais (Frontend CSS & Templates)

- [x] 3.1 Implementar estilos CSS para os 5 tipos de GitHub Alerts com bordas temáticas, ícones vetoriais SVG e contraste adaptado a Dark/Light Mode
- [x] 3.2 Implementar estilos CSS para Notas de Rodapé (`.footnotes`), Listas de Definição (`dl`, `dt`, `dd`) e fórmulas matemáticas
- [x] 3.3 Incluir assets KaTeX no template `web/templates/base.html` para renderização matemática assíncrona

## 4. Interatividade e Cópia de Código (Frontend JavaScript)

- [x] 4.1 Implementar injeção do botão de cópia de código (*Copy Button*) com evento `clipboard` e feedback visual em `web/static/js/app.js`
- [x] 4.2 Integrar inicialização do KaTeX para renderização automática de expressões matemáticas no `web/static/js/app.js`

## 5. Testes Automatizados, Cobertura e Validação de Qualidade

- [x] 5.1 Implementar e atualizar testes unitários em `internal/adapters/out/renderer/goldmark_test.go` cobrindo Footnotes, Definition Lists, Typographer, GitHub Alerts e Math
- [x] 5.2 Executar a suíte de testes com barreira de cobertura >= 80% via `make test-coverage`
- [x] 5.3 Validar a integridade geral do projeto com `make check` e compilação com `make build-all`

## 6. Documentação Visual e Demonstração no README

- [x] 6.1 Gerar mockups visuais/prints de tela em alta fidelidade com dados fictícios ilustrativos em `docs/screenshots/` (exibindo interface, alertas, Mermaid, Dark/Light mode e QR Code)
- [x] 6.2 Atualizar o `README.md` com os prints de tela, tabela de recursos visuais, exemplos das novas extensões de Markdown e badges atualizados

## 7. Conventional Commits e Finalização

- [x] 7.1 Registrar commits semânticos atômicos seguindo o padrão Conventional Commits (`feat(renderer):`, `style(ui):`, `docs(readme):`, `test(renderer):`)
- [x] 7.2 Preparar para solicitação de autorização do usuário antes de realizar o Squash Merge na branch `main`

## 8. Flag de Segurança CLI `--lan` e QR Code Condicional

- [x] 8.1 Adicionar a flag `--lan` / `-l` (padrão: `false`) no CLI Cobra (`internal/adapters/in/cli/root.go`) e na entidade de configuração (`internal/core/domain/config.go`)
- [x] 8.2 Atualizar o adaptador de servidor HTTP (`internal/adapters/in/http/server.go`) para escuta padrão em `127.0.0.1:<port>` (modo seguro) e `0.0.0.0:<port>` apenas quando `ExposeLAN` for verdadeiro
- [x] 8.3 Atualizar `PageData` e `handleDocument` para propagar `ShowQRCode = ExposeLAN`, e atualizar o template `web/templates/base.html` para renderizar o botão de QR Code condicionalmente (`{{if .ShowQRCode}}`)
- [x] 8.4 Atualizar o banner de console no terminal (`cmd/md_server/main.go`) para exibir a URL LAN somente quando `--lan` estiver ativo
- [x] 8.5 Implementar e atualizar testes unitários e de integração em `internal/adapters/in/cli/cli_test.go`, `internal/adapters/in/http/server_test.go` e `cmd/md_server/main_test.go`
- [x] 8.6 Validar a barreira de cobertura global >= 80% (`make test-coverage`) e aprovação no `make check`

## 9. Interatividade Avançada de Diagramas (Zoom e Pan Mermaid)

- [x] 9.1 Implementar barra de ferramentas flutuante (+, -, ↺) e estilos CSS para controles de diagramas no `web/static/css/style.css`
- [x] 9.2 Implementar manipulador de Pan & Zoom com eventos de mouse/touch (`pointerdown`, `pointermove`, `wheel`) e limites de escala em `web/static/js/app.js`
- [x] 9.3 Integrar o reset e re-aplicação de Zoom/Pan durante alternância de tema no `web/static/js/app.js`

## 10. Stack Tipográfica Otimizada para Markdown, Scripts e Código

- [x] 10.1 Atualizar variáveis CSS `--font-sans` e `--font-mono` em `web/static/css/style.css` com a hierarquia de fontes de alto padrão da indústria
- [x] 10.2 Ajustar estilos de tipografia de cabeçalhos, corpo de texto, tabelas, alertas e blocos de código (`.highlight`, `code`, `pre`, `.mermaid`) para máxima legibilidade e alinhamento visual
- [x] 10.3 Atualizar documentação e recursos no `README.md`

## 11. Otimização Óptica de Escaneabilidade do QR Code

- [x] 11.1 Aumentar a dimensão de renderização do QR Code para 240x240 pixels com `CorrectLevel.M` e normalização de URL em `web/static/js/app.js`
- [x] 11.2 Estilizar a moldura branca de alto contraste (*quiet zone*), badge de IP/Porta e botões de ação em `web/static/css/style.css` e `web/templates/base.html`
- [x] 11.3 Atualizar mockups e documentação no `README.md`
