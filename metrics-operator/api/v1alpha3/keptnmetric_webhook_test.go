package v1alpha3

import (
	"reflect"
	"testing"

	"github.com/keptn/lifecycle-toolkit/metrics-operator/api/common"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

func TestKeptnMetric_validateRangeInterval(t *testing.T) {

	tests := []struct {
		name string
		Spec KeptnMetricSpec
		want *field.Error
	}{
		{
			name: "bad-provider",
			Spec: KeptnMetricSpec{
				Range: RangeSpec{Interval: "5mins"},
			},
			want: field.Invalid(
				field.NewPath("spec").Child("range").Child("interval"),
				"5mins",
				common.ErrForbiddenIntervalError.Error(),
			),
		},

		{
			name: "good-provider",
			Spec: KeptnMetricSpec{
				Range: RangeSpec{Interval: "5m"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &KeptnMetric{
				ObjectMeta: metav1.ObjectMeta{Name: tt.name},
				Spec:       KeptnMetricSpec{Range: RangeSpec{Interval: tt.Spec.Range.Interval}},
			}
			if got := s.validateRangeInterval(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validateRangeInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}
