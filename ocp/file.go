package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"gopkg.in/yaml.v2"
)

// Config oc config view --minify
type Config struct {
	Clusters []Clusters `yaml:"clusters"`
}

// Clusters {range .clusters[*]}
type Clusters struct {
	Name    string            `yaml:"name"`
	Cluster map[string]string `yaml:"cluster"`
}

func setPassword(cfg Env) Env {
	cfg.Password = encrypt(askPassword())
	file := getConfigfile()

	d, err := yaml.Marshal(&cfg)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(file)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(file, d, 0777)
	if err != nil {
		panic(err)
	}

	f.Close()

	return cfg
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)

	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func getConfig(file string) Env {
	var c Env
	yamlFile, err := ioutil.ReadFile(file)

	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, &c)

	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func getContext() string {
	// var jsonParse string = "-o=jsonpath='{range .clusters[*]}{.cluster}'"
	var jsonParse string = "-o=json"

	cmd := exec.Command("oc", "config", "view", "--minify", jsonParse)

	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	if err := cmd.Start(); err != nil {
		panic(err)
	}

	errByte, _ := ioutil.ReadAll(stderr)

	if len(errByte) > 0 {
		panic(string(errByte))
	}

	stdByte, _ := ioutil.ReadAll(stdout)
	var m Config
	err := yaml.Unmarshal(stdByte, &m)

	if err != nil {
		panic(err)
	}

	return m.Clusters[0].Cluster["server"]
}
