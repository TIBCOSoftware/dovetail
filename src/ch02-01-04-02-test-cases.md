## Test IOU Smart Contract and dapp on Corda

### 1 Issue cash

Use bank's swagger UI to issue and transfer cash to charlie. 

* go to security-controller section
    * login method
        * username = user1, password = test
    * execute

* go to cash-controller section
    * /cash/flow/issue-payment method
        * Authorizarion: copy the token value from login output, without the quotes
        * Input:
            ```
            {
                "amount": {
                    "quantity": 100000,
                    "currency": "USD"
                },
                "issuerRef": "00",
                "recipient": "O=charlie,L=New York,C=US",
                "anonymous": false,
                "notary": "O=Notary,L=London,C=GB"
            }  
            ```      
        * execute, if successful, you should see output similar to following
        ```
        {
            "transactionId": "FC85B7E74E8D9A4BDD39A2D83FA2C41C4CB57A296E43A3B5D9AF08FC7CF06E1B"
        }  
        ```

### 2 Issue an IOU

Use charlie's swagger UI to issue an IOU
* go to security-controller section
    * login method
        * username = user1, password = test
    * execute

* go to query-controller section
    * /query/states
    * Authorizarion: copy the token value from login output, without the quotes
    * execute, you should see cash is in the vault

* go to main-controller section
    * /api/issueiouinitiatort method
        * Authorizarion: copy the token value from login output, without the quotes
        * Input:
        ```
        {
            "holder": "O=alice,L=New York,C=US",
            "amt": {
                "quantity": 10000,
                "currency": "USD"
            },
            "extId": "iou1"
        }
        ```
        * execute

* go to query-controller section
    * /query/states
    * Authorizarion: copy the token value from login output, without the quotes
    * execute, you should see com.example.iou.IOU is now in the vault

### 3 Transfer the IOU

Use alice's swagger to transfer IOU to bob. 

* go to security-controller section
    * login method
        * username = user1, password = test
    * execute

* go to query-controller section
    * /query/states
    * Authorizarion: copy the token value from login output, without the quotes
    * execute, you should see com.example.iou.IOU in the output, copy the linearId.id value

* go to main-controller section
    * /api/transferiouinitiator
        * Authorizarion: copy the token value from login output, without the quotes
        * Input
        ```
        {
            "iouId": {
                "externalId": "iou1",
                "id": "paste the linearId.id value here"
            },
            "newHoder": "O=bob,L=New York,C=US"
        }
        ```
        * execute

* go to query-controller section
    * /query/states
    * Authorizarion: copy the token value from login output, without the quotes
    * execute, com.example.iou.IOU is no longer in the vault

### 4 Settle the IOU

Use charlie's swagger UI to settle the IOU with cash

* go to main-controller section
    * /api/settleiouinitiator method
        * Authorizarion: copy the token value from login output, without the quotes
        * Input:
        ```
        {
            "iouId": {
                "externalId": "iou1",
                "id": "copy iou linearId.id value here"
            }
        }
        ```
    * go to query-controller section
        * /query/states
        * Authorizarion: copy the token value from login output, without the quotes
        * execute, com.example.iou.IOU is no longer in the vault, and cash amount is reduced by $100