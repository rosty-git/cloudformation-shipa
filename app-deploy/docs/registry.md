# Shipa::AppDeploy::Item Registry

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#user" title="User">User</a>" : <i>String</i>,
    "<a href="#secret" title="Secret">Secret</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#user" title="User">User</a>: <i>String</i>
<a href="#secret" title="Secret">Secret</a>: <i>String</i>
</pre>

## Properties

#### User

Registry user name

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Secret

Registry secret

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

