# Regra 03: Infraestrutura de Testes Automatizados, TDD/BDD e Barreira >= 80%

A qualidade de software e a confiabilidade do código em Go são garantidas através de testes automatizados determinísticos, estruturação BDD e verificação de cobertura rigorosa.

## 1. Ferramental e Bibliotecas
- Pacote nativo `testing` do Go.
- `github.com/stretchr/testify/assert` e `github.com/stretchr/testify/require` para asserções expressivas e legíveis.
- Evitar frameworks de BDD complexos que obscureçam o código Go nativo; adotar subtestes declarativos nativos (`t.Run`).

## 2. Padrão Declarativo BDD com Subtestes

Todo cenário de teste deve seguir o formato formal:
```go
func TestMarkdownService_RenderFile(t *testing.T) {
    t.Run("Given an existing markdown file When RenderFile is called Then it returns parsed HTML and metadata", func(t *testing.T) {
        // GIVEN: Preparação de mocks e fixtures
        ctrl := testutils.NewMockRepository()
        service := services.NewMarkdownService(ctrl)
        ctx := context.Background()

        // WHEN: Execução do caso de uso
        result, err := service.RenderFile(ctx, "docs/index.md")

        // THEN: Asserções de sucesso e contrato
        require.NoError(t, err)
        assert.NotEmpty(t, result.HTML)
        assert.Equal(t, "Index Page", result.Title)
    })

    t.Run("Given a non-existent file path When RenderFile is called Then it returns ErrFileNotFound", func(t *testing.T) {
        // GIVEN
        ctrl := testutils.NewMockRepository()
        service := services.NewMarkdownService(ctrl)
        ctx := context.Background()

        // WHEN
        _, err := service.RenderFile(ctx, "invalid/path.md")

        // THEN
        require.Error(t, err)
        assert.ErrorIs(t, err, domain.ErrFileNotFound)
    })
}
```

## 3. Barreira Inegociável de Cobertura (>= 80%)

- A suíte de testes deve manter **cobertura global mínima de 80%** calculada pelo script `scripts/coverage.sh`.
- O script gera `coverage.out` e `coverage.html`. Se a cobertura for inferior a 80%, a execução falha imediatamente com código de saída 1.
- O Quality Gate (`make check` ou `make test-coverage`) bloqueia commits ou builds que quebrem este limiar.
