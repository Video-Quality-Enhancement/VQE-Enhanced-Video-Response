package gapi

import (
	"context"
	"os"

	"cloud.google.com/go/storage"
	"golang.org/x/exp/slog"
	"google.golang.org/api/option"
)

type GoogleCloudStorage interface {
	GrantAccess(filepath, email string) error
}

type googleCloudStorage struct {
	storageClient *storage.Client
	bucket        *storage.BucketHandle
}

func NewGoogleCloudStorage() GoogleCloudStorage {
	ctx := context.Background()

	credsPath := os.Getenv("STORAGE_SA_KEY_PATH")
	// projectId := os.Getenv("PROJECT_ID")
	bucketName := os.Getenv("BUCKET_NAME")

	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credsPath))
	if err != nil {
		slog.Error("Error creating storage client", "error", err)
		panic(err)
	}

	bucket := client.Bucket(bucketName)

	return &googleCloudStorage{client, bucket}
}

func (gcs *googleCloudStorage) GrantAccess(filepath, email string) error {

	object := gcs.bucket.Object(filepath)

	acl := object.ACL()
	ctx := context.Background()
	entity := storage.ACLEntity("user-" + email)

	err := acl.Set(ctx, entity, storage.RoleReader)
	if err != nil {
		slog.Error("Error granting access to file", "filepath", filepath, "error", err)
		return err
	}

	slog.Debug("Granted access to file", "filepath", filepath)

	return nil

}
