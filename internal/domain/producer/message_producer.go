package producer

type MessageProducer interface {
	Publish(queueName string, data []byte) error
}
