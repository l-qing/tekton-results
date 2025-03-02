package config

import (
	corev1 "k8s.io/api/core/v1"
	"knative.dev/pkg/metrics"
)

const (
	// metricsTaskrunLevel determines to what level to aggregate metrics
	// for taskrun
	metricsTaskrunLevelKey = "metrics.taskrun.level"

	// metricsPipelinerunLevel determines to what level to aggregate metrics
	// for pipelinerun
	metricsPipelinerunLevelKey = "metrics.pipelinerun.level"
	// metricsDurationTaskrunType determines what type of
	// metrics to use for aggregating duration for taskrun
	metricsDurationTaskrunType = "metrics.taskrun.duration-type"
	// metricsDurationPipelinerunType determines what type of
	// metrics to use for aggregating duration for pipelinerun
	metricsDurationPipelinerunType = "metrics.pipelinerun.duration-type"

	// DefaultTaskrunLevel determines to what level to aggregate metrics
	// when it isn't specified in configmap
	DefaultTaskrunLevel = TaskrunLevelAtTask
	// TaskrunLevelAtTask specify that aggregation will be done at task level
	TaskrunLevelAtTask = "task"
	// TaskrunLevelAtNS specify that aggregation will be done at namespace level
	TaskrunLevelAtNS = "namespace"
	// DefaultPipelinerunLevel determines to what level to aggregate metrics
	// when it isn't specified in configmap
	DefaultPipelinerunLevel = PipelinerunLevelAtPipeline
	// PipelinerunLevelAtPipeline specify that aggregation will be done at
	// pipeline level
	PipelinerunLevelAtPipeline = "pipeline"
	// PipelinerunLevelAtNS specify that aggregation will be done at
	// namespace level
	PipelinerunLevelAtNS = "namespace"

	// DefaultDurationTaskrunType determines what type
	// of metrics to use when we don't specify one in
	// configmap
	DefaultDurationTaskrunType = "histogram"
	// DurationTaskrunTypeHistogram specify that histogram
	// type metrics need to be used for Duration of Taskrun
	DurationTaskrunTypeHistogram = "histogram"
	// DurationTaskrunTypeLastValue specify that lastValue or
	// gauge type metrics need to be used for Duration of Taskrun
	DurationTaskrunTypeLastValue = "lastvalue"

	// DefaultDurationPipelinerunType determines what type
	// of metrics to use when we don't specify one in
	// configmap
	DefaultDurationPipelinerunType = "histogram"
	// DurationPipelinerunTypeHistogram specify that histogram
	// type metrics need to be used for Duration of Pipelinerun
	DurationPipelinerunTypeHistogram = "histogram"
	// DurationPipelinerunTypeLastValue specify that lastValue or
	// gauge type metrics need to be used for Duration of Pipelinerun
	DurationPipelinerunTypeLastValue = "lastvalue"
)

// Metrics holds the configurations for the metrics
type Metrics struct {
	TaskrunLevel            string
	PipelinerunLevel        string
	DurationTaskrunType     string
	DurationPipelinerunType string
}

// DeepCopy copying the receiver, creating a new Metrics.
// deepcopy-gen hasn't been introduced in results repo, so handcraft here for now
func (cfg *Metrics) DeepCopy() *Metrics {
	return &Metrics{
		TaskrunLevel:            cfg.TaskrunLevel,
		PipelinerunLevel:        cfg.PipelinerunLevel,
		DurationTaskrunType:     cfg.DurationTaskrunType,
		DurationPipelinerunType: cfg.DurationPipelinerunType,
	}
}

// GetMetricsConfigName returns the name of the configmap containing all
// customizations for the storage bucket.
func GetMetricsConfigName() string {
	return metrics.ConfigMapName()
}

// Equals returns true if two Configs are identical
func (cfg *Metrics) Equals(other *Metrics) bool {
	if cfg == nil && other == nil {
		return true
	}

	if cfg == nil || other == nil {
		return false
	}

	return other.TaskrunLevel == cfg.TaskrunLevel &&
		other.PipelinerunLevel == cfg.PipelinerunLevel &&
		other.DurationTaskrunType == cfg.DurationTaskrunType &&
		other.DurationPipelinerunType == cfg.DurationPipelinerunType
}

// newMetricsFromMap returns a Config given a map corresponding to a ConfigMap
func newMetricsFromMap(cfgMap map[string]string) (*Metrics, error) {
	tc := Metrics{
		TaskrunLevel:            DefaultTaskrunLevel,
		PipelinerunLevel:        DefaultPipelinerunLevel,
		DurationTaskrunType:     DefaultDurationTaskrunType,
		DurationPipelinerunType: DefaultDurationPipelinerunType,
	}

	if taskrunLevel, ok := cfgMap[metricsTaskrunLevelKey]; ok {
		tc.TaskrunLevel = taskrunLevel
	}

	if pipelinerunLevel, ok := cfgMap[metricsPipelinerunLevelKey]; ok {
		tc.PipelinerunLevel = pipelinerunLevel
	}
	if durationTaskrun, ok := cfgMap[metricsDurationTaskrunType]; ok {
		tc.DurationTaskrunType = durationTaskrun
	}
	if durationPipelinerun, ok := cfgMap[metricsDurationPipelinerunType]; ok {
		tc.DurationPipelinerunType = durationPipelinerun
	}
	return &tc, nil
}

// NewMetricsFromConfigMap returns a Config for the given configmap
func NewMetricsFromConfigMap(config *corev1.ConfigMap) (*Metrics, error) {
	return newMetricsFromMap(config.Data)
}
