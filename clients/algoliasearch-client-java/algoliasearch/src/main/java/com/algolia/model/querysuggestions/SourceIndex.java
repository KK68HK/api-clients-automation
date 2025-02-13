// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost
// - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

package com.algolia.model.querysuggestions;

import com.fasterxml.jackson.annotation.*;
import com.fasterxml.jackson.databind.annotation.*;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

/** Configuration of an Algolia index for Query Suggestions. */
public class SourceIndex {

  @JsonProperty("indexName")
  private String indexName;

  @JsonProperty("replicas")
  private Boolean replicas;

  @JsonProperty("analyticsTags")
  private List<String> analyticsTags;

  @JsonProperty("facets")
  private List<Facet> facets;

  @JsonProperty("minHits")
  private Integer minHits;

  @JsonProperty("minLetters")
  private Integer minLetters;

  @JsonProperty("generate")
  private List<List<String>> generate;

  @JsonProperty("external")
  private List<String> external;

  public SourceIndex setIndexName(String indexName) {
    this.indexName = indexName;
    return this;
  }

  /** Name of the Algolia index to use as source for query suggestions. */
  @javax.annotation.Nonnull
  public String getIndexName() {
    return indexName;
  }

  public SourceIndex setReplicas(Boolean replicas) {
    this.replicas = replicas;
    return this;
  }

  /**
   * If true, Query Suggestions uses all replicas of the primary index to find popular searches. If
   * false, only the primary index is used.
   */
  @javax.annotation.Nullable
  public Boolean getReplicas() {
    return replicas;
  }

  public SourceIndex setAnalyticsTags(List<String> analyticsTags) {
    this.analyticsTags = analyticsTags;
    return this;
  }

  public SourceIndex addAnalyticsTags(String analyticsTagsItem) {
    if (this.analyticsTags == null) {
      this.analyticsTags = new ArrayList<>();
    }
    this.analyticsTags.add(analyticsTagsItem);
    return this;
  }

  /**
   * [Analytics tags](https://www.algolia.com/doc/api-reference/api-parameters/analyticsTags/) for
   * filtering the popular searches.
   */
  @javax.annotation.Nullable
  public List<String> getAnalyticsTags() {
    return analyticsTags;
  }

  public SourceIndex setFacets(List<Facet> facets) {
    this.facets = facets;
    return this;
  }

  public SourceIndex addFacets(Facet facetsItem) {
    if (this.facets == null) {
      this.facets = new ArrayList<>();
    }
    this.facets.add(facetsItem);
    return this;
  }

  /**
   * Facets to use as top categories with your suggestions. If provided, Query Suggestions adds the
   * top facet values to each suggestion.
   */
  @javax.annotation.Nullable
  public List<Facet> getFacets() {
    return facets;
  }

  public SourceIndex setMinHits(Integer minHits) {
    this.minHits = minHits;
    return this;
  }

  /**
   * Minimum number of hits required to be included as a suggestion. A search query must at least
   * generate `minHits` hits to be included in the Query Suggestions index. minimum: 0
   */
  @javax.annotation.Nullable
  public Integer getMinHits() {
    return minHits;
  }

  public SourceIndex setMinLetters(Integer minLetters) {
    this.minLetters = minLetters;
    return this;
  }

  /**
   * Minimum letters required to be included as a suggestion. A search query must be at least
   * `minLetters` long to be included in the Query Suggestions index. minimum: 0
   */
  @javax.annotation.Nullable
  public Integer getMinLetters() {
    return minLetters;
  }

  public SourceIndex setGenerate(List<List<String>> generate) {
    this.generate = generate;
    return this;
  }

  public SourceIndex addGenerate(List<String> generateItem) {
    if (this.generate == null) {
      this.generate = new ArrayList<>();
    }
    this.generate.add(generateItem);
    return this;
  }

  /** Get generate */
  @javax.annotation.Nullable
  public List<List<String>> getGenerate() {
    return generate;
  }

  public SourceIndex setExternal(List<String> external) {
    this.external = external;
    return this;
  }

  public SourceIndex addExternal(String externalItem) {
    if (this.external == null) {
      this.external = new ArrayList<>();
    }
    this.external.add(externalItem);
    return this;
  }

  /**
   * Algolia indices with popular searches to use as query suggestions. Records of these indices
   * must have these attributes: - `query`: search query which will be added as a suggestion -
   * `count`: measure of popularity of that search query For example, you can export popular
   * searches from an external analytics tool, such as Google Analytics or Adobe Analytics, and feed
   * this data into an external Algolia index. You can use this external index to generate query
   * suggestions until your Algolia analytics has collected enough data.
   */
  @javax.annotation.Nullable
  public List<String> getExternal() {
    return external;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    SourceIndex sourceIndex = (SourceIndex) o;
    return (
      Objects.equals(this.indexName, sourceIndex.indexName) &&
      Objects.equals(this.replicas, sourceIndex.replicas) &&
      Objects.equals(this.analyticsTags, sourceIndex.analyticsTags) &&
      Objects.equals(this.facets, sourceIndex.facets) &&
      Objects.equals(this.minHits, sourceIndex.minHits) &&
      Objects.equals(this.minLetters, sourceIndex.minLetters) &&
      Objects.equals(this.generate, sourceIndex.generate) &&
      Objects.equals(this.external, sourceIndex.external)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(indexName, replicas, analyticsTags, facets, minHits, minLetters, generate, external);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class SourceIndex {\n");
    sb.append("    indexName: ").append(toIndentedString(indexName)).append("\n");
    sb.append("    replicas: ").append(toIndentedString(replicas)).append("\n");
    sb.append("    analyticsTags: ").append(toIndentedString(analyticsTags)).append("\n");
    sb.append("    facets: ").append(toIndentedString(facets)).append("\n");
    sb.append("    minHits: ").append(toIndentedString(minHits)).append("\n");
    sb.append("    minLetters: ").append(toIndentedString(minLetters)).append("\n");
    sb.append("    generate: ").append(toIndentedString(generate)).append("\n");
    sb.append("    external: ").append(toIndentedString(external)).append("\n");
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
