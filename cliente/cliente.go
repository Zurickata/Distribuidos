package main

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

type Planet struct {
	Name    string
	Booty   int
	Captain string
}

func main() {
	// Definir el número máximo de iteraciones y el intervalo de tiempo entre solicitudes
	maxIterations := 2

	// Definir el rango específico para los datos aleatorios
	minBooty := 1
	maxBooty := 20

	// Establecer una semilla aleatoria
	rand.Seed(time.Now().UnixNano())

	// Bucle para enviar solicitudes repetidamente
	for i := 0; i < maxIterations; i++ {

		planet := string(rune('A' + rand.Intn(6)))
		booty := rand.Intn(maxBooty-minBooty+1) + minBooty
		captain := "C" + strconv.Itoa(rand.Intn(3)+1)

		message := fmt.Sprintf("%s:%d:%s", planet, booty, captain)
		sendRequest(message)

		// Generar un intervalo aleatorio entre 2 y 7 segundos
		interval := time.Duration(rand.Intn(6)+2) * time.Second
		
		// Esperar el intervalo de tiempo antes de enviar la próxima solicitud
		time.Sleep(interval)
	}

	// En teoría deberían salir la solicitudes de manera aleatoria, no de forma secuencial cada "interval" de tiempo.
}

func sendRequest(message string) {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error al conectar con el servidor:", err)
		return
	}
	defer conn.Close()

	// Enviar mensaje al servidor
	_, err = conn.Write([]byte(message))

	parts := strings.Split(message, ":")
	captain := parts[2]
	planetName := parts[0]
	fmt.Printf("Capitán %s encontró botín en Planeta %s, enviando solicitud de asignación...\n", captain, planetName)

	if err != nil {
		fmt.Println("Error al enviar el mensaje:", err)
		return
	}

	// Leer respuesta del servidor
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error al leer la respuesta del servidor:", err)
		return
	}
	response := string(buffer[:n])

	// Imprimir respuesta del servidor
	fmt.Println("Respuesta del servidor:", response)
}
