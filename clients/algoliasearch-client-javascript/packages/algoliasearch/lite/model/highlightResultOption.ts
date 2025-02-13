// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

import type { MatchLevel } from './matchLevel';

/**
 * Show highlighted section and words matched on a query.
 */
export type HighlightResultOption = {
  /**
   * Markup text with `facetQuery` matches highlighted.
   */
  value: string;

  matchLevel: MatchLevel;

  /**
   * List of words from the query that matched the object.
   */
  matchedWords: string[];

  /**
   * Whether the entire attribute value is highlighted.
   */
  fullyHighlighted?: boolean;
};
