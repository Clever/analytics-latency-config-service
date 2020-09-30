<a name="module_analytics-latency-config-service"></a>

## analytics-latency-config-service
analytics-latency-config-service client library.


* [analytics-latency-config-service](#module_analytics-latency-config-service)
    * [AnalyticsLatencyConfigService](#exp_module_analytics-latency-config-service--AnalyticsLatencyConfigService) ⏏
        * [new AnalyticsLatencyConfigService(options)](#new_module_analytics-latency-config-service--AnalyticsLatencyConfigService_new)
        * _instance_
            * [.healthCheck([options], [cb])](#module_analytics-latency-config-service--AnalyticsLatencyConfigService+healthCheck) ⇒ <code>Promise</code>
            * [.getTableLatency(request, [options], [cb])](#module_analytics-latency-config-service--AnalyticsLatencyConfigService+getTableLatency) ⇒ <code>Promise</code>
            * [.getAllLegacyConfigs([options], [cb])](#module_analytics-latency-config-service--AnalyticsLatencyConfigService+getAllLegacyConfigs) ⇒ <code>Promise</code>
        * _static_
            * [.RetryPolicies](#module_analytics-latency-config-service--AnalyticsLatencyConfigService.RetryPolicies)
                * [.Exponential](#module_analytics-latency-config-service--AnalyticsLatencyConfigService.RetryPolicies.Exponential)
                * [.Single](#module_analytics-latency-config-service--AnalyticsLatencyConfigService.RetryPolicies.Single)
                * [.None](#module_analytics-latency-config-service--AnalyticsLatencyConfigService.RetryPolicies.None)
            * [.Errors](#module_analytics-latency-config-service--AnalyticsLatencyConfigService.Errors)
                * [.BadRequest](#module_analytics-latency-config-service--AnalyticsLatencyConfigService.Errors.BadRequest) ⇐ <code>Error</code>
                * [.InternalError](#module_analytics-latency-config-service--AnalyticsLatencyConfigService.Errors.InternalError) ⇐ <code>Error</code>
                * [.NotFound](#module_analytics-latency-config-service--AnalyticsLatencyConfigService.Errors.NotFound) ⇐ <code>Error</code>
            * [.DefaultCircuitOptions](#module_analytics-latency-config-service--AnalyticsLatencyConfigService.DefaultCircuitOptions)

<a name="exp_module_analytics-latency-config-service--AnalyticsLatencyConfigService"></a>

### AnalyticsLatencyConfigService ⏏
analytics-latency-config-service client

**Kind**: Exported class  
<a name="new_module_analytics-latency-config-service--AnalyticsLatencyConfigService_new"></a>

#### new AnalyticsLatencyConfigService(options)
Create a new client object.


| Param | Type | Default | Description |
| --- | --- | --- | --- |
| options | <code>Object</code> |  | Options for constructing a client object. |
| [options.address] | <code>string</code> |  | URL where the server is located. Must provide this or the discovery argument |
| [options.discovery] | <code>bool</code> |  | Use clever-discovery to locate the server. Must provide this or the address argument |
| [options.timeout] | <code>number</code> |  | The timeout to use for all client requests, in milliseconds. This can be overridden on a per-request basis. Default is 5000ms. |
| [options.keepalive] | <code>bool</code> |  | Set keepalive to true for client requests. This sets the forever: true attribute in request. Defaults to true. |
| [options.retryPolicy] | [<code>RetryPolicies</code>](#module_analytics-latency-config-service--AnalyticsLatencyConfigService.RetryPolicies) | <code>RetryPolicies.Single</code> | The logic to determine which requests to retry, as well as how many times to retry. |
| [options.logger] | <code>module:kayvee.Logger</code> | <code>logger.New(&quot;analytics-latency-config-service-wagclient&quot;)</code> | The Kayvee logger to use in the client. |
| [options.circuit] | <code>Object</code> |  | Options for constructing the client's circuit breaker. |
| [options.circuit.forceClosed] | <code>bool</code> |  | When set to true the circuit will always be closed. Default: true. |
| [options.circuit.maxConcurrentRequests] | <code>number</code> |  | the maximum number of concurrent requests the client can make at the same time. Default: 100. |
| [options.circuit.requestVolumeThreshold] | <code>number</code> |  | The minimum number of requests needed before a circuit can be tripped due to health. Default: 20. |
| [options.circuit.sleepWindow] | <code>number</code> |  | how long, in milliseconds, to wait after a circuit opens before testing for recovery. Default: 5000. |
| [options.circuit.errorPercentThreshold] | <code>number</code> |  | the threshold to place on the rolling error rate. Once the error rate exceeds this percentage, the circuit opens. Default: 90. |

<a name="module_analytics-latency-config-service--AnalyticsLatencyConfigService+healthCheck"></a>

#### analyticsLatencyConfigService.healthCheck([options], [cb]) ⇒ <code>Promise</code>
Checks if the service is healthy

**Kind**: instance method of [<code>AnalyticsLatencyConfigService</code>](#exp_module_analytics-latency-config-service--AnalyticsLatencyConfigService)  
**Fulfill**: <code>undefined</code>  
**Reject**: [<code>BadRequest</code>](#module_analytics-latency-config-service--AnalyticsLatencyConfigService.Errors.BadRequest)  
**Reject**: [<code>InternalError</code>](#module_analytics-latency-config-service--AnalyticsLatencyConfigService.Errors.InternalError)  
**Reject**: <code>Error</code>  

| Param | Type | Description |
| --- | --- | --- |
| [options] | <code>object</code> |  |
| [options.timeout] | <code>number</code> | A request specific timeout |
| [options.span] | [<code>Span</code>](https://doc.esdoc.org/github.com/opentracing/opentracing-javascript/class/src/span.js~Span.html) | An OpenTracing span - For example from the parent request |
| [options.retryPolicy] | [<code>RetryPolicies</code>](#module_analytics-latency-config-service--AnalyticsLatencyConfigService.RetryPolicies) | A request specific retryPolicy |
| [cb] | <code>function</code> |  |

<a name="module_analytics-latency-config-service--AnalyticsLatencyConfigService+getTableLatency"></a>

#### analyticsLatencyConfigService.getTableLatency(request, [options], [cb]) ⇒ <code>Promise</code>
**Kind**: instance method of [<code>AnalyticsLatencyConfigService</code>](#exp_module_analytics-latency-config-service--AnalyticsLatencyConfigService)  
**Fulfill**: <code>Object</code>  
**Reject**: [<code>BadRequest</code>](#module_analytics-latency-config-service--AnalyticsLatencyConfigService.Errors.BadRequest)  
**Reject**: [<code>NotFound</code>](#module_analytics-latency-config-service--AnalyticsLatencyConfigService.Errors.NotFound)  
**Reject**: [<code>InternalError</code>](#module_analytics-latency-config-service--AnalyticsLatencyConfigService.Errors.InternalError)  
**Reject**: <code>Error</code>  

| Param | Type | Description |
| --- | --- | --- |
| request |  |  |
| [options] | <code>object</code> |  |
| [options.timeout] | <code>number</code> | A request specific timeout |
| [options.span] | [<code>Span</code>](https://doc.esdoc.org/github.com/opentracing/opentracing-javascript/class/src/span.js~Span.html) | An OpenTracing span - For example from the parent request |
| [options.retryPolicy] | [<code>RetryPolicies</code>](#module_analytics-latency-config-service--AnalyticsLatencyConfigService.RetryPolicies) | A request specific retryPolicy |
| [cb] | <code>function</code> |  |

<a name="module_analytics-latency-config-service--AnalyticsLatencyConfigService+getAllLegacyConfigs"></a>

#### analyticsLatencyConfigService.getAllLegacyConfigs([options], [cb]) ⇒ <code>Promise</code>
**Kind**: instance method of [<code>AnalyticsLatencyConfigService</code>](#exp_module_analytics-latency-config-service--AnalyticsLatencyConfigService)  
**Fulfill**: <code>Object</code>  
**Reject**: [<code>BadRequest</code>](#module_analytics-latency-config-service--AnalyticsLatencyConfigService.Errors.BadRequest)  
**Reject**: [<code>InternalError</code>](#module_analytics-latency-config-service--AnalyticsLatencyConfigService.Errors.InternalError)  
**Reject**: <code>Error</code>  

| Param | Type | Description |
| --- | --- | --- |
| [options] | <code>object</code> |  |
| [options.timeout] | <code>number</code> | A request specific timeout |
| [options.span] | [<code>Span</code>](https://doc.esdoc.org/github.com/opentracing/opentracing-javascript/class/src/span.js~Span.html) | An OpenTracing span - For example from the parent request |
| [options.retryPolicy] | [<code>RetryPolicies</code>](#module_analytics-latency-config-service--AnalyticsLatencyConfigService.RetryPolicies) | A request specific retryPolicy |
| [cb] | <code>function</code> |  |

<a name="module_analytics-latency-config-service--AnalyticsLatencyConfigService.RetryPolicies"></a>

#### AnalyticsLatencyConfigService.RetryPolicies
Retry policies available to use.

**Kind**: static property of [<code>AnalyticsLatencyConfigService</code>](#exp_module_analytics-latency-config-service--AnalyticsLatencyConfigService)  

* [.RetryPolicies](#module_analytics-latency-config-service--AnalyticsLatencyConfigService.RetryPolicies)
    * [.Exponential](#module_analytics-latency-config-service--AnalyticsLatencyConfigService.RetryPolicies.Exponential)
    * [.Single](#module_analytics-latency-config-service--AnalyticsLatencyConfigService.RetryPolicies.Single)
    * [.None](#module_analytics-latency-config-service--AnalyticsLatencyConfigService.RetryPolicies.None)

<a name="module_analytics-latency-config-service--AnalyticsLatencyConfigService.RetryPolicies.Exponential"></a>

##### RetryPolicies.Exponential
The exponential retry policy will retry five times with an exponential backoff.

**Kind**: static constant of [<code>RetryPolicies</code>](#module_analytics-latency-config-service--AnalyticsLatencyConfigService.RetryPolicies)  
<a name="module_analytics-latency-config-service--AnalyticsLatencyConfigService.RetryPolicies.Single"></a>

##### RetryPolicies.Single
Use this retry policy to retry a request once.

**Kind**: static constant of [<code>RetryPolicies</code>](#module_analytics-latency-config-service--AnalyticsLatencyConfigService.RetryPolicies)  
<a name="module_analytics-latency-config-service--AnalyticsLatencyConfigService.RetryPolicies.None"></a>

##### RetryPolicies.None
Use this retry policy to turn off retries.

**Kind**: static constant of [<code>RetryPolicies</code>](#module_analytics-latency-config-service--AnalyticsLatencyConfigService.RetryPolicies)  
<a name="module_analytics-latency-config-service--AnalyticsLatencyConfigService.Errors"></a>

#### AnalyticsLatencyConfigService.Errors
Errors returned by methods.

**Kind**: static property of [<code>AnalyticsLatencyConfigService</code>](#exp_module_analytics-latency-config-service--AnalyticsLatencyConfigService)  

* [.Errors](#module_analytics-latency-config-service--AnalyticsLatencyConfigService.Errors)
    * [.BadRequest](#module_analytics-latency-config-service--AnalyticsLatencyConfigService.Errors.BadRequest) ⇐ <code>Error</code>
    * [.InternalError](#module_analytics-latency-config-service--AnalyticsLatencyConfigService.Errors.InternalError) ⇐ <code>Error</code>
    * [.NotFound](#module_analytics-latency-config-service--AnalyticsLatencyConfigService.Errors.NotFound) ⇐ <code>Error</code>

<a name="module_analytics-latency-config-service--AnalyticsLatencyConfigService.Errors.BadRequest"></a>

##### Errors.BadRequest ⇐ <code>Error</code>
BadRequest

**Kind**: static class of [<code>Errors</code>](#module_analytics-latency-config-service--AnalyticsLatencyConfigService.Errors)  
**Extends**: <code>Error</code>  
**Properties**

| Name | Type |
| --- | --- |
| code |  | 
| message | <code>string</code> | 

<a name="module_analytics-latency-config-service--AnalyticsLatencyConfigService.Errors.InternalError"></a>

##### Errors.InternalError ⇐ <code>Error</code>
InternalError

**Kind**: static class of [<code>Errors</code>](#module_analytics-latency-config-service--AnalyticsLatencyConfigService.Errors)  
**Extends**: <code>Error</code>  
**Properties**

| Name | Type |
| --- | --- |
| code |  | 
| message | <code>string</code> | 

<a name="module_analytics-latency-config-service--AnalyticsLatencyConfigService.Errors.NotFound"></a>

##### Errors.NotFound ⇐ <code>Error</code>
NotFound

**Kind**: static class of [<code>Errors</code>](#module_analytics-latency-config-service--AnalyticsLatencyConfigService.Errors)  
**Extends**: <code>Error</code>  
**Properties**

| Name | Type |
| --- | --- |
| code |  | 
| message | <code>string</code> | 

<a name="module_analytics-latency-config-service--AnalyticsLatencyConfigService.DefaultCircuitOptions"></a>

#### AnalyticsLatencyConfigService.DefaultCircuitOptions
Default circuit breaker options.

**Kind**: static constant of [<code>AnalyticsLatencyConfigService</code>](#exp_module_analytics-latency-config-service--AnalyticsLatencyConfigService)  
