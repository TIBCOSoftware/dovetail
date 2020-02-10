---
title: "Create smart contract"
linkTitle: "Create smart contract"
weight: 7
description: >
  Instructions on how to create a hyperledger smart contract into TCI (TIBCO Cloud Integration)
---

Before anything download the [fabric_hw.json](../../fabric_hw.json) and [fabric_hw_client.json](../../fabric_hw_client.json) model example from our latest release.

* Inside your Develop Flogo environment on TCI select "Create" button to create a new app.
* Enter "fabric_hw" and select "create".
* Select "create a TIBCO Flogo App" option.
* Click on "Import App" option.
* Upload fabric_hw.json selecting "import all" option.


After you import fabric_hw.json you can see a very simple smart contract with 2 transactions "get_value" and "set_value".

* set_value: Logs a message "Inside get_value" and then sets a value for a given key.
* get_value: Logs a message "Setting value" and then returns the value of a given key.



For more detail information on how to create apps and model apps to our [TCI docs](https://integration.cloud.tibco.com/docs/index.html) page.






