# Glossary

| Term | Definition |
| :--- | :--------- |
| **ABS** | Anti-lock Braking System. Referenced in tell tales like `ABS_TRAILER`. |
| **ACC** | Adaptive Cruise Control. Referenced in tell tales like `ACC`. |
| **ACD** | Automatic Connection Device. A type of charging device mentioned under `batteryPackChargingDevice`. |
| **DC / AC** | Direct Current / Alternating Current. Used in relation to battery charging status (`batteryPackChargingStatus`). |
| **DID** | Data Identifier Definition. Referenced as an alternative source for trailer identification (`trailerIdentificationData`). |
| **DTCO** | Digital Tachograph. A type of tachograph (`tachographType`). G1/G2 denote generations. |
| **EBS** | Electronic Braking System. Referenced in tell tales like `EBS` and `EBS_TRAILER_1_2`. |
| **EEC** | Electronic Engine Controller. Implied context for engine torque values (e.g., "EEC1 value", "EEC2 value"). |
| **ESC** | Electronic Stability Control. Referenced in tell tales like `ESC_INDICATOR` and `ESC_SWITCHED_OFF`. |
| **GNSS** | Global Navigation Satellite System. Used for vehicle positioning (`GNSSPositionDto`). |
| **hrTotalVehicleDistance** | High Resolution Total Vehicle Distance. Accumulated distance travelled by the vehicle in meters. |
| **HVESS** | High Voltage Energy Storage System. Refers to the vehicle's main battery pack, often used in temperature fields (`hvessOutletCoolantTemperature`, `hvessTemperature`). |
| **ISO** | International Organization for Standardization. Referenced for various standards (e.g., ISO 11992 for trailer communication, ISO 3779 for VIN, ISO 15118 for charging). |
| **MTCO** | Modular Tachograph. A type of tachograph (`tachographType`). |
| **OEM** | Original Equipment Manufacturer. Refers to the vehicle manufacturer (e.g., Scania, Volvo). Used throughout the spec for non-standardized fields or behaviors. |
| **PTO** | Power Take-Off. An auxiliary device powered by the vehicle's engine/motor. Referenced in triggers (`PTO_ENABLED`, `PTO_DISABLED`) and accumulated data (`ptoActiveClass`). |
| **rFMS** | Remote Fleet Management System. The standard defining this SDK. |
| **SPN** | Suspect Parameter Number. Standardized numbers (SAE J1939) used to identify specific data parameters (e.g., used in `possibleFuelType`, `fuelType`, `selectedGearClass` descriptions). |
| **UTC** | Coordinated Universal Time. Standard time reference, used for timestamps like `requestServerDateTime`. |
| **VIN** | Vehicle Identification Number. A unique 17-character identifier for a vehicle. |
| **WGS84** | World Geodetic System 1984. The coordinate system used for GNSS latitude and longitude. |
| **WPT** | Wireless Power Transfer. A type of charging device mentioned under `batteryPackChargingDevice`.
