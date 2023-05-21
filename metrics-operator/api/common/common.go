package common

import (
	"github.com/pkg/errors"
)

const KeptnMetricProviderName = "keptn-metric"

var ErrForbiddenIntervalError = errors.New("Forbidden! KeptnMetrics should define an interval different from 5mins")
