---
title: Transaction
weight: 4603
---

# Transaction
This trigger allows user to select a predefined transaction in the selected common data model to implement, all schemas defined in the common data model will be imported into the application that utilizes this trigger.

## Settings
| Setting     | Required | Description |
|:------------|:---------|:------------|
| model       | True     | Common data model name |
| createAll   | True     | Create flows for all transactions defined in the model, or select a specific transaction |
| transaction | True     | Select the transaction to implement |


