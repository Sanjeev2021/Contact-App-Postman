package contact

//crud for contact
import (
	"errors"

	uuid "github.com/satori/go.uuid"

	//"contactApp/contact"
	contactinfo "contactApp/contactinfo"
)

type Contact struct {
	ID            uuid.UUID
	ContactName   string
	MyContactInfo []contactinfo.ContactInfo
}

// func NewContact(contactname string,mycontactinfo []contactinfo.ContactInfo) *Contact {
func NewContact(contactname string) *Contact {
	return &Contact{
		ID:          uuid.NewV4(),
		ContactName: contactname,
		//MyContactInfo: mycontactinfo,
	}
}

func FindContactID(contacts []*Contact, ID uuid.UUID) (*Contact, error) {
	for i := 0; i < len(contacts); i++ {
		if contacts[i].ID == ID {
			return contacts[i], nil
		}
	}
	return nil, errors.New("ID DOESNT EXIST")
	// doubt over there i am using error but down i am doing errors

}

func (c *Contact) UpdateContactName(name string) {
	c.ContactName = name
}

func DeleteContact(contacts []*Contact, contact *Contact) ([]*Contact, error) {
	for i := 0; i < len(contacts); i++ {
		if contacts[i] == contact {
			return append(contacts[:i], contacts[i+1:]...), nil
		}

	}
	return nil, errors.New("no contact found")
}

func (c *Contact) CreateContactInfo(cit, civ string) (*Contact, error) {
	newContactInfo := contactinfo.NewContactInfo(cit, civ)
	c.MyContactInfo = append(c.MyContactInfo, *newContactInfo)
	return nil, errors.New("Name already exist")
}
