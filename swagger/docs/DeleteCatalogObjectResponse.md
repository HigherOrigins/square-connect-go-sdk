# DeleteCatalogObjectResponse

## Properties

 Name                 | Type                         | Description                                                                                                                                                                                                                                                      | Notes                        
----------------------|------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------
 **Errors**           | [**[]ModelError**](Error.md) | Any errors that occurred during the request.                                                                                                                                                                                                                     | [optional] [default to null] 
 **DeletedObjectIds** | **[]string**                 | The IDs of all catalog objects deleted by this request. Multiple IDs may be returned when associated objects are also deleted, for example a catalog item variation will be deleted (and its ID included in this field) when its parent catalog item is deleted. | [optional] [default to null] 
 **DeletedAt**        | **string**                   | The database [timestamp](https://developer.squareup.com/docs/build-basics/working-with-dates) of this deletion in RFC 3339 format, e.g., &#x60;2016-09-04T23:59:33.123Z&#x60;.                                                                                   | [optional] [default to null] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

