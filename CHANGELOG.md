Changelog
=========
All notable changes to this project will be documented in this file.

v1.1.5-alpha
-------------

### Changed

- Fixed Trade model unmarshal


v1.1.4-alpha
-------------

### Changed

- Added buffer to `Ws` sender channel

v1.1.3-alpha
-------------

### Changed

- Fixed `Ws` lock issue

v1.1.2-alpha
-------------

### Changed

- Combined mutexes into one for `Ws` client

v1.1.1-alpha
-------------

### Changed

- Fixed all float64 that were mapped to int64
- Changed destination constants values to start from iota

v1.1.0-alpha
-------------

### Changed

- Because versioning needs all the numbers!

v1.1-alpha
-------------

### Changed

- Because Golang checksum tool sometimes can not verify the version :)

v1.0.35-alpha
-------------

### Changed

- Fixed response struct unmarshalling issue

v1.0.34-alpha
-------------

### Changed

- Fixed response struct unmarshalling issue

v1.0.33-alpha
-------------

### Changed

- Added WS trade requests

v1.0.32-alpha
-------------

### Changed

- Fixed `OrderBookWs` struct `Checksum` field by removing string json tag
- Added `InstID` to `OrderBook` event struct

v1.0.31-alpha
-------------

### Changed

- Migrated to `okx.com`

v1.0.30-alpha
-------------

### Changed

- Fixed `PlaceOrder` struct Sz field by changing it to float64
- Added `UnmarshalJSON` to `Basic` response struct

v1.0.29-alpha
-------------

### Changed

- Fixed websocket's client sender racing issue
- Fixed websocket's public `UOrderBook` racing issue
- Added `Action` field to  `OrderBook` struct
- Fixed `GetOrderDetail` function parameter

v1.0.28-alpha
-------------

### Changed

- Fixed nil pointer reference bug on websocket connection

v1.0.27-alpha
-------------

### Changed

- Fixed typo on `GetSystemTime` func
- Added `string` tag to `GetCandlesticks`

v1.0.26-alpha
------------

### Changed

- Removed extra login log entry

v1.0.25-alpha
------------

### Changed

- Fixed login process

v1.0.24-alpha
------------

### Changed

- Fixed nil assignment issue on login process

v1.0.23-alpha
------------

### Changed

- Changed login mechanism

v1.0.22-alpha
------------

### Changed

- Added place algo order
- Added cancel algo order
- Added get algo order list and history

v1.0.21-alpha
------------

### Changed

- Fixed public data models json tags

v1.0.20-alpha
------------

### Changed

- Fixed account models json tags

v1.0.19-alpha
------------

### Changed

- Fixed trade models json tags

v1.0.18-alpha
------------

### Changed

- Renamed subaccount models filed names

v1.0.17-alpha
------------

### Changed

- Fixed api requests json tag issues

v1.0.16-alpha
------------

### Changed

- Fixed rest api request json tag issues

v1.0.15-alpha
------------

### Changed

- Added status endpoint on rest

v1.0.14-alpha
------------

### Changed

- Fixed requests array type conversion

v1.0.13-alpha
------------

### Changed

- Fixed requests array type conversion

v1.0.12-alpha
------------

### Changed

- Fixed invalid channel name on `private.Order`

v1.0.11-alpha
------------

### Changed

- Better error handling on `Ws`

v1.0.10-alpha
------------

### Changed

- Fixed nil assignment on Ws client of Trade
- Added request `ID` to trade requests
- Added new event `Success` to handle trade success response

v1.0.9-alpha
------------

### Changed

- Revert json tag changes on last version
- Fixed `account_models/Config.PosMode` type

v1.0.8-alpha
------------

### Changed

- Fixed typo issue on `funding_requests.go`
- Fixed json tag on `account_requests.go`

v1.0.7-alpha
------------

### Changed

- Fixed JSON tags on `models`
- Changed `*Id` into `ID` everywhere

v1.0.6-alpha
------------

### Changed

- Fixed GET method requests issue

v1.0.5-alpha
------------

### Changed

- Forced destination on creating api client

v1.0.4-alpha
------------

### Changed

- Added github actions badges to README.md

v1.0.3-alpha
------------

### Changed

- Fixed nil assignment on api.go 
