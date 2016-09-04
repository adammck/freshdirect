package cart

import "github.com/adammck/freshdirect/iface"

// {
//   "items": [
//     {
//       "salesUnit": "EA",
//       "quantity": "1",
//       "atcItemId": "atc_mea_pid_3335058_MEA3335058_picks_love",
//       "productId": "mea_pid_3335058",
//       "categoryId": "picks_love",
//       "skuCode": "MEA3335058",
//       "pageType": "PRES_PICKS"
//     }
//   ],
//   "eventSource": "BROWSE",
//   "coremetricsPageId": "What's Good: Fresh Deals",
//   "coremetricsPageContentHierarchy": "What's Good: Fresh Deals-_--_--_-",
//   "coremetricsVirtualCategory": ""
// }
type Request struct {
	Items                           []RequestItem `json:"items"`
	CoremetricsPageContentHierarchy string        `json:"coremetricsPageContentHierarchy"`
	CoremetricsPageId               string        `json:"coremetricsPageId"`
	CoremetricsVirtualCategory      string        `json:"coremetricsVirtualCategory"`
	EventSource                     string        `json:"eventSource"`
}
type RequestItem struct {
	AtcItemId  string `json:"atcItemId"`
	CategoryId string `json:"categoryId"`
	PageType   string `json:"pageType"`
	ProductId  string `json:"productId"`
	Quantity   string `json:"quantity"`
	SalesUnit  string `json:"salesUnit"`
	SkuCode    string `json:"skuCode"`
}

// {
//   "coremetrics": [
//     [
//       "cmCreateShopAction5Tag",
//       "mea_pid_3335058",
//       "Local Angus House Recipe Burger, Raised without Antibiotics",
//       "1.0",
//       "4.99",
//       "FDW_PICKS_LOVE",
//       "MEA3335058-_--_--_-REGULAR-_-28-_-TIERED-_-0.0-_-What's Good: Fresh Deals-_--_--_--_-What's Good: Fresh Deals-_-"
//     ],
//     [
//       "cmDisplayShops"
//     ]
//   ],
//   "atcResult": [
//     {
//       "itemId": "atc_mea_pid_3335058_MEA3335058_picks_love",
//       "inCartAmount": 1.0,
//       "message": "Added to Cart",
//       "status": "SUCCESS",
//       "cartlineId": 675254879,
//       "productId": "mea_pid_3335058",
//       "categoryId": "picks_love"
//     }
//   ],
//   "pendingPopupData": null,
//   "couponStatus": {},
//   "redirectUrl": null
// }
type Response struct {
	AtcResult []struct {
		CartlineId   float64 `json:"cartlineId"`
		CategoryId   string  `json:"categoryId"`
		InCartAmount float64 `json:"inCartAmount"`
		ItemId       string  `json:"itemId"`
		Message      string  `json:"message"`
		ProductId    string  `json:"productId"`
		Status       string  `json:"status"`
	} `json:"atcResult"`
	Coremetrics      [][]string  `json:"coremetrics"`
	CouponStatus     struct{}    `json:"couponStatus"`
	PendingPopupData interface{} `json:"pendingPopupData"`
	RedirectUrl      interface{} `json:"redirectUrl"`
}

func Add(api iface.API, sku string) {
	api.Post("addtocart", Request{
		Items: []RequestItem{
			RequestItem{
				Quantity:  "1",
				SalesUnit: "EA",
				SkuCode:   sku,
			},
		},
	})
}
