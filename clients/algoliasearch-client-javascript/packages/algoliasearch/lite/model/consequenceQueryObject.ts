// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

import type { Edit } from './edit';

export type ConsequenceQueryObject = {
  /**
   * Words to remove.
   */
  remove?: string[];

  /**
   * Edits to apply.
   */
  edits?: Edit[];
};
