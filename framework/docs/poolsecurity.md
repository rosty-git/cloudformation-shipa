# Shipa::Framework::Item PoolSecurity

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#disablescan" title="DisableScan">DisableScan</a>" : <i>Boolean</i>,
    "<a href="#scanplatformlayers" title="ScanPlatformLayers">ScanPlatformLayers</a>" : <i>Boolean</i>,
    "<a href="#ignorecomponents" title="IgnoreComponents">IgnoreComponents</a>" : <i>[ String, ... ]</i>,
    "<a href="#ignorecves" title="IgnoreCVES">IgnoreCVES</a>" : <i>[ String, ... ]</i>
}
</pre>

### YAML

<pre>
<a href="#disablescan" title="DisableScan">DisableScan</a>: <i>Boolean</i>
<a href="#scanplatformlayers" title="ScanPlatformLayers">ScanPlatformLayers</a>: <i>Boolean</i>
<a href="#ignorecomponents" title="IgnoreComponents">IgnoreComponents</a>: <i>
      - String</i>
<a href="#ignorecves" title="IgnoreCVES">IgnoreCVES</a>: <i>
      - String</i>
</pre>

## Properties

#### DisableScan

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ScanPlatformLayers

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### IgnoreComponents

_Required_: No

_Type_: List of String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### IgnoreCVES

_Required_: No

_Type_: List of String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

