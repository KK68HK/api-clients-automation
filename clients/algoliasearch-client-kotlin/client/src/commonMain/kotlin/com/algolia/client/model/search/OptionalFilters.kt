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
 * Create filters to boost or demote records.   Records that match the filter are ranked higher for positive and lower for negative optional filters. In contrast to regular filters, records that don't match the optional filter are still included in the results, only their ranking is affected.
 *
 * Implementations:
 * - [List<MixedSearchFilters>] - *[OptionalFilters.of]*
 * - [String] - *[OptionalFilters.of]*
 */
@Serializable(OptionalFiltersSerializer::class)
public sealed interface OptionalFilters {

  @JvmInline
  public value class ListOfMixedSearchFiltersValue(public val value: List<MixedSearchFilters>) : OptionalFilters

  @JvmInline
  public value class StringValue(public val value: String) : OptionalFilters

  public companion object {

    /** [OptionalFilters] as [List<MixedSearchFilters>] Value. */
    public fun of(value: List<MixedSearchFilters>): OptionalFilters {
      return ListOfMixedSearchFiltersValue(value)
    }

    /** [OptionalFilters] as [String] Value. */
    public fun of(value: String): OptionalFilters {
      return StringValue(value)
    }
  }
}

internal class OptionalFiltersSerializer : KSerializer<OptionalFilters> {

  override val descriptor: SerialDescriptor = buildClassSerialDescriptor("OptionalFilters")

  override fun serialize(encoder: Encoder, value: OptionalFilters) {
    when (value) {
      is OptionalFilters.ListOfMixedSearchFiltersValue -> ListSerializer(MixedSearchFilters.serializer()).serialize(encoder, value.value)
      is OptionalFilters.StringValue -> String.serializer().serialize(encoder, value.value)
    }
  }

  override fun deserialize(decoder: Decoder): OptionalFilters {
    val codec = decoder.asJsonDecoder()
    val tree = codec.decodeJsonElement()

    // deserialize List<MixedSearchFilters>
    if (tree is JsonArray) {
      try {
        val value = codec.json.decodeFromJsonElement(ListSerializer(MixedSearchFilters.serializer()), tree)
        return OptionalFilters.of(value)
      } catch (e: Exception) {
        // deserialization failed, continue
        println("Failed to deserialize List<MixedSearchFilters> (error: ${e.message})")
      }
    }

    // deserialize String
    if (tree is JsonPrimitive) {
      try {
        val value = codec.json.decodeFromJsonElement(String.serializer(), tree)
        return OptionalFilters.of(value)
      } catch (e: Exception) {
        // deserialization failed, continue
        println("Failed to deserialize String (error: ${e.message})")
      }
    }

    throw AlgoliaClientException("Failed to deserialize json element: $tree")
  }
}
