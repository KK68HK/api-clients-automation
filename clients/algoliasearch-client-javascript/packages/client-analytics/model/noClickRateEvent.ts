// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

export type NoClickRateEvent = {
  /**
   * [Click-through rate (CTR)](https://www.algolia.com/doc/guides/search-analytics/concepts/metrics/#click-through-rate).
   */
  rate: number;

  /**
   * Number of tracked _and_ untracked searches (where the `clickAnalytics` parameter isn\'t `true`).
   */
  count: number;

  /**
   * Number of click events.
   */
  noClickCount: number;

  /**
   * Date of the event in the format YYYY-MM-DD.
   */
  date: string;
};
