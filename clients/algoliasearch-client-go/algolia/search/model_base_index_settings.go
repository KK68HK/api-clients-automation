// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package search

import (
	"encoding/json"
	"fmt"
)

// BaseIndexSettings struct for BaseIndexSettings
type BaseIndexSettings struct {
	// Creates [replicas](https://www.algolia.com/doc/guides/managing-results/refine-results/sorting/in-depth/replicas/), which are copies of a primary index with the same records but different settings.
	Replicas []string `json:"replicas,omitempty"`
	// Maximum number of hits accessible through pagination.
	PaginationLimitedTo *int32 `json:"paginationLimitedTo,omitempty"`
	// Attributes that can't be retrieved at query time.
	UnretrievableAttributes []string `json:"unretrievableAttributes,omitempty"`
	// Words for which you want to turn off [typo tolerance](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/typo-tolerance/).
	DisableTypoToleranceOnWords []string `json:"disableTypoToleranceOnWords,omitempty"`
	// Attributes in your index to which [Japanese transliteration](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/handling-natural-languages-nlp/in-depth/language-specific-configurations/#japanese-transliteration-and-type-ahead) applies. This will ensure that words indexed in Katakana or Kanji can also be searched in Hiragana.
	AttributesToTransliterate []string `json:"attributesToTransliterate,omitempty"`
	// Attributes on which to split [camel case](https://wikipedia.org/wiki/Camel_case) words.
	CamelCaseAttributes []string `json:"camelCaseAttributes,omitempty"`
	// Attributes in your index to which [word segmentation](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/handling-natural-languages-nlp/how-to/customize-segmentation/) (decompounding) applies.
	DecompoundedAttributes map[string]interface{} `json:"decompoundedAttributes,omitempty"`
	// Set the languages of your index, for language-specific processing steps such as [tokenization](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/handling-natural-languages-nlp/in-depth/tokenization/) and [normalization](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/handling-natural-languages-nlp/in-depth/normalization/).
	IndexLanguages []string `json:"indexLanguages,omitempty"`
	// Attributes for which you want to turn off [prefix matching](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/override-search-engine-defaults/#adjusting-prefix-search).
	DisablePrefixOnAttributes []string `json:"disablePrefixOnAttributes,omitempty"`
	// Incidates whether the engine compresses arrays with exclusively non-negative integers. When enabled, the compressed arrays may be reordered.
	AllowCompressionOfIntegerArray *bool `json:"allowCompressionOfIntegerArray,omitempty"`
	// Numeric attributes that can be used as [numerical filters](https://www.algolia.com/doc/guides/managing-results/rules/detecting-intent/how-to/applying-a-custom-filter-for-a-specific-query/#numerical-filters).
	NumericAttributesForFiltering []string `json:"numericAttributesForFiltering,omitempty"`
	// Controls which separators are added to an Algolia index as part of [normalization](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/handling-natural-languages-nlp/#what-does-normalization-mean). Separators are all non-letter characters except spaces and currency characters, such as $€£¥.
	SeparatorsToIndex *string `json:"separatorsToIndex,omitempty"`
	// [Attributes used for searching](https://www.algolia.com/doc/guides/managing-results/must-do/searchable-attributes/), including determining [if matches at the beginning of a word are important (ordered) or not (unordered)](https://www.algolia.com/doc/guides/managing-results/must-do/searchable-attributes/how-to/configuring-searchable-attributes-the-right-way/#understanding-word-position).
	SearchableAttributes []string `json:"searchableAttributes,omitempty"`
	// Lets you store custom data in your indices.
	UserData interface{} `json:"userData,omitempty"`
	// A list of characters and their normalized replacements to override Algolia's default [normalization](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/handling-natural-languages-nlp/in-depth/normalization/).
	CustomNormalization *map[string]map[string]string `json:"customNormalization,omitempty"`
}

type BaseIndexSettingsOption func(f *BaseIndexSettings)

func WithBaseIndexSettingsReplicas(val []string) BaseIndexSettingsOption {
	return func(f *BaseIndexSettings) {
		f.Replicas = val
	}
}

func WithBaseIndexSettingsPaginationLimitedTo(val int32) BaseIndexSettingsOption {
	return func(f *BaseIndexSettings) {
		f.PaginationLimitedTo = &val
	}
}

func WithBaseIndexSettingsUnretrievableAttributes(val []string) BaseIndexSettingsOption {
	return func(f *BaseIndexSettings) {
		f.UnretrievableAttributes = val
	}
}

func WithBaseIndexSettingsDisableTypoToleranceOnWords(val []string) BaseIndexSettingsOption {
	return func(f *BaseIndexSettings) {
		f.DisableTypoToleranceOnWords = val
	}
}

func WithBaseIndexSettingsAttributesToTransliterate(val []string) BaseIndexSettingsOption {
	return func(f *BaseIndexSettings) {
		f.AttributesToTransliterate = val
	}
}

func WithBaseIndexSettingsCamelCaseAttributes(val []string) BaseIndexSettingsOption {
	return func(f *BaseIndexSettings) {
		f.CamelCaseAttributes = val
	}
}

func WithBaseIndexSettingsDecompoundedAttributes(val map[string]interface{}) BaseIndexSettingsOption {
	return func(f *BaseIndexSettings) {
		f.DecompoundedAttributes = val
	}
}

func WithBaseIndexSettingsIndexLanguages(val []string) BaseIndexSettingsOption {
	return func(f *BaseIndexSettings) {
		f.IndexLanguages = val
	}
}

func WithBaseIndexSettingsDisablePrefixOnAttributes(val []string) BaseIndexSettingsOption {
	return func(f *BaseIndexSettings) {
		f.DisablePrefixOnAttributes = val
	}
}

func WithBaseIndexSettingsAllowCompressionOfIntegerArray(val bool) BaseIndexSettingsOption {
	return func(f *BaseIndexSettings) {
		f.AllowCompressionOfIntegerArray = &val
	}
}

func WithBaseIndexSettingsNumericAttributesForFiltering(val []string) BaseIndexSettingsOption {
	return func(f *BaseIndexSettings) {
		f.NumericAttributesForFiltering = val
	}
}

func WithBaseIndexSettingsSeparatorsToIndex(val string) BaseIndexSettingsOption {
	return func(f *BaseIndexSettings) {
		f.SeparatorsToIndex = &val
	}
}

func WithBaseIndexSettingsSearchableAttributes(val []string) BaseIndexSettingsOption {
	return func(f *BaseIndexSettings) {
		f.SearchableAttributes = val
	}
}

func WithBaseIndexSettingsUserData(val interface{}) BaseIndexSettingsOption {
	return func(f *BaseIndexSettings) {
		f.UserData = val
	}
}

func WithBaseIndexSettingsCustomNormalization(val map[string]map[string]string) BaseIndexSettingsOption {
	return func(f *BaseIndexSettings) {
		f.CustomNormalization = &val
	}
}

// NewBaseIndexSettings instantiates a new BaseIndexSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseIndexSettings(opts ...BaseIndexSettingsOption) *BaseIndexSettings {
	this := &BaseIndexSettings{}
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewBaseIndexSettingsWithDefaults instantiates a new BaseIndexSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseIndexSettingsWithDefaults() *BaseIndexSettings {
	this := &BaseIndexSettings{}
	var paginationLimitedTo int32 = 1000
	this.PaginationLimitedTo = &paginationLimitedTo
	var allowCompressionOfIntegerArray bool = false
	this.AllowCompressionOfIntegerArray = &allowCompressionOfIntegerArray
	var separatorsToIndex string = ""
	this.SeparatorsToIndex = &separatorsToIndex
	return this
}

// GetReplicas returns the Replicas field value if set, zero value otherwise.
func (o *BaseIndexSettings) GetReplicas() []string {
	if o == nil || o.Replicas == nil {
		var ret []string
		return ret
	}
	return o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseIndexSettings) GetReplicasOk() ([]string, bool) {
	if o == nil || o.Replicas == nil {
		return nil, false
	}
	return o.Replicas, true
}

// HasReplicas returns a boolean if a field has been set.
func (o *BaseIndexSettings) HasReplicas() bool {
	if o != nil && o.Replicas != nil {
		return true
	}

	return false
}

// SetReplicas gets a reference to the given []string and assigns it to the Replicas field.
func (o *BaseIndexSettings) SetReplicas(v []string) {
	o.Replicas = v
}

// GetPaginationLimitedTo returns the PaginationLimitedTo field value if set, zero value otherwise.
func (o *BaseIndexSettings) GetPaginationLimitedTo() int32 {
	if o == nil || o.PaginationLimitedTo == nil {
		var ret int32
		return ret
	}
	return *o.PaginationLimitedTo
}

// GetPaginationLimitedToOk returns a tuple with the PaginationLimitedTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseIndexSettings) GetPaginationLimitedToOk() (*int32, bool) {
	if o == nil || o.PaginationLimitedTo == nil {
		return nil, false
	}
	return o.PaginationLimitedTo, true
}

// HasPaginationLimitedTo returns a boolean if a field has been set.
func (o *BaseIndexSettings) HasPaginationLimitedTo() bool {
	if o != nil && o.PaginationLimitedTo != nil {
		return true
	}

	return false
}

// SetPaginationLimitedTo gets a reference to the given int32 and assigns it to the PaginationLimitedTo field.
func (o *BaseIndexSettings) SetPaginationLimitedTo(v int32) {
	o.PaginationLimitedTo = &v
}

// GetUnretrievableAttributes returns the UnretrievableAttributes field value if set, zero value otherwise.
func (o *BaseIndexSettings) GetUnretrievableAttributes() []string {
	if o == nil || o.UnretrievableAttributes == nil {
		var ret []string
		return ret
	}
	return o.UnretrievableAttributes
}

// GetUnretrievableAttributesOk returns a tuple with the UnretrievableAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseIndexSettings) GetUnretrievableAttributesOk() ([]string, bool) {
	if o == nil || o.UnretrievableAttributes == nil {
		return nil, false
	}
	return o.UnretrievableAttributes, true
}

// HasUnretrievableAttributes returns a boolean if a field has been set.
func (o *BaseIndexSettings) HasUnretrievableAttributes() bool {
	if o != nil && o.UnretrievableAttributes != nil {
		return true
	}

	return false
}

// SetUnretrievableAttributes gets a reference to the given []string and assigns it to the UnretrievableAttributes field.
func (o *BaseIndexSettings) SetUnretrievableAttributes(v []string) {
	o.UnretrievableAttributes = v
}

// GetDisableTypoToleranceOnWords returns the DisableTypoToleranceOnWords field value if set, zero value otherwise.
func (o *BaseIndexSettings) GetDisableTypoToleranceOnWords() []string {
	if o == nil || o.DisableTypoToleranceOnWords == nil {
		var ret []string
		return ret
	}
	return o.DisableTypoToleranceOnWords
}

// GetDisableTypoToleranceOnWordsOk returns a tuple with the DisableTypoToleranceOnWords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseIndexSettings) GetDisableTypoToleranceOnWordsOk() ([]string, bool) {
	if o == nil || o.DisableTypoToleranceOnWords == nil {
		return nil, false
	}
	return o.DisableTypoToleranceOnWords, true
}

// HasDisableTypoToleranceOnWords returns a boolean if a field has been set.
func (o *BaseIndexSettings) HasDisableTypoToleranceOnWords() bool {
	if o != nil && o.DisableTypoToleranceOnWords != nil {
		return true
	}

	return false
}

// SetDisableTypoToleranceOnWords gets a reference to the given []string and assigns it to the DisableTypoToleranceOnWords field.
func (o *BaseIndexSettings) SetDisableTypoToleranceOnWords(v []string) {
	o.DisableTypoToleranceOnWords = v
}

// GetAttributesToTransliterate returns the AttributesToTransliterate field value if set, zero value otherwise.
func (o *BaseIndexSettings) GetAttributesToTransliterate() []string {
	if o == nil || o.AttributesToTransliterate == nil {
		var ret []string
		return ret
	}
	return o.AttributesToTransliterate
}

// GetAttributesToTransliterateOk returns a tuple with the AttributesToTransliterate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseIndexSettings) GetAttributesToTransliterateOk() ([]string, bool) {
	if o == nil || o.AttributesToTransliterate == nil {
		return nil, false
	}
	return o.AttributesToTransliterate, true
}

// HasAttributesToTransliterate returns a boolean if a field has been set.
func (o *BaseIndexSettings) HasAttributesToTransliterate() bool {
	if o != nil && o.AttributesToTransliterate != nil {
		return true
	}

	return false
}

// SetAttributesToTransliterate gets a reference to the given []string and assigns it to the AttributesToTransliterate field.
func (o *BaseIndexSettings) SetAttributesToTransliterate(v []string) {
	o.AttributesToTransliterate = v
}

// GetCamelCaseAttributes returns the CamelCaseAttributes field value if set, zero value otherwise.
func (o *BaseIndexSettings) GetCamelCaseAttributes() []string {
	if o == nil || o.CamelCaseAttributes == nil {
		var ret []string
		return ret
	}
	return o.CamelCaseAttributes
}

// GetCamelCaseAttributesOk returns a tuple with the CamelCaseAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseIndexSettings) GetCamelCaseAttributesOk() ([]string, bool) {
	if o == nil || o.CamelCaseAttributes == nil {
		return nil, false
	}
	return o.CamelCaseAttributes, true
}

// HasCamelCaseAttributes returns a boolean if a field has been set.
func (o *BaseIndexSettings) HasCamelCaseAttributes() bool {
	if o != nil && o.CamelCaseAttributes != nil {
		return true
	}

	return false
}

// SetCamelCaseAttributes gets a reference to the given []string and assigns it to the CamelCaseAttributes field.
func (o *BaseIndexSettings) SetCamelCaseAttributes(v []string) {
	o.CamelCaseAttributes = v
}

// GetDecompoundedAttributes returns the DecompoundedAttributes field value if set, zero value otherwise.
func (o *BaseIndexSettings) GetDecompoundedAttributes() map[string]interface{} {
	if o == nil || o.DecompoundedAttributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.DecompoundedAttributes
}

// GetDecompoundedAttributesOk returns a tuple with the DecompoundedAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseIndexSettings) GetDecompoundedAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.DecompoundedAttributes == nil {
		return nil, false
	}
	return o.DecompoundedAttributes, true
}

// HasDecompoundedAttributes returns a boolean if a field has been set.
func (o *BaseIndexSettings) HasDecompoundedAttributes() bool {
	if o != nil && o.DecompoundedAttributes != nil {
		return true
	}

	return false
}

// SetDecompoundedAttributes gets a reference to the given map[string]interface{} and assigns it to the DecompoundedAttributes field.
func (o *BaseIndexSettings) SetDecompoundedAttributes(v map[string]interface{}) {
	o.DecompoundedAttributes = v
}

// GetIndexLanguages returns the IndexLanguages field value if set, zero value otherwise.
func (o *BaseIndexSettings) GetIndexLanguages() []string {
	if o == nil || o.IndexLanguages == nil {
		var ret []string
		return ret
	}
	return o.IndexLanguages
}

// GetIndexLanguagesOk returns a tuple with the IndexLanguages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseIndexSettings) GetIndexLanguagesOk() ([]string, bool) {
	if o == nil || o.IndexLanguages == nil {
		return nil, false
	}
	return o.IndexLanguages, true
}

// HasIndexLanguages returns a boolean if a field has been set.
func (o *BaseIndexSettings) HasIndexLanguages() bool {
	if o != nil && o.IndexLanguages != nil {
		return true
	}

	return false
}

// SetIndexLanguages gets a reference to the given []string and assigns it to the IndexLanguages field.
func (o *BaseIndexSettings) SetIndexLanguages(v []string) {
	o.IndexLanguages = v
}

// GetDisablePrefixOnAttributes returns the DisablePrefixOnAttributes field value if set, zero value otherwise.
func (o *BaseIndexSettings) GetDisablePrefixOnAttributes() []string {
	if o == nil || o.DisablePrefixOnAttributes == nil {
		var ret []string
		return ret
	}
	return o.DisablePrefixOnAttributes
}

// GetDisablePrefixOnAttributesOk returns a tuple with the DisablePrefixOnAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseIndexSettings) GetDisablePrefixOnAttributesOk() ([]string, bool) {
	if o == nil || o.DisablePrefixOnAttributes == nil {
		return nil, false
	}
	return o.DisablePrefixOnAttributes, true
}

// HasDisablePrefixOnAttributes returns a boolean if a field has been set.
func (o *BaseIndexSettings) HasDisablePrefixOnAttributes() bool {
	if o != nil && o.DisablePrefixOnAttributes != nil {
		return true
	}

	return false
}

// SetDisablePrefixOnAttributes gets a reference to the given []string and assigns it to the DisablePrefixOnAttributes field.
func (o *BaseIndexSettings) SetDisablePrefixOnAttributes(v []string) {
	o.DisablePrefixOnAttributes = v
}

// GetAllowCompressionOfIntegerArray returns the AllowCompressionOfIntegerArray field value if set, zero value otherwise.
func (o *BaseIndexSettings) GetAllowCompressionOfIntegerArray() bool {
	if o == nil || o.AllowCompressionOfIntegerArray == nil {
		var ret bool
		return ret
	}
	return *o.AllowCompressionOfIntegerArray
}

// GetAllowCompressionOfIntegerArrayOk returns a tuple with the AllowCompressionOfIntegerArray field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseIndexSettings) GetAllowCompressionOfIntegerArrayOk() (*bool, bool) {
	if o == nil || o.AllowCompressionOfIntegerArray == nil {
		return nil, false
	}
	return o.AllowCompressionOfIntegerArray, true
}

// HasAllowCompressionOfIntegerArray returns a boolean if a field has been set.
func (o *BaseIndexSettings) HasAllowCompressionOfIntegerArray() bool {
	if o != nil && o.AllowCompressionOfIntegerArray != nil {
		return true
	}

	return false
}

// SetAllowCompressionOfIntegerArray gets a reference to the given bool and assigns it to the AllowCompressionOfIntegerArray field.
func (o *BaseIndexSettings) SetAllowCompressionOfIntegerArray(v bool) {
	o.AllowCompressionOfIntegerArray = &v
}

// GetNumericAttributesForFiltering returns the NumericAttributesForFiltering field value if set, zero value otherwise.
func (o *BaseIndexSettings) GetNumericAttributesForFiltering() []string {
	if o == nil || o.NumericAttributesForFiltering == nil {
		var ret []string
		return ret
	}
	return o.NumericAttributesForFiltering
}

// GetNumericAttributesForFilteringOk returns a tuple with the NumericAttributesForFiltering field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseIndexSettings) GetNumericAttributesForFilteringOk() ([]string, bool) {
	if o == nil || o.NumericAttributesForFiltering == nil {
		return nil, false
	}
	return o.NumericAttributesForFiltering, true
}

// HasNumericAttributesForFiltering returns a boolean if a field has been set.
func (o *BaseIndexSettings) HasNumericAttributesForFiltering() bool {
	if o != nil && o.NumericAttributesForFiltering != nil {
		return true
	}

	return false
}

// SetNumericAttributesForFiltering gets a reference to the given []string and assigns it to the NumericAttributesForFiltering field.
func (o *BaseIndexSettings) SetNumericAttributesForFiltering(v []string) {
	o.NumericAttributesForFiltering = v
}

// GetSeparatorsToIndex returns the SeparatorsToIndex field value if set, zero value otherwise.
func (o *BaseIndexSettings) GetSeparatorsToIndex() string {
	if o == nil || o.SeparatorsToIndex == nil {
		var ret string
		return ret
	}
	return *o.SeparatorsToIndex
}

// GetSeparatorsToIndexOk returns a tuple with the SeparatorsToIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseIndexSettings) GetSeparatorsToIndexOk() (*string, bool) {
	if o == nil || o.SeparatorsToIndex == nil {
		return nil, false
	}
	return o.SeparatorsToIndex, true
}

// HasSeparatorsToIndex returns a boolean if a field has been set.
func (o *BaseIndexSettings) HasSeparatorsToIndex() bool {
	if o != nil && o.SeparatorsToIndex != nil {
		return true
	}

	return false
}

// SetSeparatorsToIndex gets a reference to the given string and assigns it to the SeparatorsToIndex field.
func (o *BaseIndexSettings) SetSeparatorsToIndex(v string) {
	o.SeparatorsToIndex = &v
}

// GetSearchableAttributes returns the SearchableAttributes field value if set, zero value otherwise.
func (o *BaseIndexSettings) GetSearchableAttributes() []string {
	if o == nil || o.SearchableAttributes == nil {
		var ret []string
		return ret
	}
	return o.SearchableAttributes
}

// GetSearchableAttributesOk returns a tuple with the SearchableAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseIndexSettings) GetSearchableAttributesOk() ([]string, bool) {
	if o == nil || o.SearchableAttributes == nil {
		return nil, false
	}
	return o.SearchableAttributes, true
}

// HasSearchableAttributes returns a boolean if a field has been set.
func (o *BaseIndexSettings) HasSearchableAttributes() bool {
	if o != nil && o.SearchableAttributes != nil {
		return true
	}

	return false
}

// SetSearchableAttributes gets a reference to the given []string and assigns it to the SearchableAttributes field.
func (o *BaseIndexSettings) SetSearchableAttributes(v []string) {
	o.SearchableAttributes = v
}

// GetUserData returns the UserData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseIndexSettings) GetUserData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UserData
}

// GetUserDataOk returns a tuple with the UserData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseIndexSettings) GetUserDataOk() (*interface{}, bool) {
	if o == nil || o.UserData == nil {
		return nil, false
	}
	return &o.UserData, true
}

// HasUserData returns a boolean if a field has been set.
func (o *BaseIndexSettings) HasUserData() bool {
	if o != nil && o.UserData != nil {
		return true
	}

	return false
}

// SetUserData gets a reference to the given interface{} and assigns it to the UserData field.
func (o *BaseIndexSettings) SetUserData(v interface{}) {
	o.UserData = v
}

// GetCustomNormalization returns the CustomNormalization field value if set, zero value otherwise.
func (o *BaseIndexSettings) GetCustomNormalization() map[string]map[string]string {
	if o == nil || o.CustomNormalization == nil {
		var ret map[string]map[string]string
		return ret
	}
	return *o.CustomNormalization
}

// GetCustomNormalizationOk returns a tuple with the CustomNormalization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseIndexSettings) GetCustomNormalizationOk() (*map[string]map[string]string, bool) {
	if o == nil || o.CustomNormalization == nil {
		return nil, false
	}
	return o.CustomNormalization, true
}

// HasCustomNormalization returns a boolean if a field has been set.
func (o *BaseIndexSettings) HasCustomNormalization() bool {
	if o != nil && o.CustomNormalization != nil {
		return true
	}

	return false
}

// SetCustomNormalization gets a reference to the given map[string]map[string]string and assigns it to the CustomNormalization field.
func (o *BaseIndexSettings) SetCustomNormalization(v map[string]map[string]string) {
	o.CustomNormalization = &v
}

func (o BaseIndexSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Replicas != nil {
		toSerialize["replicas"] = o.Replicas
	}
	if o.PaginationLimitedTo != nil {
		toSerialize["paginationLimitedTo"] = o.PaginationLimitedTo
	}
	if o.UnretrievableAttributes != nil {
		toSerialize["unretrievableAttributes"] = o.UnretrievableAttributes
	}
	if o.DisableTypoToleranceOnWords != nil {
		toSerialize["disableTypoToleranceOnWords"] = o.DisableTypoToleranceOnWords
	}
	if o.AttributesToTransliterate != nil {
		toSerialize["attributesToTransliterate"] = o.AttributesToTransliterate
	}
	if o.CamelCaseAttributes != nil {
		toSerialize["camelCaseAttributes"] = o.CamelCaseAttributes
	}
	if o.DecompoundedAttributes != nil {
		toSerialize["decompoundedAttributes"] = o.DecompoundedAttributes
	}
	if o.IndexLanguages != nil {
		toSerialize["indexLanguages"] = o.IndexLanguages
	}
	if o.DisablePrefixOnAttributes != nil {
		toSerialize["disablePrefixOnAttributes"] = o.DisablePrefixOnAttributes
	}
	if o.AllowCompressionOfIntegerArray != nil {
		toSerialize["allowCompressionOfIntegerArray"] = o.AllowCompressionOfIntegerArray
	}
	if o.NumericAttributesForFiltering != nil {
		toSerialize["numericAttributesForFiltering"] = o.NumericAttributesForFiltering
	}
	if o.SeparatorsToIndex != nil {
		toSerialize["separatorsToIndex"] = o.SeparatorsToIndex
	}
	if o.SearchableAttributes != nil {
		toSerialize["searchableAttributes"] = o.SearchableAttributes
	}
	if o.UserData != nil {
		toSerialize["userData"] = o.UserData
	}
	if o.CustomNormalization != nil {
		toSerialize["customNormalization"] = o.CustomNormalization
	}
	return json.Marshal(toSerialize)
}

func (o BaseIndexSettings) String() string {
	out := ""
	out += fmt.Sprintf("  replicas=%v\n", o.Replicas)
	out += fmt.Sprintf("  paginationLimitedTo=%v\n", o.PaginationLimitedTo)
	out += fmt.Sprintf("  unretrievableAttributes=%v\n", o.UnretrievableAttributes)
	out += fmt.Sprintf("  disableTypoToleranceOnWords=%v\n", o.DisableTypoToleranceOnWords)
	out += fmt.Sprintf("  attributesToTransliterate=%v\n", o.AttributesToTransliterate)
	out += fmt.Sprintf("  camelCaseAttributes=%v\n", o.CamelCaseAttributes)
	out += fmt.Sprintf("  decompoundedAttributes=%v\n", o.DecompoundedAttributes)
	out += fmt.Sprintf("  indexLanguages=%v\n", o.IndexLanguages)
	out += fmt.Sprintf("  disablePrefixOnAttributes=%v\n", o.DisablePrefixOnAttributes)
	out += fmt.Sprintf("  allowCompressionOfIntegerArray=%v\n", o.AllowCompressionOfIntegerArray)
	out += fmt.Sprintf("  numericAttributesForFiltering=%v\n", o.NumericAttributesForFiltering)
	out += fmt.Sprintf("  separatorsToIndex=%v\n", o.SeparatorsToIndex)
	out += fmt.Sprintf("  searchableAttributes=%v\n", o.SearchableAttributes)
	out += fmt.Sprintf("  userData=%v\n", o.UserData)
	out += fmt.Sprintf("  customNormalization=%v\n", o.CustomNormalization)
	return fmt.Sprintf("BaseIndexSettings {\n%s}", out)
}

type NullableBaseIndexSettings struct {
	value *BaseIndexSettings
	isSet bool
}

func (v NullableBaseIndexSettings) Get() *BaseIndexSettings {
	return v.value
}

func (v *NullableBaseIndexSettings) Set(val *BaseIndexSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseIndexSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseIndexSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseIndexSettings(val *BaseIndexSettings) *NullableBaseIndexSettings {
	return &NullableBaseIndexSettings{value: val, isSet: true}
}

func (v NullableBaseIndexSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseIndexSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
