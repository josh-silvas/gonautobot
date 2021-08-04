package virtualization

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// VirtualizationVirtualMachinesPartialUpdateReader is a Reader for the VirtualizationVirtualMachinesPartialUpdate structure.
type VirtualizationVirtualMachinesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationVirtualMachinesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationVirtualMachinesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVirtualizationVirtualMachinesPartialUpdateOK creates a VirtualizationVirtualMachinesPartialUpdateOK with default headers values
func NewVirtualizationVirtualMachinesPartialUpdateOK() *VirtualizationVirtualMachinesPartialUpdateOK {
	return &VirtualizationVirtualMachinesPartialUpdateOK{}
}

/* VirtualizationVirtualMachinesPartialUpdateOK describes a response with status code 200, with default header values.

VirtualizationVirtualMachinesPartialUpdateOK virtualization virtual machines partial update o k
*/
type VirtualizationVirtualMachinesPartialUpdateOK struct {
	Payload *models.VirtualMachineWithConfigContext
}

func (o *VirtualizationVirtualMachinesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /virtualization/virtual-machines/{id}/][%d] virtualizationVirtualMachinesPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *VirtualizationVirtualMachinesPartialUpdateOK) GetPayload() *models.VirtualMachineWithConfigContext {
	return o.Payload
}

func (o *VirtualizationVirtualMachinesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualMachineWithConfigContext)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
