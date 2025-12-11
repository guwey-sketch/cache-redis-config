export interface RedisConfig {
  host: string;
  port: number;
  password?: string;
  db?: number;
  ttl?: number;
}

export interface CacheConfig {
  redis: RedisConfig;
  defaultTtl: number;
  prefix: string;
}

export interface CacheOptions {
  key: string;
  ttl?: number;
}

export interface CacheResult<T> {
  success: boolean;
  data?: T;
  error?: Error;
}