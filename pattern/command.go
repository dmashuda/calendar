package pattern

type Command interface {
	Execute() error
	Undo() error
}
