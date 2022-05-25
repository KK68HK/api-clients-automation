package com.algolia.model.search;

import com.google.gson.annotations.SerializedName;
import java.util.Objects;

/** The response with a taskID and a deletedAt timestamp. */
public class DeletedAtResponse {

  @SerializedName("taskID")
  private Long taskID;

  @SerializedName("deletedAt")
  private String deletedAt;

  public DeletedAtResponse setTaskID(Long taskID) {
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

  public DeletedAtResponse setDeletedAt(String deletedAt) {
    this.deletedAt = deletedAt;
    return this;
  }

  /**
   * Date of deletion (ISO-8601 format).
   *
   * @return deletedAt
   */
  @javax.annotation.Nonnull
  public String getDeletedAt() {
    return deletedAt;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    DeletedAtResponse deletedAtResponse = (DeletedAtResponse) o;
    return Objects.equals(this.taskID, deletedAtResponse.taskID) && Objects.equals(this.deletedAt, deletedAtResponse.deletedAt);
  }

  @Override
  public int hashCode() {
    return Objects.hash(taskID, deletedAt);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class DeletedAtResponse {\n");
    sb.append("    taskID: ").append(toIndentedString(taskID)).append("\n");
    sb.append("    deletedAt: ").append(toIndentedString(deletedAt)).append("\n");
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
