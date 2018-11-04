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

import * as lodash from "lodash";

const keyschema = {
    "type": "object",
    "$schema": "http://json-schema.org/draft-07/schema#",
    "properties": {}
  };

const arraykeyschema = {
       "type": "array", 
       "$schema": "http://json-schema.org/draft-07/schema#", 
       "items": {
         "type": "object", 
         "properties": {}
      }
    };

@WiContrib({})
@Injectable()
export class LedgerActivityContributionHandler extends WiServiceHandlerContribution {
    constructor(private injector: Injector, private http: Http,) {
        super(injector, http);
    }
   
    value = (fieldName: string, context: IActivityContribution): any | Observable<any> => {
        let action = context.getField("operation").value;
        let asset = context.getField("assetName").value;
        let isArray = false;
        if(Boolean(context.getField("isArray"))){
            isArray = context.getField("isArray").value;
        }
        
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
                if(Boolean(conId) == false || Boolean(asset) == false)
                    return null;

                return Observable.create(observer => {
                    this.getSchemas(conId).subscribe( schemas => {
                        let schema = JSON.parse(schemas[asset]);
                        let metadata = JSON.parse(schema.description);
                        let assetIdentifier = metadata.metadata.identifiedBy;
                        if(action === "DELETE" || action == "GET") {
                            observer.next(this.createKeySchema(assetIdentifier, metadata, action, isArray));
                        } else if (action == "LOOKUP"){
                            assetIdentifier = context.getField("compositeKey").value
                            if(Boolean(assetIdentifier)){
                                let lastidx = assetIdentifier.lastIndexOf(",", assetIdentifier.length-1);
                                if(lastidx != -1)
                                    assetIdentifier = assetIdentifier.substring(0, lastidx)
                                    
                                observer.next(this.createKeySchema(assetIdentifier, metadata, action, isArray));
                            } else {
                                observer.next(null);
                            }
                        }
                        else {
                            if(isArray && schema.type == "object"){
                                observer.next(this.createArraySchema(schema));
                            }
                            else
                                observer.next(schemas[asset]);
                        }
                    });
                });
            case "identifier":
                if(Boolean(conId) == false || Boolean(asset) == false)
                    return null;
                return Observable.create(observer => {
                    this.getSchemas(conId).subscribe( schemas => {
                        let schema = JSON.parse(schemas[asset])
                        let metadata = JSON.parse(schema.description)
                        let assetIdentifier = metadata.metadata.identifiedBy
                        observer.next(assetIdentifier);
                    });
                });
            case "compositeKey":
                if(Boolean(conId) == false || Boolean(asset) == false)
                        return null;

                    return Observable.create(observer => {
                        this.getSchemas(conId).subscribe( schemas => {
                            let schema = JSON.parse(schemas[asset]);
                            let metaschema = JSON.parse(schema.description);
                            let keys = [];
                            if(metaschema.metadata.isPrimaryComposite)
                                keys.push(metaschema.metadata.identifiedBy);

                            //secondary composite keys
                            if(Boolean(metaschema.metadata.compositeKeys))
                                keys = keys.concat(metaschema.metadata.compositeKeys.split("|"));
                            
                            if(action === "LOOKUP") {
                                observer.next(keys);
                            } else {
                                observer.next(null);
                            }
                            
                        });
                    });
            case "compositeKeys":
                if(Boolean(conId) == false || Boolean(asset) == false)
                return null;

                return Observable.create(observer => {
                    this.getSchemas(conId).subscribe( schemas => {
                        let schema = JSON.parse(schemas[asset]);
                        let metaschema = JSON.parse(schema.description);
                        let keys = "";
                        
                        if(Boolean(metaschema.metadata.compositeKeys))
                            keys = metaschema.metadata.compositeKeys;
                        
                        observer.next(keys);
                    });
                });
            case "output":
                if(Boolean(conId) == false || Boolean(asset) == false)
                    return null;

                return Observable.create(observer => {
                    this.getSchemas(conId).subscribe( schemas => {
                        if(isArray || action == "LOOKUP") {
                            let schema = JSON.parse(schemas[asset]);
                            if(schema.type == "object"){         
                                observer.next(this.createArraySchema(schema));
                            }
                            else
                                observer.next(schemas[asset]);
                        } else {
                            observer.next(schemas[asset]);
                        } 
                    });
                });
            default:
                return null;
        }
    }
 
    validate = (fieldName: string, context: IActivityContribution): Observable<IValidationResult> | IValidationResult => {
        switch(fieldName){
            case "model":
                return Observable.create(observer => {
                    let vresult: IValidationResult = ValidationResult.newValidationResult();
                   // vresult.setReadOnly(Boolean(localStorage.getItem(context.appId)));
                    observer.next(vresult);
                });
            case "input":
                return Observable.create(observer => {
                    let vresult: IValidationResult = ValidationResult.newValidationResult();
                    vresult.setReadOnly(true);
                    observer.next(vresult);
                });
            case "identifier":
            case "compositeKeys":
                return Observable.create(observer => {
                    let vresult: IValidationResult = ValidationResult.newValidationResult();
                    vresult.setVisible(false);
                    observer.next(vresult);
                });
            case "isArray":
                return Observable.create(observer => {
                    let vresult: IValidationResult = ValidationResult.newValidationResult();
                   
                    vresult.setVisible(true);
                    observer.next(vresult);
                });
            case "compositeKey":
                return Observable.create(observer => {
                    let vresult: IValidationResult = ValidationResult.newValidationResult();
                    if (context.getField("operation").value === "LOOKUP")
                        vresult.setVisible(true);
                    else
                        vresult.setVisible(false);
                    observer.next(vresult);
                });
            default:
                return null; 
        }
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

    createKeySchema(key, metadata, action, isArray): string {
        let identifiers = key.split(",");

        let p = {};
        identifiers.forEach( id => {
            id = id.trim();
            for(let i=0; i<metadata.attributes.length;i++){
                if(metadata.attributes[i].name === id){
                    let datatype = metadata.attributes[i].type;
                    if(datatype === "Integer" || datatype === "Long")
                        p[id] = {type:"integer"};
                    else if(datatype === "Boolean")
                        p[id] = {type: "boolean"};
                    else 
                        p[id] = {type: "string"};

                    break;
                }
            }
        });

        let idschema 
        if(isArray){
          idschema = lodash.cloneDeep(arraykeyschema);
          idschema.items.properties = p;
        }
        else{
           idschema = lodash.cloneDeep(keyschema);
           idschema.properties = p;
        }
        if(action != "LOOKUP")
            idschema["required"] = identifiers;

        return JSON.stringify(idschema);
    }
}