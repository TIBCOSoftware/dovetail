import {NgModule} from "@angular/core"
import {HttpModule} from "@angular/http";
import {WiServiceContribution} from "wi-studio/app/contrib/wi-contrib"
import {SmartContractTriggerHandler} from "./trigger"

@NgModule({
    imports: [
        HttpModule
    ],
    providers: [
        {
            provide: WiServiceContribution,
            useClass: SmartContractTriggerHandler
        }
    ]
})
export default class SmartContractTriggerModule {

}
