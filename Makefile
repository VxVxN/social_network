# build:
# 	go build -o bin/social_network cmd/social_network/social_network.go

run:
	go run cmd/reverse_proxy_server/main.go &
	go run cmd/web_server/main.go &
	go run cmd/ajax_server/main.go &

stop:
	kill $$(ps | grep main | awk '{ print $$1 }')