package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

// Contact struct to hold contact information
type Contact struct {
    Name    string
    Phone   string
    Email   string
}

// AddressBook to hold multiple contacts
type AddressBook struct {
    Contacts []Contact
}

// Add a new contact to the address book
func (ab *AddressBook) AddContact(name, phone, email string) {
    newContact := Contact{Name: name, Phone: phone, Email: email}
    ab.Contacts = append(ab.Contacts, newContact)
    fmt.Println("Contact added successfully.")
}

// List all contacts in the address book
func (ab *AddressBook) ListContacts() {
    if len(ab.Contacts) == 0 {
        fmt.Println("No contacts found.")
        return
    }
    fmt.Println("Contacts:")
    for i, contact := range ab.Contacts {
        fmt.Printf("%d. %s - %s - %s\n", i+1, contact.Name, contact.Phone, contact.Email)
    }
}

// Delete a contact by index
func (ab *AddressBook) DeleteContact(index int) {
    if index < 0 || index >= len(ab.Contacts) {
        fmt.Println("Invalid index.")
        return
    }
    ab.Contacts = append(ab.Contacts[:index], ab.Contacts[index+1:]...)
    fmt.Println("Contact deleted successfully.")
}

// Get user input with a prompt
func getInput(prompt string) string {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print(prompt)
    input, _ := reader.ReadString('\n')
    return strings.TrimSpace(input)
}

func main() {
    addressBook := AddressBook{}

    for {
        fmt.Println("\nAddress Book Menu:")
        fmt.Println("1. Add Contact")
        fmt.Println("2. List Contacts")
        fmt.Println("3. Delete Contact")
        fmt.Println("4. Exit")
        
        choice := getInput("Choose an option (1-4): ")

        switch choice {
        case "1":
            name := getInput("Enter name: ")
            phone := getInput("Enter phone: ")
            email := getInput("Enter email: ")
            addressBook.AddContact(name, phone, email)

        case "2":
            addressBook.ListContacts()

        case "3":
            indexStr := getInput("Enter the contact number to delete: ")
            index := 0
            fmt.Sscanf(indexStr, "%d", &index)
            addressBook.DeleteContact(index - 1)

        case "4":
            fmt.Println("Exiting...")
            return

        default:
            fmt.Println("Invalid choice. Please select again.")
        }
    }
}
