##  Generate CorDapp artifacts

### 1. Export CorDapp flows
> * Export charlie application as charlie.json to artifacts/ folder
> * Export alice application as alice.json to artifacts/ folder
> * Export bob application as bob.json to artifacts/ folder

### 2. Generate CorDapp dependencies

CorDapp responder flows have compile time and runtime dependency on their corresponding initiator flows, however since flows are private to each organization, Tibco Dovetail generates a base initator class without implementation details and can be shared with business partners.

run following command from iou_tutorial folder

#### 2.1 charlie

```bash
dovetail corda dapp generate -m artifacts/charlie.json -v 1.0.0 -t artifacts/corda --namespace com.charlie.iou.flows --api
```

#### 2.2 alice

```bash
dovetail corda dapp generate -m artifacts/alice.json -v 1.0.0 -t artifacts/corda --namespace com.alice.iou.flows --api
```
### 3. Create dependency pom file for each organization

#### 3.1 charlie

copy following to artifacts/charlie.pom file

```javascript
<dependency>
    <groupId>com.alice.iou.flows</groupId>
    <artifactId>alice-api</artifactId>
    <version>1.0.0</version>
</dependency>
```

#### 3.2 alice

copy following to artifacts/alice.pom file

```javascript
<dependency>
    <groupId>com.charlie.iou.flows</groupId>
    <artifactId>charlie-api</artifactId>
    <version>1.0.0</version>
</dependency>
```

#### 3.2 bob

copy following to artifacts/bob.pom file

```javascript
<dependency>
    <groupId>com.charlie.iou.flows</groupId>
    <artifactId>charlie-api</artifactId>
    <version>1.0.0</version>
</dependency>
<dependency>
    <groupId>com.alice.iou.flows</groupId>
    <artifactId>alice-api</artifactId>
    <version>1.0.0</version>
</dependency>
```

### 4. Generate CorDapps

run following command from iou_tutorial folder

### 4.1 charlie
```bash
dovetail corda dapp generate -m artifacts/charlie.json -v 1.0.0 -t artifacts/corda --namespace com.charlie.iou.flows --dependency-file artifacts/charlie.pom
```

### 4.2 alice
```bash
dovetail corda dapp generate -m artifacts/alice.json -v 1.0.0 -t artifacts/corda --namespace com.alice.iou.flows --dependency-file alice.pom
```

### 4.3 bob
```bash
dovetail corda dapp generate -m artifacts/bob.json -v 1.0.0 -t artifacts/corda --namespace com.bob.iou.flows --dependency-file bob.pom
```