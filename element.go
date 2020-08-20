package jamswrapper

import "fmt"

type Element struct {
	ElementName         string          `json:"elementName"`
	ElementTypeName     string          `json:"elementTypeName"`
	ElementTypeID       int             `json:"elementTypeId"`
	ElementID           int             `json:"elementId"`
	ElementUID          string          `json:"elementUid"`
	OwnerUID            string          `json:"ownerUid"`
	Inherited           bool            `json:"inherited"`
	InheritedFromType   string          `json:"inheritedFromType"`
	InheritedFromID     int             `json:"inheritedFromId"`
	InheritedFromName   string          `json:"inheritedFromName"`
	ElementKind         string          `json:"elementKind"`
	ElementState        string          `json:"elementState"`
	Properties          []PropertyValue `json:"properties"`
	WriteToEntry        bool            `json:"writeToEntry"`
	WriteToHistory      bool            `json:"writeToHistory"`
	TabName             string          `json:"tabName"`
	SortOrder           int             `json:"sortOrder"`
	ValidateTypeNameSSO string          `json:"validateTypeNameSSO"`
	FormatString        string          `json:"formatString"`
	AddMenuText         string          `json:"addMenuText"`
	CategoryName        string          `json:"categoryName"`
	CategoryAddMenuText string          `json:"categoryAddMenuText"`
	CategorySortOrder   int             `json:"categorySortOrder"`
}

func GetElementsFor(connection *AuthResponse, elementOf string) ([]Element, error) {
	res := make([]Element, 50)
	target := fmt.Sprintf("%s/element/%s", destinationServer, elementOf)
	request := setupClient(connection)
	result, err := request.SetResult(&res).Get(target)
	err = manageResponse(result, err)
	return res, err
}
