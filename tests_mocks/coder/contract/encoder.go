package contract

type Encoder interface {
	Encode(string) error
	SaveToFile(string) error
}
