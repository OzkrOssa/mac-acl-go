package devices

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/ssh"
)

func Ubiquiti(ipAddress string, macAddress string, comment string) (string, error) {
	godotenv.Load()

	var user, password string = os.Getenv("USER"), os.Getenv("PASSWORD")
	fmt.Println(user, password)

	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password + "68531059$"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	sshClient, err := ssh.Dial("tcp", fmt.Sprintf("%s:22", ipAddress), sshConfig)
	if err != nil {
		panic(err)
	}
	defer sshClient.Close()

	session, err := sshClient.NewSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	output, err := session.CombinedOutput("cat /tmp/system.cfg | grep wireless.1.mac_acl.")
	if err != nil {
		panic(err)
	}

	lastNumber := 0
	for _, line := range strings.Split(string(output), "\n") {
		if strings.HasPrefix(line, "wireless.1.mac_acl.") {
			parts := strings.Split(line, ".")
			fmt.Println(len(parts))
			if len(parts) >= 4 {
				if number, err := strconv.Atoi(parts[3]); err == nil {
					lastNumber = int(math.Max(float64(lastNumber), float64(number)))
				}
			}
		}
	}

	return string("output"), nil
}
