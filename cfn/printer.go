package cfn

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/gookit/color"
	"io"
	"strings"
)

type StackPrinter struct {
	timestampFormat string
	writer          io.Writer
}

func NewStackPrinter(monocromeOutput bool, writer io.Writer) StackPrinter {
	if monocromeOutput {
		color.Disable()
	}
	return StackPrinter{
		timestampFormat: "Mon Jan 02 2006 15:04:05 MST",
		writer:          writer,
	}
}

func statusColor(s *StackSummary) color.Color {
	c := color.FgGreen
	if strings.HasSuffix(aws.StringValue(s.StackStatus), "_FAILED") {
		c = color.FgRed
	}
	return c
}

func (sp *StackPrinter) printStack(s *StackSummary) {
	scolor := statusColor(s)

	line := fmt.Sprintf("%s %s   %s",
		s.LastChangedTime.Format(sp.timestampFormat),
		scolor.Sprintf("%25.25s", aws.StringValue(s.StackStatus)),
		aws.StringValue(s.StackName),
	)
	fmt.Fprintln(sp.writer, line)
}

func (sp *StackPrinter) PrintStacks(ss []*StackSummary) {
	for _, s := range ss {
		sp.printStack(s)
	}
}
