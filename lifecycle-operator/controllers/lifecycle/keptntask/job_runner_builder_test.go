package keptntask

import (
	"testing"

	"github.com/keptn/lifecycle-toolkit/lifecycle-operator/apis/lifecycle/v1beta1"
	"github.com/stretchr/testify/require"
	v1 "k8s.io/api/core/v1"
)

func Test_getJobRunnerBuilder(t *testing.T) {
	runtimeBuilderOptions := BuilderOptions{
		funcSpec: &v1beta1.RuntimeSpec{
			Inline: v1beta1.Inline{
				Code: "some code",
			},
		},
	}
	containerBuilderOptions := BuilderOptions{
		containerSpec: &v1beta1.ContainerSpec{
			Container: &v1.Container{
				Image: "image",
			},
		},
	}
	tests := []struct {
		name    string
		options BuilderOptions
		want    JobRunnerBuilder
	}{
		{
			name:    "js builder",
			options: runtimeBuilderOptions,
			want:    NewRuntimeBuilder(runtimeBuilderOptions),
		},
		{
			name:    "container builder",
			options: containerBuilderOptions,
			want:    NewContainerBuilder(containerBuilderOptions),
		},
		{
			name:    "invalid builder",
			options: BuilderOptions{},
			want:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, NewJobRunnerBuilder(tt.options))
		})
	}
}
