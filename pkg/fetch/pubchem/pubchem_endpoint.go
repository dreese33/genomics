package pubchem

import (
	"fmt"
	"strings"

	. "github.com/dreese33/genomics/pkg/fetch/pubchem/pubchem_dirs"
	"github.com/dreese33/genomics/pkg/fetch/pubchem/pubchem_dirs/assays"
)

// Full API reference
// Stolen from - https://pubchemdocs.ncbi.nlm.nih.gov/pug-rest
//
// ****************************************************************************************************************************************************************************************************************************
//
// <input specification> = <domain>/<namespace>/<identifiers>
// <domain> = substance | compound | assay | <other inputs>
// compound domain <namespace> = cid | name | smiles | inchi | sdf | inchikey | formula | <structure search> | <xref> | listkey | <fast search>
// <structure search> = {substructure | superstructure | similarity | identity}/{smiles | inchi | sdf | cid}
// <fast search> = {fastidentity | fastsimilarity_2d | fastsimilarity_3d | fastsubstructure | fastsuperstructure}/{smiles | smarts | inchi | sdf | cid} | fastformula
// <xref> = xref / {RegistryID | RN | PubMedID | MMDBID | ProteinGI | NucleotideGI | TaxonomyID | MIMID | GeneID | ProbeID | PatentID}
// substance domain <namespace> = sid | sourceid/<source id> | sourceall/<source name> | name | <xref> | listkey
// <source name> = any valid PubChem depositor name
// assay domain <namespace> = aid | listkey | type/<assay type> | sourceall/<source name> | target/<assay target> | activity/<activity column name>
// <assay type> = all | confirmatory | doseresponse | onhold | panel | rnai | screening | summary | cellbased | biochemical | invivo | invitro | activeconcentrationspecified
// <assay target> = gi | proteinname | geneid | genesymbol | accession
// <identifiers> = comma-separated list of positive integers (e.g. cid, sid, aid) or identifier strings (source, inchikey, formula); in some cases only a single identifier string (name, smiles, xref; inchi, sdf by POST only)
// <other inputs> = sources / [substance, assay] |sourcetable | conformers | annotations/[sourcename/<source name> | heading/<heading>]
//
// ****************************************************************************************************************************************************************************************************************************
//

// PubchemDirectory is the directory for PUG REST API query
type PubchemDirectory struct {
	Dom Domain
	CD  CompoundDomain
	SS  StructureSearch
	FS  FastSearch
	REF XREF
	SD  SubstanceDomain
	OT  OtherInputs
	AD  assays.AssayDomain
	ATG assays.AssayTarget
	AT  assays.AssayType
}

// FillDefaults sets all values to -1
func (directory *PubchemDirectory) FillDefaults() {
	directory.Dom = -1
	directory.CD = -1
	directory.SS = -1
	directory.FS = -1
	directory.REF = -1
	directory.SD = -1
	directory.OT = -1
	directory.AD = -1
	directory.ATG = -1
	directory.AT = -1
}

// NewPubchemDirectory constructs a PubchemDirectory object
func NewPubchemDirectory(subDirectories ...interface{}) PubchemDirectory {
	var pugDirectory = new(PubchemDirectory)
	pugDirectory.FillDefaults()

	for _, subDirectory := range subDirectories {
		switch subDirType := subDirectory.(type) {
		case Domain:
			pugDirectory.Dom = subDirectory.(Domain)
		case CompoundDomain:
			pugDirectory.CD = subDirectory.(CompoundDomain)
		default:
			fmt.Print("directory type does not exist ", subDirType)
		}
	}

	return *pugDirectory
}

// Implementation for Directory object functions
func (directory *PubchemDirectory) ValidateDirectory() bool {
	return true
}

func (directory *PubchemDirectory) ConstructDirectory() string {
	var directories []string
	if directory.Dom != -1 {
		directories = append(directories, directory.Dom.String())
	}

	if directory.CD != -1 {
		directories = append(directories, directory.CD.String())
	}

	return strings.Join(directories, "/")
}
