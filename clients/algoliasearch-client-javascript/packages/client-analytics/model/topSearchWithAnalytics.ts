// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

export type TopSearchWithAnalytics = {
  /**
   * User query.
   */
  search: string;

  /**
   * Number of tracked _and_ untracked searches (where the `clickAnalytics` parameter isn\'t `true`).
   */
  count: number;

  /**
   * [Click-through rate (CTR)](https://www.algolia.com/doc/guides/search-analytics/concepts/metrics/#click-through-rate).
   */
  clickThroughRate: number;

  /**
   * Average [position](https://www.algolia.com/doc/guides/search-analytics/concepts/metrics/#click-position) of clicked search result.
   */
  averageClickPosition: number;

  /**
   * [Conversion rate (CR)](https://www.algolia.com/doc/guides/search-analytics/concepts/metrics/#conversion-rate).
   */
  conversionRate: number;

  /**
   * Number of tracked searches. This is the number of search requests where the `clickAnalytics` parameter is `true`.
   */
  trackedSearchCount: number;

  /**
   * Number of click events.
   */
  clickCount: number;

  /**
   * Number of converted clicks.
   */
  conversionCount: number;

  /**
   * Number of hits the search query matched.
   */
  nbHits: number;
};
