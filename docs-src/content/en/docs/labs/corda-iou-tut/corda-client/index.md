---
title: "Generate Corda Client"
linkTitle: "Generate Corda Client"
weight: 3
description: >
  generates RPC client from CorDApp flows and Smart contract metadata
---

Dovetail CLI generates RPC client from CorDApp flows and Smart contract metadata, the generated client is a standalone springboot web server that exposes REST APIs to allow applications to invoke CorDapp flows.

The generated client also contains an embedded event streaming service to pulish vault state changes to specified messaging service, at present, only TIBCO Cloud Messaging is supported, use commandline --streaming eftl to enable streaming service. 

### 1 Create dependency pom file

copy following to artifacts/iou.pom file

```xml
<dependency>
    <groupId>com.example.iou</groupId>
    <artifactId>IOU</artifactId>
    <version>1.0.0</version>
</dependency>
```

### 2. Generate Corda Client

run following command from iou_tutorial folder, a web client is generated for charlie, alice and bob. The last command is to generate a generic web client than will be used for party 'bank' to issue and transfer cash.

### 2.1. charlie

```bash
dovetail corda client generate --cordapp-json artifacts/charlie.json --smartcontract-json artifacts/IOU.json -v 1.0.0 -t artifacts/corda --cordapp-ns com.charlie.iou.flows --dependency-file artifacts/iou.pom --streaming none
```

### 2.2. alice
```bash
dovetail corda client generate --cordapp-json artifacts/alice.json --smartcontract-json artifacts/IOU.json -v 1.0.0 -t artifacts/corda --cordapp-ns com.alice.iou.flows --dependency-file artifacts/iou.pom --streaming none
```

### 2.3. bob
```bash
dovetail corda client generate --cordapp-json artifacts/bob.json --smartcontract-json artifacts/IOU.json -v 1.0.0 -t artifacts/corda --cordapp-ns com.bob.iou.flows --dependency-file artifacts/iou.pom --streaming none
```

### 2.4. bank
```bash
dovetail corda client generate -v 1.0.0 -t artifacts/corda
```