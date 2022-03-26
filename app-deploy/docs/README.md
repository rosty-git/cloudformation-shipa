# Shipa::AppDeploy::Item

An example resource schema demonstrating some basic constructs and validation rules.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "Shipa::AppDeploy::Item",
    "Properties" : {
        "<a href="#app" title="App">App</a>" : <i>String</i>,
        "<a href="#image" title="Image">Image</a>" : <i>String</i>,
        "<a href="#appconfig" title="AppConfig">AppConfig</a>" : <i><a href="appconfig.md">AppConfig</a></i>,
        "<a href="#canarysettings" title="CanarySettings">CanarySettings</a>" : <i><a href="canarysettings.md">CanarySettings</a></i>,
        "<a href="#podautoscaler" title="PodAutoScaler">PodAutoScaler</a>" : <i><a href="podautoscaler.md">PodAutoScaler</a></i>,
        "<a href="#port" title="Port">Port</a>" : <i><a href="port.md">Port</a></i>,
        "<a href="#registry" title="Registry">Registry</a>" : <i><a href="registry.md">Registry</a></i>,
        "<a href="#volumes" title="Volumes">Volumes</a>" : <i>[ <a href="volume.md">Volume</a>, ... ]</i>,
        "<a href="#shipahost" title="ShipaHost">ShipaHost</a>" : <i>String</i>,
        "<a href="#shipatoken" title="ShipaToken">ShipaToken</a>" : <i>String</i>
    }
}
</pre>

### YAML

<pre>
Type: Shipa::AppDeploy::Item
Properties:
    <a href="#app" title="App">App</a>: <i>String</i>
    <a href="#image" title="Image">Image</a>: <i>String</i>
    <a href="#appconfig" title="AppConfig">AppConfig</a>: <i><a href="appconfig.md">AppConfig</a></i>
    <a href="#canarysettings" title="CanarySettings">CanarySettings</a>: <i><a href="canarysettings.md">CanarySettings</a></i>
    <a href="#podautoscaler" title="PodAutoScaler">PodAutoScaler</a>: <i><a href="podautoscaler.md">PodAutoScaler</a></i>
    <a href="#port" title="Port">Port</a>: <i><a href="port.md">Port</a></i>
    <a href="#registry" title="Registry">Registry</a>: <i><a href="registry.md">Registry</a></i>
    <a href="#volumes" title="Volumes">Volumes</a>: <i>
      - <a href="volume.md">Volume</a></i>
    <a href="#shipahost" title="ShipaHost">ShipaHost</a>: <i>String</i>
    <a href="#shipatoken" title="ShipaToken">ShipaToken</a>: <i>String</i>
</pre>

## Properties

#### App

Shipa application name

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### Image

Application docker image

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### AppConfig

_Required_: No

_Type_: <a href="appconfig.md">AppConfig</a>

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### CanarySettings

_Required_: No

_Type_: <a href="canarysettings.md">CanarySettings</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### PodAutoScaler

_Required_: No

_Type_: <a href="podautoscaler.md">PodAutoScaler</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Port

_Required_: No

_Type_: <a href="port.md">Port</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Registry

_Required_: No

_Type_: <a href="registry.md">Registry</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Volumes

Volumes settings

_Required_: No

_Type_: List of <a href="volume.md">Volume</a>

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
