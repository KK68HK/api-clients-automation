/** Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT. */
package com.algolia.client.model.ingestion

import kotlinx.serialization.*
import kotlinx.serialization.json.*

/**
 * BigCommerceMetafield
 *
 * @param namespace The namespace of the metafield.
 * @param key The key identifier of the metafield.
 */
@Serializable
public data class BigCommerceMetafield(

  /** The namespace of the metafield. */
  @SerialName(value = "namespace") val namespace: String,

  /** The key identifier of the metafield. */
  @SerialName(value = "key") val key: String,
)
