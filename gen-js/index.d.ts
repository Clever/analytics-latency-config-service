import { Span, Tracer } from "opentracing";
import { Logger } from "kayvee";

type Callback<R> = (err: Error, result: R) => void;
type ArrayInner<R> = R extends (infer T)[] ? T : never;

interface RetryPolicy {
  backoffs(): number[];
  retry(requestOptions: {method: string}, err: Error, res: {statusCode: number}): boolean;
}

interface RequestOptions {
  timeout?: number;
  span?: Span;
  retryPolicy?: RetryPolicy;
}

interface IterResult<R> {
  map<T>(f: (r: R) => T, cb?: Callback<T[]>): Promise<T[]>;
  toArray(cb?: Callback<R[]>): Promise<R[]>;
  forEach(f: (r: R) => void, cb?: Callback<void>): Promise<void>;
  forEachAsync(f: (r: R) => void, cb?: Callback<void>): Promise<void>;
}

interface CircuitOptions {
  forceClosed?: boolean;
  maxConcurrentRequests?: number;
  requestVolumeThreshold?: number;
  sleepWindow?: number;
  errorPercentThreshold?: number;
}

interface GenericOptions {
  timeout?: number;
  keepalive?: boolean;
  retryPolicy?: RetryPolicy;
  logger?: Logger;
  tracer?: Tracer;
  circuit?: CircuitOptions;
  serviceName?: string;
}

interface DiscoveryOptions {
  discovery: true;
  address?: undefined;
}

interface AddressOptions {
  discovery?: false;
  address: string;
}

type AnalyticsLatencyConfigServiceOptions = (DiscoveryOptions | AddressOptions) & GenericOptions;

import models = AnalyticsLatencyConfigService.Models

declare class AnalyticsLatencyConfigService {
  constructor(options: AnalyticsLatencyConfigServiceOptions);

  
  healthCheck(options?: RequestOptions, cb?: Callback<void>): Promise<void>
  
  getAllLegacyConfigs(options?: RequestOptions, cb?: Callback<models.AnalyticsLatencyConfigs>): Promise<models.AnalyticsLatencyConfigs>
  
}

declare namespace AnalyticsLatencyConfigService {
  const RetryPolicies: {
    Single: RetryPolicy;
    Exponential: RetryPolicy;
    None: RetryPolicy;
  }

  const DefaultCircuitOptions: CircuitOptions;

  namespace Errors {
    interface ErrorBody {
      message: string;
      [key: string]: any;
    }

    
    class BadRequest {
  code?: models.ErrorCode;
  message?: string;

  constructor(body: ErrorBody);
}
    
    class InternalError {
  code?: models.ErrorCode;
  message?: string;

  constructor(body: ErrorBody);
}
    
  }

  namespace Models {
    
    type AnalyticsDatabase = ("RedshiftProd" | "RedshiftFast" | "RdsInternal" | "RdsExternal");
    
    type AnalyticsLatencyConfigs = any;
    
    type ErrorCode = ("InvalidID");
    
    type LatencySpec = {
  thresholds?: Thresholds;
  timestampColumn?: string;
};
    
    type SchemaConfig = {
  blacklist?: string[];
  defaultThresholds?: Thresholds;
  defaultTimestampColumn?: string;
  schemaName?: string;
  schemaOwner?: string;
  tableOverrides?: TableCheck[];
  whitelist?: string[];
};
    
    type TableCheck = any;
    
    type ThresholdTier = ("Critical" | "Major" | "Minor" | "None");
    
    type Thresholds = {
  critical?: string;
  major?: string;
  minor?: string;
};
    
  }
}

export = AnalyticsLatencyConfigService;
