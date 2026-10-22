package usecase

import (
	"bytes"
	"testing"
)

func TestReadInput(t *testing.T) {
	// 1. Simulamos uma sequência de entradas do usuário separadas por '\n':
	// - Primeiro uma string vazia (deve dar erro/mensagem)
	// - Segundo um caminho inválido (deve dar erro/mensagem)
	// - Terceiro um caminho válido (deve retornar com sucesso e encerrar a função)
	inputSimulado := "\n/caminho/invalido\n/caminho/valido\n"

	// Usamos bytes.Reader para simular o os.Stdin
	reader := bytes.NewBufferString(inputSimulado)

	// Usamos bytes.Buffer para capturar as mensagens impressas (fmt.Fprintln)
	var writer bytes.Buffer

	// Mock da função de verificação de diretório
	// Retorna true apenas para "/caminho/valido"
	mockIsDir := func(path string) bool {
		return path == "/caminho/valido"
	}

	// Executamos a função
	resultado := ReadInput(reader, &writer, mockIsDir)

	// Validações
	esperado := "/caminho/valido"
	if resultado != esperado {
		t.Errorf("Esperado '%s', mas obteve '%s'", esperado, resultado)
	}
}
