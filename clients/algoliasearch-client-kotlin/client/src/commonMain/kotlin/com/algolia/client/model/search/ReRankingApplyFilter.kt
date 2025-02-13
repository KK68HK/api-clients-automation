/** Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT. */
package com.algolia.client.model.search

import com.algolia.client.exception.AlgoliaClientException
import com.algolia.client.extensions.internal.*
import kotlinx.serialization.*
import kotlinx.serialization.builtins.*
import kotlinx.serialization.descriptors.*
import kotlinx.serialization.encoding.*
import kotlinx.serialization.json.*
import kotlin.jvm.JvmInline

/**
 * When [Dynamic Re-Ranking](https://www.algolia.com/doc/guides/algolia-ai/re-ranking/) is enabled, only records that match these filters will be affected by Dynamic Re-Ranking.
 *
 * Implementations:
 * - [List<MixedSearchFilters>] - *[ReRankingApplyFilter.of]*
 * - [String] - *[ReRankingApplyFilter.of]*
 */
@Serializable(ReRankingApplyFilterSerializer::class)
public sealed interface ReRankingApplyFilter {

  @JvmInline
  public value class ListOfMixedSearchFiltersValue(public val value: List<MixedSearchFilters>) : ReRankingApplyFilter

  @JvmInline
  public value class StringValue(public val value: String) : ReRankingApplyFilter

  public companion object {

    /** [ReRankingApplyFilter] as [List<MixedSearchFilters>] Value. */
    public fun of(value: List<MixedSearchFilters>): ReRankingApplyFilter {
      return ListOfMixedSearchFiltersValue(value)
    }

    /** [ReRankingApplyFilter] as [String] Value. */
    public fun of(value: String): ReRankingApplyFilter {
      return StringValue(value)
    }
  }
}

internal class ReRankingApplyFilterSerializer : KSerializer<ReRankingApplyFilter> {

  override val descriptor: SerialDescriptor = buildClassSerialDescriptor("ReRankingApplyFilter")

  override fun serialize(encoder: Encoder, value: ReRankingApplyFilter) {
    when (value) {
      is ReRankingApplyFilter.ListOfMixedSearchFiltersValue -> ListSerializer(MixedSearchFilters.serializer()).serialize(encoder, value.value)
      is ReRankingApplyFilter.StringValue -> String.serializer().serialize(encoder, value.value)
    }
  }

  override fun deserialize(decoder: Decoder): ReRankingApplyFilter {
    val codec = decoder.asJsonDecoder()
    val tree = codec.decodeJsonElement()

    // deserialize List<MixedSearchFilters>
    if (tree is JsonArray) {
      try {
        val value = codec.json.decodeFromJsonElement(ListSerializer(MixedSearchFilters.serializer()), tree)
        return ReRankingApplyFilter.of(value)
      } catch (e: Exception) {
        // deserialization failed, continue
        println("Failed to deserialize List<MixedSearchFilters> (error: ${e.message})")
      }
    }

    // deserialize String
    if (tree is JsonPrimitive) {
      try {
        val value = codec.json.decodeFromJsonElement(String.serializer(), tree)
        return ReRankingApplyFilter.of(value)
      } catch (e: Exception) {
        // deserialization failed, continue
        println("Failed to deserialize String (error: ${e.message})")
      }
    }

    throw AlgoliaClientException("Failed to deserialize json element: $tree")
  }
}
