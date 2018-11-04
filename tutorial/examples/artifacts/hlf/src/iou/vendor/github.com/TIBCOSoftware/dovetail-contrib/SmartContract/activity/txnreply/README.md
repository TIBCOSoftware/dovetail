---
title: Txnreply
weight: 4603
---

# Txnreply
This activity sends response to transaction initiator, reponse can be SUCCESS, SUCCESS WITH DATA or Erroor with Message.

## Schema
Inputs and Outputs:

```json
{
  "inputs": [
        {
            "name": "status",
            "type": "string",
            "required": true,
            "allowed": ["Success", "Success with Data", "Error with Message"]
        },
        {
            "name": "message",
            "type": "string",
            "required": false
        },
        {
            "name": "model",
            "type": "string",
            "required": false
        },
        {
            "name": "dataType",
            "type": "string",
            "required":false
        },
        {
            "name": "isArray",
            "type": "boolean"
        }
        {
         "name": "input",
         "type": "complex_object"
     },
     {
         "name": "userInput",
         "type": "complex_object"
     }
    ],
  
    "outputs": []
}
```

## Settings
| Setting             | Required | Description |
|:--------------------|:---------|:------------|
| model    | False    | Select the common data model, must be the same as the one selected in Trigger |
| dataType | True     | Asset types defined in common data model if a model is selected. "User Defined..." data type allows user defined json schema |
| isArray  | True     | True if input is an array  |
| status   | True     | This is the handler to underlying blockchain service, should always be mapped to $flow.containerServiceStub |
| message  | False    | Error message to return |

## Input Schema
The json schema is automatically created based on settings

## Ouput Schema
The json schema is automatically created based on settings


