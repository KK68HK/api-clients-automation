// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.search;

import com.fasterxml.jackson.annotation.*;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Objects;

/** MultipleBatchResponse */
public class MultipleBatchResponse {

  @JsonProperty("taskID")
  private Map<String, Long> taskID = new HashMap<>();

  @JsonProperty("objectIDs")
  private List<String> objectIDs = new ArrayList<>();

  public MultipleBatchResponse setTaskID(Map<String, Long> taskID) {
    this.taskID = taskID;
    return this;
  }

  public MultipleBatchResponse putTaskID(String key, Long taskIDItem) {
    this.taskID.put(key, taskIDItem);
    return this;
  }

  /**
   * List of tasksIDs per index.
   *
   * @return taskID
   */
  @javax.annotation.Nonnull
  public Map<String, Long> getTaskID() {
    return taskID;
  }

  public MultipleBatchResponse setObjectIDs(List<String> objectIDs) {
    this.objectIDs = objectIDs;
    return this;
  }

  public MultipleBatchResponse addObjectIDs(String objectIDsItem) {
    this.objectIDs.add(objectIDsItem);
    return this;
  }

  /**
   * List of objectID.
   *
   * @return objectIDs
   */
  @javax.annotation.Nonnull
  public List<String> getObjectIDs() {
    return objectIDs;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    MultipleBatchResponse multipleBatchResponse = (MultipleBatchResponse) o;
    return Objects.equals(this.taskID, multipleBatchResponse.taskID) && Objects.equals(this.objectIDs, multipleBatchResponse.objectIDs);
  }

  @Override
  public int hashCode() {
    return Objects.hash(taskID, objectIDs);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class MultipleBatchResponse {\n");
    sb.append("    taskID: ").append(toIndentedString(taskID)).append("\n");
    sb.append("    objectIDs: ").append(toIndentedString(objectIDs)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }
}
