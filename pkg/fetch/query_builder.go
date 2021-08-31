package fetch

// ResponseFormat - for HTTP requests
type ResponseFormat int32

// Response format enumeration
const (
  JSON ResponseFormat = iota
  CSV ResponseFormat = iota
  XML ResponseFormat = iota
)

// Responses - for ResponseFormat
var Responses = [3]string{"JSON", "CSV", "XML"}

// Endpoint - defines all functions that
// need to be implemented by each query builder
// for each api used in this project
type Endpoint interface {
  GetURL() string
  GetFmt() ResponseFormat
  TestEndpoint(string) bool
}

// EndpointImpl - information for url construction
type EndpointImpl struct {
  URL string
  Fmt ResponseFormat
  TestURL string
}

// NewEndpoint - implementation for Endpoint
func NewEndpoint(url string, fmt ResponseFormat, testURL string) Endpoint {
  var endpt = new(EndpointImpl)
  endpt.URL = url
  endpt.Fmt = fmt
  endpt.TestURL = testURL
  return endpt
}

// GetURL - implementation for Endpoint
func (endpoint *EndpointImpl) GetURL() string {
  return endpoint.URL
}

// GetFmt - implementation for Endpoint
func (endpoint *EndpointImpl) GetFmt() ResponseFormat {
  return endpoint.Fmt
} 

// TestEndpoint - implementation for Endpoint
func (endpoint *EndpointImpl) TestEndpoint(expected string) bool {
  body, err := RequestURL(endpoint.TestURL)
  if err != nil {
    return false
  }

  return body == expected
}
