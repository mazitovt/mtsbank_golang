package contract

type Decoder interface {
	Decode(string, string) error
	SaveToFile(string) error
}
