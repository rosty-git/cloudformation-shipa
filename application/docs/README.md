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
        "<a href="#description" title="Description">Description</a>" : <i>String</i>,
        "<a href="#teamowner" title="Teamowner">Teamowner</a>" : <i>String</i>,
        "<a href="#framework" title="Framework">Framework</a>" : <i>String</i>,
        "<a href="#plan" title="Plan">Plan</a>" : <i><a href="plan.md">Plan</a></i>,
        "<a href="#units" title="Units">Units</a>" : <i>[ <a href="unit.md">Unit</a>, ... ]</i>,
        "<a href="#cname" title="Cname">Cname</a>" : <i>[ String, ... ]</i>,
        "<a href="#ip" title="IP">IP</a>" : <i>String</i>,
        "<a href="#org" title="Org">Org</a>" : <i>String</i>,
        "<a href="#entrypoints" title="Entrypoints">Entrypoints</a>" : <i>[ <a href="entrypoint.md">Entrypoint</a>, ... ]</i>,
        "<a href="#routers" title="Routers">Routers</a>" : <i>[ <a href="router.md">Router</a>, ... ]</i>,
        "<a href="#lock" title="Lock">Lock</a>" : <i><a href="lock.md">Lock</a></i>,
        "<a href="#tags" title="Tags">Tags</a>" : <i>[ String, ... ]</i>,
        "<a href="#platform" title="Platform">Platform</a>" : <i>String</i>,
        "<a href="#status" title="Status">Status</a>" : <i>String</i>,
        "<a href="#error" title="Error">Error</a>" : <i>String</i>,
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
    <a href="#description" title="Description">Description</a>: <i>String</i>
    <a href="#teamowner" title="Teamowner">Teamowner</a>: <i>String</i>
    <a href="#framework" title="Framework">Framework</a>: <i>String</i>
    <a href="#plan" title="Plan">Plan</a>: <i><a href="plan.md">Plan</a></i>
    <a href="#units" title="Units">Units</a>: <i>
      - <a href="unit.md">Unit</a></i>
    <a href="#cname" title="Cname">Cname</a>: <i>
      - String</i>
    <a href="#ip" title="IP">IP</a>: <i>String</i>
    <a href="#org" title="Org">Org</a>: <i>String</i>
    <a href="#entrypoints" title="Entrypoints">Entrypoints</a>: <i>
      - <a href="entrypoint.md">Entrypoint</a></i>
    <a href="#routers" title="Routers">Routers</a>: <i>
      - <a href="router.md">Router</a></i>
    <a href="#lock" title="Lock">Lock</a>: <i><a href="lock.md">Lock</a></i>
    <a href="#tags" title="Tags">Tags</a>: <i>
      - String</i>
    <a href="#platform" title="Platform">Platform</a>: <i>String</i>
    <a href="#status" title="Status">Status</a>: <i>String</i>
    <a href="#error" title="Error">Error</a>: <i>String</i>
    <a href="#shipahost" title="ShipaHost">ShipaHost</a>: <i>String</i>
    <a href="#shipatoken" title="ShipaToken">ShipaToken</a>: <i>String</i>
</pre>

## Properties

#### Name

Shipa application name

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### Description

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

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

_Required_: No

_Type_: <a href="plan.md">Plan</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Units

_Required_: No

_Type_: List of <a href="unit.md">Unit</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Cname

_Required_: No

_Type_: List of String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### IP

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Org

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Entrypoints

_Required_: No

_Type_: List of <a href="entrypoint.md">Entrypoint</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Routers

_Required_: No

_Type_: List of <a href="router.md">Router</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Lock

_Required_: No

_Type_: <a href="lock.md">Lock</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Tags

_Required_: No

_Type_: List of String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Platform

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Status

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Error

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
