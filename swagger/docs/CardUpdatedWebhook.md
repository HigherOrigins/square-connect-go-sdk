# CardUpdatedWebhook

## Properties

 Name           | Type                                                     | Description                                                                  | Notes                        
----------------|----------------------------------------------------------|------------------------------------------------------------------------------|------------------------------
 **MerchantId** | **string**                                               | The ID of the target seller associated with the event.                       | [optional] [default to null] 
 **Type_**      | **string**                                               | The type of this event. The value is &#x60;\&quot;card.updated\&quot;&#x60;. | [optional] [default to null] 
 **EventId**    | **string**                                               | A unique ID for the webhook event.                                           | [optional] [default to null] 
 **CreatedAt**  | **string**                                               | The timestamp of when the webhook event was created, in RFC 3339 format.     | [optional] [default to null] 
 **Data**       | [***CardUpdatedWebhookData**](CardUpdatedWebhookData.md) |                                                                              | [optional] [default to null] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

