package byob

import (
	"time"
)

type Flaggable struct {
	FData map[string]interface{} `json:"data" mapstructure:"data" msgpack:"data"`
}

type FlaggableChangeset struct {
	Changes []FlaggableChange `json:"changes" mapstructure:"changes" msgpack:"changes"`
}

const (
	FCSet = iota
	FCDelete
)

type FlaggableChange struct {
	Type int         `json:"type" mapstructure:"type" msgpack:"type"`
	Key  string      `json:"key" mapstructure:"key" msgpack:"key"`
	Data interface{} `json:"data" mapstructure:"data" msgpack:"data"`
}

func NewFlaggableChangeset() *FlaggableChangeset {
	return &FlaggableChangeset{Changes: []FlaggableChange{}}
}

func (f *FlaggableChangeset) Set(key string, value interface{}) {
	f.Changes = append(f.Changes, FlaggableChange{
		Type: FCSet,
		Key:  key,
		Data: value,
	})
}

func (f *FlaggableChangeset) Delete(key string) {
	f.Changes = append(f.Changes, FlaggableChange{
		Type: FCDelete,
		Key:  key,
		Data: nil,
	})
}

func (f *FlaggableChangeset) Apply(flaggable *Flaggable) {
	for _, c := range f.Changes {
		if c.Type == FCSet {
			flaggable.Set(c.Key, c.Data)
		} else if c.Type == FCDelete {
			flaggable.Del(c.Key)
		}
	}
}

func NewFlaggable(data map[string]interface{}) Flaggable {
	if data == nil {
		data = make(map[string]interface{})
	}

	return Flaggable{FData: data}
}

func (f *Flaggable) Data() map[string]interface{} {
	return f.FData
}

func (f *Flaggable) SetData(data map[string]interface{}) {
	f.FData = data
}

// Set is used to store a new key/value pair exclusively for this flaggable.
// It also lazy initializes f.Keys if it was not used previously.
func (f *Flaggable) Set(key string, value interface{}) {
	if f.FData == nil {
		f.FData = make(map[string]interface{})
	}

	f.FData[key] = value
}

// Present is used to check for the existence of a key
func (f *Flaggable) Present(key string) (present bool) {
	_, present = f.FData[key]
	return
}

// Get returns the value for a key
// If the value does not exist, it returns (nil, false)
func (f *Flaggable) Get(key string) (value interface{}, present bool) {
	if f.FData == nil {
		f.FData = make(map[string]interface{})
	}

	value, present = f.FData[key]
	return
}

func (f *Flaggable) Del(key string) {
	if f.FData == nil {
		f.FData = make(map[string]interface{})
	}

	delete(f.FData, key)
}

func (f *Flaggable) Clear() {
	f.FData = make(map[string]interface{})
}

// GetString returns the value associated with the key as a string.
func (f *Flaggable) GetString(key string) (s string) {
	if val, ok := f.Get(key); ok && val != nil {
		s, _ = val.(string)
	}
	return
}

// GetBool returns the value associated with the key as a boolean.
func (f *Flaggable) GetBool(key string) (b bool) {
	if val, ok := f.Get(key); ok && val != nil {
		b, _ = val.(bool)
	}
	return
}

// GetInt returns the value associated with the key as an integer.
func (f *Flaggable) GetInt(key string) (i int) {
	if val, ok := f.Get(key); ok && val != nil {
		i, _ = val.(int)
	}
	return
}

// GetInt64 returns the value associated with the key as an integer.
func (f *Flaggable) GetInt64(key string) (i64 int64) {
	if val, ok := f.Get(key); ok && val != nil {
		i64, _ = val.(int64)
	}
	return
}

// GetFloat64 returns the value associated with the key as a float64.
func (f *Flaggable) GetFloat64(key string) (f64 float64) {
	if val, ok := f.Get(key); ok && val != nil {
		f64, _ = val.(float64)
	}
	return
}

// GetTime returns the value associated with the key as time.
func (f *Flaggable) GetTime(key string) (t time.Time) {
	if val, ok := f.Get(key); ok && val != nil {
		t, _ = val.(time.Time)
	}
	return
}

// GetDuration returns the value associated with the key as a duration.
func (f *Flaggable) GetDuration(key string) (d time.Duration) {
	if val, ok := f.Get(key); ok && val != nil {
		d, _ = val.(time.Duration)
	}
	return
}

// GetStringSlice returns the value associated with the key as a slice of strings.
func (f *Flaggable) GetStringSlice(key string) (ss []string) {
	if val, ok := f.Get(key); ok && val != nil {
		ss, _ = val.([]string)
	}
	return
}

// GetStringMap returns the value associated with the key as a map of interfaces.
func (f *Flaggable) GetStringMap(key string) (sm map[string]interface{}) {
	if val, ok := f.Get(key); ok && val != nil {
		sm, _ = val.(map[string]interface{})
	}
	return
}

// GetStringMapString returns the value associated with the key as a map of strings.
func (f *Flaggable) GetStringMapString(key string) (sms map[string]string) {
	if val, ok := f.Get(key); ok && val != nil {
		sms, _ = val.(map[string]string)
	}
	return
}

