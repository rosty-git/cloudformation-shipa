# Shipa::AppDeploy::Item AppConfig

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#team" title="Team">Team</a>" : <i>String</i>,
    "<a href="#framework" title="Framework">Framework</a>" : <i>String</i>,
    "<a href="#description" title="Description">Description</a>" : <i>String</i>,
    "<a href="#env" title="Env">Env</a>" : <i>[ String, ... ]</i>,
    "<a href="#plan" title="Plan">Plan</a>" : <i>String</i>,
    "<a href="#router" title="Router">Router</a>" : <i>String</i>,
    "<a href="#tags" title="Tags">Tags</a>" : <i>[ String, ... ]</i>
}
</pre>

### YAML

<pre>
<a href="#team" title="Team">Team</a>: <i>String</i>
<a href="#framework" title="Framework">Framework</a>: <i>String</i>
<a href="#description" title="Description">Description</a>: <i>String</i>
<a href="#env" title="Env">Env</a>: <i>
      - String</i>
<a href="#plan" title="Plan">Plan</a>: <i>String</i>
<a href="#router" title="Router">Router</a>: <i>String</i>
<a href="#tags" title="Tags">Tags</a>: <i>
      - String</i>
</pre>

## Properties

#### Team

The team name

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Framework

The framework name

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Description

The app description

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Env

The app env vars

_Required_: No

_Type_: List of String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Plan

The app plan

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Router

The app router

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Tags

The app tags

_Required_: No

_Type_: List of String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

