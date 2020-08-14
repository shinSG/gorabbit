package main

import (
	"gorabbit/gorabbit"
)

func main() {
	// hc = new(gorabbit.HTTPClient)
	conf := make(map[string]string)
	conf["username"] = "guest"
	conf["password"] = "guest"
	conf["host"] = "127.0.01"
	conf["port"] = "5672"
	conf["url"] = "http://127.0.0.1:15672/api"
	// hc := gorabbit.GetHTTPClient(conf)
	// resp, _ := hc.Get("http://127.0.0.1:15672/api/users")
	// fmt.Println(resp.StatusCode)
	// data, _ := ioutil.ReadAll(resp.Body)
	rc := new(gorabbit.RabbitClient)
	rc.Init(conf)
	rc.CreateQueue("test", "", "", false, false, false, false)
	// rc.DeleteQueue("test", false, false, false)
	rc.GetQueues("")
	// rc.Close()
	// fmt.Println(string(data))
}
