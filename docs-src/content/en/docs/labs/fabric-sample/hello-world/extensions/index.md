---
title: "Load fabric extensions"
linkTitle: "Load extensions"
weight: 6
description: >
  Instructions on how to load hyperledger fabric extensions into TCI
---

Before we start modelling our fabric application, we need to load the fabric extensions into flogo TCI.

* First log into TIBCO Cloud Integration as indicated in our [Dovetail UI](../../../../getting-started/dovetail-ui) section.
* Open "Develop Flogo" environment.
* Select the "extensions" tab and upload the fabric extensions found at our [releases](https://github.com/TIBCOSoftware/dovetail-contrib/releases) repository (NOTE: You will need to download and then upload the fabric-extension.zip and fabric-client-extension.zip files).

* IMPORTANT: For smart contract modelling only Open Source extensions can be used, for this hello world tutorial we will use the following extensions
    * fabric-extension.zip from [dovetail-contrib](https://github.com/TIBCOSoftware/dovetail-contrib).
    * [Flogo log](https://github.com/project-flogo/contrib/tree/master/activity/log) activity.
    * [Flogo return](https://github.com/project-flogo/contrib/tree/master/activity/actreturn) activity.

For more detail information go to our [dovetail-contrib](https://github.com/TIBCOSoftware/dovetail-contrib/tree/master/hyperledger-fabric/tci) page.






