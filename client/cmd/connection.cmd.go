package cmd

import (
	"bytes"
	"crypto/x509"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"

	"github.com/manifoldco/promptui"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/tasnimzotder/tchat/client/internal"
	"github.com/tasnimzotder/tchat/client/models"
	"github.com/tasnimzotder/tchat/client/pkg/crypto"
	"github.com/tasnimzotder/tchat/client/pkg/file"
)

var connectionCmd = &cobra.Command{
	Use:   "conn",
	Short: "Connection related commands",
	Long:  `Connection related commands`,
	Run:   ConnectionCmd,
}

func ConnectionCmd(cmd *cobra.Command, args []string) {
	connectionTypes := []string{
		"begin - Begin connection for other users",        // 0
		"add - Add a user to the connection list",         // 1
		"remove - remove a user from the connection list", // 2
		"list - list all users in the connection list",    // 3
		"end - End connection",                            // 4
	}

	commandTypePrompt := promptui.Select{
		Label: "Select a Connection Type",
		Items: connectionTypes,
	}

	_, commandType, err := commandTypePrompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}

	switch commandType {
	case connectionTypes[0]: // begin
		BeginConnection(cmd)
	case connectionTypes[1]: // add
		AddConnection(cmd)
	case connectionTypes[3]: // list
		ListConnections(cmd)
	}
}

func BeginConnection(cmd *cobra.Command) {
	fmt.Printf("\nThis setting will store your public RSA key on the server for other users to download \nand use to send you encrypted messages.\n\n")

	//	todo: implement
	passKey := generatePassKey(6)
	expirationTime := "3600"

	persistenceTypes := []string{
		"single use - key deleted after use",
		"temporary - key deleted after time",
		"persistent - key stored on server",
	}

	persistenceTypePrompt := promptui.Select{
		Label: "Select a Persistence Type for the public key storage",
		Items: persistenceTypes,
	}

	_, persistenceType, err := persistenceTypePrompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}

	switch persistenceType {
	case persistenceTypes[0]: // single use
		expirationTime = "0"
	case persistenceTypes[1]: // temporary
		prompt := promptui.Prompt{
			Label: "Enter the time in seconds for the key to expire",
			Validate: func(input string) error {
				_, err := strconv.Atoi(input)
				return err
			},
		}

		_expirationTime, err := prompt.Run()
		if err != nil {
			log.Fatalf("Prompt failed %v\n", err)
		}

		expirationTime = _expirationTime

	case persistenceTypes[2]: // persistent
		expirationTime = "-1"
	}

	_displayExpirationTime := ""

	if persistenceType == persistenceTypes[1] {
		_displayExpirationTime = fmt.Sprintf("%ss", expirationTime)
	} else {
		_displayExpirationTime = "-"
	}

	config, err := file.ReadFromConfigFile()
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	//	display
	fmt.Println()
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"Persistence Type",
		"Expiration Time",
		"Passkey",
		"User ID",
	})
	table.SetBorder(false)

	table.Append([]string{
		persistenceType,
		_displayExpirationTime,
		passKey,
		config.ID,
	})
	fmt.Println()

	table.Render()

	// post the data to the server
	public_key, err := file.GetPublicRSAKey()
	if err != nil {
		log.Fatalf("Failed to get public RSA key: %v", err)
	}

	var encryptioner crypto.Encryptioner = &crypto.RSAEncryption{}

	// convert the public key to string
	publicKeyBytes := x509.MarshalPKCS1PublicKey(public_key)
	publicKeyBase64 := encryptioner.EncodeBase64(publicKeyBytes)

	connectionReq := models.Connection{
		Name:        config.Name,
		UserID:      config.ID,
		PublicKey:   publicKeyBase64,
		PassKey:     passKey,
		Expiration:  expirationTime,
		Persistence: persistenceType,
	}

	err = internal.StartConnection(connectionReq)

	if err != nil {
		log.Fatalf("Failed to start connection: %v", err)
	}

	// copy the passkey to clipboard
	// clipboardData := struct {
	// 	PassKey string
	// 	ID      string
	// }{
	// 	PassKey: passKey,
	// 	ID:      config.ID,
	// }

	// _data, err := json.Marshal(clipboardData)
	// if err != nil {
	// 	log.Fatalf("Failed to marshal clipboard data: %v", err)
	// }

	// util.CopyToClipboard(string(_data))
}

func ListConnections(cmd *cobra.Command) {
	contacts, err := file.ReadFromContactFile()
	if err != nil {
		log.Fatalf("Failed to read from contacts file: %v", err)
	}

	//	display
	fmt.Println()
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"S/N",
		"Name",
		"ID",
	})
	table.SetBorder(false)

	for i, contact := range contacts {
		table.Append([]string{
			strconv.Itoa(i + 1),
			contact.Name,
			contact.ID,
		})
	}

	fmt.Println()
	table.Render()
}

func AddConnection(cmd *cobra.Command) {
	var passKey string
	var userID string
	var name string

	// userID prompt
	userIDPrompt := promptui.Prompt{
		Label: "Enter the user ID",
		Validate: func(input string) error {
			if len(input) < 1 {
				return fmt.Errorf("user id cannot be empty")
			}
			return nil
		},
	}

	_userID, err := userIDPrompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}

	userID = _userID

	// name prompt
	namePrompt := promptui.Prompt{
		Label: "Enter the user name",
		Validate: func(input string) error {
			if len(input) < 1 {
				return fmt.Errorf("user name cannot be empty")
			}
			return nil
		},
	}

	_name, err := namePrompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}

	name = _name

	// passkey prompt
	passKeyPrompt := promptui.Prompt{
		Label: "Enter the passkey",
		Validate: func(input string) error {
			if len(input) < 6 {
				return fmt.Errorf("passkey cannot be empty")
			}
			return nil
		},
	}

	_passKey, err := passKeyPrompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}

	passKey = _passKey

	connectionRequest := internal.GetConnectionRequest{
		UserID:  userID,
		PassKey: passKey,
	}

	connection, err := internal.GetConnection(connectionRequest)
	if err != nil {
		log.Fatalf("Failed to get connection: %v", err)
	}

	keyPath, err := file.StoreContactPublicKey(connection.UserID, connection.PublicKey)
	if err != nil {
		log.Fatalf("Failed to store contact public key: %v", err)
	}

	contact := models.Contact{
		ID:   connection.UserID,
		Name: name,
		Key:  keyPath,
	}

	// save contact to file
	err = file.WriteToContactFile(contact)
	if err != nil {
		log.Fatalf("Failed to write to contact file: %v", err)
	}

	fmt.Printf("Connection added successfully\n")
}

func generatePassKey(length int) string {
	const digits = "0123456789"
	var passKey bytes.Buffer

	for i := 0; i < length; i++ {
		passKey.WriteByte(digits[rand.Intn(len(digits))])
	}

	return passKey.String()
}
