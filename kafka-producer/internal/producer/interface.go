package producer

type Produce interface {
	Retry(value []byte) error
	Produce(value []byte) error
}
