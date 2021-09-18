package fetch

import "strings"

// ResponseFormat for HTTP requests
type ResponseFormat int32

// Response format enumeration
const (
  JSON ResponseFormat = iota
  CSV
  XML
)

// Responses stringified for ResponseFormat
var Responses = [3]string{"JSON", "CSV", "XML"}
func (rf ResponseFormat) String() string {
  return Responses[rf]
}

// Endpoint defines all functions that need to be implemented by each query builder
// for each api used in this project
type Endpoint interface {
  GetURL() string
  GetFmt() ResponseFormat
  GetDIR() string
  GetSearchVal() string
  GetTestURL() string

  ConstructURL() string
}

// EndpointImpl information for url construction
type EndpointImpl struct {
  URL string
  Fmt ResponseFormat
  TestURL string
  DIR string
  SearchVal string
}

// FormatURLString ensures that the string does not end in a backslash
func FormatURLString(str string) string {
  return strings.TrimSuffix(str, "/")
}

// NewEndpoint constructor for the Endpoint object
func NewEndpoint(url string, fmt ResponseFormat, testURL string, dir string, searchVal string) Endpoint {
  var endpt = new(EndpointImpl)
  endpt.URL = FormatURLString(url)
  endpt.Fmt = fmt
  endpt.TestURL = FormatURLString(testURL)
  endpt.DIR = FormatURLString(dir)
  endpt.SearchVal = FormatURLString(searchVal)
  return endpt
}

// GetURL returns endpoint url
func (endpoint *EndpointImpl) GetURL() string {
  return endpoint.URL
}

// GetFmt returns endpoint response format
func (endpoint *EndpointImpl) GetFmt() ResponseFormat {
  return endpoint.Fmt
}

// GetDIR returns url directory to be queried
func (endpoint *EndpointImpl) GetDIR() string {
  return endpoint.DIR
}

// GetSearchVal returns value the client is searching for
func (endpoint *EndpointImpl) GetSearchVal() string {
  return endpoint.SearchVal
}

// GetTestURL returns specific url as a test endpoint
func (endpoint *EndpointImpl) GetTestURL() string {
  return endpoint.TestURL
}

// ConstructURL Constructs url then searches Pubchem api
// for the information specified
func (endpoint *EndpointImpl) ConstructURL() string {
  urlVals := []string{
    endpoint.GetURL(),
    endpoint.GetDIR(),
    endpoint.GetSearchVal(),
    endpoint.GetFmt().String(),
  }

  return strings.Join(urlVals, "/")
}
