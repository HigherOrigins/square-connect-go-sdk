# CustomerCreatedWebhook

## Properties

 Name           | Type                                                             | Description                                                                                                                                            | Notes                        
----------------|------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------
 **MerchantId** | **string**                                                       | The ID of the seller associated with the event.                                                                                                        | [optional] [default to null] 
 **Type_**      | **string**                                                       | The type of event. For this object, the value is &#x60;customer.created&#x60;.                                                                         | [optional] [default to null] 
 **EventId**    | **string**                                                       | The unique ID of the event, which is used for [idempotency support](https://developer.squareup.com/docs/webhooks/step4manage#webhooks-best-practices). | [optional] [default to null] 
 **CreatedAt**  | **string**                                                       | The timestamp of when the event was created, in RFC 3339 format.                                                                                       | [optional] [default to null] 
 **Data**       | [***CustomerCreatedWebhookData**](CustomerCreatedWebhookData.md) |                                                                                                                                                        | [optional] [default to null] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

