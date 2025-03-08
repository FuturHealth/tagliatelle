package tagliatelle

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	exampleYAML = `resources:
  - ../base

images:
  - name: us-docker.pkg.dev/futurhealth/cloud-run/posthog-relay
    newTag: 1.5.0
`
	newTagPattern = `(?mi)(newTag: ["']?)(\d+\.\d+\.\d+)(["']?)$`
)

func TestCheckTagAlreadyExists(t *testing.T) {
	tt := []struct {
		name    string
		data    string
		pattern string
		tag     string
		want    bool
		err     error
	}{
		{
			name:    "tag already exists",
			data:    exampleYAML,
			pattern: newTagPattern,
			tag:     "1.5.0",
			want:    true,
			err:     nil,
		},
		{
			name:    "tag does not exist",
			data:    exampleYAML,
			pattern: newTagPattern,
			tag:     "1.5.1",
			want:    false,
		},
		{
			name:    "invalid regex",
			data:    exampleYAML,
			pattern: "invalid",
			tag:     "1.5.1",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			exists, err := checkTagAlreadyExists(&tc.data, newTagPattern, tc.tag)
			assert.Equal(t, tc.want, exists)
			assert.Equal(t, tc.err, err)
		})
	}
}
