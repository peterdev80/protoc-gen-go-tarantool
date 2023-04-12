# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [asu.proto](#asu-proto)
    - [PlcAttribute](#api_tnt-PlcAttribute)
    - [RequestPutAttributePlc](#api_tnt-RequestPutAttributePlc)
    - [TableAsu](#api_tnt-TableAsu)
  
- [history.proto](#history-proto)
    - [Attribute](#api_tnt-Attribute)
    - [QueueMessage](#api_tnt-QueueMessage)
    - [RequestAck](#api_tnt-RequestAck)
    - [RequestListAttributeBus](#api_tnt-RequestListAttributeBus)
    - [RequestPutAttributeBus](#api_tnt-RequestPutAttributeBus)
    - [ResponseAck](#api_tnt-ResponseAck)
    - [ResponseListAttributeBus](#api_tnt-ResponseListAttributeBus)
    - [ResponsePutAttributeBus](#api_tnt-ResponsePutAttributeBus)
  
    - [Action](#api_tnt-Action)
  
    - [BufferBusService](#api_tnt-BufferBusService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="asu-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## asu.proto
Описание сообщений передачи сервисом асу


<a name="api_tnt-PlcAttribute"></a>

### PlcAttribute



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| plc_id | [string](#string) |  | уникальный идентификатор оборудования |
| tag_id | [string](#string) |  |  |
| id | [string](#string) |  | идентификатор атрибута |
| timestamp | [int64](#int64) |  | unix time времени |
| value | [double](#double) |  | значение атрибута |
| action | [Action](#api_tnt-Action) |  | действия применяемые к атрибуту |






<a name="api_tnt-RequestPutAttributePlc"></a>

### RequestPutAttributePlc







<a name="api_tnt-TableAsu"></a>

### TableAsu
таблица соответствия plc unitguid


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| plc_id | [string](#string) |  |  |
| tag_id | [string](#string) |  |  |
| unit_guid | [string](#string) |  |  |
| object_id | [string](#string) |  |  |





 

 

 

 



<a name="history-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## history.proto
Описание сообщений передачи текущих данных из микросервисов в таратул


<a name="api_tnt-Attribute"></a>

### Attribute
Сообщение с текущими значениями атрибута


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| unit_guid | [string](#string) |  | уникальный идентификатор оборудования |
| object_id | [string](#string) |  |  |
| id | [string](#string) |  | идентификатор атрибута |
| timestamp | [int64](#int64) |  | unix time времени |
| value | [double](#double) |  | значение атрибута |
| actions | [Action](#api_tnt-Action) | repeated | действия применяемые к атрибуту |






<a name="api_tnt-QueueMessage"></a>

### QueueMessage
сообщение получаемое из очереди


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | уникальный идентификатор сообщения (позиция в очереди) |
| status | [string](#string) |  | статус сообщения |
| data | [Attribute](#api_tnt-Attribute) |  | атрибут |






<a name="api_tnt-RequestAck"></a>

### RequestAck
сообщение на подтверждение обработки message


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ids | [int64](#int64) | repeated | список уникальных идентификатор сообщения |






<a name="api_tnt-RequestListAttributeBus"></a>

### RequestListAttributeBus
запрос атрибут из шины


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| batch_size | [int64](#int64) |  | количество атрибутов в ответе |






<a name="api_tnt-RequestPutAttributeBus"></a>

### RequestPutAttributeBus
отправить аттрибут в шину


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| attrs | [Attribute](#api_tnt-Attribute) | repeated | пакет атрибутов |






<a name="api_tnt-ResponseAck"></a>

### ResponseAck







<a name="api_tnt-ResponseListAttributeBus"></a>

### ResponseListAttributeBus
ответ на запрос атрибут


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| messages | [QueueMessage](#api_tnt-QueueMessage) | repeated | список атрибут |






<a name="api_tnt-ResponsePutAttributeBus"></a>

### ResponsePutAttributeBus
результат отправки атрибута в шину





 


<a name="api_tnt-Action"></a>

### Action
действие которые необходимо совершить с сообщением

| Name | Number | Description |
| ---- | ------ | ----------- |
| ACTION_ACTION_UNKNOWN_UNSPECIFIED | 0 | не определено |
| ACTION_ACTION_CLICKHOUSE | 1 | передавать в clickhouse |


 

 


<a name="api_tnt-BufferBusService"></a>

### BufferBusService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| PutAttributeBus | [RequestPutAttributeBus](#api_tnt-RequestPutAttributeBus) | [ResponsePutAttributeBus](#api_tnt-ResponsePutAttributeBus) |  |
| ListAttributeBus | [RequestListAttributeBus](#api_tnt-RequestListAttributeBus) | [ResponseListAttributeBus](#api_tnt-ResponseListAttributeBus) |  |
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

