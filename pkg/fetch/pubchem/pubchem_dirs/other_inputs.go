package pubchem_dirs

// OtherInputs are other values to query by
type OtherInputs int32

// OtherInputs enumeration
const (
	Sources OtherInputs = iota
	SourceTable
	Conformers
	Annotations
	SourceName
	Heading
)

var otherInputs = [6]string{
	"sources",
	"sourcetable",
	"conformers",
	"annotations",
	"sourcename",
	"heading",
}

func (ot OtherInputs) String() string {
	return otherInputs[ot]
}
