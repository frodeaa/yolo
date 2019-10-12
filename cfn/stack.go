package cfn

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"sort"
	"time"
)

var filter = []*string{
	aws.String("CREATE_IN_PROGRESS"),
	aws.String("CREATE_FAILED"),
	aws.String("CREATE_COMPLETE"),
	aws.String("ROLLBACK_IN_PROGRESS"),
	aws.String("ROLLBACK_FAILED"),
	aws.String("ROLLBACK_COMPLETE"),
	aws.String("DELETE_IN_PROGRESS"),
	aws.String("DELETE_FAILED"),
	aws.String("UPDATE_IN_PROGRESS"),
	aws.String("UPDATE_COMPLETE_CLEANUP_IN_PROGRESS"),
	aws.String("UPDATE_COMPLETE"),
	aws.String("UPDATE_ROLLBACK_IN_PROGRESS"),
	aws.String("UPDATE_ROLLBACK_FAILED"),
	aws.String("UPDATE_ROLLBACK_COMPLETE_CLEANUP_IN_PROGRESS"),
	aws.String("UPDATE_ROLLBACK_COMPLETE"),
	aws.String("REVIEW_IN_PROGRESS"),
}

type StackSummary struct {
	cloudformation.StackSummary
	LastChangedTime *time.Time
}

func lastChangedTime(s *cloudformation.StackSummary) *time.Time {
	if s.LastUpdatedTime != nil {
		return s.LastUpdatedTime
	}
	return s.CreationTime
}

func (c *Context) ListStacks() ([]*StackSummary, error) {
	input := &cloudformation.ListStacksInput{
		StackStatusFilter: filter,
	}

	resp, err := c.cfn.ListStacks(input)

	if err != nil {
		return nil, err
	}

	res := make([]*StackSummary, 0, len(resp.StackSummaries))
	for _, s := range resp.StackSummaries {
		res = append(res, &StackSummary{*s, lastChangedTime(s)})
	}

	sort.Slice(res, func(i, j int) bool {
		// old to new
		return res[i].LastChangedTime.Before(*res[j].LastChangedTime)
	})

	return res, nil
}
