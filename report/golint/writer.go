package golint

import (
	"fmt"
	"github.com/securego/gosec/v2/report/core"
	"io"
	"strings"
)

//WriteReport write a report in golint format to the output writer
func WriteReport(w io.Writer, data *core.ReportInfo) error {
	// Output Sample:
	// /tmp/main.go:11:14: [CWE-310] RSA keys should be at least 2048 bits (Rule:G403, Severity:MEDIUM, Confidence:HIGH)

	for _, issue := range data.Issues {
		what := issue.What
		if issue.Cwe.ID != "" {
			what = fmt.Sprintf("[CWE-%s] %s", issue.Cwe.ID, issue.What)
		}

		// issue.Line uses "start-end" format for multiple line detection.
		lines := strings.Split(issue.Line, "-")
		start := lines[0]

		_, err := fmt.Fprintf(w, "%s:%s:%s: %s (Rule:%s, Severity:%s, Confidence:%s)\n",
			issue.File,
			start,
			issue.Col,
			what,
			issue.RuleID,
			issue.Severity.String(),
			issue.Confidence.String(),
		)
		if err != nil {
			return err
		}
	}
	return nil
}
