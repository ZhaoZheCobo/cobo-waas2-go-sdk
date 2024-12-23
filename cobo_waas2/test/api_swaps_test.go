/*
Cobo Wallet as a Service 2.0

Testing SwapsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package cobo_waas2

import (
	"context"
	"testing"

	coboWaas2 "github.com/CoboGlobal/cobo-waas2-go-sdk/cobo_waas2"
	"github.com/CoboGlobal/cobo-waas2-go-sdk/cobo_waas2/crypto"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func Test_cobo_waas2_SwapsAPIService(t *testing.T) {

	configuration := coboWaas2.NewConfiguration()
	apiClient := coboWaas2.NewAPIClient(configuration)
	ctx := context.WithValue(context.Background(), coboWaas2.ContextServerHost, "https://api[.xxxx].cobo.com/v2")
	ctx = context.WithValue(ctx, coboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})

	t.Run("Test SwapsAPIService CreateQuote", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SwapsAPI.CreateQuote(ctx).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwapsAPIService CreateSwapActivity", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SwapsAPI.CreateSwapActivity(ctx).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwapsAPIService GetSwapActivity", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var activityId string

		resp, httpRes, err := apiClient.SwapsAPI.GetSwapActivity(ctx, activityId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwapsAPIService GetSwapSummary", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SwapsAPI.GetSwapSummary(ctx).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwapsAPIService ListEnableTokenPairs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SwapsAPI.ListEnableTokenPairs(ctx).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwapsAPIService ListSwapActivities", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SwapsAPI.ListSwapActivities(ctx).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}