package pubchem_dirs

// Domain specifies the data we are searching for
type Domain int32

// Domain enumeration
const (
	Substance Domain = iota
	Compound
	Assay
)

var domains = [3]string{"substance", "compound", "assay"}

func (d Domain) String() string {
	return domains[d]
}
