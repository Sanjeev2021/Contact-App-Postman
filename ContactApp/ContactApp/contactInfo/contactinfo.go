package contactinfo

//CRUD OPERATIONS IN CONTACTINFO

import (
	"errors"

	uuid "github.com/satori/go.uuid"

)

// type of contact
type ContactInfo struct {
	ID               uuid.UUID
	contactInfoType  string
	contactInfoValue string
}

func NewContactInfo(contactinfotype, contactinfovalue string) *ContactInfo {
	return &ContactInfo{
		ID:               uuid.NewV4(),
		contactInfoType:  contactinfotype,
		contactInfoValue: contactinfovalue,
	}
}

//func ReadNewContact() ([]ContactInfo , error) {

//}



// func (c *ContactInfo) CreateContactInfo(contactinfotype string) (*ContactInfo, error) {
// 	if !c.IsAdmin {
// 		return nil, errors.New(c.FirstName + "not admin")
// 	}
// 	_, contactExist := FindContactInfo(c.Mycontacts, contactinfotype)
// 	if !contactExist {
// 		return nil, errors.New("Contact already exists")
// 	}
	
// 	c.Mycontacts = append(c.Mycontacts, newContactInfo)
// 	return newContactInfo, nil
// }


func FindContactInfo(list []ContactInfo, ID uuid.UUID) (*ContactInfo, error) {
	for i := 0; i < len(list); i++ {
		if list[i].ID == ID {
			return &list[i], nil
		}
	}
	return nil, errors.New("No contact info found")
}

func (c *ContactInfo)UpdateContact(list []ContactInfo, ID uuid.UUID, field string, value string) (*ContactInfo, error) {
	contact, err := FindContactInfo(list, ID)
	if err != nil {
		return nil, err
	}
	switch field {
	case "contactInfoType":
		contact.contactInfoType = value
		return contact, nil
	case "contactInfoValue":
		contact.contactInfoValue = value
		return contact, nil
	default:
		return nil, errors.New("Invalid field")
	}
}

func (c *ContactInfo) DeleteContact(list []ContactInfo, ID uuid.UUID) ([]ContactInfo, error) {
	for i := 0; i < len(list); i++ {
		if list[i].ID == ID {
			return append(list[:i], list[i+1:]...), nil
		}

	}
	return nil, errors.New("USER does not exist")
}
