package scheduler

import (
	"context"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/stretchr/testify/assert"
)

type AwsRoleProviderMock struct {
	AwsKey          string
	AwsSecret       string
	AwsSessionToken string
	AwsProviderName string
}

func (a *AwsRoleProviderMock) Retrieve() (aws.Credentials, error) {
	return aws.Credentials{
		AccessKeyID:     a.AwsKey,
		SecretAccessKey: a.AwsSecret,
		SessionToken:    a.AwsSessionToken,
	}, nil
}

func (a *AwsRoleProviderMock) ExpiresAt() time.Time {
	return time.Time{}
}

func (a *AwsRoleProviderMock) RetrieveWithContext(context.Context) (aws.Credentials, error) {
	return aws.Credentials{
		AccessKeyID:     a.AwsKey,
		SecretAccessKey: a.AwsSecret,
		SessionToken:    a.AwsSessionToken,
	}, nil
}

func (a *AwsRoleProviderMock) IsExpired() bool { return false }

// TODO: uncomment this test after figuring out how to create a mock compatible with WebIdentityRoleProvider
//func TestPopulationForAwsRoleToAssumeSetsValueOfKeys(t *testing.T) {
//	stateEnvVars := make(map[string]string)
//	commandEnvVars := make(map[string]string)
//
//	x := AwsRoleProviderMock{
//		AwsKey:          "statekey",
//		AwsSecret:       "stateSecret",
//		AwsSessionToken: "stateSessionToken",
//	}.(stscreds.WebIdentityRoleProvider)
//	job := Job{
//		StateEnvProvider: &x,
//		CommandEnvProvider: AwsRoleProviderMock{
//			AwsKey:          "commandkey",
//			AwsSecret:       "commandSecret",
//			AwsSessionToken: "commandSessionToken",
//		},
//
//		StateEnvVars:   stateEnvVars,
//		CommandEnvVars: commandEnvVars,
//	}
//
//	job.PopulateAwsCredentialsEnvVarsForJob()
//	assert.Equal(t, job.CommandEnvVars["AWS_ACCESS_KEY_ID"], "KEY")
//	assert.Equal(t, job.CommandEnvVars["AWS_SECRET_ACCESS_KEY"], "SECRET")
//	assert.Equal(t, job.CommandEnvVars["AWS_SESSION_TOKEN"], "TOKEN")
//	assert.Equal(t, job.StateEnvVars["AWS_ACCESS_KEY_ID"], "KEY")
//	assert.Equal(t, job.StateEnvVars["AWS_SECRET_ACCESS_KEY"], "SECRET")
//	assert.Equal(t, job.StateEnvVars["AWS_SESSION_TOKEN"], "TOKEN")
//}

func TestPopulationForNoAwsRoleToAssumeDoesNotSetValueOfKeys(t *testing.T) {
	stateEnvVars := make(map[string]string)
	commandEnvVars := make(map[string]string)

	job := Job{
		StateEnvVars:   stateEnvVars,
		CommandEnvVars: commandEnvVars,
	}

	job.PopulateAwsCredentialsEnvVarsForJob()
	assert.NotContains(t, job.CommandEnvVars, "AWS_ACCESS_KEY_ID")
	assert.NotContains(t, job.CommandEnvVars, "AWS_SECRET_ACCESS_KEY")
	assert.NotContains(t, job.CommandEnvVars, "AWS_SESSION_TOKEN")
	assert.NotContains(t, job.StateEnvVars, "AWS_ACCESS_KEY_ID")
	assert.NotContains(t, job.StateEnvVars, "AWS_SECRET_ACCESS_KEY")
	assert.NotContains(t, job.StateEnvVars, "AWS_SESSION_TOKEN")
}
