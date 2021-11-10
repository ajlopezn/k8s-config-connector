// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Config Connector and manual
//     changes will be clobbered when the file is regenerated.
//
// ----------------------------------------------------------------------------

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type AlertpolicyAggregations struct {
	/* The alignment period for per-time
	series alignment. If present,
	alignmentPeriod must be at least
	60 seconds. After per-time series
	alignment, each time series will
	contain data points only on the
	period boundaries. If
	perSeriesAligner is not specified
	or equals ALIGN_NONE, then this
	field is ignored. If
	perSeriesAligner is specified and
	does not equal ALIGN_NONE, then
	this field must be defined;
	otherwise an error is returned. */
	// +optional
	AlignmentPeriod *string `json:"alignmentPeriod,omitempty"`

	/* The approach to be used to combine
	time series. Not all reducer
	functions may be applied to all
	time series, depending on the
	metric type and the value type of
	the original time series.
	Reduction may change the metric
	type of value type of the time
	series.Time series data must be
	aligned in order to perform cross-
	time series reduction. If
	crossSeriesReducer is specified,
	then perSeriesAligner must be
	specified and not equal ALIGN_NONE
	and alignmentPeriod must be
	specified; otherwise, an error is
	returned. Possible values: ["REDUCE_NONE", "REDUCE_MEAN", "REDUCE_MIN", "REDUCE_MAX", "REDUCE_SUM", "REDUCE_STDDEV", "REDUCE_COUNT", "REDUCE_COUNT_TRUE", "REDUCE_COUNT_FALSE", "REDUCE_FRACTION_TRUE", "REDUCE_PERCENTILE_99", "REDUCE_PERCENTILE_95", "REDUCE_PERCENTILE_50", "REDUCE_PERCENTILE_05"]. */
	// +optional
	CrossSeriesReducer *string `json:"crossSeriesReducer,omitempty"`

	/* The set of fields to preserve when
	crossSeriesReducer is specified.
	The groupByFields determine how
	the time series are partitioned
	into subsets prior to applying the
	aggregation function. Each subset
	contains time series that have the
	same value for each of the
	grouping fields. Each individual
	time series is a member of exactly
	one subset. The crossSeriesReducer
	is applied to each subset of time
	series. It is not possible to
	reduce across different resource
	types, so this field implicitly
	contains resource.type. Fields not
	specified in groupByFields are
	aggregated away. If groupByFields
	is not specified and all the time
	series have the same resource
	type, then the time series are
	aggregated into a single output
	time series. If crossSeriesReducer
	is not defined, this field is
	ignored. */
	// +optional
	GroupByFields []string `json:"groupByFields,omitempty"`

	/* The approach to be used to align
	individual time series. Not all
	alignment functions may be applied
	to all time series, depending on
	the metric type and value type of
	the original time series.
	Alignment may change the metric
	type or the value type of the time
	series.Time series data must be
	aligned in order to perform cross-
	time series reduction. If
	crossSeriesReducer is specified,
	then perSeriesAligner must be
	specified and not equal ALIGN_NONE
	and alignmentPeriod must be
	specified; otherwise, an error is
	returned. Possible values: ["ALIGN_NONE", "ALIGN_DELTA", "ALIGN_RATE", "ALIGN_INTERPOLATE", "ALIGN_NEXT_OLDER", "ALIGN_MIN", "ALIGN_MAX", "ALIGN_MEAN", "ALIGN_COUNT", "ALIGN_SUM", "ALIGN_STDDEV", "ALIGN_COUNT_TRUE", "ALIGN_COUNT_FALSE", "ALIGN_FRACTION_TRUE", "ALIGN_PERCENTILE_99", "ALIGN_PERCENTILE_95", "ALIGN_PERCENTILE_50", "ALIGN_PERCENTILE_05", "ALIGN_PERCENT_CHANGE"]. */
	// +optional
	PerSeriesAligner *string `json:"perSeriesAligner,omitempty"`
}

type AlertpolicyConditionAbsent struct {
	/* Specifies the alignment of data points in
	individual time series as well as how to
	combine the retrieved time series together
	(such as when aggregating multiple streams
	on each resource to a single stream for each
	resource or when aggregating streams across
	all members of a group of resources).
	Multiple aggregations are applied in the
	order specified. */
	// +optional
	Aggregations []AlertpolicyAggregations `json:"aggregations,omitempty"`

	/* The amount of time that a time series must
	fail to report new data to be considered
	failing. Currently, only values that are a
	multiple of a minute--e.g. 60s, 120s, or 300s
	--are supported. */
	Duration string `json:"duration"`

	/* A filter that identifies which time series
	should be compared with the threshold.The
	filter is similar to the one that is
	specified in the
	MetricService.ListTimeSeries request (that
	call is useful to verify the time series
	that will be retrieved / processed) and must
	specify the metric type and optionally may
	contain restrictions on resource type,
	resource labels, and metric labels. This
	field may not exceed 2048 Unicode characters
	in length. */
	// +optional
	Filter *string `json:"filter,omitempty"`

	/* The number/percent of time series for which
	the comparison must hold in order for the
	condition to trigger. If unspecified, then
	the condition will trigger if the comparison
	is true for any of the time series that have
	been identified by filter and aggregations. */
	// +optional
	Trigger *AlertpolicyTrigger `json:"trigger,omitempty"`
}

type AlertpolicyConditionMonitoringQueryLanguage struct {
	/* The amount of time that a time series must
	violate the threshold to be considered
	failing. Currently, only values that are a
	multiple of a minute--e.g., 0, 60, 120, or
	300 seconds--are supported. If an invalid
	value is given, an error will be returned.
	When choosing a duration, it is useful to
	keep in mind the frequency of the underlying
	time series data (which may also be affected
	by any alignments specified in the
	aggregations field); a good duration is long
	enough so that a single outlier does not
	generate spurious alerts, but short enough
	that unhealthy states are detected and
	alerted on quickly. */
	Duration string `json:"duration"`

	/* Monitoring Query Language query that outputs a boolean stream. */
	Query string `json:"query"`

	/* The number/percent of time series for which
	the comparison must hold in order for the
	condition to trigger. If unspecified, then
	the condition will trigger if the comparison
	is true for any of the time series that have
	been identified by filter and aggregations,
	or by the ratio, if denominator_filter and
	denominator_aggregations are specified. */
	// +optional
	Trigger *AlertpolicyTrigger `json:"trigger,omitempty"`
}

type AlertpolicyConditionThreshold struct {
	/* Specifies the alignment of data points in
	individual time series as well as how to
	combine the retrieved time series together
	(such as when aggregating multiple streams
	on each resource to a single stream for each
	resource or when aggregating streams across
	all members of a group of resources).
	Multiple aggregations are applied in the
	order specified.This field is similar to the
	one in the MetricService.ListTimeSeries
	request. It is advisable to use the
	ListTimeSeries method when debugging this
	field. */
	// +optional
	Aggregations []AlertpolicyAggregations `json:"aggregations,omitempty"`

	/* The comparison to apply between the time
	series (indicated by filter and aggregation)
	and the threshold (indicated by
	threshold_value). The comparison is applied
	on each time series, with the time series on
	the left-hand side and the threshold on the
	right-hand side. Only COMPARISON_LT and
	COMPARISON_GT are supported currently. Possible values: ["COMPARISON_GT", "COMPARISON_GE", "COMPARISON_LT", "COMPARISON_LE", "COMPARISON_EQ", "COMPARISON_NE"]. */
	Comparison string `json:"comparison"`

	/* Specifies the alignment of data points in
	individual time series selected by
	denominatorFilter as well as how to combine
	the retrieved time series together (such as
	when aggregating multiple streams on each
	resource to a single stream for each
	resource or when aggregating streams across
	all members of a group of resources).When
	computing ratios, the aggregations and
	denominator_aggregations fields must use the
	same alignment period and produce time
	series that have the same periodicity and
	labels.This field is similar to the one in
	the MetricService.ListTimeSeries request. It
	is advisable to use the ListTimeSeries
	method when debugging this field. */
	// +optional
	DenominatorAggregations []AlertpolicyDenominatorAggregations `json:"denominatorAggregations,omitempty"`

	/* A filter that identifies a time series that
	should be used as the denominator of a ratio
	that will be compared with the threshold. If
	a denominator_filter is specified, the time
	series specified by the filter field will be
	used as the numerator.The filter is similar
	to the one that is specified in the
	MetricService.ListTimeSeries request (that
	call is useful to verify the time series
	that will be retrieved / processed) and must
	specify the metric type and optionally may
	contain restrictions on resource type,
	resource labels, and metric labels. This
	field may not exceed 2048 Unicode characters
	in length. */
	// +optional
	DenominatorFilter *string `json:"denominatorFilter,omitempty"`

	/* The amount of time that a time series must
	violate the threshold to be considered
	failing. Currently, only values that are a
	multiple of a minute--e.g., 0, 60, 120, or
	300 seconds--are supported. If an invalid
	value is given, an error will be returned.
	When choosing a duration, it is useful to
	keep in mind the frequency of the underlying
	time series data (which may also be affected
	by any alignments specified in the
	aggregations field); a good duration is long
	enough so that a single outlier does not
	generate spurious alerts, but short enough
	that unhealthy states are detected and
	alerted on quickly. */
	Duration string `json:"duration"`

	/* A filter that identifies which time series
	should be compared with the threshold.The
	filter is similar to the one that is
	specified in the
	MetricService.ListTimeSeries request (that
	call is useful to verify the time series
	that will be retrieved / processed) and must
	specify the metric type and optionally may
	contain restrictions on resource type,
	resource labels, and metric labels. This
	field may not exceed 2048 Unicode characters
	in length. */
	// +optional
	Filter *string `json:"filter,omitempty"`

	/* A value against which to compare the time
	series. */
	// +optional
	ThresholdValue *float64 `json:"thresholdValue,omitempty"`

	/* The number/percent of time series for which
	the comparison must hold in order for the
	condition to trigger. If unspecified, then
	the condition will trigger if the comparison
	is true for any of the time series that have
	been identified by filter and aggregations,
	or by the ratio, if denominator_filter and
	denominator_aggregations are specified. */
	// +optional
	Trigger *AlertpolicyTrigger `json:"trigger,omitempty"`
}

type AlertpolicyConditions struct {
	/* A condition that checks that a time series
	continues to receive new data points. */
	// +optional
	ConditionAbsent *AlertpolicyConditionAbsent `json:"conditionAbsent,omitempty"`

	/* A Monitoring Query Language query that outputs a boolean stream. */
	// +optional
	ConditionMonitoringQueryLanguage *AlertpolicyConditionMonitoringQueryLanguage `json:"conditionMonitoringQueryLanguage,omitempty"`

	/* A condition that compares a time series against a
	threshold. */
	// +optional
	ConditionThreshold *AlertpolicyConditionThreshold `json:"conditionThreshold,omitempty"`

	/* A short name or phrase used to identify the
	condition in dashboards, notifications, and
	incidents. To avoid confusion, don't use the same
	display name for multiple conditions in the same
	policy. */
	DisplayName string `json:"displayName"`

	/* The unique resource name for this condition.
	Its syntax is:
	projects/[PROJECT_ID]/alertPolicies/[POLICY_ID]/conditions/[CONDITION_ID]
	[CONDITION_ID] is assigned by Stackdriver Monitoring when
	the condition is created as part of a new or updated alerting
	policy. */
	// +optional
	Name *string `json:"name,omitempty"`
}

type AlertpolicyDenominatorAggregations struct {
	/* The alignment period for per-time
	series alignment. If present,
	alignmentPeriod must be at least
	60 seconds. After per-time series
	alignment, each time series will
	contain data points only on the
	period boundaries. If
	perSeriesAligner is not specified
	or equals ALIGN_NONE, then this
	field is ignored. If
	perSeriesAligner is specified and
	does not equal ALIGN_NONE, then
	this field must be defined;
	otherwise an error is returned. */
	// +optional
	AlignmentPeriod *string `json:"alignmentPeriod,omitempty"`

	/* The approach to be used to combine
	time series. Not all reducer
	functions may be applied to all
	time series, depending on the
	metric type and the value type of
	the original time series.
	Reduction may change the metric
	type of value type of the time
	series.Time series data must be
	aligned in order to perform cross-
	time series reduction. If
	crossSeriesReducer is specified,
	then perSeriesAligner must be
	specified and not equal ALIGN_NONE
	and alignmentPeriod must be
	specified; otherwise, an error is
	returned. Possible values: ["REDUCE_NONE", "REDUCE_MEAN", "REDUCE_MIN", "REDUCE_MAX", "REDUCE_SUM", "REDUCE_STDDEV", "REDUCE_COUNT", "REDUCE_COUNT_TRUE", "REDUCE_COUNT_FALSE", "REDUCE_FRACTION_TRUE", "REDUCE_PERCENTILE_99", "REDUCE_PERCENTILE_95", "REDUCE_PERCENTILE_50", "REDUCE_PERCENTILE_05"]. */
	// +optional
	CrossSeriesReducer *string `json:"crossSeriesReducer,omitempty"`

	/* The set of fields to preserve when
	crossSeriesReducer is specified.
	The groupByFields determine how
	the time series are partitioned
	into subsets prior to applying the
	aggregation function. Each subset
	contains time series that have the
	same value for each of the
	grouping fields. Each individual
	time series is a member of exactly
	one subset. The crossSeriesReducer
	is applied to each subset of time
	series. It is not possible to
	reduce across different resource
	types, so this field implicitly
	contains resource.type. Fields not
	specified in groupByFields are
	aggregated away. If groupByFields
	is not specified and all the time
	series have the same resource
	type, then the time series are
	aggregated into a single output
	time series. If crossSeriesReducer
	is not defined, this field is
	ignored. */
	// +optional
	GroupByFields []string `json:"groupByFields,omitempty"`

	/* The approach to be used to align
	individual time series. Not all
	alignment functions may be applied
	to all time series, depending on
	the metric type and value type of
	the original time series.
	Alignment may change the metric
	type or the value type of the time
	series.Time series data must be
	aligned in order to perform cross-
	time series reduction. If
	crossSeriesReducer is specified,
	then perSeriesAligner must be
	specified and not equal ALIGN_NONE
	and alignmentPeriod must be
	specified; otherwise, an error is
	returned. Possible values: ["ALIGN_NONE", "ALIGN_DELTA", "ALIGN_RATE", "ALIGN_INTERPOLATE", "ALIGN_NEXT_OLDER", "ALIGN_MIN", "ALIGN_MAX", "ALIGN_MEAN", "ALIGN_COUNT", "ALIGN_SUM", "ALIGN_STDDEV", "ALIGN_COUNT_TRUE", "ALIGN_COUNT_FALSE", "ALIGN_FRACTION_TRUE", "ALIGN_PERCENTILE_99", "ALIGN_PERCENTILE_95", "ALIGN_PERCENTILE_50", "ALIGN_PERCENTILE_05", "ALIGN_PERCENT_CHANGE"]. */
	// +optional
	PerSeriesAligner *string `json:"perSeriesAligner,omitempty"`
}

type AlertpolicyDocumentation struct {
	/* The text of the documentation, interpreted according to mimeType.
	The content may not exceed 8,192 Unicode characters and may not
	exceed more than 10,240 bytes when encoded in UTF-8 format,
	whichever is smaller. */
	// +optional
	Content *string `json:"content,omitempty"`

	/* The format of the content field. Presently, only the value
	"text/markdown" is supported. */
	// +optional
	MimeType *string `json:"mimeType,omitempty"`
}

type AlertpolicyTrigger struct {
	/* The absolute number of time series
	that must fail the predicate for the
	condition to be triggered. */
	// +optional
	Count *int `json:"count,omitempty"`

	/* The percentage of time series that
	must fail the predicate for the
	condition to be triggered. */
	// +optional
	Percent *float64 `json:"percent,omitempty"`
}

type MonitoringAlertPolicySpec struct {
	/* How to combine the results of multiple conditions to
	determine if an incident should be opened. Possible values: ["AND", "OR", "AND_WITH_MATCHING_RESOURCE"]. */
	Combiner string `json:"combiner"`

	/* A list of conditions for the policy. The conditions are combined by
	AND or OR according to the combiner field. If the combined conditions
	evaluate to true, then an incident is created. A policy can have from
	one to six conditions. */
	Conditions []AlertpolicyConditions `json:"conditions"`

	/* A short name or phrase used to identify the policy in
	dashboards, notifications, and incidents. To avoid confusion, don't use
	the same display name for multiple policies in the same project. The
	name is limited to 512 Unicode characters. */
	DisplayName string `json:"displayName"`

	/* Documentation that is included with notifications and incidents related
	to this policy. Best practice is for the documentation to include information
	to help responders understand, mitigate, escalate, and correct the underlying
	problems detected by the alerting policy. Notification channels that have
	limited capacity might not show this documentation. */
	// +optional
	Documentation *AlertpolicyDocumentation `json:"documentation,omitempty"`

	/* Whether or not the policy is enabled. The default is true. */
	// +optional
	Enabled *bool `json:"enabled,omitempty"`

	/*  */
	// +optional
	NotificationChannels []v1alpha1.ResourceRef `json:"notificationChannels,omitempty"`

	/* Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

type AlertpolicyCreationRecordStatus struct {
	/* When the change occurred. */
	MutateTime string `json:"mutateTime,omitempty"`

	/* The email address of the user making the change. */
	MutatedBy string `json:"mutatedBy,omitempty"`
}

type MonitoringAlertPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   MonitoringAlertPolicy's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* A read-only record of the creation of the alerting policy.
	If provided in a call to create or update, this field will
	be ignored. */
	CreationRecord []AlertpolicyCreationRecordStatus `json:"creationRecord,omitempty"`
	/* The unique resource name for this policy.
	Its syntax is: projects/[PROJECT_ID]/alertPolicies/[ALERT_POLICY_ID]. */
	Name string `json:"name,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	ObservedGeneration int `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MonitoringAlertPolicy is the Schema for the monitoring API
// +k8s:openapi-gen=true
type MonitoringAlertPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MonitoringAlertPolicySpec   `json:"spec,omitempty"`
	Status MonitoringAlertPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MonitoringAlertPolicyList contains a list of MonitoringAlertPolicy
type MonitoringAlertPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MonitoringAlertPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MonitoringAlertPolicy{}, &MonitoringAlertPolicyList{})
}
