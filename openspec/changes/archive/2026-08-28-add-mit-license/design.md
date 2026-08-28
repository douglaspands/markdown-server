## Context

Consulte [`proposal.md`](file:///home/douglas/Workspace/gemini/markdown-server/openspec/changes/add-mit-license/proposal.md) para a motivação do negócio e contexto de licenciamento. O projeto **md_server (Markdown Viewer)** adota Clean Architecture em Go 1.22+ e disponibiliza binários estáticos e código-fonte aberto. A documentação ([`README.md`](file:///home/douglas/Workspace/gemini/markdown-server/README.md#L7-L161)) já declara a distribuição sob a Licença MIT e referencia o arquivo `LICENSE`, o qual deve ser formalmente disponibilizado na raiz do repositório.

## Goals / Non-Goals

**Goals:**
- Criar o arquivo `LICENSE` na raiz do repositório contendo o texto padrão canônico da Licença MIT (aprovada pela Open Source Initiative - OSI).
- Definir a linha de copyright padronizada: `Copyright (c) 2026 Markdown Viewer Contributors`.
- Garantir a integridade das referências de licença no `README.md` e nos metadados de distribuição do projeto.

**Non-Goals:**
- Alterar o código-fonte Go, contratos de interfaces (ports) ou serviços em `internal/`.
- Adicionar ou alterar dependências no `go.mod`.
- Adotar esquemas de licenciamento duplo ou restritivos (GPL, AGPL).

## Decisions

### 1. Adoção da Licença MIT Oficial (OSI-Approved)
- **Decisão**: Utilizar o texto canônico e inalterado da Licença MIT padrão.
- **Racional**: A Licença MIT é uma das mais consagradas e liberais no ecossistema de software livre e de código aberto. Permite uso comercial, modificação, distribuição e sublicenciamento sem restrições complexas, exigindo apenas a manutenção do aviso de copyright e isenção de garantia.
- **Alternativas Consideradas**:
  - *Apache 2.0*: Mais extensa e com cláusulas explícitas de patente, desnecessárias para o escopo utilitário do `md_server`.
  - *BSD 3-Clause*: Similar em termos práticos, porém menos ubíqua do que a MIT para ferramentas de desenvolvimento rápido em Go.

### 2. Atribuição de Direitos Autorais Coletiva
- **Decisão**: Especificar `Copyright (c) 2026 Markdown Viewer Contributors` no cabeçalho do arquivo `LICENSE`.
- **Racional**: Permite acolher contribuições da comunidade mantendo atribuição coletiva e clara para o projeto.

```mermaid
flowchart TD
    A[Repositório md_server] --> B[LICENSE (MIT Text & Copyright)]
    A --> C[README.md (Badge & Link)]
    A --> D[Distribuição / Releases]
    B --- E[Permissões: Uso Comercial, Modificação, Distribuição]
    B --- F[Condições: Inclusão do Copyright e Licença]
    B --- G[Limitações: Isenção de Garantia e Responsabilidade]
```

## Risks / Trade-offs

- **[Inconsistência de quebra de linha ou codificação]** $\rightarrow$ *Mitigação*: Criar o arquivo `LICENSE` estritamente em UTF-8 com quebras de linha padrão Unix (LF) e sem caracteres de controle especiais.
- **[Ruptura de links relativos no README]** $\rightarrow$ *Mitigação*: Validar que o link `[LICENSE](LICENSE)` no `README.md` aponta diretamente para o arquivo recém-criado.
