## MODIFIED Requirements

### Requirement: Servir e renderizar documentos Markdown locais com interface minimalista (Markdown Viewer)
O sistema SHALL processar arquivos Markdown no diretório raiz fornecido, convertê-los em HTML semântico com suporte a GitHub Flavored Markdown (GFM), GitHub Alerts (Admonitions), Notas de Rodapé (Footnotes), Fórmulas Matemáticas LaTeX, Listas de Definição, Tipografia Inteligente e botão de cópia de código, exibindo-os em uma interface web minimalista inspirada no md-reader sob a marca "Markdown Viewer", com ícone vetorial exclusivo, favicon na aba do navegador, alternância de temas (Dark/Light Mode), menu lateral retrátil full-height e divisor maleável de redimensionamento manual (*resizable sidebar splitter*) com persistência de largura no navegador.

#### Scenario: Visualização de arquivo Markdown principal com branding Markdown Viewer e Ícone
- **GIVEN** que o servidor `md_server` está em execução apontando para uma pasta com o arquivo `README.md`
- **WHEN** o usuário acessa a raiz `http://localhost:8080/` no navegador
- **THEN** o sistema SHALL renderizar o conteúdo do `README.md` com tipografia moderna, exibir no cabeçalho a marca "Markdown Viewer" acompanhada do ícone vetorial e carregar o favicon na aba do navegador.

#### Scenario: Tipografia de leitura otimizada para prosa e documentação Markdown
- **GIVEN** que o usuário está visualizando qualquer documento Markdown renderizado
- **WHEN** o navegador renderiza o texto, títulos, parágrafos e listas
- **THEN** o sistema SHALL utilizar a pilha tipográfica sem-serifa moderna recomendada (`-apple-system, BlinkMacSystemFont, "Segoe UI", "Noto Sans", Helvetica, Arial, sans-serif`), proporcionando nitidez visual, excelente legibilidade e suporte completo a emojis do sistema.

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

---

### Requirement: Destaque de sintaxe de código-fonte colorido (Syntax Highlighting)
O sistema DEVE aplicar colorização e destaque de sintaxe em blocos de código-fonte formatados com especificadores de linguagem (Go, Python, JavaScript, TypeScript, JSON, YAML, HTML, Shell, SQL, etc.), utilizando a hierarquia de fontes monoespaçadas mais recomendada para desenvolvedores.

#### Scenario: Destaque de código com linguagem reconhecida
- **GIVEN** que o documento Markdown possui um bloco de código Go cercado por ```go
- **WHEN** a página é renderizada
- **THEN** o sistema SHALL colorir palavras-chave, strings, funções e comentários com classes de estilo de alto contraste e legibilidade, respeitando o tema visual selecionado (Dark/Light).

#### Scenario: Fallback para blocos de código genéricos sem linguagem declarada
- **GIVEN** que o documento Markdown possui um bloco de código cercado por ``` sem identificador de linguagem
- **WHEN** a página é renderizada
- **THEN** o sistema SHALL exibir o bloco com fonte monoespaçada, fundo destacado e formatação preservada sem quebras de layout.

#### Scenario: Tipografia monoespaçada otimizada para blocos de código e scripts
- **GIVEN** que blocos de código, tags `<code>`, scripts ou snippets de terminal são exibidos na interface
- **WHEN** os elementos são desenhados no navegador
- **THEN** o sistema SHALL aplicar a pilha tipográfica monoespaçada otimizada (`"JetBrains Mono", "Fira Code", "Cascadia Code", "SF Mono", Consolas, "Roboto Mono", monospace`), garantindo alinhamento de colunas, distinção nítida de caracteres (`0`/`O`, `1`/`l`/`I`) e renderização nítida.

---

### Requirement: Renderização de diagramas Mermaid em tempo real
O sistema SHALL identificar blocos de código com identificador de linguagem `mermaid`, processá-los no parser e renderizá-los dinamicamente como diagramas vetoriais (SVG) estilizados na interface web, integrando-os com a alternância de temas (Dark/Light), fornecendo fallback seguro para sintaxes incorretas e suportando navegação interativa fluida com controles de Zoom (+ / - / ↺) e Pan por arrasto com mouse ou toque.

#### Scenario: Renderização de fluxo ou diagrama Mermaid válido
- **GIVEN** que um arquivo Markdown contém um bloco de código cercado com ```mermaid (ex: fluxo flowchart TD, diagrama de sequência ou grafo de classes)
- **WHEN** o usuário carrega a página correspondente no navegador
- **THEN** o sistema SHALL processar o bloco e exibir o diagrama visual SVG estilizado e interativo em vez do texto bruto.

#### Scenario: Tratamento de erro em diagrama Mermaid com sintaxe inválida
- **GIVEN** que um arquivo Markdown contém um bloco ```mermaid com erro de digitação ou sintaxe inválida
- **WHEN** o navegador tenta processar e renderizar o diagrama
- **THEN** o sistema SHALL capturar o erro, exibir um alerta visual claro no local do diagrama indicando a falha de sintaxe e disponibilizar um botão para inspecionar o código-fonte original.

#### Scenario: Navegação interativa com Zoom e Pan em diagramas Mermaid
- **GIVEN** que um diagrama Mermaid renderizado está visível na tela
- **WHEN** o usuário clica nos botões de ampliação (+) ou redução (-) ou utiliza a roda do mouse sobre o diagrama
- **THEN** o sistema SHALL aplicar escala proporcional suave no diagrama SVG e permitir que o usuário arraste (*pan*) o diagrama livremente com o cursor do mouse em formato de mão (`grab`/`grabbing`) ou gesto de toque.

#### Scenario: Redefinição da visualização do diagrama
- **GIVEN** que o usuário alterou o nível de zoom ou arrastou a posição de um diagrama Mermaid
- **WHEN** o usuário clica no botão de reset (↺) na barra de ferramentas flutuante do diagrama
- **THEN** o sistema SHALL redefinir o zoom para 100% e centralizar novamente o diagrama no container original.

---

### Requirement: Compartilhamento facilitado e acesso móvel via QR Code
O sistema SHALL disponibilizar no cabeçalho da interface um botão que abre um modal interativo contendo o **QR Code** gerado com o endereço IP acessível na rede local (LAN), exibindo o botão exclusivamente quando a aplicação for iniciada com o argumento `--lan` ativo.

#### Scenario: Exibição do modal de QR Code para escaneamento móvel
- **GIVEN** que o usuário está visualizando qualquer documento Markdown na aplicação executada com a flag `--lan` em uma máquina com interface de rede local ativa (ex: IP `192.168.1.50` na porta `8080`)
- **WHEN** o usuário clica no botão "QR Code" no cabeçalho
- **THEN** o sistema SHALL abrir um modal centralizado exibindo o QR Code gerado contendo o endereço IP da rede local (ex: `http://192.168.1.50:8080/documento.md`), com o campo de texto contendo a mesma URL de rede e botão para copiar para a área de transferência.

#### Scenario: Otimização de escaneabilidade com quiet zone branca e tamanho ampliado
- **GIVEN** que o modal de QR Code é aberto em qualquer tema visual (incluindo Dark Mode)
- **WHEN** o QR Code é renderizado na tela
- **THEN** o sistema SHALL exibir o código em tamanho ampliado (240x240 pixels) dentro de uma moldura branca isolada (*quiet zone*) com contraste máximo (21:1) e nível de correção de erros calibrado (`CorrectLevel.M`), assegurando reconhecimento óptico instantâneo pelas câmeras de smartphones.

#### Scenario: Exibição em destaque do endereço IP para conexão direta
- **GIVEN** que o usuário abriu o modal de QR Code
- **WHEN** o modal é renderizado
- **THEN** o sistema SHALL exibir o endereço IP e a porta em badge de texto em destaque e campo legível com botão "Copiar", permitindo conexão manual rápida sem esforço.

#### Scenario: Ocultação do botão de QR Code no modo padrão local (seguro)
- **GIVEN** que o `md_server` é iniciado sem o argumento `--lan` (padrão local seguro)
- **WHEN** o usuário acessa a aplicação no navegador em `http://localhost:8080`
- **THEN** o sistema SHALL omitir/ocultar o botão "QR Code" no cabeçalho da interface web e manter o servidor escutando apenas em `127.0.0.1`.

#### Scenario: Fallback para URL local quando nenhuma interface externa estiver disponível
- **GIVEN** que a máquina servidora está desconectada de qualquer rede local (offline/sem adaptador de rede externo ativo)
- **WHEN** o usuário abre o modal de QR Code
- **THEN** o sistema SHALL utilizar a URL de host local atual (`window.location.origin` ou `http://localhost:<porta>`) como fallback seguro para a geração do QR Code e cópia de link.

#### Scenario: Fechamento do modal de QR Code por clique externo ou botão
- **GIVEN** que o modal de QR Code está visível na tela
- **WHEN** o usuário clica no botão de fechar (X) ou em qualquer área externa ao modal
- **THEN** o sistema SHALL ocultar o modal suavemente e retornar o foco para a leitura do documento.

---

### Requirement: Inicialização por Duplo Clique no Windows e CLI Multiplataforma
O sistema SHALL suportar execução por duplo clique em ambientes Windows sem requisições de argumentos mantendo escuta local segura por padrão (`127.0.0.1`), e fornecer uma interface de linha de comando completa para Windows e Linux/AMD64 com suporte a parâmetros (`--dir`, `--port`, `--open`, `--lan`).

#### Scenario: Inicialização por duplo clique no Windows (Zero-Config)
- **GIVEN** que o usuário executa o binário `md_server.exe` no Windows através de duplo clique no Windows Explorer
- **WHEN** a aplicação inicia sem flags de linha de comando
- **THEN** o sistema SHALL assumir o diretório corrente como raiz, alocar a porta padrão `8080` (ou a próxima porta disponível), iniciar o servidor HTTP escutando em loopback local (`127.0.0.1`), exibir a URL de acesso Local no console e disparar automaticamente a abertura da URL no navegador padrão do sistema operacional.

#### Scenario: Execução via CLI especificando diretório raiz e porta personalizada
- **GIVEN** que o usuário executa `md_server --dir ./documentacao --port 9090 --open=false`
- **WHEN** o comando é processado
- **THEN** o sistema SHALL servir os arquivos localizados em `./documentacao` na porta `9090`, não abrir o navegador automaticamente e exibir no console a mensagem de servidor pronto com a URL local `http://localhost:9090`.

#### Scenario: Execução via CLI com exposição em rede local
- **GIVEN** que o usuário executa `md_server --dir ./docs --port 9090 --lan`
- **WHEN** o comando é processado
- **THEN** o sistema SHALL vincular a escuta em todas as interfaces (`0.0.0.0:9090`), habilitar o QR Code na interface e exibir no console as URLs de acesso Local (`http://localhost:9090`) e de Rede Local (`http://<lan-ip>:9090`).

#### Scenario: Inspeção de versão da aplicação via comando CLI
- **GIVEN** que o binário foi compilado com metadados de versão via ldflags
- **WHEN** o usuário executa `md_server version` ou `md_server --version`
- **THEN** o sistema SHALL exibir a versão semântica oficial (ou `dev`), o hash do commit e a data de compilação no formato padronizado.
