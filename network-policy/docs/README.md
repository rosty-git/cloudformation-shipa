# Shipa::NetworkPolicy::Item

An example resource schema demonstrating some basic constructs and validation rules.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "Shipa::NetworkPolicy::Item",
    "Properties" : {
        "<a href="#app" title="App">App</a>" : <i>String</i>,
        "<a href="#ingress" title="Ingress">Ingress</a>" : <i><a href="networkpolicyconfig.md">NetworkPolicyConfig</a></i>,
        "<a href="#egress" title="Egress">Egress</a>" : <i><a href="networkpolicyconfig.md">NetworkPolicyConfig</a></i>,
        "<a href="#restartapp" title="RestartApp">RestartApp</a>" : <i>Boolean</i>,
        "<a href="#shipahost" title="ShipaHost">ShipaHost</a>" : <i>String</i>,
        "<a href="#shipatoken" title="ShipaToken">ShipaToken</a>" : <i>String</i>
    }
}
</pre>

### YAML

<pre>
Type: Shipa::NetworkPolicy::Item
Properties:
    <a href="#app" title="App">App</a>: <i>String</i>
    <a href="#ingress" title="Ingress">Ingress</a>: <i><a href="networkpolicyconfig.md">NetworkPolicyConfig</a></i>
    <a href="#egress" title="Egress">Egress</a>: <i><a href="networkpolicyconfig.md">NetworkPolicyConfig</a></i>
    <a href="#restartapp" title="RestartApp">RestartApp</a>: <i>Boolean</i>
    <a href="#shipahost" title="ShipaHost">ShipaHost</a>: <i>String</i>
    <a href="#shipatoken" title="ShipaToken">ShipaToken</a>: <i>String</i>
</pre>

## Properties

#### App

Shipa application name

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### Ingress

_Required_: No

_Type_: <a href="networkpolicyconfig.md">NetworkPolicyConfig</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Egress

_Required_: No

_Type_: <a href="networkpolicyconfig.md">NetworkPolicyConfig</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### RestartApp

Application restart flag

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ShipaHost

Shipa server host

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ShipaToken

Shipa access token

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, Ref returns the App.
