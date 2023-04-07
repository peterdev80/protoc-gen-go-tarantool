# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [demo.proto](#demo-proto)
    - [AttrValue](#api_tnt-AttrValue)
    - [GetMessageValue](#api_tnt-GetMessageValue)
    - [RequestAck](#api_tnt-RequestAck)
    - [RequestGetAttrValue](#api_tnt-RequestGetAttrValue)
    - [RequestPutAttrValue](#api_tnt-RequestPutAttrValue)
    - [ResponcePutAttrValue](#api_tnt-ResponcePutAttrValue)
    - [ResponseAck](#api_tnt-ResponseAck)
    - [ResponseGetAttrValue](#api_tnt-ResponseGetAttrValue)
  
    - [Action](#api_tnt-Action)
  
    - [BufferBusService](#api_tnt-BufferBusService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="demo-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## demo.proto
Описание сообщений передачи текущих данных из микросервисов в таратул


<a name="api_tnt-AttrValue"></a>

### AttrValue
Сообщение с текщими значениями атрибута


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| unit_guid | [string](#string) |  | уникальный индификатор оборудования |
| object_id | [string](#string) |  |  |
| attr_id | [string](#string) |  |  |
| timestamp | [int64](#int64) |  |  |
| value | [double](#double) |  |  |
| action | [Action](#api_tnt-Action) |  |  |






<a name="api_tnt-GetMessageValue"></a>

### GetMessageValue



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | уникальный индификатор оборудования |
| status | [string](#string) |  |  |
| data | [AttrValue](#api_tnt-AttrValue) |  |  |






<a name="api_tnt-RequestAck"></a>

### RequestAck



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ids | [int64](#int64) | repeated |  |






<a name="api_tnt-RequestGetAttrValue"></a>

### RequestGetAttrValue



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| batch_size | [int64](#int64) |  |  |






<a name="api_tnt-RequestPutAttrValue"></a>

### RequestPutAttrValue



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| attr_values | [AttrValue](#api_tnt-AttrValue) | repeated |  |






<a name="api_tnt-ResponcePutAttrValue"></a>

### ResponcePutAttrValue







<a name="api_tnt-ResponseAck"></a>

### ResponseAck







<a name="api_tnt-ResponseGetAttrValue"></a>

### ResponseGetAttrValue



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| attr_values | [GetMessageValue](#api_tnt-GetMessageValue) | repeated |  |





 


<a name="api_tnt-Action"></a>

### Action


| Name | Number | Description |
| ---- | ------ | ----------- |
| ACTION_ACTION_UNKNOW_UNSPECIFIED | 0 |  |
| ACTION_ACTION_CLICKHOUSE | 1 |  |


 

 


<a name="api_tnt-BufferBusService"></a>

### BufferBusService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Put | [RequestPutAttrValue](#api_tnt-RequestPutAttrValue) | [ResponcePutAttrValue](#api_tnt-ResponcePutAttrValue) |  |
| Get | [RequestGetAttrValue](#api_tnt-RequestGetAttrValue) | [ResponseGetAttrValue](#api_tnt-ResponseGetAttrValue) |  |
| Ack | [RequestAck](#api_tnt-RequestAck) | [ResponseAck](#api_tnt-ResponseAck) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

