package pubchem

// BaseURL for PUG REST API
const BaseURL = "https://pubchem.ncbi.nlm.nih.gov/rest/pug"

// TestURL for PUG REST API
const TestURL = "https://pubchem.ncbi.nlm.nih.gov/rest/pug/compound/cid/1,2,3,4,5/property/MolecularFormula,MolecularWeight,CanonicalSMILES/JSON"

// GetCompoundByName retrieves compound from DB by name or chemical formula
const GetCompoundByName = "compound/name"

// GetCompoundByCID retrieves compound from DB by cid
const GetCompoundByCID = "compound/cid"
