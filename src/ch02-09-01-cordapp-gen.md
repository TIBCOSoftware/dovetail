## 6. Test IOU Smart Contract on Corda

## 6.4. Generate CorDapp artifacts

### 6.4.1. Export CorDapp flows
> * Export charlie application as charlie.json to artifacts/ folder
> * Export alice application as alice.json to artifacts/ folder
> * Export bob application as bob.json to artifacts/ folder

### 6.4.2. Generate CorDapp dependencies

CorDapp responder flows have compile and runtime dependency on their corresponding initiator flows, however since flows are private to each organization, Tibco Dovetail generates a base initator class without implementation details and can be shared with business partners.

run following command from iou_tutorial folder

```bash
dovetail dapp generate -b corda -m artifiacts/charlie.json -v 1.0.0 -t artifacts/corda/charlie --namespace com.charlie.iou.flows --api

dovetail dapp generate -b corda -m artifacts/alice.json -v 1.0.0 -t artifacts/corda/alice --namespace com.alice.iou.flows --api
```
### 6.4.3. Create dependency pom file for each organization

#### 6.4.3.1 charlie

copy following to artifacts/charlie.pom file

```javascript
<dependency>
    <groupId>com.alice.iou.flows</groupId>
    <artifactId>alice_iou-api</artifactId>
    <version>[1.0.0, )</version>
</dependency>
```

#### 6.4.3.2 alice

copy following to artifacts/alice.pom file

```javascript
<dependency>
    <groupId>com.charlie.iou.flows</groupId>
    <artifactId>charlie_iou-api</artifactId>
    <version>[1.0.0, )</version>
</dependency>
```

#### 6.4.3.2 bob

copy following to artifacts/bob.pom file

```javascript
<dependency>
    <groupId>com.charlie.iou.flows</groupId>
    <artifactId>charlie_iou-api</artifactId>
    <version>[1.0.0, )</version>
</dependency>
<dependency>
    <groupId>com.alice.iou.flows</groupId>
    <artifactId>alice_iou-api</artifactId>
    <version>[1.0.0, )</version>
</dependency>
```

### 6.4.3. Generate CorDapps

run following command from iou_tutorial folder

```bash
dovetail dapp generate -b corda -m artifacts/charlie.json -v 1.0.0 -t artifacts/corda/charlie --namespace com.charlie.iou.flows --dependency-file artifacts/charlie.pom

dovetail dapp generate -b corda -m artifacts/alice.json -v 1.0.0 -t artifacts/corda/alice --namespace com.alice.iou.flows --dependency-file alice.pom

dovetail dapp generate -b corda -m artifacts/bob.json -v 1.0.0 -t corda/bob --namespace com.bob.iou.flows --dependency-file bob.pom
```