# RepCheck - чекер (ваших) зависимостей

- [Установка](#установка)
- [Использование](#использование)
- [Флаги](#флаги)
- [Пример 1](#пример-1)
- [Пример 2](#пример-2)

## Установка

```bash
go install github.com/vellun/repcheck
```

## Использование

```bash
repcheck [repo url] [flags]
```

## Флаги

```
--json        — вывод информации в формате json
--noindirect  — не выводить непрямые зависимости
```

## Пример 1

### Ввод

```bash
repcheck https://github.com/golang/example
```

### Вывод

```bash
Module: golang.org/x/example
Go version: 1.18

Deps to update:
golang.org/x/tools: v0.14.0 -> v0.33.0
golang.org/x/mod: v0.13.0 -> v0.24.0 // indirect
golang.org/x/sys: v0.13.0 -> v0.33.0 // indirect
```

## Пример 2

### Ввод

```bash
repcheck https://github.com/golang/example --json
```

### Вывод

```bash
{
  "module": {
    "Name": "golang.org/x/example",
    "GoVersion": "1.18"
  },
  "deps": [
    {
      "Path": "golang.org/x/tools",
      "CurVersion": "v0.14.0",
      "UpdateVersion": "v0.33.0",
      "IsIndirect": false
    },
    {
      "Path": "golang.org/x/mod",
      "CurVersion": "v0.13.0",
      "UpdateVersion": "v0.24.0",
      "IsIndirect": true
    },
    {
      "Path": "golang.org/x/sys",
      "CurVersion": "v0.13.0",
      "UpdateVersion": "v0.33.0",
      "IsIndirect": true
    }
  ]
}
```
