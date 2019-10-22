## Dovetail Studio Requirements

### Docker
Docker is needed to run Dovetail Studio, you can find the installation details here:

* [Docker](https://docs.docker.com/)
* [Docker Compose](https://docs.docker.com/compose/install/)


## Hyperledger Fabric Requirements

### Hyperledger Fabric SDK
If you are developing for Hyperledger Fabric, below is a link to its installation instructions.

* [Hyperledger Fabric](https://hyperledger-fabric.readthedocs.io/en/release-1.3/install.html)

## R3 Corda Requirements

### Java Programming Language and Tools
If you are developing for R3 Corda you will need Java runtime.

* IMPORTANT: [Java SE Development Kit 8](https://www.oracle.com/technetwork/java/javase/downloads/index.html) is required
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
        <version>[0.2.0, )</version>
    </dependency>
```

 > Project Dovetail™ library for R3 Corda is available [here](https://github.com/TIBCOSoftware/dovetail-contrib/releases/download/v0.2.0/dovetail-corda-0.2.0.jar), run following command to install it to your local Maven repository

```
mkdir dovetail_corda
cd dovetail_corda
curl -OL https://github.com/TIBCOSoftware/dovetail-contrib/releases/download/v0.2.0/dovetail-corda-0.2.0.jar
mvn org.apache.maven.plugins:maven-install-plugin:2.5.2:install-file -DgeneratePom=true -DgroupId=com.tibco.dovetail -DartifactId=dovetail-corda -Dversion=0.2.0 -Dfile=dovetail_corda_v0.2.0.jar -Dpackaging=jar
```