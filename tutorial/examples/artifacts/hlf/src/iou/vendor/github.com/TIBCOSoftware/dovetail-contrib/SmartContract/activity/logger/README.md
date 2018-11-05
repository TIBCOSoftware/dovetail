---
title: Logger
weight: 4603
---

# Logger
This activity sends logging messages to underlying blockchain log framework

## Schema
Inputs and Outputs:

```json
{
  "inputs": [
            {
            "name": "logLevel",
            "type": "string",
            "required": true,
            "allowed": ["DEBUG", "INFO", "WARNING", "ERROR"]
           },
           {
            "name": "message",
            "type": "string",
            "required": true
           },
           {
            "name": "errorCode",
            "type": "string",
            "required": false
           },
           {
            "name": "containerServiceStub",
            "type": "any",
            "required":true
           }
    ],
  
    "outputs": []
}
```

## Settings
| Setting                  | Required | Description |
|:-------------------------|:---------|:------------|
| logLevel                 | True     | Logging level |
| message                  | True     | Logging message |
| errorCode                | False    | Error code |
| containerServiceStub     | True     | This is the handler to underlying blockchain service, should always be mapped to $flow.containerServiceStub |



