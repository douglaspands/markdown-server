## Why

O projeto **md_server (Markdown Viewer)** é distribuído como uma ferramenta utilitária e servidor web de código aberto. Embora o documento [`README.md`](file:///home/douglas/Workspace/gemini/markdown-server/README.md) já referencie a licença MIT e exiba um badge correspondente, o arquivo formal `LICENSE` contendo o texto jurídico padrão da licença MIT e os termos legais de permissão, limitação de responsabilidade e aviso de copyright ainda não está presente na raiz do repositório.

A inclusão explícita do arquivo `LICENSE` é fundamental para formalizar a distribuição pública do software, assegurar segurança jurídica aos usuários e contribuidores e cumprir com os padrões da comunidade open source.

## What Changes

- **Adição do arquivo `LICENSE`**: Criação do arquivo de licença formal na raiz do repositório com o texto canônico da Licença MIT (OSI-approved), definindo o ano atual e o detentor dos direitos autorais.
- **Validação de Conformidade**: Assegurar a coerência das referências à licença no [`README.md`](file:///home/douglas/Workspace/gemini/markdown-server/README.md#L7-L161) e nos metadados de distribuição do projeto.

## Capabilities

### Modified Capabilities

- `project-foundation`: Adiciona o requisito explícito de licenciamento e distribuição open source sob a Licença MIT com o arquivo `LICENSE` presente na raiz do repositório.

## Impact

- **Arquivos Afetados**: Criação de `LICENSE` na raiz do projeto.
- **APIs e Código**: Nenhum impacto funcional em APIs Go, serviços internos ou adaptadores.
- **Compatibilidade e Qualidade**: Sem quebras de compatibilidade; atende integralmente aos quality gates (`make check`).
