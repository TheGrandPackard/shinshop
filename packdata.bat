cd webserver

go-bindata-assetfs -ignore=\\.DS_Store -pkg template templates/...
move bindata_assetfs.go template/templatedata.go

go-bindata-assetfs -ignore=\\.DS_Store -pkg webserver web/...
move bindata_assetfs.go webdata.go

go-bindata-assetfs -ignore=\\.DS_Store -pkg rest rest/map/...
move bindata_assetfs.go rest/mapdata.go

cd ..
go run main.go
