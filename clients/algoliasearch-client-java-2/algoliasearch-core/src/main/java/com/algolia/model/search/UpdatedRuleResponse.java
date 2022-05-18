package com.algolia.model.search;

import com.google.gson.annotations.SerializedName;
import java.util.Objects;

/** UpdatedRuleResponse */
public class UpdatedRuleResponse {

  @SerializedName("objectID")
  private String objectID;

  @SerializedName("updatedAt")
  private String updatedAt;

  @SerializedName("taskID")
  private Long taskID;

  public UpdatedRuleResponse setObjectID(String objectID) {
    this.objectID = objectID;
    return this;
  }

  /**
   * Unique identifier of the object.
   *
   * @return objectID
   */
  @javax.annotation.Nonnull
  public String getObjectID() {
    return objectID;
  }

  public UpdatedRuleResponse setUpdatedAt(String updatedAt) {
    this.updatedAt = updatedAt;
    return this;
  }

  /**
   * Date of last update (ISO-8601 format).
   *
   * @return updatedAt
   */
  @javax.annotation.Nonnull
  public String getUpdatedAt() {
    return updatedAt;
  }

  public UpdatedRuleResponse setTaskID(Long taskID) {
    this.taskID = taskID;
    return this;
  }

  /**
   * taskID of the task to wait for.
   *
   * @return taskID
   */
  @javax.annotation.Nonnull
  public Long getTaskID() {
    return taskID;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    UpdatedRuleResponse updatedRuleResponse = (UpdatedRuleResponse) o;
    return (
      Objects.equals(this.objectID, updatedRuleResponse.objectID) &&
      Objects.equals(this.updatedAt, updatedRuleResponse.updatedAt) &&
      Objects.equals(this.taskID, updatedRuleResponse.taskID)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(objectID, updatedAt, taskID);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class UpdatedRuleResponse {\n");
    sb.append("    objectID: ").append(toIndentedString(objectID)).append("\n");
    sb
      .append("    updatedAt: ")
      .append(toIndentedString(updatedAt))
      .append("\n");
    sb.append("    taskID: ").append(toIndentedString(taskID)).append("\n");
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
