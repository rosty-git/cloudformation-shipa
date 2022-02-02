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
    "<a href="#containerpolicy" title="ContainerPolicy">ContainerPolicy</a>" : <i><a href="poolcontainerpolicy.md">PoolContainerPolicy</a></i>,
    "<a href="#nodeselectors" title="NodeSelectors">NodeSelectors</a>" : <i><a href="nodeselectors.md">NodeSelectors</a></i>,
    "<a href="#podautoscaler" title="PodAutoScaler">PodAutoScaler</a>" : <i><a href="podautoscaler.md">PodAutoScaler</a></i>,
    "<a href="#domainpolicy" title="DomainPolicy">DomainPolicy</a>" : <i><a href="domainpolicy.md">DomainPolicy</a></i>,
    "<a href="#appautodiscovery" title="AppAutoDiscovery">AppAutoDiscovery</a>" : <i><a href="appautodiscovery.md">AppAutoDiscovery</a></i>,
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
<a href="#containerpolicy" title="ContainerPolicy">ContainerPolicy</a>: <i><a href="poolcontainerpolicy.md">PoolContainerPolicy</a></i>
<a href="#nodeselectors" title="NodeSelectors">NodeSelectors</a>: <i><a href="nodeselectors.md">NodeSelectors</a></i>
<a href="#podautoscaler" title="PodAutoScaler">PodAutoScaler</a>: <i><a href="podautoscaler.md">PodAutoScaler</a></i>
<a href="#domainpolicy" title="DomainPolicy">DomainPolicy</a>: <i><a href="domainpolicy.md">DomainPolicy</a></i>
<a href="#appautodiscovery" title="AppAutoDiscovery">AppAutoDiscovery</a>: <i><a href="appautodiscovery.md">AppAutoDiscovery</a></i>
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

#### ContainerPolicy

_Required_: No

_Type_: <a href="poolcontainerpolicy.md">PoolContainerPolicy</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### NodeSelectors

_Required_: No

_Type_: <a href="nodeselectors.md">NodeSelectors</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### PodAutoScaler

_Required_: No

_Type_: <a href="podautoscaler.md">PodAutoScaler</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### DomainPolicy

_Required_: No

_Type_: <a href="domainpolicy.md">DomainPolicy</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### AppAutoDiscovery

_Required_: No

_Type_: <a href="appautodiscovery.md">AppAutoDiscovery</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### NetworkPolicy

_Required_: No

_Type_: <a href="poolnetworkpolicy.md">PoolNetworkPolicy</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

