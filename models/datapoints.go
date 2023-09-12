package models

import (
	"database/sql"
	"strconv"
)

type DataPoints map[string]map[string]any

type DataPointDescriptor struct {
	deviceId string
	attr     string
	value    string
	_type    string
}

func (config *DataPoints) UpdateConfig(db *sql.DB, deviceId string) error {
	for key, elem := range *config {
		rslt, err := db.Exec("update config set value = ? where attr = ? and device_id = ?", elem["value"], key, deviceId)
		if err == nil {
			aff, _ := rslt.RowsAffected()
			if aff == 0 {
				_, err := db.Exec("insert into config (value, attr, type, device_id) values (?, ?, ?, ?)", elem["value"], elem["type"], key, deviceId)
				if err != nil {
					return err
				}
			}
		} else {
			return err
		}
	}
	return nil
}

func GetDataPoints(db *sql.DB, deviceId *string) (DataPoints, error) {
	rslt, err := db.Query("select device_id, attr, value, type from device, config where device.id = config.device_id and device.id = ?", *deviceId)
	if err != nil {
		return nil, err
	}

	result := make(DataPoints)
	for rslt.Next() {
		row := new(DataPointDescriptor)
		rslt.Scan(&row.deviceId, &row.attr, &row.value, &row._type)
		elem := make(map[string]any)
		if row._type == "bool" {
			bval, _ := strconv.ParseBool(row.value)
			elem["value"] = bval
		} else if row._type == "int32" {
			val, _ := strconv.Atoi(row.value)
			elem["value"] = val
		} else if row._type == "int64" {
			val, _ := strconv.ParseInt(row.value, 10, 64)
			elem["value"] = val
		} else {
			elem["value"] = row.value
		}
		elem["type"] = row._type
		result[row.attr] = elem
	}

	return result, nil
}
