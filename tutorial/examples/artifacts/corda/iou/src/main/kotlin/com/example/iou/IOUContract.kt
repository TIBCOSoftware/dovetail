package com.example.iou

import com.tibco.dovetail.container.corda.CordaCommandDataWithData
import com.tibco.dovetail.container.corda.CordaFlowContract
import net.corda.core.contracts.Contract;
import net.corda.core.transactions.LedgerTransaction
import java.io.InputStream
import net.corda.core.serialization.CordaSerializable
import java.util.Currency

@CordaSerializable
public class IOUContract : CordaFlowContract(), Contract {
    companion object {
        @JvmStatic
        val IOU_CONTRACT_ID = "com.example.iou.IOUContract"
    } 
    
    @CordaSerializable
    class IssueIOU( val iou: com.example.iou.IOU  , val transactionId: String  , val timestamp: String  ) : CordaCommandDataWithData() {
        init {
            
            putData("iou", iou)
            putData("transactionId", transactionId)
            putData("timestamp", timestamp)
            putData("command", "com.example.iou.IssueIOU")
        }
    }
    
    @CordaSerializable
    class TransferIOU( val iou: com.example.iou.IOU  , val newOwner: net.corda.core.identity.Party  , val transactionId: String  , val timestamp: String  ) : CordaCommandDataWithData() {
        init {
            
            putData("iou", iou)
            putData("newOwner", newOwner)
            putData("transactionId", transactionId)
            putData("timestamp", timestamp)
            putData("command", "com.example.iou.TransferIOU")
        }
    }
    
    override fun getResourceHash(): String {
        return "705ae2d320d8ebac6fcd5a0b50d06447129f38f5eed44031b837db28da82974f"
    }

    override fun getTransactionJson(): InputStream {
        return this.javaClass.getResourceAsStream("transactions.json")
    }

    override fun verify(tx: LedgerTransaction) {
        verifyTransaction(tx)
    }
}