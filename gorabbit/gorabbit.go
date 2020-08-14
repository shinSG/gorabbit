package gorabbit

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/streadway/amqp"
)

// RabbitClient rabbitmq client
type RabbitClient struct {
	AmqpConn *amqp.Connection
	HTTPConn *HTTPClient
	BaseURL  string
	ch       *amqp.Channel
}

// URLDict URL mapping
var URLDict = map[string]string{
	"user":       "/users",
	"permission": "/permissions",
	"vhost":      "/vhosts",
	"exchange":   "/exchanges",
	"channel":    "/channels",
	"queue":      "/queues",
}

// Log info
func Log(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		log.Fatalf("%s: %s", msg, err)
	}
}

// Init Get RabbitMQ Client
func (rc *RabbitClient) Init(conf map[string]string) {
	rc.BaseURL = conf["url"]
	host := conf["host"]
	port := conf["port"]
	// SSLPort := conf["sslport"]
	userName := conf["username"]
	password := conf["password"]
	amqpConnStr := fmt.Sprint("amqp://" + userName + ":" + password + "@" + host + ":" + port + "/")
	conn, err := amqp.Dial(amqpConnStr)
	Log(err, "[Connection] RabbitMQ connect error!")
	rc.AmqpConn = conn
	rc.HTTPConn = GetHTTPClient(conf)
	rc.ch, err = rc.AmqpConn.Channel()
	// runtime.SetFinalizer(rc, func(rc *RabbitClient) { rc.Close() })
}

// Close close all connection
func (rc *RabbitClient) Close() {
	fmt.Println("Start close connection")
	rc.ch.Close()
	rc.AmqpConn.Close()
	rc.HTTPConn.Client.CloseIdleConnections()
	fmt.Println("All Closed!!!")
}

// func (rc *RabbitClient) reconnect() {
// 	if rc.ch.
// }

// CreateQueue xxx
func (rc *RabbitClient) CreateQueue(name, key, exchange string, durable, autoDelete, exclusive, noWait bool) amqp.Queue {
	queue, err := rc.ch.QueueDeclare(name, durable, autoDelete, exclusive, noWait, nil)
	if err != nil {
		Log(err, "Create Queue Error")
	}
	if key != "" && exchange != "" {
		rc.ch.QueueBind(name, key, exchange, noWait, nil)
	}
	return queue
}

// DeleteQueue del queue
func (rc *RabbitClient) DeleteQueue(name string, ifUnused, ifEmpty, noWait bool) {
	MsgNo, err := rc.ch.QueueDelete(name, ifUnused, ifEmpty, noWait)
	if err != nil {
		Log(err, "Create Queue Error")
	}
	fmt.Println(MsgNo)
}

// CreateExchange create an exchange
func (rc *RabbitClient) CreateExchange(name, kind string, durable, autoDelete, internal, noWait bool) bool {
	err := rc.ch.ExchangeDeclare(name, kind, durable, autoDelete, internal, noWait, nil)
	if err != nil {
		Log(err, "Create Queue Error")
		return false
	}
	return true
}

// GetRespBody xxx
func GetRespBody(resp *http.Response) string {
	cont, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(cont)
}

// GetQueues get queue list
func (rc *RabbitClient) GetQueues(name string) {
	url := rc.BaseURL + URLDict["queue"]
	if name != "" {
		url = url + "/" + name
	}
	resp, _ := rc.HTTPConn.Get(url)
	body := GetRespBody(resp)
	fmt.Println(body)
}

//GetExchanges get exchange list
func (rc *RabbitClient) GetExchanges(name string) {
	url := rc.BaseURL + URLDict["queue"]
	if name != "" {
		url = url + "/" + name
	}
	resp, _ := rc.HTTPConn.Get(url)
	body := GetRespBody(resp)
	fmt.Println(body)
}

// func (rc *RabbitClient) Disconnect() {
// 	rc.ch.Disconnect()
// }

// // Do send http request
// func (rc *RabbitClient) Do(method string, url string, body io.Reader) (*http.Response, error) {
// 	client = http.Client{}
// }

// // GetUsers get users list
// func (rc *RabbitClient) GetUsers() []map[string]string {
// 	url := rc.BaseURL + URLDict["user"]
// 	req, err := http.NewRequest("GET", url, nil)
// 	req.Header.Set("content-type", "application/json")
// 	Log(err, "[User] Get User Info Failed:")
// 	respData := resp.Body
// }
