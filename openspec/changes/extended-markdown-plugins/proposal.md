## Why

O `md_server` atualmente oferece suporte básico ao GitHub Flavored Markdown (tabelas, task lists, strikethrough, autolinks) e diagramas Mermaid. No entanto, documentações técnicas modernas, repositórios do GitHub e bases de conhecimento ricas utilizam extensivamente recursos adicionais como **GitHub Alerts / Callouts** (`[!NOTE]`, `[!TIP]`, `[!WARNING]`, `[!IMPORTANT]`, `[!CAUTION]`), **Notas de Rodapé** (*Footnotes*), **Fórmulas Matemáticas / LaTeX**, **Listas de Definição**, **Tipografia Inteligente**, **Emoji Shortcodes** e **Botão de Cópia em Blocos de Código**.

Além disso, em ambientes corporativos e redes restritas, políticas de segurança exigem que servidores locais **não fiquem expostos na rede local (LAN)** por padrão. O comportamento seguro deve ser escutar estritamente em loopback local (`127.0.0.1`) e ocultar o botão de QR Code, permitindo a exposição em rede (`0.0.0.0`) e ativação do QR Code exclusivamente sob demanda via argumento explícito de linha de comando (`--lan` / `-l`).

## What Changes

- **GitHub Alerts / Callouts (Admonitions)**: Suporte completo aos 5 tipos de avisos nativos do GitHub (`> [!NOTE]`, `> [!TIP]`, `> [!IMPORTANT]`, `> [!WARNING]`, `> [!CAUTION]`), renderizados com bordas coloridas, ícones SVG específicos, tipografia destacada e adaptação aos temas Claro e Escuro.
- **Notas de Rodapé (Footnotes GFM)**: Ativação da extensão `extension.Footnote` do Goldmark para referências de notas no texto (`[^1]`) e seção de notas estruturada com links bidirecionais de retorno no final do documento.
- **Listas de Definição (Definition Lists)**: Ativação da extensão `extension.DefinitionList` para tags semânticas `<dl>`, `<dt>` e `<dd>` com estilização moderna.
- **Tipografia Aprimorada (Typographer)**: Ativação de `extension.Typographer` para pontuação inteligente (travessões `—`, meias-riscas `–`, reticências `…` e aspas tipográficas).
- **Fórmulas Matemáticas e LaTeX (Math Rendering)**: Suporte a fórmulas matemáticas inline (`$E=mc^2$`) e blocos de equações (`$$ ... $$`) com renderização vetorial ultrarrápida via KaTeX.
- **Emoji Shortcodes**: Reconhecimento de códigos de emoji populares do GitHub (ex: `:rocket:`, `:white_check_mark:`, `:bulb:`, `:sparkles:`).
- **Botão de Cópia em Blocos de Código (Code Copy Action)**: Inclusão de botão interativo e discreto de cópia para a área de transferência em todos os blocos de código formatados com feedback visual imediato.
- **Controle Seguro de Exposição de Rede Local (`--lan`)**: Inclusão da flag CLI `--lan` (shorthand `-l`, padrão `false`). Por padrão, o servidor escuta em `127.0.0.1` e o botão de QR Code fica oculto; ao passar `--lan`, o servidor escuta em `0.0.0.0`, exibe a URL LAN no console e habilita o botão de QR Code no cabeçalho.
- **Documentação Visual com Prints de Tela (README.md)**: Atualização do `README.md` com prévias visuais em alta definição e mockups de tela (com dados fictícios ilustrativos) apresentando a interface do Markdown Viewer, suporte a temas (Dark/Light), alertas GitHub, diagramas Mermaid, fórmulas e menu lateral redimensionável.

## Capabilities

### New Capabilities
<!-- Nenhuma nova capacidade raiz necessária; expande os requisitos existentes de fundação e renderização -->

### Modified Capabilities
- `project-foundation`: Modifica e expande os requisitos de renderização Markdown para incluir GitHub Alerts/Callouts, Notas de Rodapé (Footnotes), Fórmulas Matemáticas KaTeX, Listas de Definição, Tipografia Inteligente, Emojis, Botão de Cópia em blocos de código, flag de controle seguro de exposição de rede local (`--lan`) com ativação condicional do QR Code e prévias visuais no README.

## Impact

- **Backend (Go)**:
  - `internal/core/domain/config.go`: Adição do campo `ExposeLAN bool` na entidade de configuração.
  - `internal/adapters/in/cli/root.go`: Adição da flag `--lan` / `-l` (default: `false`).
  - `internal/adapters/in/http/server.go` e `handlers.go`: Escuta condicional (`127.0.0.1` vs `0.0.0.0`) e propagação de `ShowQRCode` no `PageData`.
  - `cmd/md_server/main.go`: Exibição de URL LAN no banner apenas quando `--lan` for habilitado.
  - `internal/adapters/out/renderer/goldmark.go`: Extensões Goldmark e GitHub Alerts.
- **Frontend (HTML/CSS/JS)**:
  - `web/templates/base.html`: Renderização condicional do botão de QR Code (`{{if .ShowQRCode}}`).
  - `web/static/css/style.css`: Estilos de alerts, footnotes, KaTeX e copy button.
  - `web/static/js/app.js`: Cópia de código, KaTeX e inicialização segura de QR Code.
- **Documentação e Assets (`docs/screenshots/` e `README.md`)**:
  - Mockups visuais vetoriais em `docs/screenshots/` e atualização do `README.md` documentando a nova flag `--lan`.
