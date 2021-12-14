# Shipa::Cluster::Item ClusterResources

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#frameworks" title="Frameworks">Frameworks</a>" : <i>[ <a href="framework.md">Framework</a>, ... ]</i>,
    "<a href="#ingresscontrollers" title="IngressControllers">IngressControllers</a>" : <i>[ <a href="ingresscontroller.md">IngressController</a>, ... ]</i>
}
</pre>

### YAML

<pre>
<a href="#frameworks" title="Frameworks">Frameworks</a>: <i>
      - <a href="framework.md">Framework</a></i>
<a href="#ingresscontrollers" title="IngressControllers">IngressControllers</a>: <i>
      - <a href="ingresscontroller.md">IngressController</a></i>
</pre>

## Properties

#### Frameworks

_Required_: No

_Type_: List of <a href="framework.md">Framework</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### IngressControllers

_Required_: No

_Type_: List of <a href="ingresscontroller.md">IngressController</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

