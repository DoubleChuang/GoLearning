package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

func httpGet(path string) (string, error) {
	resp, err := http.Get(path)
	if err != nil {
		fmt.Println("Fail:", err)
		flag.PrintDefaults()
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Fail:", err)
	}
	switch s := resp.StatusCode; {
	case s >= 300 && s < 400:
		return "", errors.New(resp.Status)
	case s >= 400 && s < 500:
		return "", errors.New(resp.Status)
	case s >= 200 && s < 300:
	default:
		return "", errors.New(resp.Status)
	}
	return string(body), nil
}

var (
	help bool
	url  string
	port string
	user string
	pswd string
)

func init() {
	flag.BoolVar(&help, "h", false, "help")
	flag.StringVar(&url, "u", "localhost", "Get url")
	flag.StringVar(&port, "p", "80", "Get port")
	flag.StringVar(&user, "s", "admin", "User")
	flag.StringVar(&pswd, "w", "admin", "Password")
}

func main() {
	flag.Parse()
	if help {
		flag.Usage()
	} else {
		fmt.Println(httpGet("http://" + user + ":" + pswd + "@" + url + ":" + port + "/Telnetd.cgi?Telnetd=Open"))
	}
}
