# Regra 07: Disciplinas de Engenharia de Agentes em Sistemas Compostos (Harness, Loop & Graph Engineering)

O Antigravity CLI (AGY) opera como um sistema composto autônomo baseado em três disciplinas fundamentais da engenharia de agentes de inteligência artificial:

## 1. Harness Engineering (Arcabouço Operacional & Scaffolding de IA)
- **Definição**: Conjunto de diretrizes determinísticas (`.agent/rules/`, `AGENTS.md`, `GEMINI.md`), matriz de permissões pré-autorizadas (`.agent/settings.json`), guardrails de segurança, contexto otimizado e limites de sandbox que governam a operação autônoma, previsível e segura do agente.
- **Diretriz**:
  - Eliminar prompts repetitivos de confirmação através de permissões pré-autorizadas (`autoApprove` em `.agent/settings.json`).
  - O agente deve respeitar rigorosamente as fronteiras de sandbox do workspace.

## 2. Loop Engineering (Ciclos Cognitivos & Auto-Validação ReAct/Reflection)
- **Definição**: Ciclo cognitivo e iterativo de execução, reflexão e auto-validação contínua do agente (ReAct / Reflection loops):
  1. *Inspeção Cirúrgica*: Leitura precisa de arquivos e requisitos (`view_file`, `grep_search`).
  2. *Intervenção Pontual*: Edição atômica de código (`replace_file_content` / `write_to_file`).
  3. *Execução de Testes Automatizados*: Rodar a suíte de testes (`make test`).
  4. *Diagnóstico e Reflexão*: Análise de eventuais falhas ou quebra de cobertura.
  5. *Correção e Validação Final*: Ajuste imediato do código até atingir aprovação plena no Quality Gate (`make check`).
- **Diretriz**: Prevenir loops infinitos através de diagnósticos estruturados e convergência rápida com critérios de parada claros.

## 3. Graph Engineering (State Graphs & DAG de Raciocínio)
- **Definição**: Modelagem e orquestração do fluxo cognitivo/operacional do agente como um grafo direcionado acíclico (DAG) de tarefas, estados e decisões:
  - *Decomposição Topológica*: Resolução estrita de dependências entre tarefas (Ex: Proposta -> Especificação/Design -> Implementação -> Testes -> Validação -> Documentação).
  - *Transições Determinísticas de Estado*: Controle explícito de transições entre fases (`ready`, `in_progress`, `blocked`, `done`).
  - *Delegação para Subagentes*: Roteamento e isolamento de pesquisas exploratórias e tarefas concorrentes em subagentes dedicados (`research`, `self`), mantendo a janela de contexto do agente orquestrador limpa, concisa e focada.
