module github.com/met-slewis/ShapeServiceGo

go 1.19

replace (
	github.com/MetServiceDev/WeatherEventLib => ../weatherEvents/WeatherEventLib
	github.com/met-slewis/WeatherUMS => ../weatherEvents/WeatherUMS
)

require (
	github.com/MetServiceDev/WeatherEventLib v0.0.0-00010101000000-000000000000
	github.com/paulmach/orb v0.8.0
)

require (
	github.com/ctessum/geom v0.2.12 // indirect
	github.com/ctessum/polyclip-go v1.1.0 // indirect
	github.com/gonum/floats v0.0.0-20181209220543-c233463c7e82 // indirect
	github.com/gonum/internal v0.0.0-20181124074243-f884aa714029 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	github.com/sony/sonyflake v1.1.0 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
	gonum.org/v1/gonum v0.9.3 // indirect
)
