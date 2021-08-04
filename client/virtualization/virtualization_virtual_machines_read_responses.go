package virtualization

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// VirtualizationVirtualMachinesReadReader is a Reader for the VirtualizationVirtualMachinesRead structure.
type VirtualizationVirtualMachinesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationVirtualMachinesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationVirtualMachinesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVirtualizationVirtualMachinesReadOK creates a VirtualizationVirtualMachinesReadOK with default headers values
func NewVirtualizationVirtualMachinesReadOK() *VirtualizationVirtualMachinesReadOK {
	return &VirtualizationVirtualMachinesReadOK{}
}

/* VirtualizationVirtualMachinesReadOK describes a response with status code 200, with default header values.

VirtualizationVirtualMachinesReadOK virtualization virtual machines read o k
*/
type VirtualizationVirtualMachinesReadOK struct {
	Payload *models.VirtualMachineWithConfigContext
}

func (o *VirtualizationVirtualMachinesReadOK) Error() string {
	return fmt.Sprintf("[GET /virtualization/virtual-machines/{id}/][%d] virtualizationVirtualMachinesReadOK  %+v", 200, o.Payload)
}
func (o *VirtualizationVirtualMachinesReadOK) GetPayload() *models.VirtualMachineWithConfigContext {
	return o.Payload
}

func (o *VirtualizationVirtualMachinesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualMachineWithConfigContext)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
