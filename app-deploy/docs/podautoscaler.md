# Shipa::AppDeploy::Item PodAutoScaler

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#maxreplicas" title="MaxReplicas">MaxReplicas</a>" : <i>Integer</i>,
    "<a href="#minreplicas" title="MinReplicas">MinReplicas</a>" : <i>Integer</i>,
    "<a href="#targetcpuutilizationpercentage" title="TargetCPUUtilizationPercentage">TargetCPUUtilizationPercentage</a>" : <i>Integer</i>
}
</pre>

### YAML

<pre>
<a href="#maxreplicas" title="MaxReplicas">MaxReplicas</a>: <i>Integer</i>
<a href="#minreplicas" title="MinReplicas">MinReplicas</a>: <i>Integer</i>
<a href="#targetcpuutilizationpercentage" title="TargetCPUUtilizationPercentage">TargetCPUUtilizationPercentage</a>: <i>Integer</i>
</pre>

## Properties

#### MaxReplicas

Max replicas

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### MinReplicas

Min replicas

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### TargetCPUUtilizationPercentage

Target CPU utilization percentage

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

