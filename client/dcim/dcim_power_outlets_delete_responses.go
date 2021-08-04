package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimPowerOutletsDeleteReader is a Reader for the DcimPowerOutletsDelete structure.
type DcimPowerOutletsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerOutletsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimPowerOutletsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimPowerOutletsDeleteNoContent creates a DcimPowerOutletsDeleteNoContent with default headers values
func NewDcimPowerOutletsDeleteNoContent() *DcimPowerOutletsDeleteNoContent {
	return &DcimPowerOutletsDeleteNoContent{}
}

/* DcimPowerOutletsDeleteNoContent describes a response with status code 204, with default header values.

DcimPowerOutletsDeleteNoContent dcim power outlets delete no content
*/
type DcimPowerOutletsDeleteNoContent struct {
}

func (o *DcimPowerOutletsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/power-outlets/{id}/][%d] dcimPowerOutletsDeleteNoContent ", 204)
}

func (o *DcimPowerOutletsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
