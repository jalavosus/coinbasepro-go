package coinbasepro

import (
	"strings"

	"github.com/pkg/errors"
)

//go:generate stringer -type=OrderType,OrderSide,StopOrderType,TimeInForceOption,SelfTradePreventionFlag -linecomment -output enums_string.go

type OrderType uint8

func (o OrderType) MarshalJSON() ([]byte, error) {
	return []byte(o.String()), nil
}

func (o *OrderType) UnmarshalJSON(data []byte) error {
	raw := string(data)
	switch strings.ToLower(raw) {
	case LimitOrder.String():
		*o = LimitOrder
	case MarketOrder.String():
		*o = MarketOrder
	default:
		return errors.Errorf("invalid OrderType %[1]s", raw)
	}

	return nil
}

const (
	// LimitOrder is a constant for the "limit" order type.
	LimitOrder OrderType = iota // limit
	// MarketOrder is a constant for the "market" order type.
	MarketOrder // market
)

type OrderSide uint8

func (o OrderSide) MarshalJSON() ([]byte, error) {
	return []byte(o.String()), nil
}

func (o *OrderSide) UnmarshalJSON(data []byte) error {
	raw := string(data)
	switch strings.ToLower(raw) {
	case BuyOrder.String():
		*o = BuyOrder
	case SellOrder.String():
		*o = SellOrder
	default:
		return errors.Errorf("invalid OrderSide %[1]s", raw)
	}

	return nil
}

const (
	BuyOrder  OrderSide = iota // buy
	SellOrder                  // sell
)

type StopOrderType uint8

func (s StopOrderType) MarshalJSON() ([]byte, error) {
	return []byte(s.String()), nil
}

func (s *StopOrderType) UnmarshalJSON(data []byte) error {
	raw := string(data)
	switch strings.ToLower(raw) {
	case StopLoss.String():
		*s = StopLoss
	case StopEntry.String():
		*s = StopEntry
	default:
		return errors.Errorf("invalid StopOrderType %[1]s", raw)
	}

	return nil
}

const (
	StopLoss  StopOrderType = iota // loss
	StopEntry                      // entry
)

type TimeInForceOption uint8

func (t TimeInForceOption) MarshalJSON() ([]byte, error) {
	return []byte(t.String()), nil
}

func (t *TimeInForceOption) UnmarshalJSON(data []byte) error {
	raw := string(data)
	switch strings.ToLower(raw) {
	case GTC.String():
		*t = GTC
	case GTT.String():
		*t = GTT
	case IOC.String():
		*t = IOC
	case FOK.String():
		*t = FOK
	default:
		return errors.Errorf("invalid TimeInForceOption %[1]s", raw)
	}

	return nil
}

const (
	// GTC Good till canceled
	GTC TimeInForceOption = iota // GTC
	// GTT Good till time
	GTT // GTT
	// IOC Immediate or cancel
	IOC // IOC
	// FOK Fill or kill
	FOK // FOK
)

type SelfTradePreventionFlag uint8

func (s SelfTradePreventionFlag) MarshalJSON() ([]byte, error) {
	return []byte(s.String()), nil
}

func (s *SelfTradePreventionFlag) UnmarshalJSON(data []byte) error {
	raw := string(data)
	switch strings.ToLower(raw) {
	case DC.String():
		*s = DC
	case CO.String():
		*s = CO
	case CN.String():
		*s = CN
	case CB.String():
		*s = CB
	default:
		return errors.Errorf("invalid SelfTradePreventionFlag %[1]s", raw)
	}

	return nil
}

const (
	// DC Decrease and cancel (default)
	DC SelfTradePreventionFlag = iota // dc
	// CO Cancel oldest
	CO // co
	// CN Cancel newest
	CN // cn
	// CB cancel both
	CB // cb
)
