# SessionSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Login** | Pointer to **string** |  | [optional] 
**Useragent** | Pointer to **string** |  | [optional] 

## Methods

### NewSessionSession

`func NewSessionSession() *SessionSession`

NewSessionSession instantiates a new SessionSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionSessionWithDefaults

`func NewSessionSessionWithDefaults() *SessionSession`

NewSessionSessionWithDefaults instantiates a new SessionSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogin

`func (o *SessionSession) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *SessionSession) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *SessionSession) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *SessionSession) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetUseragent

`func (o *SessionSession) GetUseragent() string`

GetUseragent returns the Useragent field if non-nil, zero value otherwise.

### GetUseragentOk

`func (o *SessionSession) GetUseragentOk() (*string, bool)`

GetUseragentOk returns a tuple with the Useragent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseragent

`func (o *SessionSession) SetUseragent(v string)`

SetUseragent sets Useragent field to given value.

### HasUseragent

`func (o *SessionSession) HasUseragent() bool`

HasUseragent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


