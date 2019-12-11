## Generate Corda Artifacts

> * Maven should be installed on your system, and mvn is in your PATH
> * Dovetail CLI should installed and in your PATH, it is also included in the cli folder in the corda.zip file (for mac only)
> * Latest relese of dovetail-corda-x.x.x.jar should be installed to your local maven repo, it is also included in corda.zip, you can run cli/mvninstall.sh to install the jar file

Run following command from iou_tutorial directory, R3 Corda contract  kotlin-IOUContract-1.0.0.jar is generated and written to artifacts/corda/iou folder, and installed in your local maven repository.

```
cd iou_tutorial
dovetail corda contract generate -m artifacts/IOU.json -v 1.0.0 -t artifacts/corda --namespace com.example.iou
```