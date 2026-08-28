## Why

O `md_server` atualmente oferece suporte básico ao GitHub Flavored Markdown (tabelas, task lists, strikethrough, autolinks) e diagramas Mermaid. No entanto, documentações técnicas modernas, repositórios do GitHub e bases de conhecimento ricas utilizam extensivamente recursos adicionais como **GitHub Alerts / Callouts** (`[!NOTE]`, `[!TIP]`, `[!WARNING]`, `[!IMPORTANT]`, `[!CAUTION]`), **Notas de Rodapé** (*Footnotes*), **Fórmulas Matemáticas / LaTeX**, **Listas de Definição**, **Tipografia Inteligente**, **Emoji Shortcodes** e **Botão de Cópia em Blocos de Código**.

Adicionar essas extensões alinha a experiência de visualização com o padrão oficial do GitHub Markdown, mantendo a performance extrema do Goldmark em Go e a arquitetura zero-configuração.

## What Changes

- **GitHub Alerts / Callouts (Admonitions)**: Suporte completo aos 5 tipos de avisos nativos do GitHub (`> [!NOTE]`, `> [!TIP]`, `> [!IMPORTANT]`, `> [!WARNING]`, `> [!CAUTION]`), renderizados com bordas coloridas, ícones SVG específicos, tipografia destacada e adaptação aos temas Claro e Escuro.
- **Notas de Rodapé (Footnotes GFM)**: Ativação da extensão `extension.Footnote` do Goldmark para referências de notas no texto (`[^1]`) e seção de notas estruturada com links bidirecionais de retorno no final do documento.
- **Listas de Definição (Definition Lists)**: Ativação da extensão `extension.DefinitionList` para tags semânticas `<dl>`, `<dt>` e `<dd>` com estilização moderna.
- **Tipografia Aprimorada (Typographer)**: Ativação de `extension.Typographer` para pontuação inteligente (travessões `—`, meias-riscas `–`, reticências `…` e aspas tipográficas).
- **Fórmulas Matemáticas e LaTeX (Math Rendering)**: Suporte a fórmulas matemáticas inline (`$E=mc^2$`) e blocos de equações (`$$ ... $$`) com renderização vetorial ultrarrápida via KaTeX.
- **Emoji Shortcodes**: Reconhecimento de códigos de emoji populares do GitHub (ex: `:rocket:`, `:white_check_mark:`, `:bulb:`, `:sparkles:`).
- **Botão de Cópia em Blocos de Código (Code Copy Action)**: Inclusão de botão interativo e discreto de cópia para a área de transferência em todos os blocos de código formatados com feedback visual imediato.
- **Documentação Visual com Prints de Tela (README.md)**: Atualização do `README.md` com prévias visuais em alta definição e mockups de tela (com dados fictícios ilustrativos) apresentando a interface do Markdown Viewer, suporte a temas (Dark/Light), alertas GitHub, diagramas Mermaid, fórmulas e menu lateral redimensionável.

## Capabilities

### New Capabilities
<!-- Nenhuma nova capacidade raiz necessária; expande os requisitos existentes de fundação e renderização -->

### Modified Capabilities
- `project-foundation`: Modifica e expande os requisitos de renderização Markdown para incluir GitHub Alerts/Callouts, Notas de Rodapé (Footnotes), Fórmulas Matemáticas KaTeX, Listas de Definição, Tipografia Inteligente, Emojis, Botão de Cópia em blocos de código e prévias visuais no README.

## Impact

- **Backend (Go)**:
  - `internal/adapters/out/renderer/goldmark.go`: Registro e configuração das extensões Goldmark (`extension.Footnote`, `extension.DefinitionList`, `extension.Typographer`), parser/transformer para GitHub Alerts e suporte a delimitadores de fórmulas matemáticas.
  - `internal/adapters/out/renderer/goldmark_test.go`: Testes unitários para todas as novas extensões e renderizações garantindo cobertura inegociável >= 80%.
- **Frontend (HTML/CSS/JS)**:
  - `web/templates/base.html`: Inclusão dos assets KaTeX para fórmulas matemáticas.
  - `web/static/css/style.css`: Estilos para GitHub Alerts (`.markdown-alert`, `.markdown-alert-note`, etc.), notas de rodapé (`.footnotes`), listas de definição, fórmulas matemáticas e botão de cópia de código.
  - `web/static/js/app.js`: Lógica de clique para cópia de blocos de código e renderização de expressões matemáticas KaTeX.
- **Documentação e Assets (`docs/screenshots/` e `README.md`)**:
  - Criação de mockups/prints visuais vetoriais em `docs/screenshots/` exibindo as funcionalidades e atualização do `README.md` com as imagens e guia de novos plugins.
