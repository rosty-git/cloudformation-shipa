# Shipa::Framework::Item NetworkPolicyRule

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#id" title="ID">ID</a>" : <i>String</i>,
    "<a href="#enabled" title="Enabled">Enabled</a>" : <i>Boolean</i>,
    "<a href="#description" title="Description">Description</a>" : <i>String</i>,
    "<a href="#ports" title="Ports">Ports</a>" : <i>[ <a href="networkport.md">NetworkPort</a>, ... ]</i>,
    "<a href="#peers" title="Peers">Peers</a>" : <i>[ <a href="networkpeer.md">NetworkPeer</a>, ... ]</i>,
    "<a href="#allowedapps" title="AllowedApps">AllowedApps</a>" : <i>[ String, ... ]</i>,
    "<a href="#allowedpools" title="AllowedPools">AllowedPools</a>" : <i>[ String, ... ]</i>
}
</pre>

### YAML

<pre>
<a href="#id" title="ID">ID</a>: <i>String</i>
<a href="#enabled" title="Enabled">Enabled</a>: <i>Boolean</i>
<a href="#description" title="Description">Description</a>: <i>String</i>
<a href="#ports" title="Ports">Ports</a>: <i>
      - <a href="networkport.md">NetworkPort</a></i>
<a href="#peers" title="Peers">Peers</a>: <i>
      - <a href="networkpeer.md">NetworkPeer</a></i>
<a href="#allowedapps" title="AllowedApps">AllowedApps</a>: <i>
      - String</i>
<a href="#allowedpools" title="AllowedPools">AllowedPools</a>: <i>
      - String</i>
</pre>

## Properties

#### ID

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Enabled

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Description

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Ports

_Required_: No

_Type_: List of <a href="networkport.md">NetworkPort</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Peers

_Required_: No

_Type_: List of <a href="networkpeer.md">NetworkPeer</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### AllowedApps

_Required_: No

_Type_: List of String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### AllowedPools

_Required_: No

_Type_: List of String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

