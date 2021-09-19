package pubchem_dirs

// CompoundDomain search by criteria
type CompoundDomain int32

// CompoundDomain types enumeration
const (
	Cid CompoundDomain = iota
	Name
	Smiles
	Inchi
	Sdf
	Inchikey
	Formula
	Listkey
)

var compoundDomains = [8]string{
	"cid", "name", "smiles", "inchi", "sdf", "inchikey", "formula", "listkey",
}

func (cd CompoundDomain) String() string {
	return compoundDomains[cd]
}
