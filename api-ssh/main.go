import (
	// ...
	"fmt"
	"golang.org/x/crypto/ssh"
)

func LoginSSH(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan nilai parameter ip, ssh_key, dan auth_method
	ip := r.FormValue("ip")
	sshKey := r.FormValue("ssh_key")
	authMethod := r.FormValue("auth_method")

	// Membuat konfigurasi SSH
	config := &ssh.ClientConfig{
		User: "username", // Ganti dengan username SSH Anda
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Memilih metode otentikasi yang akan digunakan
	switch authMethod {
	case "password":
		password := r.FormValue("password")
		config.Auth = []ssh.AuthMethod{
			ssh.Password(password),
		}
	case "key":
		// Memuat kunci SSH dari sshKey
		key, err := ssh.ParsePrivateKey([]byte(sshKey))
		if err != nil {
			// Mengembalikan response dengan status "error"
			response := map[string]interface{}{
				"error":   fmt.Sprintf("Failed to parse SSH key: %v", err),
				"status":  "error",
			}
			render.JSON(w, r, response)
			return
		}
		config.Auth = []ssh.AuthMethod{
			ssh.PublicKeys(key),
		}
	default:
		// Mengembalikan response dengan status "error" jika metode otentikasi tidak valid
		response := map[string]interface{}{
			"error":   "Invalid authentication method",
			"status":  "error",
		}
		render.JSON(w, r, response)
		return
	}

	// Membuat koneksi SSH
	client, err := ssh.Dial("tcp", ip+":22", config)
	if err != nil {
		// Mengembalikan response dengan status "error"
		response := map[string]interface{}{
			"error":   fmt.Sprintf("Failed to connect to SSH server: %v", err),
			"status":  "error",
		}
		render.JSON(w, r, response)
		return
	}

	// Melakukan proses login ssh menggunakan nilai parameter yang diberikan
	// ...

	// Mendapatkan IP server
	serverIP, _, _ := net.SplitHostPort(r.RemoteAddr)

	// Mengembalikan response
	response := map[string]interface{}{
		"ip":      serverIP,
		"status":  "success",
	}
	render.JSON(w, r, response)
}