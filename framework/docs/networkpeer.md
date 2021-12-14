# Shipa::Framework::Item NetworkPeer

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#podselector" title="PodSelector">PodSelector</a>" : <i><a href="networkpeerselector.md">NetworkPeerSelector</a></i>,
    "<a href="#namespaceselector" title="NamespaceSelector">NamespaceSelector</a>" : <i><a href="networkpeerselector.md">NetworkPeerSelector</a></i>,
    "<a href="#ipblock" title="IPBlock">IPBlock</a>" : <i>[ String, ... ]</i>
}
</pre>

### YAML

<pre>
<a href="#podselector" title="PodSelector">PodSelector</a>: <i><a href="networkpeerselector.md">NetworkPeerSelector</a></i>
<a href="#namespaceselector" title="NamespaceSelector">NamespaceSelector</a>: <i><a href="networkpeerselector.md">NetworkPeerSelector</a></i>
<a href="#ipblock" title="IPBlock">IPBlock</a>: <i>
      - String</i>
</pre>

## Properties

#### PodSelector

_Required_: No

_Type_: <a href="networkpeerselector.md">NetworkPeerSelector</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### NamespaceSelector

_Required_: No

_Type_: <a href="networkpeerselector.md">NetworkPeerSelector</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### IPBlock

_Required_: No

_Type_: List of String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

