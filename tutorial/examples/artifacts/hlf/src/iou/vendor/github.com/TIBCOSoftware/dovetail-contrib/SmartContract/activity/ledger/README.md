---
title: Ledger
weight: 4603
---

# Ledger
This activity reads and writes to a ledger, supported operations: "PUT", "DELETE", "GET","LOOKUP"

## Schema
Inputs and Outputs:

```json
{
  "inputs": [
            {
                "name": "model",
                "type": "string"
            },
           {
                "name": "assetName",
                "type": "string"
           },
           {
            "name": "operation",
            "type": "string",
            "allowed": ["PUT", "DELETE", "GET","LOOKUP"]
           },
           {
            "name": "isArray",
            "type": "boolean"
           },
           {
            "name": "compositeKey",
            "type": "string"
           },
           {
            "name": "input",
            "type": "complex_object",
            "required": false
           },
           {
            "name": "containerServiceStub",
            "type": "any",
            "required":true
           }
    ],
  
    "outputs": [
        {
            "name": "output",
            "type": "complex_object",
            "required": false
        }
    ]
}
```

## Settings
| Setting              | Required | Description |
|:---------------------|:---------|:------------|
| model                | True    | Select the common data model, must be the same as the one selected in Trigger |
| assetName            | True    | Select the asset to be read or write |
| operation            | True    | Supported actions: "PUT", "DELETE", "GET","LOOKUP", LOOKUP is for partial key lookup when the primary key is composite key, e.g., for composite key1, key2, key3, you can lookup assets using key1, or key1 and key2 |
| isArray              | True    | True for batch operations  |
| containerServiceStub | True    | This is the handler to underlying blockchain service, should always be mapped to $flow.containerServiceStub |

## Input Schema
The json schema is automatically created based on asset and isArray values

## Ouput Schema
The json schema is automatically created based on asset and isArray values


