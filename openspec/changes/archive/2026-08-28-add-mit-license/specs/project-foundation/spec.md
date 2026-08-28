## ADDED Requirements

### Requirement: Licenciamento Open Source sob a Licença MIT
O repositório do projeto `md_server` SHALL disponibilizar formalmente o arquivo `LICENSE` em sua raiz, contendo a íntegra dos termos canônicos da Licença MIT (aprovada pela Open Source Initiative - OSI), com declaração expressa de direitos autorais, concessão irrestrita de uso/cópia/modificação/redistribuição e termo padrão de isenção de garantia e limitação de responsabilidade.

#### Scenario: Disponibilidade do arquivo LICENSE canônico na raiz do projeto
- **GIVEN** que o repositório do `md_server` é inspecionado ou distribuído
- **WHEN** qualquer usuário, contribuidor ou ferramenta automatizada de conformidade de código aberto consulta a raiz do projeto
- **THEN** o sistema SHALL conter o arquivo `LICENSE` em formato de texto simples legível, contendo os termos integrais e padronizados da Licença MIT e a linha de copyright (`Copyright (c) 2026 Markdown Viewer Contributors`).

#### Scenario: Coerência entre documentação e termos de licenciamento
- **GIVEN** que o [`README.md`](file:///home/douglas/Workspace/gemini/markdown-server/README.md) declara a distribuição do software sob a Licença MIT e exibe badge indicativo
- **WHEN** o leitor da documentação navega até a seção de licença ou clica no link de referência
- **THEN** a documentação SHALL apontar diretamente para o arquivo `LICENSE` existente na raiz, mantendo conformidade e integridade jurídica no projeto.
