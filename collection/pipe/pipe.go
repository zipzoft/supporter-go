package pipe

type Pipe interface {
	Handle(interface{}) (interface{}, error)
}
