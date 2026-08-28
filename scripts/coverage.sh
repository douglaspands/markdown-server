#!/usr/bin/env bash

set -euo pipefail

# Limiar inegociável de cobertura mínima
COVERAGE_THRESHOLD=80.0
COVERAGE_OUT="coverage.out"
COVERAGE_HTML="coverage.html"

echo "================================================================="
echo " Executando Suíte de Testes com Análise de Cobertura Global"
echo "================================================================="

# Executa todos os testes com cobertura atômica
go test -race -coverprofile="${COVERAGE_OUT}" -covermode=atomic ./...

# Gera relatório em HTML para inspeção visual
go tool cover -html="${COVERAGE_OUT}" -o "${COVERAGE_HTML}"

# Calcula a cobertura total
TOTAL_COVERAGE=$(go tool cover -func="${COVERAGE_OUT}" | grep 'total:' | awk '{print $3}' | sed 's/%//')

echo "================================================================="
echo " Cobertura Global Obtida : ${TOTAL_COVERAGE}%"
echo " Limiar Mínimo Requerido : ${COVERAGE_THRESHOLD}%"
echo " Relatório HTML gerado em: ${COVERAGE_HTML}"
echo "================================================================="

# Validação com comparação numérica
COVERAGE_PASSED=$(awk -v total="${TOTAL_COVERAGE}" -v threshold="${COVERAGE_THRESHOLD}" 'BEGIN { if (total >= threshold) print "1"; else print "0"; }')

if [ "${COVERAGE_PASSED}" -eq "1" ]; then
    echo " [SUCESSO] Barreira de cobertura superada com sucesso! (${TOTAL_COVERAGE}% >= ${COVERAGE_THRESHOLD}%)"
    exit 0
else
    echo " [FALHA] Cobertura insuficiente! Obtido: ${TOTAL_COVERAGE}%, Mínimo Requerido: ${COVERAGE_THRESHOLD}%"
    exit 1
fi
