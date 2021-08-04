package users

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// UsersPermissionsBulkPartialUpdateReader is a Reader for the UsersPermissionsBulkPartialUpdate structure.
type UsersPermissionsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersPermissionsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersPermissionsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUsersPermissionsBulkPartialUpdateOK creates a UsersPermissionsBulkPartialUpdateOK with default headers values
func NewUsersPermissionsBulkPartialUpdateOK() *UsersPermissionsBulkPartialUpdateOK {
	return &UsersPermissionsBulkPartialUpdateOK{}
}

/* UsersPermissionsBulkPartialUpdateOK describes a response with status code 200, with default header values.

UsersPermissionsBulkPartialUpdateOK users permissions bulk partial update o k
*/
type UsersPermissionsBulkPartialUpdateOK struct {
	Payload *models.ObjectPermission
}

func (o *UsersPermissionsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /users/permissions/][%d] usersPermissionsBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *UsersPermissionsBulkPartialUpdateOK) GetPayload() *models.ObjectPermission {
	return o.Payload
}

func (o *UsersPermissionsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ObjectPermission)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
