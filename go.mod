module example

//replace FishCache => ./
replace lru => ./FishCache/lru

replace consistenthash => ./FishCache/consistenthash

replace FishCache => ./FishCache

go 1.14

require (
	FishCache v0.0.0-00010101000000-000000000000
	consistenthash v0.0.0-00010101000000-000000000000 // indirect
)
