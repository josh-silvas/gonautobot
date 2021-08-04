package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimConsolePortTemplatesDeleteReader is a Reader for the DcimConsolePortTemplatesDelete structure.
type DcimConsolePortTemplatesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsolePortTemplatesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimConsolePortTemplatesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimConsolePortTemplatesDeleteNoContent creates a DcimConsolePortTemplatesDeleteNoContent with default headers values
func NewDcimConsolePortTemplatesDeleteNoContent() *DcimConsolePortTemplatesDeleteNoContent {
	return &DcimConsolePortTemplatesDeleteNoContent{}
}

/* DcimConsolePortTemplatesDeleteNoContent describes a response with status code 204, with default header values.

DcimConsolePortTemplatesDeleteNoContent dcim console port templates delete no content
*/
type DcimConsolePortTemplatesDeleteNoContent struct {
}

func (o *DcimConsolePortTemplatesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/console-port-templates/{id}/][%d] dcimConsolePortTemplatesDeleteNoContent ", 204)
}

func (o *DcimConsolePortTemplatesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
