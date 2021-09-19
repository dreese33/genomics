package pubchem_dirs

// XREF for cross references information
type XREF int32

// XREF enumeration
const (
	RegistryID XREF = iota
	RN
	PubMedID
	MMDBID
	ProteinGI
	NucleotideGI
	TaxonomyID
	MIMID
	GeneID
	ProbeID
	PatentID
)

var xrefs = [11]string{
	"RegistryID",
	"RN",
	"PubMedID",
	"MMDBID",
	"ProteinGI",
	"NucleotideGI",
	"TaxonomyID",
	"MIMID",
	"GeneID",
	"ProbeID",
	"PatentID",
}

func (xr XREF) String() string {
	return xrefs[xr]
}
