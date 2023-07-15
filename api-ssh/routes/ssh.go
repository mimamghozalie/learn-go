package routes

import (
	"fmt"
	"net"
	"net/http"
)

func LoginSSH(c *gin.Context) {
	// Mendapatkan nilai parameter ip, ssh_key, dan auth_method
	ip := c.PostForm("ip")
	sshKey := c.PostForm("ssh_key")
	authMethod := c.PostForm("auth_method")

	// Membuat konfigurasi SSH
	config := &ssh.ClientConfig{
		User:            "username", // Ganti dengan username SSH Anda
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Memilih metode otentikasi yang akan digunakan
	switch authMethod {
	case "password":
		password := c.PostForm("password")
		config.Auth = []ssh.AuthMethod{
			ssh.Password(password),
		}
	case "key":
		// Memuat kunci SSH dari sshKey
		key, err := ssh.ParsePrivateKey([]byte(sshKey))
		if err != nil {
			// Mengembalikan response dengan status "error"
			response := map[string]interface{}{
				"error":  fmt.Sprintf("Failed to parse SSH key: %v", err),
				"status": "error",
			}
			c.JSON(http.StatusInternalServerError, response)
			return
		}
		config.Auth = []ssh.AuthMethod{
			ssh.PublicKeys(key),
		}
	default:
		// Mengembalikan response dengan status "error" jika metode otentikasi tidak valid
		response := map[string]interface{}{
			"error":  "Invalid authentication method",
			"status": "error",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Membuat koneksi SSH
	client, err := ssh.Dial("tcp", ip+":22", config)
	if err != nil {
		// Mengembalikan response dengan status "error"
		response := map[string]interface{}{
			"error":  fmt.Sprintf("Failed to connect to SSH server: %v", err),
			"status": "error",
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// Melakukan proses login ssh menggunakan nilai parameter yang diberikan
	// ...

	// Mendapatkan IP server
	serverIP, _, _ := net.SplitHostPort(c.Request.RemoteAddr)

	// Mengembalikan response
	response := map[string]interface{}{
		"ip":     serverIP,
		"status": "success",
	}
	c.JSON(http.StatusOK, response)
}
