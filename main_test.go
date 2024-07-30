// main_test.go
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test case 1: Memastikan server berjalan dan mengembalikan status HTTP 200 OK
func TestServerStatus(t *testing.T) {
	// Buat server uji
	server := httptest.NewServer(http.FileServer(http.Dir("./static")))
	defer server.Close()

	// Kirim permintaan GET ke server
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Periksa status code dari respons
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
}

// Test case 2: Memastikan file yang tersedia disajikan dengan benar
func TestFileServing(t *testing.T) {
	// Buat server uji
	server := httptest.NewServer(http.FileServer(http.Dir("./static")))
	defer server.Close()

	// Kirim permintaan GET untuk file tertentu
	filePath := "/index.html" // Ganti dengan file yang sesuai jika perlu
	resp, err := http.Get(server.URL + filePath)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Periksa status code dari respons
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	// Periksa apakah file yang diinginkan ada di dalam respons
	// Untuk memastikan konten, kamu bisa memeriksa body respons
	// Misalnya, jika kamu ingin memastikan konten tertentu ada di dalam file
	// content, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	t.Fatalf("Failed to read response body: %v", err)
	// }
	// if !strings.Contains(string(content), "konten yang diharapkan") {
	// 	t.Errorf("Expected file content to include 'konten yang diharapkan'")
	// }
}
