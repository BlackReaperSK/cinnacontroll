// commons/utility/listener.go

package utility

import (
	"fmt"
	"net"
)

// Função para tratar a conexão
func HandleConnection(conn net.Conn) {
	defer conn.Close()
	// Loop para receber e descartar os dados continuamente
	buffer := make([]byte, 1024)
	for {
		_, err := conn.Read(buffer)
		if err != nil {
			break
		}
	}

	fmt.Printf("Conexão encerrada de %s\n", conn.RemoteAddr())
}
