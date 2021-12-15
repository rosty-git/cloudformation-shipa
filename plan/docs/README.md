# Shipa::Plan::Item

An example resource schema demonstrating some basic constructs and validation rules.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "Shipa::Plan::Item",
    "Properties" : {
        "<a href="#name" title="Name">Name</a>" : <i>String</i>,
        "<a href="#memory" title="Memory">Memory</a>" : <i>String</i>,
        "<a href="#swap" title="Swap">Swap</a>" : <i>String</i>,
        "<a href="#cpushare" title="CPUShare">CPUShare</a>" : <i>Integer</i>,
        "<a href="#default" title="Default">Default</a>" : <i>Boolean</i>,
        "<a href="#public" title="Public">Public</a>" : <i>Boolean</i>,
        "<a href="#org" title="Org">Org</a>" : <i>String</i>,
        "<a href="#teams" title="Teams">Teams</a>" : <i>[ String, ... ]</i>,
        "<a href="#shipahost" title="ShipaHost">ShipaHost</a>" : <i>String</i>,
        "<a href="#shipatoken" title="ShipaToken">ShipaToken</a>" : <i>String</i>
    }
}
</pre>

### YAML

<pre>
Type: Shipa::Plan::Item
Properties:
    <a href="#name" title="Name">Name</a>: <i>String</i>
    <a href="#memory" title="Memory">Memory</a>: <i>String</i>
    <a href="#swap" title="Swap">Swap</a>: <i>String</i>
    <a href="#cpushare" title="CPUShare">CPUShare</a>: <i>Integer</i>
    <a href="#default" title="Default">Default</a>: <i>Boolean</i>
    <a href="#public" title="Public">Public</a>: <i>Boolean</i>
    <a href="#org" title="Org">Org</a>: <i>String</i>
    <a href="#teams" title="Teams">Teams</a>: <i>
      - String</i>
    <a href="#shipahost" title="ShipaHost">ShipaHost</a>: <i>String</i>
    <a href="#shipatoken" title="ShipaToken">ShipaToken</a>: <i>String</i>
</pre>

## Properties

#### Name

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### Memory

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Swap

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### CPUShare

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Default

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Public

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Org

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Teams

_Required_: No

_Type_: List of String

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

When you pass the logical ID of this resource to the intrinsic `Ref` function, Ref returns the Name.
