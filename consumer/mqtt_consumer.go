package consumer

import (
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
)

type Consumer struct {
	client        MQTT.Client
	Channel        chan string
}

func NewConsumer(opts map[string]interface{}) (consumer *Consumer){
	clientOptions := MQTT.NewClientOptions()
	brokerURI := fmt.Sprintf("%v:%v", opts["host"],opts["port"])
	userName := opts["username"].(string)
	password := opts["password"].(string)

	clientOptions.AddBroker(brokerURI)
	clientOptions.SetUsername(userName)
	clientOptions.SetPassword(password)
	clientOptions.SetAutoReconnect(true)
	
	return &Consumer{
		client: MQTT.NewClient(clientOptions),
		Channel: make(chan string),
	}

}

func (c *Consumer) Connect(){
	if token := c.client.Connect(); token.Wait() && token.Error() != nil {
		log.WithFields(log.Fields{
        "error":    token.Error(),
    	}).Fatal("Could not conect to MQTT broker.")
		
	}
	log.Info("Connected to MQTT broker.")
}

func (c *Consumer) Subscribe(topic string){
	if c.client.IsConnected(){
		c.client.Subscribe(topic, 0, func(client MQTT.Client, msg MQTT.Message) {
			c.Channel <- string(msg.Payload()) 
		})
	}else{

	}
}

func (c *Consumer) Disconnect(){
	if c.client.IsConnected(){
		c.client.Disconnect(1000)		
	}
	
}