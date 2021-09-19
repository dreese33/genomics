package fetch

type Directory interface {
	ValidateDirectory() bool
	ConstructDirectory() string
}
