package oswriter

type OSWriter interface {
	WriteJsonToFile(any) error
}
