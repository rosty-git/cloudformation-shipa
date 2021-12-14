# Shipa::Framework::Item PoolNode

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#drivers" title="Drivers">Drivers</a>" : <i>[ String, ... ]</i>,
    "<a href="#autoscale" title="AutoScale">AutoScale</a>" : <i><a href="poolautoscale.md">PoolAutoScale</a></i>
}
</pre>

### YAML

<pre>
<a href="#drivers" title="Drivers">Drivers</a>: <i>
      - String</i>
<a href="#autoscale" title="AutoScale">AutoScale</a>: <i><a href="poolautoscale.md">PoolAutoScale</a></i>
</pre>

## Properties

#### Drivers

_Required_: No

_Type_: List of String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### AutoScale

_Required_: No

_Type_: <a href="poolautoscale.md">PoolAutoScale</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

