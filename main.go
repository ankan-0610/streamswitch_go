package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const delayThreshold = 500 * time.Millisecond // Define a delay threshold

var providerIndex = 0 // Start with the first provider
var textIndex = 0      // Index of the text to stream

// streamText handles the streaming of text from the selected provider
func streamText(w http.ResponseWriter, r *http.Request) {

	for {
		currentProvider := selectProvider(providerIndex)
		urlWithIndex := fmt.Sprintf("%s?index=%d", currentProvider.URL, textIndex)

		resp, err := http.Get(urlWithIndex)
		if err == nil {
			defer resp.Body.Close()
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusOK)

			buf := make([]byte, 1024)
            loop:
			for {
				chunkChan := make(chan []byte)
				errChan := make(chan error)

				// Read from the response body in a separate goroutine
				go func() {
					n, err := resp.Body.Read(buf)
					if n > 0 {
						chunkChan <- buf[:n]
					} else {
						errChan <- err
					}
				}()

				select {
				case chunk := <-chunkChan:
					// Received a chunk of text
					w.Write(chunk)
					fmt.Println(string(chunk))
					if f, ok := w.(http.Flusher); ok {
						f.Flush()
					}
                    textIndex += 1
				case err := <-errChan:
					if err == io.EOF {
						return
					}
					if err != nil {
						log.Printf("Error reading from provider %s: %v", currentProvider.Name, err)
						// Move to the next provider in a round-robin fashion
		                providerIndex = (providerIndex + 1) % len(providers)
                        break loop
					}
				case <-time.After(delayThreshold):
					// If no chunk is received within the delay threshold, switch provider
					log.Printf("Provider %s exceeded delay threshold, switching to another provider", currentProvider.Name)
                    // Move to the next provider in a round-robin fashion
		            providerIndex = (providerIndex + 1) % len(providers)
					break loop
				}
			}
		} else {
			log.Printf("Failed to connect to provider %s, error: %v", currentProvider.Name, err)
		}
        if textIndex==4{
            break
        }
	}

	// log.Fatal("All providers failed")
}

func main() {
	http.HandleFunc("/stream", streamText)
	log.Fatal(http.ListenAndServe(":8080", nil))
}