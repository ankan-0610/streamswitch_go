.PHONY: all

all: run_text_streamer1 run_text_streamer2 run_go_server curl_stream

run_text_streamer1:
	@tmux new-session -d -s mysession -n text_streamer1 "python3 text_streamer1.py"

run_text_streamer2:
	@tmux new-window -t mysession:1 -n text_streamer2 "python3 text_streamer2.py"

run_go_server:
	@tmux new-window -t mysession:2 -n go_server "go run ."

curl_stream:
	@sleep 5 # Give the server time to start
	@tmux new-window -t mysession:3 -n curl_stream "curl localhost:8080/stream"
	@tmux attach -t mysession