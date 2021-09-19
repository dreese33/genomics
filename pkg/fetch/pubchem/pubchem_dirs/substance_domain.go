package pubchem_dirs

// SubstanceDomain
type SubstanceDomain int32

// SubstanceDomain enumeration
const (
	SID SubstanceDomain = iota
	SourceID
	SourceAll
	SubstanceName
	ListKey
)

var substanceDomains = [5]string{
	"sid", "sourceid", "sourceall", "name", "listkey",
}

func (sd SubstanceDomain) String() string {
	return substanceDomains[sd]
}
