package other

type Custom int

//go:generate smock -debug
type Bar interface {
	Do(func(Custom) Custom) Custom
}
