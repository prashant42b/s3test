package s3

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var sess *session.Session
var err error

const (
	region     = "us-east-1"     // Specify your AWS region
	bucketName = "testbucketjmz" // Specify your bucket name
)

func InitAWSSession() {

	// Initialize a session using the default credential provider chain
	sess, err = session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		fmt.Println("Failed to create session:", err)
		return
	}
	fmt.Println("AWS Connection Established")

}

var BucketLS []string

func FetchS3BucketContent() {

	// Create an S3 service client
	s3Svc := s3.New(sess)

	// List objects in the bucket
	listObjectsResp, err := s3Svc.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		fmt.Println("Failed to list objects:", err)
		return
	}

	for _, obj := range listObjectsResp.Contents {
		BucketLS = append(BucketLS, *obj.Key)
	}
}
