// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
// ignore_for_file: unused_element
import 'package:algolia_client_search/src/model/scope_type.dart';
import 'package:algolia_client_search/src/model/operation_type.dart';

import 'package:json_annotation/json_annotation.dart';

part 'operation_index_params.g.dart';

@JsonSerializable()
final class OperationIndexParams {
  /// Returns a new [OperationIndexParams] instance.
  const OperationIndexParams({
    required this.operation,
    required this.destination,
    this.scope,
  });

  @JsonKey(name: r'operation')
  final OperationType operation;

  /// Algolia index name.
  @JsonKey(name: r'destination')
  final String destination;

  /// **This only applies to the _copy_ operation.**  If you omit `scope`, the copy command copies all records, settings, synonyms, and rules.  If you specify `scope`, only the specified scopes are copied.
  @JsonKey(name: r'scope')
  final List<ScopeType>? scope;

  @override
  bool operator ==(Object other) =>
      identical(this, other) ||
      other is OperationIndexParams &&
          other.operation == operation &&
          other.destination == destination &&
          other.scope == scope;

  @override
  int get hashCode =>
      operation.hashCode + destination.hashCode + scope.hashCode;

  factory OperationIndexParams.fromJson(Map<String, dynamic> json) =>
      _$OperationIndexParamsFromJson(json);

  Map<String, dynamic> toJson() => _$OperationIndexParamsToJson(this);

  @override
  String toString() {
    return toJson().toString();
  }
}
