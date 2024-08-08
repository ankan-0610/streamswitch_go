from flask import Flask, Response, request
import time

app = Flask(__name__)

# List of texts to stream
texts = ["Hello, world!", "This is a text stream.", "Streaming text line by line.", "Enjoy the stream!"]

def generate_text(idx):
    """Generator that yields the specific text line."""
    yield f"{texts[idx]} <= from streamer 1\n"
    time.sleep(1)  # Simulate a delay in streaming

@app.route('/')
def stream_text():
    index = int(request.args.get('index', 0))  # Get the index from the query parameter
    response = Response(generate_text(index % len(texts)), mimetype='text/plain')
    return response

if __name__ == '__main__':
    app.run(port=5000)