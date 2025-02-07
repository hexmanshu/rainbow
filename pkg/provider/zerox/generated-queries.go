// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package zerox

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

// getOptionsOtokensOToken includes the requested fields of the GraphQL type OToken.
type getOptionsOtokensOToken struct {
	Id              string                                      `json:"id"`
	Symbol          string                                      `json:"symbol"`
	Name            string                                      `json:"name"`
	Decimals        int                                         `json:"decimals"`
	StrikeAsset     getOptionsOtokensOTokenStrikeAssetERC20     `json:"strikeAsset"`
	UnderlyingAsset getOptionsOtokensOTokenUnderlyingAssetERC20 `json:"underlyingAsset"`
	CollateralAsset getOptionsOtokensOTokenCollateralAssetERC20 `json:"collateralAsset"`
	StrikePrice     string                                      `json:"strikePrice"`
	IsPut           bool                                        `json:"isPut"`
	ExpiryTimestamp string                                      `json:"expiryTimestamp"`
}

// GetId returns getOptionsOtokensOToken.Id, and is useful for accessing the field via an interface.
func (v *getOptionsOtokensOToken) GetId() string { return v.Id }

// GetSymbol returns getOptionsOtokensOToken.Symbol, and is useful for accessing the field via an interface.
func (v *getOptionsOtokensOToken) GetSymbol() string { return v.Symbol }

// GetName returns getOptionsOtokensOToken.Name, and is useful for accessing the field via an interface.
func (v *getOptionsOtokensOToken) GetName() string { return v.Name }

// GetDecimals returns getOptionsOtokensOToken.Decimals, and is useful for accessing the field via an interface.
func (v *getOptionsOtokensOToken) GetDecimals() int { return v.Decimals }

// GetStrikeAsset returns getOptionsOtokensOToken.StrikeAsset, and is useful for accessing the field via an interface.
func (v *getOptionsOtokensOToken) GetStrikeAsset() getOptionsOtokensOTokenStrikeAssetERC20 {
	return v.StrikeAsset
}

// GetUnderlyingAsset returns getOptionsOtokensOToken.UnderlyingAsset, and is useful for accessing the field via an interface.
func (v *getOptionsOtokensOToken) GetUnderlyingAsset() getOptionsOtokensOTokenUnderlyingAssetERC20 {
	return v.UnderlyingAsset
}

// GetCollateralAsset returns getOptionsOtokensOToken.CollateralAsset, and is useful for accessing the field via an interface.
func (v *getOptionsOtokensOToken) GetCollateralAsset() getOptionsOtokensOTokenCollateralAssetERC20 {
	return v.CollateralAsset
}

// GetStrikePrice returns getOptionsOtokensOToken.StrikePrice, and is useful for accessing the field via an interface.
func (v *getOptionsOtokensOToken) GetStrikePrice() string { return v.StrikePrice }

// GetIsPut returns getOptionsOtokensOToken.IsPut, and is useful for accessing the field via an interface.
func (v *getOptionsOtokensOToken) GetIsPut() bool { return v.IsPut }

// GetExpiryTimestamp returns getOptionsOtokensOToken.ExpiryTimestamp, and is useful for accessing the field via an interface.
func (v *getOptionsOtokensOToken) GetExpiryTimestamp() string { return v.ExpiryTimestamp }

// getOptionsOtokensOTokenCollateralAssetERC20 includes the requested fields of the GraphQL type ERC20.
type getOptionsOtokensOTokenCollateralAssetERC20 struct {
	Id       string `json:"id"`
	Symbol   string `json:"symbol"`
	Decimals int    `json:"decimals"`
}

// GetId returns getOptionsOtokensOTokenCollateralAssetERC20.Id, and is useful for accessing the field via an interface.
func (v *getOptionsOtokensOTokenCollateralAssetERC20) GetId() string { return v.Id }

// GetSymbol returns getOptionsOtokensOTokenCollateralAssetERC20.Symbol, and is useful for accessing the field via an interface.
func (v *getOptionsOtokensOTokenCollateralAssetERC20) GetSymbol() string { return v.Symbol }

// GetDecimals returns getOptionsOtokensOTokenCollateralAssetERC20.Decimals, and is useful for accessing the field via an interface.
func (v *getOptionsOtokensOTokenCollateralAssetERC20) GetDecimals() int { return v.Decimals }

// getOptionsOtokensOTokenStrikeAssetERC20 includes the requested fields of the GraphQL type ERC20.
type getOptionsOtokensOTokenStrikeAssetERC20 struct {
	Id       string `json:"id"`
	Symbol   string `json:"symbol"`
	Decimals int    `json:"decimals"`
}

// GetId returns getOptionsOtokensOTokenStrikeAssetERC20.Id, and is useful for accessing the field via an interface.
func (v *getOptionsOtokensOTokenStrikeAssetERC20) GetId() string { return v.Id }

// GetSymbol returns getOptionsOtokensOTokenStrikeAssetERC20.Symbol, and is useful for accessing the field via an interface.
func (v *getOptionsOtokensOTokenStrikeAssetERC20) GetSymbol() string { return v.Symbol }

// GetDecimals returns getOptionsOtokensOTokenStrikeAssetERC20.Decimals, and is useful for accessing the field via an interface.
func (v *getOptionsOtokensOTokenStrikeAssetERC20) GetDecimals() int { return v.Decimals }

// getOptionsOtokensOTokenUnderlyingAssetERC20 includes the requested fields of the GraphQL type ERC20.
type getOptionsOtokensOTokenUnderlyingAssetERC20 struct {
	Id       string `json:"id"`
	Symbol   string `json:"symbol"`
	Decimals int    `json:"decimals"`
}

// GetId returns getOptionsOtokensOTokenUnderlyingAssetERC20.Id, and is useful for accessing the field via an interface.
func (v *getOptionsOtokensOTokenUnderlyingAssetERC20) GetId() string { return v.Id }

// GetSymbol returns getOptionsOtokensOTokenUnderlyingAssetERC20.Symbol, and is useful for accessing the field via an interface.
func (v *getOptionsOtokensOTokenUnderlyingAssetERC20) GetSymbol() string { return v.Symbol }

// GetDecimals returns getOptionsOtokensOTokenUnderlyingAssetERC20.Decimals, and is useful for accessing the field via an interface.
func (v *getOptionsOtokensOTokenUnderlyingAssetERC20) GetDecimals() int { return v.Decimals }

// getOptionsResponse is returned by getOptions on success.
type getOptionsResponse struct {
	Otokens []getOptionsOtokensOToken `json:"otokens"`
}

// GetOtokens returns getOptionsResponse.Otokens, and is useful for accessing the field via an interface.
func (v *getOptionsResponse) GetOtokens() []getOptionsOtokensOToken { return v.Otokens }

func getOptions(
	ctx context.Context,
	client graphql.Client,
) (*getOptionsResponse, error) {
	var err error

	var retval getOptionsResponse
	err = client.MakeRequest(
		ctx,
		"getOptions",
		`
query getOptions {
	otokens {
		id
		symbol
		name
		decimals
		strikeAsset {
			id
			symbol
			decimals
		}
		underlyingAsset {
			id
			symbol
			decimals
		}
		collateralAsset {
			id
			symbol
			decimals
		}
		strikePrice
		isPut
		expiryTimestamp
	}
}
`,
		&retval,
		nil,
	)
	return &retval, err
}
