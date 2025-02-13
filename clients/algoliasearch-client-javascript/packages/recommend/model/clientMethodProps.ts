// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

import type { RecommendModels } from './recommendModels';
import type { SearchRecommendRulesParams } from './searchRecommendRulesParams';

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
 * Properties for the `deleteRecommendRule` method.
 */
export type DeleteRecommendRuleProps = {
  /**
   * Index on which to perform the request.
   */
  indexName: string;
  /**
   * [Recommend models](https://www.algolia.com/doc/guides/algolia-recommend/overview/#recommend-models).
   */
  model: RecommendModels;
  /**
   * Unique record (object) identifier.
   */
  objectID: string;
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
 * Properties for the `getRecommendRule` method.
 */
export type GetRecommendRuleProps = {
  /**
   * Index on which to perform the request.
   */
  indexName: string;
  /**
   * [Recommend models](https://www.algolia.com/doc/guides/algolia-recommend/overview/#recommend-models).
   */
  model: RecommendModels;
  /**
   * Unique record (object) identifier.
   */
  objectID: string;
};

/**
 * Properties for the `getRecommendStatus` method.
 */
export type GetRecommendStatusProps = {
  /**
   * Index on which to perform the request.
   */
  indexName: string;
  /**
   * [Recommend models](https://www.algolia.com/doc/guides/algolia-recommend/overview/#recommend-models).
   */
  model: RecommendModels;
  /**
   * Unique identifier of a task. Numeric value (up to 64bits).
   */
  taskID: number;
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
 * Properties for the `searchRecommendRules` method.
 */
export type SearchRecommendRulesProps = {
  /**
   * Index on which to perform the request.
   */
  indexName: string;
  /**
   * [Recommend models](https://www.algolia.com/doc/guides/algolia-recommend/overview/#recommend-models).
   */
  model: RecommendModels;
  searchRecommendRulesParams?: SearchRecommendRulesParams;
};
