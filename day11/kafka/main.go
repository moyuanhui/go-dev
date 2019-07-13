package main

import "github.com/Shopify/sarama"

func main() {

	config := sarama.NewConfig()
	config.Producer.RequireAcks = sarama.WaitForAll
}
