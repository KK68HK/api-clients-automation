// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

import type { Condition } from './condition';
import type { Consequence } from './consequence';
import type { RuleResponseMetadata } from './ruleResponseMetadata';

/**
 * Rule object.
 */
export type RuleResponse = {
  _metadata?: RuleResponseMetadata;

  /**
   * Unique identifier for a rule object.
   */
  objectID: string;

  /**
   * [Conditions](https://www.algolia.com/doc/guides/managing-results/rules/rules-overview/#conditions) required to activate a rule. You can use up to 25 conditions per rule.
   */
  conditions?: Condition[];

  consequence?: Consequence;

  /**
   * Description of the rule\'s purpose. This can be helpful for display in the Algolia dashboard.
   */
  description?: string;

  /**
   * Indicates whether to enable the rule. If it isn\'t enabled, it isn\'t applied at query time.
   */
  enabled?: boolean;
};
