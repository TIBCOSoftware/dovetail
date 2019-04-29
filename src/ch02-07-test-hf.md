## 6. Test IOU Smart Contract on Hyperledger Fabric

We need four docker images in order for "dev mode" to run against the supplied
docker compose script.  If you installed the ``fabric-samples`` repo clone and
followed the instructions to [install binaries](https://hyperledger-fabric.readthedocs.io/en/latest/install.html), then
you should have the necessary Docker images installed locally.


> If you have downloaded the full tutorial files, jump to step 2

## 1. Copy Hyperledger Fabric Network


Copy sample Hyperledger Fabric Network to your network/fabric directory

```
cd iou_tutorial/network && \
curl -OL https://TIBCOSoftware.github.io/dovetail/tutorials/iou/fabric.zip && \
unzip fabric.zip && \
rm fabric.zip
```


## 2. open a terminal, and execute following command to start up the network

```
cd fabric
docker-compose -f docker-compose-simple.yaml up
```

open a terminal window, run ```docker ps```, make sure container cli, chaincode, peer, orderer and couchdb are running

## 3. open another terminal, execute following commands to connect to chaincode container and compile/start chaincode

```
docker exec -it chaincode bash
cd iou
go build
CORE_PEER_ADDRESS=peer:7052 CORE_CHAINCODE_ID_NAME=iou:0 ./iou
```

## 4. open a third terminal, execute following commands to install/instantiate chaincode from cli 

```
docker exec -it cli bash
peer chaincode install -p chaincodedev/chaincode/iou -n iou -v 0
peer chaincode instantiate -n iou -v 0 -c '{"Args":[]}' -C myc
```

## 5. from the third (cli) terminal,execute following commands to invoke IssueIOU and getIOU transactions

```
peer chaincode invoke -n iou -c '{"Args":["com.example.iou.IssueIOU","{\\"issuer\\":\\"charlie\\",\\"owner\\":\\"alice\\",\\"amt\\":{\\"quantity\\":10000,\\"currency\\":\\"USD\\"},\\"linearId\\":\\"testiou\\"}"]}' -C myc
peer chaincode query -n iou -c '{"Args":["com.example.iou.getIOU","testiou"]}' -C myc
```

You should see the IOU returned.

## 6. from the cli terminal,execute following commands to invoke TransferIOU and getIOU transactions

```
peer chaincode invoke -n iou -c '{"Args":["com.example.iou.TransferIOU","{\\"linearId\\":\\"testiou\\"}", "bob"]}' -C myc
peer chaincode query -n iou -c '{"Args":["com.example.iou.getIOU","testiou"]}' -C myc
```

You should see the IOU now with new owner "bob" returned

## 7. query iou

let's add another IOU issued by charlie

```
peer chaincode invoke -n iou -c '{"Args":["com.example.iou.IssueIOU","{\\"issuer\\":\\"charlie\\",\\"owner\\":\\"john\\",\\"amt\\":{\\"quantity\\":20000,\\"currency\\":\\"USD\\"},\\"linearId\\":\\"testioujohn\\"}"]}' -C myc
peer chaincode query -n iou -c '{"Args":["com.example.iou.getIOUIssuedBy", "charlie"]}' -C myc
```

You should see two IOUs returned, both issued by charlie.

## 8. shutdown network

```
docker-compose -f docker-compose-simple.yaml down
```