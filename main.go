/*
 * @Author: cuihaonan
 * @Email: devcui@outlook.com
 * @Date: 2021-04-29 13:43:26
 * @LastEditTime: 2021-04-29 13:57:07
 * @LastEditors: cuihaonan
 * @Description: Basic description
 * @FilePath: /helm-char-push/main.go
 * @LICENSE: NONE
 */
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	filePath = flag.String("file", GetEnvString("YAML_FILE_PATH", "Chart.yaml"), "yaml file path")
)

type ChartYaml struct {
	ApiVersion   string
	Name         string
	Description  string
	Version      string
	AppVersion   string
	Dependencies []struct {
		Name       string
		Repository string
		Version    string
	}
	Icon        string
	Maintainers []struct {
		Email string
		Name  string
	}
}

func main() {
	flag.Parse()
	chart := ChartYaml{}
	b, e := ioutil.ReadFile(*filePath)
	if e != nil {
		panic(e)
	}
	yaml.Unmarshal(b, &chart)
	fmt.Printf("%s-%s.tgz", chart.Name, chart.Version)
}

func GetEnvString(key string, defaultValue string) string {
	if envVal, ok := os.LookupEnv(key); ok {
		return envVal
	}
	return defaultValue
}
