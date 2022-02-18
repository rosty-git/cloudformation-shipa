# Shipa::Application::Item

An example resource schema demonstrating some basic constructs and validation rules.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "Shipa::Application::Item",
    "Properties" : {
        "<a href="#name" title="Name">Name</a>" : <i>String</i>,
        "<a href="#teamowner" title="Teamowner">Teamowner</a>" : <i>String</i>,
        "<a href="#framework" title="Framework">Framework</a>" : <i>String</i>,
        "<a href="#plan" title="Plan">Plan</a>" : <i>String</i>,
        "<a href="#tags" title="Tags">Tags</a>" : <i>[ String, ... ]</i>,
        "<a href="#shipahost" title="ShipaHost">ShipaHost</a>" : <i>String</i>,
        "<a href="#shipatoken" title="ShipaToken">ShipaToken</a>" : <i>String</i>
    }
}
</pre>

### YAML

<pre>
Type: Shipa::Application::Item
Properties:
    <a href="#name" title="Name">Name</a>: <i>String</i>
    <a href="#teamowner" title="Teamowner">Teamowner</a>: <i>String</i>
    <a href="#framework" title="Framework">Framework</a>: <i>String</i>
    <a href="#plan" title="Plan">Plan</a>: <i>String</i>
    <a href="#tags" title="Tags">Tags</a>: <i>
      - String</i>
    <a href="#shipahost" title="ShipaHost">ShipaHost</a>: <i>String</i>
    <a href="#shipatoken" title="ShipaToken">ShipaToken</a>: <i>String</i>
</pre>

## Properties

#### Name

Shipa application name

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### Teamowner

Shipa application teamowner

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Framework

Shipa application framework

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Plan

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Tags

_Required_: Yes

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
