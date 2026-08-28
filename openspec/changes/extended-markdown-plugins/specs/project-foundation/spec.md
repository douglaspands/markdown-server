## MODIFIED Requirements

### Requirement: Servir e renderizar documentos Markdown locais com interface minimalista (Markdown Viewer)
O sistema SHALL processar arquivos Markdown no diretório raiz fornecido, convertê-los em HTML semântico com suporte a GitHub Flavored Markdown (GFM), GitHub Alerts (Admonitions), Notas de Rodapé (Footnotes), Fórmulas Matemáticas LaTeX, Listas de Definição, Tipografia Inteligente e botão de cópia de código, exibindo-os em uma interface web minimalista inspirada no md-reader sob a marca "Markdown Viewer", com ícone vetorial exclusivo, favicon na aba do navegador, alternância de temas (Dark/Light Mode), menu lateral retrátil full-height e divisor maleável de redimensionamento manual (*resizable sidebar splitter*) com persistência de largura no navegador.

#### Scenario: Visualização de arquivo Markdown principal com branding Markdown Viewer e Ícone
- **GIVEN** que o servidor `md_server` está em execução apontando para uma pasta com o arquivo `README.md`
- **WHEN** o usuário acessa a raiz `http://localhost:8080/` no navegador
- **THEN** o sistema SHALL renderizar o conteúdo do `README.md` com tipografia moderna, exibir no cabeçalho a marca "Markdown Viewer" acompanhada do ícone vetorial e carregar o favicon na aba do navegador.

#### Scenario: Atendimento da rota de Favicon do navegador
- **GIVEN** que o servidor `md_server` está ativo
- **WHEN** o navegador requisita `/favicon.ico` ou `/static/img/icon.svg`
- **THEN** o sistema SHALL responder com status HTTP 200 e os bytes correspondentes do ícone com cabeçalho `Content-Type` adequado.

#### Scenario: Recolhimento e expansão do menu lateral retrátil
- **GIVEN** que a página está aberta no navegador com o menu lateral expandido
- **WHEN** o usuário clica no botão de alternância (toggle sidebar) no cabeçalho
- **THEN** a barra lateral SHALL colapsar suavemente, expandindo a área de leitura para a largura total da janela e salvando o estado no navegador.

#### Scenario: Redimensionamento manual da barra lateral via divisor arrastável (Splitter Handle)
- **GIVEN** que a página está aberta em dispositivo desktop com a barra lateral expandida
- **WHEN** o usuário clica e arrasta o divisor lateral posicionado entre a barra e a área de leitura
- **THEN** a largura da barra lateral SHALL acompanhar o movimento do cursor do mouse em tempo real com cursor visual `col-resize`, respeitando a largura mínima de 180px e máxima de 600px (ou 50vw), salvando a largura escolhida no `localStorage` do navegador para recarregamentos futuros.

#### Scenario: Manutenção da altura contínua (100vh) da barra lateral sem cortes
- **GIVEN** um documento curto com apenas poucas linhas de texto ou uma pasta com poucos itens
- **WHEN** a página é renderizada na janela do navegador
- **THEN** a barra lateral, seu divisor de redimensionamento, seu plano de fundo e sua borda divisória SHALL ocupar 100% da altura da tela (`100vh`) até a base da janela, sem deixar espaços vazios truncados.

#### Scenario: Tratamento amigável para documento Markdown inexistente
- **GIVEN** que o servidor `md_server` está ativo
- **WHEN** o usuário requisita a URL de um arquivo que não existe no diretório (ex: `/documento-inexistente.md`)
- **THEN** o sistema SHALL responder com status HTTP 404 e exibir uma página visualmente integrada informando que o documento não foi encontrado e oferecendo link de retorno para a página inicial.

#### Scenario: Prevenção contra Directory Traversal e acesso fora da pasta raiz
- **GIVEN** que o servidor `md_server` está servindo a pasta `/meu-projeto/docs`
- **WHEN** uma requisição tenta acessar caminhos externos utilizando sequências de escape (ex: `/../../etc/passwd` ou `..\..\Windows\win.ini`)
- **THEN** o sistema SHALL sanitizar o caminho, rejeitar a requisição com status HTTP 403 (Forbidden) ou 400 (Bad Request) e impedir qualquer leitura fora do diretório raiz.

#### Scenario: Renderização de GitHub Alerts (Admonitions)
- **GIVEN** que o documento Markdown contém blocos de aviso no padrão GitHub (ex: `> [!NOTE]`, `> [!TIP]`, `> [!IMPORTANT]`, `> [!WARNING]`, `> [!CAUTION]`)
- **WHEN** a página é renderizada no navegador
- **THEN** o sistema SHALL converter cada bloco em um container estilizado com borda lateral com a cor correspondente, ícone visual SVG padronizado e título do alerta em destaque.

#### Scenario: Renderização de Notas de Rodapé (Footnotes GFM)
- **GIVEN** que o documento Markdown contém referências de notas de rodapé no corpo do texto (ex: `[^1]`) e a definição da nota (ex: `[^1]: Detalhes adicionais`)
- **WHEN** a página é renderizada
- **THEN** o sistema SHALL criar links numéricos sobrescritos clicáveis que direcionam para a seção inferior de notas de rodapé (`#fn:1`), incluindo links de retorno suave (`#fnref:1`).

#### Scenario: Renderização de Fórmulas Matemáticas e LaTeX
- **GIVEN** que o documento contém expressões matemáticas inline (ex: `$E=mc^2$`) ou em bloco fechado (ex: `$$\int_0^1 f(x)dx$$`)
- **WHEN** a página carrega
- **THEN** o sistema SHALL processar e exibir as equações renderizadas via KaTeX com tipografia matemática de alta definição.

#### Scenario: Renderização de Listas de Definição e Tipografia Inteligente
- **GIVEN** que o documento Markdown contém termos e definições (ex: `Termo\n: Definição`) e pontuações especiais (aspas retas, `---`, `...`)
- **WHEN** o conteúdo é convertido pelo parser
- **THEN** o sistema SHALL gerar tags semânticas `<dl>`, `<dt>` e `<dd>` e substituir automaticamente os caracteres por pontuação tipográfica inteligente (travessão `—`, reticências `…`, aspas curvas).

#### Scenario: Cópia rápida de código-fonte via botão interativo
- **GIVEN** que o usuário está visualizando qualquer bloco de código-fonte formatado com Chroma
- **WHEN** o usuário clica no botão "Copiar" posicionado no canto superior do bloco
- **THEN** o sistema SHALL copiar o código-fonte puro para a área de transferência do sistema operacional e exibir temporariamente a confirmação visual "Copiado!".

#### Scenario: Documentação visual no README com prévias de tela
- **GIVEN** que a aplicação possui interface gráfica minimalista e plugins avançados de renderização
- **WHEN** o usuário ou desenvolvedor lê o `README.md` no repositório
- **THEN** o documento SHALL exibir capturas de tela e mockups visuais ilustrativos com dados fictícios apresentando o visualizador em funcionamento, demonstrando o divisor de redimensionamento, os alertas do GitHub, diagramas Mermaid e o modal de QR Code.
