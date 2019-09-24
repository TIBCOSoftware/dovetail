## 6. Test IOU Smart Contract on Corda

## 6.4. Generate Deploy CorDapps

### 6.4.1 Copy IOU CorDapp deployment script

Copy deployment script to your network/corda directory

```
cd iou_tutorial/network/corda
```

```
curl -OL https://TIBCOSoftware.github.io/dovetail/tutorials/iou/corda.zip && \
unzip corda.zip && \
rm corda.zip
```

### 6.4.2 Deploy nodes

From the network/corda directory, run following command, it will take a few minutes.

```bash
./gradlew clean deployNodes
```

### 6.4.2 Run nodes

Run following command to start up all nodes, it will take a few minutes.

```bash
cd build/nodes
./runnodes
```
