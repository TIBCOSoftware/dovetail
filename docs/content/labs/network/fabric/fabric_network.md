# Test IOU chaincode in Hyperledger Fabric Network

We need four docker images in order for "dev mode" to run against the supplied
docker compose script.  If you installed the ``fabric-samples`` repo clone and
followed the instructions to [download-platform-specific-binaries](http://hyperledger-fabric.readthedocs.io/en/latest/samples.html#download-platform-specific-binaries), then
you should have the necessary Docker images installed locally.

## 1. open a terminal, and execute following command to start up the network
> docker-compose -f docker-compose-simple.yaml up

open a terminal window, run ```docker ps```, make sure container cli, chaincode, peer, orderer and couchdb are running

## 2. open another terminal, execute following commands to connect to chaincode container and compile/start chaincode

> docker exec -it chaincode bash

> cd iou

> go build

> CORE_PEER_ADDRESS=peer:7052 CORE_CHAINCODE_ID_NAME=iou:0 ./iou

## 3. open a third termial, execute following commands to install/instantiate chaincoce from cli 

> docker exec -it cli bash

> peer chaincode install -p chaincodedev/chaincode/iou -n iou -v 0

> peer chaincode instantiate -n iou -v 0 -c '{"Args":[]}' -C myc

## 4. from the third (cli) terminal,execute following commands to invoke IssueIOU and getIOU transactions
> peer chaincode invoke -n iou -c '{"Args":["com.example.iou.IssueIOU","{\\"issuer\\":\\"charlie\\",\\"owner\\":\\"alice\\",\\"amt\\":{\\"quantity\\":10000,\\"currency\\":\\"USD\\"},\\"linearId\\":\\"testiou\\"}"]}' -C myc

> peer chaincode query -n iou -c '{"Args":["com.example.iou.getIOU","testiou"]}' -C myc

You should see the IOU returned.

## 5. from the cli terminal,execute following commands to invoke TransferIOU and getIOU transactions
> peer chaincode invoke -n iou -c '{"Args":["com.example.iou.TransferIOU","{\\"linearId\\":\\"testiou\\"}", "bob"]}' -C myc

> peer chaincode query -n iou -c '{"Args":["com.example.iou.getIOU","testiou"]}' -C myc

You should see the IOU now with new owner "bob" returned

## 6. query iou

let's add another IOU issued by charlie

> peer chaincode invoke -n iou -c '{"Args":["com.example.iou.IssueIOU","{\\"issuer\\":\\"charlie\\",\\"owner\\":\\"john\\",\\"amt\\":{\\"quantity\\":20000,\\"currency\\":\\"USD\\"},\\"linearId\\":\\"testioujohn\\"}"]}' -C myc

> peer chaincode query -n iou -c '{"Args":["com.example.iou.getIOUIssuedBy", "charlie"]}' -C myc

You should see two IOUs returned, both issued by charlie.

## 7. shutdown network
> docker-compose -f docker-compose-simple.yaml down

