package handlers

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	target := r.URL.Query().Get("target")

	// Command Injection - выполнение команды без санитизации ввода (CWE-78)
	cmdStr := "ping -c 1 " + target
	cmd := exec.Command("sh", "-c", cmdStr)
	
	output, err := cmd.CombinedOutput()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error: %v\nOutput: %s", err, output)))
		return
	}

	w.Write(output)
}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Query().Get("file")

	// Path Traversal - чтение файла по пути от пользователя без защиты (CWE-22)
	data, err := os.ReadFile(fileName)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	w.Write(data)
}
