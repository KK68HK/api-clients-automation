/** Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT. */
package com.algolia.client.model.search

import kotlinx.serialization.*
import kotlinx.serialization.json.*

/**
 * BatchResponse
 *
 * @param taskID Unique identifier of a task. A successful API response means that a task was added to a queue. It might not run immediately. You can check the task's progress with the `task` operation and this `taskID`.
 * @param objectIDs Unique object (record) identifiers.
 */
@Serializable
public data class BatchResponse(

  /** Unique identifier of a task. A successful API response means that a task was added to a queue. It might not run immediately. You can check the task's progress with the `task` operation and this `taskID`.  */
  @SerialName(value = "taskID") val taskID: Long,

  /** Unique object (record) identifiers. */
  @SerialName(value = "objectIDs") val objectIDs: List<String>,
)
