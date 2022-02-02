# Shipa::Framework::Item AppAutoDiscovery

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#appselector" title="AppSelector">AppSelector</a>" : <i>[ <a href="appselectorlabels.md">AppSelectorLabels</a>, ... ]</i>,
    "<a href="#suffix" title="Suffix">Suffix</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#appselector" title="AppSelector">AppSelector</a>: <i>
      - <a href="appselectorlabels.md">AppSelectorLabels</a></i>
<a href="#suffix" title="Suffix">Suffix</a>: <i>String</i>
</pre>

## Properties

#### AppSelector

_Required_: No

_Type_: List of <a href="appselectorlabels.md">AppSelectorLabels</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Suffix

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

