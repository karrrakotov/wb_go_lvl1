Сколько существует способов задать переменную типа slice или map?

Ответ:
```
Существует 3 способа:

1. Используя литералы:
Slice: s := []int{1, 2, 3}
Map: m := map[string]int{"a": 1, "b": 2}

2. Используя make:
Slice: s := make([]int, 5, 10)
Map: m := make(map[string]int)

3. Используя инициализацию пустой переменной:
Slice: var s []int
Map: var m map[string]int

```