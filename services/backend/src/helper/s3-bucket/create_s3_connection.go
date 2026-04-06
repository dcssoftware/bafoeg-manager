package s3bucket

import (
	"context"

	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func CreateS3BucketConnection() (*minio.Client, error) {
	endpoint := configuration.S3Bucket.Endpoint
	accessKeyID := configuration.S3Bucket.AccessKeyID
	secretAccessKey := configuration.S3Bucket.AccessKeySecret
	useSSL := configuration.S3Bucket.UseSSL

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, err
	}

	err = createS3BucketIfNotExists(
		context.Background(),
		minioClient,
		configuration.S3StoragePaths.ProfilePictureBucket,
		"",
	)

	if err != nil {
		return nil, err
	}

	err = createS3BucketIfNotExists(
		context.Background(),
		minioClient,
		configuration.S3StoragePaths.ApplicationDataBucket,
		"",
	)

	if err != nil {
		return nil, err
	}

	return minioClient, nil
}

func createS3BucketIfNotExists(ctx context.Context, minioClient *minio.Client, bucketName, location string) error {

	err := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {

		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			return nil
		} else {
			return err
		}

	}

	return nil
}
