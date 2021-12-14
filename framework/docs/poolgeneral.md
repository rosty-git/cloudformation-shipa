# Shipa::Framework::Item PoolGeneral

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#setup" title="Setup">Setup</a>" : <i><a href="poolsetup.md">PoolSetup</a></i>,
    "<a href="#plan" title="Plan">Plan</a>" : <i><a href="poolplan.md">PoolPlan</a></i>,
    "<a href="#security" title="Security">Security</a>" : <i><a href="poolsecurity.md">PoolSecurity</a></i>,
    "<a href="#access" title="Access">Access</a>" : <i><a href="poolserviceaccess.md">PoolServiceAccess</a></i>,
    "<a href="#services" title="Services">Services</a>" : <i><a href="poolserviceaccess.md">PoolServiceAccess</a></i>,
    "<a href="#router" title="Router">Router</a>" : <i>String</i>,
    "<a href="#volumes" title="Volumes">Volumes</a>" : <i>[ String, ... ]</i>,
    "<a href="#appquota" title="AppQuota">AppQuota</a>" : <i><a href="poolappquota.md">PoolAppQuota</a></i>,
    "<a href="#containerpolicy" title="ContainerPolicy">ContainerPolicy</a>" : <i><a href="poolcontainerpolicy.md">PoolContainerPolicy</a></i>,
    "<a href="#networkpolicy" title="NetworkPolicy">NetworkPolicy</a>" : <i><a href="poolnetworkpolicy.md">PoolNetworkPolicy</a></i>
}
</pre>

### YAML

<pre>
<a href="#setup" title="Setup">Setup</a>: <i><a href="poolsetup.md">PoolSetup</a></i>
<a href="#plan" title="Plan">Plan</a>: <i><a href="poolplan.md">PoolPlan</a></i>
<a href="#security" title="Security">Security</a>: <i><a href="poolsecurity.md">PoolSecurity</a></i>
<a href="#access" title="Access">Access</a>: <i><a href="poolserviceaccess.md">PoolServiceAccess</a></i>
<a href="#services" title="Services">Services</a>: <i><a href="poolserviceaccess.md">PoolServiceAccess</a></i>
<a href="#router" title="Router">Router</a>: <i>String</i>
<a href="#volumes" title="Volumes">Volumes</a>: <i>
      - String</i>
<a href="#appquota" title="AppQuota">AppQuota</a>: <i><a href="poolappquota.md">PoolAppQuota</a></i>
<a href="#containerpolicy" title="ContainerPolicy">ContainerPolicy</a>: <i><a href="poolcontainerpolicy.md">PoolContainerPolicy</a></i>
<a href="#networkpolicy" title="NetworkPolicy">NetworkPolicy</a>: <i><a href="poolnetworkpolicy.md">PoolNetworkPolicy</a></i>
</pre>

## Properties

#### Setup

_Required_: No

_Type_: <a href="poolsetup.md">PoolSetup</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Plan

_Required_: No

_Type_: <a href="poolplan.md">PoolPlan</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Security

_Required_: No

_Type_: <a href="poolsecurity.md">PoolSecurity</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Access

_Required_: No

_Type_: <a href="poolserviceaccess.md">PoolServiceAccess</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Services

_Required_: No

_Type_: <a href="poolserviceaccess.md">PoolServiceAccess</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Router

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Volumes

_Required_: No

_Type_: List of String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### AppQuota

_Required_: No

_Type_: <a href="poolappquota.md">PoolAppQuota</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ContainerPolicy

_Required_: No

_Type_: <a href="poolcontainerpolicy.md">PoolContainerPolicy</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### NetworkPolicy

_Required_: No

_Type_: <a href="poolnetworkpolicy.md">PoolNetworkPolicy</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

