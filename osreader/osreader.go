package osreader

type OSReader interface {
	ReadLines() ([]string, error)
}
