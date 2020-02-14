---
title: "Create client app"
linkTitle: "Create client app"
weight: 9
description: >
  Instructions on how to create a client app into TCI (TIBCO Cloud Integration)
---

Before anything download the [fabric_hw_client.json](../../fabric_hw_client.json) model example.

* Inside your Develop Flogo environment on TCI select "Create" button to create a new app.
* Enter "fabric_hw_client" and select "create".
* Select "create a TIBCO Flogo App" option.
* Click on "Import App" option.
* Upload fabric_hw_client.json selecting "import all" option.


After you import fabric_hw_client.json you can see a very simple rest application with 2 transactions "get_request" and "set_request".

* get_request: Logs a message and then calls the smart contract fabric_hw to store a value for a given key.
* get_request: Logs a message and then calls the smart contract fabric_hw to return the value of a given key.



For more detail information on how to create apps and model apps to our [TCI docs](https://integration.cloud.tibco.com/docs/index.html) page.






