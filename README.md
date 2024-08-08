# Text Streaming Service with Dynamic Provider Switching

This project demonstrates a text streaming service where text is dynamically streamed from multiple providers (implemented in Python) using a backend server written in Go.
The Go server dynamically switches between providers based on response delays, ensuring that the streaming continues smoothly even if one of the providers is slow or unresponsive.

## Setup Instructions
Run the following command

```bash
make
```


## Overview
The Go server serves as the backend for the text streaming service. It fetches text from multiple Python-based providers, dynamically switching between them based on their response times.
The providers return blocks of text, one line at a time, with simulated delays. 


## Dynamic Switching Logic

### 1. Providers:
The providers (`text_streamer1.py` and `text_streamer2.py`) are simple Flask applications that return lines of text with a delay between each line. 
Each request to a provider includes an `index` parameter, specifying which line to return.

### 2. Round-Robin selection:
The Go server uses a round-robin approach to switch between providers. 
The providerIndex tracks the current provider, and the textIndex tracks the line of text to request.

### 3. Delay Threshold:
The server sets a delay threshold (delayThreshold), which defines the maximum acceptable delay for receiving a response from a provider. 
If the delay exceeds this threshold, the server switches to the next provider.

## Contributing
We welcome suggestions for new features or improvements. If you have an idea, please open an issue to discuss it.

## Acknowledgement:
I would like to thank Alle and their team for providing this opportunity to exercise my Golang skills.
