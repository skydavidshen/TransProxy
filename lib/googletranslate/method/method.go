package method

type Method interface {
	Content(resourceUrl string) ([]byte, error)
}
