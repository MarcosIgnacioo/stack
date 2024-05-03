package stack

import (
	"errors"
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}

func (n *Node) String() string {
	prev := n.Prev
	next := n.Next
	var prevValue int
	var nextValue int
	if prev != nil {
		prevValue = prev.Value
	}
	if next != nil {
		nextValue = next.Value
	}
	return fmt.Sprintf("%v <- %v -> %v", prevValue, n.Value, nextValue)

}

func NewNode(i int) *Node {
	return &Node{Value: i, Next: nil, Prev: nil}
}

type Stack struct {
	Head  *Node
	Size  int
	Limit int
}

func NewStack() *Stack {
	return &Stack{Head: nil, Size: 0, Limit: 100}
}

func (s *Stack) Print() {
	curr := s.Head
	for i := 0; i < s.Size; i++ {
		fmt.Println("|", curr, "|")
		curr = curr.Prev
	}
}

// En el caso de que no haya ningun nodo en el stack el nuevo nodo lo asigna como si fuera la cabeza, y hace que el nodo previo sea nulo junto al siguiente
// En el caso contrario, si el tama;o del stack no ha excedido el limite aun, signfica que podemos meter un nuevo nodo y que ya hay por lo menos un elemento dentro del stack. Asignamos a la cabezaa el nodo siguiente como el nuevo nodo, y hacemos que el nuevo nodo, su nodo previo haga referencia al head en ese momento, y ya al final hacemos que el head se actualice a ser el nuevo nodo ingresado
func (s *Stack) Push(i int) {
	if s.Size == 0 {
		s.Size++
		s.Head = NewNode(i)
		s.Head.Next = nil
		s.Head.Prev = nil
		return
	}
	if s.Size < s.Limit {
		s.Size++
		s.Head.Next = NewNode(i)
		s.Head.Next.Prev = s.Head
		s.Head = s.Head.Next
		return
	}
}

// Obtenemos el valor de la cabeza y si el tamano es mayor a 0 signfica que se
func (s *Stack) Pop() interface{} {
	if s.Size == 0 {
		return errors.New("stackunderflow")
	}
	popped := s.Head.Value
	s.Size--
	s.Head = s.Head.Prev
	s.Head.Next = nil
	return popped
}

func (s *Stack) IsFull() bool {
	return s.Size == s.Limit
}

func (s *Stack) IsEmpty() bool {
	return s.Size == 0
}

func (s *Stack) Empty() {
	s.Head = nil
	s.Size = 0
}

func (s *Stack) Reverse() {
	if s.Size <= 1 {
		return
	}
	// Guardamos en una variable axuliar al nodo previo
	// Intercambiamos el lugar del next y del prev
	// Si el valor del next del nodo actual no es nil signfica que aun hay mas elementos (Es next el que nos indica el elemento que corresponde porque ya hemos hecho el intercambio de las variables)
	// Si no es nil el valor del tope de la pila pasa a ser el siguiente elemento
	for {
		prev := s.Head.Prev
		s.Head.Prev = s.Head.Next
		s.Head.Next = prev
		if s.Head.Next == nil {
			break
		}
		s.Head = s.Head.Next
	}
}

func (s *Stack) FillStack() {
	for i := 0; i < s.Limit; i++ {
		s.Push(i + 1)
	}
}
