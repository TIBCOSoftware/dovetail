<p align="center">
  <img src ="images/noutfoundprojectflogo.png" alt ="TODO : Replace with project Dovetail image" />
</p>

<p align="center" >
  <b>Project Dovetail™ is an Open Source application for blockchain smart contracts</b>
</p>

<p align="center">
  <b>Documentation Status :</b> <img src="https://codebuild.us-east-1.amazonaws.com/badges?uuid=eyJlbmNyeXB0ZWREYXRhIjoiSFZraVI4YTY2UHFua09RaTNnQTZKYTVlUnp0RktPVk9CeWM3cmlMODNraEllUk9UNG9KSTgxUjUyQ0pPOUZtcUVyRHVoaWI3WEh3VDNMRll5WjZUREVzPSIsIml2UGFyYW1ldGVyU3BlYyI6IjZMZUh4MXFWSVlZZHdjeTAiLCJtYXRlcmlhbFNldFNlcmlhbCI6MX0%3D&branch=master"/>
  <img src="https://img.shields.io/badge/license-BSD%20style-blue.svg"/>
  </a>
</p>


<p align="center">
  <a href="#project-dovetail">Dovetail™</a> | <a href="#dovetail-core">Architecture</a> | <a href="#what-are-smart-contracts">Blockchains smart contracts</a> | <a href="#contributing">Contributing</a> | <a href="#license">License</a>
</p>

<br/>

Developing blockchain solutions with today’s technology is challenging, given the lack of tooling and standardization. Project Dovetail™ by TIBCO LABS™ addresses these issues by providing a graphical interface for modeling smart contracts, making them easier to write, visualize, test, and audit, all without deep programming experience. Logic is abstracted from low-level code, and contracts can be deployed into different blockchain stacks with little or no change. On-chain and off-chain computation can become more seamless, time to market is improved, and the risk of technology lock-in is reduced. Project Dovetail makes your smart contracts smarter. 
<br/>

<p align="center">
  <img src ="images/eventhandlers.png" />
</p>

# What are Smart Contracts

There are many definitions and descriptions of smart contracts, as each blockchain framework tends to implement this capability in its own way (if it does implement this feature at all).  However, generally speaking, smart contracts (at least from an enterprise / permissioned perspective) may be thought of as application or business logic (defined in code) that runs within the context of a blockchain network itself.  The idea is to automate the processing of blockchain transactions across the network, use the logic to determine if a transaction will get written to the ledger, and to maintain accuracy, compliance, and trust.  In essence, smart contracts represent a method for controlling how changes are made to the underlying blockchain, in a non-centralized and (potentially) untrustworthy environment.  Supply chain transaction tracking, healthcare revenue cycle management, consumer contract execution, and government interactions are all areas in which smart contracts (business logic) may be applied.  Remember, however, that "smart contracts" are not necessarily "smart" (depends on the code) nor "contracts" (as they may not be viewed as legal contracts).  Thus, the term can be a bit misleading :)


# Project Dovetail

Project Dovetail™ is a framework that allows for the end to end design, development, testing, and deployment of blockchain smart contracts.  Project Dovetail™ allows you to develop smart contracts based on a series of models, helping to:

* **Increase security** since the modeling abstraction layer will allow for reusable and tested code derived from the model.
* **Simplify** development via an extendable abstraction layer (flow model)
* **Decouple** your code from the underlying blockchain technology
* **Reduce** the amount of code needed
* Increase visibility and audit-ability
* Allow customizable modeling for your industry
* **Expose** a better UI for the design of smart contracts

# Dovetail Core

Dovetail is based on TIBCO FLOGO™, an event-driven app framework used to develop apps for the cloud & IoT edge. It can also be thought of as a lightweight *app kernel* used by open source & commercial solutions like Dovetail here. The trigger used is based on the cli to generate or transpile your smart contract logic into blockahin technologies languages : R3 Corda, HyperLedger Fabric...

Dovetail Core provides the following key benefits:

⛓ **Action chaining** enables communication between one or more capabilities in a single, sub 10MB binary!<br/>
🏗 **Common contribution model** build activities and triggers that can be leveraged by all capabilities<br/>
🔨 **Extensible** easily extend the capabilities available by building your own action using the common interfaces<br/>

## Dovetail Core Contribution Model

Dovetail™ Core exposes three principal contribution interfaces that enable developers to build common capabilities and functionality. These contribution interfaces include:

* **Trigger Interface** a common interface for building event-consumers that dispatch events to one or more actions. The Kafka subscriber is an example of a trigger.
* **Activity Interface** a common interface for exposing common application logic in a reusable manner. Think of this as a function, such as write to database, publish to Kafka, etc that can be used by all Dovetail apps.

# Repos

Project Dovetail consists of the following sub-projects available as separate repos:

* [dovetail-cli](https://github.com/TIBCOSoftware/dovetail-cli):  Command line tools for building Dovetail apps & extensions
* [dovetail-contrib](https://github.com/TIBCOSoftware/dovetail-contrib): Dovetail contributions/extensions

# Dovetail Flows

Dovetail Flows provides smart contract logic design capabilities and includes the following key highlights.

🌈 **Painless development** Visual modeler with step-back debugging capabilities & elegant DSL<br/>
⚙️ **Ultra-light process engine** for conditional flow control


## Getting Started

We've made getting started with Dovetail Flows as easy as possible. The current set of tooling is designed for:

- Serverless function developers
- Cloud-native microservices developers
- IoT Solutions developers
- <a href="#golang-api">Go Developers</a>

### Zero-code Developers

If your background is in or you prefer to develop your apps using zero-coding environments, then read on, because we’ve got something special for you.

Flows Web UI is available via [Docker Hub](https://hub.docker.com/r/flogo/flogo-docker) or [Dovetail.io](http://flogo.io). The Docker image contains the Flows Web UI along with all required components to begin developing, testing and building deployable artifacts right from your web browser.

To report any issues with the Issue tracker on this project.

![Dovetail Web In Action](images/flogo-web2.gif)

## Getting Started

We’ve made building powerful streaming pipelines as easy as possible. Develop your smart contracts pipelines using:

- A simple, clean JSON-based DSL
- Golang API

See the sample below of an aggregation pipeline (for brevity, the triggers and metadata of the resource has been omitted). Also don’t forget to check out the examples in the [project-flogo/stream](https://github.com/project-flogo/stream/tree/master/examples) repo.

```json
  "stages": [
    {
      "ref": "github.com/TIBCOSoftware/dovetail-contrib/activity/aggregate",
      "settings": {
        "function": "sum",
        "windowType": "timeTumbling",
        "windowSize": "5000"
      },
      "input": {
        "value": "=$.input"
      }
    },
    {
      "ref": "github.com/TIBCOSoftware/dovetail-contrib/activity/log",
      "input": {
        "message": "=$.result"
      }
    }
  ]
```

# The CLI

The CLI is used to build all applications that leverage the JSON-based DSL. If you’re using the Go API to build your apps, feel free to just `go build` your stuff without the flogo CLI.

Getting started with the CLI couldn't be any easier (refer to [Dovetail CLI](https://github.com/TIBCOSoftware/dovetail-cli) repo for detail instructions and dependencies):

* Install the CLI
```bash
go get -u github.com/TIBCOSoftware/dovetail-cli/...
```

* Create & build your app
<img src="images/dovetail-cli.gif" width="70%"/>

* **dovetail** the main CLI for creating and building  and testing your blockchains smart contracts

If you're interested in building your own contribution(s), refer to the [Dovetail Documentation](https://tibcosoftware.github.io/flogo/) or join us on the [project-flogo/Lobby Gitter Channel](https://gitter.im/project-flogo/Lobby?utm_source=share-link&utm_medium=link&utm_campaign=share-link).

# Contributing
Want to contribute to Project Dovetail? We've made it easy, all you need to do is fork the repository you intend to contribute to, make your changes and create a Pull Request! Once the pull request has been created, you'll be prompted to sign the CLA (Contributor License Agreement) online.

Not sure where to start? No problem, here are a few suggestions:

* [dovetail-contrib](https://github.com/TIBCOSoftware/dovetail-contrib): This repository contains all of the contributions, such as activities, triggers, etc. Perhaps there is something missing? Create a new activity or trigger or fix a bug in an existing activity or trigger.
* Browse all of the Project Dovetail repositories and look for issues tagged `kind/help-wanted` or `good first issue`

If you have any questions, feel free to post an issue and tag it as a question, email flogo-oss@tibco.com or chat with the team and community:

* The [project-flogo/Lobby](https://gitter.im/project-flogo/Lobby) Gitter channel should be used for general discussions, start here for all things Dovetail!
* The [project-flogo/developers](https://gitter.im/project-flogo/developers) Gitter channel should be used for developer/contributor focused conversations. 

For additional details, refer to the [Contribution Guidelines](https://github.com/TIBCOSoftware/flogo/blob/master/CONTRIBUTING.md).

# License 
The top level flogo repo, consisting of flow samples & documentation, is licensed licensed under a BSD-style license. Refer to [LICENSE](https://github.com/TIBCOSoftware/flogo/blob/master/LICENSE) for license text.

Dovetail source code in [dovetail-cli](https://github.com/TIBCOSoftware/dovetail-cli), [flogo-lib](https://github.com/TIBCOSoftware/flogo-lib), [dovetail-contrib](https://github.com/TIBCOSoftware/dovetail-contrib), [flogo-services](https://github.com/TIBCOSoftware/flogo-services) are all licensed under a BSD-style license, refer to [LICENSE](https://github.com/TIBCOSoftware/flogo/blob/master/LICENSE) 

## Usage Guidelines

We’re excited that you’re using Project Dovetail to power your project(s). Please adhere to the [usage guidelines](http://flogo.io/brand-guidelines) when referencing the use of Project Dovetail within your project(s) and don't forget to let others know you're using Project Dovetail by proudly displaying one of the following badges or the Flynn logo, found in the [branding](branding) folder of this project.
