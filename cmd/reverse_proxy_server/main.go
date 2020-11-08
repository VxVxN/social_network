package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"

	cnfg "github.com/VxVxN/social_network/app/config"
	"github.com/VxVxN/social_network/app/log"
	"github.com/VxVxN/social_network/cmd/reverse_proxy_server/context"
)

var protocol = "http://"

func getProxyURL(url string) string {
	var host string
	if strings.Contains(url, "/ajax/") {
		host = protocol + cnfg.Config.AJAXServerHostname + ":" + strconv.Itoa(cnfg.Config.AJAXServerPort)
		return host
	}

	host = protocol + cnfg.Config.WebServerHostname + ":" + strconv.Itoa(cnfg.Config.WebServerPort)
	return host
}

func serveReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
	url, _ := url.Parse(target)

	// create the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(url)

	// Update the headers to allow for SSL redirection
	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = url.Host

	// Note that ServeHttp is non blocking and uses a go routine under the hood
	proxy.ServeHTTP(res, req)
}

func handleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {
	url := getProxyURL(req.URL.String())

	serveReverseProxy(url, res, req)
}

func main() {
	context := &context.Context{Log: log.Init("reverse_proxy_server.log", false)}

	context.Log.Info.Println("Reverse proxy server start.")

	http.HandleFunc("/", handleRequestAndRedirect)

	port := ":" + strconv.Itoa(cnfg.Config.ReverseProxyServerPort)
	if err := http.ListenAndServe(port, nil); err != nil {
		context.Log.Fatal.Printf("Failed to listen and serve port: %v. Error: %v", port, err)
		panic(err)
	}
}
