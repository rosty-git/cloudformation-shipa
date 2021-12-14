# Shipa::Framework::Item NetworkPolicyConfig

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#policymode" title="PolicyMode">PolicyMode</a>" : <i>String</i>,
    "<a href="#customrules" title="CustomRules">CustomRules</a>" : <i>[ <a href="networkpolicyrule.md">NetworkPolicyRule</a>, ... ]</i>,
    "<a href="#shiparules" title="ShipaRules">ShipaRules</a>" : <i>[ <a href="networkpolicyrule.md">NetworkPolicyRule</a>, ... ]</i>,
    "<a href="#shiparulesenabled" title="ShipaRulesEnabled">ShipaRulesEnabled</a>" : <i>[ String, ... ]</i>
}
</pre>

### YAML

<pre>
<a href="#policymode" title="PolicyMode">PolicyMode</a>: <i>String</i>
<a href="#customrules" title="CustomRules">CustomRules</a>: <i>
      - <a href="networkpolicyrule.md">NetworkPolicyRule</a></i>
<a href="#shiparules" title="ShipaRules">ShipaRules</a>: <i>
      - <a href="networkpolicyrule.md">NetworkPolicyRule</a></i>
<a href="#shiparulesenabled" title="ShipaRulesEnabled">ShipaRulesEnabled</a>: <i>
      - String</i>
</pre>

## Properties

#### PolicyMode

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### CustomRules

_Required_: No

_Type_: List of <a href="networkpolicyrule.md">NetworkPolicyRule</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ShipaRules

_Required_: No

_Type_: List of <a href="networkpolicyrule.md">NetworkPolicyRule</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ShipaRulesEnabled

_Required_: No

_Type_: List of String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

