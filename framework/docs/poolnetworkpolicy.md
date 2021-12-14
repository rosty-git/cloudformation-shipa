# Shipa::Framework::Item PoolNetworkPolicy

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#ingress" title="Ingress">Ingress</a>" : <i><a href="networkpolicyconfig.md">NetworkPolicyConfig</a></i>,
    "<a href="#egress" title="Egress">Egress</a>" : <i><a href="networkpolicyconfig.md">NetworkPolicyConfig</a></i>,
    "<a href="#disableapppolicies" title="DisableAppPolicies">DisableAppPolicies</a>" : <i>Boolean</i>
}
</pre>

### YAML

<pre>
<a href="#ingress" title="Ingress">Ingress</a>: <i><a href="networkpolicyconfig.md">NetworkPolicyConfig</a></i>
<a href="#egress" title="Egress">Egress</a>: <i><a href="networkpolicyconfig.md">NetworkPolicyConfig</a></i>
<a href="#disableapppolicies" title="DisableAppPolicies">DisableAppPolicies</a>: <i>Boolean</i>
</pre>

## Properties

#### Ingress

_Required_: No

_Type_: <a href="networkpolicyconfig.md">NetworkPolicyConfig</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Egress

_Required_: No

_Type_: <a href="networkpolicyconfig.md">NetworkPolicyConfig</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### DisableAppPolicies

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

