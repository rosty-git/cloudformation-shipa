# Shipa::Framework::Item

Congratulations on starting development!

Next steps:

1. Populate the JSON schema describing your resource, `shipa-framework-item.json`
2. The RPDK will automatically generate the correct resource model from the
   schema whenever the project is built via Make.
   You can also do this manually with the following command: `cfn-cli generate`
3. Implement your resource handlers by adding code to provision your resources in your resource handler's methods.

Please don't modify files `model.go and main.go`, as they will be automatically overwritten.


# development

First need to activate python virtual env
```bash
   source ../venv/bin/activate
```

To re-generate code after updating JSON schema
```bash
  cfn generate
```
   
Submit

```bash
$ cfn submit -v --region eu-west-1
```

List versions
```bash
$ aws cloudformation list-type-versions --type "RESOURCE" --type-name "Shipa::Framework::Item"
```


Set default version

```bash
$ aws cloudformation set-type-default-version \
  --type "RESOURCE" \
  --type-name "Shipa::Framework::Item" \
  --version-id "00000004"
```

Describe registration type
```bash
$ aws cloudformation describe-type-registration \
--registration-token a8ecadb8-3dcb-4e68-be9c-dbc85f375e57
```


Provision the resource in a CloudFormation stack (FAILING)
```bash
$ aws cloudformation create-stack --region eu-west-1 \
  --template-body "file://stack.json" \
  --stack-name "shipa-2"
  
  
An error occurred (ValidationError) when calling the CreateStack operation: Template format error: Unrecognized resource types: [Shipa::Framework::Item]

UPD:
{
    "StackId": "arn:aws:cloudformation:eu-west-1:859488020228:stack/shipa/2ff12150-571e-11ec-bb72-0250966f7939"
}

{
    "StackId": "arn:aws:cloudformation:eu-west-1:859488020228:stack/shipa/74091300-5798-11ec-9a97-06eb9983c2a9"
}


shipa-2
{
    "StackId": "arn:aws:cloudformation:eu-west-1:859488020228:stack/shipa-2/eecef250-579b-11ec-a04b-0a7ad719d0d7"
}

```

```bash
$ aws cloudformation describe-type --type-name "Shipa::Framework::Item" --type RESOURCE
```


Describe stack

```bash
$ aws cloudformation describe-stacks --region eu-west-1
```

List stack events
```bash
$ aws cloudformation describe-stack-events --stack-name shipa > events.log
```

{
"StackId": "arn:aws:cloudformation:eu-west-1:859488020228:stack/shipa-1/4ffa3320-579a-11ec-a32d-06f40cc54a2d"
}
