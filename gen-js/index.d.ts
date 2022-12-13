import { Logger } from "kayvee";

type Callback<R> = (err: Error, result: R) => void;
type ArrayInner<R> = R extends (infer T)[] ? T : never;

interface RetryPolicy {
  backoffs(): number[];
  retry(requestOptions: {method: string}, err: Error, res: {statusCode: number}): boolean;
}

interface RequestOptions {
  timeout?: number;
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

  close(): void;
  
  healthCheck(options?: RequestOptions, cb?: Callback<void>): Promise<void>
  
  getTableLatency(request: models.GetTableLatencyRequest, options?: RequestOptions, cb?: Callback<models.GetTableLatencyResponse>): Promise<models.GetTableLatencyResponse>
  
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
    
    class NotFound {
  code?: models.ErrorCode;
  message?: string;

  constructor(body: ErrorBody);
}
    
  }

  namespace Models {
    
    type AnalyticsDatabase = ("RedshiftFast" | "RdsInternal" | "RdsExternal" | "Snowflake");
    
    type AnalyticsLatencyConfigs = any;
    
    type ErrorCode = ("InvalidID");
    
    type GetTableLatencyRequest = {
  database: AnalyticsDatabase;
  schema: string;
  table: string;
};
    
    type GetTableLatencyResponse = {
  database: AnalyticsDatabase;
  latency?: number;
  owner: string;
  schema: string;
  table: string;
  thresholds: Thresholds;
};
    
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
    
    type ThresholdTier = ("Critical" | "Major" | "Minor" | "Refresh" | "None");
    
    type Thresholds = {
  critical?: string;
  major?: string;
  minor?: string;
  refresh?: string;
};
    
  }
}

export = AnalyticsLatencyConfigService;
