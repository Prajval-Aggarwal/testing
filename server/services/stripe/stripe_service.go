package stripeservice

import (
	"main/server/response"
	"main/server/utils"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
	"github.com/stripe/stripe-go/v72/paymentmethod"
)

// stripe payment service
func StripePayment(OrderPrice int64, cardNumber string, expMonth string, expYear string, cvc string, ctx *gin.Context) (pi, pi1 *stripe.PaymentIntent) {

	stripe.Key = "sk_test_51MvCYxSGxKXiPagaKdfa8MM2nYhjysJ41IUESqCLjca0meSTlzal4wbqMFZDbpTa5w1YXvdwygU8yMYbBecfgLCC00Yrx2WfFF"

	pm, err := paymentmethod.New(&stripe.PaymentMethodParams{
		Type: stripe.String("card"),
		Card: &stripe.PaymentMethodCardParams{
			Number:   stripe.String(cardNumber),
			ExpMonth: stripe.String(expMonth),
			ExpYear:  stripe.String(expYear),
			CVC:      stripe.String(cvc),
		},
	})
	if err != nil {
		response.ShowResponse("Error creating card details", utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	params := &stripe.PaymentIntentParams{
		Amount:             stripe.Int64(int64(OrderPrice) * 100),
		Currency:           stripe.String("inr"),
		Description:        stripe.String("Payment"),
		PaymentMethod:      stripe.String(pm.ID),
		ConfirmationMethod: stripe.String(string(stripe.PaymentIntentConfirmationMethodAutomatic)),
	}
	pi, err = paymentintent.New(params)
	if err != nil {
		response.ShowResponse("Error processing payment", utils.HTTP_SERVICE_UNAVAILABLE, utils.FAILURE, nil, ctx)
		return
	}

	params1 := &stripe.PaymentIntentConfirmParams{
		PaymentMethod: stripe.String("pm_card_visa"),
		CaptureMethod: stripe.String(string(stripe.PaymentIntentConfirmationMethodAutomatic)),
	}

	pi1, err = paymentintent.Confirm(pi.ID, params1)
	if err != nil {
		response.ShowResponse("Error confirming payment", utils.HTTP_SERVICE_UNAVAILABLE, utils.FAILURE, nil, ctx)
		return
	}

	switch pi1.Status {
	case "succeeded":
		response.ShowResponse("Success", utils.HTTP_OK, "Payment processed Successfully", "", ctx)
		return
	case "requires_payment_method":
		response.ShowResponse("Requires Payment Method", utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	case "requires_action":

		if pi1.Status == "requires_action" && pi1.NextAction != nil {
			switch pi1.NextAction.Type {
			case "use_stripe_sdk":

				response.ShowResponse(
					"Success",
					int64(utils.HTTP_OK),
					"Payment processed Successfully , Here is your client secret",
					pi1.ClientSecret,
					ctx,
				)
			}
		}
	default:
		response.ShowResponse("Payment requires more actions", utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	return pi, pi1

}
