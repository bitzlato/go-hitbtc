package hitbtc

import (
	"errors"
	"fmt"
)

var ErrMalformedErrorResponse = errors.New("malformed error response")

type APIError struct {
	Code        int    `json:"code"`
	Message     string `json:"message,omitempty"`
	Description string `json:"description,omitempty"`
}

func IsAPIError(v interface{}) bool {
	_, ok := v.(*APIError)
	return ok
}

func (e *APIError) Error() string {
	return fmt.Sprintf("HitBTC <APIError> code=%d, message=%q, description=%q", e.Code, e.Message, e.Description)
}

/*
   Error code	HTTP Status Code	Message	                                    Note
   403	        401	                Action is forbidden for account
   429	        429	                Too many requests	                        Action is being rate limited for account
   500	        500	                Internal Server Error
   503	        503	                Service Unavailable	                        Try it again later
   504	        504	                Gateway Timeout	                            Check the result of your request later
   1001	        401	                Authorization required
   1002	        401	                Authorization required or has been failed	Check that authorization data provided
   1003	        403	                Action forbidden for this API key	        Check permissions for API key
   1004	        401	                Unsupported authorization method	        Use Basic authentication
   2001	        400	                Symbol not found
   2002	        400	                Currency not found
   2010	        400	                Quantity not a valid number
   2011	        400	                Quantity too low
   2012	        400	                Bad quantity	                            More details in description
   2020	        400	                Price not a valid number
   2021	        400	                Price too low
   2022	        400	                Bad price	                                More details in description
   20001	    400	                Insufficient funds	                        Insufficient funds for creating order or any account operation
   20002	    400	                Order not found	                            Attempt to get active order that not existing, filled, canceled or expired. Attempt to cancel not existing order. Attempt to cancel already filled or expired order.
   20003	    400	                Limit exceeded	                            Withdrawal limit exceeded
   20004	    400	                Transaction not found	                    Requested transaction not found
   20005	    400	                Payout not found
   20006	    400	                Payout already committed
   20007	    400	                Payout already rolled back
   20008	    400	                Duplicate clientOrderId
   20009	    400	                Price and quantity not changed
   20010	    400	                Exchange temporary closed	                Exchange market temporary closed on symbol
   20011	    400	                Payout address is invalid
   20014	    400	                Offchain for this payout is unavailable	    Address do not belong to exchange or belongs to same user or unavailable for currency
   20032	    400	                Margin account or position not found	    Create margin account before place orders or account for requested symbol not found
   20033	    400	                Position not changed	                    Existed marginBalance equal requested value
   20034	    400	                Position in close only state
   20040	    400	                Margin trading forbidden
   20080	    400	                Internal order execution deadline exceeded	Order not placed.
   10001	    400	                Validation error	                        Input not valid, see more in message field
   10021	    400	                User disabled	                            Sub-Account API. User cant be changed
*/
