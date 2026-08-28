## Why

O repositório remoto no GitHub (`douglaspands/markdown-server`) está atualmente com a branch `feature/extended-markdown-plugins` configurada como branch padrão (*default branch*), em vez da branch canônica `main`. Essa inconsistência faz com que novos Pull Requests, clones do repositório e a página inicial do GitHub apontem para uma branch de feature antiga e desatualizada.

É necessário corrigir essa configuração no GitHub, definindo a branch `main` como a branch padrão definitiva do projeto e assegurando que esteja totalmente atualizada e sincronizada com o último commit (`93be511`).

## What Changes

- **Definição da Branch Padrão no GitHub**: Alterar a *default branch* do repositório `douglaspands/markdown-server` para `main` utilizando a API / CLI do GitHub (`gh repo edit`).
- **Sincronização e Validação**: Verificar e garantir que a branch `main` remota e local estejam perfeitamente alinhadas com o commit mais recente.
- **Remoção de Branches Remotas Obsoletas**: Limpar branches de feature já mescladas no GitHub (`origin/feature/extended-markdown-plugins`, `origin/feature/add-mit-license`) para manter a árvore do repositório limpa e linear.

## Capabilities

### New Capabilities
<!-- Nenhuma nova capacidade funcional adicionada -->

### Modified Capabilities
<!-- Nenhuma especificação de comportamento alterada (governança e infraestrutura de repositório com skip_specs: true) -->

## Impact

- **Repositório GitHub**: Acesso padrão da interface web e comandos `git clone` apontarão diretamente para `main`.
- **Pull Requests**: Novos PRs terão a `main` como branch base padrão automática.
- **Código da Aplicação**: Nenhum impacto ou alteração no código Go, portas, serviços ou assets da aplicação.
