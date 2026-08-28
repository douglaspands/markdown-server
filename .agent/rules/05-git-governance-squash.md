# Regra 05: Governança Git, Higiene de Repositório e Squash Merge

A governança do controle de versão preserva o histórico do projeto linear, compreensível, seguro e auditável.

## 1. Higiene Rigorosa e Economia de Contexto
- Manter o `.gitignore` permanentemente atualizado para impedir que binários (`bin/`, `dist/`, `*.exe`), relatórios (`coverage.out`, `coverage.html`) e caches poluam o repositório.
- A presença de arquivos de lixo ou binários no repositório degrada as ferramentas de busca do agente (`grep_search`, `find_by_name`, `list_dir`) e polui a janela de contexto.
- Prevenção estrita contra commits acidentais de credenciais, certificados e segredos (`.env`, `*.pem`, `*.key`).

## 2. Padrão Conventional Commits
Todas as mensagens de commit devem seguir o padrão Conventional Commits:
- `feat:` Nova funcionalidade para o usuário ou capacidade do sistema.
- `fix:` Correção de bug.
- `refactor:` Alteração de código que não adiciona recurso nem corrige bug.
- `test:` Adição ou ajuste de testes automatizados.
- `docs:` Alterações em documentações, especificações ou README.
- `chore:` Tarefas de manutenção, dependências ou tooling.
- `perf:` Melhoria de performance.
- `ci:` Alterações em pipelines de CI/CD ou GitHub Actions.

Exemplo: `feat(server): add auto browser open support for windows double-click`

## 3. Isolamento em Feature Branch e Permissão Obrigatória
- Toda nova especificação ou mudança deve ser desenvolvida em uma branch dedicada: `git checkout -b feature/<nome-da-mudanca>`.
- O agente de IA **NUNCA deve realizar merge na branch `main` ou abrir Pull Requests sem a autorização prévia e explícita do usuário**.

## 4. Estratégia de Integração Exclusivamente SQUASH
- Ao autorizar a integração na branch `main`, a estratégia obrigatória é o **Squash Merge**:
  ```bash
  git checkout main
  git merge --squash feature/<nome-da-mudanca>
  git commit -m "feat(<nome-da-mudanca>): <resumo consolidado das mudanças>"
  ```
- **Racional Técnico**: O Squash Merge consolida múltiplos commits intermediários em um único commit atômico e testado na `main`. Isso garante histórico limpo, linear, legível e 100% compatível com operações de `git bisect` e `git revert`.
