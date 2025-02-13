<?php

// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

namespace Algolia\AlgoliaSearch\Model\Abtesting;

/**
 * CurrenciesValue Class Doc Comment.
 *
 * @category Class
 */
class CurrenciesValue extends \Algolia\AlgoliaSearch\Model\AbstractModel implements ModelInterface, \ArrayAccess, \JsonSerializable
{
    /**
     * Array of property to type mappings. Used for (de)serialization.
     *
     * @var string[]
     */
    protected static $modelTypes = [
        'currency' => 'string',
        'revenue' => 'float',
        'mean' => 'float',
        'standardDeviation' => 'float',
    ];

    /**
     * Array of property to format mappings. Used for (de)serialization.
     *
     * @var string[]
     */
    protected static $modelFormats = [
        'currency' => null,
        'revenue' => 'double',
        'mean' => 'double',
        'standardDeviation' => 'double',
    ];

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name.
     *
     * @var string[]
     */
    protected static $attributeMap = [
        'currency' => 'currency',
        'revenue' => 'revenue',
        'mean' => 'mean',
        'standardDeviation' => 'standardDeviation',
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses).
     *
     * @var string[]
     */
    protected static $setters = [
        'currency' => 'setCurrency',
        'revenue' => 'setRevenue',
        'mean' => 'setMean',
        'standardDeviation' => 'setStandardDeviation',
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests).
     *
     * @var string[]
     */
    protected static $getters = [
        'currency' => 'getCurrency',
        'revenue' => 'getRevenue',
        'mean' => 'getMean',
        'standardDeviation' => 'getStandardDeviation',
    ];

    /**
     * Associative array for storing property values.
     *
     * @var mixed[]
     */
    protected $container = [];

    /**
     * Constructor.
     *
     * @param mixed[] $data Associated array of property values
     */
    public function __construct(array $data = null)
    {
        if (isset($data['currency'])) {
            $this->container['currency'] = $data['currency'];
        }
        if (isset($data['revenue'])) {
            $this->container['revenue'] = $data['revenue'];
        }
        if (isset($data['mean'])) {
            $this->container['mean'] = $data['mean'];
        }
        if (isset($data['standardDeviation'])) {
            $this->container['standardDeviation'] = $data['standardDeviation'];
        }
    }

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name.
     *
     * @return array
     */
    public static function attributeMap()
    {
        return self::$attributeMap;
    }

    /**
     * Array of property to type mappings. Used for (de)serialization.
     *
     * @return array
     */
    public static function modelTypes()
    {
        return self::$modelTypes;
    }

    /**
     * Array of property to format mappings. Used for (de)serialization.
     *
     * @return array
     */
    public static function modelFormats()
    {
        return self::$modelFormats;
    }

    /**
     * Array of attributes to setter functions (for deserialization of responses).
     *
     * @return array
     */
    public static function setters()
    {
        return self::$setters;
    }

    /**
     * Array of attributes to getter functions (for serialization of requests).
     *
     * @return array
     */
    public static function getters()
    {
        return self::$getters;
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        return [];
    }

    /**
     * Validate all the properties in the model
     * return true if all passed.
     *
     * @return bool True if all properties are valid
     */
    public function valid()
    {
        return 0 === count($this->listInvalidProperties());
    }

    /**
     * Gets currency.
     *
     * @return null|string
     */
    public function getCurrency()
    {
        return $this->container['currency'] ?? null;
    }

    /**
     * Sets currency.
     *
     * @param null|string $currency currency code
     *
     * @return self
     */
    public function setCurrency($currency)
    {
        $this->container['currency'] = $currency;

        return $this;
    }

    /**
     * Gets revenue.
     *
     * @return null|float
     */
    public function getRevenue()
    {
        return $this->container['revenue'] ?? null;
    }

    /**
     * Sets revenue.
     *
     * @param null|float $revenue revenue for this currency
     *
     * @return self
     */
    public function setRevenue($revenue)
    {
        $this->container['revenue'] = $revenue;

        return $this;
    }

    /**
     * Gets mean.
     *
     * @return null|float
     */
    public function getMean()
    {
        return $this->container['mean'] ?? null;
    }

    /**
     * Sets mean.
     *
     * @param null|float $mean mean for this currency
     *
     * @return self
     */
    public function setMean($mean)
    {
        $this->container['mean'] = $mean;

        return $this;
    }

    /**
     * Gets standardDeviation.
     *
     * @return null|float
     */
    public function getStandardDeviation()
    {
        return $this->container['standardDeviation'] ?? null;
    }

    /**
     * Sets standardDeviation.
     *
     * @param null|float $standardDeviation standard deviation for this currency
     *
     * @return self
     */
    public function setStandardDeviation($standardDeviation)
    {
        $this->container['standardDeviation'] = $standardDeviation;

        return $this;
    }

    /**
     * Returns true if offset exists. False otherwise.
     *
     * @param int $offset Offset
     *
     * @return bool
     */
    public function offsetExists($offset)
    {
        return isset($this->container[$offset]);
    }

    /**
     * Gets offset.
     *
     * @param int $offset Offset
     *
     * @return null|mixed
     */
    public function offsetGet($offset)
    {
        return $this->container[$offset] ?? null;
    }

    /**
     * Sets value based on offset.
     *
     * @param null|int $offset Offset
     * @param mixed    $value  Value to be set
     */
    public function offsetSet($offset, $value)
    {
        if (is_null($offset)) {
            $this->container[] = $value;
        } else {
            $this->container[$offset] = $value;
        }
    }

    /**
     * Unsets offset.
     *
     * @param int $offset Offset
     */
    public function offsetUnset($offset)
    {
        unset($this->container[$offset]);
    }
}
