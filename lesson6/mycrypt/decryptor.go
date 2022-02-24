package mycrypt

type Decryptor struct {
	fileHash   string
	hashString string
	fileSource string
	fileSigned string
}

//func NewDecryptor(fileSource string, fileHashSign string) (enc *FileEncoder, err error) {
//	hashString, err := ioutil.ReadFile(fileSource)
//	if err != nil {
//		return
//	}
//
//	enc = &FileEncoder{fileSource}
//	return
//}
