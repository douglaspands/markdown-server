## 1. Descoberta de Rede e Ajustes no Servidor HTTP (Backend Go)

- [x] 1.1 Implementar utilitário de detecção de IP IPv4 não-loopback da rede local (`DetectLANIP`) com testes unitários cobrindo cenários com e sem interface externa
- [x] 1.2 Atualizar o adaptador de servidor HTTP (`internal/adapters/in/http/server.go`) para escuta em todas as interfaces (`0.0.0.0:<port>`), expondo métodos para obter URL local e URL de LAN
- [x] 1.3 Atualizar `PageData` e `handleDocument` para propagar o endereço/URL de rede local para os templates HTML
- [x] 1.4 Atualizar o banner de inicialização no terminal (`cmd/md_server/main.go`) para exibir as URLs de acesso Local e de Rede Local (LAN)

## 2. Divisor Arrastável e Redimensionamento da Barra Lateral (Frontend UI)

- [x] 2.1 Adicionar o elemento divisor (`#sidebar-resizer`) no template `web/templates/base.html` entre a barra lateral e a área principal de leitura
- [x] 2.2 Implementar estilos CSS para o divisor de redimensionamento (`.sidebar-resizer`), cursor `col-resize`, classe `.resizing` e limites visuais em `web/static/css/style.css`
- [x] 2.3 Implementar a lógica JavaScript (`initSidebarResizer`) em `web/static/js/app.js` com eventos de mouse/touch, cálculo de largura com limites (180px - 600px / 50vw), persistência em `localStorage` e integração suave com o recolhimento (`.collapsed`)

## 3. Geração de QR Code e Compartilhamento com IP da Rede Local (Frontend UI)

- [x] 3.1 Injetar os metadados de URL de rede local (`data-lan-url`) no template `web/templates/base.html`
- [x] 3.2 Atualizar a função `initQRCodeModal` em `web/static/js/app.js` para compor e gerar o QR Code e o campo copiável utilizando o IP de rede local, com fallback seguro para a URL atual do navegador

## 4. Correção e Renderização Confiável de Diagramas Mermaid (Frontend UI)

- [x] 4.1 Corrigir a inicialização e integração de Mermaid.js em `web/static/js/mermaid.min.js` e `web/static/js/app.js`, garantindo execução de `mermaid.run()` / `mermaid.init()`
- [x] 4.2 Garantir suporte à alternância de tema dinâmico (Dark/Light) na renderização e re-renderização de diagramas Mermaid
- [x] 4.3 Implementar tratamento visual de erro para blocos Mermaid com sintaxe inválida

## 5. Testes Automatizados, Cobertura e Validação Final

- [x] 5.1 Implementar e atualizar testes unitários e de integração em `internal/adapters/in/http/server_test.go` e `renderer_test.go` cobrindo o novo `PageData`, a escuta em rede, métodos de URL e blocos Mermaid
- [x] 5.2 Executar a suíte de testes completa e verificar a barreira inegociável de cobertura >= 80% via `make test-coverage`
- [x] 5.3 Validar a integridade geral do projeto e qualidade de código com `make check`
