# Regra 02: Padrão Arquitetural e Estrutura Canônica em Go (Clean Architecture)

Todo código Go desenvolvido neste projeto deve seguir rigorosamente os princípios de Clean Architecture (Ports & Adapters / Hexagonal Architecture), garantindo isolamento total do domínio, desacoplamento de infraestrutura e testabilidade determinística.

## 1. Layout Canônico de Diretórios

```
md_server/
├── cmd/
│   └── md_server/           # Ponto de entrada (CLI & Composition Root)
│       └── main.go          # Parse de flags, instanciação de adapters/services e arranque
├── internal/
│   ├── core/
│   │   ├── domain/          # Entidades puras e regras de negócio sem dependências externas
│   │   ├── ports/           # Interfaces formais de entrada (Driver) e saída (Driven)
│   │   └── services/        # Casos de uso implementando as portas de entrada
│   ├── adapters/
│   │   ├── in/              # Adaptadores de entrada (CLI handlers, HTTP web handlers, middleware)
│   │   └── out/             # Adaptadores de saída (FS scanner, Goldmark renderer, Browser launcher)
│   ├── version/             # Metadados de versão (Version, Commit, Date injetados via ldflags)
│   └── testutils/           # Mocks, fixtures e builders auxiliares para testes
├── pkg/                     # (Opcional) Pacotes explicitamente exportáveis para outros módulos
├── scripts/                 # Scripts de automação e validação (coverage.sh)
├── web/                     # Assets embutidos (templates HTML, CSS, JS Mermaid)
└── .github/workflows/       # Pipelines de CI e Release multiplataforma
```

## 2. Regras Fundamentais de Design

1. **Regra de Isolamento Estrito (`internal/`)**:
   - Toda lógica de negócio, entidades, portas, serviços e adaptadores privados residem sob `internal/`.
   - Nenhum módulo externo pode importar código privado.

2. **Injeção de Dependências Explícita (Constructor Injection)**:
   - Todas as dependências devem ser passadas explicitamente através de funções construtoras tipadas (ex: `NewMarkdownService(repo ports.FileRepository, renderer ports.MarkdownRenderer)`).
   - É terminantemente proibido o uso de reflexão em tempo de execução, contêineres de DI mágicos ou variáveis globais de estado compartilhado.

3. **Propagação Idiomática de Contexto (`context.Context`)**:
   - Toda função ou método que realize operações de I/O (leitura de disco, requisições HTTP), concorrência ou ciclo de vida deve receber `ctx context.Context` como seu **primeiro parâmetro**.
   - O encerramento da aplicação deve utilizar `context.WithCancel` ou `signal.NotifyContext` para garantir graceful shutdown seguro em servidores e workers.

4. **Direção das Dependências**:
   - `domain` não depende de ninguém.
   - `ports` depende apenas de `domain` e stdlib.
   - `services` implementa as portas de entrada e depende das portas de saída e de `domain`.
   - `adapters` implementam portas de saída ou chamam portas de entrada via `services`.
   - `cmd/md_server` é o Composition Root que conecta todas as partes.
