package main

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Planet struct {
	Name    string
	Booty   int
	Captain string
}

var planets = make(map[string]int)
var mutex sync.Mutex

func main() {
	initializePlanets()
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
		return
	}
	defer ln.Close()
	fmt.Println("Servidor en ejecución en el puerto 8080...")
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error al aceptar la conexión:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func initializePlanets() {
	// Inicializar planetas con botines aleatorios
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 6; i++ {
		planetName := string(rune('A' + i))
		planets[planetName] = rand.Intn(11)
	}

	printPlanetStatus()
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Leer mensaje del cliente
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error al leer el mensaje del cliente:", err)
		return
	}
	message := string(buffer[:n])

	// Descompone el mensaje del cliente en una Lista
	parts := strings.Split(message, ":")

	if len(parts) != 3 {
		fmt.Println("Mensaje del cliente con formato incorrecto:", message)
		return
	}
	planetName := parts[0]
	booty, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Println("Error al convertir botín a entero:", err)
		return
	}
	captain := parts[2]

	fmt.Printf("Recepción de solicitud desde el Planeta %s, del capitan %s\n", planetName, captain)

	// Asignar botín al planeta adecuado
	assignBooty(booty)

	// Mostrar estado actual de las asignaciones de los planetas
	printPlanetStatus()

	// Responder al cliente con el planeta asignado
	response := fmt.Sprintf("Botín asignado a planeta %s para capitán %s\n", planetName, captain)
	conn.Write([]byte(response))
}

func assignBooty(booty int) {
	// Lógica para asignar botín al planeta adecuado
	// (en este caso, simplemente se asigna al planeta con menos botines)
	minBooty := 9999
	var targetPlanet string
	for name, booty := range planets {
		if booty < minBooty {
			minBooty = booty
			targetPlanet = name

			fmt.Printf("Botín asignado al Planeta %s, cantidad actual: %d\n", targetPlanet, minBooty)
		}
	}
	mutex.Lock()
	defer mutex.Unlock()
	planets[targetPlanet] += booty
}

func printPlanetStatus() {
	// Imprimir estado actual de las asignaciones de los planetas
	status := "Estado actual de las asignaciones de los planetas: "
	for name, booty := range planets {
		status += fmt.Sprintf("P%s: %d, ", name, booty)
	}
	fmt.Println(status[:len(status)-2]) // Eliminar la última coma y espacio
}
