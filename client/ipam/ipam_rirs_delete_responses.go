package ipam

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// IpamRirsDeleteReader is a Reader for the IpamRirsDelete structure.
type IpamRirsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamRirsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIpamRirsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamRirsDeleteNoContent creates a IpamRirsDeleteNoContent with default headers values
func NewIpamRirsDeleteNoContent() *IpamRirsDeleteNoContent {
	return &IpamRirsDeleteNoContent{}
}

/* IpamRirsDeleteNoContent describes a response with status code 204, with default header values.

IpamRirsDeleteNoContent ipam rirs delete no content
*/
type IpamRirsDeleteNoContent struct {
}

func (o *IpamRirsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/rirs/{id}/][%d] ipamRirsDeleteNoContent ", 204)
}

func (o *IpamRirsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
