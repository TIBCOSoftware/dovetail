## 5. Generate Corda Artifacts

Run following command from iou_tutorial directory, R3 Corda contract is written to artifacts/corda folder, put kotlin-IOU-1.0.0.jar on your classpath to develop your CorDapp. A sample IOUApp has been created and available [here](https://github.com/TIBCOSoftware/dovetail/blob/master/docs/content/labs/network/corda)

```
cd iou_tutorial
dovetail contract generate -b corda -m artifacts/IOU.json -v 1.0.0 -t artifacts/ --namespace com.example.iou
```
