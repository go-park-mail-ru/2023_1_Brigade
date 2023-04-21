package consumer

type Usecase interface {
	ConsumeMessage() []byte
	StartConsumeMessages()
}
