package jamswrapper

import (
	"encoding/json"
	"fmt"
)

type PropertyValue struct {
	PropertyName      string          `json:"propertyName"`
	CategoryName      string          `json:"categoryName"`
	CategorySortOrder int             `json:"categorySortOrder"`
	DisplayName       string          `json:"displayName"`
	PropertyID        int             `json:"propertyId"`
	TypeName          string          `json:"typeName"`
	TypeNameSSO       string          `json:"typeNameSSO"`
	ToolTip           string          `json:"toolTip"`
	Description       string          `json:"description"`
	Editor            string          `json:"editor"`
	MergeOption       string          `json:"mergeOption"`
	SortOrder         int             `json:"sortOrder"`
	InheritedFromType string          `json:"inheritedFromType"`
	InheritedFromID   int             `json:"inheritedFromId"`
	InheritedFromName string          `json:"inheritedFromName"`
	ReadOnly          bool            `json:"readOnly"`
	Browsable         bool            `json:"browsable"`
	DefaultValue      json.RawMessage `json:"defaultValue,omitempty"`
	CurrentValue      json.RawMessage `json:"currentValue,omitempty"`
}

func GetPropertiesFor(connection *AuthResponse, propertyOf string) ([]PropertyValue, error) {
	res := make([]PropertyValue, 50)
	target := fmt.Sprintf("%s/propertydefinition/%s", destinationServer, propertyOf)
	request := setupClient(connection)
	result, err := request.SetResult(&res).Get(target)
	err = manageResponse(result, err)
	return res, err
}
