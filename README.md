# kinvo-investments

Project to parse https://app.kinvo.com.br/meus-produtos export.

The result will be a suggestion to balance all products.

You can create a second sheet to set products to filter final results.

## Build

```shell
go build internal/cmd/kinvo.go
```

## Usage

**From sources**

```shell
go run internal/cmd/kinvo.go -f file.xlsx -b brokerNameMatcher -c assetClassMatcher
```

**From binary**

Windows:
```shell
./kinvo.exe /f "file.xlsx" /b brokerNameMatcher /c assetClassMatcher
```

Linux:
```shell
./kinvo -f "file.xlsx" -b brokerNameMatcher -c assetClassMatcher
```