package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimPowerFeedsDeleteReader is a Reader for the DcimPowerFeedsDelete structure.
type DcimPowerFeedsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerFeedsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimPowerFeedsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimPowerFeedsDeleteNoContent creates a DcimPowerFeedsDeleteNoContent with default headers values
func NewDcimPowerFeedsDeleteNoContent() *DcimPowerFeedsDeleteNoContent {
	return &DcimPowerFeedsDeleteNoContent{}
}

/* DcimPowerFeedsDeleteNoContent describes a response with status code 204, with default header values.

DcimPowerFeedsDeleteNoContent dcim power feeds delete no content
*/
type DcimPowerFeedsDeleteNoContent struct {
}

func (o *DcimPowerFeedsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/power-feeds/{id}/][%d] dcimPowerFeedsDeleteNoContent ", 204)
}

func (o *DcimPowerFeedsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
