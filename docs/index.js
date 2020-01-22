const data = require('../catalog.json');
const table = require('markdown-table');

function tableByType(type) {
    let t = [['Image', 'Args', 'Description', 'Source']];
    data.filter(r => r.type === type).forEach(r => {
        let desc = r.description;
        if (r.demo === true) {
            desc = '[Demo] ' + desc;
        }
        const source = `[Link](${r.source})`;
        t.push([r.image, r.args, desc, source]);
    });
    return table(t);
}

const README = `
<!---
DO NOT EDIT. Generated by: "cd docs; npm run gen-docs"
-->

# KPT Functions Catalog

This repository documents a catalog of functions implementing [Configuration Functions Specification][spec].

These functions can be implemented using any toolchain such as the [KPT Functions SDK][sdk].

## Sources

See [definition of source functions][source].

${tableByType('source')}

## Sinks

See [definition of sink functions][sink].

${tableByType('sink')}

## Validators

${tableByType('validator')}

## Generators

${tableByType('generator')}

## Transformers

${tableByType('transformer')}

## Miscellaneous

${tableByType('misc')}

[spec]: https://github.com/kubernetes-sigs/kustomize/blob/master/cmd/config/docs/api-conventions/functions-spec.md
[source]: https://github.com/GoogleContainerTools/kpt-functions-sdk/blob/master/docs/concepts.md#source-function
[sink]: https://github.com/GoogleContainerTools/kpt-functions-sdk/blob/master/docs/concepts.md#sink-function
[sdk]: https://github.com/GoogleContainerTools/kpt-functions-sdk
`;

console.log(README);