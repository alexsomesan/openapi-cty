# openapi-cty
A Go library for synthesizing [github.com/zclconf/go-cty](https://github.com/zclconf/go-cty) types types from an OpenAPI document

**WORK IN PROGRESS**

## How to use it

You need an OpenAPI v2 spec file to begin with. You can either load it from local storage or retrieve it from some remote location.
```golang
file := filepath.Join("path", "to", "swagger.json")

input, err := ioutil.ReadFile(sfile)

if err != nil {
  return fmt.Errorf("failed to load definition file: %s : %s", file, err)
}
```
Once you've done that, you can create the type foundry this library offers using the following:
```golang
f, err := foundry.NewFoundryFromSpecV2(input)

if err != nil {
  return fmt.Errorf("failed to initialize type foundry", err)
}
```
You can now synthesize types using the `GetTypeByID()` function of the Foundry interface. `ID` is the name by which the type is identified in the *Definitions* section of the OpenAPI spec document.
```golang
id := "io.k8s.api.apps.v1.Deployment"

t, err := f.GetTypeByID(id)

if err != nil {
  return fmt.Errorf("failed to get type %s : %s", id, err)
}
```
