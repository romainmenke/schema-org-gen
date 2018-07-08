package schemaorggo

import "encoding/json"

// WPSideBar see : https://schema.org/WPSideBar
type WPSideBar struct {
	WebPageElement

	typeContext

	// CssSelector see : https://pending.schema.org/cssSelector
	// A CSS selector, e.g. of a SpeakableSpecification (see: https://schema.org/SpeakableSpecification) or WebPageElement (see: https://schema.org/WebPageElement). In the latter case, multiple matches within a page can constitute a single conceptual &quot;Web page element&quot;.
	// types : CssSelectorType
	CssSelector []interface{} `json:"cssSelector,omitempty"`

	// Xpath see : https://pending.schema.org/xpath
	// An XPath, e.g. of a SpeakableSpecification (see: https://schema.org/SpeakableSpecification) or WebPageElement (see: https://schema.org/WebPageElement). In the latter case, multiple matches within a page can constitute a single conceptual &quot;Web page element&quot;.
	// types : XPathType
	Xpath []interface{} `json:"xpath,omitempty"`
}

func (v WPSideBar) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.WebPageElement.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.CssSelector
		if len(v.CssSelector) == 1 {
			value = v.CssSelector[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["cssSelector"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Xpath
		if len(v.Xpath) == 1 {
			value = v.Xpath[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["xpath"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v WPSideBar) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "WPSideBar"

	return data, nil
}

func (v WPSideBar) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
