package assays

// AssayTarget query for procedures with additional subdirectories
type AssayTarget int32

// AssayTarget enumeration
const (
	GI AssayTarget = iota
	ProteinName
	GeneID
	GeneSymbol
	Accession
)

var assayTargets = [5]string {
	"gi",
	"proteinname",
	"geneid",
	"genesymbol",
	"accession",
}

func (at AssayTarget) String() string {
	return assayTargets[at]
}