package pubchem

import (
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
