package checker

import (
	"github.com/treeverse/oasdiff/diff"
)

func RequestParameterMinItemsSetCheck(diffReport *diff.Diff, operationsSources *diff.OperationsSourcesMap, config Config) Changes {
	result := make(Changes, 0)
	if diffReport.PathsDiff == nil {
		return result
	}
	for path, pathItem := range diffReport.PathsDiff.Modified {
		if pathItem.OperationsDiff == nil {
			continue
		}
		for operation, operationItem := range pathItem.OperationsDiff.Modified {
			if operationItem.ParametersDiff == nil {
				continue
			}
			for paramLocation, paramDiffs := range operationItem.ParametersDiff.Modified {
				for paramName, paramDiff := range paramDiffs {
					if paramDiff.SchemaDiff == nil {
						continue
					}
					minItemsDiff := paramDiff.SchemaDiff.MinItemsDiff
					if minItemsDiff == nil {
						continue
					}
					if minItemsDiff.From != nil ||
						minItemsDiff.To == nil {
						continue
					}

					source := (*operationsSources)[operationItem.Revision]
					result = append(result, ApiChange{
						Id:          "request-parameter-min-items-set",
						Level:       WARN,
						Text:        config.Localize("request-parameter-min-items-set", ColorizedValue(paramLocation), ColorizedValue(paramName), ColorizedValue(minItemsDiff.To)),
						Comment:     config.Localize("request-parameter-min-items-set-comment"),
						Operation:   operation,
						OperationId: operationItem.Revision.OperationID,
						Path:        path,
						Source:      source,
					})
				}
			}
		}
	}
	return result
}
