package pubchem_dirs

// StructureSearch search by chemical structure
type StructureSearch int32

// StructureSearch enumeration
const (
	Substructure StructureSearch = iota
	Superstructure
	Similarity
	Identity
)

var structureSearches = [4]string{"substructure", "superstructure", "similarity", "identity"}

func (ss StructureSearch) String() string {
	return structureSearches[ss]
}
