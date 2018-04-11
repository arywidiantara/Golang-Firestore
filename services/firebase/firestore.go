package firebase

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"golang-firestore/services/notification"

	firebase "firebase.google.com/go"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

/**
 * @brief      Sends a data to firestore.
 *
 * @param      subdomain     The subdomain
 * @param      id_customer   The identifier customer
 * @param      dataCustomer  The data customer
 *
 * @return     object
 */
func SendDataToFirestore(subdomain string, id_customer int, dataCustomer map[string]interface{}) error {
	fmt.Println(dataCustomer)
	directory, err := os.Getwd()
	if err != nil {
		log.Println(err)
		fmt.Println(err)
		notification.SendNotificationError(err)
		return err
	}

	ctx := context.Background()
	sa := option.WithCredentialsFile(directory + "/env.json")

	// Use a service account
	fmt.Println(sa)
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Println(err)
		fmt.Println(err)
		notification.SendNotificationError(err)
		return err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Println(err)
		fmt.Println(err)
		notification.SendNotificationError(err)
		return err
	}

	_, err = client.Collection(subdomain).Doc("global").Collection("customers").Doc(strconv.Itoa(id_customer)).Set(ctx, dataCustomer)
	if err != nil {
		log.Println(err)
		fmt.Println(err)
		notification.SendNotificationError(err)
		return err
	}

	return nil
}
