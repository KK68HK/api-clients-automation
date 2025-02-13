/** Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT. */
package com.algolia.client.model.analytics

import kotlinx.serialization.*
import kotlinx.serialization.json.*

/**
 * GetConversationRateResponse
 *
 * @param rate [Click-through rate (CTR)](https://www.algolia.com/doc/guides/search-analytics/concepts/metrics/#click-through-rate).
 * @param trackedSearchCount Number of tracked searches. This is the number of search requests where the `clickAnalytics` parameter is `true`.
 * @param conversionCount Number of converted clicks.
 * @param dates Conversion events.
 */
@Serializable
public data class GetConversationRateResponse(

  /** [Click-through rate (CTR)](https://www.algolia.com/doc/guides/search-analytics/concepts/metrics/#click-through-rate).  */
  @SerialName(value = "rate") val rate: Double,

  /** Number of tracked searches. This is the number of search requests where the `clickAnalytics` parameter is `true`. */
  @SerialName(value = "trackedSearchCount") val trackedSearchCount: Int,

  /** Number of converted clicks. */
  @SerialName(value = "conversionCount") val conversionCount: Int,

  /** Conversion events. */
  @SerialName(value = "dates") val dates: List<ConversionRateEvent>,
)
