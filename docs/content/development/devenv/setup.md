---
title: Setup your Dev environment
weight: 4604
---

1. Install [Go](https://golang.org/doc/install) and create workspace and setup GOPATH environment viarable as instructed
2. Install [Docker](https://docs.docker.com/) and [Docker Compose](https://docs.docker.com/compose/install/)
3. Install [Hyperledger Fabric](https://hyperledger-fabric.readthedocs.io/en/release-1.3/install.html)
4. Install [Hyperledger Composer CLI](https://hyperledger.github.io/composer/v0.19/installing/installing-index)
5. Install [Visio Studio Code](https://code.visualstudio.com/docs/setup/setup-overview) and Hyperledger Composer plugin
6. Install [Maven](https://maven.apache.org/install.html)
   * make sure you have access to following jars
   * com.tibco.dovetail:dovetail-corda:0.0.1 is available [here](https://github.com/TIBCOSoftware/dovetail/blob/master/docs/content/labs/network/corda)
    ```
        <dependency>
            <groupId>org.jetbrains.kotlin</groupId>
            <artifactId>kotlin-stdlib-jre8</artifactId>
            <version>1.1.60</version>
        </dependency>
        <dependency>
            <groupId>net.corda</groupId>
            <artifactId>corda-core</artifactId>
            <version>[2.0.0,)</version>
        </dependency>
        <dependency>
            <groupId>net.corda</groupId>
            <artifactId>corda-finance</artifactId>
            <version>[2.0.0,)</version>
        </dependency>
        <dependency>
            <groupId>com.tibco.dovetail</groupId>
            <artifactId>dovetail-corda</artifactId>
            <version>0.0.1</version>
        </dependency>
    ```
7. Install [flogo-lib](https://github.com/TIBCOSoftware/flogo-lib)
8. Install [flogo-contrib](https://github.com/TIBCOSoftware/flogo-contrib)
9. Install [go-bindata](https://github.com/jteeuwen/go-bindata)
10. Intall [govendor](https://github.com/kardianos/govendor)
11. Install [Project Dovetail™ Studio]()
   - after starting up the studio, you can upload contrib-smartcontract.zip from the Extensions tab
12. Install [dovetail-cli](https://github.com/TIBCOSoftware/dovetail-cli/blob/master/)

