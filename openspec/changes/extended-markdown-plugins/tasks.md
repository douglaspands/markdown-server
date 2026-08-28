## 1. Governança Git e Isolamento da Feature Branch

- [ ] 1.1 Inicializar o repositório Git (se ainda não inicializado) e criar a branch de trabalho `feature/extended-markdown-plugins`
- [ ] 1.2 Garantir `.gitignore` íntegro e alinhado aos padrões do projeto (ocultando binários, relatórios de cobertura e builds temporários)

## 2. Extensões Nativas do Goldmark e Alertas GitHub (Backend Go)

- [ ] 2.1 Habilitar `extension.Footnote`, `extension.DefinitionList` e `extension.Typographer` no construtor `GoldmarkRenderer`
- [ ] 2.2 Implementar transformador/processador para blocos de aviso no padrão GitHub Alerts (`[!NOTE]`, `[!TIP]`, `[!IMPORTANT]`, `[!WARNING]`, `[!CAUTION]`) com ícones SVG
- [ ] 2.3 Adicionar suporte à preservação de delimitadores de fórmulas matemáticas LaTeX (`$...$` e `$$...$$`) na pipeline de renderização

## 3. Estilização e Componentes Visuais (Frontend CSS & Templates)

- [ ] 3.1 Implementar estilos CSS para os 5 tipos de GitHub Alerts com bordas temáticas, ícones vetoriais SVG e contraste adaptado a Dark/Light Mode
- [ ] 3.2 Implementar estilos CSS para Notas de Rodapé (`.footnotes`), Listas de Definição (`dl`, `dt`, `dd`) e fórmulas matemáticas
- [ ] 3.3 Incluir assets KaTeX no template `web/templates/base.html` para renderização matemática assíncrona

## 4. Interatividade e Cópia de Código (Frontend JavaScript)

- [ ] 4.1 Implementar injeção do botão de cópia de código (*Copy Button*) com evento `clipboard` e feedback visual em `web/static/js/app.js`
- [ ] 4.2 Integrar inicialização do KaTeX para renderização automática de expressões matemáticas no `web/static/js/app.js`

## 5. Testes Automatizados, Cobertura e Validação de Qualidade

- [ ] 5.1 Implementar e atualizar testes unitários em `internal/adapters/out/renderer/goldmark_test.go` cobrindo Footnotes, Definition Lists, Typographer, GitHub Alerts e Math
- [ ] 5.2 Executar a suíte de testes com barreira de cobertura >= 80% via `make test-coverage`
- [ ] 5.3 Validar a integridade geral do projeto com `make check` e compilação com `make build-all`

## 6. Documentação Visual e Demonstração no README

- [ ] 6.1 Gerar mockups visuais/prints de tela em alta fidelidade com dados fictícios ilustrativos em `docs/screenshots/` (exibindo interface, alertas, Mermaid, Dark/Light mode e QR Code)
- [ ] 6.2 Atualizar o `README.md` com os prints de tela, tabela de recursos visuais, exemplos das novas extensões de Markdown e badges atualizados

## 7. Conventional Commits e Finalização

- [ ] 7.1 Registrar commits semânticos atômicos seguindo o padrão Conventional Commits (`feat(renderer):`, `style(ui):`, `docs(readme):`, `test(renderer):`)
- [ ] 7.2 Preparar para solicitação de autorização do usuário antes de realizar o Squash Merge na branch `main`
