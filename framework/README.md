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
   
