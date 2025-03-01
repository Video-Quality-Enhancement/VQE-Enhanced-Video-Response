package config

import (
	"context"
	"os"
	"time"

	firebase "firebase.google.com/go/v4"
	"golang.org/x/exp/slog"
	"google.golang.org/api/option"
)

type FirebaseClient interface {
	GetEmail(uid string) (string, error)
}

type firebaseClient struct {
	app *firebase.App
}

func NewFirebaseClient() FirebaseClient {

	opt := option.WithCredentialsFile(os.Getenv("FIREBASE_SA_KEY_PATH"))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		slog.Error("error initializing app: %v\n", err)
	}

	return &firebaseClient{app}
}

func (c *firebaseClient) GetEmail(uid string) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := c.app.Auth(ctx)
	if err != nil {
		slog.Error("error getting Auth client", "error", err)
		return "", err
	}

	user, err := client.GetUser(ctx, uid)
	if err != nil {
		slog.Error("error getting user from firebase client", "error", err)
		return "", err
	}

	return user.Email, nil

}
