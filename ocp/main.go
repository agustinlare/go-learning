package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

// Context - stop asking me for the comment
type Context struct {
	Name     string `yaml:"name"`
	Url      string `yaml:"url"`
	Contexto string `yaml:"contexto"`
}

// Env - stop aspking me for the the comment
type Env struct {
	Password string    `yaml:"password"`
	Context  []Context `yaml:"env"`
}

func main() {
	var argsLens int = len(os.Args)
	configFile := getConfigfile()

	if !fileExists(configFile) {
		fmt.Println(configFile)
		panic("No existe archivo de configuracion.")
	}

	var resp []byte
	n := "default"

	switch {
	case argsLens == 1:
		resp = loginCluster(selectCluster(configFile), n)
	case argsLens > 2 && os.Args[1] == "-n":
		resp = ocCaller([]string{"project", os.Args[2]})
	case argsLens == 2 && os.Args[1] != "project":
		var cfg Env
		var url string

		cfg = getConfig(configFile)
		for _, v := range cfg.Context {
			if v.Name == os.Args[1] {
				url = v.Url
				break
			}
		}

		if url == "" {
			log.Println("WARNING\t| Contexto \" " + os.Args[1] + " \" no encontrado.")
		}

		resp = loginCluster(url, os.Args[1])
	default:
		var args []string

		for _, v := range os.Args[1:] {
			args = append(args, v)
		}

		resp = ocCaller(args)
	}

	fmt.Println(string(resp))
}

func loginCluster(c string, n string) []byte {
	var namespace string = "--namespace=" + n
	var server = "--server=" + c

	username := "--username=" + getUsername()
	password := "--password=" + decrypt()

	resp := ocCaller([]string{"login", username, password, server, "--insecure-skip-tls-verify=true", namespace})

	return resp
}

func selectCluster(file string) string {
	var cfg Env
	var o int
	var separator string = "-----------------------------------------------------------"
	var header string = "N | NAME\t| URL"

	cfg = getConfig(file)
	fmt.Println("Ingrese opcion:\n" + header + "\n" + separator)

	for i, v := range cfg.Context {
		fmt.Println(i, "|", v.Name, "\t|", v.Url)
	}

	fmt.Println(separator)
	fmt.Scan(&o)

	return cfg.Context[o].Url
}

func ocCaller(args []string) []byte {
	var stdByte []byte
	cmd := exec.Command("oc", args...)
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	if err := cmd.Start(); err != nil {
		panic(err)
	}

	errByte, _ := ioutil.ReadAll(stderr)

	if len(errByte) > 0 {
		stdByte = tryThis(string(errByte))
	}

	stdByte, _ = ioutil.ReadAll(stdout)

	return stdByte
}
