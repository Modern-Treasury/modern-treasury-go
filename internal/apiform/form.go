package apiform

type Marshaler interface {
	MarshalMultipart() ([]byte, error)
}
