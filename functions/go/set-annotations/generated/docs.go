

// Code generated by "mdtogo"; DO NOT EDIT.
package generated

var SetAnnotationsShort = `The ` + "`" + `set-annotations` + "`" + ` function adds a list of annotations to all resources.
Annotations are commonly used in Kubernetes for attaching arbitrary metadata to
KRM resources.

For example, annotations can be used in the following scenarios:

- Provide information for controllers. (e.g. gce ingress controller will only
  take actions on ` + "`" + `Ingress` + "`" + ` resources with
  annotation ` + "`" + `kubernetes.io/ingress.class: gce` + "`" + `)
- Tools store information for later use. (e.g. ` + "`" + `kubectl apply` + "`" + ` stores what a
  user applied previously in
  annotation ` + "`" + `kubectl.kubernetes.io/last-applied-configuration` + "`" + `)`
var SetAnnotationsLong = `
There are 2 kinds of ` + "`" + `functionConfig` + "`" + ` supported by this function:

- ` + "`" + `ConfigMap` + "`" + `
- A custom resource of kind ` + "`" + `SetAnnotations` + "`" + `

To use a ` + "`" + `ConfigMap` + "`" + ` as the ` + "`" + `functionConfig` + "`" + `, the desired annotations must be
specified in the ` + "`" + `data` + "`" + ` field.

To add 2 annotations ` + "`" + `color: orange` + "`" + ` and ` + "`" + `fruit: apple` + "`" + ` to all resources:

  apiVersion: v1
  kind: ConfigMap
  metadata:
    name: my-config
  data:
    color: orange
    fruit: apple

To use a ` + "`" + `SetAnnotations` + "`" + ` custom resource as the ` + "`" + `functionConfig` + "`" + `, the desired
annotations must be specified in the ` + "`" + `annotations` + "`" + ` field. Sometimes you have
resources (especially custom resources) that have annotations fields in fields
other than the [defaults][commonannotations], you can specify such annotations
fields using ` + "`" + `additionalAnnotationFields` + "`" + `. It will be used jointly with the
[defaults][commonannotations].

` + "`" + `additionalAnnotationFields` + "`" + ` has following fields:

- ` + "`" + `group` + "`" + `: Select the resources by API version group. Will select all groups if
  omitted.
- ` + "`" + `version` + "`" + `: Select the resources by API version. Will select all versions if
  omitted.
- ` + "`" + `kind` + "`" + `: Select the resources by resource kind. Will select all kinds if
  omitted.
- ` + "`" + `path` + "`" + `: Specify the path to the field that the value will be updated. This
  field is required.
- ` + "`" + `create` + "`" + `: If it's set to true, the field specified will be created if it
  doesn't exist. Otherwise, the function will only update the existing field.

To add 2 annotations ` + "`" + `color: orange` + "`" + ` and ` + "`" + `fruit: apple` + "`" + ` to all built-in
resources and the path ` + "`" + `data.selector.annotations` + "`" + ` in ` + "`" + `MyOwnKind` + "`" + ` resource, we
use the following ` + "`" + `functionConfig` + "`" + `:

  apiVersion: fn.kpt.dev/v1alpha1
  kind: SetAnnotations
  metadata:
    name: my-config
  annotations:
    color: orange
    fruit: apple
  additionalAnnotationFields:
    - path: data/selector/annotations
      kind: MyOwnKind
      create: true
`
