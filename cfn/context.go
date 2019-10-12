package cfn

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

type Context struct {
	cfn *cloudformation.CloudFormation
}

func NewContext() *Context {
	sess := session.Must(
		session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}))
	context := &Context{}
	context.cfn = cloudformation.New(sess)
	return context
}
