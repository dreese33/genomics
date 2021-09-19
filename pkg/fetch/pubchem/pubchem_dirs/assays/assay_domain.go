package assays

// AssayDomain queries for procedures
type AssayDomain int32

// AssayDomain enumeration
const (
	AID AssayDomain = iota
	ListKey
	Type
	SourceAll
	Target
	Activity
)

var assayDomains = [6]string{
	"aid",
	"listkey",
	"type",
	"sourceall",
	"target",
	"activity",
}

func (ad AssayDomain) String() string {
	return assayDomains[ad]
}
