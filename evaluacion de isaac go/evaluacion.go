package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func Clear() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func readerStrings(message string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s ", message)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Hubo un error, intente nuevamente", err)
		return ""
	}
	input = strings.TrimSpace(input)
	return input
}

func readerInt(message string) int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s ", message)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Print("Hubo un error, intente nuevamente\n")
			continue
		}
		input = strings.TrimSpace(input)
		inputConv, errConv := strconv.Atoi(input)
		if errConv != nil {
			fmt.Print("Por favor ingrese un valor valido\n")
			continue
		}
		return inputConv
	}
}

type node_song struct {
	data *song
	next *node_song
}

type song struct {
	titulo  string
	artista string

	Next *song
}
type linkelist struct {
	cabeza *node_song
	tamaño uint32
}

func (lista *linkelist) Agregar(valor *song) {
	NuevoNodo := &node_song{data: valor}
	if lista.cabeza == nil {
		lista.cabeza = NuevoNodo
	} else {
		nuevo := lista.cabeza
		for nuevo.next != nil {
			nuevo = nuevo.next

		}
		nuevo.next = NuevoNodo

	}
	lista.tamaño++
}
func (lista *linkelist) Eliminar(titulo string) {
	re := readerInt("Estas seguro? \n1.Si\n2.No ")
	if re == 1 {
		if lista.cabeza == nil {
			fmt.Print("Lista vacia\n")

			readerStrings("Dale enter para continuar")
			return
		}
		if lista.cabeza.data.titulo == titulo {
			lista.cabeza = lista.cabeza.next
			return
		}
		prev := lista.cabeza
		current := lista.cabeza.next
		for current != nil && current.data.titulo != titulo {
			prev = current
			current = current.next
		}
		if current == nil {
			fmt.Print("No se encontro\n")
			readerStrings("Dale enter para continuar")
		}
		prev.next = current.next
		lista.tamaño--

	} else {
		return
	}

}
func (lista *linkelist) buscar(titulo string) {
	actual := lista.cabeza
	for actual != nil {
		if actual.data.titulo == titulo {
			fmt.Printf("Titulo: %s, Artista: %s\n", actual.data.titulo, actual.data.artista)
			readerStrings("Dale enter para continuar")
			return
		}
		actual = actual.next
	}
	fmt.Print("No se encontró la canción\n")
	readerStrings("Dale enter para continuar")
}

func main() {
	lista := linkelist{}
	for {
		fmt.Print("-----MUSICA----\n1.Agregar \n2.Eliminar \n3.Buscar \n4.Mostrar \n5.Salir\n")
		r := readerInt("Ingrese una opcion")

		if r == 5 {
			break
		}
		if r == 1 {
			Clear()
			titulo := readerStrings("Ingrese titulo: ")
			artista := readerStrings("Ingrese artista: ")

			cancion := &song{titulo: titulo, artista: artista}
			lista.Agregar(cancion)
			Clear()

		}
		if r == 2 {
			Clear()
			lista.Imprimir()
			titulo := readerStrings("Que titulo desea eliminar?")
			lista.Eliminar(titulo)
			Clear()
		}
		if r == 3 {
			Clear()
			titulo := readerStrings("Ingrese el titulo que desea buscar")
			lista.buscar(titulo)
			Clear()

		}
		if r == 4 {
			Clear()
			lista.Imprimir()
			readerStrings("\nDale enter para continuar")
			Clear()
		}
	}

}
func (lista *linkelist) Imprimir() {
	Nuevo := lista.cabeza
	for Nuevo != nil {
		fmt.Printf("Titulo:%s ,Artista:%s\n", Nuevo.data.titulo, Nuevo.data.artista)
		Nuevo = Nuevo.next
	}
	fmt.Print("")

}
