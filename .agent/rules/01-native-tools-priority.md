# Regra 01: Prioridade Mandatória de Ferramentas Nativas (Native Tool Grounding)

Esta regra é mandatória para qualquer agente de IA que opere neste repositório. O agente DEVE priorizar ferramentas de I/O estruturadas em vez de comandos de shell/terminal para manipular o sistema de arquivos e inspecionar código.

## 1. Mapeamento de Ferramentas Nativas

| Operação Desejada | Ferramenta Nativa Obrigatória | Prática ESTRITAMENTE PROIBIDA no Terminal |
| :--- | :--- | :--- |
| **Criar novos arquivos** | `write_to_file` | `cat << 'EOF' > ...`, `echo "..." > ...`, `touch ...` |
| **Editar arquivos existentes** | `replace_file_content` | `sed -i ...`, `awk`, `cat > ...`, `patch` |
| **Inspecionar / Ler arquivos** | `view_file` (com `StartLine`/`EndLine`) | `cat ...`, `head -n ...`, `tail -n ...`, `less` |
| **Buscar padrões / termos** | `grep_search` (com filtros `Includes`) | `grep -r ...`, `rg ...`, `find ... -exec grep` |
| **Localizar arquivos por nome** | `find_by_name` (com `Pattern`/`Extensions`) | `find . -name ...`, `ls -R` |
| **Listar diretórios** | `list_dir` | `ls`, `ls -la`, `dir`, `tree` |

## 2. Uso Restrito do Terminal (`run_command`)

O terminal deve ser acionado **exclusivamente** para tarefas do ciclo de vida de compilação, testes e governança:
- **Makefile**: `make help`, `make setup`, `make dev`, `make run`, `make test`, `make test-coverage`, `make lint`, `make fmt`, `make check`, `make build`, `make build-all`, `make clean`.
- **Ferramentas Go**: `go test`, `go mod tidy`, `go vet`, `go generate`.
- **Git**: `git status`, `git branch`, `git checkout`, `git add`, `git commit`, `git diff`.
- **OpenSpec CLI**: `openspec status`, `openspec instructions`, `openspec validate`, `openspec archive`.

## 3. Economia de Tokens e Curação de Contexto
- Nunca leia arquivos inteiros desnecessariamente. Use `view_file` delimitando o intervalo de linhas (`StartLine` e `EndLine`).
- Ao editar, faça substituições cirúrgicas através de `replace_file_content` garantindo correspondência exata de indentação e contexto.
- Não duplique grandes blocos de código já presentes no disco dentro das respostas de texto. Utilize links clicáveis no formato `[arquivo.ext](file:///caminho/arquivo.ext#L1-L20)`.
