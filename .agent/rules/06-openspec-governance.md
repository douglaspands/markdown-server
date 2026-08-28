# Regra 06: Governança de Especificações com OpenSpec (Alinhamento PO & QA)

O framework OpenSpec é a fonte única de verdade para governança de requisitos, rastreabilidade arquitetural e alinhamento contínuo entre Product Owner (PO) e Quality Assurance (QA).

## 1. Fluxo de Vida de uma Mudança OpenSpec
1. **Proposta (`openspec new change <nome>` / `proposal.md`)**: Define o problema, a motivação de negócio e as capacidades introduzidas ou modificadas.
2. **Especificação Delta (`specs/<capability-path>/spec.md`)**: Define o contrato comportamental do sistema sob ótica PO/QA.
3. **Design Técnico (`design.md`)**: Define a arquitetura interna, escolhas técnicas, portas, serviços e adaptadores.
4. **Plano de Tarefas (`tasks.md`)**: Detalha as tarefas sequenciais de implementação, testes com barreira >= 80% e verificação final.
5. **Aplicação (`/opsx-apply`)**: Implementação passo a passo com testes e validações.
6. **Arquivamento (`/opsx-archive`)**: Sincronização das specs delta para as specs principais e arquivamento formal da mudança.

## 2. Diretrizes Colaborativas para Especificações (`spec.md`)

### Para o Product Owner (PO)
- Redação em linguagem clara, ubíqua e acessível ao negócio, sem detalhes de implementação interna (como nomes de funções ou structs Go).
- Seção `## Purpose` obrigatória com no mínimo 50 caracteres explicando o valor entregue.
- Requisitos funcionais declarados como `### Requirement: <Nome do Requisito>` com critérios de aceitação objetivos.

### Para o Quality Assurance (QA) e Automação de Testes
- Estruturação formal de cenários BDD/Gherkin com 4 hashtags `#### Scenario: <Nome do Cenário>`.
- Bullets padronizados:
  - `- **GIVEN** <estado inicial do sistema ou contexto>`
  - `- **WHEN** <ação do usuário ou evento>`
  - `- **THEN** <comportamento esperado e verificável>`
- Cobertura obrigatória de:
  - Caminho Feliz (Happy Path).
  - Fluxos Alternativos e de Borda (Edge Cases).
  - Tratamento de Erros e Exceções (Falhas de I/O, arquivos corrompidos, caminhos inválidos).
- Os cenários devem ser inequívocos para permitir mapeamento 1:1 com subtestes em Go (`t.Run`) ou testes E2E.
