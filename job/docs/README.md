# Shipa::Job::Item

An example resource schema demonstrating some basic constructs and validation rules.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "Shipa::Job::Item",
    "Properties" : {
        "<a href="#name" title="Name">Name</a>" : <i>String</i>,
        "<a href="#framework" title="Framework">Framework</a>" : <i>String</i>,
        "<a href="#containers" title="Containers">Containers</a>" : <i>[ <a href="jobcontainer.md">JobContainer</a>, ... ]</i>,
        "<a href="#policy" title="Policy">Policy</a>" : <i><a href="policy.md">Policy</a></i>,
        "<a href="#backofflimit" title="BackoffLimit">BackoffLimit</a>" : <i>Integer</i>,
        "<a href="#completions" title="Completions">Completions</a>" : <i>Integer</i>,
        "<a href="#parallelism" title="Parallelism">Parallelism</a>" : <i>Integer</i>,
        "<a href="#suspend" title="Suspend">Suspend</a>" : <i>Boolean</i>,
        "<a href="#description" title="Description">Description</a>" : <i>String</i>,
        "<a href="#team" title="Team">Team</a>" : <i>String</i>,
        "<a href="#type" title="Type">Type</a>" : <i>String</i>,
        "<a href="#version" title="Version">Version</a>" : <i>String</i>,
        "<a href="#shipahost" title="ShipaHost">ShipaHost</a>" : <i>String</i>,
        "<a href="#shipatoken" title="ShipaToken">ShipaToken</a>" : <i>String</i>
    }
}
</pre>

### YAML

<pre>
Type: Shipa::Job::Item
Properties:
    <a href="#name" title="Name">Name</a>: <i>String</i>
    <a href="#framework" title="Framework">Framework</a>: <i>String</i>
    <a href="#containers" title="Containers">Containers</a>: <i>
      - <a href="jobcontainer.md">JobContainer</a></i>
    <a href="#policy" title="Policy">Policy</a>: <i><a href="policy.md">Policy</a></i>
    <a href="#backofflimit" title="BackoffLimit">BackoffLimit</a>: <i>Integer</i>
    <a href="#completions" title="Completions">Completions</a>: <i>Integer</i>
    <a href="#parallelism" title="Parallelism">Parallelism</a>: <i>Integer</i>
    <a href="#suspend" title="Suspend">Suspend</a>: <i>Boolean</i>
    <a href="#description" title="Description">Description</a>: <i>String</i>
    <a href="#team" title="Team">Team</a>: <i>String</i>
    <a href="#type" title="Type">Type</a>: <i>String</i>
    <a href="#version" title="Version">Version</a>: <i>String</i>
    <a href="#shipahost" title="ShipaHost">ShipaHost</a>: <i>String</i>
    <a href="#shipatoken" title="ShipaToken">ShipaToken</a>: <i>String</i>
</pre>

## Properties

#### Name

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### Framework

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Containers

_Required_: Yes

_Type_: List of <a href="jobcontainer.md">JobContainer</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Policy

_Required_: Yes

_Type_: <a href="policy.md">Policy</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### BackoffLimit

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Completions

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Parallelism

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Suspend

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Description

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Team

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Type

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Version

_Required_: No

_Type_: String

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
