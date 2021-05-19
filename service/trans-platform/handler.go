package trans_platform

type Handler interface {
	Translate(to, text string) (string, error)
}
