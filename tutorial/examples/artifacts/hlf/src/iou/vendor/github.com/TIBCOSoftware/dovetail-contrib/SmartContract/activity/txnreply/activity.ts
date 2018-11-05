import {Observable} from "rxjs/Observable";
import {Injectable, Injector, Inject} from "@angular/core";
import {Http} from "@angular/http";
import {
    WiContrib,
    WiServiceHandlerContribution,
    IValidationResult,
    ValidationResult,
    IActivityContribution,
    WiContributionUtils,
    IConnectorContribution
} from "wi-studio/app/contrib/wi-contrib";


@WiContrib({})
@Injectable()
export class TxnResponseActivityContributionHandler extends WiServiceHandlerContribution {
    constructor(private injector: Injector, private http: Http,) {
        super(injector, http);
    }
   
    value = (fieldName: string, context: IActivityContribution): any | Observable<any> => {
        let conId = context.getField("model").value;
        
        switch(fieldName){
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
            case "dataType":
                if(Boolean(conId)) {
                    return Observable.create(observer => {
                        WiContributionUtils.getConnection(this.http, conId)
                                            .map(data => data)
                                            .subscribe(data => {
                                                let types = [];
                                                if(Boolean(data)){
                                                    for (let setting of data.settings) {
                                                        if (setting.name === "assets" || setting.name === "concepts") {
                                                            types = types.concat(setting.value);
                                                        }
                                                    }
                                                    types = types.concat("User Defined...");
                                                    observer.next(types);
                                                } else {
                                                    observer.next(["User Defined..."]);
                                                }
                                            });
                    });
                } else {
                    return ["User Defined..."];
                }
            case "input":
                let datatype = context.getField("dataType").value;
                if(Boolean(conId) == false || Boolean(datatype) == false || datatype == "User Defined...")
                        return null;

                let isArray = context.getField("isArray").value;
                return Observable.create(observer => {
                    this.getAssetSchemas(conId).subscribe( schemas => {
                        let schema = JSON.parse(schemas[datatype]);      
                        if(isArray && schema.type == "object"){
                            observer.next(this.createArraySchema(schema));
                        }
                        else
                            observer.next(schemas[datatype]);
                        
                    });
                });
    
            default:
                return null;
        }   
    }
 
    validate = (fieldName: string, context: IActivityContribution): Observable<IValidationResult> | IValidationResult => {
        let status = context.getField("status").value;
        let datatype = context.getField("dataType").value
        switch(fieldName) {
            case "dataType":
            case "model":
                return Observable.create(observer => {
                    let vresult: IValidationResult = ValidationResult.newValidationResult();
                    if(status == "Success with Data"){
                        vresult.setVisible(true);
                    } else {
                        vresult.setVisible(false);
                    }
                    
                    observer.next(vresult);
                });
            case "input":
            case "isArray":
                return Observable.create(observer => {
                    let vresult: IValidationResult = ValidationResult.newValidationResult();
                    if(status == "Success with Data" && datatype != "User Defined..."){
                        vresult.setVisible(true);
                    } else {
                        vresult.setVisible(false);
                    }
                    
                    observer.next(vresult);
                });
            case "userInput":
                return Observable.create(observer => {
                    let vresult: IValidationResult = ValidationResult.newValidationResult();
                    if(status == "Success with Data" && datatype == "User Defined..."){
                        vresult.setVisible(true);
                    } else {
                        vresult.setVisible(false);
                    }
                    
                    observer.next(vresult);
                });
            case "message":
                return Observable.create(observer => {
                    let vresult: IValidationResult = ValidationResult.newValidationResult();
                    if(status == "Error with Message"){
                        vresult.setVisible(true);
                    } else {
                        vresult.setVisible(false);
                    }
                    
                    observer.next(vresult);
                });
            default:
                return null;
        }
    }

    getAssetSchemas(conId):  Observable<any> {
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

    createArraySchema(schema): string {
        let newSchema = {};
        newSchema["$schema"] = schema["$schema"];
        newSchema["title"] = schema["title"];
        newSchema["type"] = "array"
        newSchema["items"] = {type: "object", properties:{}};
        newSchema["items"].properties = schema.properties;
        newSchema["description"] = schema.description;
        return JSON.stringify(newSchema);
    }
}