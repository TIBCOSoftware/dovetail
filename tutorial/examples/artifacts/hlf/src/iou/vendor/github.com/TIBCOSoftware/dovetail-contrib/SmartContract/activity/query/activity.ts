import {Observable} from "rxjs/Observable";
import {Injectable, Injector, Inject} from "@angular/core";
import {Http} from "@angular/http";
import {
    WiContrib,
    WiServiceHandlerContribution,
    IValidationResult,
    IActivityContribution,
    WiContributionUtils,
    IConnectorContribution
} from "wi-studio/app/contrib/wi-contrib";

import * as lodash from "lodash";

@WiContrib({})
@Injectable()
export class QueryActivityContributionHandler extends WiServiceHandlerContribution {
    constructor(private injector: Injector, private http: Http,) {
        super(injector, http);
    }
   
    value = (fieldName: string, context: IActivityContribution): any | Observable<any> => {
        let asset = context.getField("assetName").value;
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

            case "assetName":
                if(Boolean(conId) == false)
                    return null;

                return Observable.create(observer => {
                    WiContributionUtils.getConnection(this.http, conId)
                                        .map(data => data)
                                        .subscribe(data => {
                                            for (let setting of data.settings) {
                                                if (setting.name === "assets") {
                                                    observer.next(setting.value);
                                                    break;
                                                }
                                            }
                                        });
                });
            case "input":
                let field = context.getField("params").value;
                if(Boolean(field)){
                    let gridData = field.value;
                    if(Boolean(gridData)){
                        let schema = {$schema: "http://json-schema.org/draft-04/schema#", type: "object", properties: {}}
                        let params = JSON.parse(gridData)
                        for (let i=0; i<params.length; i++)
                            schema.properties[params[i].paramName] = {type: params[i].type}

                        return JSON.stringify(schema);
                    } else {
                        return null;
                    }
                } else {
                    return null;
                } 
            case "output":
                if(Boolean(conId) == false || Boolean(asset) == false)
                    return null;

                return Observable.create(observer => {
                    this.getSchemas(conId).subscribe( schemas => {
                        let schema = JSON.parse(schemas[asset]);
                        
                        if(schema.type == "object"){
                            let newSchema = {};
                            newSchema["$schema"] = schema["$schema"];
                            newSchema["title"] = schema["title"];
                            newSchema["type"] = "array"
                            newSchema["items"] = {type: "object", properties:{}};
                            newSchema["items"].properties = schema.properties;
                            newSchema["description"] = schema.description;
                            observer.next(JSON.stringify(newSchema));
                        }
                        else {
                            observer.next(schemas[asset]);
                        }
                    });
                });
            default:
                return null;
        }
    }
 
    validate = (fieldName: string, context: IActivityContribution): Observable<IValidationResult> | IValidationResult => {
        return null; 
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