package smtp

import (
	"fmt"
	"net"
	"net/smtp"
	"time"

	smtpmock "github.com/mocktools/go-smtp-mock/v2"
)

func RunAndSendEmail(from string, to string, data string) {
	server := smtpmock.New(smtpmock.ConfigurationAttr{
		LogToStdout:       true,
		LogServerActivity: true,
	})

	if err := server.Start(); err != nil {
		fmt.Println(err)
	}

	hostAddress, portNumber := "127.0.0.1", server.PortNumber()

	address := fmt.Sprintf("%s:%d", hostAddress, portNumber)
	timeout := time.Duration(2) * time.Second

	connection, _ := net.DialTimeout("tcp", address, timeout)
	client, _ := smtp.NewClient(connection, hostAddress)

	client.Mail(from)
	client.Rcpt(to)

	wr, _ := client.Data()
	wr.Write([]byte(data))
	wr.Close()

	client.Quit()
	client.Close()

	fmt.Println(server.Messages()[0])

	if err := server.Stop(); err != nil {
		fmt.Println(err)
	}
}
