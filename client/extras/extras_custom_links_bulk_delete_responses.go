package extras

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExtrasCustomLinksBulkDeleteReader is a Reader for the ExtrasCustomLinksBulkDelete structure.
type ExtrasCustomLinksBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasCustomLinksBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewExtrasCustomLinksBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasCustomLinksBulkDeleteNoContent creates a ExtrasCustomLinksBulkDeleteNoContent with default headers values
func NewExtrasCustomLinksBulkDeleteNoContent() *ExtrasCustomLinksBulkDeleteNoContent {
	return &ExtrasCustomLinksBulkDeleteNoContent{}
}

/* ExtrasCustomLinksBulkDeleteNoContent describes a response with status code 204, with default header values.

ExtrasCustomLinksBulkDeleteNoContent extras custom links bulk delete no content
*/
type ExtrasCustomLinksBulkDeleteNoContent struct {
}

func (o *ExtrasCustomLinksBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /extras/custom-links/][%d] extrasCustomLinksBulkDeleteNoContent ", 204)
}

func (o *ExtrasCustomLinksBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
