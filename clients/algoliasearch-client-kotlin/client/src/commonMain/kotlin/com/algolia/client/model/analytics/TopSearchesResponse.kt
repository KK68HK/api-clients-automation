/** Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT. */
package com.algolia.client.model.analytics

import kotlinx.serialization.*
import kotlinx.serialization.json.*

/**
 * TopSearchesResponse
 *
 * @param searches Top searches with their hits count.
 */
@Serializable
public data class TopSearchesResponse(

  /** Top searches with their hits count. */
  @SerialName(value = "searches") val searches: List<TopSearch>,
) : GetTopSearchesResponse
