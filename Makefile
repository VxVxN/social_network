# build:
# 	go build -o bin/social_network cmd/social_network/social_network.go

run:
	go run cmd/reverse_proxy_server/reverse_proxy_server.go &
	go run cmd/web_server/web_server.go &
	go run cmd/ajax_server/ajax_server.go &

stop:
	kill $$(ps | grep reverse_proxy_s | awk '{ print $$1 }')
	kill $$(ps | grep web_server | awk '{ print $$1 }')
	kill $$(ps | grep ajax_server | awk '{ print $$1 }')