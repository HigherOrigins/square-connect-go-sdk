# {{classname}}

All URIs are relative to *https://connect.squareup.com*

 Method                                                   | HTTP request                        | Description      
----------------------------------------------------------|-------------------------------------|------------------
 [**ListMerchants**](MerchantsApi.md#ListMerchants)       | **Get** /v2/merchants               | ListMerchants    
 [**RetrieveMerchant**](MerchantsApi.md#RetrieveMerchant) | **Get** /v2/merchants/{merchant_id} | RetrieveMerchant 

# **ListMerchants**

> ListMerchantsResponse ListMerchants(ctx, optional)
> ListMerchants

Provides details about the merchant associated with a given access token. The access token used to connect your
application to a Square seller is associated with a single merchant. That means that `ListMerchants` returns a list with
a single `Merchant` object. You can specify your personal access token to get your own merchant information or specify
an OAuth token to get the information for the merchant that granted your application access. If you know the merchant
ID, you can also use the [RetrieveMerchant](api-endpoint:Merchants-RetrieveMerchant) endpoint to retrieve the merchant
information.

### Required Parameters

 Name         | Type                               | Description                                                                 | Notes                
--------------|------------------------------------|-----------------------------------------------------------------------------|----------------------
 **ctx**      | **context.Context**                | context for authentication, logging, cancellation, deadlines, tracing, etc. 
 **optional** | ***MerchantsApiListMerchantsOpts** | optional parameters                                                         | nil if no parameters 

### Optional Parameters

Optional parameters are passed through a pointer to a MerchantsApiListMerchantsOpts struct
Name | Type | Description | Notes
------------- | ------------- | ------------- | -------------
**cursor** | **optional.Int32**| The cursor generated by the previous response. |

### Return type

[**ListMerchantsResponse**](ListMerchantsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveMerchant**

> RetrieveMerchantResponse RetrieveMerchant(ctx, merchantId)
> RetrieveMerchant

Retrieves the `Merchant` object for the given `merchant_id`.

### Required Parameters

 Name           | Type                | Description                                                                                                                                                     | Notes 
----------------|---------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------|-------
 **ctx**        | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.                                                                                     
 **merchantId** | **string**          | The ID of the merchant to retrieve. If the string \&quot;me\&quot; is supplied as the ID, then retrieve the merchant that is currently accessible to this call. |

### Return type

[**RetrieveMerchantResponse**](RetrieveMerchantResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

