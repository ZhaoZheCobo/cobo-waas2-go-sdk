/*
Cobo Wallet as a Service 2.0

Testing TransactionsAPIService

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

func Test_cobo_waas2_TransactionsAPIService(t *testing.T) {

	configuration := coboWaas2.NewConfiguration()
	apiClient := coboWaas2.NewAPIClient(configuration)
	ctx := context.WithValue(context.Background(), coboWaas2.ContextServerHost, "https://api[.xxxx].cobo.com/v2")
	ctx = context.WithValue(ctx, coboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})

	t.Run("Test TransactionsAPIService BroadcastSignedTransactions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TransactionsAPI.BroadcastSignedTransactions(ctx).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransactionsAPIService CancelTransactionById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var transactionId string

		resp, httpRes, err := apiClient.TransactionsAPI.CancelTransactionById(ctx, transactionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransactionsAPIService CheckLoopTransfers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TransactionsAPI.CheckLoopTransfers(ctx).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransactionsAPIService CreateContractCallTransaction", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TransactionsAPI.CreateContractCallTransaction(ctx).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransactionsAPIService CreateMessageSignTransaction", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TransactionsAPI.CreateMessageSignTransaction(ctx).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransactionsAPIService CreateTransferTransaction", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TransactionsAPI.CreateTransferTransaction(ctx).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransactionsAPIService DropTransactionById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var transactionId string

		resp, httpRes, err := apiClient.TransactionsAPI.DropTransactionById(ctx, transactionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransactionsAPIService EstimateFee", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TransactionsAPI.EstimateFee(ctx).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransactionsAPIService GetTransactionById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var transactionId string

		resp, httpRes, err := apiClient.TransactionsAPI.GetTransactionById(ctx, transactionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransactionsAPIService ListTransactions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TransactionsAPI.ListTransactions(ctx).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransactionsAPIService ResendTransactionById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var transactionId string

		resp, httpRes, err := apiClient.TransactionsAPI.ResendTransactionById(ctx, transactionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransactionsAPIService SpeedupTransactionById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var transactionId string

		resp, httpRes, err := apiClient.TransactionsAPI.SpeedupTransactionById(ctx, transactionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
