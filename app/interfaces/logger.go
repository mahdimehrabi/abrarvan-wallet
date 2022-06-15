package interfaces

//our logger interface that every part
//of our application use that for getting logger
//for saving time I keep this interface simple
type Logger interface {
	Error(err string)
}
