## 5. Generate Hyperledger Fabric Blockchain Artifacts

Run following command from iou_tutorial directory Hyperledger Fabric chaincode is written to artifacts/hlf folder

For testing, transaction security support is not enabled.

```bash
dovetail contract generate -b fabric -m artifacts/IOU.json -v 1.0.0 -t artifacts/hlf
```
