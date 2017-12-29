# extendMap
Golang: extend map[string]interface{} that often use with JSON.

Need to merge 2 maps/json in golang.

For example:-

* map/json#1
```javascript
map[string]interface{}{
    "key1": "value1",
    "key2": "value2",
}
```
```json
{
    "key1": "value1",
    "key2": "value2"
}
```

* map/json#2
```javascript
map[string]interface{}{
    "key1": "value1",
    "key2": "value2",
}
```
```json
{
    "key3": "value3"
}
```

* Merged result:-
```javascript
map[string]interface{}{
    "key1": "value1",
    "key2": "value2",
    "key3": "value3",
}
```
```json
{
    "key1": "value1",
    "key2": "value2",
    "key3": "value3"
}
```

So, we can use:-
```javascript
m1 := map[string]interface{}{
    "key1": "value1",
    "key2": "value2",
}
m2 := map[string]interface{}{
    "key3": "value3",
}
result := ExtendMap(m1, m2)
```

See the complex thing in extend_map_test.go
