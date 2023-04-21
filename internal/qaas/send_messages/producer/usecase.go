package producer

type Usecase interface {
	ProduceMessage(message []byte) error
}
