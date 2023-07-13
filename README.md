# Config Maker

A tool for making configuration from file and environments conveniently.

## Usage

full example please go to `example/demo01`

### Installation
```bash
$ go get github.com/yukiouma/cfg-maker
```

### Assumption
We want to store the configurations using a struct like this:
```go
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

```

most of information will be store in a file, for example in a yaml file like this:
```yaml
server:
  grpc:
    host: 0.0.0.0
data:
  db:
    host: 127.0.0.1
    user: root
    database: hello
```
The rest of messages may be set into environment variables, like this:
```bash
echo $password
123456
echo $asroot
true
echo $cpulimit
0.5
echo $mysqlport
3306
echo $serverport
8080
```

And then, we can use cfg-reader to build the configuration struct easliy.

We should read the configuration file first, to initialize the struct
```go
reader := cfgmaker.New(&Bootstrap{})
reader.ReadFromYamlFile("./config.yaml")
```

If we want to set the env variable `$password` as the field of `Bootstrap.Data.DB.Password`, we can declare field using the chain of json tag `data.db.passwor`, like this
```go
reader.ReadFromEnv("data.db.password", "password")
```

Finally we can read all the imformation as follow:
```go
func main() {
	reader := cfgmaker.New(&Bootstrap{})
	reader.ReadFromYamlFile("./config.yaml").
		ReadFromEnv("data.db.password", "password").
		ReadFromEnv("asroot", "asroot").
		ReadFromEnv("cpulimit", "cpulimit").
		ReadFromEnv("data.db.port", "mysqlport").
		ReadFromEnv("server.grpc.port", "serverport")
	if err := reader.Err(); err != nil {
		log.Fatal(err)
	}
    cfg := reader1.Get().(*Bootstrap)
	fmt.Printf("%#v\n", cfg)
}
```



## TODO
* support TOML
* support XML
* support Map and Slice