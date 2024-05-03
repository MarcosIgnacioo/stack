package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	// Crear una nueva pila
	s := NewStack()

	// Verificar que la pila está vacía al inicio
	if !s.IsEmpty() {
		t.Errorf("Se esperaba que la pila esté vacía, pero no lo está")
	}

	// Agregar elementos a la pila
	s.Push(2)
	s.Push(3)

	// Verificar si la pila está vacía después de agregar elementos
	if s.IsEmpty() {
		t.Errorf("Se esperaba que la pila no esté vacía, pero está vacía")
	}

	// Verificar si la pila está llena
	if s.IsFull() {
		t.Errorf("Se esperaba que la pila no esté llena, pero está llena")
	}

	// Revisar el tamaño de la pila
	if s.Size != 2 {
		t.Errorf("El tamaño de la pila es incorrecto. Se esperaba 3, pero se obtuvo %d", s.Size)
	}

	// Revisar el funcionamiento de Pop
	popped := s.Pop()
	if popped != 3 {
		t.Errorf("El elemento desapilado es incorrecto. Se esperaba 3, pero se obtuvo %d", popped)
	}

	// Verificar si la pila está vacía después de desapilar
	if s.IsEmpty() {
		t.Errorf("Se esperaba que la pila no esté vacía después de desapilar, pero está vacía")
	}

	// Vaciar la pila
	s.Empty()

	// Verificar si la pila está vacía después de vaciarla
	if !s.IsEmpty() {
		t.Errorf("Se esperaba que la pila esté vacía después de vaciarla, pero no lo está")
	}

	s = NewStack()

	// Agregar elementos a la pila
	s.Push(1)
	s.Push(2)
	s.Push(3)

	// Invertir la pila
	s.Reverse()

	// Comprobar si los elementos se invirtieron correctamente
	expected := []int{3, 2, 1}
	for _, exp := range expected {
		if s.Pop() != exp {
			t.Errorf("Error al invertir la pila. Se esperaba %d, pero se obtuvo otro elemento", exp)
		}
	}
}
