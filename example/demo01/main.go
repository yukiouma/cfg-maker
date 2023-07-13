package main

import (
	"fmt"
	"log"
	"os"

	cfgreader "github.com/yukiouma/cfg-maker"
)

func init() {
	os.Setenv("password", "123456")
	os.Setenv("asroot", "true")
	os.Setenv("cpulimit", "0.5")
	os.Setenv("mysqlport", "3306")
	os.Setenv("serverport", "8080")
}

func main() {
	reader1 := cfgreader.New(&Bootstrap{})
	reader1.ReadFromYamlFile("./config.yaml").
		ReadFromEnv("data.db.password", "password").
		ReadFromEnv("asroot", "asroot").
		ReadFromEnv("cpulimit", "cpulimit").
		ReadFromEnv("data.db.port", "mysqlport").
		ReadFromEnv("server.grpc.port", "serverport")
	if err := reader1.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", reader1.Get().(*Bootstrap))

	reader2 := cfgreader.New(&Bootstrap{})
	reader2.ReadFromJsonFile("./config.json").
		ReadFromEnv("data.db.password", "password").
		ReadFromEnv("asroot", "asroot").
		ReadFromEnv("cpulimit", "cpulimit").
		ReadFromEnv("data.db.port", "mysqlport").
		ReadFromEnv("server.grpc.port", "serverport")
	if err := reader2.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", reader2.Get().(*Bootstrap))
}
