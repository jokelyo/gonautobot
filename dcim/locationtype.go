package dcim

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/josh-silvas/gonautobot/core"
)

type (
	// LocationType : Represents a location type in Nautobot.
	LocationType struct {
		ID           uuid.UUID      `json:"id"`
		ContentTypes []string       `json:"content_types"`
		Created      time.Time      `json:"created"`
		CustomFields map[string]any `json:"custom_fields"`
		Description  string         `json:"description"`
		Display      string         `json:"display"`
		LastUpdated  time.Time      `json:"last_updated"`
		Name         string         `json:"name"`
		NaturalSlug  string         `json:"natural_slug"`
		Nestable     bool           `json:"nestable"`
		NotesURL     string         `json:"notes_url"`
		ObjectType   string         `json:"object_type"`
		Parent       *LocationType  `json:"parent"`
		TreeDepth    *int           `json:"tree_depth"`
		URL          string         `json:"url"`
	}

	// NewLocationType : Represents a new location type to be created in Nautobot.
	NewLocationType struct {
		Name         string         `json:"name"`
		ContentTypes []string       `json:"content_types,omitempty"`
		CustomFields map[string]any `json:"custom_fields,omitempty"`
		Description  string         `json:"description,omitempty"`
		Nestable     bool           `json:"nestable,omitempty"`
		Parent       *string        `json:"parent,omitempty"`
	}
)

// LocationTypeGet : Get a LocationType by UUID identifier.
func (c *Client) LocationTypeGet(id uuid.UUID) (*LocationType, error) {
	if id == uuid.Nil {
		return nil, errors.New("LocationTypeGet.error.ID(ID is missing or nil)")
	}
	req, err := c.Request(http.MethodGet, fmt.Sprintf("dcim/location-types/%s/", id), nil, nil)
	if err != nil {
		return nil, err
	}

	ret := new(LocationType)
	return ret, c.UnmarshalDo(req, ret)
}

// LocationTypeFilter : Get a list of LocationTypes based on query parameters.
func (c *Client) LocationTypeFilter(q *url.Values) ([]LocationType, error) {
	locationTypes := make([]LocationType, 0)
	return locationTypes, core.Paginate[LocationType](c.Client, "dcim/location-types/", q, &locationTypes)
}

// LocationTypeAll : Get all LocationTypes in Nautobot.
func (c *Client) LocationTypeAll() ([]LocationType, error) {
	locationTypes := make([]LocationType, 0)
	return locationTypes, core.Paginate[LocationType](c.Client, "dcim/location-types/", nil, &locationTypes)
}

// LocationTypeCreate : Generate a new LocationType record in Nautobot.
func (c *Client) LocationTypeCreate(obj NewLocationType) (LocationType, error) {
	var lt LocationType
	req, err := c.Request(http.MethodPost, "dcim/location-types/", obj, nil)
	if err != nil {
		return lt, err
	}

	if err := c.UnmarshalDo(req, &lt); err != nil {
		return lt, fmt.Errorf("LocationTypeCreate.error.UnmarshalDo(%w)", err)
	}
	return lt, nil
}

// LocationTypeUpdate : Update an existing LocationType record in Nautobot.
func (c *Client) LocationTypeUpdate(id uuid.UUID, patch map[string]any) (LocationType, error) {
	var lt LocationType
	if id == uuid.Nil {
		return lt, errors.New("LocationTypeUpdate.error.ID(ID is missing or nil)")
	}
	req, err := c.Request(http.MethodPatch, fmt.Sprintf("dcim/location-types/%s/", id), patch, nil)
	if err != nil {
		return lt, err
	}
	if err := c.UnmarshalDo(req, &lt); err != nil {
		return lt, fmt.Errorf("LocationTypeUpdate.error.UnmarshalDo(%w)", err)
	}
	return lt, nil
}

// LocationTypeDelete : Delete a LocationType by UUID identifier.
func (c *Client) LocationTypeDelete(id uuid.UUID) error {
	if id == uuid.Nil {
		return errors.New("LocationTypeDelete.error.ID(ID is missing or nil)")
	}
	req, err := c.Request(http.MethodDelete, fmt.Sprintf("dcim/location-types/%s/", id), nil, nil)
	if err != nil {
		return err
	}
	return c.UnmarshalDo(req, nil)
}
