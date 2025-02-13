/** Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT. */
package com.algolia.client.model.search

import com.algolia.client.extensions.internal.*
import kotlinx.serialization.*
import kotlinx.serialization.descriptors.*
import kotlinx.serialization.encoding.*
import kotlinx.serialization.json.*

/**
 * A single hit.
 *
 * @param objectID Unique object identifier.
 * @param highlightResult Show highlighted section and words matched on a query.
 * @param snippetResult Snippeted attributes show parts of the matched attributes. Only returned when attributesToSnippet is non-empty.
 * @param rankingInfo
 * @param distinctSeqID
 */
@Serializable(HitSerializer::class)
public data class Hit(

  /** Unique object identifier. */
  val objectID: String,

  /** Show highlighted section and words matched on a query. */
  val highlightResult: Map<kotlin.String, HighlightResult>? = null,

  /** Snippeted attributes show parts of the matched attributes. Only returned when attributesToSnippet is non-empty. */
  val snippetResult: Map<kotlin.String, SnippetResult>? = null,

  val rankingInfo: RankingInfo? = null,

  val distinctSeqID: Int? = null,

  val additionalProperties: Map<String, JsonElement>? = null,
)

internal object HitSerializer : KSerializer<Hit> {

  override val descriptor: SerialDescriptor = buildClassSerialDescriptor("Hit") {
    element<String>("objectID")
    element<Map<kotlin.String, HighlightResult>>("_highlightResult")
    element<Map<kotlin.String, SnippetResult>>("_snippetResult")
    element<RankingInfo>("_rankingInfo")
    element<Int>("_distinctSeqID")
  }

  override fun deserialize(decoder: Decoder): Hit {
    val input = decoder.asJsonDecoder()
    val tree = input.decodeJsonObject()
    return Hit(
      objectID = tree.getValue("objectID").let { input.json.decodeFromJsonElement(it) },
      highlightResult = tree["_highlightResult"]?.let { input.json.decodeFromJsonElement(it) },
      snippetResult = tree["_snippetResult"]?.let { input.json.decodeFromJsonElement(it) },
      rankingInfo = tree["_rankingInfo"]?.let { input.json.decodeFromJsonElement(it) },
      distinctSeqID = tree["_distinctSeqID"]?.let { input.json.decodeFromJsonElement(it) },
      additionalProperties = tree.filterKeys { it !in descriptor.elementNames },
    )
  }

  override fun serialize(encoder: Encoder, value: Hit) {
    val output = encoder.asJsonEncoder()
    val json = buildJsonObject {
      put("objectID", output.json.encodeToJsonElement(value.objectID))
      value.highlightResult?.let { put("_highlightResult", output.json.encodeToJsonElement(it)) }
      value.snippetResult?.let { put("_snippetResult", output.json.encodeToJsonElement(it)) }
      value.rankingInfo?.let { put("_rankingInfo", output.json.encodeToJsonElement(it)) }
      value.distinctSeqID?.let { put("_distinctSeqID", output.json.encodeToJsonElement(it)) }
      value.additionalProperties?.onEach { (key, element) -> put(key, element) }
    }
    (encoder as JsonEncoder).encodeJsonElement(json)
  }
}
