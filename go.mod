module example

//replace FishCache => ./
replace lru => ./FishCache/lru

replace consistenthash => ./FishCache/consistenthash

replace FishCache => ./FishCache

replace singleflight => ./FishCache/singleflight

replace fishcachepb => ./FishCache/fishcachepb

go 1.14

require (
	FishCache v0.0.0-00010101000000-000000000000
	consistenthash v0.0.0-00010101000000-000000000000 // indirect
	fishcachepb v0.0.0-00010101000000-000000000000 // indirect
	github.com/golang/protobuf v1.4.2 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	singleflight v0.0.0-00010101000000-000000000000 // indirect
)
