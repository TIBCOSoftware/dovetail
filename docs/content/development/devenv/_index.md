---
title: Environment Prerequisites
weight: 4200
chapter: true
pre: "<i class=\"fas fa-code\" aria-hidden=\"true\"></i> "
---

### Go Programming Language and Tools
Project Dovetail™ CLI is written in Go Programming Language, and it uses a few Go tools to package dependencies and resources

* [Go](https://golang.org/doc/install) version 1.11.x is required


## Hyperledger Fabric SDK
If you are developing for Hyperledger Fabric, below is a link to its installation instructions

* [Hyperledger Fabric](https://hyperledger-fabric.readthedocs.io/en/release-1.3/install.html)

### Java Programming Language and Tools
Project Dovetail™ provides smart contract flow engine for distributed ledger platform that requires Java runtime, such as R3 Corda.

* [Java SE Development Kit 8 and above](https://www.oracle.com/technetwork/java/javase/downloads/index.html) is required
* [Maven](https://maven.apache.org/install.html)

If you are developing for R3 Corda, following jars must be available in your local or public Maven Repository.

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
            <version>0.1.0</version>
        </dependency>
    ```
 *** com.tibco.dovetail:dovetail-corda:0.1.0 jar is available [here](https://github.com/TIBCOSoftware/dovetail-java-lib/releases/tag/v0.1.0)

### Docker

* [Docker](https://docs.docker.com/)
* [Docker Compose](https://docs.docker.com/compose/install/)