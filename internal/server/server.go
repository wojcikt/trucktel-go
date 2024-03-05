package server

import (
	"bytes"
	"encoding/json"
	"github.com/wojcikt/trucktel-go/internal/shm"
	"log"
	"net/http"
	"time"
)

func New() *http.Server {
	mem, err := shm.Open(shm.DefaultFileName)
	if err != nil {
		log.Fatalf("failed to open shm: %v\n", err)
	}

	handler := newHandler(mem)

	return &http.Server{
		Handler:      handler,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}

func newHandler(mem shm.SharedMemory) http.HandlerFunc {
	var data data

	go func() {
		var values shm.Values
		for {
			err := mem.Read(&values)
			if err != nil {
				log.Printf("failed to read telemetry: %v\n", err)
				continue
			}

			data.update(&values)

			time.Sleep(500 * time.Millisecond)
		}
	}()

	return func(w http.ResponseWriter, r *http.Request) {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			log.Printf("failed to encode to JSON: %v\n", err)
			return
		}
	}
}

type data struct {
	AaAcceleration            vec3f
	Acceleration              vec3f
	AdblueCapacity            float32
	AdblueWarningFactor       float32
	AirPressureEmergency      float32
	AirPressureWarning        float32
	AvAcceleration            vec3f
	BatteryVoltageWarning     float32
	CabinAA                   vec3f
	CabinAV                   vec3f
	CabinOffsetRotation       vec3f
	CabinOffset               vec3f
	CabinPosition             vec3f
	Cargo                     string
	CargoDamage               float32
	CargoId                   string
	CargoMass                 float32
	CityDst                   string
	CityDstId                 string
	CitySrc                   string
	CitySrcId                 string
	CompDst                   string
	CompDstId                 string
	CompSrc                   string
	CompSrcId                 string
	Coordinate                vec3d
	EngineRpmMax              float32
	Ferry                     bool
	FerryPayAmount            int64
	FerrySourceId             string
	FerrySourceName           string
	FerryTargetId             string
	FerryTargetName           string
	FineAmount                int64
	FineOffence               string
	Fined                     bool
	FuelCapacity              float32
	FuelWarningFactor         float32
	Game                      uint32
	GameTelemetryVersionMajor uint32
	GameTelemetryVersionMinor uint32
	GameVersionMajor          uint32
	GameVersionMinor          uint32
	GearDifferential          float32
	GearRatiosForward         [24]float32
	GearRatiosReverse         [8]float32
	Gears                     uint32
	GearsReverse              uint32
	HeadOffsetRotation        vec3f
	HeadOffset                vec3f
	HeadPosition              vec3f
	IsCargoLoaded             bool
	JobCancelled              bool
	JobCancelledPenalty       int64
	JobDelivered              bool
	JobDeliveredAutoloadUsed  bool
	JobDeliveredAutoparkUsed  bool
	JobDeliveredCargoDamage   float32
	JobDeliveredDeliveryTime  uint32
	JobDeliveredDistanceKm    float32
	JobDeliveredEarnedXp      int32
	JobDeliveredRevenue       int64
	JobFinished               bool
	JobFinishedTime           uint32
	JobIncome                 uint64
	JobMarket                 string
	JobStartingTime           uint32
	LvAcceleration            vec3f
	MaxTrailerCount           uint32
	MultiplayerTimeOffset     int64
	OilPressureWarning        float32
	OnJob                     bool
	Paused                    bool
	PlannedDistanceKm         uint32
	Refuel                    bool
	RefuelAmount              float32
	RefuelPayed               bool
	RenderTime                uint64
	RestStop                  int32
	RetarderStepCount         uint32
	Rotation                  vec3d
	Scale                     float32
	SdkActive                 bool
	SelectorCount             uint32
	ShifterType               string
	SimulatedTime             uint64
	SpecialJob                bool
	TelemetryPluginRevision   uint32
	Time                      uint64
	TimeAbs                   uint32
	TimeAbsDelivery           uint32
	Tollgate                  bool
	TollgatePayAmount         int64
	Train                     bool
	TrainPayAmount            int64
	TrainSourceId             string
	TrainSourceName           string
	TrainTargetId             string
	TrainTargetName           string
	Truck                     struct {
		Adblue                   float32
		AdblueWarning            bool
		AirPressure              float32
		AirPressureEmergency     bool
		AirPressureWarning       bool
		BatteryVoltage           float32
		BatteryVoltageWarning    bool
		BlinkerLeft              blinker
		BlinkerRight             blinker
		BrakeTemperature         float32
		Brand                    string
		BrandId                  string
		CruiseControl            bool
		CruiseControlSpeed       float32
		DifferentialLock         bool
		ElectricEnabled          bool
		EngineEnabled            bool
		EngineRpm                float32
		Fuel                     float32
		FuelAvgConsumption       float32
		FuelRange                float32
		FuelWarning              bool
		GameBrake                float32
		GameClutch               float32
		GameSteer                float32
		GameThrottle             float32
		Gear                     int32
		GearDashboard            int32
		HShifterBitmask          [32]uint32
		HShifterPosition         [32]uint32
		HShifterResulting        [32]int32
		HookPosition             vec3f
		Id                       string
		LicensePlate             string
		LicensePlateCountry      string
		LicensePlateCountryId    string
		LiftAxle                 bool
		LiftAxleIndicator        bool
		LightsAuxFront           uint32
		LightsAuxRoof            uint32
		LightsBeacon             bool
		LightsBeamHigh           bool
		LightsBeamLow            bool
		LightsBrake              bool
		LightsDashboard          float32
		LightsHazard             bool
		LightsParking            bool
		LightsReverse            bool
		MotorBrake               bool
		Name                     string
		Odometer                 float32
		OilPressure              float32
		OilPressureWarning       bool
		OilTemperature           float32
		ParkBrake                bool
		RetarderBrake            uint32
		RouteDistance            float32
		RouteTime                float32
		ShifterSlot              uint32
		ShifterToggle            [2]bool
		Speed                    float32
		SpeedLimit               float32
		TrailerLiftAxle          bool
		TrailerLiftAxleIndicator bool
		UserBrake                float32
		UserClutch               float32
		UserSteer                float32
		UserThrottle             float32
		WaterTemperature         float32
		WaterTemperatureWarning  bool
		WearCabin                float32
		WearChassis              float32
		WearEngine               float32
		WearTransmission         float32
		WearWheels               float32
		Wheel                    struct {
			Count                uint32
			Lift                 [16]float32
			LiftOffset           [16]float32
			Liftable             [16]bool
			OnGround             [16]bool
			PositionX            [16]float32
			PositionY            [16]float32
			PositionZ            [16]float32
			Powered              [16]bool
			Radius               [16]float32
			Rotation             [16]float32
			Simulated            [16]bool
			Steerable            [16]bool
			Steering             [16]float32
			Substance            [16]uint32
			SuspensionDeflection [16]float32
			Velocity             [16]float32
		}
		Wipers bool
	}
	UnitCount               uint32
	UnitMass                float32
	WaterTemperatureWarning float32
}

type blinker struct {
	Active bool
	On     bool
}

type vec3f struct {
	X float32
	Y float32
	Z float32
}

type vec3d struct {
	X float64
	Y float64
	Z float64
}

func (d *data) update(v *shm.Values) {
	d.SdkActive = v.SdkActive
	d.Paused = v.Paused
	d.Time = v.Time
	d.SimulatedTime = v.SimulatedTime
	d.RenderTime = v.RenderTime
	d.MultiplayerTimeOffset = v.MultiplayerTimeOffset

	d.TelemetryPluginRevision = v.ScsValues.TelemetryPluginRevision
	d.GameVersionMajor = v.ScsValues.GameVersionMajor
	d.GameVersionMinor = v.ScsValues.GameVersionMinor
	d.Game = v.ScsValues.Game
	d.GameTelemetryVersionMajor = v.ScsValues.GameTelemetryVersionMajor
	d.GameTelemetryVersionMinor = v.ScsValues.GameTelemetryVersionMinor

	d.TimeAbs = v.CommonUI.TimeAbs

	d.Truck.ShifterSlot = v.TruckUI.ShifterSlot
	d.Truck.RetarderBrake = v.TruckUI.RetarderBrake
	d.Truck.LightsAuxFront = v.TruckUI.LightsAuxFront
	d.Truck.LightsAuxRoof = v.TruckUI.LightsAuxRoof
	d.Truck.Wheel.Substance = v.TruckUI.TruckWheelSubstance
	d.Truck.HShifterPosition = v.TruckUI.HShifterPosition
	d.Truck.HShifterBitmask = v.TruckUI.HShifterBitmask

	d.JobDeliveredDeliveryTime = v.GameplayUI.JobDeliveredDeliveryTime
	d.JobStartingTime = v.GameplayUI.JobStartingTime
	d.JobFinishedTime = v.GameplayUI.JobFinishedTime

	d.RestStop = v.CommonI.RestStop

	d.Truck.Gear = v.TruckI.Gear
	d.Truck.GearDashboard = v.TruckI.GearDashboard
	d.Truck.HShifterResulting = v.TruckI.HShifterResulting

	d.JobDeliveredEarnedXp = v.GameplayI.JobDeliveredEarnedXp

	d.Scale = v.CommonF.Scale

	d.FuelCapacity = v.ConfigF.FuelCapacity
	d.FuelWarningFactor = v.ConfigF.FuelWarningFactor
	d.AdblueCapacity = v.ConfigF.AdblueCapacity
	d.AdblueWarningFactor = v.ConfigF.AdblueWarningFactor
	d.AirPressureWarning = v.ConfigF.AirPressureWarning
	d.AirPressureEmergency = v.ConfigF.AirPressureEmergency
	d.OilPressureWarning = v.ConfigF.OilPressureWarning
	d.WaterTemperatureWarning = v.ConfigF.WaterTemperatureWarning
	d.BatteryVoltageWarning = v.ConfigF.BatteryVoltageWarning
	d.EngineRpmMax = v.ConfigF.EngineRpmMax
	d.GearDifferential = v.ConfigF.GearDifferential
	d.CargoMass = v.ConfigF.CargoMass
	d.Truck.Wheel.Radius = v.ConfigF.TruckWheelRadius
	d.GearRatiosForward = v.ConfigF.GearRatiosForward
	d.GearRatiosReverse = v.ConfigF.GearRatiosReverse
	d.UnitMass = v.ConfigF.UnitMass

	d.Truck.Speed = v.TruckF.Speed
	d.Truck.EngineRpm = v.TruckF.EngineRpm
	d.Truck.UserSteer = v.TruckF.UserSteer
	d.Truck.UserThrottle = v.TruckF.UserThrottle
	d.Truck.UserBrake = v.TruckF.UserBrake
	d.Truck.UserClutch = v.TruckF.UserClutch
	d.Truck.GameSteer = v.TruckF.GameSteer
	d.Truck.GameThrottle = v.TruckF.GameThrottle
	d.Truck.GameBrake = v.TruckF.GameBrake
	d.Truck.GameClutch = v.TruckF.GameClutch
	d.Truck.CruiseControlSpeed = v.TruckF.CruiseControlSpeed
	d.Truck.AirPressure = v.TruckF.AirPressure
	d.Truck.BrakeTemperature = v.TruckF.BrakeTemperature
	d.Truck.Fuel = v.TruckF.Fuel
	d.Truck.FuelAvgConsumption = v.TruckF.FuelAvgConsumption
	d.Truck.FuelRange = v.TruckF.FuelRange
	d.Truck.Adblue = v.TruckF.Adblue
	d.Truck.OilPressure = v.TruckF.OilPressure
	d.Truck.OilTemperature = v.TruckF.OilTemperature
	d.Truck.WaterTemperature = v.TruckF.WaterTemperature
	d.Truck.BatteryVoltage = v.TruckF.BatteryVoltage
	d.Truck.LightsDashboard = v.TruckF.LightsDashboard
	d.Truck.WearEngine = v.TruckF.WearEngine
	d.Truck.WearTransmission = v.TruckF.WearTransmission
	d.Truck.WearCabin = v.TruckF.WearCabin
	d.Truck.WearChassis = v.TruckF.WearChassis
	d.Truck.WearWheels = v.TruckF.WearWheels
	d.Truck.Odometer = v.TruckF.TruckOdometer
	d.Truck.RouteDistance = v.TruckF.RouteDistance
	d.Truck.RouteTime = v.TruckF.RouteTime
	d.Truck.SpeedLimit = v.TruckF.SpeedLimit
	d.Truck.Wheel.SuspensionDeflection = v.TruckF.TruckWheelSuspensionDeflection
	d.Truck.Wheel.Velocity = v.TruckF.TruckWheelVelocity
	d.Truck.Wheel.Steering = v.TruckF.TruckWheelSteering
	d.Truck.Wheel.Rotation = v.TruckF.TruckWheelRotation
	d.Truck.Wheel.Lift = v.TruckF.TruckWheelLift
	d.Truck.Wheel.LiftOffset = v.TruckF.TruckWheelLiftOffset

	d.JobDeliveredCargoDamage = v.GameplayF.JobDeliveredCargoDamage
	d.JobDeliveredDistanceKm = v.GameplayF.JobDeliveredDistanceKm
	d.RefuelAmount = v.GameplayF.RefuelAmount

	d.CargoDamage = v.JobF.CargoDamage

	d.Truck.Wheel.Steerable = v.ConfigB.TruckWheelSteerable
	d.Truck.Wheel.Simulated = v.ConfigB.TruckWheelSimulated
	d.Truck.Wheel.Powered = v.ConfigB.TruckWheelPowered
	d.Truck.Wheel.Liftable = v.ConfigB.TruckWheelLiftable
	d.IsCargoLoaded = v.ConfigB.IsCargoLoaded
	d.SpecialJob = v.ConfigB.SpecialJob

	d.Truck.ParkBrake = v.TruckB.ParkBrake
	d.Truck.MotorBrake = v.TruckB.MotorBrake
	d.Truck.AirPressureWarning = v.TruckB.AirPressureWarning
	d.Truck.AirPressureEmergency = v.TruckB.AirPressureEmergency
	d.Truck.FuelWarning = v.TruckB.FuelWarning
	d.Truck.AdblueWarning = v.TruckB.AdblueWarning
	d.Truck.OilPressureWarning = v.TruckB.OilPressureWarning
	d.Truck.WaterTemperatureWarning = v.TruckB.WaterTemperatureWarning
	d.Truck.BatteryVoltageWarning = v.TruckB.BatteryVoltageWarning
	d.Truck.ElectricEnabled = v.TruckB.ElectricEnabled
	d.Truck.EngineEnabled = v.TruckB.EngineEnabled
	d.Truck.Wipers = v.TruckB.Wipers
	d.Truck.BlinkerLeft.Active = v.TruckB.BlinkerLeftActive
	d.Truck.BlinkerLeft.On = v.TruckB.BlinkerLeftOn
	d.Truck.BlinkerRight.Active = v.TruckB.BlinkerRightActive
	d.Truck.BlinkerRight.On = v.TruckB.BlinkerRightOn
	d.Truck.LightsParking = v.TruckB.LightsParking
	d.Truck.LightsBeamLow = v.TruckB.LightsBeamLow
	d.Truck.LightsBeamHigh = v.TruckB.LightsBeamHigh
	d.Truck.LightsBeacon = v.TruckB.LightsBeacon
	d.Truck.LightsBrake = v.TruckB.LightsBrake
	d.Truck.LightsReverse = v.TruckB.LightsReverse
	d.Truck.LightsHazard = v.TruckB.LightsHazard
	d.Truck.CruiseControl = v.TruckB.CruiseControl
	d.Truck.Wheel.OnGround = v.TruckB.TruckWheelOnGround
	d.Truck.ShifterToggle = v.TruckB.ShifterToggle
	d.Truck.DifferentialLock = v.TruckB.DifferentialLock
	d.Truck.LiftAxle = v.TruckB.LiftAxle
	d.Truck.LiftAxleIndicator = v.TruckB.LiftAxleIndicator
	d.Truck.TrailerLiftAxle = v.TruckB.TrailerLiftAxle
	d.Truck.TrailerLiftAxleIndicator = v.TruckB.TrailerLiftAxleIndicator

	d.JobDeliveredAutoparkUsed = v.GameplayB.JobDeliveredAutoparkUsed
	d.JobDeliveredAutoloadUsed = v.GameplayB.JobDeliveredAutoloadUsed

	d.Truck.BrandId = toString(v.ConfigS.TruckBrandId[:])
	d.Truck.Brand = toString(v.ConfigS.TruckBrand[:])
	d.Truck.Id = toString(v.ConfigS.TruckId[:])
	d.Truck.Name = toString(v.ConfigS.TruckName[:])
	d.CargoId = toString(v.ConfigS.CargoId[:])
	d.Cargo = toString(v.ConfigS.Cargo[:])
	d.CityDstId = toString(v.ConfigS.CityDstId[:])
	d.CityDst = toString(v.ConfigS.CityDst[:])
	d.CompDstId = toString(v.ConfigS.CompDstId[:])
	d.CompDst = toString(v.ConfigS.CompDst[:])
	d.CitySrcId = toString(v.ConfigS.CitySrcId[:])
	d.CitySrc = toString(v.ConfigS.CitySrc[:])
	d.CompSrcId = toString(v.ConfigS.CompSrcId[:])
	d.CompSrc = toString(v.ConfigS.CompSrc[:])
	d.ShifterType = toString(v.ConfigS.ShifterType[:])
	d.Truck.LicensePlate = toString(v.ConfigS.TruckLicensePlate[:])
	d.Truck.LicensePlateCountryId = toString(v.ConfigS.TruckLicensePlateCountryId[:])
	d.Truck.LicensePlateCountry = toString(v.ConfigS.TruckLicensePlateCountry[:])
	d.JobMarket = toString(v.ConfigS.JobMarket[:])

	d.FineOffence = toString(v.GameplayS.FineOffence[:])
	d.FerrySourceName = toString(v.GameplayS.FerrySourceName[:])
	d.FerryTargetName = toString(v.GameplayS.FerryTargetName[:])
	d.FerrySourceId = toString(v.GameplayS.FerrySourceId[:])
	d.FerryTargetId = toString(v.GameplayS.FerryTargetId[:])
	d.TrainSourceName = toString(v.GameplayS.TrainSourceName[:])
	d.TrainTargetName = toString(v.GameplayS.TrainTargetName[:])
	d.TrainSourceId = toString(v.GameplayS.TrainSourceId[:])
	d.TrainTargetId = toString(v.GameplayS.TrainTargetId[:])

	d.JobIncome = v.ConfigULL.JobIncome

	d.JobCancelledPenalty = v.GameplayLL.JobCancelledPenalty
	d.JobDeliveredRevenue = v.GameplayLL.JobDeliveredRevenue
	d.FineAmount = v.GameplayLL.FineAmount
	d.TollgatePayAmount = v.GameplayLL.TollgatePayAmount
	d.FerryPayAmount = v.GameplayLL.FerryPayAmount
	d.TrainPayAmount = v.GameplayLL.TrainPayAmount

	d.OnJob = v.SpecialB.OnJob
	d.JobFinished = v.SpecialB.JobFinished
	d.JobCancelled = v.SpecialB.JobCancelled
	d.JobDelivered = v.SpecialB.JobDelivered
	d.Fined = v.SpecialB.Fined
	d.Tollgate = v.SpecialB.Tollgate
	d.Ferry = v.SpecialB.Ferry
	d.Train = v.SpecialB.Train
	d.Refuel = v.SpecialB.Refuel
	d.RefuelPayed = v.SpecialB.RefuelPayed
}

func toString(b []byte) string {
	length := bytes.IndexByte(b, 0)
	return string(b[:length])
}
