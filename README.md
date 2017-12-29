# ExtendMap
Golang: extend map[string]interface{} that often use with JSON.

Need to merge 2 maps/json in golang.

For example:-

* map/json#1
```javascript
map[string]interface{}{
    "key1": "value1",
    "dup1": map[string]interface{}{
        "key2": "value2",
    },
}
```
```json
{
    "key1": "value1",
    "dup1": {
        "key2": "value2"
    }
}
```

* map/json#2
```javascript
map[string]interface{}{
    "key3": "value3",
    "dup1": map[string]interface{}{
        "key4": "value4",
    },
}
```
```json
{
    "key3": "value3",
    "dup1": {
        "key4": "value4"
    }
}
```

* Merged result:-
```javascript
map[string]interface{}{
    "key1": "value1",
    "key3": "value3",
    "dup1": map[string]interface{}{
        "key2": "value2",
        "key4": "value4",
    },
}
```
```json
{
    "key1": "value1",
    "key3": "value3",
    "dup1": {
        "key2": "value2",
        "key4": "value4"
    }
}
```

## Usage:
```javascript
m1 := map[string]interface{}{
    "key1": "value1",
    "dup1": map[string]interface{}{
        "key2": "value2",
    },
}
m2 := map[string]interface{}{
    "key3": "value3",
    "dup1": map[string]interface{}{
        "key4": "value4",
    },
}
result := ExtendMap(m1, m2)
```

See the complex thing in extend_map_test.go
