---
title: "IOU Smart Contract Data Model"
linkTitle: "Model Ledger Asset"
weight: 1
description: >
    Define contract state data schema that should be stored on the ledger
---

<p><video width="480" height="320" controls="controls">
    <source src="videos/iou_asset_schema.mp4" type="video/mp4">
</video></p>

* Start up and open Dovetail Studio WebUI
* Go to Connection tab
* Click "Add Connection"
  * Choose "Define Ledger Asset Schema" from the list
  * Enter display name, e.g. "my first IOU"
  * Enter asset name "IOU"
  * Enter module name "com.example.iou"
  * Select asset type "Linear State"
    * LinearId attribute will be automatically added to the pre-defined fields
  * Enter addition fields
    * issuer, select "Party" from the type dropdown menu
    * holder, select "Party" from the type dropdown menu
    * amt, select "Amount<Courrency>" from the type dropdown menu
  * Define party roles
    * add "issuer", leave other fields as default
    * add "holder", leave other fields as default
  * Click "Donee" button
