## Why

Atualmente, a barra lateral do Markdown Viewer possui largura estática (280px), impossibilitando que usuários ajustem o espaço de acordo com nomes longos de arquivos ou árvores de diretórios profundas. Além disso, o recurso de compartilhamento e acesso via QR Code utiliza o endereço de loopback (`127.0.0.1` ou `localhost`), impedindo que smartphones, tablets e outros computadores conectados à mesma rede local (LAN / Wi-Fi) consigam carregar os documentos ao escanear o código ou clicar no link. Por fim, identificou-se que os diagramas Mermaid não estavam sendo renderizados visualmente como SVG devido a conflitos de inicialização assíncrona do script no frontend.

Tornar a barra lateral redimensionável melhora a ergonomia de leitura, a detecção de IP local viabiliza o uso imediato do QR Code em múltiplos dispositivos e a correção da renderização de diagramas Mermaid assegura a visualização precisa de fluxos, diagramas de sequência e arquitetura na documentação.

## What Changes

- **Barra Lateral Redimensionável (Draggable Resizer)**: Inclusão de um divisor arrastável (*splitter / resize handle*) entre a barra lateral e a área de conteúdo principal, permitindo ajuste manual da largura com o mouse, cursor visual indicativo (`col-resize`), limites mínimos (180px) e máximos (600px / 50vw), e persistência da largura selecionada no `localStorage`.
- **Servidor HTTP em Múltiplas Interfaces (0.0.0.0)**: Alteração do bind do servidor HTTP de `127.0.0.1` para todas as interfaces (`0.0.0.0` / `:port`), permitindo que clientes da rede local acessem a aplicação.
- **Detecção Automática do IP da Rede Local (LAN)**: Implementação de utilitário/serviço no backend para identificar o endereço IP IPv4 não-loopback da interface de rede ativa da máquina servidora.
- **QR Code e Compartilhamento com IP da Rede Local**: Atualização da geração do QR Code e do campo de URL copiável no modal para utilizar o IP da rede local (ex: `http://192.168.1.50:8080/caminho`), garantindo que smartphones e tablets acessem a página instantaneamente.
- **Correção da Renderização de Diagramas Mermaid**: Correção da biblioteca e rotina de inicialização no frontend (`app.js` e `mermaid.min.js`), garantindo que blocos ```mermaid sejam convertidos em diagramas SVG vetoriais renderizados, com sincronização ao tema ativo (Dark/Light) e tratamento amigável de sintaxes inválidas.
- **Logs e Banner do Terminal Informativos**: Exibição clara no banner de inicialização do terminal da URL local (`http://localhost:8080`) e da URL de rede local (`http://<lan-ip>:8080`).

## Capabilities

### New Capabilities
<!-- Nenhuma nova capacidade raiz necessária; expande os requisitos existentes de fundação e interface -->

### Modified Capabilities
- `project-foundation`: Atualiza os requisitos da barra lateral para incluir redimensionamento dinâmico por arraste com persistência de largura, atualiza o requisito de compartilhamento por QR Code e escuta HTTP para utilizar endereços IP acessíveis na rede local (LAN), e aprimora os requisitos de renderização e tratamento de erro de diagramas Mermaid.

## Impact

- **Frontend**:
  - `web/templates/base.html`: Inclusão do elemento divisor arrastável (`#sidebar-resizer`), injeção do IP da rede local via atributos de dados e garantia do carregamento correto do Mermaid.
  - `web/static/css/style.css`: Estilos e classes para o divisor de redimensionamento (`.sidebar-resizer`), cursor `col-resize`, manipulação de seleção durante o arraste e estilos para diagramas Mermaid e alertas de erro.
  - `web/static/js/app.js`: Lógica de *drag-and-drop* para ajuste da largura da barra lateral com persistência no `localStorage`, adaptação da geração do QR Code para usar o endereço LAN e chamada explícita de `mermaid.run()` / `initMermaid`.
  - `web/static/js/mermaid.min.js`: Fornecimento de carregador e execução confiável da engine Mermaid com suporte a temas e tratamento de exceções.
- **Backend (Go)**:
  - `internal/adapters/in/http/server.go`: Configuração do `net.Listen` para escutar em todas as interfaces de rede (`0.0.0.0` ou `:porta`) e inclusão do método para obter o IP/URL de LAN.
  - `internal/adapters/out/renderer/goldmark.go`: Preservação e formatação dos blocos `<div class="mermaid">` para consumo sem perdas pelo parser do Mermaid.
  - `internal/core/domain/config.go` e `internal/core/domain/document.go` / `PageData`: Propagação das informações de rede local para a camada de apresentação.
  - `cmd/md_server/main.go`: Atualização das mensagens informativas exibidas no console na inicialização.
- **Testes & Cobertura**:
  - Testes unitários para detecção do IP de rede local, renderização de blocos Mermaid e testes de integração para o servidor HTTP e templates garantindo >= 80% de cobertura.
