package pubchem_dirs

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

var fastSearches = [5]string{
  "fastidentity", "fastsimilarity_2d", "fastsimilarity_3d", "fastsubstructure", "fastsuperstructure",
}

func (fs FastSearch) String() string {
  return fastSearches[fs]
}