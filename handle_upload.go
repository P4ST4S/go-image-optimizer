package main

import (
	"fmt"
	"image"
	"net/http"
	"time"
)

// Semaphore to limit concurrent uploads (max 5 simultaneous)
var uploadSemaphore = make(chan struct{}, 5)

// handleUpload handles image upload and processing
func handleUpload(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	fmt.Printf("🔵 Request received: %s %s\n", r.Method, r.URL.Path)

	if r.Method != http.MethodPost {
		fmt.Printf("❌ Method not allowed: %s\n", r.Method)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Try to acquire semaphore slot
	select {
	case uploadSemaphore <- struct{}{}:
		// Got a slot, continue processing
		fmt.Printf("✅ Semaphore acquired (slots available: %d/5)\n", 5-len(uploadSemaphore))
		defer func() {
			<-uploadSemaphore
			fmt.Printf("🔓 Semaphore released (slots available: %d/5)\n", 5-len(uploadSemaphore))
		}()
	default:
		// No slot available, reject request
		fmt.Printf("⛔ Server too busy - all 5 slots occupied\n")
		http.Error(w, "Server too busy, try again later", http.StatusServiceUnavailable)
		return
	}

	fmt.Printf("📦 Parsing multipart form...\n")
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		fmt.Printf("❌ Failed to parse form: %v\n", err)
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	fmt.Printf("📂 Retrieving file from form...\n")
	file, header, err := r.FormFile("image")
	if err != nil {
		fmt.Printf("❌ Error retrieving file: %v\n", err)
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}
	defer file.Close()
	fmt.Printf("📁 File received: %s (%d bytes)\n", header.Filename, header.Size)

	fmt.Printf("🖼️  Decoding image...\n")
	img, format, err := image.Decode(file)
	if err != nil {
		fmt.Printf("❌ Invalid image format: %v\n", err)
		http.Error(w, "Invalid image format", http.StatusBadRequest)
		return
	}
	fmt.Printf("📸 Image decoded: Format %s\n", format)

	fmt.Printf("🔄 Resizing image...\n")
	buf, err := ResizeImage(img, 300, 80)
	if err != nil {
		fmt.Printf("❌ Error during compression: %v\n", err)
		http.Error(w, "Error during compression", http.StatusInternalServerError)
		return
	}
	fmt.Printf("✨ Image resized successfully (%d bytes)\n", buf.Len())

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", fmt.Sprint(buf.Len()))
	bytesWritten, err := w.Write(buf.Bytes())
	if err != nil {
		fmt.Printf("❌ Error writing response: %v\n", err)
		return
	}

	fmt.Printf("✅ Request completed in %s (%d bytes sent)\n", time.Since(start), bytesWritten)
}
