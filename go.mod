module github.com/met-slewis/ShapeServiceGo

go 1.19

replace (
	github.com/MetServiceDev/WeatherEventLib => ../WeatherEventLib
	github.com/met-slewis/WeatherUMS => ../WeatherUMS
)

require (
	github.com/MetServiceDev/WeatherEventLib v0.0.0-00010101000000-000000000000
	github.com/paulmach/orb v0.8.0
)

require (
	github.com/sirupsen/logrus v1.9.0 // indirect
	github.com/sony/sonyflake v1.1.0 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
)
