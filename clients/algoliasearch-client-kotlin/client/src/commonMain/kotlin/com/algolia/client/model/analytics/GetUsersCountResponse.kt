/** Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT. */
package com.algolia.client.model.analytics

import kotlinx.serialization.*
import kotlinx.serialization.json.*

/**
 * GetUsersCountResponse
 *
 * @param count Number of occurrences.
 * @param dates User count.
 */
@Serializable
public data class GetUsersCountResponse(

  /** Number of occurrences. */
  @SerialName(value = "count") val count: Int,

  /** User count. */
  @SerialName(value = "dates") val dates: List<UserWithDate>,
)
