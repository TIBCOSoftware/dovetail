import {Observable} from "rxjs/Observable";
import {Injectable, Injector, Inject} from "@angular/core";
import {Http} from "@angular/http";
import {
    WiContrib,
    WiServiceHandlerContribution,
    IValidationResult,
    ValidationResult,
    IFieldDefinition,
    IActivityContribution,
    ActionResult,
    IActionResult,
    WiContributionUtils,
    WiContribModelService,
    IConnectorContribution
} from "wi-studio/app/contrib/wi-contrib";


@WiContrib({})
@Injectable()
export class LoggerActivityContributionHandler extends WiServiceHandlerContribution {
  /*  constructor(private injector: Injector, private http: Http,) {
        super(injector, http);
    }
   */

  constructor(private injector: Injector, private http: Http) {
    super(injector, http);
}
    value = (fieldName: string, context: IActivityContribution): any | Observable<any> => {
        return null;
    }
 
    validate = (fieldName: string, context: IActivityContribution): Observable<IValidationResult> | IValidationResult => {
        let level = context.getField("logLevel").value;
        switch(fieldName) {
            case "errorCode":
                return Observable.create(observer => {
                    let vresult: IValidationResult = ValidationResult.newValidationResult();
                    if(level == "ERROR"){
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

}