package CfgMaker

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	_JSON = "json"
)

type CfgMaker interface {
	ReadFromYamlFile(string) CfgMaker
	ReadFromJsonFile(string) CfgMaker
	ReadFromEnv(key, envKey string) CfgMaker
	Get() any
	Err() error
}

// initialize reader, please provide a pointer which is not nil
func NewMaker(data any) CfgMaker {
	var c = new(cfg)
	rv := reflect.ValueOf(data)
	// if receive a nil or not a pointer, set error
	if rv.Kind() != reflect.Pointer || rv.IsNil() {
		c.err = fmt.Errorf("please provide a pointer which is not nil")
		return c
	}

	c.err = initFields(rv)
	c.data = data
	return c
}

type cfg struct {
	data any
	err  error
}

func (c *cfg) ReadFromYamlFile(dir string) CfgMaker {
	if c.err != nil {
		return c
	}
	data, err := os.ReadFile(dir)
	if err != nil {
		c.err = err
		return c
	}
	if c.err = yaml.Unmarshal(data, c.data); c.err != nil {
		return c
	}
	c.err = initFields(reflect.ValueOf(c.data))
	return c
}

func (c *cfg) ReadFromJsonFile(dir string) CfgMaker {
	if c.err != nil {
		return c
	}
	data, err := os.ReadFile(dir)
	if err != nil {
		c.err = err
		return c
	}
	if c.err = json.Unmarshal(data, c.data); c.err != nil {
		return c
	}
	c.err = initFields(reflect.ValueOf(c.data))
	return c
}

// set value to specify field using environment variable
func (c *cfg) ReadFromEnv(key, envKey string) CfgMaker {
	if c.err != nil {
		return c
	}
	value := os.Getenv(envKey)
	v, err := c.locateField(strings.Split(key, ".")...)
	if err != nil {
		c.err = err
		return c
	}
	if !v.CanSet() {
		c.err = fmt.Errorf("error: field %s is not addressable", key)
		return c
	}
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		var val int64
		val, _ = strconv.ParseInt(value, 10, 64)
		v.SetInt(val)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		var val uint64
		val, _ = strconv.ParseUint(value, 10, 64)
		v.SetUint(val)
	case reflect.Float32, reflect.Float64:
		var val float64
		val, _ = strconv.ParseFloat(value, 64)
		v.SetFloat(val)
	case reflect.Bool:
		v.SetBool(value == "true")
	default:
		c.err = fmt.Errorf("field should be string, integer, float or boolean, current \"%s\" is %s", key, v.Kind())
	}
	return c
}

// find field location by the tag chains
func (c *cfg) locateField(tags ...string) (v reflect.Value, err error) {
	if c.err != nil {
		err = c.err
		return
	}
	tailIndex := len(tags) - 1
	key := tags[tailIndex]
	tags = tags[:tailIndex]
	rv := reflect.ValueOf(c.data)
	for _, tag := range tags {
		rv = reflect.Indirect(rv)
		kind := rv.Kind()
		if rv.Kind() != reflect.Struct {
			c.err = fmt.Errorf("error: type %s is invalid, prefix tags should map to pointer or struct", kind)
			return
		}
		for i := 0; i < rv.NumField(); i++ {
			field := rv.Type().Field(i)
			fieldTag := field.Tag.Get(_JSON)
			if strings.Contains(fieldTag, tag) {
				rv = rv.Field(i)
				break
			}
		}
	}
	rv = reflect.Indirect(rv)
	for i := 0; i < rv.NumField(); i++ {
		field := rv.Type().Field(i)
		tagValue := field.Tag.Get("json")
		if strings.Contains(tagValue, key) {
			v = rv.Field(i)
			break
		}
	}
	return
}

func (c *cfg) Get() any {
	return c.data
}

func (c *cfg) Err() error {
	return c.err
}

// initialize all the fields which are nil pointer
func initFields(rv reflect.Value) error {
	rv = reflect.Indirect(rv)
	if !rv.CanSet() {
		return fmt.Errorf("field %s is not addressable", rv.Type().Name())
	}
	if rv.Kind() != reflect.Struct {
		return nil
	}
	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i)
		if !rv.CanSet() {
			return fmt.Errorf("field %s is not addressable", field.Type().Name())
		}
		if field.Kind() == reflect.Pointer && field.IsNil() {
			field.Set(reflect.New(field.Type().Elem()))
		}
		if err := initFields(field); err != nil {
			return err
		}
	}
	return nil
}
