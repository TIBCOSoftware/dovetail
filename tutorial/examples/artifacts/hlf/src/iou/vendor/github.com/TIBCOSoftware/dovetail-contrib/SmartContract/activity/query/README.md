---
title: Hyperledger Fabric Query
weight: 4603
---

# Query
This activity supports Hyperledger Fabric CouchDB query

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
                "type": "string",
                "required": true
           },
           {
            "name": "params",
            "type": "complex_object",
            "required": false,
            "display":{
                "name": "Define parameters used in query string",
                "type": "params",
                "schema": "{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"paramName\":{\"type\":\"string\"},\"type\":{\"type\":{\"enum\":[\"string\",\"number\",\"boolean\"]}}}}}"
            }
       },
           {
            "name": "queryString",
            "type": "string",
            "required": true
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
            "required": true
        }
    ]
}
```

## Settings
| Setting              | Required | Description |
|:---------------------|:---------|:------------|
| model                | True    | Select the common data model, must be the same as the one selected in Trigger |
| assetName            | True    | Select the asset to query|
| params               | False   | Define input parameters |
| queryString          | True    | CouchDB query string, use _$paramName (paramName is defined in params) for input parameters |
| containerServiceStub | True    | This is the handler to underlying blockchain service, should always be mapped to $flow.containerServiceStub |

## Input Schema
The json schema is automatically created

## Ouput Schema
The json schema is automatically created 


