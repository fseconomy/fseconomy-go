package data

import (
	"encoding/xml"
	"github.com/fseconomy/fseconomy-go/internal/data"
	"github.com/fseconomy/fseconomy-go/types"
	"strconv"
)

type Facility struct {
	Icao           string        `xml:"Icao"`
	Location       string        `xml:"Location"`
	Carrier        string        `xml:"Carrier"`
	CommodityNames string        `xml:"CommodityNames"`
	Gates          int           `xml:"Gates"`
	JobsPublic     types.FseBool `xml:"JobsPublic"`
	Destinations   string        `xml:"Destinations"`
	Fbo            string        `xml:"Fbo"`
	Status         string        `xml:"Status"`
}

type Fbo struct {
	FboId             int           `xml:"FboId"`
	Status            string        `xml:"Status"`
	Airport           string        `xml:"Airport"`
	Name              string        `xml:"Name"`
	Owner             string        `xml:"Owner"`
	Icao              string        `xml:"Icao"`
	Location          string        `xml:"Location"`
	Lots              int           `xml:"Lots"`
	RepairShop        types.FseBool `xml:"RepairShop"`
	Gates             int           `xml:"Gates"`
	GatesRented       int           `xml:"GatesRented"`
	Fuel100LL         int           `xml:"Fuel100LL"`
	FuelJetA          int           `xml:"FuelJetA"`
	BuildingMaterials int           `xml:"BuildingMaterials"`
	Supplies          int           `xml:"Supplies"`
	SuppliesPerDay    int           `xml:"SuppliesPerDay"`
	SuppliedDays      int           `xml:"SuppliedDays"`
	SellPrice         float64       `xml:"SellPrice"`
	Fuel100LLGal      int           `xml:"Fuel100LLGal"`
	FuelJetAGal       int           `xml:"FuelJetAGal"`
	Price100LLGal     float64       `xml:"Price100LLGal"`
	PriceJetAGal      float64       `xml:"PriceJetAGal"`
}

type FboMonthlySummary struct {
	Owner                             string  `xml:"Owner"`
	Icao                              string  `xml:"ICAO"`
	Month                             int     `xml:"Month"`
	Year                              int     `xml:"Year"`
	AssignmentRentalExpenses          float64 `xml:"Assignment_Rental_Expenses"`
	AssignmentIncome                  float64 `xml:"Assignment_Income"`
	AssignmentExpenses                float64 `xml:"Assignment_Expenses"`
	AssignmentPilotFees               float64 `xml:"Assignment_Pilot_Fees"`
	AssignmentAdditionalCrewFees      float64 `xml:"Assignment_Additional_Crew_Fees"`
	AssignmentGroundCrewFees          float64 `xml:"Assignment_Ground_Crew_Fees"`
	AssignmentBookingFees             float64 `xml:"Assignment_Booking_Fees"`
	AircraftOpsRentalIncome           float64 `xml:"Aircraft_Ops_Rental_Income"`
	AircraftOpsRefueling100LL         float64 `xml:"Aircraft_Ops_Refueling_100_LL"`
	AircraftOpsRefuelingJetA          float64 `xml:"Aircraft_Ops_Refueling_JetA"`
	AircraftOpsLandingFees            float64 `xml:"Aircraft_Ops_Landing_Fees"`
	AircraftOpsExpensesForMaintenance float64 `xml:"Aircraft_Ops_Expenses_for_Maintenance"`
	AircraftOpsEquipmentInstallation  float64 `xml:"Aircraft_Ops_Equipment_Installation"`
	AircraftSold                      float64 `xml:"Aircraft_Sold"`
	AircraftBought                    float64 `xml:"Aircraft_Bought"`
	FboOpsRefueling100LL              float64 `xml:"Fbo_Ops_Refueling_100_LL"`
	FboOpsRefuelingJetA               float64 `xml:"Fbo_Ops_Refueling_JetA"`
	FboOpsGroundCrewFees              float64 `xml:"Fbo_Ops_Ground_Crew_Fees"`
	FboOpsRepairShopIncome            float64 `xml:"Fbo_Ops_Repair_Shop_Income"`
	FboOpsRepairShopExpenses          float64 `xml:"Fbo_Ops_Repair_Shop_Expenses"`
	FboOpsEquipmentInstallation       float64 `xml:"Fbo_Ops_Equipment_Installation"`
	FboOpsEquipmentExpenses           float64 `xml:"Fbo_Ops_Equipment_Expenses"`
	PtRentIncome                      float64 `xml:"PT_Rent_Income"`
	PtRentExpenses                    float64 `xml:"PT_Rent_Expenses"`
	FboSold                           float64 `xml:"FBO_Sold"`
	FboBought                         float64 `xml:"FBO_Bought"`
	GoodsBoughtWholesale100LL         float64 `xml:"Goods_Bought_Wholesale_100LL"`
	GoodsBoughtWholesaleJetA          float64 `xml:"Goods_Bought_Wholesale_JetA"`
	GoodsBoughtBuildingMaterials      float64 `xml:"Goods_Bought_Building_Materials"`
	GoodsBoughtSupplies               float64 `xml:"Goods_Bought_Supplies"`
	GoodsSoldWholesale100LL           float64 `xml:"Goods_Sold_Wholesale_100LL"`
	GoodsSoldWholesaleJetA            float64 `xml:"Goods_Sold_Wholesale_JetA"`
	GoodsSoldBuildingMaterials        float64 `xml:"Goods_Sold_Building_Materials"`
	GoodsSoldSupplies                 float64 `xml:"Goods_Sold_Supplies"`
	GroupPayments                     float64 `xml:"Group_Payments"`
	GroupDeletion                     float64 `xml:"Group_Deletion"`
	NetTotal                          float64 `xml:"Net_Total"`
	CurrentOps                        int     `xml:"Current_Ops"`
	AvgOps                            int     `xml:"Avg_Ops"`
}

// FacilitiesByKey extracts data from the Facilities By Key data feed
func (d *Data) FacilitiesByKey() ([]*Facility, error) {
	keys, err := d.Keys()
	if err != nil {
		return nil, err
	}
	feed, err := data.GetDataFeed("Facilities By Key")
	if err != nil {
		return nil, err
	}
	resp, err := feed.QueryFeed(nil, keys)
	if err != nil {
		return nil, err
	}
	var items struct {
		Facility []*Facility `xml:"Facility"`
	}
	err = xml.Unmarshal(resp.ByteData, &items)
	if err != nil {
		return nil, err
	}
	return items.Facility, nil
}

// FbosByKey extracts data from the Facilities By Key data feed
func (d *Data) FbosByKey() ([]*Fbo, error) {
	keys, err := d.Keys()
	if err != nil {
		return nil, err
	}
	feed, err := data.GetDataFeed("FBOs By Key")
	if err != nil {
		return nil, err
	}
	resp, err := feed.QueryFeed(nil, keys)
	if err != nil {
		return nil, err
	}
	var items struct {
		Fbo []*Fbo `xml:"FBO"`
	}
	err = xml.Unmarshal(resp.ByteData, &items)
	if err != nil {
		return nil, err
	}
	return items.Fbo, nil
}

// FbosForSale extracts data from the FBOs By Key data feed
func (d *Data) FbosForSale() ([]*Fbo, error) {
	keys, err := d.Keys()
	if err != nil {
		return nil, err
	}
	feed, err := data.GetDataFeed("FBOs For Sale")
	if err != nil {
		return nil, err
	}
	resp, err := feed.QueryFeed(nil, keys)
	if err != nil {
		return nil, err
	}
	var items struct {
		Fbo []*Fbo `xml:"FBO"`
	}
	err = xml.Unmarshal(resp.ByteData, &items)
	if err != nil {
		return nil, err
	}
	return items.Fbo, nil
}

// FboMonthlySummaryByIcao extracts data from the FBO Monthly Summary by ICAO data feed
func (d *Data) FboMonthlySummaryByIcao(month int, year int, icao string) ([]*FboMonthlySummary, error) {
	keys, err := d.Keys()
	if err != nil {
		return nil, err
	}
	feed, err := data.GetDataFeed("FBO Monthly Summary By ICAO")
	if err != nil {
		return nil, err
	}
	resp, err := feed.QueryFeed(map[string]string{"month": strconv.Itoa(month), "year": strconv.Itoa(year), "icao": icao}, keys)
	if err != nil {
		return nil, err
	}
	var items struct {
		FboMonthlySummary []*FboMonthlySummary `xml:"FboMonthlySummary"`
	}
	err = xml.Unmarshal(resp.ByteData, &items)
	if err != nil {
		return nil, err
	}
	return items.FboMonthlySummary, nil
}
