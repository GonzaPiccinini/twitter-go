package secretsmanager

import (
	"encoding/json"
	"fmt"

	"github.com/GonzaPiccinini/twitter-go/awsgo"
	"github.com/GonzaPiccinini/twitter-go/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func GetSecret(secretName string) (models.Secret, error) {
	var secrets models.Secret
	fmt.Println("-> Getting secrets: " + secretName)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	key, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})
	if err != nil {
		fmt.Println("Error getting secrets: " + err.Error())
		return secrets, err
	}

	if err = json.Unmarshal([]byte(*key.SecretString), &secrets); err != nil {
		fmt.Println("Error unmarshaling secrets: " + err.Error())
		return secrets, err
	}

	fmt.Println("-> Successful secret reading " + secretName)
	return secrets, nil
}
