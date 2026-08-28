## Why

O `md_server` atualmente oferece suporte básico ao GitHub Flavored Markdown (tabelas, task lists, strikethrough, autolinks) e diagramas Mermaid. No entanto, documentações técnicas modernas, repositórios do GitHub e bases de conhecimento ricas utilizam extensivamente recursos adicionais como **GitHub Alerts / Callouts** (`[!NOTE]`, `[!TIP]`, `[!WARNING]`, `[!IMPORTANT]`, `[!CAUTION]`), **Notas de Rodapé** (*Footnotes*), **Fórmulas Matemáticas / LaTeX**, **Listas de Definição**, **Tipografia Inteligente**, **Emoji Shortcodes**, **Botão de Cópia em Blocos de Código**, **Exploração Interativa de Diagramas com Zoom e Pan** e **Stack Tipográfica Recomendada para Leitura e Códigos**.

Uma tipografia cuidadosamente selecionada faz toda a diferença na experiência de leitura de documentos técnicos: uma stack de fontes sem-serifa moderna para prosa garante máxima clareza e ritmo visual, enquanto uma hierarquia de fontes monoespaçadas consagradas (JetBrains Mono, Fira Code, Cascadia Code, SF Mono, Consolas) oferece legibilidade impecável para blocos de código-fonte, scripts, Chroma highlighting e delimitadores de diagramas.

Além disso, diagramas complexos de arquitetura e fluxogramas extensos frequentemente ultrapassam a largura padrão da viewport ou contêm detalhes minuciosos que exigem ampliação e navegação fluida por arrasto. Adicionalmente, em ambientes corporativos e redes restritas, políticas de segurança exigem que servidores locais **não fiquem expostos na rede local (LAN)** por padrão, escutando em loopback local (`127.0.0.1`) e ativando a exposição externa e QR Code sob demanda com `--lan` / `-l`.

## What Changes

- **GitHub Alerts / Callouts (Admonitions)**: Suporte completo aos 5 tipos de avisos nativos do GitHub (`> [!NOTE]`, `> [!TIP]`, `> [!IMPORTANT]`, `> [!WARNING]`, `> [!CAUTION]`), renderizados com bordas coloridas, ícones SVG específicos, tipografia destacada e adaptação aos temas Claro e Escuro.
- **Notas de Rodapé (Footnotes GFM)**: Ativação da extensão `extension.Footnote` do Goldmark para referências de notas no texto (`[^1]`) e seção de notas estruturada com links bidirecionais de retorno no final do documento.
- **Listas de Definição (Definition Lists)**: Ativação da extensão `extension.DefinitionList` para tags semânticas `<dl>`, `<dt>` e `<dd>` com estilização moderna.
- **Tipografia Aprimorada (Typographer)**: Ativação de `extension.Typographer` para pontuação inteligente (travessões `—`, meias-riscas `–`, reticências `…` e aspas tipográficas).
- **Stack Tipográfica Otimizada para Markdown e Código**: Configuração da hierarquia de fontes mais recomendada para documentação técnica e Markdown: stack de sistema moderna (`-apple-system, BlinkMacSystemFont, "Segoe UI", "Noto Sans", Helvetica, Arial, sans-serif`) para texto e títulos, e pilha monoespaçada de desenvolvedor (`"JetBrains Mono", "Fira Code", "Cascadia Code", "SF Mono", Consolas, "Roboto Mono", monospace`) para códigos, scripts e Chroma syntax highlighting.
- **Fórmulas Matemáticas e LaTeX (Math Rendering)**: Suporte a fórmulas matemáticas inline (`$E=mc^2$`) e blocos de equações (`$$ ... $$`) com renderização vetorial ultrarrápida via KaTeX.
- **Emoji Shortcodes**: Reconhecimento de códigos de emoji populares do GitHub (ex: `:rocket:`, `:white_check_mark:`, `:bulb:`, `:sparkles:`).
- **Botão de Cópia em Blocos de Código (Code Copy Action)**: Inclusão de botão interativo e discreto de cópia para a área de transferência em todos os blocos de código formatados com feedback visual imediato.
- **Zoom e Pan Interativo para Diagramas (Mermaid Navigation)**: Controles flutuantes de ampliação (+), redução (-), reset (↺) e navegação por clique e arraste (*drag-and-pan*) intuitivo em todos os diagramas Mermaid gerados, com suporte a roda do mouse (wheel/Ctrl+wheel) e toque móvel.
- **Controle Seguro de Exposição de Rede Local (`--lan`)**: Inclusão da flag CLI `--lan` (shorthand `-l`, padrão `false`). Por padrão, o servidor escuta em `127.0.0.1` e o botão de QR Code fica oculto; ao passar `--lan`, o servidor escuta em `0.0.0.0`, exibe a URL LAN no console e habilita o botão de QR Code no cabeçalho.
- **Motor Canônico de QR Code (Conformidade ISO/IEC 18004) e Acesso Móvel**: Substituição da implementação de QR Code por um motor canônico universal e matematicamente exato (com codificação Reed-Solomon autêntica, cálculo polinomial de correção de erros e máscaras de dados padrão ISO/IEC 18004), renderizado em 240x240px com moldura branca de alto contraste (*quiet zone 21:1*), nível de correção `CorrectLevel.M` (15%) e badge de IP para digitação manual, garantindo que qualquer câmera de smartphone (iOS, Android, Google Lens) decodifique o endereço instantaneamente.
- **Adaptação Responsiva de Título no Celular ("MD Viewer")**: Ajuste do título da marca no cabeçalho e da tag `<title>` em dispositivos móveis e telas estreitas ($\le 768\text{px}$) para **"MD Viewer"**, evitando quebra de linha ou sobreposição com os botões de ação e preservando a ergonomia móvel.
- **Documentação Visual com Prints de Tela (README.md)**: Atualização do `README.md` com prévias visuais em alta definição e mockups de tela (com dados fictícios ilustrativos) apresentando a interface do Markdown Viewer, suporte a temas (Dark/Light), alertas GitHub, diagramas Mermaid com controles de zoom, fórmulas e menu lateral redimensionável.

## Capabilities

### New Capabilities
<!-- Nenhuma nova capacidade raiz necessária; expande os requisitos existentes de fundação e renderização -->

### Modified Capabilities
- `project-foundation`: Modifica e expande os requisitos de renderização Markdown para incluir GitHub Alerts/Callouts, Notas de Rodapé (Footnotes), Fórmulas Matemáticas KaTeX, Listas de Definição, Tipografia Inteligente, stack tipográfica moderna para leitura e código-fonte/scripts, Emojis, Botão de Cópia em blocos de código, controles de Zoom e Pan em diagramas Mermaid, flag de controle seguro de exposição de rede local (`--lan`) com ativação condicional do QR Code, motor canônico de QR Code ISO/IEC 18004 de alta escaneabilidade, adaptação de título compacto ("MD Viewer") no celular e prévias visuais no README.

## Impact

- **Backend (Go)**:
  - `internal/core/domain/config.go`: Adição do campo `ExposeLAN bool` na entidade de configuração.
  - `internal/adapters/in/cli/root.go`: Adição da flag `--lan` / `-l` (default: `false`).
  - `internal/adapters/in/http/server.go` e `handlers.go`: Escuta condicional (`127.0.0.1` vs `0.0.0.0`) e propagação de `ShowQRCode` no `PageData`.
  - `cmd/md_server/main.go`: Exibição de URL LAN no banner apenas quando `--lan` for habilitado.
  - `internal/adapters/out/renderer/goldmark.go`: Extensões Goldmark e GitHub Alerts.
- **Frontend (HTML/CSS/JS)**:
  - `web/templates/base.html`: Estrutura responsiva para o título da marca com variantes Desktop ("Markdown Viewer") e Mobile ("MD Viewer").
  - `web/static/css/style.css`: Media query para alternância fluida entre `.brand-full` e `.brand-short` em telas móveis e ajuste de padding/tamanho.
  - `web/static/js/qrcode.min.js`: Motor canônico standalone autêntico baseado na especificação ISO/IEC 18004 com Reed-Solomon e SVG nativo.
  - `web/static/js/app.js`: Geração do QR Code em 240x240px com `CorrectLevel.M`, normalização de URL simplificada e manipulador de cópia rápida.
- **Documentação e Assets (`docs/screenshots/` e `README.md`)**:
  - Mockups visuais vetoriais em `docs/screenshots/` e atualização do `README.md` documentando a nova flag `--lan` e a navegação em diagramas Mermaid.
