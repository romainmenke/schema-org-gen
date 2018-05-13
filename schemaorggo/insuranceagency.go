package schemaorggo

import "encoding/json"

// InsuranceAgency see : https://schema.org/InsuranceAgency
type InsuranceAgency struct {
	FinancialService

	typeContext

	// FeesAndCommissionsSpecification see : https://schema.org/feesAndCommissionsSpecification
	// Description of fees, commissions, and other terms applied either to a class of financial product, or by a financial service organization.
	// types : Text URL
	FeesAndCommissionsSpecification []string `json:"feesAndCommissionsSpecification,omitempty"`
}

func (v InsuranceAgency) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.FinancialService.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.FeesAndCommissionsSpecification
		if len(v.FeesAndCommissionsSpecification) == 1 {
			value = v.FeesAndCommissionsSpecification[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["feesAndCommissionsSpecification"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v InsuranceAgency) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "InsuranceAgency"

	return data, nil
}

func (v InsuranceAgency) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
