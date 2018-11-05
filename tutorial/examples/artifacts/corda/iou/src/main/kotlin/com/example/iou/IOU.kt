package com.example.iou

import net.corda.core.identity.AbstractParty
import java.util.Currency
import net.corda.core.serialization.CordaSerializable



@CordaSerializable
class IOU(var issuer: net.corda.core.identity.Party ,var owner: net.corda.core.identity.Party ,var amt: net.corda.core.contracts.Amount<Currency> ,override val linearId: net.corda.core.contracts.UniqueIdentifier ) : net.corda.core.contracts.LinearState {
   
    override val participants : List<AbstractParty>
        get(){ 
            val participants = ArrayList<AbstractParty>()
            	participants.add(issuer)
	participants.add(owner)
            return participants
        }

    override fun toString() : String{
         var json:String = "{"
         
         
         json = json + "\"issuer\":" + "\"" +issuer!!.toString() + "\"" + ","
         
         
         json = json + "\"owner\":" + "\"" +owner!!.toString() + "\"" + ","
         
         
         json = json + "\"amt\":" + "{\"quantity\":" + amt!!.quantity + ", \"currency\":\"" + amt!!.token.currencyCode + "\"}" + ","
         
         
         json = json + "\"linearId\":" + "\"" + linearId.toString() + "\"" + ","
         
         json = json.dropLast(1)
         json = json + "}"
         return json
    }

    override fun equals(other: Any?): Boolean  {
         if(other is IOU) {
            val to = other as (IOU)
            var isEqual: Boolean = true
        
            
                
                if (!issuer!!.equals(to.issuer))
                    return false
                
                else 
                    isEqual = true
            
        
            
                
                if (!owner!!.equals(to.owner))
                    return false
                
                else 
                    isEqual = true
            
        
            
                
                if (!amt!!.equals(to.amt))
                    return false
                
                else 
                    isEqual = true
            
        
            
                
                if (!linearId!!.equals(to.linearId))
                    return false
                
                else 
                    isEqual = true
            
        
            return isEqual
        } else {
            return false
        }
    }
}
