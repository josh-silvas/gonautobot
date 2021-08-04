package users

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// UsersPermissionsCreateReader is a Reader for the UsersPermissionsCreate structure.
type UsersPermissionsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersPermissionsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUsersPermissionsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUsersPermissionsCreateCreated creates a UsersPermissionsCreateCreated with default headers values
func NewUsersPermissionsCreateCreated() *UsersPermissionsCreateCreated {
	return &UsersPermissionsCreateCreated{}
}

/* UsersPermissionsCreateCreated describes a response with status code 201, with default header values.

UsersPermissionsCreateCreated users permissions create created
*/
type UsersPermissionsCreateCreated struct {
	Payload *models.ObjectPermission
}

func (o *UsersPermissionsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /users/permissions/][%d] usersPermissionsCreateCreated  %+v", 201, o.Payload)
}
func (o *UsersPermissionsCreateCreated) GetPayload() *models.ObjectPermission {
	return o.Payload
}

func (o *UsersPermissionsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ObjectPermission)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
