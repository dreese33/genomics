package pubchem

// BaseURL for PUG REST API
const BaseURL = "https://pubchem.ncbi.nlm.nih.gov/rest/pug"

// TestURL for PUG REST API
const TestURL = "https://pubchem.ncbi.nlm.nih.gov/rest/pug/compound/cid/1,2,3,4,5/property/MolecularFormula,MolecularWeight,CanonicalSMILES/JSON"


// Domain specifies the data we are searching for
type Domain int32

// Domain enumeration
const (
  Substance Domain = iota
  Compound
  Assay
)

// Domains stringified for Domain
var Domains = [3]string{"substance", "compound", "assay"}
func (d Domain) String() string {
  return Domains[d]
}


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

// CompoundDomains stringified CompoundDomain
var CompoundDomains = [8]string{
  "cid", "name", "smiles", "inchi", "sdf", "inchikey", "formula", "listkey",
}

func (cd CompoundDomain) String() string {
  return CompoundDomains[cd]
}


// StructureSearch search by chemical structure
type StructureSearch int32

// StructureSearch enumeration
const (
  Substructure StructureSearch = iota
  Superstructure
  Similarity
  Identity
)

// StructureSearches stringified StructureSearch
var StructureSearches = [4]string{"substructure", "superstructure", "similarity", "identity"}
func (ss StructureSearch) String() string {
  return StructureSearches[ss]
}


// FastSearch info retrieval
type FastSearch int32

// FastSearch enumeration
const (
  FastIdentity FastSearch = iota
  FastSimilarity2D
  FastSimilarity3D
  FastSubstructure
  FastSuperstructure
)

// FastSearches stringified FastSearch
var FastSearches = [5]string{
  "fastidentity", "fastsimilarity_2d", "fastsimilarity_3d", "fastsubstructure", "fastsuperstructure",
}

func (fs FastSearch) String() string {
  return FastSearches[fs]
}


// GetCompoundByName retrieves compound from DB by name or chemical formula
const GetCompoundByName = "compound/name"

// GetCompoundByCID retrieves compound from DB by cid
const GetCompoundByCID = "compound/cid"