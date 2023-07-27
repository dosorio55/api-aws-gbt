package secretmanager

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/dosorio55/gambitApi/awsgo"
	"github.com/dosorio55/gambitApi/models"
)

// GetSecret retrieves the secret from AWS Secrets Manager
func GetSecret(SecretName string) (models.SecretStructJSON, error) {
	var SecretStructJSON models.SecretStructJSON

	fmt.Println("SecretName: ", SecretName)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)

	password, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(SecretName),
	})

	if err != nil {
		return SecretStructJSON, err
	} 

	err = json.Unmarshal([]byte(*password.SecretString), &SecretStructJSON)

	if err != nil {
		fmt.Println("Error unmarshalling secret string: ", err)
		return SecretStructJSON, err
	}

	return SecretStructJSON, nil


}
