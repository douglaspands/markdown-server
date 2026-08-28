# Especificação de Requisitos e Cenários BDD: Fundação do md_server (Markdown Viewer)

## Purpose

Fornece um servidor web local e ferramenta de linha de comando (CLI) multiplataforma (Windows e Linux/AMD64) para renderização instantânea de diretórios com arquivos Markdown sob a identidade visual "Markdown Viewer", com ícone vetorial moderno e favicon, suporte a diagramas Mermaid, destaque de sintaxe Chroma, navegação fluida por links internos, menu lateral retrátil full-height (100vh), interface responsiva para smartphones e tablets, compartilhamento rápido via QR Code e inicialização facilitada por duplo clique com abertura automática do navegador.

## ADDED Requirements

### Requirement: Servir e renderizar documentos Markdown locais com interface minimalista (Markdown Viewer)
O sistema DEVE processar arquivos Markdown no diretório raiz fornecido, convertê-los em HTML semântico com suporte a GitHub Flavored Markdown (GFM) e exibi-los em uma interface web minimalista inspirada no md-reader sob a marca "Markdown Viewer", com ícone vetorial exclusivo, favicon na aba do navegador, alternância de temas (Dark/Light Mode) e menu lateral retrátil full-height.

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

#### Scenario: Manutenção da altura contínua (100vh) da barra lateral sem cortes
- **GIVEN** um documento curto com apenas poucas linhas de texto ou uma pasta com poucos itens
- **WHEN** a página é renderizada na janela do navegador
- **THEN** a barra lateral, seu plano de fundo e sua borda divisória SHALL ocupar 100% da altura da tela (`100vh`) até a base da janela, sem deixar espaços vazios truncados.

#### Scenario: Tratamento amigável para documento Markdown inexistente
- **GIVEN** que o servidor `md_server` está ativo
- **WHEN** o usuário requisita a URL de um arquivo que não existe no diretório (ex: `/documento-inexistente.md`)
- **THEN** o sistema SHALL responder com status HTTP 404 e exibir uma página visualmente integrada informando que o documento não foi encontrado e oferecendo link de retorno para a página inicial.

#### Scenario: Prevenção contra Directory Traversal e acesso fora da pasta raiz
- **GIVEN** que o servidor `md_server` está servindo a pasta `/meu-projeto/docs`
- **WHEN** uma requisição tenta acessar caminhos externos utilizando sequências de escape (ex: `/../../etc/passwd` ou `..\..\Windows\win.ini`)
- **THEN** o sistema SHALL sanitizar o caminho, rejeitar a requisição com status HTTP 403 (Forbidden) ou 400 (Bad Request) e impedir qualquer leitura fora do diretório raiz.

---

### Requirement: Responsividade e compatibilidade total com smartphones e tablets
O sistema DEVE adaptar sua interface dinamicamente para dispositivos móveis com telas de toque, transformando a navegação em uma experiência fluida de gaveta (drawer) e otimizando a densidade de leitura.

#### Scenario: Acesso à aplicação a partir de smartphone ou tablet
- **GIVEN** que o usuário acessa a aplicação através de um navegador móvel (largura de tela inferior a 768px)
- **WHEN** a página carrega
- **THEN** a barra lateral SHALL iniciar oculta em formato de gaveta deslizante (*drawer*), permitindo que o usuário abra o menu tocando no botão do cabeçalho e feche ao selecionar um documento ou tocar no fundo com overlay.

---

### Requirement: Compartilhamento facilitado e acesso móvel via QR Code
O sistema DEVE disponibilizar no cabeçalho da interface um botão que abre um modal interativo contendo o **QR Code** da URL do documento atual, permitindo que outros dispositivos na rede local acessem a documentação instantaneamente.

#### Scenario: Exibição do modal de QR Code para escaneamento móvel
- **GIVEN** que o usuário está visualizando qualquer documento Markdown na aplicação
- **WHEN** o usuário clica no botão "QR Code" no cabeçalho
- **THEN** o sistema SHALL abrir um modal centralizado exibindo o QR Code gerado da URL atual com a URL legível e botão para copiar ou fechar.

#### Scenario: Fechamento do modal de QR Code por clique externo ou botão
- **GIVEN** que o modal de QR Code está visível na tela
- **WHEN** o usuário clica no botão de fechar (X) ou em qualquer área externa ao modal
- **THEN** o sistema SHALL ocultar o modal suavemente e retornar o foco para a leitura do documento.

---

### Requirement: Renderização de diagramas Mermaid em tempo real
O sistema DEVE identificar blocos de código com identificador de linguagem `mermaid` e renderizá-los dinamicamente como diagramas vetoriais (SVG) na interface web.

#### Scenario: Renderização de fluxo ou diagrama Mermaid válido
- **GIVEN** que um arquivo Markdown contém um bloco de código cercado com ```mermaid (ex: fluxo flowchart TD ou diagrama de sequência)
- **WHEN** o usuário carrega a página correspondente no navegador
- **THEN** o sistema SHALL processar o bloco e exibir o diagrama visual SVG estilizado e interativo em vez do texto bruto.

#### Scenario: Tratamento de erro em diagrama Mermaid com sintaxe inválida
- **GIVEN** que um arquivo Markdown contém um bloco ```mermaid com erro de digitação ou sintaxe inválida
- **WHEN** o navegador tenta processar e renderizar o diagrama
- **THEN** o sistema SHALL capturar o erro, exibir um alerta visual claro no local do diagrama indicando a falha de sintaxe e disponibilizar um botão para inspecionar o código-fonte original.

---

### Requirement: Destaque de sintaxe de código-fonte colorido (Syntax Highlighting)
O sistema DEVE aplicar colorização e destaque de sintaxe em blocos de código-fonte formatados com especificadores de linguagem (Go, Python, JavaScript, TypeScript, JSON, YAML, HTML, Shell, SQL, etc.).

#### Scenario: Destaque de código com linguagem reconhecida
- **GIVEN** que o documento Markdown possui um bloco de código Go cercado por ```go
- **WHEN** a página é renderizada
- **THEN** o sistema SHALL colorir palavras-chave, strings, funções e comentários com classes de estilo de alto contraste e legibilidade, respeitando o tema visual selecionado (Dark/Light).

#### Scenario: Fallback para blocos de código genéricos sem linguagem declarada
- **GIVEN** que o documento Markdown possui um bloco de código cercado por ``` sem identificador de linguagem
- **WHEN** a página é renderizada
- **THEN** o sistema SHALL exibir o bloco com fonte monoespaçada, fundo destacado e formatação preservada sem quebras de layout.

---

### Requirement: Navegação fluida por links internos e árvore lateral de arquivos
O sistema DEVE permitir a navegação contínua entre arquivos Markdown locais através de cliques em links relativos e pela árvore de diretórios exibida na barra lateral (*Sidebar Tree*).

#### Scenario: Clique em link relativo apontando para outro arquivo Markdown
- **GIVEN** que o documento em exibição possui um link Markdown relativo `[Guia de Instalação](./instalacao.md)`
- **WHEN** o usuário clica no link na interface web
- **THEN** o sistema SHALL redirecionar para a rota correspondente e carregar o arquivo `instalacao.md` renderizado sem perda de contexto ou erro de rota.

#### Scenario: Navegação e ancoragem interna na mesma página
- **GIVEN** que o documento possui links de âncora para títulos internos (ex: `[Ir para Requisitos](#requisitos)`)
- **WHEN** o usuário clica na âncora
- **THEN** o navegador SHALL rolar suavemente até a seção correspondente sem recarregar a página.

---

### Requirement: Inicialização por Duplo Clique no Windows e CLI Multiplataforma
O sistema DEVE suportar execução por duplo clique em ambientes Windows sem requisições de argumentos e fornecer uma interface de linha de comando completa para Windows e Linux/AMD64.

#### Scenario: Inicialização por duplo clique no Windows (Zero-Config)
- **GIVEN** que o usuário executa o binário `md_server.exe` no Windows através de duplo clique no Windows Explorer
- **WHEN** a aplicação inicia sem flags de linha de comando
- **THEN** o sistema SHALL assumir o diretório corrente como raiz, alocar a porta padrão `8080` (ou a próxima porta disponível), iniciar o servidor HTTP e disparar automaticamente a abertura da URL no navegador padrão do sistema operacional.

#### Scenario: Execução via CLI especificando diretório raiz e porta personalizada
- **GIVEN** que o usuário executa `md_server --dir ./documentacao --port 9090 --open=false`
- **WHEN** o comando é processado
- **THEN** o sistema SHALL servir os arquivos localizados em `./documentacao` na porta `9090`, não abrir o navegador automaticamente e exibir no console a mensagem de servidor pronto com a URL `http://localhost:9090`.

#### Scenario: Inspeção de versão da aplicação via comando CLI
- **GIVEN** que o binário foi compilado com metadados de versão via ldflags
- **WHEN** o usuário executa `md_server version` ou `md_server --version`
- **THEN** o sistema SHALL exibir a versão semântica oficial (ou `dev`), o hash do commit e a data de compilação no formato padronizado.

---

### Requirement: Serviço de Domínio de Referência e Endpoint de Saúde (Health Check)
O sistema DEVE fornecer um serviço de domínio desacoplado para verificação de integridade operacional exposto em endpoint HTTP REST.

#### Scenario: Consulta de saúde operacional da aplicação
- **GIVEN** que o servidor `md_server` está em execução
- **WHEN** um cliente HTTP realiza uma requisição `GET /api/health`
- **THEN** o sistema SHALL responder com status HTTP 200 OK e corpo JSON contendo `{"status": "UP", "version": "<versao>", "uptime": "<tempo_ativo>"}`.

---

### Requirement: Encerramento Gracioso do Servidor (Graceful Shutdown)
O sistema DEVE capturar sinais do sistema operacional e encerrar suas atividades de forma graciosa sem abortar conexões pendentes.

#### Scenario: Recebimento de sinal de encerramento do sistema
- **GIVEN** que o servidor HTTP possui conexões ativas em processamento
- **WHEN** o processo recebe um sinal `SIGINT` (Ctrl+C) ou `SIGTERM`
- **THEN** o sistema SHALL parar de aceitar novas requisições, concluir as requisições em trânsito dentro do tempo limite de tolerância (timeout de 5 segundos), liberar a porta de rede e finalizar o processo com código de saída 0.
