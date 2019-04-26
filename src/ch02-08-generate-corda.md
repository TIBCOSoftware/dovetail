## 5. Generate Corda Artifacts

Run following command from tutorial directory, R3 Corda contract is written to tutorial/artifacts/corda folder, put kotlin-IOU-1.0.0.jar on your classpath to develop your CorDapp. A sample IOUApp has been created and availble [here](https://github.com/TIBCOSoftware/dovetail/blob/master/docs/content/labs/network/corda)

```
dovetail contract generate -b corda -m artifacts/IOU.json -v 1.0.0 -t artifacts/ --namespace com.example.iou
```
