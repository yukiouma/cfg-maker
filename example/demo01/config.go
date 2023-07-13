package main

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
