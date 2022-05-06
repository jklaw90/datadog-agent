// +build linux darwin

package filesystem

import (
	"context"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/go-multierror"
)

func getFileSystemInfo() (interface{}, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	/* Grab filesystem data from df	*/
	cmd := exec.CommandContext(ctx, "df", dfOptions...)

	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("df failed to collect filesystem data: %s", err)
	}
	if out != nil {
		return parseDfOutput(string(out))
	}
	return nil, fmt.Errorf("df failed to collect filesystem data")
}

func parseDfOutput(out string) (interface{}, error) {
	var aggregateError error
	lines := strings.Split(out, "\n")
	var fileSystemInfo = make([]interface{}, len(lines)-2)
	for i, line := range lines[1:] {
		values := regexp.MustCompile(`\s+`).Split(line, -1)
		if len(values) >= expectedLength {
			var err error
			fileSystemInfo[i], err = updatefileSystemInfo(values)
			if err != nil {
				aggregateError = multierror.Append(aggregateError, err)
			}
		}
	}
	return fileSystemInfo, aggregateError
}