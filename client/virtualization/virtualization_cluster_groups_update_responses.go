package virtualization

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// VirtualizationClusterGroupsUpdateReader is a Reader for the VirtualizationClusterGroupsUpdate structure.
type VirtualizationClusterGroupsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClusterGroupsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationClusterGroupsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVirtualizationClusterGroupsUpdateOK creates a VirtualizationClusterGroupsUpdateOK with default headers values
func NewVirtualizationClusterGroupsUpdateOK() *VirtualizationClusterGroupsUpdateOK {
	return &VirtualizationClusterGroupsUpdateOK{}
}

/* VirtualizationClusterGroupsUpdateOK describes a response with status code 200, with default header values.

VirtualizationClusterGroupsUpdateOK virtualization cluster groups update o k
*/
type VirtualizationClusterGroupsUpdateOK struct {
	Payload *models.ClusterGroup
}

func (o *VirtualizationClusterGroupsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /virtualization/cluster-groups/{id}/][%d] virtualizationClusterGroupsUpdateOK  %+v", 200, o.Payload)
}
func (o *VirtualizationClusterGroupsUpdateOK) GetPayload() *models.ClusterGroup {
	return o.Payload
}

func (o *VirtualizationClusterGroupsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
