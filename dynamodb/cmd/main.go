package main

import (
	"dynamodb/model"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gofrs/uuid"
	"github.com/guregu/dynamo"
	"github.com/k0kubun/pp"
)

func main() {
	creds := credentials.NewStaticCredentials("AWS_ACCESS_KEY_ID", "AWS_SECRET_ACCESS_KEY", "token")
	config := &aws.Config{
		Credentials: creds,
		Region:      aws.String("us-west-2"),
		Endpoint:    aws.String("http://localhost:8000"),
	}
	db := dynamo.New(session.New(), config)

	id, _ := uuid.NewV4()
	user := &model.User{
		ID:       id.String(),
		Username: "hieudeptrai",
		Fullname: "Hieu Phan",
		ImageID:  "456",
	}
	table := db.Table(user.TableName())
	err := table.Put(user).Run()
	if err != nil {
		panic(err)
	}

	results := []model.User{}
	table.Scan().Filter("'username' = ?", "hieudeptrai").All(&results)
	pp.Println(results)
}
