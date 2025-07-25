// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/compute/RegionAutoscaler.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc_next/services/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package compute

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/pkg/tpgresource"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/pkg/verify"
)

const ComputeRegionAutoscalerSchemaName string = "google_compute_region_autoscaler"

func ResourceComputeRegionAutoscaler() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"autoscaling_policy": {
				Type:     schema.TypeList,
				Required: true,
				Description: `The configuration parameters for the autoscaling algorithm. You can
define one or more of the policies for an autoscaler: cpuUtilization,
customMetricUtilizations, and loadBalancingUtilization.

If none of these are specified, the default will be to autoscale based
on cpuUtilization to 0.6 or 60%.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"max_replicas": {
							Type:     schema.TypeInt,
							Required: true,
							Description: `The maximum number of instances that the autoscaler can scale up
to. This is required when creating or updating an autoscaler. The
maximum number of replicas should not be lower than minimal number
of replicas.`,
						},
						"min_replicas": {
							Type:     schema.TypeInt,
							Required: true,
							Description: `The minimum number of replicas that the autoscaler can scale down
to. This cannot be less than 0. If not provided, autoscaler will
choose a default value depending on maximum number of instances
allowed.`,
						},
						"cooldown_period": {
							Type:     schema.TypeInt,
							Optional: true,
							Description: `The number of seconds that the autoscaler should wait before it
starts collecting information from a new instance. This prevents
the autoscaler from collecting information when the instance is
initializing, during which the collected usage would not be
reliable. The default time autoscaler waits is 60 seconds.

Virtual machine initialization times might vary because of
numerous factors. We recommend that you test how long an
instance may take to initialize. To do this, create an instance
and time the startup process.`,
							Default: 60,
						},
						"cpu_utilization": {
							Type:     schema.TypeList,
							Computed: true,
							Optional: true,
							Description: `Defines the CPU utilization policy that allows the autoscaler to
scale based on the average CPU utilization of a managed instance
group.`,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"target": {
										Type:     schema.TypeFloat,
										Required: true,
										Description: `The target CPU utilization that the autoscaler should maintain.
Must be a float value in the range (0, 1]. If not specified, the
default is 0.6.

If the CPU level is below the target utilization, the autoscaler
scales down the number of instances until it reaches the minimum
number of instances you specified or until the average CPU of
your instances reaches the target utilization.

If the average CPU is above the target utilization, the autoscaler
scales up until it reaches the maximum number of instances you
specified or until the average utilization reaches the target
utilization.`,
									},
									"predictive_method": {
										Type:     schema.TypeString,
										Optional: true,
										Description: `Indicates whether predictive autoscaling based on CPU metric is enabled. Valid values are:

- NONE (default). No predictive method is used. The autoscaler scales the group to meet current demand based on real-time metrics.

- OPTIMIZE_AVAILABILITY. Predictive autoscaling improves availability by monitoring daily and weekly load patterns and scaling out ahead of anticipated demand.`,
										Default: "NONE",
									},
								},
							},
						},
						"load_balancing_utilization": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Configuration parameters of autoscaling based on a load balancer.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"target": {
										Type:     schema.TypeFloat,
										Required: true,
										Description: `Fraction of backend capacity utilization (set in HTTP(s) load
balancing configuration) that autoscaler should maintain. Must
be a positive float value. If not defined, the default is 0.8.`,
									},
								},
							},
						},
						"metric": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Configuration parameters of autoscaling based on a custom metric.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:     schema.TypeString,
										Required: true,
										Description: `The identifier (type) of the Stackdriver Monitoring metric.
The metric cannot have negative values.

The metric must have a value type of INT64 or DOUBLE.`,
									},
									"filter": {
										Type:     schema.TypeString,
										Optional: true,
										Description: `A filter string to be used as the filter string for
a Stackdriver Monitoring TimeSeries.list API call.
This filter is used to select a specific TimeSeries for
the purpose of autoscaling and to determine whether the metric
is exporting per-instance or per-group data.

You can only use the AND operator for joining selectors.
You can only use direct equality comparison operator (=) without
any functions for each selector.
You can specify the metric in both the filter string and in the
metric field. However, if specified in both places, the metric must
be identical.

The monitored resource type determines what kind of values are
expected for the metric. If it is a gce_instance, the autoscaler
expects the metric to include a separate TimeSeries for each
instance in a group. In such a case, you cannot filter on resource
labels.

If the resource type is any other value, the autoscaler expects
this metric to contain values that apply to the entire autoscaled
instance group and resource label filtering can be performed to
point autoscaler at the correct TimeSeries to scale upon.
This is called a per-group metric for the purpose of autoscaling.

If not specified, the type defaults to gce_instance.

You should provide a filter that is selective enough to pick just
one TimeSeries for the autoscaled group or for each of the instances
(if you are using gce_instance resource type). If multiple
TimeSeries are returned upon the query execution, the autoscaler
will sum their respective values to obtain its scaling value.`,
									},
									"single_instance_assignment": {
										Type:     schema.TypeFloat,
										Optional: true,
										Description: `If scaling is based on a per-group metric value that represents the
total amount of work to be done or resource usage, set this value to
an amount assigned for a single instance of the scaled group.
The autoscaler will keep the number of instances proportional to the
value of this metric, the metric itself should not change value due
to group resizing.

For example, a good metric to use with the target is
'pubsub.googleapis.com/subscription/num_undelivered_messages'
or a custom metric exporting the total number of requests coming to
your instances.

A bad example would be a metric exporting an average or median
latency, since this value can't include a chunk assignable to a
single instance, it could be better used with utilization_target
instead.`,
									},
									"target": {
										Type:     schema.TypeFloat,
										Optional: true,
										Description: `The target value of the metric that autoscaler should
maintain. This must be a positive value. A utilization
metric scales number of virtual machines handling requests
to increase or decrease proportionally to the metric.

For example, a good metric to use as a utilizationTarget is
www.googleapis.com/compute/instance/network/received_bytes_count.
The autoscaler will work to keep this value constant for each
of the instances.`,
									},
									"type": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: verify.ValidateEnum([]string{"GAUGE", "DELTA_PER_SECOND", "DELTA_PER_MINUTE", ""}),
										Description: `Defines how target utilization value is expressed for a
Stackdriver Monitoring metric. Possible values: ["GAUGE", "DELTA_PER_SECOND", "DELTA_PER_MINUTE"]`,
									},
								},
							},
						},
						"mode": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Defines operating mode for this policy.`,
							Default:     "ON",
						},
						"scale_down_control": {
							Type:     schema.TypeList,
							Optional: true,
							Description: `Defines scale down controls to reduce the risk of response latency
and outages due to abrupt scale-in events`,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"max_scaled_down_replicas": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: `A nested object resource.`,
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"fixed": {
													Type:     schema.TypeInt,
													Optional: true,
													Description: `Specifies a fixed number of VM instances. This must be a positive
integer.`,
													AtLeastOneOf: []string{"autoscaling_policy.0.scale_down_control.0.max_scaled_down_replicas.0.fixed", "autoscaling_policy.0.scale_down_control.0.max_scaled_down_replicas.0.percent"},
												},
												"percent": {
													Type:     schema.TypeInt,
													Optional: true,
													Description: `Specifies a percentage of instances between 0 to 100%, inclusive.
For example, specify 80 for 80%.`,
													AtLeastOneOf: []string{"autoscaling_policy.0.scale_down_control.0.max_scaled_down_replicas.0.fixed", "autoscaling_policy.0.scale_down_control.0.max_scaled_down_replicas.0.percent"},
												},
											},
										},
										AtLeastOneOf: []string{"autoscaling_policy.0.scale_down_control.0.max_scaled_down_replicas", "autoscaling_policy.0.scale_down_control.0.time_window_sec"},
									},
									"time_window_sec": {
										Type:     schema.TypeInt,
										Optional: true,
										Description: `How long back autoscaling should look when computing recommendations
to include directives regarding slower scale down, as described above.`,
										AtLeastOneOf: []string{"autoscaling_policy.0.scale_down_control.0.max_scaled_down_replicas", "autoscaling_policy.0.scale_down_control.0.time_window_sec"},
									},
								},
							},
						},
						"scale_in_control": {
							Type:     schema.TypeList,
							Optional: true,
							Description: `Defines scale in controls to reduce the risk of response latency
and outages due to abrupt scale-in events`,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"max_scaled_in_replicas": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: `A nested object resource.`,
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"fixed": {
													Type:     schema.TypeInt,
													Optional: true,
													Description: `Specifies a fixed number of VM instances. This must be a positive
integer.`,
													AtLeastOneOf: []string{"autoscaling_policy.0.scale_in_control.0.max_scaled_in_replicas.0.fixed", "autoscaling_policy.0.scale_in_control.0.max_scaled_in_replicas.0.percent"},
												},
												"percent": {
													Type:     schema.TypeInt,
													Optional: true,
													Description: `Specifies a percentage of instances between 0 to 100%, inclusive.
For example, specify 80 for 80%.`,
													AtLeastOneOf: []string{"autoscaling_policy.0.scale_in_control.0.max_scaled_in_replicas.0.fixed", "autoscaling_policy.0.scale_in_control.0.max_scaled_in_replicas.0.percent"},
												},
											},
										},
										AtLeastOneOf: []string{"autoscaling_policy.0.scale_in_control.0.max_scaled_in_replicas", "autoscaling_policy.0.scale_in_control.0.time_window_sec"},
									},
									"time_window_sec": {
										Type:     schema.TypeInt,
										Optional: true,
										Description: `How long back autoscaling should look when computing recommendations
to include directives regarding slower scale down, as described above.`,
										AtLeastOneOf: []string{"autoscaling_policy.0.scale_in_control.0.max_scaled_in_replicas", "autoscaling_policy.0.scale_in_control.0.time_window_sec"},
									},
								},
							},
						},
						"scaling_schedules": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: `Scaling schedules defined for an autoscaler. Multiple schedules can be set on an autoscaler and they can overlap.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"duration_sec": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: `The duration of time intervals (in seconds) for which this scaling schedule will be running. The minimum allowed value is 300.`,
									},
									"min_required_replicas": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: `Minimum number of VM instances that autoscaler will recommend in time intervals starting according to schedule.`,
									},
									"schedule": {
										Type:        schema.TypeString,
										Required:    true,
										Description: `The start timestamps of time intervals when this scaling schedule should provide a scaling signal. This field uses the extended cron format (with an optional year field).`,
									},
									"description": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: `A description of a scaling schedule.`,
									},
									"disabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: `A boolean value that specifies if a scaling schedule can influence autoscaler recommendations. If set to true, then a scaling schedule has no effect.`,
										Default:     false,
									},
									"time_zone": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: `The time zone to be used when interpreting the schedule. The value of this field must be a time zone name from the tz database: http://en.wikipedia.org/wiki/Tz_database.`,
										Default:     "UTC",
									},
								},
							},
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateGCEName,
				Description: `Name of the resource. The name must be 1-63 characters long and match
the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
first character must be a lowercase letter, and all following
characters must be a dash, lowercase letter, or digit, except the last
character, which cannot be a dash.`,
			},
			"target": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `URL of the managed instance group that this autoscaler will scale.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `An optional description of this resource.`,
			},
			"region": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `URL of the region where the instance group resides.`,
			},
			"creation_timestamp": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Creation timestamp in RFC3339 text format.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		UseJSONNumber: true,
	}
}
