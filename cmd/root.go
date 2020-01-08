/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"plugin"
	"github.com/spf13/viper"
	"github.com/spf13/cobra"
	Consumer "github.com/nikhilfernandes/mqtt-consumer/consumer"
)

var cfgFile string
var topic string
var sink string

type Connector interface {
	Sink(msg string)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mqtt-consumer",
	Short: "Go consumer to consume messages from mqtt broker.",
	Long: `A MQTT consumer service written in Golang, that interacts with an MQTT broker and 
	sinks the messages to a kafka broker.`,
	
	Run: func(cmd *cobra.Command, args []string) { 
		consumer := Consumer.NewConsumer(viper.GetStringMap("emqx"))
		topic, _:= cmd.Flags().GetString("topic")
		sink, _:= cmd.Flags().GetString("sink")
		connector := parseSinkToGetPlug(sink)
		
		consumer.Connect()
		consumer.Subscribe(topic)
		for message := range consumer.Channel {
			connector.Sink(message) 
		}

	},
}


func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is current directory config.yml)")
	rootCmd.Flags().StringVarP(&topic, "topic", "t", "", "MQTT broker topic name to consume from. (required)")
	rootCmd.Flags().StringVarP(&sink, "sink", "s", "stdout", "Sink the messages consumed from MQTT broker too. (default is stdout)")
	rootCmd.MarkFlagRequired("topic")
	
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {	
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func parseSinkToGetPlug(sink string) (connector Connector){
	var mod string
	switch sink {
	case "stdout":
		mod = "./stdout-sink/connector.so"
	case "kafka":
		mod = "./kafka-sink/connector.so"
	default:
		fmt.Println("specify the correct sink (stdout,kafka)")
		os.Exit(1)
	}
	plug, _ := plugin.Open(mod)
	symConnector, _ := plug.Lookup("Connector")
	connector, _ = symConnector.(Connector)
	return connector
}
