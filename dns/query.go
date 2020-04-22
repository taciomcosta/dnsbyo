package dns

type Query struct {
	Name   string
	QClass Class
	QType  Type
}

type Class uint16
type Type uint16

const (
	QClassIN = 1
)

const (
	QTypeA = 1
)
