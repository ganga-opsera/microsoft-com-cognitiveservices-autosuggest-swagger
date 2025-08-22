package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// ErrorResponse represents the ErrorResponse schema from the OpenAPI specification
type ErrorResponse struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Potentialaction []Action `json:"potentialAction,omitempty"`
	Preferredclickthroughurl string `json:"preferredClickthroughUrl,omitempty"`
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Adaptivecard string `json:"adaptiveCard,omitempty"`
	Immediateaction []Action `json:"immediateAction,omitempty"`
}

// Response represents the Response schema from the OpenAPI specification
type Response struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
}

// Suggestions represents the Suggestions schema from the OpenAPI specification
type Suggestions struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Immediateaction []Action `json:"immediateAction,omitempty"`
	Potentialaction []Action `json:"potentialAction,omitempty"`
	Preferredclickthroughurl string `json:"preferredClickthroughUrl,omitempty"`
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Adaptivecard string `json:"adaptiveCard,omitempty"`
	Querycontext QueryContext `json:"queryContext,omitempty"` // Defines the query context that Bing used for the request.
}

// Identifiable represents the Identifiable schema from the OpenAPI specification
type Identifiable struct {
	TypeField string `json:"_type"`
}

// ResponseBase represents the ResponseBase schema from the OpenAPI specification
type ResponseBase struct {
	TypeField string `json:"_type"`
}

// SearchAction represents the SearchAction schema from the OpenAPI specification
type SearchAction struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Potentialaction []Action `json:"potentialAction,omitempty"`
	Preferredclickthroughurl string `json:"preferredClickthroughUrl,omitempty"`
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Adaptivecard string `json:"adaptiveCard,omitempty"`
	Immediateaction []Action `json:"immediateAction,omitempty"`
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Disclaimer string `json:"disclaimer,omitempty"`
	About []Thing `json:"about,omitempty"` // For internal use only.
	Copyrightyear int `json:"copyrightYear,omitempty"`
	Thumbnailurl string `json:"thumbnailUrl,omitempty"` // The URL to a thumbnail of the item.
	Copyrightholder Thing `json:"copyrightHolder,omitempty"` // Defines a thing.
	Creator Thing `json:"creator,omitempty"` // Defines a thing.
	Provider []Thing `json:"provider,omitempty"` // The source of the creative work.
	Genre []string `json:"genre,omitempty"`
	Mainentity Thing `json:"mainEntity,omitempty"` // Defines a thing.
	Commentcount int `json:"commentCount,omitempty"`
	Discussionurl string `json:"discussionUrl,omitempty"`
	Mentions []Thing `json:"mentions,omitempty"` // For internal use only.
	Text string `json:"text,omitempty"` // Text content of this creative work
	Isaccessibleforfree bool `json:"isAccessibleForFree,omitempty"`
	Isfamilyfriendly bool `json:"isFamilyFriendly,omitempty"`
	Headline string `json:"headLine,omitempty"`
	Serviceurl string `json:"serviceUrl,omitempty"`
	Displayname string `json:"displayName,omitempty"`
	Istopaction bool `json:"isTopAction,omitempty"`
	Result []Thing `json:"result,omitempty"`
}

// Thing represents the Thing schema from the OpenAPI specification
type Thing struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Adaptivecard string `json:"adaptiveCard,omitempty"`
	Immediateaction []Action `json:"immediateAction,omitempty"`
	Potentialaction []Action `json:"potentialAction,omitempty"`
	Preferredclickthroughurl string `json:"preferredClickthroughUrl,omitempty"`
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource.
}

// CreativeWork represents the CreativeWork schema from the OpenAPI specification
type CreativeWork struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Potentialaction []Action `json:"potentialAction,omitempty"`
	Preferredclickthroughurl string `json:"preferredClickthroughUrl,omitempty"`
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Adaptivecard string `json:"adaptiveCard,omitempty"`
	Immediateaction []Action `json:"immediateAction,omitempty"`
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
}

// Answer represents the Answer schema from the OpenAPI specification
type Answer struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Potentialaction []Action `json:"potentialAction,omitempty"`
	Preferredclickthroughurl string `json:"preferredClickthroughUrl,omitempty"`
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Adaptivecard string `json:"adaptiveCard,omitempty"`
	Immediateaction []Action `json:"immediateAction,omitempty"`
}

// QueryContext represents the QueryContext schema from the OpenAPI specification
type QueryContext struct {
	Originalquery string `json:"originalQuery"` // The query string as specified in the request.
	TypeField string `json:"_type"`
	Adultintent bool `json:"adultIntent,omitempty"` // A Boolean value that indicates whether the specified query has adult intent. The value is true if the query has adult intent; otherwise, false.
	Alterationoverridequery string `json:"alterationOverrideQuery,omitempty"` // The query string to use to force Bing to use the original string. For example, if the query string is "saling downwind", the override query string will be "+saling downwind". Remember to encode the query string which results in "%2Bsaling+downwind". This field is included only if the original query string contains a spelling mistake.
	Alteredquery string `json:"alteredQuery,omitempty"` // The query string used by Bing to perform the query. Bing uses the altered query string if the original query string contained spelling mistakes. For example, if the query string is "saling downwind", the altered query string will be "sailing downwind". This field is included only if the original query string contains a spelling mistake.
	Askuserforlocation bool `json:"askUserForLocation,omitempty"` // A Boolean value that indicates whether Bing requires the user's location to provide accurate results. If you specified the user's location by using the X-MSEdge-ClientIP and X-Search-Location headers, you can ignore this field. For location aware queries, such as "today's weather" or "restaurants near me" that need the user's location to provide accurate results, this field is set to true. For location aware queries that include the location (for example, "Seattle weather"), this field is set to false. This field is also set to false for queries that are not location aware, such as "best sellers".
	Istransactional bool `json:"isTransactional,omitempty"`
}

// SearchResultsAnswer represents the SearchResultsAnswer schema from the OpenAPI specification
type SearchResultsAnswer struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Immediateaction []Action `json:"immediateAction,omitempty"`
	Potentialaction []Action `json:"potentialAction,omitempty"`
	Preferredclickthroughurl string `json:"preferredClickthroughUrl,omitempty"`
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Adaptivecard string `json:"adaptiveCard,omitempty"`
}

// SuggestionsSuggestionGroup represents the SuggestionsSuggestionGroup schema from the OpenAPI specification
type SuggestionsSuggestionGroup struct {
	Searchsuggestions []SearchAction `json:"searchSuggestions"`
	TypeField string `json:"_type"`
	Name string `json:"name"`
}

// Action represents the Action schema from the OpenAPI specification
type Action struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Adaptivecard string `json:"adaptiveCard,omitempty"`
	Immediateaction []Action `json:"immediateAction,omitempty"`
	Potentialaction []Action `json:"potentialAction,omitempty"`
	Preferredclickthroughurl string `json:"preferredClickthroughUrl,omitempty"`
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Isaccessibleforfree bool `json:"isAccessibleForFree,omitempty"`
	Isfamilyfriendly bool `json:"isFamilyFriendly,omitempty"`
	Headline string `json:"headLine,omitempty"`
	Disclaimer string `json:"disclaimer,omitempty"`
	About []Thing `json:"about,omitempty"` // For internal use only.
	Copyrightyear int `json:"copyrightYear,omitempty"`
	Thumbnailurl string `json:"thumbnailUrl,omitempty"` // The URL to a thumbnail of the item.
	Copyrightholder Thing `json:"copyrightHolder,omitempty"` // Defines a thing.
	Creator Thing `json:"creator,omitempty"` // Defines a thing.
	Provider []Thing `json:"provider,omitempty"` // The source of the creative work.
	Genre []string `json:"genre,omitempty"`
	Mainentity Thing `json:"mainEntity,omitempty"` // Defines a thing.
	Commentcount int `json:"commentCount,omitempty"`
	Discussionurl string `json:"discussionUrl,omitempty"`
	Mentions []Thing `json:"mentions,omitempty"` // For internal use only.
	Text string `json:"text,omitempty"` // Text content of this creative work
}

// Error represents the Error schema from the OpenAPI specification
type Error struct {
	Value string `json:"value,omitempty"` // The parameter's value in the request that was not valid.
	TypeField string `json:"_type"`
	Code string `json:"code"` // The error code that identifies the category of error.
	Message string `json:"message"` // A description of the error.
	Moredetails string `json:"moreDetails,omitempty"` // A description that provides additional information about the error.
	Parameter string `json:"parameter,omitempty"` // The parameter in the request that caused the error.
}
