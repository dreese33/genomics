package assays

// AssayType queries for procedures
type AssayType int32

// AssayDomain enumeration
const (
	ALL AssayType = iota
	Confirmatory
	DoseResponse
	OnHold
	Panel
	RNAI
	Screening
	Summary
	CellBased
	BioChemical
	Invivo
	Invitro
	ActiveConcentrationSpecified
)

var AssayTypes = [13]string {
	"all",
	"confirmatory",
	"doseresponse",
	"onhold",
	"panel",
	"rnai",
	"screening",
	"summary",
	"cellbased",
	"biochemical",
	"invivo",
	"invitro",
	"activeconcentrationspecified",
}

func (at AssayType) String() string {
	return AssayTypes[at]
}