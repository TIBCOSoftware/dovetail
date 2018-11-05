import {Injectable, Injector} from "@angular/core";
import {Http} from "@angular/http";
import {Observable} from "rxjs/Observable";
import {
    WiContrib,
    WiServiceHandlerContribution,
    IValidationResult,
    ValidationResult,
    ActionResult,
    IActionResult,
    ICreateFlowActionContext,
    CreateFlowActionResult,
    WiContribModelService,
    WiContributionUtils,
    IConnectorContribution
} from "wi-studio/app/contrib/wi-contrib";
import { ITriggerContribution, IFieldDefinition, IConnectionAllowedValue, MODE } from "wi-studio/common/models/contrib";
import * as lodash from "lodash";

@WiContrib({})
@Injectable()
export class SmartContractTriggerHandler extends WiServiceHandlerContribution {
    
    constructor(private injector: Injector, private http: Http, private contribModelService: WiContribModelService) {
        super(injector, http, contribModelService);
    }
    
    value = (fieldName: string, context: ITriggerContribution): Observable<any> | any => {
        let conId = context.getField("model").value;
        
        switch(fieldName) {
            case "model":
                let connectionRefs = [];
                
                return Observable.create(observer => {
                    WiContributionUtils.getConnections(this.http, "SmartContract").subscribe((data: IConnectorContribution[]) => {
                        data.forEach(connection => {
                            if ((<any>connection).isValid) {
                                for(let i=0; i < connection.settings.length; i++) {
                                    if(connection.settings[i].name === "name"){
                                        connectionRefs.push({
                                            "unique_id": WiContributionUtils.getUniqueId(connection),
                                            "name": connection.settings[i].value
                                        });
                                        break;
                                    }
                                }
                            }
                        });
                        observer.next(connectionRefs);
                    });
                });
                
            case "transaction":
                if(Boolean(conId) == false)
                    return null;

                return Observable.create(observer => {
                    WiContributionUtils.getConnection(this.http, conId)
                                        .map(data => data)
                                        .subscribe(data => {
                                            for (let setting of data.settings) {
                                                if (setting.name === "transactions") {
                                                    observer.next(setting.value);
                                                    break;
                                                }
                                            }
                                        });
                });
            case "transactions":
                if(Boolean(conId) == false)
                    return null;

                return Observable.create(observer => {
                    WiContributionUtils.getConnection(this.http, conId)
                                        .map(data => data)
                                        .subscribe(data => {
                                            for (let setting of data.settings) {
                                                if (setting.name === "transactions") {
                                                    observer.next(JSON.stringify(setting.value));
                                                    break;
                                                }
                                            }
                                        });
                });
            case "concepts":
                if(Boolean(conId) == false)
                    return null;

                return Observable.create(observer => {
                    WiContributionUtils.getConnection(this.http, conId)
                                        .map(data => data)
                                        .subscribe(data => {
                                            for (let setting of data.settings) {
                                                if (setting.name === "concepts") {
                                                    observer.next(JSON.stringify(setting.value));
                                                    break;
                                                }
                                            }
                                        });
                });
            case "assets":
                if(Boolean(conId) == false)
                    return null;

                return Observable.create(observer => {
                    WiContributionUtils.getConnection(this.http, conId)
                                        .map(data => data)
                                        .subscribe(data => {
                                            for (let setting of data.settings) {
                                                if (setting.name === "assets") {
                                                    observer.next(JSON.stringify(setting.value));
                                                    break;
                                                }
                                            }
                                        });
                });
            case "schemas":
                if(Boolean(conId) == false)
                    return null;

                return Observable.create(observer => {
                    WiContributionUtils.getConnection(this.http, conId)
                                        .map(data => data)
                                        .subscribe(data => {
                                            for (let setting of data.settings) {
                                                if (setting.name === "schemas") {
                                                    observer.next(JSON.stringify(setting.value));
                                                    break;
                                                }
                                            }
                                        });
                });
            case "transactionInput":
                let txn = context.getField("transaction").value;
                if(Boolean(conId) == false || Boolean(txn) == false)
                    return null;

                return Observable.create(observer => {
                    let txnRefs = [];
                    this.getSchemas(conId).subscribe( schemas => {
                       // console.log(schemas[txn]);
                        observer.next(schemas[txn]);
                    });
                });
            default: 
                return null;
        }
            
    }

    validate = (fieldName: string, context: ITriggerContribution): Observable<IValidationResult> | IValidationResult => {
        if(context.getMode() === MODE.WIZARD) {
            switch(fieldName){
                case "transaction":
                    let createAll = context.getField("createAll").value;
                    return Observable.create(observer => {
                        let vresult: IValidationResult = ValidationResult.newValidationResult();
                        vresult.setVisible(!createAll);
                        if(!createAll && Boolean(context.getField("transaction").value) == false)
                            vresult.setValid(false).setError("REQUIRED_VALUE_NOT_SET", "Please select a transaction");

                        observer.next(vresult);
                    });
                
                default:
                    return Observable.create(observer => {
                        let vresult: IValidationResult = ValidationResult.newValidationResult();
                        observer.next(vresult);
                    });
            }
        } else {
            switch(fieldName){
                case "createAll":
                case "assets":
                case "transactions":
                case "schemas":
                case "concepts":
                    return Observable.create(observer => {
                        let vresult: IValidationResult = ValidationResult.newValidationResult();
                        vresult.setVisible(false);
                        observer.next(vresult);
                    });
                default:
                    return Observable.create(observer => {
                        let vresult: IValidationResult = ValidationResult.newValidationResult();
                        let conId = context.getField("model").value;
                        
                        if(Boolean(conId)){
                            vresult.setReadOnly(true);
                        } else {
                            vresult.setReadOnly(false);
                        }
                        observer.next(vresult);
                    });
            }
        }
    }

    action = (actionId: string, context: ICreateFlowActionContext): Observable<IActionResult> | IActionResult => {
       
        let result = CreateFlowActionResult.newActionResult();
        let conId = context.getField("model").value;
        let createAll = context.getField("createAll").value;
        let flows = []
        return Observable.create(observer => {
                WiContributionUtils.getConnection(this.http, conId)
                            .map(data => data)
                            .subscribe(data => {
                                let txns = [];
                                let schemas = new Map();
                                for (let setting of data.settings) {
                                    if (setting.name === "transactions") {
                                        txns = setting.value;
                                    } else if(setting.name === "schemas") {
                                        setting.value.map(item => schemas[item[0]] = item[1]);
                                    }
                                }
                             
                                if (context.handler && context.handler.settings) {
                                    if(createAll){
                                        for(let i=0; i<txns.length; i++) {
                                            flows.push(this.createFlow(context, conId, txns[i], schemas, result));
                                        }
                                    } else {
                                        let txn = context.getField("transaction").value
                                        flows.push(this.createFlow(context, conId, txn, schemas, result));                                    
                                    }
                                }
        
                                let actionResult = ActionResult.newActionResult().setSuccess(true).setResult(result);
                                observer.next(actionResult);
                            });
        });
    }

    createFlow(context, conId, txn, schemas, result) : string{
        let modelService = this.getModelService();
        let trigger = modelService.createTriggerElement("SmartContract/SmartContractTrigger");
        if (trigger) {
            for(let t = 0; t < trigger.settings.length; t++) {
                if (trigger.settings[t].name === "model" ) {
                    trigger.settings[t].value = conId;
                    break;
                }
            }
            for (let s = 0; s < trigger.handler.settings.length; s++) {
                if (trigger.handler.settings[s].name === "transaction") {
                    trigger.handler.settings[s].value = txn;
                    break;
                } 
            }
            for (let j = 0; j < trigger.outputs.length; j++) {
                if (trigger.outputs[j].name === "transactionInput") {
                    trigger.outputs[j].value =  {
                        "value": schemas[txn],
                        "metadata": txn
                    };
                    break;
                }
            }
        }

        let flowName = context.getFlowName();
        if(context.getField("createAll").value)
            flowName = flowName + "_" + txn;

        let flowModel = modelService.createFlow(flowName, context.getFlowDescription());
        let reply = modelService.createFlowElement("SmartContract/txnreply");
        let flow = flowModel.addFlowElement(reply);
        result = result.addTriggerFlowMapping(lodash.cloneDeep(trigger), lodash.cloneDeep(flowModel));
        return flowName;
    }
    getSchemas(conId):  Observable<any> {
        let schemas = new Map();
        return Observable.create(observer => {
            WiContributionUtils.getConnection(this.http, conId)
                            .map(data => data)
                            .subscribe(data => {
                                let schemas = new Map();
                                for (let setting of data.settings) {
                                    if(setting.name === "schemas") {
                                        setting.value.map(item => schemas[item[0]] = item[1]);
                                        observer.next(schemas);
                                        break;
                                    }
                                }
                            });
                        });
    }
}