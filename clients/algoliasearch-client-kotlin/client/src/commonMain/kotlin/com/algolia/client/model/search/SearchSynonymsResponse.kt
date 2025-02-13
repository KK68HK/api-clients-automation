/** Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT. */
package com.algolia.client.model.search

import com.algolia.client.extensions.internal.*
import kotlinx.serialization.*
import kotlinx.serialization.descriptors.*
import kotlinx.serialization.encoding.*
import kotlinx.serialization.json.*

/**
 * SearchSynonymsResponse
 *
 * @param hits Synonym objects.
 * @param nbHits Number of hits the search query matched.
 */
@Serializable(SearchSynonymsResponseSerializer::class)
public data class SearchSynonymsResponse(

  /** Synonym objects. */
  val hits: List<SynonymHit>,

  /** Number of hits the search query matched. */
  val nbHits: Int,

  val additionalProperties: Map<String, JsonElement>? = null,
)

internal object SearchSynonymsResponseSerializer : KSerializer<SearchSynonymsResponse> {

  override val descriptor: SerialDescriptor = buildClassSerialDescriptor("SearchSynonymsResponse") {
    element<List<SynonymHit>>("hits")
    element<Int>("nbHits")
  }

  override fun deserialize(decoder: Decoder): SearchSynonymsResponse {
    val input = decoder.asJsonDecoder()
    val tree = input.decodeJsonObject()
    return SearchSynonymsResponse(
      hits = tree.getValue("hits").let { input.json.decodeFromJsonElement(it) },
      nbHits = tree.getValue("nbHits").let { input.json.decodeFromJsonElement(it) },
      additionalProperties = tree.filterKeys { it !in descriptor.elementNames },
    )
  }

  override fun serialize(encoder: Encoder, value: SearchSynonymsResponse) {
    val output = encoder.asJsonEncoder()
    val json = buildJsonObject {
      put("hits", output.json.encodeToJsonElement(value.hits))
      put("nbHits", output.json.encodeToJsonElement(value.nbHits))
      value.additionalProperties?.onEach { (key, element) -> put(key, element) }
    }
    (encoder as JsonEncoder).encodeJsonElement(json)
  }
}
