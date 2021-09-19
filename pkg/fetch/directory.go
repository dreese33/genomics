package fetch

type Directory interface {
	ValidateDirectory()
	ConstructDirectory()
}
