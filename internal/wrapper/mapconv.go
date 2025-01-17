package wrapper

import (
	"strconv"
	"strings"
)

type typeConversionMap map[string]any

func (cmap typeConversionMap) getKey(key string) (any, bool) {
	keys := strings.SplitN(key, ".", 1)
	if len(keys) == 1 {
		val, ok := cmap[keys[0]]
		return val, ok
	}
	value, ok := cmap[keys[0]]
	if !ok {
		return nil, false
	}
	if kv_, ok := value.(typeConversionMap); ok {
		return kv_.getKey(keys[1])
	}
	return nil, false
}

func (cmap typeConversionMap) GetInt8(key string, def int8) int8 {
	value, ok := cmap.getKey(key)
	if !ok {
		return def
	}
	switch val := value.(type) {
	case string:
		intVal, err := strconv.ParseInt(val, 10, 0)
		if err != nil {
			return def
		}
		return int8(intVal)
	case int8:
		return val
	}
	return def
}

func (cmap typeConversionMap) GetUint8(key string, def uint8) uint8 {
	value, ok := cmap.getKey(key)
	if !ok {
		return def
	}
	switch val := value.(type) {
	case string:
		intVal, err := strconv.ParseUint(val, 10, 0)
		if err != nil {
			return def
		}
		return uint8(intVal)
	case uint8:
		return val
	}
	return def
}

func (cmap typeConversionMap) GetInt16(key string, def int16) int16 {
	value, ok := cmap.getKey(key)
	if !ok {
		return def
	}
	switch val := value.(type) {
	case string:
		intVal, err := strconv.ParseInt(val, 10, 0)
		if err != nil {
			return def
		}
		return int16(intVal)
	case int16:
		return val
	}
	return def
}

func (cmap typeConversionMap) GetUint16(key string, def uint16) uint16 {
	value, ok := cmap.getKey(key)
	if !ok {
		return def
	}
	switch val := value.(type) {
	case string:
		intVal, err := strconv.ParseUint(val, 10, 0)
		if err != nil {
			return def
		}
		return uint16(intVal)
	case uint16:
		return val
	}
	return def
}

func (cmap typeConversionMap) GetInt32(key string, def int32) int32 {
	value, ok := cmap.getKey(key)
	if !ok {
		return def
	}
	switch val := value.(type) {
	case string:
		intVal, err := strconv.ParseInt(val, 10, 0)
		if err != nil {
			return def
		}
		return int32(intVal)
	case int32:
		return val
	}
	return def
}

func (cmap typeConversionMap) GetUint32(key string, def uint32) uint32 {
	value, ok := cmap.getKey(key)
	if !ok {
		return def
	}
	switch val := value.(type) {
	case string:
		intVal, err := strconv.ParseUint(val, 10, 0)
		if err != nil {
			return def
		}
		return uint32(intVal)
	case uint32:
		return val
	}
	return def
}

func (cmap typeConversionMap) GetInt64(key string, def int64) int64 {
	value, ok := cmap.getKey(key)
	if !ok {
		return def
	}
	switch val := value.(type) {
	case string:
		intVal, err := strconv.ParseInt(val, 10, 0)
		if err != nil {
			return def
		}
		return int64(intVal)
	case int64:
		return val
	}
	return def
}

func (cmap typeConversionMap) GetUint64(key string, def uint64) uint64 {
	value, ok := cmap.getKey(key)
	if !ok {
		return def
	}
	switch val := value.(type) {
	case string:
		intVal, err := strconv.ParseUint(val, 10, 0)
		if err != nil {
			return def
		}
		return uint64(intVal)
	case uint64:
		return val
	}
	return def
}

func (cmap typeConversionMap) GetBool(key string, def bool) bool {
	value, ok := cmap.getKey(key)
	if !ok {
		return def
	}
	switch val := value.(type) {
	case string:
		boolVal, err := strconv.ParseBool(val)
		if err != nil {
			return def
		}
		return boolVal
	case bool:
		return val
	}
	return def
}

func (cmap typeConversionMap) GetString(key string, def string) string {
	value, ok := cmap.getKey(key)
	if !ok {
		return def
	}
	switch val := value.(type) {
	case string:
		return val
	case int8:
		return strconv.FormatInt(int64(val), 10)
	case uint8:
		return strconv.FormatUint(uint64(val), 10)
	case int16:
		return strconv.FormatInt(int64(val), 10)
	case uint16:
		return strconv.FormatUint(uint64(val), 10)
	case int32:
		return strconv.FormatInt(int64(val), 10)
	case uint32:
		return strconv.FormatUint(uint64(val), 10)
	case int64:
		return strconv.FormatInt(int64(val), 10)
	case uint64:
		return strconv.FormatUint(uint64(val), 10)
	case bool:
		return strconv.FormatBool(val)
	}
	return def
}

func (cmap typeConversionMap) GetMap(key string, def Map) Map {
	value, ok := cmap.getKey(key)
	if !ok {
		return def
	}
	if val, ok := value.(map[string]any); ok {
		return typeConversionMap(val)
	}
	return def
}
