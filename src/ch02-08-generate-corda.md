## 5. Generate Corda Artifacts

Run following command from iou_tutorial directory, R3 Corda contract is written to artifacts/corda folder, put kotlin-IOU-1.0.0.jar on your classpath to develop your CorDapp. A sample IOUApp has been created and available [here](tutorials/iou/iou_tutorial.zip) inside iou_tutorial/network/corda folder

```
cd iou_tutorial
dovetail corda contract generate -m artifacts/IOU.json -v 1.0.0 -t artifacts/corda --namespace com.example.iou
```