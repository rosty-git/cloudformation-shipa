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

Validating your resource specification...
Explicitly specify value for tagging
Resource schema is valid.
Validating your resource schema...
Packaging Go project
Creating CloudFormationManagedUploadInfrastructure
CloudFormationManagedUploadInfrastructure stack was successfully created
Successfully submitted type. Waiting for registration with token 'a8ecadb8-3dcb-4e68-be9c-dbc85f375e57' to complete.
Registration complete.
{'ProgressStatus': 'COMPLETE', 'Description': 'Deployment is currently in DEPLOY_STAGE of status COMPLETED; ', 'TypeArn': 'arn:aws:cloudformation:eu-west-1:859488020228:type/resource/Shipa-Framework-Item', 'TypeVersionArn': 'arn:aws:cloudformation:eu-west-1:859488020228:type/resource/Shipa-Framework-Item/00000001', 'ResponseMetadata': {'RequestId': 'a615398c-06af-4138-99b5-5a90ee5fdeb9', 'HTTPStatusCode': 200, 'HTTPHeaders': {'x-amzn-requestid': 'a615398c-06af-4138-99b5-5a90ee5fdeb9', 'content-type': 'text/xml', 'content-length': '679', 'date': 'Tue, 07 Dec 2021 05:16:44 GMT'}, 'RetryAttempts': 0}}
```

List versions
```bash
$ aws cloudformation list-type-versions \
  --type "RESOURCE" \
  --type-name "Shipa::Framework::Item"
  
{
    "TypeVersionSummaries": [
        {
            "Type": "RESOURCE",
            "TypeName": "Shipa::Framework::Item",
            "VersionId": "00000001",
            "IsDefaultVersion": true,
            "Arn": "arn:aws:cloudformation:eu-west-1:859488020228:type/resource/Shipa-Framework-Item/00000001",
            "TimeCreated": "2021-12-07T05:16:42.851000+00:00",
            "Description": "An example resource schema demonstrating some basic constructs and validation rules."
        }
    ]
}

```


Set default version

```bash
$ aws cloudformation set-type-default-version \
  --type "RESOURCE" \
  --type-name "Shipa::Framework::Item" \
  --version-id "00000002"
```

Describe registration type
```bash
$ aws cloudformation describe-type-registration \
--registration-token a8ecadb8-3dcb-4e68-be9c-dbc85f375e57

{
    "ProgressStatus": "COMPLETE",
    "Description": "Deployment is currently in DEPLOY_STAGE of status COMPLETED; ",
    "TypeArn": "arn:aws:cloudformation:eu-west-1:859488020228:type/resource/Shipa-Framework-Item",
    "TypeVersionArn": "arn:aws:cloudformation:eu-west-1:859488020228:type/resource/Shipa-Framework-Item/00000001"
}

```


Provision the resource in a CloudFormation stack (FAILING)
```bash
$ aws cloudformation create-stack --region eu-west-1 \
  --template-body "file://stack.json" \
  --stack-name "shipa"
  
  
An error occurred (ValidationError) when calling the CreateStack operation: Template format error: Unrecognized resource types: [Shipa::Framework::Item]

UPD:
{
    "StackId": "arn:aws:cloudformation:eu-west-1:859488020228:stack/shipa/2ff12150-571e-11ec-bb72-0250966f7939"
}


```

```bash
$ aws cloudformation describe-type --type-name "Shipa::Framework::Item" --type RESOURCE

{
"Arn": "arn:aws:cloudformation:eu-west-1:859488020228:type/resource/Shipa-Framework-Item/00000001",
"Type": "RESOURCE",
"TypeName": "Shipa::Framework::Item",
"DefaultVersionId": "00000001",
"IsDefaultVersion": true,
"Description": "An example resource schema demonstrating some basic constructs and validation rules.",
"Schema": "{\n    \"typeName\": \"Shipa::Framework::Item\",\n    \"description\": \"An example resource schema demonstrating some basic constructs and validation rules.\",\n    \"properties\": {\n        \"Name\": {\n            \"description\": \"Shipa framework name\",\n            \"type\": \"string\"\n        },\n        \"ShipaHost\": {\n            \"description\": \"Shipa server host\",\n            \"type\": \"string\"\n        },\n        \"ShipaToken\": {\n            \"description\": \"Shipa access token\",\n            \"type\": \"string\"\n        }\n    },\n    \"additionalProperties\": false,\n    \"createOnlyProperties\": [\n        \"/properties/Name\"\n    ],\n    \"primaryIdentifier\": [\n        \"/properties/Name\"\n    ],\n    \"required\": [\n        \"ShipaHost\",\n        \"ShipaToken\"\n    ]\n}\n",
"ProvisioningType": "NON_PROVISIONABLE",
"DeprecatedStatus": "LIVE",
"LoggingConfig": {
"LogRoleArn": "arn:aws:iam::859488020228:role/CloudFormationManagedUplo-LogAndMetricsDeliveryRol-F4BKOU761DJ5",
"LogGroupName": "shipa-framework-item-logs"
},
"Visibility": "PRIVATE",
"LastUpdated": "2021-12-07T05:16:42.851000+00:00",
"TimeCreated": "2021-12-07T05:16:42.851000+00:00"
}

```


Describe stack

```bash
$ aws cloudformation describe-stacks --region eu-west-1


 "Stacks": [
        {
            "StackId": "arn:aws:cloudformation:eu-west-1:859488020228:stack/shipa/2ff12150-571e-11ec-bb72-0250966f7939",
            "StackName": "shipa",
            "Description": "Shipa stack",
            "CreationTime": "2021-12-07T17:51:51.606000+00:00",
            "DeletionTime": "2021-12-07T17:51:59.382000+00:00",
            "RollbackConfiguration": {},
            "StackStatus": "ROLLBACK_COMPLETE",
            "DisableRollback": false,
            "NotificationARNs": [],
            "Tags": [],
            "DriftInformation": {
                "StackDriftStatus": "NOT_CHECKED"
            }
        },

```