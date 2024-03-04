package shm

import (
	"encoding/binary"
	"io"
)

const (
	stringSize    = 64
	substanceSize = 25
)

func (v *Values) Read(r io.Reader) error {
	return binary.Read(r, binary.LittleEndian, v)
}

type Values struct {
	//----- START OF FIRST ZONE AT OFFSET 0 -----//

	SdkActive             bool
	_                     [3]byte
	Paused                bool
	_                     [3]byte
	Time                  uint64
	SimulatedTime         uint64
	RenderTime            uint64
	MultiplayerTimeOffset int64

	//----- END OF FIRST ZONE AT OFFSET 39 -----//
	//----- START OF SECOND ZONE AT OFFSET 40 -----//

	ScsValues struct {
		TelemetryPluginRevision   uint32
		GameVersionMajor          uint32
		GameVersionMinor          uint32
		Game                      uint32
		GameTelemetryVersionMajor uint32
		GameTelemetryVersionMinor uint32
	}
	CommonUI struct {
		TimeAbs uint32
	}
	ConfigUI struct {
		Gears             uint32
		GearsReverse      uint32
		RetarderStepCount uint32
		TruckWheelCount   uint32
		SelectorCount     uint32
		TimeAbsDelivery   uint32
		MaxTrailerCount   uint32
		UnitCount         uint32
		PlannedDistanceKm uint32
	}
	TruckUI struct {
		ShifterSlot         uint32
		RetarderBrake       uint32
		LightsAuxFront      uint32
		LightsAuxRoof       uint32
		TruckWheelSubstance [16]uint32
		HShifterPosition    [32]uint32
		HShifterBitmask     [32]uint32
	}
	GameplayUI struct {
		JobDeliveredDeliveryTime uint32
		JobStartingTime          uint32
		JobFinishedTime          uint32
	}
	_ [48]byte

	//----- END OF SECOND ZONE AT OFFSET 499 -----//
	//----- START OF Third ZONE AT OFFSET 500 -----//

	CommonI struct {
		RestStop int32
	}
	TruckI struct {
		Gear              int32
		GearDashboard     int32
		HShifterResulting [32]int32
	}
	GameplayI struct {
		JobDeliveredEarnedXp int32
	}
	_ [56]byte

	//----- END OF third ZONE AT OFFSET 699 -----//
	//----- START OF FOURTH ZONE AT OFFSET 700 -----//

	CommonF struct {
		Scale float32
	}
	ConfigF struct {
		FuelCapacity            float32
		FuelWarningFactor       float32
		AdblueCapacity          float32
		AdblueWarningFactor     float32
		AirPressureWarning      float32
		AirPressureEmergency    float32
		OilPressureWarning      float32
		WaterTemperatureWarning float32
		BatteryVoltageWarning   float32
		EngineRpmMax            float32
		GearDifferential        float32
		CargoMass               float32
		TruckWheelRadius        [16]float32
		GearRatiosForward       [24]float32
		GearRatiosReverse       [8]float32
		UnitMass                float32
	}
	TruckF struct {
		Speed                    float32
		EngineRpm                float32
		UserSteer                float32
		UserThrottle             float32
		UserBrake                float32
		UserClutch               float32
		GameSteer                float32
		GameThrottle             float32
		GameBrake                float32
		GameClutch               float32
		CruiseControlSpeed       float32
		AirPressure              float32
		BrakeTemperature         float32
		Fuel                     float32
		FuelAvgConsumption       float32
		FuelRange                float32
		Adblue                   float32
		OilPressure              float32
		OilTemperature           float32
		WaterTemperature         float32
		BatteryVoltage           float32
		LightsDashboard          float32
		WearEngine               float32
		WearTransmission         float32
		WearCabin                float32
		WearChassis              float32
		WearWheels               float32
		TruckOdometer            float32
		RouteDistance            float32
		RouteTime                float32
		SpeedLimit               float32
		TruckWheelSuspDeflection [16]float32
		TruckWheelVelocity       [16]float32
		TruckWheelSteering       [16]float32
		TruckWheelRotation       [16]float32
		TruckWheelLift           [16]float32
		TruckWheelLiftOffset     [16]float32
	}
	GameplayF struct {
		JobDeliveredCargoDamage float32
		JobDeliveredDistanceKm  float32
		RefuelAmount            float32
	}
	JobF struct {
		CargoDamage float32
	}
	_ [28]byte

	//----- END OF FOURTH ZONE AT OFFSET 1499 -----//
	//----- START OF FIFTH ZONE AT OFFSET 1500 -----//

	ConfigB struct {
		TruckWheelSteerable [16]bool
		TruckWheelSimulated [16]bool
		TruckWheelPowered   [16]bool
		TruckWheelLiftable  [16]bool
		IsCargoLoaded       bool
		SpecialJob          bool
	}
	TruckB struct {
		ParkBrake                bool
		MotorBrake               bool
		AirPressureWarning       bool
		AirPressureEmergency     bool
		FuelWarning              bool
		AdblueWarning            bool
		OilPressureWarning       bool
		WaterTemperatureWarning  bool
		BatteryVoltageWarning    bool
		ElectricEnabled          bool
		EngineEnabled            bool
		Wipers                   bool
		BlinkerLeftActive        bool
		BlinkerRightActive       bool
		BlinkerLeftOn            bool
		BlinkerRightOn           bool
		LightsParking            bool
		LightsBeamLow            bool
		LightsBeamHigh           bool
		LightsBeacon             bool
		LightsBrake              bool
		LightsReverse            bool
		LightsHazard             bool
		CruiseControl            bool
		TruckWheelOnGround       [16]bool
		ShifterToggle            [2]bool
		DifferentialLock         bool
		LiftAxle                 bool
		LiftAxleIndicator        bool
		TrailerLiftAxle          bool
		TrailerLiftAxleIndicator bool
	}
	GameplayB struct {
		JobDeliveredAutoparkUsed bool
		JobDeliveredAutoloadUsed bool
	}
	_ [25]byte

	//----- END OF FIFTH ZONE AT OFFSET 1639 -----//
	//----- START OF SIXTH ZONE AT OFFSET 1640 -----//

	ConfigFV struct {
		CabinPositionX      float32
		CabinPositionY      float32
		CabinPositionZ      float32
		HeadPositionX       float32
		HeadPositionY       float32
		HeadPositionZ       float32
		TruckHookPositionX  float32
		TruckHookPositionY  float32
		TruckHookPositionZ  float32
		TruckWheelPositionX [16]float32
		TruckWheelPositionY [16]float32
		TruckWheelPositionZ [16]float32
	}
	TruckFV struct {
		LvAccelerationX float32
		LvAccelerationY float32
		LvAccelerationZ float32
		AvAccelerationX float32
		AvAccelerationY float32
		AvAccelerationZ float32
		AccelerationX   float32
		AccelerationY   float32
		AccelerationZ   float32
		AaAccelerationX float32
		AaAccelerationY float32
		AaAccelerationZ float32
		CabinAVX        float32
		CabinAVY        float32
		CabinAVZ        float32
		CabinAAX        float32
		CabinAAY        float32
		CabinAAZ        float32
	}
	_ [60]byte

	//----- END OF SIXTH ZONE AT OFFSET 1999 -----//
	//----- START OF 7TH ZONE AT OFFSET 2000 -----//

	TruckFP struct {
		CabinOffsetX         float32
		CabinOffsetY         float32
		CabinOffsetZ         float32
		CabinOffsetRotationX float32
		CabinOffsetRotationY float32
		CabinOffsetRotationZ float32
		HeadOffsetX          float32
		HeadOffsetY          float32
		HeadOffsetZ          float32
		HeadOffsetRotationX  float32
		HeadOffsetRotationY  float32
		HeadOffsetRotationZ  float32
	}
	_ [152]byte

	//----- END OF 7TH ZONE AT OFFSET 2199 -----//
	//----- START OF 8TH ZONE AT OFFSET 2200 -----//

	TruckDP struct {
		CoordinateX float64
		CoordinateY float64
		CoordinateZ float64
		RotationX   float64
		RotationY   float64
		RotationZ   float64
	}
	_ [52]byte

	//----- END OF 8TH ZONE AT OFFSET 2299 -----//
	//----- START OF 9TH ZONE AT OFFSET 2300 -----//

	ConfigS struct {
		TruckBrandId               [stringSize]byte
		TruckBrand                 [stringSize]byte
		TruckId                    [stringSize]byte
		TruckName                  [stringSize]byte
		CargoId                    [stringSize]byte
		Cargo                      [stringSize]byte
		CityDstId                  [stringSize]byte
		CityDst                    [stringSize]byte
		CompDstId                  [stringSize]byte
		CompDst                    [stringSize]byte
		CitySrcId                  [stringSize]byte
		CitySrc                    [stringSize]byte
		CompSrcId                  [stringSize]byte
		CompSrc                    [stringSize]byte
		ShifterType                [16]byte
		TruckLicensePlate          [stringSize]byte
		TruckLicensePlateCountryId [stringSize]byte
		TruckLicensePlateCountry   [stringSize]byte
		JobMarket                  [32]byte
	}
	GameplayS struct {
		FineOffence     [32]byte
		FerrySourceName [stringSize]byte
		FerryTargetName [stringSize]byte
		FerrySourceId   [stringSize]byte
		FerryTargetId   [stringSize]byte
		TrainSourceName [stringSize]byte
		TrainTargetName [stringSize]byte
		TrainSourceId   [stringSize]byte
		TrainTargetId   [stringSize]byte
	}
	_ [20]byte

	//----- END OF 9TH ZONE AT OFFSET 3999 -----//
	//----- START OF 10TH ZONE AT OFFSET 4000 -----//

	ConfigULL struct {
		JobIncome uint64
	}
	_ [192]byte

	//----- END OF 10TH ZONE AT OFFSET 4199 -----//
	//----- START OF 11TH ZONE AT OFFSET 4200 -----//

	GameplayLL struct {
		JobCancelledPenalty int64
		JobDeliveredRevenue int64
		FineAmount          int64
		TollgatePayAmount   int64
		FerryPayAmount      int64
		TrainPayAmount      int64
	}
	_ [52]byte

	//----- END OF 11TH ZONE AT OFFSET 4299 -----//
	//----- START OF 12TH ZONE AT OFFSET 4300 -----//

	SpecialB struct {
		OnJob        bool
		JobFinished  bool
		JobCancelled bool
		JobDelivered bool
		Fined        bool
		Tollgate     bool
		Ferry        bool
		Train        bool
		Refuel       bool
		RefuelPayed  bool
	}
	_ [90]byte

	//----- END OF 12TH ZONE AT OFFSET 4399 -----//
	//----- START OF 13TH ZONE AT OFFSET 4400 -----//

	Substances struct {
		Substance [substanceSize][stringSize]byte
	}
	//----- END OF 13TH ZONE AT OFFSET 5999 -----//
	//----- START OF 14TH ZONE AT OFFSET 6000 -----//

	Trailer struct {
		Trailer [10]ScsTrailer
	}

	//----- END OF 14TH ZONE AT OFFSET 21619 -----//
}

type ScsTrailer struct {

	//----- START OF FIRST ZONE AT OFFSET 0 -----//

	ConB struct {
		WheelSteerable [16]bool
		WheelSimulated [16]bool
		WheelPowered   [16]bool
		WheelLiftable  [16]bool
	}
	ComB struct {
		WheelOnGround [16]bool
		Attached      bool
	}
	_ [3]byte

	//----- END OF FIRST ZONE AT OFFSET 83 -----//
	//----- START OF SECOND ZONE AT OFFSET 84 -----//

	ComUI struct {
		WheelSubstance [16]uint32
	}
	ConUI struct {
		WheelCount uint32
	}

	//----- END OF SECOND ZONE AT OFFSET 151 -----//
	//----- START OF THIRD ZONE AT OFFSET 152 -----//

	ComF struct {
		CargoDamage               float32
		WearChassis               float32
		WearWheels                float32
		WearBody                  float32
		WheelSuspensionDeflection [16]float32
		WheelVelocity             [16]float32
		WheelSteering             [16]float32
		WheelRotation             [16]float32
		WheelLift                 [16]float32
		WheelLiftOffset           [16]float32
	}
	ConF struct {
		WheelRadius [16]float32
	}

	//----- END OF THIRD ZONE AT OFFSET 615 -----//
	//----- START OF 4TH ZONE AT OFFSET 616 -----//

	ComFV struct {
		LinearVelocityX      float32
		LinearVelocityY      float32
		LinearVelocityZ      float32
		AngularVelocityX     float32
		AngularVelocityY     float32
		AngularVelocityZ     float32
		LinearAccelerationX  float32
		LinearAccelerationY  float32
		LinearAccelerationZ  float32
		AngularAccelerationX float32
		AngularAccelerationY float32
		AngularAccelerationZ float32
	}
	ConFV struct {
		HookPositionX  float32
		HookPositionY  float32
		HookPositionZ  float32
		WheelPositionX [16]float32
		WheelPositionY [16]float32
		WheelPositionZ [16]float32
	}
	_ [4]byte

	//----- END OF 4TH ZONE AT OFFSET 871 -----//
	//----- START OF 5TH ZONE AT OFFSET 872 -----//

	ComDP struct {
		WorldX    float64
		WorldY    float64
		WorldZ    float64
		RotationX float64
		RotationY float64
		RotationZ float64
	}

	//----- END OF 5TH ZONE AT OFFSET 919 -----//
	//----- START OF 6TH ZONE AT OFFSET 920 -----//

	ConS struct {
		Id                    [stringSize]byte
		CargoAccessoryId      [stringSize]byte
		BodyType              [stringSize]byte
		BrandId               [stringSize]byte
		Brand                 [stringSize]byte
		Name                  [stringSize]byte
		ChainType             [stringSize]byte
		LicensePlate          [stringSize]byte
		LicensePlateCountry   [stringSize]byte
		LicensePlateCountryId [stringSize]byte
	}

	//----- END OF 6TH ZONE AT OFFSET 1559 -----//
}
