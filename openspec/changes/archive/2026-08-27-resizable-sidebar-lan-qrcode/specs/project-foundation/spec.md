## MODIFIED Requirements

### Requirement: Servir e renderizar documentos Markdown locais com interface minimalista (Markdown Viewer)
O sistema SHALL processar arquivos Markdown no diretório raiz fornecido, convertê-los em HTML semântico com suporte a GitHub Flavored Markdown (GFM) e exibi-los em uma interface web minimalista inspirada no md-reader sob a marca "Markdown Viewer", com ícone vetorial exclusivo, favicon na aba do navegador, alternância de temas (Dark/Light Mode), menu lateral retrátil full-height e divisor maleável de redimensionamento manual (*resizable sidebar splitter*) com persistência de largura no navegador.

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

---

### Requirement: Compartilhamento facilitado e acesso móvel via QR Code
O sistema SHALL disponibilizar no cabeçalho da interface um botão que abre um modal interativo contendo o **QR Code** gerado com o endereço IP acessível na rede local (LAN), permitindo que smartphones, tablets e outros dispositivos na mesma rede Wi-Fi/Ethernet acessem a documentação instantaneamente sem restrições de loopback (`127.0.0.1` ou `localhost`).

#### Scenario: Exibição do modal de QR Code para escaneamento móvel
- **GIVEN** que o usuário está visualizando qualquer documento Markdown na aplicação executada em uma máquina com interface de rede local ativa (ex: IP `192.168.1.50` na porta `8080`)
- **WHEN** o usuário clica no botão "QR Code" no cabeçalho
- **THEN** o sistema SHALL abrir um modal centralizado exibindo o QR Code gerado contendo o endereço IP da rede local (ex: `http://192.168.1.50:8080/documento.md`), com o campo de texto contendo a mesma URL de rede e botão para copiar para a área de transferência.

#### Scenario: Fallback para URL local quando nenhuma interface externa estiver disponível
- **GIVEN** que a máquina servidora está desconectada de qualquer rede local (offline/sem adaptador de rede externo ativo)
- **WHEN** o usuário abre o modal de QR Code
- **THEN** o sistema SHALL utilizar a URL de host local atual (`window.location.origin` ou `http://localhost:<porta>`) como fallback seguro para a geração do QR Code e cópia de link.

#### Scenario: Fechamento do modal de QR Code por clique externo ou botão
- **GIVEN** que o modal de QR Code está visível na tela
- **WHEN** o usuário clica no botão de fechar (X) ou em qualquer área externa ao modal
- **THEN** o sistema SHALL ocultar o modal suavemente e retornar o foco para a leitura do documento.

---

### Requirement: Renderização de diagramas Mermaid em tempo real
O sistema SHALL identificar blocos de código com identificador de linguagem `mermaid`, processá-los no parser e renderizá-los dinamicamente como diagramas vetoriais (SVG) estilizados na interface web, integrando-os com a alternância de temas (Dark/Light) e fornecendo fallback seguro para sintaxes incorretas.

#### Scenario: Renderização de fluxo ou diagrama Mermaid válido
- **GIVEN** que um arquivo Markdown contém um bloco de código cercado com ```mermaid (ex: fluxo flowchart TD, diagrama de sequência ou grafo de classes)
- **WHEN** o usuário carrega a página correspondente no navegador
- **THEN** o sistema SHALL processar o bloco e exibir o diagrama visual SVG estilizado e interativo em vez do texto bruto.

#### Scenario: Tratamento de erro em diagrama Mermaid com sintaxe inválida
- **GIVEN** que um arquivo Markdown contém um bloco ```mermaid com erro de digitação ou sintaxe inválida
- **WHEN** o navegador tenta processar e renderizar o diagrama
- **THEN** o sistema SHALL capturar o erro, exibir um alerta visual claro no local do diagrama indicando a falha de sintaxe e disponibilizar um botão para inspecionar o código-fonte original.

---

### Requirement: Inicialização por Duplo Clique no Windows e CLI Multiplataforma
O sistema SHALL suportar execução por duplo clique em ambientes Windows sem requisições de argumentos, vincular a escuta HTTP em todas as interfaces de rede (`0.0.0.0`) para possibilitar acesso local e via LAN, e fornecer uma interface de linha de comando completa para Windows e Linux/AMD64 exibindo as URLs local e de rede.

#### Scenario: Inicialização por duplo clique no Windows (Zero-Config)
- **GIVEN** que o usuário executa o binário `md_server.exe` no Windows através de duplo clique no Windows Explorer
- **WHEN** a aplicação inicia sem flags de linha de comando
- **THEN** o sistema SHALL assumir o diretório corrente como raiz, alocar a porta padrão `8080` (ou a próxima porta disponível), iniciar o servidor HTTP escutando em todas as interfaces de rede (`0.0.0.0`), exibir as URLs de acesso Local e de Rede Local no console e disparar automaticamente a abertura da URL no navegador padrão do sistema operacional.

#### Scenario: Execução via CLI especificando diretório raiz e porta personalizada
- **GIVEN** que o usuário executa `md_server --dir ./documentacao --port 9090 --open=false`
- **WHEN** o comando é processado
- **THEN** o sistema SHALL servir os arquivos localizados em `./documentacao` na porta `9090`, não abrir o navegador automaticamente e exibir no console as mensagens de servidor pronto com a URL local `http://localhost:9090` e a URL de rede local `http://<lan-ip>:9090`.

#### Scenario: Inspeção de versão da aplicação via comando CLI
- **GIVEN** que o binário foi compilado com metadados de versão via ldflags
- **WHEN** o usuário executa `md_server version` ou `md_server --version`
- **THEN** o sistema SHALL exibir a versão semântica oficial (ou `dev`), o hash do commit e a data de compilação no formato padronizado.
