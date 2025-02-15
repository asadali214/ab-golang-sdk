package advancedbilling

import (
    "context"
    "fmt"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/maxio-com/ab-golang-sdk/errors"
    "github.com/maxio-com/ab-golang-sdk/models"
    "net/http"
)

// SubscriptionGroupStatusController represents a controller struct.
type SubscriptionGroupStatusController struct {
    baseController
}

// NewSubscriptionGroupStatusController creates a new instance of SubscriptionGroupStatusController.
// It takes a baseController as a parameter and returns a pointer to the SubscriptionGroupStatusController.
func NewSubscriptionGroupStatusController(baseController baseController) *SubscriptionGroupStatusController {
    subscriptionGroupStatusController := SubscriptionGroupStatusController{baseController: baseController}
    return &subscriptionGroupStatusController
}

// CancelSubscriptionsInGroup takes context, uid, body as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// This endpoint will immediately cancel all subscriptions within the specified group. The group is identified by it's `uid` passed in the URL. To successfully cancel the group, the primary subscription must be on automatic billing. The group members as well must be on automatic billing or they must be prepaid.
// In order to cancel a subscription group while also charging for any unbilled usage on metered or prepaid components, the `charge_unbilled_usage=true` parameter must be included in the request.
func (s *SubscriptionGroupStatusController) CancelSubscriptionsInGroup(
    ctx context.Context,
    uid string,
    body *models.CancelGroupedSubscriptionsRequest) (
    *http.Response,
    error) {
    req := s.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/subscription_groups/%v/cancel.json", uid),
    )
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    
    context, err := req.Call()
    if err != nil {
        return context.Response, err
    }
    return context.Response, err
}

// InitiateDelayedCancellationForGroup takes context, uid as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// This endpoint will schedule all subscriptions within the specified group to be canceled at the end of their billing period. The group is identified by it's uid passed in the URL.
// All subscriptions in the group must be on automatic billing in order to successfully cancel them, and the group must not be in a "past_due" state.
func (s *SubscriptionGroupStatusController) InitiateDelayedCancellationForGroup(
    ctx context.Context,
    uid string) (
    *http.Response,
    error) {
    req := s.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/subscription_groups/%v/delayed_cancel.json", uid),
    )
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    
    context, err := req.Call()
    if err != nil {
        return context.Response, err
    }
    return context.Response, err
}

// CancelDelayedCancellationForGroup takes context, uid as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Removing the delayed cancellation on a subscription group will ensure that the subscriptions do not get canceled at the end of the period. The request will reset the `cancel_at_end_of_period` flag to false on each member in the group.
func (s *SubscriptionGroupStatusController) CancelDelayedCancellationForGroup(
    ctx context.Context,
    uid string) (
    *http.Response,
    error) {
    req := s.prepareRequest(
      ctx,
      "DELETE",
      fmt.Sprintf("/subscription_groups/%v/delayed_cancel.json", uid),
    )
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    
    context, err := req.Call()
    if err != nil {
        return context.Response, err
    }
    return context.Response, err
}

// ReactivateSubscriptionGroup takes context, uid, body as parameters and
// returns an models.ApiResponse with models.ReactivateSubscriptionGroupResponse data and
// an error if there was an issue with the request or response.
// This endpoint will attempt to reactivate or resume a cancelled subscription group. Upon reactivation, any canceled invoices created after the beginning of the primary subscription's billing period will be reopened and payment will be attempted on them. If the subscription group is being reactivated (as opposed to resumed), new charges will also be assessed for the new billing period.
// Whether a subscription group is reactivated (a new billing period is created) or resumed (the current billing period is respected) will depend on the parameters that are sent with the request as well as the date of the request relative to the primary subscription's period.
// ## Reactivating within the current period
// If a subscription group is cancelled and reactivated within the primary subscription's current period, we can choose to either start a new billing period or maintain the existing one. If we want to maintain the existing billing period the `resume=true` option must be passed in request parameters.
// An exception to the above are subscriptions that are on calendar billing. These subscriptions cannot be reactivated within the current period. If the `resume=true` option is not passed the request will return an error.
// The `resume_members` option is ignored in this case. All eligible group members will be automatically resumed.
// ## Reactivating beyond the current period
// In this case, a subscription group can only be reactivated with a new billing period. If the `resume=true` option is passed it will be ignored.
// Member subscriptions can have billing periods that are longer than the primary (e.g. a monthly primary with annual group members). If the primary subscription in a group cannot be reactivated within the current period, but other group members can be, passing `resume_members=true` will resume the existing billing period for eligible group members. The primary subscription will begin a new billing period.
// For calendar billing subscriptions, the new billing period created will be a partial one, spanning from the date of reactivation to the next corresponding calendar renewal date.
func (s *SubscriptionGroupStatusController) ReactivateSubscriptionGroup(
    ctx context.Context,
    uid string,
    body *models.ReactivateSubscriptionGroupRequest) (
    models.ApiResponse[models.ReactivateSubscriptionGroupResponse],
    error) {
    req := s.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/subscription_groups/%v/reactivate.json", uid),
    )
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    
    var result models.ReactivateSubscriptionGroupResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ReactivateSubscriptionGroupResponse](decoder)
    return models.NewApiResponse(result, resp), err
}
