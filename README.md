<p align="center">
  <img src ="images/TIBCO Labs final with TM2-08_email.png" />
</p>

<p align="center" >
  <b>Project Dovetail™ is an Open Source tool for modelling blockchain smart contracts</b>
</p>

[![Documentation Status](https://travis-ci.org/TIBCOSoftware/dovetail.svg?branch=master)](https://travis-ci.org/TIBCOSoftware/dovetail)<img src="https://img.shields.io/badge/license-BSD%20style-blue.svg"/>

<p align="center">
  <a href="#project-dovetail">Dovetail™</a> | <a href="#dovetail-core">Architecture</a> | <a href="#what-are-smart-contracts">Blockchains smart contracts</a> | <a href="#contributing">Contributing</a> | <a href="#license">License</a>
</p>

<br/>

Developing blockchain solutions with today’s technology is challenging given the lack of tooling and standardization. Project Dovetail™ by TIBCO LABS™ addresses these issues by providing a graphical interface for modeling smart contracts, making them easier to write, visualize, test, and audit, all without deep programming experience. 

Logic is abstracted from low-level code, and contracts can be developed for different blockchain stacks with little or no code. 
On-chain and off-chain computation can become more seamless, time to market is improved, and the risk of technology lock-in is reduced. Project Dovetail makes your smart contracts smarter. 
<br/>

<p align="center">
  <img src ="images/eventhandlers.png" />
</p>

# What are Smart Contracts

There are many definitions and descriptions of smart contracts, as each blockchain framework tends to implement this capability in its own way (if it does implement this feature at all). However, generally speaking, smart contracts (at least from an enterprise / permissioned perspective) may be thought of as application or business logic (defined in code) that runs within the context of a blockchain network itself.

The idea is to automate the processing of blockchain transactions across the network, use the logic to determine if a transaction will get written to the ledger, and to maintain accuracy, compliance, and trust.  In essence, smart contracts represent a method for controlling how changes are made to the underlying blockchain, in a non-centralized and (potentially) untrustworthy environment.  Supply chain transaction tracking, healthcare revenue cycle management, consumer contract execution, and government interactions are all areas in which smart contracts (business logic) may be applied.
Remember, however, that "smart contracts" are not necessarily "smart" (depends on the code) nor "contracts" (as they may not be viewed as legal contracts).  Thus, the term can be a bit misleading :)


# Project Dovetail

Project Dovetail™ is a framework that allows for the end to end design, development, testing, and deployment of blockchain smart contracts.  Project Dovetail™ allows you to develop smart contracts based on a series of models, helping to:

* **Increase security** since the modeling abstraction layer will allow for reusable and tested code derived from the model.
* **Simplify** development via an extendable abstraction layer (flow model)
* **Decouple** your code from the underlying blockchain technology
* **Reduce** the amount of code needed
* **Increase visibility** and audit-ability
* **Allow customizable modeling** for your industry
* **Expose** a better UI for the design of smart contracts

# Dovetail Core

Dovetail is based on TIBCO FLOGO™, an event-driven app framework used to develop apps for the cloud & IoT edge. It can also be thought of as a lightweight *app kernel* used by open source & commercial solutions like Dovetail here. The trigger used is based on the cli to generate or transpile your smart contract logic into blockchain technologies languages : R3 Corda, HyperLedger Fabric...

Dovetail Core provides the following key benefits:

⛓ **Action chaining** enables communication between one or more capabilities in a single, sub 10MB binary!<br/>
🏗 **Common contribution model** build activities and triggers that can be leveraged by all capabilities<br/>
🔨 **Extensible** easily extend the capabilities available by building your own action using the common interfaces<br/>

## Dovetail Core Contribution Model

Dovetail™ Core exposes three principal contribution interfaces that enable developers to build common capabilities and functionality. These contribution interfaces include:

* **Connector Interface** a common interface for importing predefined schemas into Dovetail, Hyperledger Composer Connector is an example.
* **Trigger Interface** a common interface for building event-consumers that dispatch events to one or more actions. The Smart Contract TXN Trigger is an example of a trigger.
* **Activity Interface** a common interface for exposing common application logic in a reusable manner. Think of this as a function, such as write to ledger, publish events, etc that can be used by all Dovetail apps.


# Dovetail Flows

Dovetail Flows provides smart contract logic design capabilities and includes the following key highlights.

🌈 **Painless development** Visual modeler with step-back debugging capabilities & elegant DSL<br/>
⚙️ **Ultra-light process engine** for conditional flow control


### Zero-code Developers

If your background is in or you prefer to develop your smart contracts using zero-coding environments, then read on, because we’ve got something special for you.

Flows Web UI is available via [Dovetail releases page](https://github.com/TIBCOSoftware/dovetail/releases).

To report any issues with the Issue tracker on this project.

# Dovetail Documentation

Dovetail documentation can be found on the [documentation page](https://tibcosoftware.github.io/dovetail/) and the source code in the [github page](https://github.com/TIBCOSoftware/dovetail).

# Contributing
Want to contribute to Project Dovetail? We've made it easy, all you need to do is fork the repository you intend to contribute to, make your changes and create a Pull Request! Once the pull request has been created, you'll be prompted to sign the CLA (Contributor License Agreement) online.

Not sure where to start? No problem, here are a few suggestions:

* [dovetail-contrib](https://github.com/TIBCOSoftware/dovetail-contrib): This repository contains all of the contributions, such as activities, triggers, etc. Perhaps there is something missing? Create a new activity or trigger or fix a bug in an existing activity or trigger.
* Browse all of the Project Dovetail repositories and look for issues tagged `kind/help-wanted` or `good first issue`

If you have any questions, feel free to post an issue and tag it as a question or email <tibcolabs@tibco.com>. 

For additional details, refer to the [“Contribution Guidelines”](https://github.com/TIBCOSoftware/dovetail/blob/master/CONTRIBUTING.md).

# License 
The top level flogo repo, consisting of flow samples & documentation, is licensed licensed under a BSD-style license. Refer to [LICENSE](https://github.com/TIBCOSoftware/dovetail/blob/master/LICENSE) for license text.

Dovetail source code in [dovetail-cli](https://github.com/TIBCOSoftware/dovetail-cli), [dovetail-contrib](https://github.com/TIBCOSoftware/dovetail-contrib), [dovetail-java-lib](https://github.com/TIBCOSoftware/dovetail-java-lib) are all licensed under a BSD-style license, refer to [LICENSE](https://github.com/TIBCOSoftware/dovetail/blob/master/LICENSE) 

