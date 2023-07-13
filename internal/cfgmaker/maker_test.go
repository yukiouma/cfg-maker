package CfgMaker

import (
	"os"
	"testing"
)

func TestMakeConfig(t *testing.T) {
	os.Setenv("password", "123456")
	os.Setenv("asroot", "true")
	os.Setenv("cpulimit", "0.5")
	os.Setenv("mysqlport", "3306")
	os.Setenv("serverport", "8080")
	jsonFileDir := "/root/playground/golang/misc-playground/cfg-maker/example/demo01/config.json"
	yamlFileDir := "/root/playground/golang/misc-playground/cfg-maker/example/demo01/config.yaml"
	bs1 := new(Bootstrap)
	m1 := NewMaker(bs1)
	m1.ReadFromYamlFile(yamlFileDir).
		ReadFromEnv("data.db.password", "password").
		ReadFromEnv("asroot", "asroot").
		ReadFromEnv("cpulimit", "cpulimit").
		ReadFromEnv("data.db.port", "mysqlport").
		ReadFromEnv("server.grpc.port", "serverport")

	if err := m1.Err(); err != nil {
		t.Fatalf("make config failed because: %v", err)
	}

	bs2 := new(Bootstrap)
	m2 := NewMaker(bs2)
	m2.ReadFromYamlFile(jsonFileDir).
		ReadFromEnv("data.db.password", "password").
		ReadFromEnv("asroot", "asroot").
		ReadFromEnv("cpulimit", "cpulimit").
		ReadFromEnv("data.db.port", "mysqlport").
		ReadFromEnv("server.grpc.port", "serverport")

	if err := m2.Err(); err != nil {
		t.Fatalf("make config failed because: %v", err)
	}
}

type Bootstrap struct {
	AsRoot   bool    `json:"asroot,omitempty"`
	CPULimit float64 `json:"cpulimit,omitempty"`
	Server   *Server `json:"server,omitempty"`
	Data     *Data   `json:"data,omitempty"`
}

type Data struct {
	DB *DB `json:"db,omitempty"`
}

type DB struct {
	Host     string `json:"host,omitempty"`
	Port     int    `json:"port,omitempty"`
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
	Database string `json:"database,omitempty"`
}

type Server struct {
	Grpc GRPC `json:"grpc,omitempty"`
}

type GRPC struct {
	Host string `json:"host,omitempty"`
	Port int    `json:"port,omitempty"`
}
