// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

/**
 * Properties for the `del` method.
 */
export type DelProps = {
  /**
   * Path of the endpoint, anything after \"/1\" must be specified.
   */
  path: string;
  /**
   * Query parameters to apply to the current query.
   */
  parameters?: Record<string, any>;
};

/**
 * Properties for the `deleteABTest` method.
 */
export type DeleteABTestProps = {
  /**
   * Unique A/B test ID.
   */
  id: number;
};

/**
 * Properties for the `get` method.
 */
export type GetProps = {
  /**
   * Path of the endpoint, anything after \"/1\" must be specified.
   */
  path: string;
  /**
   * Query parameters to apply to the current query.
   */
  parameters?: Record<string, any>;
};

/**
 * Properties for the `getABTest` method.
 */
export type GetABTestProps = {
  /**
   * Unique A/B test ID.
   */
  id: number;
};

/**
 * Properties for the `listABTests` method.
 */
export type ListABTestsProps = {
  /**
   * Position of the starting record. Used for paging. 0 is the first record.
   */
  offset?: number;
  /**
   * Number of records to return (page size).
   */
  limit?: number;
  /**
   * Only return A/B tests for indices starting with this prefix.
   */
  indexPrefix?: string;
  /**
   * Only return A/B tests for indices ending with this suffix.
   */
  indexSuffix?: string;
};

/**
 * Properties for the `post` method.
 */
export type PostProps = {
  /**
   * Path of the endpoint, anything after \"/1\" must be specified.
   */
  path: string;
  /**
   * Query parameters to apply to the current query.
   */
  parameters?: Record<string, any>;
  /**
   * Parameters to send with the custom request.
   */
  body?: Record<string, any>;
};

/**
 * Properties for the `put` method.
 */
export type PutProps = {
  /**
   * Path of the endpoint, anything after \"/1\" must be specified.
   */
  path: string;
  /**
   * Query parameters to apply to the current query.
   */
  parameters?: Record<string, any>;
  /**
   * Parameters to send with the custom request.
   */
  body?: Record<string, any>;
};

/**
 * Properties for the `stopABTest` method.
 */
export type StopABTestProps = {
  /**
   * Unique A/B test ID.
   */
  id: number;
};
